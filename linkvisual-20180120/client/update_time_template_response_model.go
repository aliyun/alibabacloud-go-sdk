// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTimeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTimeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTimeTemplateResponseBody) *UpdateTimeTemplateResponse
	GetBody() *UpdateTimeTemplateResponseBody
}

type UpdateTimeTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTimeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTimeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimeTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTimeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTimeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTimeTemplateResponse) GetBody() *UpdateTimeTemplateResponseBody {
	return s.Body
}

func (s *UpdateTimeTemplateResponse) SetHeaders(v map[string]*string) *UpdateTimeTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTimeTemplateResponse) SetStatusCode(v int32) *UpdateTimeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTimeTemplateResponse) SetBody(v *UpdateTimeTemplateResponseBody) *UpdateTimeTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateTimeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
