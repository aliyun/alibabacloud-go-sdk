// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppViewTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppViewTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppViewTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppViewTemplateResponseBody) *CreateAppViewTemplateResponse
	GetBody() *CreateAppViewTemplateResponseBody
}

type CreateAppViewTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppViewTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppViewTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppViewTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppViewTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppViewTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppViewTemplateResponse) GetBody() *CreateAppViewTemplateResponseBody {
	return s.Body
}

func (s *CreateAppViewTemplateResponse) SetHeaders(v map[string]*string) *CreateAppViewTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppViewTemplateResponse) SetStatusCode(v int32) *CreateAppViewTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppViewTemplateResponse) SetBody(v *CreateAppViewTemplateResponseBody) *CreateAppViewTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAppViewTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
