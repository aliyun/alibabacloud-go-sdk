// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WakeUpClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *WakeUpClusterResponseBody
	GetTaskId() *string
}

type WakeUpClusterResponseBody struct {
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// T-67d7ec016ce37c0106000***
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s WakeUpClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WakeUpClusterResponseBody) GoString() string {
	return s.String()
}

func (s *WakeUpClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WakeUpClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *WakeUpClusterResponseBody) SetRequestId(v string) *WakeUpClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *WakeUpClusterResponseBody) SetTaskId(v string) *WakeUpClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *WakeUpClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
