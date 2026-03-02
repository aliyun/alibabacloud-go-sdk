// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileSignatureResult interface {
	dara.Model
	String() string
	GoString() string
	SetExpiredTime(v int64) *UploadFileSignatureResult
	GetExpiredTime() *int64
	SetFileUrl(v string) *UploadFileSignatureResult
	GetFileUrl() *string
	SetHost(v string) *UploadFileSignatureResult
	GetHost() *string
	SetKeyId(v string) *UploadFileSignatureResult
	GetKeyId() *string
	SetPath(v string) *UploadFileSignatureResult
	GetPath() *string
	SetPolicy(v string) *UploadFileSignatureResult
	GetPolicy() *string
	SetSignature(v string) *UploadFileSignatureResult
	GetSignature() *string
	SetThumbUrl(v string) *UploadFileSignatureResult
	GetThumbUrl() *string
}

type UploadFileSignatureResult struct {
	ExpiredTime *int64  `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	FileUrl     *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	Host        *string `json:"host,omitempty" xml:"host,omitempty"`
	KeyId       *string `json:"keyId,omitempty" xml:"keyId,omitempty"`
	Path        *string `json:"path,omitempty" xml:"path,omitempty"`
	Policy      *string `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
	ThumbUrl    *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
}

func (s UploadFileSignatureResult) String() string {
	return dara.Prettify(s)
}

func (s UploadFileSignatureResult) GoString() string {
	return s.String()
}

func (s *UploadFileSignatureResult) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *UploadFileSignatureResult) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadFileSignatureResult) GetHost() *string {
	return s.Host
}

func (s *UploadFileSignatureResult) GetKeyId() *string {
	return s.KeyId
}

func (s *UploadFileSignatureResult) GetPath() *string {
	return s.Path
}

func (s *UploadFileSignatureResult) GetPolicy() *string {
	return s.Policy
}

func (s *UploadFileSignatureResult) GetSignature() *string {
	return s.Signature
}

func (s *UploadFileSignatureResult) GetThumbUrl() *string {
	return s.ThumbUrl
}

func (s *UploadFileSignatureResult) SetExpiredTime(v int64) *UploadFileSignatureResult {
	s.ExpiredTime = &v
	return s
}

func (s *UploadFileSignatureResult) SetFileUrl(v string) *UploadFileSignatureResult {
	s.FileUrl = &v
	return s
}

func (s *UploadFileSignatureResult) SetHost(v string) *UploadFileSignatureResult {
	s.Host = &v
	return s
}

func (s *UploadFileSignatureResult) SetKeyId(v string) *UploadFileSignatureResult {
	s.KeyId = &v
	return s
}

func (s *UploadFileSignatureResult) SetPath(v string) *UploadFileSignatureResult {
	s.Path = &v
	return s
}

func (s *UploadFileSignatureResult) SetPolicy(v string) *UploadFileSignatureResult {
	s.Policy = &v
	return s
}

func (s *UploadFileSignatureResult) SetSignature(v string) *UploadFileSignatureResult {
	s.Signature = &v
	return s
}

func (s *UploadFileSignatureResult) SetThumbUrl(v string) *UploadFileSignatureResult {
	s.ThumbUrl = &v
	return s
}

func (s *UploadFileSignatureResult) Validate() error {
	return dara.Validate(s)
}
