// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAITemplateResponseBody) *DeleteAITemplateResponse
	GetBody() *DeleteAITemplateResponseBody
}

type DeleteAITemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAITemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAITemplateResponse) GetBody() *DeleteAITemplateResponseBody {
	return s.Body
}

func (s *DeleteAITemplateResponse) SetHeaders(v map[string]*string) *DeleteAITemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAITemplateResponse) SetStatusCode(v int32) *DeleteAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAITemplateResponse) SetBody(v *DeleteAITemplateResponseBody) *DeleteAITemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAITemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
