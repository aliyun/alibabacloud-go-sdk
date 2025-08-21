// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindPurchasedDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUnbindPurchasedDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUnbindPurchasedDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchUnbindPurchasedDevicesResponseBody) *BatchUnbindPurchasedDevicesResponse
	GetBody() *BatchUnbindPurchasedDevicesResponseBody
}

type BatchUnbindPurchasedDevicesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUnbindPurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUnbindPurchasedDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindPurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchUnbindPurchasedDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUnbindPurchasedDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUnbindPurchasedDevicesResponse) GetBody() *BatchUnbindPurchasedDevicesResponseBody {
	return s.Body
}

func (s *BatchUnbindPurchasedDevicesResponse) SetHeaders(v map[string]*string) *BatchUnbindPurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponse) SetStatusCode(v int32) *BatchUnbindPurchasedDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponse) SetBody(v *BatchUnbindPurchasedDevicesResponseBody) *BatchUnbindPurchasedDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchUnbindPurchasedDevicesResponse) Validate() error {
	return dara.Validate(s)
}
