// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTimeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTimeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateTimeTemplateResponseBody) *CreateTimeTemplateResponse
	GetBody() *CreateTimeTemplateResponseBody
}

type CreateTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTimeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTimeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTimeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTimeTemplateResponse) GetBody() *CreateTimeTemplateResponseBody {
	return s.Body
}

func (s *CreateTimeTemplateResponse) SetHeaders(v map[string]*string) *CreateTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTimeTemplateResponse) SetStatusCode(v int32) *CreateTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTimeTemplateResponse) SetBody(v *CreateTimeTemplateResponseBody) *CreateTimeTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateTimeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
