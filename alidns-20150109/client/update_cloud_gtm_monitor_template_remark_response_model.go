// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmMonitorTemplateRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmMonitorTemplateRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmMonitorTemplateRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmMonitorTemplateRemarkResponseBody) *UpdateCloudGtmMonitorTemplateRemarkResponse
	GetBody() *UpdateCloudGtmMonitorTemplateRemarkResponseBody
}

type UpdateCloudGtmMonitorTemplateRemarkResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmMonitorTemplateRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmMonitorTemplateRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmMonitorTemplateRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) GetBody() *UpdateCloudGtmMonitorTemplateRemarkResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmMonitorTemplateRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) SetStatusCode(v int32) *UpdateCloudGtmMonitorTemplateRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) SetBody(v *UpdateCloudGtmMonitorTemplateRemarkResponseBody) *UpdateCloudGtmMonitorTemplateRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmMonitorTemplateRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
