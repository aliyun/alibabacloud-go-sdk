// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetJobTemplateResponseBody) *GetJobTemplateResponse
	GetBody() *GetJobTemplateResponseBody
}

type GetJobTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetJobTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobTemplateResponse) GetBody() *GetJobTemplateResponseBody {
	return s.Body
}

func (s *GetJobTemplateResponse) SetHeaders(v map[string]*string) *GetJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetJobTemplateResponse) SetStatusCode(v int32) *GetJobTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobTemplateResponse) SetBody(v *GetJobTemplateResponseBody) *GetJobTemplateResponse {
	s.Body = v
	return s
}

func (s *GetJobTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
