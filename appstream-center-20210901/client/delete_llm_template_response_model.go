// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLlmTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLlmTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLlmTemplateResponseBody) *DeleteLlmTemplateResponse
	GetBody() *DeleteLlmTemplateResponseBody
}

type DeleteLlmTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLlmTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLlmTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteLlmTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLlmTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLlmTemplateResponse) GetBody() *DeleteLlmTemplateResponseBody {
	return s.Body
}

func (s *DeleteLlmTemplateResponse) SetHeaders(v map[string]*string) *DeleteLlmTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteLlmTemplateResponse) SetStatusCode(v int32) *DeleteLlmTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLlmTemplateResponse) SetBody(v *DeleteLlmTemplateResponseBody) *DeleteLlmTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteLlmTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
