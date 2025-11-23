// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantTemplateAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantTemplateAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantTemplateAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *GrantTemplateAuthorityResponseBody) *GrantTemplateAuthorityResponse
	GetBody() *GrantTemplateAuthorityResponseBody
}

type GrantTemplateAuthorityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantTemplateAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantTemplateAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantTemplateAuthorityResponse) GoString() string {
	return s.String()
}

func (s *GrantTemplateAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantTemplateAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantTemplateAuthorityResponse) GetBody() *GrantTemplateAuthorityResponseBody {
	return s.Body
}

func (s *GrantTemplateAuthorityResponse) SetHeaders(v map[string]*string) *GrantTemplateAuthorityResponse {
	s.Headers = v
	return s
}

func (s *GrantTemplateAuthorityResponse) SetStatusCode(v int32) *GrantTemplateAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantTemplateAuthorityResponse) SetBody(v *GrantTemplateAuthorityResponseBody) *GrantTemplateAuthorityResponse {
	s.Body = v
	return s
}

func (s *GrantTemplateAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
