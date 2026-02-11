// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertTemplateResponseBody) *UpdateAlertTemplateResponse
	GetBody() *UpdateAlertTemplateResponseBody
}

type UpdateAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertTemplateResponse) GetBody() *UpdateAlertTemplateResponseBody {
	return s.Body
}

func (s *UpdateAlertTemplateResponse) SetHeaders(v map[string]*string) *UpdateAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertTemplateResponse) SetStatusCode(v int32) *UpdateAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertTemplateResponse) SetBody(v *UpdateAlertTemplateResponseBody) *UpdateAlertTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
