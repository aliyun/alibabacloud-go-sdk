// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsMiniprogramTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMdsMiniprogramTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMdsMiniprogramTaskResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMdsMiniprogramTaskResponseBodyResultContent) *CreateMdsMiniprogramTaskResponseBody
	GetResultContent() *CreateMdsMiniprogramTaskResponseBodyResultContent
	SetResultMessage(v string) *CreateMdsMiniprogramTaskResponseBody
	GetResultMessage() *string
}

type CreateMdsMiniprogramTaskResponseBody struct {
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMdsMiniprogramTaskResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMdsMiniprogramTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsMiniprogramTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMdsMiniprogramTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsMiniprogramTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMdsMiniprogramTaskResponseBody) GetResultContent() *CreateMdsMiniprogramTaskResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMdsMiniprogramTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMdsMiniprogramTaskResponseBody) SetRequestId(v string) *CreateMdsMiniprogramTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBody) SetResultCode(v string) *CreateMdsMiniprogramTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBody) SetResultContent(v *CreateMdsMiniprogramTaskResponseBodyResultContent) *CreateMdsMiniprogramTaskResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBody) SetResultMessage(v string) *CreateMdsMiniprogramTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMdsMiniprogramTaskResponseBodyResultContent struct {
	Data      *CreateMdsMiniprogramTaskResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMdsMiniprogramTaskResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsMiniprogramTaskResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContent) GetData() *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	return s.Data
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContent) SetData(v *CreateMdsMiniprogramTaskResponseBodyResultContentData) *CreateMdsMiniprogramTaskResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContent) SetRequestId(v string) *CreateMdsMiniprogramTaskResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type CreateMdsMiniprogramTaskResponseBodyResultContentData struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMdsMiniprogramTaskResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsMiniprogramTaskResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) SetContent(v string) *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) SetErrorCode(v string) *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) SetRequestId(v string) *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) SetResultMsg(v string) *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) SetSuccess(v bool) *CreateMdsMiniprogramTaskResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *CreateMdsMiniprogramTaskResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
