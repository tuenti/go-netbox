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

package dcim

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

// NewDcimDevicesCreateParams creates a new DcimDevicesCreateParams object
// with the default values initialized.
func NewDcimDevicesCreateParams() *DcimDevicesCreateParams {
	var ()
	return &DcimDevicesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesCreateParamsWithTimeout creates a new DcimDevicesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDevicesCreateParamsWithTimeout(timeout time.Duration) *DcimDevicesCreateParams {
	var ()
	return &DcimDevicesCreateParams{

		timeout: timeout,
	}
}

// NewDcimDevicesCreateParamsWithContext creates a new DcimDevicesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDevicesCreateParamsWithContext(ctx context.Context) *DcimDevicesCreateParams {
	var ()
	return &DcimDevicesCreateParams{

		Context: ctx,
	}
}

// NewDcimDevicesCreateParamsWithHTTPClient creates a new DcimDevicesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDevicesCreateParamsWithHTTPClient(client *http.Client) *DcimDevicesCreateParams {
	var ()
	return &DcimDevicesCreateParams{
		HTTPClient: client,
	}
}

/*DcimDevicesCreateParams contains all the parameters to send to the API endpoint
for the dcim devices create operation typically these are written to a http.Request
*/
type DcimDevicesCreateParams struct {

	/*Data*/
	Data *models.WritableDevice

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim devices create params
func (o *DcimDevicesCreateParams) WithTimeout(timeout time.Duration) *DcimDevicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices create params
func (o *DcimDevicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices create params
func (o *DcimDevicesCreateParams) WithContext(ctx context.Context) *DcimDevicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices create params
func (o *DcimDevicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices create params
func (o *DcimDevicesCreateParams) WithHTTPClient(client *http.Client) *DcimDevicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices create params
func (o *DcimDevicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim devices create params
func (o *DcimDevicesCreateParams) WithData(data *models.WritableDevice) *DcimDevicesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim devices create params
func (o *DcimDevicesCreateParams) SetData(data *models.WritableDevice) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
