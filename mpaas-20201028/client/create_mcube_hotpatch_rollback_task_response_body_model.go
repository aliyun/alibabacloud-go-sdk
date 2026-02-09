// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeHotpatchRollbackTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateHotpatchRollbackTaskResult(v *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) *CreateMcubeHotpatchRollbackTaskResponseBody
	GetCreateHotpatchRollbackTaskResult() *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult
	SetRequestId(v string) *CreateMcubeHotpatchRollbackTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeHotpatchRollbackTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeHotpatchRollbackTaskResponseBody
	GetResultMessage() *string
}

type CreateMcubeHotpatchRollbackTaskResponseBody struct {
	CreateHotpatchRollbackTaskResult *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult `json:"CreateHotpatchRollbackTaskResult,omitempty" xml:"CreateHotpatchRollbackTaskResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeHotpatchRollbackTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchRollbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) GetCreateHotpatchRollbackTaskResult() *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult {
	return s.CreateHotpatchRollbackTaskResult
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) SetCreateHotpatchRollbackTaskResult(v *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) *CreateMcubeHotpatchRollbackTaskResponseBody {
	s.CreateHotpatchRollbackTaskResult = v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) SetRequestId(v string) *CreateMcubeHotpatchRollbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) SetResultCode(v string) *CreateMcubeHotpatchRollbackTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) SetResultMessage(v string) *CreateMcubeHotpatchRollbackTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBody) Validate() error {
	if s.CreateHotpatchRollbackTaskResult != nil {
		if err := s.CreateHotpatchRollbackTaskResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult struct {
	// example:
	//
	// 06D5CA0C-F5D4-5D64-987E-D221C88AED29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// 1543
	RollbackTaskId *string `json:"RollbackTaskId,omitempty" xml:"RollbackTaskId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) GetRollbackTaskId() *string {
	return s.RollbackTaskId
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) SetRequestId(v string) *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) SetResultMsg(v string) *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) SetRollbackTaskId(v string) *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult {
	s.RollbackTaskId = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) SetSuccess(v bool) *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeHotpatchRollbackTaskResponseBodyCreateHotpatchRollbackTaskResult) Validate() error {
	return dara.Validate(s)
}
