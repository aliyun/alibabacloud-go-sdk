// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationTemplateResponseBody) *CreateApplicationTemplateResponse
	GetBody() *CreateApplicationTemplateResponseBody
}

type CreateApplicationTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationTemplateResponse) GetBody() *CreateApplicationTemplateResponseBody {
	return s.Body
}

func (s *CreateApplicationTemplateResponse) SetHeaders(v map[string]*string) *CreateApplicationTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationTemplateResponse) SetStatusCode(v int32) *CreateApplicationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationTemplateResponse) SetBody(v *CreateApplicationTemplateResponseBody) *CreateApplicationTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
