// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationOssUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessid(v string) *GetOperationOssUploadPolicyResponseBody
	GetAccessid() *string
	SetEncodedPolicy(v string) *GetOperationOssUploadPolicyResponseBody
	GetEncodedPolicy() *string
	SetExpireTime(v string) *GetOperationOssUploadPolicyResponseBody
	GetExpireTime() *string
	SetFileDir(v string) *GetOperationOssUploadPolicyResponseBody
	GetFileDir() *string
	SetHost(v string) *GetOperationOssUploadPolicyResponseBody
	GetHost() *string
	SetRequestId(v string) *GetOperationOssUploadPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *GetOperationOssUploadPolicyResponseBody
	GetSignature() *string
}

type GetOperationOssUploadPolicyResponseBody struct {
	// example:
	//
	// hObpgEXoca42****
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAaMC0wNy0wMlQxKToyMDoxMS44ODRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNTI0Mjg4MDBdLFsic3RhcnRzLXdpdGgiLCIka2V5IiwiMTIxOTU0MTE2MTIxMzA1Ny9PRkZMSU5FX1RSQU5TRkVSLzE1OTM2ODg1MTE4ODMi****
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	// example:
	//
	// 1593688811881
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1219541161213157/OFFLINE_TRANSFER/159368851****
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// OSS Endpoint。
	//
	// example:
	//
	// //***-basic-cert.oss-cn-***.aliyuncs.com/
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-40EC-8035-4B12FEFD7D011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// pNVECGkyL0tl4bKXekV5ErZ****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetOperationOssUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationOssUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationOssUploadPolicyResponseBody) GetAccessid() *string {
	return s.Accessid
}

func (s *GetOperationOssUploadPolicyResponseBody) GetEncodedPolicy() *string {
	return s.EncodedPolicy
}

func (s *GetOperationOssUploadPolicyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetOperationOssUploadPolicyResponseBody) GetFileDir() *string {
	return s.FileDir
}

func (s *GetOperationOssUploadPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetOperationOssUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationOssUploadPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GetOperationOssUploadPolicyResponseBody) SetAccessid(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Accessid = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetEncodedPolicy(v string) *GetOperationOssUploadPolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetExpireTime(v string) *GetOperationOssUploadPolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetFileDir(v string) *GetOperationOssUploadPolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetHost(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetRequestId(v string) *GetOperationOssUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) SetSignature(v string) *GetOperationOssUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GetOperationOssUploadPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
