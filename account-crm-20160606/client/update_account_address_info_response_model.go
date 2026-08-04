// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountAddressInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccountAddressInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccountAddressInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccountAddressInfoResponseBody) *UpdateAccountAddressInfoResponse
	GetBody() *UpdateAccountAddressInfoResponseBody
}

type UpdateAccountAddressInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountAddressInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountAddressInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountAddressInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountAddressInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccountAddressInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccountAddressInfoResponse) GetBody() *UpdateAccountAddressInfoResponseBody {
	return s.Body
}

func (s *UpdateAccountAddressInfoResponse) SetHeaders(v map[string]*string) *UpdateAccountAddressInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountAddressInfoResponse) SetStatusCode(v int32) *UpdateAccountAddressInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountAddressInfoResponse) SetBody(v *UpdateAccountAddressInfoResponseBody) *UpdateAccountAddressInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAccountAddressInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
