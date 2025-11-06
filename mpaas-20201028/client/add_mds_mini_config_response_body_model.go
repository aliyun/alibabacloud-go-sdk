// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMdsMiniConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMdsMiniConfigResponseBody
	GetRequestId() *string
	SetResultCode(v string) *AddMdsMiniConfigResponseBody
	GetResultCode() *string
	SetResultContent(v *AddMdsMiniConfigResponseBodyResultContent) *AddMdsMiniConfigResponseBody
	GetResultContent() *AddMdsMiniConfigResponseBodyResultContent
	SetResultMessage(v string) *AddMdsMiniConfigResponseBody
	GetResultMessage() *string
}

type AddMdsMiniConfigResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *AddMdsMiniConfigResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s AddMdsMiniConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMdsMiniConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddMdsMiniConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMdsMiniConfigResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *AddMdsMiniConfigResponseBody) GetResultContent() *AddMdsMiniConfigResponseBodyResultContent {
	return s.ResultContent
}

func (s *AddMdsMiniConfigResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *AddMdsMiniConfigResponseBody) SetRequestId(v string) *AddMdsMiniConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMdsMiniConfigResponseBody) SetResultCode(v string) *AddMdsMiniConfigResponseBody {
	s.ResultCode = &v
	return s
}

func (s *AddMdsMiniConfigResponseBody) SetResultContent(v *AddMdsMiniConfigResponseBodyResultContent) *AddMdsMiniConfigResponseBody {
	s.ResultContent = v
	return s
}

func (s *AddMdsMiniConfigResponseBody) SetResultMessage(v string) *AddMdsMiniConfigResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *AddMdsMiniConfigResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMdsMiniConfigResponseBodyResultContent struct {
	Data      *AddMdsMiniConfigResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMdsMiniConfigResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s AddMdsMiniConfigResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *AddMdsMiniConfigResponseBodyResultContent) GetData() *AddMdsMiniConfigResponseBodyResultContentData {
	return s.Data
}

func (s *AddMdsMiniConfigResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMdsMiniConfigResponseBodyResultContent) SetData(v *AddMdsMiniConfigResponseBodyResultContentData) *AddMdsMiniConfigResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContent) SetRequestId(v string) *AddMdsMiniConfigResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMdsMiniConfigResponseBodyResultContentData struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMdsMiniConfigResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s AddMdsMiniConfigResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) SetContent(v string) *AddMdsMiniConfigResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) SetErrorCode(v string) *AddMdsMiniConfigResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) SetRequestId(v string) *AddMdsMiniConfigResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) SetResultMsg(v string) *AddMdsMiniConfigResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) SetSuccess(v bool) *AddMdsMiniConfigResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *AddMdsMiniConfigResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
