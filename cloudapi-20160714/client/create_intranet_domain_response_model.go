// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntranetDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntranetDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntranetDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntranetDomainResponseBody) *CreateIntranetDomainResponse
	GetBody() *CreateIntranetDomainResponseBody
}

type CreateIntranetDomainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntranetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntranetDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntranetDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntranetDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntranetDomainResponse) GetBody() *CreateIntranetDomainResponseBody {
	return s.Body
}

func (s *CreateIntranetDomainResponse) SetHeaders(v map[string]*string) *CreateIntranetDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateIntranetDomainResponse) SetStatusCode(v int32) *CreateIntranetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntranetDomainResponse) SetBody(v *CreateIntranetDomainResponseBody) *CreateIntranetDomainResponse {
	s.Body = v
	return s
}

func (s *CreateIntranetDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
