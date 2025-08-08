// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcdpZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMcdpZoneResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcdpZoneResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMcdpZoneResponseBodyResultContent) *QueryMcdpZoneResponseBody
	GetResultContent() *QueryMcdpZoneResponseBodyResultContent
	SetResultMessage(v string) *QueryMcdpZoneResponseBody
	GetResultMessage() *string
}

type QueryMcdpZoneResponseBody struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMcdpZoneResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcdpZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpZoneResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcdpZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcdpZoneResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcdpZoneResponseBody) GetResultContent() *QueryMcdpZoneResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMcdpZoneResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcdpZoneResponseBody) SetRequestId(v string) *QueryMcdpZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcdpZoneResponseBody) SetResultCode(v string) *QueryMcdpZoneResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcdpZoneResponseBody) SetResultContent(v *QueryMcdpZoneResponseBodyResultContent) *QueryMcdpZoneResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMcdpZoneResponseBody) SetResultMessage(v string) *QueryMcdpZoneResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcdpZoneResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMcdpZoneResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcdpZoneResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMcdpZoneResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMcdpZoneResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *QueryMcdpZoneResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *QueryMcdpZoneResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *QueryMcdpZoneResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcdpZoneResponseBodyResultContent) SetCode(v string) *QueryMcdpZoneResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *QueryMcdpZoneResponseBodyResultContent) SetData(v string) *QueryMcdpZoneResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *QueryMcdpZoneResponseBodyResultContent) SetMessage(v string) *QueryMcdpZoneResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *QueryMcdpZoneResponseBodyResultContent) SetSuccess(v bool) *QueryMcdpZoneResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *QueryMcdpZoneResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
