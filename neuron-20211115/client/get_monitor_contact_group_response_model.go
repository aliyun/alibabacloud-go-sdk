// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonitorContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonitorContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContactGroup) *GetMonitorContactGroupResponse
	GetBody() *MonitorContactGroup
}

type GetMonitorContactGroupResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContactGroup `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorContactGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonitorContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonitorContactGroupResponse) GetBody() *MonitorContactGroup {
	return s.Body
}

func (s *GetMonitorContactGroupResponse) SetHeaders(v map[string]*string) *GetMonitorContactGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorContactGroupResponse) SetStatusCode(v int32) *GetMonitorContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorContactGroupResponse) SetBody(v *MonitorContactGroup) *GetMonitorContactGroupResponse {
	s.Body = v
	return s
}

func (s *GetMonitorContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
