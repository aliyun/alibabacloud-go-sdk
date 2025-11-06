// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCubecardAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCubecardAppsResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListCubecardAppsResponseBody
	GetResultCode() *string
	SetResultContent(v *ListCubecardAppsResponseBodyResultContent) *ListCubecardAppsResponseBody
	GetResultContent() *ListCubecardAppsResponseBodyResultContent
	SetResultMessage(v string) *ListCubecardAppsResponseBody
	GetResultMessage() *string
}

type ListCubecardAppsResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListCubecardAppsResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListCubecardAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCubecardAppsResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListCubecardAppsResponseBody) GetResultContent() *ListCubecardAppsResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListCubecardAppsResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListCubecardAppsResponseBody) SetRequestId(v string) *ListCubecardAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCubecardAppsResponseBody) SetResultCode(v string) *ListCubecardAppsResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListCubecardAppsResponseBody) SetResultContent(v *ListCubecardAppsResponseBodyResultContent) *ListCubecardAppsResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListCubecardAppsResponseBody) SetResultMessage(v string) *ListCubecardAppsResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListCubecardAppsResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCubecardAppsResponseBodyResultContent struct {
	Data *ListCubecardAppsResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCubecardAppsResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsResponseBodyResultContent) GetData() *ListCubecardAppsResponseBodyResultContentData {
	return s.Data
}

func (s *ListCubecardAppsResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCubecardAppsResponseBodyResultContent) SetData(v *ListCubecardAppsResponseBodyResultContentData) *ListCubecardAppsResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContent) SetRequestId(v string) *ListCubecardAppsResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCubecardAppsResponseBodyResultContentData struct {
	Content []*ListCubecardAppsResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
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

func (s ListCubecardAppsResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsResponseBodyResultContentData) GetContent() []*ListCubecardAppsResponseBodyResultContentDataContent {
	return s.Content
}

func (s *ListCubecardAppsResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCubecardAppsResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCubecardAppsResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListCubecardAppsResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *ListCubecardAppsResponseBodyResultContentData) SetContent(v []*ListCubecardAppsResponseBodyResultContentDataContent) *ListCubecardAppsResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentData) SetErrorCode(v string) *ListCubecardAppsResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentData) SetRequestId(v string) *ListCubecardAppsResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentData) SetResultMsg(v string) *ListCubecardAppsResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentData) SetSuccess(v bool) *ListCubecardAppsResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCubecardAppsResponseBodyResultContentDataContent struct {
	// example:
	//
	// ALIPUB97DB9F1011141
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// app name
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListCubecardAppsResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsResponseBodyResultContentDataContent) GetAppId() *string {
	return s.AppId
}

func (s *ListCubecardAppsResponseBodyResultContentDataContent) GetAppName() *string {
	return s.AppName
}

func (s *ListCubecardAppsResponseBodyResultContentDataContent) SetAppId(v string) *ListCubecardAppsResponseBodyResultContentDataContent {
	s.AppId = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentDataContent) SetAppName(v string) *ListCubecardAppsResponseBodyResultContentDataContent {
	s.AppName = &v
	return s
}

func (s *ListCubecardAppsResponseBodyResultContentDataContent) Validate() error {
	return dara.Validate(s)
}
