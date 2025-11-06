// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsCubeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMdsCubeTemplateResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMdsCubeTemplateResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteMdsCubeTemplateResponseBodyResultContent) *DeleteMdsCubeTemplateResponseBody
	GetResultContent() *DeleteMdsCubeTemplateResponseBodyResultContent
	SetResultMessage(v string) *DeleteMdsCubeTemplateResponseBody
	GetResultMessage() *string
}

type DeleteMdsCubeTemplateResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteMdsCubeTemplateResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMdsCubeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsCubeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMdsCubeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMdsCubeTemplateResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMdsCubeTemplateResponseBody) GetResultContent() *DeleteMdsCubeTemplateResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteMdsCubeTemplateResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMdsCubeTemplateResponseBody) SetRequestId(v string) *DeleteMdsCubeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBody) SetResultCode(v string) *DeleteMdsCubeTemplateResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBody) SetResultContent(v *DeleteMdsCubeTemplateResponseBodyResultContent) *DeleteMdsCubeTemplateResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBody) SetResultMessage(v string) *DeleteMdsCubeTemplateResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMdsCubeTemplateResponseBodyResultContent struct {
	Data *DeleteMdsCubeTemplateResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMdsCubeTemplateResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsCubeTemplateResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContent) GetData() *DeleteMdsCubeTemplateResponseBodyResultContentData {
	return s.Data
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContent) SetData(v *DeleteMdsCubeTemplateResponseBodyResultContentData) *DeleteMdsCubeTemplateResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContent) SetRequestId(v string) *DeleteMdsCubeTemplateResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMdsCubeTemplateResponseBodyResultContentData struct {
	// example:
	//
	// ""
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

func (s DeleteMdsCubeTemplateResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsCubeTemplateResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) SetContent(v string) *DeleteMdsCubeTemplateResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) SetErrorCode(v string) *DeleteMdsCubeTemplateResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) SetRequestId(v string) *DeleteMdsCubeTemplateResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) SetResultMsg(v string) *DeleteMdsCubeTemplateResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) SetSuccess(v bool) *DeleteMdsCubeTemplateResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *DeleteMdsCubeTemplateResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
