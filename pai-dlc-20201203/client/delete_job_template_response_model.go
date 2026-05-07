// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobTemplateResponseBody) *DeleteJobTemplateResponse
	GetBody() *DeleteJobTemplateResponseBody
}

type DeleteJobTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobTemplateResponse) GetBody() *DeleteJobTemplateResponseBody {
	return s.Body
}

func (s *DeleteJobTemplateResponse) SetHeaders(v map[string]*string) *DeleteJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobTemplateResponse) SetStatusCode(v int32) *DeleteJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobTemplateResponse) SetBody(v *DeleteJobTemplateResponseBody) *DeleteJobTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
