// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMdsWhitelistContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMdsWhitelistContentResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMdsWhitelistContentResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteMdsWhitelistContentResponseBodyResultContent) *DeleteMdsWhitelistContentResponseBody
	GetResultContent() *DeleteMdsWhitelistContentResponseBodyResultContent
	SetResultMessage(v string) *DeleteMdsWhitelistContentResponseBody
	GetResultMessage() *string
}

type DeleteMdsWhitelistContentResponseBody struct {
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteMdsWhitelistContentResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMdsWhitelistContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsWhitelistContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMdsWhitelistContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMdsWhitelistContentResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMdsWhitelistContentResponseBody) GetResultContent() *DeleteMdsWhitelistContentResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteMdsWhitelistContentResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMdsWhitelistContentResponseBody) SetRequestId(v string) *DeleteMdsWhitelistContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBody) SetResultCode(v string) *DeleteMdsWhitelistContentResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBody) SetResultContent(v *DeleteMdsWhitelistContentResponseBodyResultContent) *DeleteMdsWhitelistContentResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBody) SetResultMessage(v string) *DeleteMdsWhitelistContentResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMdsWhitelistContentResponseBodyResultContent struct {
	Data      *DeleteMdsWhitelistContentResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMdsWhitelistContentResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsWhitelistContentResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContent) GetData() *DeleteMdsWhitelistContentResponseBodyResultContentData {
	return s.Data
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContent) SetData(v *DeleteMdsWhitelistContentResponseBodyResultContentData) *DeleteMdsWhitelistContentResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContent) SetRequestId(v string) *DeleteMdsWhitelistContentResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMdsWhitelistContentResponseBodyResultContentData struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMdsWhitelistContentResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMdsWhitelistContentResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) SetContent(v string) *DeleteMdsWhitelistContentResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) SetErrorCode(v string) *DeleteMdsWhitelistContentResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) SetResultMsg(v string) *DeleteMdsWhitelistContentResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) SetSuccess(v bool) *DeleteMdsWhitelistContentResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *DeleteMdsWhitelistContentResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
