// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UnbindTemplateResponseBody) *UnbindTemplateResponse
	GetBody() *UnbindTemplateResponseBody
}

type UnbindTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindTemplateResponse) GoString() string {
	return s.String()
}

func (s *UnbindTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindTemplateResponse) GetBody() *UnbindTemplateResponseBody {
	return s.Body
}

func (s *UnbindTemplateResponse) SetHeaders(v map[string]*string) *UnbindTemplateResponse {
	s.Headers = v
	return s
}

func (s *UnbindTemplateResponse) SetStatusCode(v int32) *UnbindTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindTemplateResponse) SetBody(v *UnbindTemplateResponseBody) *UnbindTemplateResponse {
	s.Body = v
	return s
}

func (s *UnbindTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
