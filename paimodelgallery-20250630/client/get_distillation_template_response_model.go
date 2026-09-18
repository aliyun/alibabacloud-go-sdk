// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDistillationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDistillationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDistillationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetDistillationTemplateResponseBody) *GetDistillationTemplateResponse
	GetBody() *GetDistillationTemplateResponseBody
}

type GetDistillationTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDistillationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDistillationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDistillationTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDistillationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDistillationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDistillationTemplateResponse) GetBody() *GetDistillationTemplateResponseBody {
	return s.Body
}

func (s *GetDistillationTemplateResponse) SetHeaders(v map[string]*string) *GetDistillationTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDistillationTemplateResponse) SetStatusCode(v int32) *GetDistillationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDistillationTemplateResponse) SetBody(v *GetDistillationTemplateResponseBody) *GetDistillationTemplateResponse {
	s.Body = v
	return s
}

func (s *GetDistillationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
