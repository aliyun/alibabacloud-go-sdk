// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSystemPropertyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSystemPropertyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSystemPropertyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateSystemPropertyTemplateResponseBody) *CreateSystemPropertyTemplateResponse
	GetBody() *CreateSystemPropertyTemplateResponseBody
}

type CreateSystemPropertyTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSystemPropertyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSystemPropertyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSystemPropertyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSystemPropertyTemplateResponse) GetBody() *CreateSystemPropertyTemplateResponseBody {
	return s.Body
}

func (s *CreateSystemPropertyTemplateResponse) SetHeaders(v map[string]*string) *CreateSystemPropertyTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateSystemPropertyTemplateResponse) SetStatusCode(v int32) *CreateSystemPropertyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSystemPropertyTemplateResponse) SetBody(v *CreateSystemPropertyTemplateResponseBody) *CreateSystemPropertyTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateSystemPropertyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
