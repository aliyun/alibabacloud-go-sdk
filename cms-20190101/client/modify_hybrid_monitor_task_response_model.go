// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridMonitorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridMonitorTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridMonitorTaskResponseBody) *ModifyHybridMonitorTaskResponse
	GetBody() *ModifyHybridMonitorTaskResponseBody
}

type ModifyHybridMonitorTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridMonitorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridMonitorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridMonitorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridMonitorTaskResponse) GetBody() *ModifyHybridMonitorTaskResponseBody {
	return s.Body
}

func (s *ModifyHybridMonitorTaskResponse) SetHeaders(v map[string]*string) *ModifyHybridMonitorTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridMonitorTaskResponse) SetStatusCode(v int32) *ModifyHybridMonitorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridMonitorTaskResponse) SetBody(v *ModifyHybridMonitorTaskResponseBody) *ModifyHybridMonitorTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridMonitorTaskResponse) Validate() error {
	return dara.Validate(s)
}
