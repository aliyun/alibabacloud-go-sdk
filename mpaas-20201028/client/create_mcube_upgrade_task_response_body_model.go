// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTaskResult(v *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) *CreateMcubeUpgradeTaskResponseBody
	GetCreateTaskResult() *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult
	SetRequestId(v string) *CreateMcubeUpgradeTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeUpgradeTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeUpgradeTaskResponseBody
	GetResultMessage() *string
}

type CreateMcubeUpgradeTaskResponseBody struct {
	CreateTaskResult *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult `json:"CreateTaskResult,omitempty" xml:"CreateTaskResult,omitempty" type:"Struct"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode       *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage    *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeUpgradeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradeTaskResponseBody) GetCreateTaskResult() *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	return s.CreateTaskResult
}

func (s *CreateMcubeUpgradeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradeTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeUpgradeTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeUpgradeTaskResponseBody) SetCreateTaskResult(v *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) *CreateMcubeUpgradeTaskResponseBody {
	s.CreateTaskResult = v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBody) SetRequestId(v string) *CreateMcubeUpgradeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBody) SetResultCode(v string) *CreateMcubeUpgradeTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBody) SetResultMessage(v string) *CreateMcubeUpgradeTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeUpgradeTaskResponseBodyCreateTaskResult struct {
	ErrorCode     *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg     *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UpgradeTaskId *string `json:"upgradeTaskId,omitempty" xml:"upgradeTaskId,omitempty"`
}

func (s CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) GetUpgradeTaskId() *string {
	return s.UpgradeTaskId
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) SetErrorCode(v string) *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) SetRequestId(v string) *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) SetResultMsg(v string) *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) SetSuccess(v bool) *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) SetUpgradeTaskId(v string) *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult {
	s.UpgradeTaskId = &v
	return s
}

func (s *CreateMcubeUpgradeTaskResponseBodyCreateTaskResult) Validate() error {
	return dara.Validate(s)
}
