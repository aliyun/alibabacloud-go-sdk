// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultDomainResponse
	GetStatusCode() *int32
	SetBody(v *GetDefaultDomainResponseBody) *GetDefaultDomainResponse
	GetBody() *GetDefaultDomainResponseBody
}

type GetDefaultDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultDomainResponse) GetBody() *GetDefaultDomainResponseBody {
	return s.Body
}

func (s *GetDefaultDomainResponse) SetHeaders(v map[string]*string) *GetDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultDomainResponse) SetStatusCode(v int32) *GetDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultDomainResponse) SetBody(v *GetDefaultDomainResponseBody) *GetDefaultDomainResponse {
	s.Body = v
	return s
}

func (s *GetDefaultDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
