// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTemplateAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeTemplateAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeTemplateAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *RevokeTemplateAuthorityResponseBody) *RevokeTemplateAuthorityResponse
	GetBody() *RevokeTemplateAuthorityResponseBody
}

type RevokeTemplateAuthorityResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeTemplateAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTemplateAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeTemplateAuthorityResponse) GoString() string {
	return s.String()
}

func (s *RevokeTemplateAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeTemplateAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeTemplateAuthorityResponse) GetBody() *RevokeTemplateAuthorityResponseBody {
	return s.Body
}

func (s *RevokeTemplateAuthorityResponse) SetHeaders(v map[string]*string) *RevokeTemplateAuthorityResponse {
	s.Headers = v
	return s
}

func (s *RevokeTemplateAuthorityResponse) SetStatusCode(v int32) *RevokeTemplateAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTemplateAuthorityResponse) SetBody(v *RevokeTemplateAuthorityResponseBody) *RevokeTemplateAuthorityResponse {
	s.Body = v
	return s
}

func (s *RevokeTemplateAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
