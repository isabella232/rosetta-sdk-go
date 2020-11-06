// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package client

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// Linger please
var (
	_ _context.Context
)

// AccountAPIService AccountAPI service
type AccountAPIService service

// AccountBalance Get an array of all AccountBalances for an AccountIdentifier and the
// BlockIdentifier at which the balance lookup was performed. The BlockIdentifier must always be
// returned because some consumers of account balance data need to know specifically at which block
// the balance was calculated to compare balances they compute from operations with the balance
// returned by the node. It is important to note that making a balance request for an account
// without populating the SubAccountIdentifier should not result in the balance of all possible
// SubAccountIdentifiers being returned. Rather, it should result in the balance pertaining to no
// SubAccountIdentifiers being returned (sometimes called the liquid balance). To get all balances
// associated with an account, it may be necessary to perform multiple balance requests with unique
// AccountIdentifiers. It is also possible to perform a historical balance lookup (if the server
// supports it) by passing in an optional BlockIdentifier.
func (a *AccountAPIService) AccountBalance(
	ctx _context.Context,
	accountBalanceRequest *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error, error) {
	var (
		localVarPostBody interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/balance"
	localVarHeaderParams := make(map[string]string)

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = accountBalanceRequest

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarPostBody, localVarHeaderParams)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, nil, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	switch localVarHTTPResponse.StatusCode {
	case _nethttp.StatusOK:
		var v types.AccountBalanceResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, err
		}

		return &v, nil, nil
	case _nethttp.StatusInternalServerError:
		var v types.Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, nil, err
		}

		return nil, &v, fmt.Errorf("%+v", v)
	case _nethttp.StatusBadGateway,
		_nethttp.StatusServiceUnavailable,
		_nethttp.StatusGatewayTimeout:
		return nil, nil, fmt.Errorf(
			"%w: code: %d body: %s",
			ErrRetriable,
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	default:
		return nil, nil, fmt.Errorf(
			"invalid status code: %d body: %s",
			localVarHTTPResponse.StatusCode,
			string(localVarBody),
		)
	}
}
