// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubePublicTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeResult(v *ChangeMcubePublicTaskStatusResponseBodyChangeResult) *ChangeMcubePublicTaskStatusResponseBody
	GetChangeResult() *ChangeMcubePublicTaskStatusResponseBodyChangeResult
	SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetResultMessage() *string
}

type ChangeMcubePublicTaskStatusResponseBody struct {
	ChangeResult  *ChangeMcubePublicTaskStatusResponseBodyChangeResult `json:"ChangeResult,omitempty" xml:"ChangeResult,omitempty" type:"Struct"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ChangeMcubePublicTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetChangeResult() *ChangeMcubePublicTaskStatusResponseBodyChangeResult {
	return s.ChangeResult
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetChangeResult(v *ChangeMcubePublicTaskStatusResponseBodyChangeResult) *ChangeMcubePublicTaskStatusResponseBody {
	s.ChangeResult = v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetResultCode(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetResultMessage(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeMcubePublicTaskStatusResponseBodyChangeResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMcubePublicTaskStatusResponseBodyChangeResult) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusResponseBodyChangeResult) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) SetErrorCode(v string) *ChangeMcubePublicTaskStatusResponseBodyChangeResult {
	s.ErrorCode = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBodyChangeResult {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) SetResultMsg(v string) *ChangeMcubePublicTaskStatusResponseBodyChangeResult {
	s.ResultMsg = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) SetSuccess(v bool) *ChangeMcubePublicTaskStatusResponseBodyChangeResult {
	s.Success = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyChangeResult) Validate() error {
	return dara.Validate(s)
}
