// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCSTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCSTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCSTemplateResponseBody) *CreateRCSTemplateResponse
	GetBody() *CreateRCSTemplateResponseBody
}

type CreateRCSTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCSTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCSTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateRCSTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCSTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCSTemplateResponse) GetBody() *CreateRCSTemplateResponseBody {
	return s.Body
}

func (s *CreateRCSTemplateResponse) SetHeaders(v map[string]*string) *CreateRCSTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateRCSTemplateResponse) SetStatusCode(v int32) *CreateRCSTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCSTemplateResponse) SetBody(v *CreateRCSTemplateResponseBody) *CreateRCSTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateRCSTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
