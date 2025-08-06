// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDefaultAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDefaultAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDefaultAITemplateResponseBody) *GetDefaultAITemplateResponse
	GetBody() *GetDefaultAITemplateResponseBody
}

type GetDefaultAITemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultAITemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDefaultAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDefaultAITemplateResponse) GetBody() *GetDefaultAITemplateResponseBody {
	return s.Body
}

func (s *GetDefaultAITemplateResponse) SetHeaders(v map[string]*string) *GetDefaultAITemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultAITemplateResponse) SetStatusCode(v int32) *GetDefaultAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultAITemplateResponse) SetBody(v *GetDefaultAITemplateResponseBody) *GetDefaultAITemplateResponse {
	s.Body = v
	return s
}

func (s *GetDefaultAITemplateResponse) Validate() error {
	return dara.Validate(s)
}
