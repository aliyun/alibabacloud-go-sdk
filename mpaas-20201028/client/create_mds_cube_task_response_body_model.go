// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMdsCubeTaskResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMdsCubeTaskResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMdsCubeTaskResponseBodyResultContent) *CreateMdsCubeTaskResponseBody
	GetResultContent() *CreateMdsCubeTaskResponseBodyResultContent
	SetResultMessage(v string) *CreateMdsCubeTaskResponseBody
	GetResultMessage() *string
}

type CreateMdsCubeTaskResponseBody struct {
	// example:
	//
	// 0CC8A9CB-9BA3-13FD-A404-6E2E7461881A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                     `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMdsCubeTaskResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMdsCubeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTaskResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMdsCubeTaskResponseBody) GetResultContent() *CreateMdsCubeTaskResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMdsCubeTaskResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMdsCubeTaskResponseBody) SetRequestId(v string) *CreateMdsCubeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBody) SetResultCode(v string) *CreateMdsCubeTaskResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBody) SetResultContent(v *CreateMdsCubeTaskResponseBodyResultContent) *CreateMdsCubeTaskResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMdsCubeTaskResponseBody) SetResultMessage(v string) *CreateMdsCubeTaskResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeTaskResponseBodyResultContent struct {
	Data *CreateMdsCubeTaskResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 0CC8A9CB-9BA3-13FD-A404-6E2E7461881A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMdsCubeTaskResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTaskResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTaskResponseBodyResultContent) GetData() *CreateMdsCubeTaskResponseBodyResultContentData {
	return s.Data
}

func (s *CreateMdsCubeTaskResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTaskResponseBodyResultContent) SetData(v *CreateMdsCubeTaskResponseBodyResultContentData) *CreateMdsCubeTaskResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContent) SetRequestId(v string) *CreateMdsCubeTaskResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeTaskResponseBodyResultContentData struct {
	// example:
	//
	// 1010
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// NONE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 0CC8A9CB-9BA3-13FD-A404-6E2E7461881A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMdsCubeTaskResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTaskResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) SetContent(v string) *CreateMdsCubeTaskResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) SetErrorCode(v string) *CreateMdsCubeTaskResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) SetRequestId(v string) *CreateMdsCubeTaskResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) SetResultMsg(v string) *CreateMdsCubeTaskResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) SetSuccess(v bool) *CreateMdsCubeTaskResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *CreateMdsCubeTaskResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
