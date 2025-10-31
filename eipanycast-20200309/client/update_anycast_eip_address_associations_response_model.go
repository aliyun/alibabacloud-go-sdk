// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAnycastEipAddressAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAnycastEipAddressAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAnycastEipAddressAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAnycastEipAddressAssociationsResponseBody) *UpdateAnycastEipAddressAssociationsResponse
	GetBody() *UpdateAnycastEipAddressAssociationsResponseBody
}

type UpdateAnycastEipAddressAssociationsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAnycastEipAddressAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAnycastEipAddressAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAnycastEipAddressAssociationsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAnycastEipAddressAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAnycastEipAddressAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAnycastEipAddressAssociationsResponse) GetBody() *UpdateAnycastEipAddressAssociationsResponseBody {
	return s.Body
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetHeaders(v map[string]*string) *UpdateAnycastEipAddressAssociationsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetStatusCode(v int32) *UpdateAnycastEipAddressAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponse) SetBody(v *UpdateAnycastEipAddressAssociationsResponseBody) *UpdateAnycastEipAddressAssociationsResponse {
	s.Body = v
	return s
}

func (s *UpdateAnycastEipAddressAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
