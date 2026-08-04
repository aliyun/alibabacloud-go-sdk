// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgAccountAddressInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgAccountAddressInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgAccountAddressInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgAccountAddressInfoResponseBody) *UpdateAgAccountAddressInfoResponse
	GetBody() *UpdateAgAccountAddressInfoResponseBody
}

type UpdateAgAccountAddressInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgAccountAddressInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgAccountAddressInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgAccountAddressInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgAccountAddressInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgAccountAddressInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgAccountAddressInfoResponse) GetBody() *UpdateAgAccountAddressInfoResponseBody {
	return s.Body
}

func (s *UpdateAgAccountAddressInfoResponse) SetHeaders(v map[string]*string) *UpdateAgAccountAddressInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgAccountAddressInfoResponse) SetStatusCode(v int32) *UpdateAgAccountAddressInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgAccountAddressInfoResponse) SetBody(v *UpdateAgAccountAddressInfoResponseBody) *UpdateAgAccountAddressInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateAgAccountAddressInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
