// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigTemplateResponseBody) *CreateConfigTemplateResponse
	GetBody() *CreateConfigTemplateResponseBody
}

type CreateConfigTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigTemplateResponse) GetBody() *CreateConfigTemplateResponseBody {
	return s.Body
}

func (s *CreateConfigTemplateResponse) SetHeaders(v map[string]*string) *CreateConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigTemplateResponse) SetStatusCode(v int32) *CreateConfigTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigTemplateResponse) SetBody(v *CreateConfigTemplateResponseBody) *CreateConfigTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateConfigTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
