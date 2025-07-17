// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordLifecycleActionHeartbeatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecordLifecycleActionHeartbeatResponseBody
	GetRequestId() *string
}

type RecordLifecycleActionHeartbeatResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecordLifecycleActionHeartbeatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecordLifecycleActionHeartbeatResponseBody) GoString() string {
	return s.String()
}

func (s *RecordLifecycleActionHeartbeatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecordLifecycleActionHeartbeatResponseBody) SetRequestId(v string) *RecordLifecycleActionHeartbeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecordLifecycleActionHeartbeatResponseBody) Validate() error {
	return dara.Validate(s)
}
