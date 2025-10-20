// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainBrandResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainBrandResponseBody) *UpdateDomainBrandResponse
	GetBody() *UpdateDomainBrandResponseBody
}

type UpdateDomainBrandResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainBrandResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainBrandResponse) GetBody() *UpdateDomainBrandResponseBody {
	return s.Body
}

func (s *UpdateDomainBrandResponse) SetHeaders(v map[string]*string) *UpdateDomainBrandResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainBrandResponse) SetStatusCode(v int32) *UpdateDomainBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainBrandResponse) SetBody(v *UpdateDomainBrandResponseBody) *UpdateDomainBrandResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
