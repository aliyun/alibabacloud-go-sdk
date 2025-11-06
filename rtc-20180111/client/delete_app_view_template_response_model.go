// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppViewTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppViewTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppViewTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppViewTemplateResponseBody) *DeleteAppViewTemplateResponse
	GetBody() *DeleteAppViewTemplateResponseBody
}

type DeleteAppViewTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppViewTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppViewTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppViewTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppViewTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppViewTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppViewTemplateResponse) GetBody() *DeleteAppViewTemplateResponseBody {
	return s.Body
}

func (s *DeleteAppViewTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppViewTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppViewTemplateResponse) SetStatusCode(v int32) *DeleteAppViewTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppViewTemplateResponse) SetBody(v *DeleteAppViewTemplateResponseBody) *DeleteAppViewTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAppViewTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
