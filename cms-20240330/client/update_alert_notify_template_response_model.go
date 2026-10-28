// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertNotifyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertNotifyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertNotifyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertNotifyTemplateResponseBody) *UpdateAlertNotifyTemplateResponse
	GetBody() *UpdateAlertNotifyTemplateResponseBody
}

type UpdateAlertNotifyTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertNotifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertNotifyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertNotifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertNotifyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertNotifyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertNotifyTemplateResponse) GetBody() *UpdateAlertNotifyTemplateResponseBody {
	return s.Body
}

func (s *UpdateAlertNotifyTemplateResponse) SetHeaders(v map[string]*string) *UpdateAlertNotifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertNotifyTemplateResponse) SetStatusCode(v int32) *UpdateAlertNotifyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertNotifyTemplateResponse) SetBody(v *UpdateAlertNotifyTemplateResponseBody) *UpdateAlertNotifyTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertNotifyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
