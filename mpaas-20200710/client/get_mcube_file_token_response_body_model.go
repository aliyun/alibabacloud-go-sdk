// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeFileTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetFileTokenResult(v *GetMcubeFileTokenResponseBodyGetFileTokenResult) *GetMcubeFileTokenResponseBody
	GetGetFileTokenResult() *GetMcubeFileTokenResponseBodyGetFileTokenResult
	SetRequestId(v string) *GetMcubeFileTokenResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMcubeFileTokenResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetMcubeFileTokenResponseBody
	GetResultMessage() *string
}

type GetMcubeFileTokenResponseBody struct {
	GetFileTokenResult *GetMcubeFileTokenResponseBodyGetFileTokenResult `json:"GetFileTokenResult,omitempty" xml:"GetFileTokenResult,omitempty" type:"Struct"`
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMcubeFileTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeFileTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcubeFileTokenResponseBody) GetGetFileTokenResult() *GetMcubeFileTokenResponseBodyGetFileTokenResult {
	return s.GetFileTokenResult
}

func (s *GetMcubeFileTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeFileTokenResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMcubeFileTokenResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMcubeFileTokenResponseBody) SetGetFileTokenResult(v *GetMcubeFileTokenResponseBodyGetFileTokenResult) *GetMcubeFileTokenResponseBody {
	s.GetFileTokenResult = v
	return s
}

func (s *GetMcubeFileTokenResponseBody) SetRequestId(v string) *GetMcubeFileTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcubeFileTokenResponseBody) SetResultCode(v string) *GetMcubeFileTokenResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMcubeFileTokenResponseBody) SetResultMessage(v string) *GetMcubeFileTokenResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMcubeFileTokenResponseBody) Validate() error {
	if s.GetFileTokenResult != nil {
		if err := s.GetFileTokenResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcubeFileTokenResponseBodyGetFileTokenResult struct {
	FileToken *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken `json:"FileToken,omitempty" xml:"FileToken,omitempty" type:"Struct"`
	ResultMsg *string                                                   `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcubeFileTokenResponseBodyGetFileTokenResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeFileTokenResponseBodyGetFileTokenResult) GoString() string {
	return s.String()
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) GetFileToken() *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	return s.FileToken
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) SetFileToken(v *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) *GetMcubeFileTokenResponseBodyGetFileTokenResult {
	s.FileToken = v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) SetResultMsg(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResult {
	s.ResultMsg = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) SetSuccess(v bool) *GetMcubeFileTokenResponseBodyGetFileTokenResult {
	s.Success = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResult) Validate() error {
	if s.FileToken != nil {
		if err := s.FileToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Dir       *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GoString() string {
	return s.String()
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetAccessid() *string {
	return s.Accessid
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetDir() *string {
	return s.Dir
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetExpire() *string {
	return s.Expire
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetHost() *string {
	return s.Host
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetPolicy() *string {
	return s.Policy
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) GetSignature() *string {
	return s.Signature
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetAccessid(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Accessid = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetDir(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Dir = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetExpire(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Expire = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetHost(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Host = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetPolicy(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Policy = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) SetSignature(v string) *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken {
	s.Signature = &v
	return s
}

func (s *GetMcubeFileTokenResponseBodyGetFileTokenResultFileToken) Validate() error {
	return dara.Validate(s)
}
