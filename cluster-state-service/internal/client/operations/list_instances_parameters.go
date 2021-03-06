// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListInstancesParams creates a new ListInstancesParams object
// with the default values initialized.
func NewListInstancesParams() *ListInstancesParams {
	var ()
	return &ListInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListInstancesParamsWithTimeout creates a new ListInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListInstancesParamsWithTimeout(timeout time.Duration) *ListInstancesParams {
	var ()
	return &ListInstancesParams{

		timeout: timeout,
	}
}

// NewListInstancesParamsWithContext creates a new ListInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListInstancesParamsWithContext(ctx context.Context) *ListInstancesParams {
	var ()
	return &ListInstancesParams{

		Context: ctx,
	}
}

/*ListInstancesParams contains all the parameters to send to the API endpoint
for the list instances operation typically these are written to a http.Request
*/
type ListInstancesParams struct {

	/*Cluster
	  Cluster name or ARN to filter instances by

	*/
	Cluster *string
	/*Status
	  Status to filter instances by

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list instances params
func (o *ListInstancesParams) WithTimeout(timeout time.Duration) *ListInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list instances params
func (o *ListInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list instances params
func (o *ListInstancesParams) WithContext(ctx context.Context) *ListInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list instances params
func (o *ListInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCluster adds the cluster to the list instances params
func (o *ListInstancesParams) WithCluster(cluster *string) *ListInstancesParams {
	o.SetCluster(cluster)
	return o
}

// SetCluster adds the cluster to the list instances params
func (o *ListInstancesParams) SetCluster(cluster *string) {
	o.Cluster = cluster
}

// WithStatus adds the status to the list instances params
func (o *ListInstancesParams) WithStatus(status *string) *ListInstancesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list instances params
func (o *ListInstancesParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Cluster != nil {

		// query param cluster
		var qrCluster string
		if o.Cluster != nil {
			qrCluster = *o.Cluster
		}
		qCluster := qrCluster
		if qCluster != "" {
			if err := r.SetQueryParam("cluster", qCluster); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
