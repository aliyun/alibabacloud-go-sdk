// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmMonitorTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmMonitorTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmMonitorTemplateResponseBody) *UpdateCloudGtmMonitorTemplateResponse
	GetBody() *UpdateCloudGtmMonitorTemplateResponseBody
}

type UpdateCloudGtmMonitorTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmMonitorTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmMonitorTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmMonitorTemplateResponse) GetBody() *UpdateCloudGtmMonitorTemplateResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmMonitorTemplateResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmMonitorTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateResponse) SetStatusCode(v int32) *UpdateCloudGtmMonitorTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateResponse) SetBody(v *UpdateCloudGtmMonitorTemplateResponseBody) *UpdateCloudGtmMonitorTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
