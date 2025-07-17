// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePromptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePromptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePromptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeletePromptTemplateResponseBody) *DeletePromptTemplateResponse
	GetBody() *DeletePromptTemplateResponseBody
}

type DeletePromptTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePromptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePromptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePromptTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeletePromptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePromptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePromptTemplateResponse) GetBody() *DeletePromptTemplateResponseBody {
	return s.Body
}

func (s *DeletePromptTemplateResponse) SetHeaders(v map[string]*string) *DeletePromptTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeletePromptTemplateResponse) SetStatusCode(v int32) *DeletePromptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePromptTemplateResponse) SetBody(v *DeletePromptTemplateResponseBody) *DeletePromptTemplateResponse {
	s.Body = v
	return s
}

func (s *DeletePromptTemplateResponse) Validate() error {
	return dara.Validate(s)
}
