// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpAimResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMcdpAimResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcdpAimResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMcdpAimResponseBodyResultContent) *QueryMcdpAimResponseBody
	GetResultContent() *QueryMcdpAimResponseBodyResultContent
	SetResultMessage(v string) *QueryMcdpAimResponseBody
	GetResultMessage() *string
}

type QueryMcdpAimResponseBody struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMcdpAimResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcdpAimResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpAimResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcdpAimResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcdpAimResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcdpAimResponseBody) GetResultContent() *QueryMcdpAimResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMcdpAimResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcdpAimResponseBody) SetRequestId(v string) *QueryMcdpAimResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcdpAimResponseBody) SetResultCode(v string) *QueryMcdpAimResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcdpAimResponseBody) SetResultContent(v *QueryMcdpAimResponseBodyResultContent) *QueryMcdpAimResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMcdpAimResponseBody) SetResultMessage(v string) *QueryMcdpAimResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcdpAimResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMcdpAimResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcdpAimResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpAimResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMcdpAimResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *QueryMcdpAimResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *QueryMcdpAimResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *QueryMcdpAimResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcdpAimResponseBodyResultContent) SetCode(v string) *QueryMcdpAimResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *QueryMcdpAimResponseBodyResultContent) SetData(v string) *QueryMcdpAimResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *QueryMcdpAimResponseBodyResultContent) SetMessage(v string) *QueryMcdpAimResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *QueryMcdpAimResponseBodyResultContent) SetSuccess(v bool) *QueryMcdpAimResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *QueryMcdpAimResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
