// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMdsCubeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMdsCubeResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateMdsCubeResourceResponseBody
	GetResultCode() *string
	SetResultContent(v *UpdateMdsCubeResourceResponseBodyResultContent) *UpdateMdsCubeResourceResponseBody
	GetResultContent() *UpdateMdsCubeResourceResponseBodyResultContent
	SetResultMessage(v string) *UpdateMdsCubeResourceResponseBody
	GetResultMessage() *string
}

type UpdateMdsCubeResourceResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *UpdateMdsCubeResourceResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateMdsCubeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMdsCubeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMdsCubeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMdsCubeResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateMdsCubeResourceResponseBody) GetResultContent() *UpdateMdsCubeResourceResponseBodyResultContent {
	return s.ResultContent
}

func (s *UpdateMdsCubeResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateMdsCubeResourceResponseBody) SetRequestId(v string) *UpdateMdsCubeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBody) SetResultCode(v string) *UpdateMdsCubeResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBody) SetResultContent(v *UpdateMdsCubeResourceResponseBodyResultContent) *UpdateMdsCubeResourceResponseBody {
	s.ResultContent = v
	return s
}

func (s *UpdateMdsCubeResourceResponseBody) SetResultMessage(v string) *UpdateMdsCubeResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMdsCubeResourceResponseBodyResultContent struct {
	Data *UpdateMdsCubeResourceResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMdsCubeResourceResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateMdsCubeResourceResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *UpdateMdsCubeResourceResponseBodyResultContent) GetData() *UpdateMdsCubeResourceResponseBodyResultContentData {
	return s.Data
}

func (s *UpdateMdsCubeResourceResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMdsCubeResourceResponseBodyResultContent) SetData(v *UpdateMdsCubeResourceResponseBodyResultContentData) *UpdateMdsCubeResourceResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContent) SetRequestId(v string) *UpdateMdsCubeResourceResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMdsCubeResourceResponseBodyResultContentData struct {
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

func (s UpdateMdsCubeResourceResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMdsCubeResourceResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) SetContent(v string) *UpdateMdsCubeResourceResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) SetErrorCode(v string) *UpdateMdsCubeResourceResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) SetRequestId(v string) *UpdateMdsCubeResourceResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) SetResultMsg(v string) *UpdateMdsCubeResourceResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) SetSuccess(v bool) *UpdateMdsCubeResourceResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *UpdateMdsCubeResourceResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
