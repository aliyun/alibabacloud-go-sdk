// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardTemplateResponseBody) *CreateStandardTemplateResponse
	GetBody() *CreateStandardTemplateResponseBody
}

type CreateStandardTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardTemplateResponse) GetBody() *CreateStandardTemplateResponseBody {
	return s.Body
}

func (s *CreateStandardTemplateResponse) SetHeaders(v map[string]*string) *CreateStandardTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardTemplateResponse) SetStatusCode(v int32) *CreateStandardTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardTemplateResponse) SetBody(v *CreateStandardTemplateResponseBody) *CreateStandardTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateStandardTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
