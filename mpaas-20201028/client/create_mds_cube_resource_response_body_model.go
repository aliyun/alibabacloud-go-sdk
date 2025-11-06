// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMdsCubeResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMdsCubeResourceResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMdsCubeResourceResponseBodyResultContent) *CreateMdsCubeResourceResponseBody
	GetResultContent() *CreateMdsCubeResourceResponseBodyResultContent
	SetResultMessage(v string) *CreateMdsCubeResourceResponseBody
	GetResultMessage() *string
}

type CreateMdsCubeResourceResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMdsCubeResourceResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMdsCubeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMdsCubeResourceResponseBody) GetResultContent() *CreateMdsCubeResourceResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMdsCubeResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMdsCubeResourceResponseBody) SetRequestId(v string) *CreateMdsCubeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBody) SetResultCode(v string) *CreateMdsCubeResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBody) SetResultContent(v *CreateMdsCubeResourceResponseBodyResultContent) *CreateMdsCubeResourceResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMdsCubeResourceResponseBody) SetResultMessage(v string) *CreateMdsCubeResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeResourceResponseBodyResultContent struct {
	Data *CreateMdsCubeResourceResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMdsCubeResourceResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeResourceResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeResourceResponseBodyResultContent) GetData() *CreateMdsCubeResourceResponseBodyResultContentData {
	return s.Data
}

func (s *CreateMdsCubeResourceResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeResourceResponseBodyResultContent) SetData(v *CreateMdsCubeResourceResponseBodyResultContentData) *CreateMdsCubeResourceResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContent) SetRequestId(v string) *CreateMdsCubeResourceResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeResourceResponseBodyResultContentData struct {
	// example:
	//
	// 1019
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// NONE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
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

func (s CreateMdsCubeResourceResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeResourceResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) SetContent(v string) *CreateMdsCubeResourceResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) SetErrorCode(v string) *CreateMdsCubeResourceResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) SetRequestId(v string) *CreateMdsCubeResourceResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) SetResultMsg(v string) *CreateMdsCubeResourceResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) SetSuccess(v bool) *CreateMdsCubeResourceResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *CreateMdsCubeResourceResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
