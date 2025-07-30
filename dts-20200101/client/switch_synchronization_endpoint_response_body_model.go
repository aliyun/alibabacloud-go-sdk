// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSynchronizationEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *SwitchSynchronizationEndpointResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SwitchSynchronizationEndpointResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SwitchSynchronizationEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SwitchSynchronizationEndpointResponseBody
	GetSuccess() *string
	SetTaskId(v string) *SwitchSynchronizationEndpointResponseBody
	GetTaskId() *string
}

type SwitchSynchronizationEndpointResponseBody struct {
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
	// 3232F84C-C961-4811-B014-4EA7A27C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID. You must specify the task ID when you call the [DescribeEndpointSwitchStatus](https://help.aliyun.com/document_detail/135598.html) operation to query the execution status of the task.
	//
	// example:
	//
	// 11****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchSynchronizationEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SwitchSynchronizationEndpointResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SwitchSynchronizationEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchSynchronizationEndpointResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SwitchSynchronizationEndpointResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrCode(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrMessage(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetRequestId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetSuccess(v string) *SwitchSynchronizationEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetTaskId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
