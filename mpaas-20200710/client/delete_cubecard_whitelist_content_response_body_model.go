// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCubecardWhitelistContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCubecardWhitelistContentResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteCubecardWhitelistContentResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteCubecardWhitelistContentResponseBodyResultContent) *DeleteCubecardWhitelistContentResponseBody
	GetResultContent() *DeleteCubecardWhitelistContentResponseBodyResultContent
	SetResultMessage(v string) *DeleteCubecardWhitelistContentResponseBody
	GetResultMessage() *string
}

type DeleteCubecardWhitelistContentResponseBody struct {
	RequestId     *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteCubecardWhitelistContentResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteCubecardWhitelistContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCubecardWhitelistContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCubecardWhitelistContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCubecardWhitelistContentResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteCubecardWhitelistContentResponseBody) GetResultContent() *DeleteCubecardWhitelistContentResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteCubecardWhitelistContentResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteCubecardWhitelistContentResponseBody) SetRequestId(v string) *DeleteCubecardWhitelistContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBody) SetResultCode(v string) *DeleteCubecardWhitelistContentResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBody) SetResultContent(v *DeleteCubecardWhitelistContentResponseBodyResultContent) *DeleteCubecardWhitelistContentResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBody) SetResultMessage(v string) *DeleteCubecardWhitelistContentResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCubecardWhitelistContentResponseBodyResultContent struct {
	Data      *DeleteCubecardWhitelistContentResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCubecardWhitelistContentResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteCubecardWhitelistContentResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContent) GetData() *DeleteCubecardWhitelistContentResponseBodyResultContentData {
	return s.Data
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContent) SetData(v *DeleteCubecardWhitelistContentResponseBodyResultContentData) *DeleteCubecardWhitelistContentResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContent) SetRequestId(v string) *DeleteCubecardWhitelistContentResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCubecardWhitelistContentResponseBodyResultContentData struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCubecardWhitelistContentResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCubecardWhitelistContentResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) GetContent() *string {
	return s.Content
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) SetContent(v string) *DeleteCubecardWhitelistContentResponseBodyResultContentData {
	s.Content = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) SetErrorCode(v string) *DeleteCubecardWhitelistContentResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) SetResultMsg(v string) *DeleteCubecardWhitelistContentResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) SetSuccess(v bool) *DeleteCubecardWhitelistContentResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *DeleteCubecardWhitelistContentResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
