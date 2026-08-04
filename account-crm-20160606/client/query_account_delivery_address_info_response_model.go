// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountDeliveryAddressInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountDeliveryAddressInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountDeliveryAddressInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountDeliveryAddressInfoResponseBody) *QueryAccountDeliveryAddressInfoResponse
	GetBody() *QueryAccountDeliveryAddressInfoResponseBody
}

type QueryAccountDeliveryAddressInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountDeliveryAddressInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountDeliveryAddressInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountDeliveryAddressInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountDeliveryAddressInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountDeliveryAddressInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountDeliveryAddressInfoResponse) GetBody() *QueryAccountDeliveryAddressInfoResponseBody {
	return s.Body
}

func (s *QueryAccountDeliveryAddressInfoResponse) SetHeaders(v map[string]*string) *QueryAccountDeliveryAddressInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponse) SetStatusCode(v int32) *QueryAccountDeliveryAddressInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponse) SetBody(v *QueryAccountDeliveryAddressInfoResponseBody) *QueryAccountDeliveryAddressInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAccountDeliveryAddressInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
