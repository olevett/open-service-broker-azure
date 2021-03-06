package blueprint

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssignmentOperationsClient is the blueprint Client
type AssignmentOperationsClient struct {
	BaseClient
}

// NewAssignmentOperationsClient creates an instance of the AssignmentOperationsClient client.
func NewAssignmentOperationsClient() AssignmentOperationsClient {
	return NewAssignmentOperationsClientWithBaseURI(DefaultBaseURI)
}

// NewAssignmentOperationsClientWithBaseURI creates an instance of the AssignmentOperationsClient client.
func NewAssignmentOperationsClientWithBaseURI(baseURI string) AssignmentOperationsClient {
	return AssignmentOperationsClient{NewWithBaseURI(baseURI)}
}

// Get get a Blueprint assignment operation.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the assignment.
// assignmentOperationName - name of the assignment operation.
func (client AssignmentOperationsClient) Get(ctx context.Context, scope string, assignmentName string, assignmentOperationName string) (result AssignmentOperation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentOperationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, assignmentName, assignmentOperationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentOperationsClient) GetPreparer(ctx context.Context, scope string, assignmentName string, assignmentOperationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName":          autorest.Encode("path", assignmentName),
		"assignmentOperationName": autorest.Encode("path", assignmentOperationName),
		"scope":                   scope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations/{assignmentOperationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentOperationsClient) GetResponder(resp *http.Response) (result AssignmentOperation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list Operations for given blueprint assignment within a subscription.
// Parameters:
// scope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the assignment.
func (client AssignmentOperationsClient) List(ctx context.Context, scope string, assignmentName string) (result AssignmentOperationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentOperationsClient.List")
		defer func() {
			sc := -1
			if result.aol.Response.Response != nil {
				sc = result.aol.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, assignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aol.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.aol, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentOperationsClient) ListPreparer(ctx context.Context, scope string, assignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName": autorest.Encode("path", assignmentName),
		"scope":          scope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentOperationsClient) ListResponder(resp *http.Response) (result AssignmentOperationList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AssignmentOperationsClient) listNextResults(ctx context.Context, lastResults AssignmentOperationList) (result AssignmentOperationList, err error) {
	req, err := lastResults.assignmentOperationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssignmentOperationsClient) ListComplete(ctx context.Context, scope string, assignmentName string) (result AssignmentOperationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentOperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, assignmentName)
	return
}
