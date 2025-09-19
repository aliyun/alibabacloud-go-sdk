// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessid(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetAccessid() *string
	SetExpire(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetExpire() *string
	SetHost(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetHost() *string
	SetKey(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetSecurityToken() *string
	SetSignature(v string) *DescribeCustomizedDictUploadInfoResponseBody
	GetSignature() *string
}

type DescribeCustomizedDictUploadInfoResponseBody struct {
	// The AccessKey ID that is required to access the OSS object.
	//
	// example:
	//
	// yourAccessKeyID
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// The time when the OSS signature expires. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1719921470
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// https://aegis-update-static-file.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The key of the OSS object.
	//
	// example:
	//
	// HC_CUSTOMIZED_DICT/176618589410****.tmp
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The policy of the OSS bucket.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNC0wNy0wMlQxMTo1Nzo1MC44MzJaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNDA5NjBdLFsiZXEiLCIka2V5IiwiSENfQ1VTVE9NSVpFRF9ESUNUXC8xNzY2MTg1ODk0MTA0Njc1LnRtc****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDEDEEE7-AC25-559E-8C12-5168B139****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token.
	//
	// example:
	//
	// ***
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The OSS signature.
	//
	// example:
	//
	// mWGRgn0CtdbVf8UuJbTXOmo2****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeCustomizedDictUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetAccessid() *string {
	return s.Accessid
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetExpire() *string {
	return s.Expire
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetAccessid(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Accessid = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetExpire(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetHost(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetKey(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetPolicy(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetRequestId(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetSecurityToken(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) SetSignature(v string) *DescribeCustomizedDictUploadInfoResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeCustomizedDictUploadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
