// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicDictUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessid(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetAccessid() *string
	SetExpire(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetExpire() *string
	SetHost(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetHost() *string
	SetKey(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribeDynamicDictUploadInfoResponseBody
	GetSignature() *string
}

type DescribeDynamicDictUploadInfoResponseBody struct {
	// The AccessKey ID that is used to access OSS.
	//
	// example:
	//
	// yourAccessKeyID
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// The validity period of the signature. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1719919893
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// https://aegis-update-static-file.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the OSS object.
	//
	// example:
	//
	// DegradePool_Offset_****
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The OSS security policy.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNC0wNy0wMlQxMTozMTozMy40MjlaIiwiY29uZGl0aW9********
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature that is calculated based on **AccessKeySecret*	- and **Policy**. When you call an OSS API operation, OSS uses the signature information to check the validity of the POST request.
	//
	// example:
	//
	// wBiwkhd5LGcLzijtc3FhI****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeDynamicDictUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicDictUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetAccessid() *string {
	return s.Accessid
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetExpire() *string {
	return s.Expire
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDynamicDictUploadInfoResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetAccessid(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Accessid = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetExpire(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetHost(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetKey(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetPolicy(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetRequestId(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) SetSignature(v string) *DescribeDynamicDictUploadInfoResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeDynamicDictUploadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
