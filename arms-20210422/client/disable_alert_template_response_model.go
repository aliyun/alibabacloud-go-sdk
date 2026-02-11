// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlertTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAlertTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAlertTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DisableAlertTemplateResponseBody) *DisableAlertTemplateResponse
	GetBody() *DisableAlertTemplateResponseBody
}

type DisableAlertTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAlertTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAlertTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAlertTemplateResponse) GetBody() *DisableAlertTemplateResponseBody {
	return s.Body
}

func (s *DisableAlertTemplateResponse) SetHeaders(v map[string]*string) *DisableAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *DisableAlertTemplateResponse) SetStatusCode(v int32) *DisableAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAlertTemplateResponse) SetBody(v *DisableAlertTemplateResponseBody) *DisableAlertTemplateResponse {
	s.Body = v
	return s
}

func (s *DisableAlertTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
