// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSInfoForCardTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOSSInfoForCardTemplateResponseBody
	GetCode() *string
	SetData(v *GetOSSInfoForCardTemplateResponseBodyData) *GetOSSInfoForCardTemplateResponseBody
	GetData() *GetOSSInfoForCardTemplateResponseBodyData
	SetRequestId(v string) *GetOSSInfoForCardTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOSSInfoForCardTemplateResponseBody
	GetSuccess() *bool
}

type GetOSSInfoForCardTemplateResponseBody struct {
	// 请求状态码。
	//
	// - OK：代表请求成功。
	//
	// - 其他错误码，请参见[API错误码](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回数据。
	Data *GetOSSInfoForCardTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 请求ID。
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 调用接口是否成功。取值：
	//
	// - **true**：调用成功。
	//
	// - **false**：调用失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOSSInfoForCardTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOSSInfoForCardTemplateResponseBody) GetData() *GetOSSInfoForCardTemplateResponseBodyData {
	return s.Data
}

func (s *GetOSSInfoForCardTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOSSInfoForCardTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetCode(v string) *GetOSSInfoForCardTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetData(v *GetOSSInfoForCardTemplateResponseBodyData) *GetOSSInfoForCardTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetRequestId(v string) *GetOSSInfoForCardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) SetSuccess(v bool) *GetOSSInfoForCardTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOSSInfoForCardTemplateResponseBodyData struct {
	// 签名使用的AccessKey ID。
	//
	// example:
	//
	// LTAI************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// 阿里云账号ID。
	//
	// example:
	//
	// 168**********184
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// OSS文件保存桶名称。
	//
	// example:
	//
	// alicom-cardsms-resources
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// 超时时间戳。单位：秒。
	//
	// example:
	//
	// 1634209418
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 访问地址。
	//
	// example:
	//
	// http://***.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 签名策略。
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0wMy0yNlQwMzo0NDoyMy4xNTlaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCIxNDIwNjM0******************
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// 短信签名。
	//
	// example:
	//
	// 阿里云
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// 策略路径。
	//
	// example:
	//
	// 168**********184
	StartPath *string `json:"StartPath,omitempty" xml:"StartPath,omitempty"`
}

func (s GetOSSInfoForCardTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOSSInfoForCardTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetAliUid() *string {
	return s.AliUid
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) GetStartPath() *string {
	return s.StartPath
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetAccessKeyId(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetAliUid(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetBucket(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetExpireTime(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetHost(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetPolicy(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetSignature(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) SetStartPath(v string) *GetOSSInfoForCardTemplateResponseBodyData {
	s.StartPath = &v
	return s
}

func (s *GetOSSInfoForCardTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
