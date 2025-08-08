// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubePublicTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetResultCode() *string
	SetResultContent(v *ChangeMcubePublicTaskStatusResponseBodyResultContent) *ChangeMcubePublicTaskStatusResponseBody
	GetResultContent() *ChangeMcubePublicTaskStatusResponseBodyResultContent
	SetResultMessage(v string) *ChangeMcubePublicTaskStatusResponseBody
	GetResultMessage() *string
}

type ChangeMcubePublicTaskStatusResponseBody struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                               `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ChangeMcubePublicTaskStatusResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                               `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ChangeMcubePublicTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetResultContent() *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	return s.ResultContent
}

func (s *ChangeMcubePublicTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetResultCode(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetResultContent(v *ChangeMcubePublicTaskStatusResponseBodyResultContent) *ChangeMcubePublicTaskStatusResponseBody {
	s.ResultContent = v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) SetResultMessage(v string) *ChangeMcubePublicTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChangeMcubePublicTaskStatusResponseBodyResultContent struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMcubePublicTaskStatusResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) SetData(v string) *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) SetErrorCode(v string) *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	s.ErrorCode = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) SetRequestId(v string) *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) SetResultMsg(v string) *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	s.ResultMsg = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) SetSuccess(v bool) *ChangeMcubePublicTaskStatusResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
