// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDomainMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDomainMetaResponse
	GetStatusCode() *int32
	SetBody(v *CreateDomainMetaResponseBody) *CreateDomainMetaResponse
	GetBody() *CreateDomainMetaResponseBody
}

type CreateDomainMetaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDomainMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDomainMetaResponse) GetBody() *CreateDomainMetaResponseBody {
	return s.Body
}

func (s *CreateDomainMetaResponse) SetHeaders(v map[string]*string) *CreateDomainMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainMetaResponse) SetStatusCode(v int32) *CreateDomainMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainMetaResponse) SetBody(v *CreateDomainMetaResponseBody) *CreateDomainMetaResponse {
	s.Body = v
	return s
}

func (s *CreateDomainMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
