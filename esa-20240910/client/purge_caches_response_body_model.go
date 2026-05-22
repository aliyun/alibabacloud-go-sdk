// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PurgeCachesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *PurgeCachesResponseBody
	GetTaskId() *string
}

type PurgeCachesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID, which is returned when you create a refresh or preheat task.
	//
	// example:
	//
	// 15940956620
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PurgeCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurgeCachesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PurgeCachesResponseBody) SetRequestId(v string) *PurgeCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurgeCachesResponseBody) SetTaskId(v string) *PurgeCachesResponseBody {
	s.TaskId = &v
	return s
}

func (s *PurgeCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
