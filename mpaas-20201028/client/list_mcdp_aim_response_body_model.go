// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcdpAimResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListMcdpAimResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcdpAimResponseBody
	GetResultCode() *string
	SetResultContent(v *ListMcdpAimResponseBodyResultContent) *ListMcdpAimResponseBody
	GetResultContent() *ListMcdpAimResponseBodyResultContent
	SetResultMessage(v string) *ListMcdpAimResponseBody
	GetResultMessage() *string
}

type ListMcdpAimResponseBody struct {
	RequestId     *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                               `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListMcdpAimResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                               `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcdpAimResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcdpAimResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcdpAimResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcdpAimResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcdpAimResponseBody) GetResultContent() *ListMcdpAimResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListMcdpAimResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcdpAimResponseBody) SetRequestId(v string) *ListMcdpAimResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcdpAimResponseBody) SetResultCode(v string) *ListMcdpAimResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcdpAimResponseBody) SetResultContent(v *ListMcdpAimResponseBodyResultContent) *ListMcdpAimResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListMcdpAimResponseBody) SetResultMessage(v string) *ListMcdpAimResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcdpAimResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcdpAimResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcdpAimResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListMcdpAimResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListMcdpAimResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *ListMcdpAimResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *ListMcdpAimResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *ListMcdpAimResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcdpAimResponseBodyResultContent) SetCode(v string) *ListMcdpAimResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *ListMcdpAimResponseBodyResultContent) SetData(v string) *ListMcdpAimResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *ListMcdpAimResponseBodyResultContent) SetMessage(v string) *ListMcdpAimResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *ListMcdpAimResponseBodyResultContent) SetSuccess(v bool) *ListMcdpAimResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *ListMcdpAimResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
