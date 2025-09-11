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
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetOSSInfoForCardTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	return dara.Validate(s)
}

type GetOSSInfoForCardTemplateResponseBodyData struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAIxetqt1Dg****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 599333677478****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// alicom-cardsms-resources
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 1634209418
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The hostname.
	//
	// example:
	//
	// https://alicom-cardsms-resources.oss-cn-zhangjiakou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The signature policy.
	//
	// example:
	//
	// eyJxxx0=
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The signature.
	//
	// example:
	//
	// Aliyun
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The path of the policy.
	//
	// example:
	//
	// 1631792777
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
