// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordLifecycleActionHeartbeatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecordLifecycleActionHeartbeatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecordLifecycleActionHeartbeatResponse
	GetStatusCode() *int32
	SetBody(v *RecordLifecycleActionHeartbeatResponseBody) *RecordLifecycleActionHeartbeatResponse
	GetBody() *RecordLifecycleActionHeartbeatResponseBody
}

type RecordLifecycleActionHeartbeatResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecordLifecycleActionHeartbeatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecordLifecycleActionHeartbeatResponse) String() string {
	return dara.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatResponse) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecordLifecycleActionHeartbeatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecordLifecycleActionHeartbeatResponse) GetBody() *RecordLifecycleActionHeartbeatResponseBody {
	return s.Body
}

func (s *RecordLifecycleActionHeartbeatResponse) SetHeaders(v map[string]*string) *RecordLifecycleActionHeartbeatResponse {
	s.Headers = v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponse) SetStatusCode(v int32) *RecordLifecycleActionHeartbeatResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponse) SetBody(v *RecordLifecycleActionHeartbeatResponseBody) *RecordLifecycleActionHeartbeatResponse {
	s.Body = v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponse) Validate() error {
	return dara.Validate(s)
}
