// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMiniTaskResult(v *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) *CreateMcubeMiniTaskResponseBody
	GetCreateMiniTaskResult() *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult
	SetRequestId(v string) *CreateMcubeMiniTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeMiniTaskResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeMiniTaskResponseBody
	GetResultMessage() *string
}

type CreateMcubeMiniTaskResponseBody struct {
	CreateMiniTaskResult *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult `json:"CreateMiniTaskResult,omitempty" xml:"CreateMiniTaskResult,omitempty" type:"Struct"`
	RequestId            *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode           *string                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage        *string                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeMiniTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniTaskResponseBody) GetCreateMiniTaskResult() *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult {
	return s.CreateMiniTaskResult
}

func (s *CreateMcubeMiniTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeMiniTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeMiniTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeMiniTaskResponseBody) SetCreateMiniTaskResult(v *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) *CreateMcubeMiniTaskResponseBody {
	s.CreateMiniTaskResult = v
	return s
}

func (s *CreateMcubeMiniTaskResponseBody) SetRequestId(v string) *CreateMcubeMiniTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBody) SetResultCode(v string) *CreateMcubeMiniTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBody) SetResultMessage(v string) *CreateMcubeMiniTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult struct {
	MiniTaskId *string `json:"MiniTaskId,omitempty" xml:"MiniTaskId,omitempty"`
	ResultMsg  *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) GetMiniTaskId() *string {
	return s.MiniTaskId
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) SetMiniTaskId(v string) *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult {
	s.MiniTaskId = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) SetResultMsg(v string) *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) SetSuccess(v bool) *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeMiniTaskResponseBodyCreateMiniTaskResult) Validate() error {
	return dara.Validate(s)
}
