// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPurchasedDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindPurchasedDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindPurchasedDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnbindPurchasedDeviceResponseBody) *UnbindPurchasedDeviceResponse
	GetBody() *UnbindPurchasedDeviceResponseBody
}

type UnbindPurchasedDeviceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindPurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindPurchasedDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindPurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindPurchasedDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindPurchasedDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindPurchasedDeviceResponse) GetBody() *UnbindPurchasedDeviceResponseBody {
	return s.Body
}

func (s *UnbindPurchasedDeviceResponse) SetHeaders(v map[string]*string) *UnbindPurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindPurchasedDeviceResponse) SetStatusCode(v int32) *UnbindPurchasedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindPurchasedDeviceResponse) SetBody(v *UnbindPurchasedDeviceResponseBody) *UnbindPurchasedDeviceResponse {
	s.Body = v
	return s
}

func (s *UnbindPurchasedDeviceResponse) Validate() error {
	return dara.Validate(s)
}
