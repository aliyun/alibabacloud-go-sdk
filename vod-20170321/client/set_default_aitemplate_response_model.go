// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultAITemplateResponseBody) *SetDefaultAITemplateResponse
	GetBody() *SetDefaultAITemplateResponseBody
}

type SetDefaultAITemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultAITemplateResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultAITemplateResponse) GetBody() *SetDefaultAITemplateResponseBody {
	return s.Body
}

func (s *SetDefaultAITemplateResponse) SetHeaders(v map[string]*string) *SetDefaultAITemplateResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultAITemplateResponse) SetStatusCode(v int32) *SetDefaultAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultAITemplateResponse) SetBody(v *SetDefaultAITemplateResponseBody) *SetDefaultAITemplateResponse {
	s.Body = v
	return s
}

func (s *SetDefaultAITemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
