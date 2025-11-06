// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMdsCubeTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeMdsCubeTaskStatusResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ChangeMdsCubeTaskStatusResponseBody
	GetResultCode() *string
	SetResultContent(v *ChangeMdsCubeTaskStatusResponseBodyResultContent) *ChangeMdsCubeTaskStatusResponseBody
	GetResultContent() *ChangeMdsCubeTaskStatusResponseBodyResultContent
	SetResultMessage(v string) *ChangeMdsCubeTaskStatusResponseBody
	GetResultMessage() *string
}

type ChangeMdsCubeTaskStatusResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ChangeMdsCubeTaskStatusResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ChangeMdsCubeTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeMdsCubeTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMdsCubeTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMdsCubeTaskStatusResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ChangeMdsCubeTaskStatusResponseBody) GetResultContent() *ChangeMdsCubeTaskStatusResponseBodyResultContent {
	return s.ResultContent
}

func (s *ChangeMdsCubeTaskStatusResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ChangeMdsCubeTaskStatusResponseBody) SetRequestId(v string) *ChangeMdsCubeTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBody) SetResultCode(v string) *ChangeMdsCubeTaskStatusResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBody) SetResultContent(v *ChangeMdsCubeTaskStatusResponseBodyResultContent) *ChangeMdsCubeTaskStatusResponseBody {
	s.ResultContent = v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBody) SetResultMessage(v string) *ChangeMdsCubeTaskStatusResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeMdsCubeTaskStatusResponseBodyResultContent struct {
	Data *ChangeMdsCubeTaskStatusResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeMdsCubeTaskStatusResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ChangeMdsCubeTaskStatusResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContent) GetData() *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	return s.Data
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContent) SetData(v *ChangeMdsCubeTaskStatusResponseBodyResultContentData) *ChangeMdsCubeTaskStatusResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContent) SetRequestId(v string) *ChangeMdsCubeTaskStatusResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeMdsCubeTaskStatusResponseBodyResultContentData struct {
	// example:
	//
	// success
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

func (s ChangeMdsCubeTaskStatusResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ChangeMdsCubeTaskStatusResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) SetContent(v string) *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) SetErrorCode(v string) *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) SetRequestId(v string) *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) SetResultMsg(v string) *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) SetSuccess(v bool) *ChangeMdsCubeTaskStatusResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *ChangeMdsCubeTaskStatusResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
