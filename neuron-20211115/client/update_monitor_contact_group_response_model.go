// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContactGroup) *UpdateMonitorContactGroupResponse
	GetBody() *MonitorContactGroup
}

type UpdateMonitorContactGroupResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContactGroup `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorContactGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorContactGroupResponse) GetBody() *MonitorContactGroup {
	return s.Body
}

func (s *UpdateMonitorContactGroupResponse) SetHeaders(v map[string]*string) *UpdateMonitorContactGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorContactGroupResponse) SetStatusCode(v int32) *UpdateMonitorContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorContactGroupResponse) SetBody(v *MonitorContactGroup) *UpdateMonitorContactGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
