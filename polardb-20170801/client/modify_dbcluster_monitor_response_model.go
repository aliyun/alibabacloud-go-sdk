// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterMonitorResponseBody) *ModifyDBClusterMonitorResponse
	GetBody() *ModifyDBClusterMonitorResponseBody
}

type ModifyDBClusterMonitorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterMonitorResponse) GetBody() *ModifyDBClusterMonitorResponseBody {
	return s.Body
}

func (s *ModifyDBClusterMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMonitorResponse) SetStatusCode(v int32) *ModifyDBClusterMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMonitorResponse) SetBody(v *ModifyDBClusterMonitorResponseBody) *ModifyDBClusterMonitorResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterMonitorResponse) Validate() error {
	return dara.Validate(s)
}
