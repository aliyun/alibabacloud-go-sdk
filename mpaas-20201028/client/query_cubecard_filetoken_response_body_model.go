// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubecardFiletokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryCubecardFiletokenResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryCubecardFiletokenResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryCubecardFiletokenResponseBodyResultContent) *QueryCubecardFiletokenResponseBody
	GetResultContent() *QueryCubecardFiletokenResponseBodyResultContent
	SetResultMessage(v string) *QueryCubecardFiletokenResponseBody
	GetResultMessage() *string
}

type QueryCubecardFiletokenResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode    *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryCubecardFiletokenResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryCubecardFiletokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCubecardFiletokenResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryCubecardFiletokenResponseBody) GetResultContent() *QueryCubecardFiletokenResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryCubecardFiletokenResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryCubecardFiletokenResponseBody) SetRequestId(v string) *QueryCubecardFiletokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBody) SetResultCode(v string) *QueryCubecardFiletokenResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBody) SetResultContent(v *QueryCubecardFiletokenResponseBodyResultContent) *QueryCubecardFiletokenResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryCubecardFiletokenResponseBody) SetResultMessage(v string) *QueryCubecardFiletokenResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCubecardFiletokenResponseBodyResultContent struct {
	Data *QueryCubecardFiletokenResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCubecardFiletokenResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenResponseBodyResultContent) GetData() *QueryCubecardFiletokenResponseBodyResultContentData {
	return s.Data
}

func (s *QueryCubecardFiletokenResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCubecardFiletokenResponseBodyResultContent) SetData(v *QueryCubecardFiletokenResponseBodyResultContentData) *QueryCubecardFiletokenResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContent) SetRequestId(v string) *QueryCubecardFiletokenResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCubecardFiletokenResponseBodyResultContentData struct {
	Content *QueryCubecardFiletokenResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// NONE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCubecardFiletokenResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) GetContent() *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	return s.Content
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) SetContent(v *QueryCubecardFiletokenResponseBodyResultContentDataContent) *QueryCubecardFiletokenResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) SetErrorCode(v string) *QueryCubecardFiletokenResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) SetResultMsg(v string) *QueryCubecardFiletokenResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) SetSuccess(v bool) *QueryCubecardFiletokenResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCubecardFiletokenResponseBodyResultContentDataContent struct {
	// example:
	//
	// LTAI5tFgfNzJFDn4Y4BhoRnc
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// example:
	//
	// /home/ecs-assist-user/proof/alert
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 1760583447
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// 172.23.129.55
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// Permit
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// yKplu9LQwgKBTDhxp0YozAeCy9c=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s QueryCubecardFiletokenResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s QueryCubecardFiletokenResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetAccessid() *string {
	return s.Accessid
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetDir() *string {
	return s.Dir
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetExpire() *string {
	return s.Expire
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetHost() *string {
	return s.Host
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetPolicy() *string {
	return s.Policy
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) GetSignature() *string {
	return s.Signature
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetAccessid(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Accessid = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetDir(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Dir = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetExpire(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Expire = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetHost(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Host = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetPolicy(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Policy = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) SetSignature(v string) *QueryCubecardFiletokenResponseBodyResultContentDataContent {
	s.Signature = &v
	return s
}

func (s *QueryCubecardFiletokenResponseBodyResultContentDataContent) Validate() error {
	return dara.Validate(s)
}
