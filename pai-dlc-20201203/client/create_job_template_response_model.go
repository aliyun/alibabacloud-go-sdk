// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse
	GetBody() *CreateJobTemplateResponseBody
}

type CreateJobTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobTemplateResponse) GetBody() *CreateJobTemplateResponseBody {
	return s.Body
}

func (s *CreateJobTemplateResponse) SetHeaders(v map[string]*string) *CreateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateJobTemplateResponse) SetStatusCode(v int32) *CreateJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobTemplateResponse) SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
