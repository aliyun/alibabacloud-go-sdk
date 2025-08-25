// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateVhostResult(v *CreateMcubeVhostResponseBodyCreateVhostResult) *CreateMcubeVhostResponseBody
	GetCreateVhostResult() *CreateMcubeVhostResponseBodyCreateVhostResult
	SetRequestId(v string) *CreateMcubeVhostResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeVhostResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeVhostResponseBody
	GetResultMessage() *string
}

type CreateMcubeVhostResponseBody struct {
	CreateVhostResult *CreateMcubeVhostResponseBodyCreateVhostResult `json:"CreateVhostResult,omitempty" xml:"CreateVhostResult,omitempty" type:"Struct"`
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode        *string                                        `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage     *string                                        `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeVhostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeVhostResponseBody) GetCreateVhostResult() *CreateMcubeVhostResponseBodyCreateVhostResult {
	return s.CreateVhostResult
}

func (s *CreateMcubeVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeVhostResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeVhostResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeVhostResponseBody) SetCreateVhostResult(v *CreateMcubeVhostResponseBodyCreateVhostResult) *CreateMcubeVhostResponseBody {
	s.CreateVhostResult = v
	return s
}

func (s *CreateMcubeVhostResponseBody) SetRequestId(v string) *CreateMcubeVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeVhostResponseBody) SetResultCode(v string) *CreateMcubeVhostResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeVhostResponseBody) SetResultMessage(v string) *CreateMcubeVhostResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeVhostResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeVhostResponseBodyCreateVhostResult struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeVhostResponseBodyCreateVhostResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeVhostResponseBodyCreateVhostResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) GetData() *string {
	return s.Data
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) SetData(v string) *CreateMcubeVhostResponseBodyCreateVhostResult {
	s.Data = &v
	return s
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) SetResultMsg(v string) *CreateMcubeVhostResponseBodyCreateVhostResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) SetSuccess(v bool) *CreateMcubeVhostResponseBodyCreateVhostResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeVhostResponseBodyCreateVhostResult) Validate() error {
	return dara.Validate(s)
}
