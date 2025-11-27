// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceMonitorResponseBody) *ModifyDBInstanceMonitorResponse
	GetBody() *ModifyDBInstanceMonitorResponseBody
}

type ModifyDBInstanceMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceMonitorResponse) GetBody() *ModifyDBInstanceMonitorResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceMonitorResponse) SetStatusCode(v int32) *ModifyDBInstanceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceMonitorResponse) SetBody(v *ModifyDBInstanceMonitorResponseBody) *ModifyDBInstanceMonitorResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
