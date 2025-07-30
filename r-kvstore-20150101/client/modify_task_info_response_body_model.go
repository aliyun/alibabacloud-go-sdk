// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyTaskInfoResponseBody
	GetErrorCode() *string
	SetErrorTaskId(v string) *ModifyTaskInfoResponseBody
	GetErrorTaskId() *string
	SetRequestId(v string) *ModifyTaskInfoResponseBody
	GetRequestId() *string
	SetSuccessCount(v string) *ModifyTaskInfoResponseBody
	GetSuccessCount() *string
}

type ModifyTaskInfoResponseBody struct {
	// The error code returned if a task fails.
	//
	// example:
	//
	// mst.errorcode.success.errormessage
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the failed task. This parameter is returned if a task fails.
	//
	// example:
	//
	// t-0mqaj5hnyofczv****
	ErrorTaskId *string `json:"ErrorTaskId,omitempty" xml:"ErrorTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2B98499B-E62B-56D4-8D7F-3D6D4DB260F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of completed tasks.
	//
	// example:
	//
	// 2
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s ModifyTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTaskInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyTaskInfoResponseBody) GetErrorTaskId() *string {
	return s.ErrorTaskId
}

func (s *ModifyTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTaskInfoResponseBody) GetSuccessCount() *string {
	return s.SuccessCount
}

func (s *ModifyTaskInfoResponseBody) SetErrorCode(v string) *ModifyTaskInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyTaskInfoResponseBody) SetErrorTaskId(v string) *ModifyTaskInfoResponseBody {
	s.ErrorTaskId = &v
	return s
}

func (s *ModifyTaskInfoResponseBody) SetRequestId(v string) *ModifyTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTaskInfoResponseBody) SetSuccessCount(v string) *ModifyTaskInfoResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ModifyTaskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
