// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetAITemplateResponseBody) *GetAITemplateResponse
	GetBody() *GetAITemplateResponseBody
}

type GetAITemplateResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAITemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAITemplateResponse) GetBody() *GetAITemplateResponseBody {
	return s.Body
}

func (s *GetAITemplateResponse) SetHeaders(v map[string]*string) *GetAITemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAITemplateResponse) SetStatusCode(v int32) *GetAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITemplateResponse) SetBody(v *GetAITemplateResponseBody) *GetAITemplateResponse {
	s.Body = v
	return s
}

func (s *GetAITemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
