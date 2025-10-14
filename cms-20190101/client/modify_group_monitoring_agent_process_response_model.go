// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGroupMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGroupMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGroupMonitoringAgentProcessResponseBody) *ModifyGroupMonitoringAgentProcessResponse
	GetBody() *ModifyGroupMonitoringAgentProcessResponseBody
}

type ModifyGroupMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGroupMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGroupMonitoringAgentProcessResponse) GetBody() *ModifyGroupMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *ModifyGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *ModifyGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponse) SetStatusCode(v int32) *ModifyGroupMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponse) SetBody(v *ModifyGroupMonitoringAgentProcessResponseBody) *ModifyGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
