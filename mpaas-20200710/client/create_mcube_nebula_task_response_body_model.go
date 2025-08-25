// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMcubeNebulaTaskResult(v *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) *CreateMcubeNebulaTaskResponseBody
	GetCreateMcubeNebulaTaskResult() *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult
	SetRequestId(v string) *CreateMcubeNebulaTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeNebulaTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeNebulaTaskResponseBody
	GetResultMessage() *string
}

type CreateMcubeNebulaTaskResponseBody struct {
	CreateMcubeNebulaTaskResult *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult `json:"CreateMcubeNebulaTaskResult,omitempty" xml:"CreateMcubeNebulaTaskResult,omitempty" type:"Struct"`
	RequestId                   *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                  *string                                                       `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage               *string                                                       `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeNebulaTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaTaskResponseBody) GetCreateMcubeNebulaTaskResult() *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	return s.CreateMcubeNebulaTaskResult
}

func (s *CreateMcubeNebulaTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeNebulaTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeNebulaTaskResponseBody) SetCreateMcubeNebulaTaskResult(v *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) *CreateMcubeNebulaTaskResponseBody {
	s.CreateMcubeNebulaTaskResult = v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBody) SetRequestId(v string) *CreateMcubeNebulaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBody) SetResultCode(v string) *CreateMcubeNebulaTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBody) SetResultMessage(v string) *CreateMcubeNebulaTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	NebulaTaskId *string `json:"NebulaTaskId,omitempty" xml:"NebulaTaskId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg    *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GetNebulaTaskId() *string {
	return s.NebulaTaskId
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) SetErrorCode(v string) *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) SetNebulaTaskId(v string) *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	s.NebulaTaskId = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) SetRequestId(v string) *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) SetResultMsg(v string) *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) SetSuccess(v bool) *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeNebulaTaskResponseBodyCreateMcubeNebulaTaskResult) Validate() error {
	return dara.Validate(s)
}
