// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMiniResult(v *CreateMcubeMiniAppResponseBodyCreateMiniResult) *CreateMcubeMiniAppResponseBody
	GetCreateMiniResult() *CreateMcubeMiniAppResponseBodyCreateMiniResult
	SetRequestId(v string) *CreateMcubeMiniAppResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeMiniAppResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeMiniAppResponseBody
	GetResultMessage() *string
}

type CreateMcubeMiniAppResponseBody struct {
	CreateMiniResult *CreateMcubeMiniAppResponseBodyCreateMiniResult `json:"CreateMiniResult,omitempty" xml:"CreateMiniResult,omitempty" type:"Struct"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode       *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage    *string                                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeMiniAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniAppResponseBody) GetCreateMiniResult() *CreateMcubeMiniAppResponseBodyCreateMiniResult {
	return s.CreateMiniResult
}

func (s *CreateMcubeMiniAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeMiniAppResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeMiniAppResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeMiniAppResponseBody) SetCreateMiniResult(v *CreateMcubeMiniAppResponseBodyCreateMiniResult) *CreateMcubeMiniAppResponseBody {
	s.CreateMiniResult = v
	return s
}

func (s *CreateMcubeMiniAppResponseBody) SetRequestId(v string) *CreateMcubeMiniAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBody) SetResultCode(v string) *CreateMcubeMiniAppResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBody) SetResultMessage(v string) *CreateMcubeMiniAppResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeMiniAppResponseBodyCreateMiniResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeMiniAppResponseBodyCreateMiniResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniAppResponseBodyCreateMiniResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) GetData() *string {
	return s.Data
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) SetData(v string) *CreateMcubeMiniAppResponseBodyCreateMiniResult {
	s.Data = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) SetResultMsg(v string) *CreateMcubeMiniAppResponseBodyCreateMiniResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) SetSuccess(v bool) *CreateMcubeMiniAppResponseBodyCreateMiniResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeMiniAppResponseBodyCreateMiniResult) Validate() error {
	return dara.Validate(s)
}
