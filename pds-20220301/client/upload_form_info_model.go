// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFormInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *UploadFormInfo
	GetBucketName() *string
	SetEndpoint(v string) *UploadFormInfo
	GetEndpoint() *string
	SetFormData(v map[string]*string) *UploadFormInfo
	GetFormData() map[string]*string
	SetObjectKey(v string) *UploadFormInfo
	GetObjectKey() *string
	SetOssAccessKeyId(v string) *UploadFormInfo
	GetOssAccessKeyId() *string
	SetOssEndPoint(v string) *UploadFormInfo
	GetOssEndPoint() *string
	SetOssSecurityToken(v string) *UploadFormInfo
	GetOssSecurityToken() *string
	SetPolicy(v string) *UploadFormInfo
	GetPolicy() *string
	SetSignature(v string) *UploadFormInfo
	GetSignature() *string
}

type UploadFormInfo struct {
	BucketName       *string            `json:"bucket_name,omitempty" xml:"bucket_name,omitempty"`
	Endpoint         *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FormData         map[string]*string `json:"form_data,omitempty" xml:"form_data,omitempty"`
	ObjectKey        *string            `json:"object_key,omitempty" xml:"object_key,omitempty"`
	OssAccessKeyId   *string            `json:"oss_access_key_id,omitempty" xml:"oss_access_key_id,omitempty"`
	OssEndPoint      *string            `json:"oss_end_point,omitempty" xml:"oss_end_point,omitempty"`
	OssSecurityToken *string            `json:"oss_security_token,omitempty" xml:"oss_security_token,omitempty"`
	Policy           *string            `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature        *string            `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s UploadFormInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadFormInfo) GoString() string {
	return s.String()
}

func (s *UploadFormInfo) GetBucketName() *string {
	return s.BucketName
}

func (s *UploadFormInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UploadFormInfo) GetFormData() map[string]*string {
	return s.FormData
}

func (s *UploadFormInfo) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *UploadFormInfo) GetOssAccessKeyId() *string {
	return s.OssAccessKeyId
}

func (s *UploadFormInfo) GetOssEndPoint() *string {
	return s.OssEndPoint
}

func (s *UploadFormInfo) GetOssSecurityToken() *string {
	return s.OssSecurityToken
}

func (s *UploadFormInfo) GetPolicy() *string {
	return s.Policy
}

func (s *UploadFormInfo) GetSignature() *string {
	return s.Signature
}

func (s *UploadFormInfo) SetBucketName(v string) *UploadFormInfo {
	s.BucketName = &v
	return s
}

func (s *UploadFormInfo) SetEndpoint(v string) *UploadFormInfo {
	s.Endpoint = &v
	return s
}

func (s *UploadFormInfo) SetFormData(v map[string]*string) *UploadFormInfo {
	s.FormData = v
	return s
}

func (s *UploadFormInfo) SetObjectKey(v string) *UploadFormInfo {
	s.ObjectKey = &v
	return s
}

func (s *UploadFormInfo) SetOssAccessKeyId(v string) *UploadFormInfo {
	s.OssAccessKeyId = &v
	return s
}

func (s *UploadFormInfo) SetOssEndPoint(v string) *UploadFormInfo {
	s.OssEndPoint = &v
	return s
}

func (s *UploadFormInfo) SetOssSecurityToken(v string) *UploadFormInfo {
	s.OssSecurityToken = &v
	return s
}

func (s *UploadFormInfo) SetPolicy(v string) *UploadFormInfo {
	s.Policy = &v
	return s
}

func (s *UploadFormInfo) SetSignature(v string) *UploadFormInfo {
	s.Signature = &v
	return s
}

func (s *UploadFormInfo) Validate() error {
	return dara.Validate(s)
}
