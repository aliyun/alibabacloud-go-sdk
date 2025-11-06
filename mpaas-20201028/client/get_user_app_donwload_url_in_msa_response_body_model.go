// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppDonwloadUrlInMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserAppDonwloadUrlInMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetUserAppDonwloadUrlInMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) *GetUserAppDonwloadUrlInMsaResponseBody
	GetResultContent() *GetUserAppDonwloadUrlInMsaResponseBodyResultContent
	SetResultMessage(v string) *GetUserAppDonwloadUrlInMsaResponseBody
	GetResultMessage() *string
}

type GetUserAppDonwloadUrlInMsaResponseBody struct {
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetUserAppDonwloadUrlInMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetUserAppDonwloadUrlInMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppDonwloadUrlInMsaResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) GetResultContent() *GetUserAppDonwloadUrlInMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) SetRequestId(v string) *GetUserAppDonwloadUrlInMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) SetResultCode(v string) *GetUserAppDonwloadUrlInMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) SetResultContent(v *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) *GetUserAppDonwloadUrlInMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) SetResultMessage(v string) *GetUserAppDonwloadUrlInMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppDonwloadUrlInMsaResponseBodyResultContent struct {
	Code    *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserAppDonwloadUrlInMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppDonwloadUrlInMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) GetData() *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData {
	return s.Data
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) SetCode(v string) *GetUserAppDonwloadUrlInMsaResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) SetData(v *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) *GetUserAppDonwloadUrlInMsaResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) SetMessage(v string) *GetUserAppDonwloadUrlInMsaResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) SetSuccess(v bool) *GetUserAppDonwloadUrlInMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAppDonwloadUrlInMsaResponseBodyResultContentData struct {
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) GetFilename() *string {
	return s.Filename
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) GetUrl() *string {
	return s.Url
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) SetFilename(v string) *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData {
	s.Filename = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) SetUrl(v string) *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData {
	s.Url = &v
	return s
}

func (s *GetUserAppDonwloadUrlInMsaResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
