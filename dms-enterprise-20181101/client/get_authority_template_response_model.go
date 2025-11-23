// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorityTemplateResponseBody) *GetAuthorityTemplateResponse
	GetBody() *GetAuthorityTemplateResponseBody
}

type GetAuthorityTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorityTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorityTemplateResponse) GetBody() *GetAuthorityTemplateResponseBody {
	return s.Body
}

func (s *GetAuthorityTemplateResponse) SetHeaders(v map[string]*string) *GetAuthorityTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorityTemplateResponse) SetStatusCode(v int32) *GetAuthorityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorityTemplateResponse) SetBody(v *GetAuthorityTemplateResponseBody) *GetAuthorityTemplateResponse {
	s.Body = v
	return s
}

func (s *GetAuthorityTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
