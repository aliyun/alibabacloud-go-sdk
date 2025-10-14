// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVendorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVendorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVendorResponse
	GetStatusCode() *int32
	SetBody(v *CreateVendorResponseBody) *CreateVendorResponse
	GetBody() *CreateVendorResponseBody
}

type CreateVendorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVendorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVendorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVendorResponse) GoString() string {
	return s.String()
}

func (s *CreateVendorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVendorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVendorResponse) GetBody() *CreateVendorResponseBody {
	return s.Body
}

func (s *CreateVendorResponse) SetHeaders(v map[string]*string) *CreateVendorResponse {
	s.Headers = v
	return s
}

func (s *CreateVendorResponse) SetStatusCode(v int32) *CreateVendorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVendorResponse) SetBody(v *CreateVendorResponseBody) *CreateVendorResponse {
	s.Body = v
	return s
}

func (s *CreateVendorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
