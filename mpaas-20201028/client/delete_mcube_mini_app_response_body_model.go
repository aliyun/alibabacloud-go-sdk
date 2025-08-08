// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcubeMiniAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMiniResult(v *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) *DeleteMcubeMiniAppResponseBody
	GetDeleteMiniResult() *DeleteMcubeMiniAppResponseBodyDeleteMiniResult
	SetRequestId(v string) *DeleteMcubeMiniAppResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcubeMiniAppResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *DeleteMcubeMiniAppResponseBody
	GetResultMessage() *string
}

type DeleteMcubeMiniAppResponseBody struct {
	DeleteMiniResult *DeleteMcubeMiniAppResponseBodyDeleteMiniResult `json:"DeleteMiniResult,omitempty" xml:"DeleteMiniResult,omitempty" type:"Struct"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode       *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage    *string                                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcubeMiniAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeMiniAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcubeMiniAppResponseBody) GetDeleteMiniResult() *DeleteMcubeMiniAppResponseBodyDeleteMiniResult {
	return s.DeleteMiniResult
}

func (s *DeleteMcubeMiniAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcubeMiniAppResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcubeMiniAppResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcubeMiniAppResponseBody) SetDeleteMiniResult(v *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) *DeleteMcubeMiniAppResponseBody {
	s.DeleteMiniResult = v
	return s
}

func (s *DeleteMcubeMiniAppResponseBody) SetRequestId(v string) *DeleteMcubeMiniAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBody) SetResultCode(v string) *DeleteMcubeMiniAppResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBody) SetResultMessage(v string) *DeleteMcubeMiniAppResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteMcubeMiniAppResponseBodyDeleteMiniResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcubeMiniAppResponseBodyDeleteMiniResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcubeMiniAppResponseBodyDeleteMiniResult) GoString() string {
	return s.String()
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) GetData() *string {
	return s.Data
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) SetData(v string) *DeleteMcubeMiniAppResponseBodyDeleteMiniResult {
	s.Data = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) SetResultMsg(v string) *DeleteMcubeMiniAppResponseBodyDeleteMiniResult {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) SetSuccess(v bool) *DeleteMcubeMiniAppResponseBodyDeleteMiniResult {
	s.Success = &v
	return s
}

func (s *DeleteMcubeMiniAppResponseBodyDeleteMiniResult) Validate() error {
	return dara.Validate(s)
}
