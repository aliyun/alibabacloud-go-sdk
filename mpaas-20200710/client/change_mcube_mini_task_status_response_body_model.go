// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeMiniTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeMiniTaskStatusResult(v *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) *ChangeMcubeMiniTaskStatusResponseBody
	GetChangeMiniTaskStatusResult() *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult
	SetRequestId(v string) *ChangeMcubeMiniTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ChangeMcubeMiniTaskStatusResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ChangeMcubeMiniTaskStatusResponseBody
	GetResultMessage() *string
}

type ChangeMcubeMiniTaskStatusResponseBody struct {
	ChangeMiniTaskStatusResult *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult `json:"ChangeMiniTaskStatusResult,omitempty" xml:"ChangeMiniTaskStatusResult,omitempty" type:"Struct"`
	RequestId                  *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                 *string                                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage              *string                                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ChangeMcubeMiniTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeMiniTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) GetChangeMiniTaskStatusResult() *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult {
	return s.ChangeMiniTaskStatusResult
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) SetChangeMiniTaskStatusResult(v *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) *ChangeMcubeMiniTaskStatusResponseBody {
	s.ChangeMiniTaskStatusResult = v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) SetRequestId(v string) *ChangeMcubeMiniTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) SetResultCode(v string) *ChangeMcubeMiniTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) SetResultMessage(v string) *ChangeMcubeMiniTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) GoString() string {
	return s.String()
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) GetData() *string {
	return s.Data
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) SetData(v string) *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult {
	s.Data = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) SetResultMsg(v string) *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult {
	s.ResultMsg = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) SetSuccess(v bool) *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult {
	s.Success = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusResponseBodyChangeMiniTaskStatusResult) Validate() error {
	return dara.Validate(s)
}
