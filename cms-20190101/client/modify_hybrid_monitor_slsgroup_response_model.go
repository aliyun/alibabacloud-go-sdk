// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorSLSGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridMonitorSLSGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridMonitorSLSGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridMonitorSLSGroupResponseBody) *ModifyHybridMonitorSLSGroupResponse
	GetBody() *ModifyHybridMonitorSLSGroupResponseBody
}

type ModifyHybridMonitorSLSGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridMonitorSLSGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridMonitorSLSGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorSLSGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorSLSGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridMonitorSLSGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridMonitorSLSGroupResponse) GetBody() *ModifyHybridMonitorSLSGroupResponseBody {
	return s.Body
}

func (s *ModifyHybridMonitorSLSGroupResponse) SetHeaders(v map[string]*string) *ModifyHybridMonitorSLSGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponse) SetStatusCode(v int32) *ModifyHybridMonitorSLSGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponse) SetBody(v *ModifyHybridMonitorSLSGroupResponseBody) *ModifyHybridMonitorSLSGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridMonitorSLSGroupResponse) Validate() error {
	return dara.Validate(s)
}
