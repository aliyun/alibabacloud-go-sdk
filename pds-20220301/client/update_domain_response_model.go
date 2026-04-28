// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainResponse
	GetStatusCode() *int32
	SetBody(v *Domain) *UpdateDomainResponse
	GetBody() *Domain
}

type UpdateDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Domain            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainResponse) GetBody() *Domain {
	return s.Body
}

func (s *UpdateDomainResponse) SetHeaders(v map[string]*string) *UpdateDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainResponse) SetStatusCode(v int32) *UpdateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainResponse) SetBody(v *Domain) *UpdateDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
