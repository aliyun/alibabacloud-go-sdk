// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignWuyingServerPrivateAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassignWuyingServerPrivateAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassignWuyingServerPrivateAddressesResponse
	GetStatusCode() *int32
	SetBody(v *UnassignWuyingServerPrivateAddressesResponseBody) *UnassignWuyingServerPrivateAddressesResponse
	GetBody() *UnassignWuyingServerPrivateAddressesResponseBody
}

type UnassignWuyingServerPrivateAddressesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignWuyingServerPrivateAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignWuyingServerPrivateAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassignWuyingServerPrivateAddressesResponse) GoString() string {
	return s.String()
}

func (s *UnassignWuyingServerPrivateAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassignWuyingServerPrivateAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassignWuyingServerPrivateAddressesResponse) GetBody() *UnassignWuyingServerPrivateAddressesResponseBody {
	return s.Body
}

func (s *UnassignWuyingServerPrivateAddressesResponse) SetHeaders(v map[string]*string) *UnassignWuyingServerPrivateAddressesResponse {
	s.Headers = v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesResponse) SetStatusCode(v int32) *UnassignWuyingServerPrivateAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesResponse) SetBody(v *UnassignWuyingServerPrivateAddressesResponseBody) *UnassignWuyingServerPrivateAddressesResponse {
	s.Body = v
	return s
}

func (s *UnassignWuyingServerPrivateAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
