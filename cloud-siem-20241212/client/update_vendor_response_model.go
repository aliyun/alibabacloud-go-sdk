// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVendorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVendorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVendorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVendorResponseBody) *UpdateVendorResponse
	GetBody() *UpdateVendorResponseBody
}

type UpdateVendorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVendorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVendorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVendorResponse) GoString() string {
	return s.String()
}

func (s *UpdateVendorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVendorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVendorResponse) GetBody() *UpdateVendorResponseBody {
	return s.Body
}

func (s *UpdateVendorResponse) SetHeaders(v map[string]*string) *UpdateVendorResponse {
	s.Headers = v
	return s
}

func (s *UpdateVendorResponse) SetStatusCode(v int32) *UpdateVendorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVendorResponse) SetBody(v *UpdateVendorResponseBody) *UpdateVendorResponse {
	s.Body = v
	return s
}

func (s *UpdateVendorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
