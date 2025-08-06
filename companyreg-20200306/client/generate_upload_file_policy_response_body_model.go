// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *GenerateUploadFilePolicyResponseBody
	GetAccessId() *string
	SetEncodedPolicy(v string) *GenerateUploadFilePolicyResponseBody
	GetEncodedPolicy() *string
	SetExpireTime(v string) *GenerateUploadFilePolicyResponseBody
	GetExpireTime() *string
	SetFileDir(v string) *GenerateUploadFilePolicyResponseBody
	GetFileDir() *string
	SetFileUrl(v string) *GenerateUploadFilePolicyResponseBody
	GetFileUrl() *string
	SetHost(v string) *GenerateUploadFilePolicyResponseBody
	GetHost() *string
	SetRequestId(v string) *GenerateUploadFilePolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *GenerateUploadFilePolicyResponseBody
	GetSignature() *string
}

type GenerateUploadFilePolicyResponseBody struct {
	// OSSAccessKeyId
	//
	// example:
	//
	// hObpgEXoca42qH3V
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMS0xMi0wNlQwNjoxOTowMi40MjdaIiwiY29uZGl0aW9ucyI6W1siZXEiLCIkYnVja2V0IiwibXNlYS1jYWlzaHVpIl1dfQ==
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	// example:
	//
	// 1638169824405
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// company_apply_card/company_change_city/1577930895198750/1638170049657
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// example:
	//
	// https://
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// OSS Endpoint。
	//
	// example:
	//
	// //companyapply.oss-cn-zhangjiakou.aliyuncs.com/
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// EB809CAB-82F7-5843-A42F-356770CD4914
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// jykxhmskIF24sLlxc1GafU/eQMU=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadFilePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *GenerateUploadFilePolicyResponseBody) GetEncodedPolicy() *string {
	return s.EncodedPolicy
}

func (s *GenerateUploadFilePolicyResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GenerateUploadFilePolicyResponseBody) GetFileDir() *string {
	return s.FileDir
}

func (s *GenerateUploadFilePolicyResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GenerateUploadFilePolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *GenerateUploadFilePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUploadFilePolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GenerateUploadFilePolicyResponseBody) SetAccessId(v string) *GenerateUploadFilePolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetEncodedPolicy(v string) *GenerateUploadFilePolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetExpireTime(v string) *GenerateUploadFilePolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileDir(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileUrl(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetHost(v string) *GenerateUploadFilePolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetSignature(v string) *GenerateUploadFilePolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
