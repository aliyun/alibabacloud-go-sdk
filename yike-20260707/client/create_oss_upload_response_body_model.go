// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetObjectKey(v string) *CreateOssUploadResponseBody
	GetObjectKey() *string
	SetOssUrl(v string) *CreateOssUploadResponseBody
	GetOssUrl() *string
	SetRequestId(v string) *CreateOssUploadResponseBody
	GetRequestId() *string
	SetUploadCredentials(v *CreateOssUploadResponseBodyUploadCredentials) *CreateOssUploadResponseBody
	GetUploadCredentials() *CreateOssUploadResponseBodyUploadCredentials
}

type CreateOssUploadResponseBody struct {
	// OSS ObjectKey
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The **OSS presigned GET URL**, which is valid for 24 hours and uses the public endpoint. If signature generation fails, null is returned, which does not affect direct upload with the credentials.
	//
	// example:
	//
	// https://yike-shared-pre-2.oss-cn-shanghai.aliyuncs.com/yike-v1/ge2tqn-o5zv6m/2026/11/18/13/upload/media/hkyo9v003.mp4?Expires=1795095160&OSSAccessKeyId=STS.Nxxxxxxxxxxxxxxxxxxxxxxxx&Signature=xxxxxxxxxxx%3D
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The **request ID**.
	//
	// example:
	//
	// F99F2AE5-***-BBD18F101455
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The upload credentials. Some fields may be empty depending on the CredentialType.
	UploadCredentials *CreateOssUploadResponseBodyUploadCredentials `json:"UploadCredentials,omitempty" xml:"UploadCredentials,omitempty" type:"Struct"`
}

func (s CreateOssUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOssUploadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssUploadResponseBody) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *CreateOssUploadResponseBody) GetOssUrl() *string {
	return s.OssUrl
}

func (s *CreateOssUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOssUploadResponseBody) GetUploadCredentials() *CreateOssUploadResponseBodyUploadCredentials {
	return s.UploadCredentials
}

func (s *CreateOssUploadResponseBody) SetObjectKey(v string) *CreateOssUploadResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *CreateOssUploadResponseBody) SetOssUrl(v string) *CreateOssUploadResponseBody {
	s.OssUrl = &v
	return s
}

func (s *CreateOssUploadResponseBody) SetRequestId(v string) *CreateOssUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssUploadResponseBody) SetUploadCredentials(v *CreateOssUploadResponseBodyUploadCredentials) *CreateOssUploadResponseBody {
	s.UploadCredentials = v
	return s
}

func (s *CreateOssUploadResponseBody) Validate() error {
	if s.UploadCredentials != nil {
		if err := s.UploadCredentials.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOssUploadResponseBodyUploadCredentials struct {
	// The STS temporary AccessKey ID.
	//
	// example:
	//
	// STS.Nxxxxxxxxxxxxxxxxxxxxxxxx
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The STS temporary AccessKey secret.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// OSS Bucket
	//
	// example:
	//
	// yike-shared-cn-shanghai-1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The expiration time of the credentials. The time is in UTC and follows the ISO 8601 standard. The actual return value prevails.
	//
	// example:
	//
	// 2026-11-20T01:32:40Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The PostObject policy, which is Base64-encoded. This field is returned when POST_POLICY is used.
	//
	// example:
	//
	// -
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The AccessKey ID used for signing. This field is returned when POST_POLICY is used.
	//
	// example:
	//
	// -
	PostAccessKeyId *string `json:"PostAccessKeyId,omitempty" xml:"PostAccessKeyId,omitempty"`
	// The destination endpoint for POST requests. This field is returned when POST_POLICY is used.
	//
	// example:
	//
	// -
	PostEndpoint *string `json:"PostEndpoint,omitempty" xml:"PostEndpoint,omitempty"`
	// The OSS region, such as cn-shanghai.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The STS security token.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The PostObject signature. This field is returned when POST_POLICY is used.
	//
	// example:
	//
	// -
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The presigned PUT URL. This field is returned when PRESIGNED_URL is used.
	//
	// example:
	//
	// -
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s CreateOssUploadResponseBodyUploadCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateOssUploadResponseBodyUploadCredentials) GoString() string {
	return s.String()
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetBucket() *string {
	return s.Bucket
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetPolicy() *string {
	return s.Policy
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetPostAccessKeyId() *string {
	return s.PostAccessKeyId
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetPostEndpoint() *string {
	return s.PostEndpoint
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetRegion() *string {
	return s.Region
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetSignature() *string {
	return s.Signature
}

func (s *CreateOssUploadResponseBodyUploadCredentials) GetSignedUrl() *string {
	return s.SignedUrl
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetAccessKeyId(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetAccessKeySecret(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetBucket(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.Bucket = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetExpiration(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.Expiration = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetPolicy(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.Policy = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetPostAccessKeyId(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.PostAccessKeyId = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetPostEndpoint(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.PostEndpoint = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetRegion(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.Region = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetSecurityToken(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.SecurityToken = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetSignature(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.Signature = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) SetSignedUrl(v string) *CreateOssUploadResponseBodyUploadCredentials {
	s.SignedUrl = &v
	return s
}

func (s *CreateOssUploadResponseBodyUploadCredentials) Validate() error {
	return dara.Validate(s)
}
