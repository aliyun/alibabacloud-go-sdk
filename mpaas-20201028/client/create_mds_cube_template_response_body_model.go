// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMdsCubeTemplateResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMdsCubeTemplateResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMdsCubeTemplateResponseBodyResultContent) *CreateMdsCubeTemplateResponseBody
	GetResultContent() *CreateMdsCubeTemplateResponseBodyResultContent
	SetResultMessage(v string) *CreateMdsCubeTemplateResponseBody
	GetResultMessage() *string
}

type CreateMdsCubeTemplateResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMdsCubeTemplateResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMdsCubeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTemplateResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMdsCubeTemplateResponseBody) GetResultContent() *CreateMdsCubeTemplateResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMdsCubeTemplateResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMdsCubeTemplateResponseBody) SetRequestId(v string) *CreateMdsCubeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBody) SetResultCode(v string) *CreateMdsCubeTemplateResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBody) SetResultContent(v *CreateMdsCubeTemplateResponseBodyResultContent) *CreateMdsCubeTemplateResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMdsCubeTemplateResponseBody) SetResultMessage(v string) *CreateMdsCubeTemplateResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeTemplateResponseBodyResultContent struct {
	Data *CreateMdsCubeTemplateResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMdsCubeTemplateResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTemplateResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTemplateResponseBodyResultContent) GetData() *CreateMdsCubeTemplateResponseBodyResultContentData {
	return s.Data
}

func (s *CreateMdsCubeTemplateResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTemplateResponseBodyResultContent) SetData(v *CreateMdsCubeTemplateResponseBodyResultContentData) *CreateMdsCubeTemplateResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContent) SetRequestId(v string) *CreateMdsCubeTemplateResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMdsCubeTemplateResponseBodyResultContentData struct {
	// example:
	//
	// success
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// None
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

func (s CreateMdsCubeTemplateResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeTemplateResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) SetContent(v string) *CreateMdsCubeTemplateResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) SetErrorCode(v string) *CreateMdsCubeTemplateResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) SetRequestId(v string) *CreateMdsCubeTemplateResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) SetResultMsg(v string) *CreateMdsCubeTemplateResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) SetSuccess(v bool) *CreateMdsCubeTemplateResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *CreateMdsCubeTemplateResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
