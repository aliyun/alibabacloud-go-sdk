// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHealthCheckUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateHealthCheckUrlResponseBody
	GetCode() *int32
	SetHealthCheckURL(v string) *UpdateHealthCheckUrlResponseBody
	GetHealthCheckURL() *string
	SetMessage(v string) *UpdateHealthCheckUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHealthCheckUrlResponseBody
	GetRequestId() *string
}

type UpdateHealthCheckUrlResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The health check URL.
	//
	// example:
	//
	// http://127.0.0.1:8080/_ehc.html
	HealthCheckURL *string `json:"HealthCheckURL,omitempty" xml:"HealthCheckURL,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-*****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHealthCheckUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHealthCheckUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHealthCheckUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateHealthCheckUrlResponseBody) GetHealthCheckURL() *string {
	return s.HealthCheckURL
}

func (s *UpdateHealthCheckUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHealthCheckUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHealthCheckUrlResponseBody) SetCode(v int32) *UpdateHealthCheckUrlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetHealthCheckURL(v string) *UpdateHealthCheckUrlResponseBody {
	s.HealthCheckURL = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetMessage(v string) *UpdateHealthCheckUrlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) SetRequestId(v string) *UpdateHealthCheckUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHealthCheckUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
