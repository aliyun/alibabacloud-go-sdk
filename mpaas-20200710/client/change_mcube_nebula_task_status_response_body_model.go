// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeNebulaTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeMcubeNebulaTaskStatusResult(v *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) *ChangeMcubeNebulaTaskStatusResponseBody
	GetChangeMcubeNebulaTaskStatusResult() *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult
	SetRequestId(v string) *ChangeMcubeNebulaTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ChangeMcubeNebulaTaskStatusResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ChangeMcubeNebulaTaskStatusResponseBody
	GetResultMessage() *string
}

type ChangeMcubeNebulaTaskStatusResponseBody struct {
	ChangeMcubeNebulaTaskStatusResult *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult `json:"ChangeMcubeNebulaTaskStatusResult,omitempty" xml:"ChangeMcubeNebulaTaskStatusResult,omitempty" type:"Struct"`
	RequestId                         *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                        *string                                                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                     *string                                                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ChangeMcubeNebulaTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeNebulaTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) GetChangeMcubeNebulaTaskStatusResult() *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult {
	return s.ChangeMcubeNebulaTaskStatusResult
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) SetChangeMcubeNebulaTaskStatusResult(v *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) *ChangeMcubeNebulaTaskStatusResponseBody {
	s.ChangeMcubeNebulaTaskStatusResult = v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) SetRequestId(v string) *ChangeMcubeNebulaTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) SetResultCode(v string) *ChangeMcubeNebulaTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) SetResultMessage(v string) *ChangeMcubeNebulaTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) GoString() string {
	return s.String()
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) SetErrorCode(v string) *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult {
	s.ErrorCode = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) SetRequestId(v string) *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) SetResultMsg(v string) *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult {
	s.ResultMsg = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) SetSuccess(v bool) *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult {
	s.Success = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusResponseBodyChangeMcubeNebulaTaskStatusResult) Validate() error {
	return dara.Validate(s)
}
