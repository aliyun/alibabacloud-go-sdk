// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePullToPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLivePullToPushResponseBody
	GetDescription() *string
	SetRequestId(v string) *CreateLivePullToPushResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *CreateLivePullToPushResponseBody
	GetRetCode() *int32
	SetTaskId(v string) *CreateLivePullToPushResponseBody
	GetTaskId() *string
}

type CreateLivePullToPushResponseBody struct {
	// The description of the custom rule.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The code that is returned for the request.
	//
	// >
	//
	// 	- 0 is returned if the request is normal.
	//
	// 	- For information about codes that are returned when exceptions occur, see the following Error codes table.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fd245384-4067-4f91-9d75-9666a6bc9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLivePullToPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePullToPushResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePullToPushResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePullToPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivePullToPushResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *CreateLivePullToPushResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateLivePullToPushResponseBody) SetDescription(v string) *CreateLivePullToPushResponseBody {
	s.Description = &v
	return s
}

func (s *CreateLivePullToPushResponseBody) SetRequestId(v string) *CreateLivePullToPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePullToPushResponseBody) SetRetCode(v int32) *CreateLivePullToPushResponseBody {
	s.RetCode = &v
	return s
}

func (s *CreateLivePullToPushResponseBody) SetTaskId(v string) *CreateLivePullToPushResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLivePullToPushResponseBody) Validate() error {
	return dara.Validate(s)
}
