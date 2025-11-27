// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPurchasedDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPurchasedDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPurchasedDeviceResponse
	GetStatusCode() *int32
	SetBody(v *BindPurchasedDeviceResponseBody) *BindPurchasedDeviceResponse
	GetBody() *BindPurchasedDeviceResponseBody
}

type BindPurchasedDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPurchasedDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindPurchasedDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPurchasedDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPurchasedDeviceResponse) GetBody() *BindPurchasedDeviceResponseBody {
	return s.Body
}

func (s *BindPurchasedDeviceResponse) SetHeaders(v map[string]*string) *BindPurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindPurchasedDeviceResponse) SetStatusCode(v int32) *BindPurchasedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPurchasedDeviceResponse) SetBody(v *BindPurchasedDeviceResponseBody) *BindPurchasedDeviceResponse {
	s.Body = v
	return s
}

func (s *BindPurchasedDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
