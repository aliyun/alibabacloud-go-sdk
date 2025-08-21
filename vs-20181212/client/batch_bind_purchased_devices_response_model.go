// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindPurchasedDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchBindPurchasedDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchBindPurchasedDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchBindPurchasedDevicesResponseBody) *BatchBindPurchasedDevicesResponse
	GetBody() *BatchBindPurchasedDevicesResponseBody
}

type BatchBindPurchasedDevicesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchBindPurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchBindPurchasedDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchBindPurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchBindPurchasedDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchBindPurchasedDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchBindPurchasedDevicesResponse) GetBody() *BatchBindPurchasedDevicesResponseBody {
	return s.Body
}

func (s *BatchBindPurchasedDevicesResponse) SetHeaders(v map[string]*string) *BatchBindPurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchBindPurchasedDevicesResponse) SetStatusCode(v int32) *BatchBindPurchasedDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchBindPurchasedDevicesResponse) SetBody(v *BatchBindPurchasedDevicesResponseBody) *BatchBindPurchasedDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchBindPurchasedDevicesResponse) Validate() error {
	return dara.Validate(s)
}
