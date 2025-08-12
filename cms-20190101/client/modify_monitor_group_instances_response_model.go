// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyMonitorGroupInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyMonitorGroupInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyMonitorGroupInstancesResponseBody) *ModifyMonitorGroupInstancesResponse
	GetBody() *ModifyMonitorGroupInstancesResponseBody
}

type ModifyMonitorGroupInstancesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyMonitorGroupInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyMonitorGroupInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyMonitorGroupInstancesResponse) GetBody() *ModifyMonitorGroupInstancesResponseBody {
	return s.Body
}

func (s *ModifyMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *ModifyMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *ModifyMonitorGroupInstancesResponse) SetStatusCode(v int32) *ModifyMonitorGroupInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponse) SetBody(v *ModifyMonitorGroupInstancesResponseBody) *ModifyMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

func (s *ModifyMonitorGroupInstancesResponse) Validate() error {
	return dara.Validate(s)
}
