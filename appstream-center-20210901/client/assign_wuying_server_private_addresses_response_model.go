// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignWuyingServerPrivateAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignWuyingServerPrivateAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignWuyingServerPrivateAddressesResponse
	GetStatusCode() *int32
	SetBody(v *AssignWuyingServerPrivateAddressesResponseBody) *AssignWuyingServerPrivateAddressesResponse
	GetBody() *AssignWuyingServerPrivateAddressesResponseBody
}

type AssignWuyingServerPrivateAddressesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignWuyingServerPrivateAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignWuyingServerPrivateAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignWuyingServerPrivateAddressesResponse) GoString() string {
	return s.String()
}

func (s *AssignWuyingServerPrivateAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignWuyingServerPrivateAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignWuyingServerPrivateAddressesResponse) GetBody() *AssignWuyingServerPrivateAddressesResponseBody {
	return s.Body
}

func (s *AssignWuyingServerPrivateAddressesResponse) SetHeaders(v map[string]*string) *AssignWuyingServerPrivateAddressesResponse {
	s.Headers = v
	return s
}

func (s *AssignWuyingServerPrivateAddressesResponse) SetStatusCode(v int32) *AssignWuyingServerPrivateAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignWuyingServerPrivateAddressesResponse) SetBody(v *AssignWuyingServerPrivateAddressesResponseBody) *AssignWuyingServerPrivateAddressesResponse {
	s.Body = v
	return s
}

func (s *AssignWuyingServerPrivateAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
