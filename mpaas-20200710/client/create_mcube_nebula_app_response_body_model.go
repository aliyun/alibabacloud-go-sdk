// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateNebulaAppResult(v *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) *CreateMcubeNebulaAppResponseBody
	GetCreateNebulaAppResult() *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult
	SetRequestId(v string) *CreateMcubeNebulaAppResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeNebulaAppResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeNebulaAppResponseBody
	GetResultMessage() *string
}

type CreateMcubeNebulaAppResponseBody struct {
	CreateNebulaAppResult *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult `json:"CreateNebulaAppResult,omitempty" xml:"CreateNebulaAppResult,omitempty" type:"Struct"`
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode            *string                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage         *string                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeNebulaAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaAppResponseBody) GetCreateNebulaAppResult() *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult {
	return s.CreateNebulaAppResult
}

func (s *CreateMcubeNebulaAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaAppResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeNebulaAppResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeNebulaAppResponseBody) SetCreateNebulaAppResult(v *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) *CreateMcubeNebulaAppResponseBody {
	s.CreateNebulaAppResult = v
	return s
}

func (s *CreateMcubeNebulaAppResponseBody) SetRequestId(v string) *CreateMcubeNebulaAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBody) SetResultCode(v string) *CreateMcubeNebulaAppResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBody) SetResultMessage(v string) *CreateMcubeNebulaAppResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) SetErrorCode(v string) *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) SetRequestId(v string) *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) SetResultMsg(v string) *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) SetSuccess(v bool) *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult {
	s.Success = &v
	return s
}

func (s *CreateMcubeNebulaAppResponseBodyCreateNebulaAppResult) Validate() error {
	return dara.Validate(s)
}
