// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationEndpointAddressResponseBody) *ModifyApplicationEndpointAddressResponse
	GetBody() *ModifyApplicationEndpointAddressResponseBody
}

type ModifyApplicationEndpointAddressResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationEndpointAddressResponse) GetBody() *ModifyApplicationEndpointAddressResponseBody {
	return s.Body
}

func (s *ModifyApplicationEndpointAddressResponse) SetHeaders(v map[string]*string) *ModifyApplicationEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationEndpointAddressResponse) SetStatusCode(v int32) *ModifyApplicationEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationEndpointAddressResponse) SetBody(v *ModifyApplicationEndpointAddressResponseBody) *ModifyApplicationEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
