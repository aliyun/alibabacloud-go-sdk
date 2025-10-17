// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppAgentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppAgentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppAgentTemplateResponseBody) *DeleteAppAgentTemplateResponse
	GetBody() *DeleteAppAgentTemplateResponseBody
}

type DeleteAppAgentTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppAgentTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppAgentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppAgentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppAgentTemplateResponse) GetBody() *DeleteAppAgentTemplateResponseBody {
	return s.Body
}

func (s *DeleteAppAgentTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppAgentTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppAgentTemplateResponse) SetStatusCode(v int32) *DeleteAppAgentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppAgentTemplateResponse) SetBody(v *DeleteAppAgentTemplateResponseBody) *DeleteAppAgentTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAppAgentTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
