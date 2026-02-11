// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertTemplateResponseBody) *CreateAlertTemplateResponse
	GetBody() *CreateAlertTemplateResponseBody
}

type CreateAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertTemplateResponse) GetBody() *CreateAlertTemplateResponseBody {
	return s.Body
}

func (s *CreateAlertTemplateResponse) SetHeaders(v map[string]*string) *CreateAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertTemplateResponse) SetStatusCode(v int32) *CreateAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertTemplateResponse) SetBody(v *CreateAlertTemplateResponseBody) *CreateAlertTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAlertTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
