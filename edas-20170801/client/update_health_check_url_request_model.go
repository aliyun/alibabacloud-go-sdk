// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateHealthCheckUrlRequest
	GetAppId() *string
	SetHcURL(v string) *UpdateHealthCheckUrlRequest
	GetHcURL() *string
}

type UpdateHealthCheckUrlRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d-43ff-***************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The health check URL of the application. The URL must start with `http://`, and can be up to 255 characters in length. Example: `http://127.0.0.1:8080/_ehc.html`. If this parameter is not specified, the health check URL of the application is not changed.
	//
	// example:
	//
	// http://127.0.0.1:8080/_ehc.html
	HcURL *string `json:"hcURL,omitempty" xml:"hcURL,omitempty"`
}

func (s UpdateHealthCheckUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateHealthCheckUrlRequest) GetHcURL() *string {
	return s.HcURL
}

func (s *UpdateHealthCheckUrlRequest) SetAppId(v string) *UpdateHealthCheckUrlRequest {
	s.AppId = &v
	return s
}

func (s *UpdateHealthCheckUrlRequest) SetHcURL(v string) *UpdateHealthCheckUrlRequest {
	s.HcURL = &v
	return s
}

func (s *UpdateHealthCheckUrlRequest) Validate() error {
	return dara.Validate(s)
}
