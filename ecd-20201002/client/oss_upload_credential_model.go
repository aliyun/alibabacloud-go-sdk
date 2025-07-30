// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssUploadCredential interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *OssUploadCredential
	GetAccessKeyId() *string
	SetEndpoint(v string) *OssUploadCredential
	GetEndpoint() *string
	SetFilePath(v string) *OssUploadCredential
	GetFilePath() *string
	SetOssPolicy(v string) *OssUploadCredential
	GetOssPolicy() *string
	SetOssSignature(v string) *OssUploadCredential
	GetOssSignature() *string
	SetStsToken(v string) *OssUploadCredential
	GetStsToken() *string
}

type OssUploadCredential struct {
	AccessKeyId  *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Endpoint     *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FilePath     *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	OssPolicy    *string `json:"OssPolicy,omitempty" xml:"OssPolicy,omitempty"`
	OssSignature *string `json:"OssSignature,omitempty" xml:"OssSignature,omitempty"`
	StsToken     *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s OssUploadCredential) String() string {
	return dara.Prettify(s)
}

func (s OssUploadCredential) GoString() string {
	return s.String()
}

func (s *OssUploadCredential) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *OssUploadCredential) GetEndpoint() *string {
	return s.Endpoint
}

func (s *OssUploadCredential) GetFilePath() *string {
	return s.FilePath
}

func (s *OssUploadCredential) GetOssPolicy() *string {
	return s.OssPolicy
}

func (s *OssUploadCredential) GetOssSignature() *string {
	return s.OssSignature
}

func (s *OssUploadCredential) GetStsToken() *string {
	return s.StsToken
}

func (s *OssUploadCredential) SetAccessKeyId(v string) *OssUploadCredential {
	s.AccessKeyId = &v
	return s
}

func (s *OssUploadCredential) SetEndpoint(v string) *OssUploadCredential {
	s.Endpoint = &v
	return s
}

func (s *OssUploadCredential) SetFilePath(v string) *OssUploadCredential {
	s.FilePath = &v
	return s
}

func (s *OssUploadCredential) SetOssPolicy(v string) *OssUploadCredential {
	s.OssPolicy = &v
	return s
}

func (s *OssUploadCredential) SetOssSignature(v string) *OssUploadCredential {
	s.OssSignature = &v
	return s
}

func (s *OssUploadCredential) SetStsToken(v string) *OssUploadCredential {
	s.StsToken = &v
	return s
}

func (s *OssUploadCredential) Validate() error {
	return dara.Validate(s)
}
