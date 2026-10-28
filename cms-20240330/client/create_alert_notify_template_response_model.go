// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertNotifyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertNotifyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertNotifyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertNotifyTemplateResponseBody) *CreateAlertNotifyTemplateResponse
	GetBody() *CreateAlertNotifyTemplateResponseBody
}

type CreateAlertNotifyTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertNotifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertNotifyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertNotifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertNotifyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertNotifyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertNotifyTemplateResponse) GetBody() *CreateAlertNotifyTemplateResponseBody {
	return s.Body
}

func (s *CreateAlertNotifyTemplateResponse) SetHeaders(v map[string]*string) *CreateAlertNotifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertNotifyTemplateResponse) SetStatusCode(v int32) *CreateAlertNotifyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertNotifyTemplateResponse) SetBody(v *CreateAlertNotifyTemplateResponseBody) *CreateAlertNotifyTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAlertNotifyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
