// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUploadPolicyResponseBody
	GetCode() *string
	SetData(v *GetUploadPolicyResponseBodyData) *GetUploadPolicyResponseBody
	GetData() *GetUploadPolicyResponseBodyData
	SetHttpStatusCode(v int32) *GetUploadPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUploadPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUploadPolicyResponseBody
	GetSuccess() *bool
}

type GetUploadPolicyResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetUploadPolicyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 90C68329-A75C-5449-A928-4D0BAD7AA0FA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUploadPolicyResponseBody) GetData() *GetUploadPolicyResponseBodyData {
	return s.Data
}

func (s *GetUploadPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUploadPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUploadPolicyResponseBody) SetCode(v string) *GetUploadPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadPolicyResponseBody) SetData(v *GetUploadPolicyResponseBodyData) *GetUploadPolicyResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadPolicyResponseBody) SetHttpStatusCode(v int32) *GetUploadPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadPolicyResponseBody) SetMessage(v string) *GetUploadPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadPolicyResponseBody) SetRequestId(v string) *GetUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadPolicyResponseBody) SetSuccess(v bool) *GetUploadPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetUploadPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUploadPolicyResponseBodyData struct {
	// example:
	//
	// material/INPUT_CHAT_BACKGROUND_PIC/Mt.CPQHBSWQS5QU2
	OssKey    *string                                   `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
	OssPolicy *GetUploadPolicyResponseBodyDataOssPolicy `json:"ossPolicy,omitempty" xml:"ossPolicy,omitempty" type:"Struct"`
}

func (s GetUploadPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUploadPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponseBodyData) GetOssKey() *string {
	return s.OssKey
}

func (s *GetUploadPolicyResponseBodyData) GetOssPolicy() *GetUploadPolicyResponseBodyDataOssPolicy {
	return s.OssPolicy
}

func (s *GetUploadPolicyResponseBodyData) SetOssKey(v string) *GetUploadPolicyResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *GetUploadPolicyResponseBodyData) SetOssPolicy(v *GetUploadPolicyResponseBodyDataOssPolicy) *GetUploadPolicyResponseBodyData {
	s.OssPolicy = v
	return s
}

func (s *GetUploadPolicyResponseBodyData) Validate() error {
	if s.OssPolicy != nil {
		if err := s.OssPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUploadPolicyResponseBodyDataOssPolicy struct {
	// accessId。
	//
	// example:
	//
	// LTBI5123ADDJdsadCidYA8kw3
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// example:
	//
	// material/INPUT_CHAT_BACKGROUND_PIC/Mt.CPQHBSWQS5QU2/
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// example:
	//
	// 1761551667
	Expire *string `json:"expire,omitempty" xml:"expire,omitempty"`
	// example:
	//
	// daily-avatar-property.oss-cn-beijing.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMi0yNVQwMzowMDoyNC4xNDNaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNTM2ODcwOTEyMF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJ0ZW1wXC8xNzQwNDQ4ODI0MTQxLnppcCJdXX0=
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// example:
	//
	// I2KcV3CFloyRr94WhefmVEuNiv0=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetUploadPolicyResponseBodyDataOssPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetUploadPolicyResponseBodyDataOssPolicy) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetAccessId() *string {
	return s.AccessId
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetDir() *string {
	return s.Dir
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetExpire() *string {
	return s.Expire
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetHost() *string {
	return s.Host
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetPolicy() *string {
	return s.Policy
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) GetSignature() *string {
	return s.Signature
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetAccessId(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.AccessId = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetDir(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.Dir = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetExpire(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.Expire = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetHost(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.Host = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetPolicy(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.Policy = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) SetSignature(v string) *GetUploadPolicyResponseBodyDataOssPolicy {
	s.Signature = &v
	return s
}

func (s *GetUploadPolicyResponseBodyDataOssPolicy) Validate() error {
	return dara.Validate(s)
}
