// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContactGroupCreateCmd) *CreateMonitorContactGroupResponse
	GetBody() *MonitorContactGroupCreateCmd
}

type CreateMonitorContactGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContactGroupCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorContactGroupResponse) GetBody() *MonitorContactGroupCreateCmd {
	return s.Body
}

func (s *CreateMonitorContactGroupResponse) SetHeaders(v map[string]*string) *CreateMonitorContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorContactGroupResponse) SetStatusCode(v int32) *CreateMonitorContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorContactGroupResponse) SetBody(v *MonitorContactGroupCreateCmd) *CreateMonitorContactGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
