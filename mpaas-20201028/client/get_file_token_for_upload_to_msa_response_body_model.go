// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileTokenForUploadToMsaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetFileTokenForUploadToMsaResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetFileTokenForUploadToMsaResponseBody
	GetResultCode() *string
	SetResultContent(v *GetFileTokenForUploadToMsaResponseBodyResultContent) *GetFileTokenForUploadToMsaResponseBody
	GetResultContent() *GetFileTokenForUploadToMsaResponseBodyResultContent
	SetResultMessage(v string) *GetFileTokenForUploadToMsaResponseBody
	GetResultMessage() *string
}

type GetFileTokenForUploadToMsaResponseBody struct {
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *GetFileTokenForUploadToMsaResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetFileTokenForUploadToMsaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileTokenForUploadToMsaResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileTokenForUploadToMsaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileTokenForUploadToMsaResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetFileTokenForUploadToMsaResponseBody) GetResultContent() *GetFileTokenForUploadToMsaResponseBodyResultContent {
	return s.ResultContent
}

func (s *GetFileTokenForUploadToMsaResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetFileTokenForUploadToMsaResponseBody) SetRequestId(v string) *GetFileTokenForUploadToMsaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBody) SetResultCode(v string) *GetFileTokenForUploadToMsaResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBody) SetResultContent(v *GetFileTokenForUploadToMsaResponseBodyResultContent) *GetFileTokenForUploadToMsaResponseBody {
	s.ResultContent = v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBody) SetResultMessage(v string) *GetFileTokenForUploadToMsaResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileTokenForUploadToMsaResponseBodyResultContent struct {
	Content   *GetFileTokenForUploadToMsaResponseBodyResultContentContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	ErrorCode *string                                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                     `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *string                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileTokenForUploadToMsaResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s GetFileTokenForUploadToMsaResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) GetContent() *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	return s.Content
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) GetSuccess() *string {
	return s.Success
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) SetContent(v *GetFileTokenForUploadToMsaResponseBodyResultContentContent) *GetFileTokenForUploadToMsaResponseBodyResultContent {
	s.Content = v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) SetErrorCode(v string) *GetFileTokenForUploadToMsaResponseBodyResultContent {
	s.ErrorCode = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) SetRequestId(v string) *GetFileTokenForUploadToMsaResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) SetResultMsg(v string) *GetFileTokenForUploadToMsaResponseBodyResultContent {
	s.ResultMsg = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) SetSuccess(v string) *GetFileTokenForUploadToMsaResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type GetFileTokenForUploadToMsaResponseBodyResultContentContent struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetFileTokenForUploadToMsaResponseBodyResultContentContent) String() string {
	return dara.Prettify(s)
}

func (s GetFileTokenForUploadToMsaResponseBodyResultContentContent) GoString() string {
	return s.String()
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetAccessid() *string {
	return s.Accessid
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetDir() *string {
	return s.Dir
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetExpire() *string {
	return s.Expire
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetHost() *string {
	return s.Host
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetPolicy() *string {
	return s.Policy
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) GetSignature() *string {
	return s.Signature
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetAccessid(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Accessid = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetDir(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Dir = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetExpire(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Expire = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetHost(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Host = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetPolicy(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Policy = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) SetSignature(v string) *GetFileTokenForUploadToMsaResponseBodyResultContentContent {
	s.Signature = &v
	return s
}

func (s *GetFileTokenForUploadToMsaResponseBodyResultContentContent) Validate() error {
	return dara.Validate(s)
}
