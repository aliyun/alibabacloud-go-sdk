// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskInstanceLogResponseBody
	GetRequestId() *string
	SetTaskInstanceLog(v string) *GetTaskInstanceLogResponseBody
	GetTaskInstanceLog() *string
}

type GetTaskInstanceLogResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The run log of the instance.
	//
	// example:
	//
	// This is running log
	TaskInstanceLog *string `json:"TaskInstanceLog,omitempty" xml:"TaskInstanceLog,omitempty"`
}

func (s GetTaskInstanceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskInstanceLogResponseBody) GetTaskInstanceLog() *string {
	return s.TaskInstanceLog
}

func (s *GetTaskInstanceLogResponseBody) SetRequestId(v string) *GetTaskInstanceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskInstanceLogResponseBody) SetTaskInstanceLog(v string) *GetTaskInstanceLogResponseBody {
	s.TaskInstanceLog = &v
	return s
}

func (s *GetTaskInstanceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
