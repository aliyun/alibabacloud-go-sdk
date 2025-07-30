// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySynchronizationObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifySynchronizationObjectResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifySynchronizationObjectResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ModifySynchronizationObjectResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySynchronizationObjectResponseBody
	GetSuccess() *string
	SetTaskId(v string) *ModifySynchronizationObjectResponseBody
	GetTaskId() *string
}

type ModifySynchronizationObjectResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 902DDCDE-C755-4458-85DA-DF9A323C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the task that changes the objects. You must specify the task ID when you call the DescribeSynchronizationObjectModifyStatus operation to query the status and progress of the task.
	//
	// example:
	//
	// tl911uvi25z****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifySynchronizationObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySynchronizationObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifySynchronizationObjectResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifySynchronizationObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySynchronizationObjectResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySynchronizationObjectResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifySynchronizationObjectResponseBody) SetErrCode(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrMessage(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetRequestId(v string) *ModifySynchronizationObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetSuccess(v string) *ModifySynchronizationObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetTaskId(v string) *ModifySynchronizationObjectResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
