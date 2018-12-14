// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package virtualization

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

	models "github.com/tuenti/go-netbox/netbox/models"
)

// NewVirtualizationClustersCreateParams creates a new VirtualizationClustersCreateParams object
// with the default values initialized.
func NewVirtualizationClustersCreateParams() *VirtualizationClustersCreateParams {
	var ()
	return &VirtualizationClustersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersCreateParamsWithTimeout creates a new VirtualizationClustersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClustersCreateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersCreateParams {
	var ()
	return &VirtualizationClustersCreateParams{

		timeout: timeout,
	}
}

// NewVirtualizationClustersCreateParamsWithContext creates a new VirtualizationClustersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClustersCreateParamsWithContext(ctx context.Context) *VirtualizationClustersCreateParams {
	var ()
	return &VirtualizationClustersCreateParams{

		Context: ctx,
	}
}

// NewVirtualizationClustersCreateParamsWithHTTPClient creates a new VirtualizationClustersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClustersCreateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersCreateParams {
	var ()
	return &VirtualizationClustersCreateParams{
		HTTPClient: client,
	}
}

/*VirtualizationClustersCreateParams contains all the parameters to send to the API endpoint
for the virtualization clusters create operation typically these are written to a http.Request
*/
type VirtualizationClustersCreateParams struct {

	/*Data*/
	Data *models.WritableCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithContext(ctx context.Context) *VirtualizationClustersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithData(data *models.WritableCluster) *VirtualizationClustersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
