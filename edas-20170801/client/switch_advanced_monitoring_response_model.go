// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchAdvancedMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchAdvancedMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchAdvancedMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *SwitchAdvancedMonitoringResponseBody) *SwitchAdvancedMonitoringResponse
	GetBody() *SwitchAdvancedMonitoringResponseBody
}

type SwitchAdvancedMonitoringResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchAdvancedMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchAdvancedMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchAdvancedMonitoringResponse) GoString() string {
	return s.String()
}

func (s *SwitchAdvancedMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchAdvancedMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchAdvancedMonitoringResponse) GetBody() *SwitchAdvancedMonitoringResponseBody {
	return s.Body
}

func (s *SwitchAdvancedMonitoringResponse) SetHeaders(v map[string]*string) *SwitchAdvancedMonitoringResponse {
	s.Headers = v
	return s
}

func (s *SwitchAdvancedMonitoringResponse) SetStatusCode(v int32) *SwitchAdvancedMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchAdvancedMonitoringResponse) SetBody(v *SwitchAdvancedMonitoringResponseBody) *SwitchAdvancedMonitoringResponse {
	s.Body = v
	return s
}

func (s *SwitchAdvancedMonitoringResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
