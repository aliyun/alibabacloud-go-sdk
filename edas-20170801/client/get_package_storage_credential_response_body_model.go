// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageStorageCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPackageStorageCredentialResponseBody
	GetCode() *int32
	SetCredential(v *GetPackageStorageCredentialResponseBodyCredential) *GetPackageStorageCredentialResponseBody
	GetCredential() *GetPackageStorageCredentialResponseBodyCredential
	SetMessage(v string) *GetPackageStorageCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPackageStorageCredentialResponseBody
	GetRequestId() *string
}

type GetPackageStorageCredentialResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The STS credential.
	Credential *GetPackageStorageCredentialResponseBodyCredential `json:"Credential,omitempty" xml:"Credential,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPackageStorageCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPackageStorageCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPackageStorageCredentialResponseBody) GetCredential() *GetPackageStorageCredentialResponseBodyCredential {
	return s.Credential
}

func (s *GetPackageStorageCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPackageStorageCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPackageStorageCredentialResponseBody) SetCode(v int32) *GetPackageStorageCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetCredential(v *GetPackageStorageCredentialResponseBodyCredential) *GetPackageStorageCredentialResponseBody {
	s.Credential = v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetMessage(v string) *GetPackageStorageCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) SetRequestId(v string) *GetPackageStorageCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBody) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPackageStorageCredentialResponseBodyCredential struct {
	// The AccessKey ID of your account.
	//
	// example:
	//
	// <yourAccessKeyId>
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret of your account.
	//
	// example:
	//
	// <yourAccessKeySecret>
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// edas-bj
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The time when the STS credential expires. Example: 2019-11-10T07:20:19Z.
	//
	// example:
	//
	// 2019-11-10T07:20:19Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The object key prefix in Object Storage Service (OSS).
	//
	// example:
	//
	// release-pkg/117274586608****
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The private endpoint of OSS.
	//
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	OssInternalEndpoint *string `json:"OssInternalEndpoint,omitempty" xml:"OssInternalEndpoint,omitempty"`
	// The public endpoint of OSS.
	//
	// example:
	//
	// oss-cn-beijing.aliyuncs.com
	OssPublicEndpoint *string `json:"OssPublicEndpoint,omitempty" xml:"OssPublicEndpoint,omitempty"`
	// The VPC endpoint of OSS.
	//
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	OssVpcEndpoint *string `json:"OssVpcEndpoint,omitempty" xml:"OssVpcEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security token issued by STS.
	//
	// example:
	//
	// <yourSecurityToken>
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetPackageStorageCredentialResponseBodyCredential) String() string {
	return dara.Prettify(s)
}

func (s GetPackageStorageCredentialResponseBodyCredential) GoString() string {
	return s.String()
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetBucket() *string {
	return s.Bucket
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetExpiration() *string {
	return s.Expiration
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetOssInternalEndpoint() *string {
	return s.OssInternalEndpoint
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetOssPublicEndpoint() *string {
	return s.OssPublicEndpoint
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetOssVpcEndpoint() *string {
	return s.OssVpcEndpoint
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPackageStorageCredentialResponseBodyCredential) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetAccessKeyId(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.AccessKeyId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetAccessKeySecret(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.AccessKeySecret = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetBucket(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.Bucket = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetExpiration(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.Expiration = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetKeyPrefix(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.KeyPrefix = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetOssInternalEndpoint(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.OssInternalEndpoint = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetOssPublicEndpoint(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.OssPublicEndpoint = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetOssVpcEndpoint(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.OssVpcEndpoint = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetRegionId(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.RegionId = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) SetSecurityToken(v string) *GetPackageStorageCredentialResponseBodyCredential {
	s.SecurityToken = &v
	return s
}

func (s *GetPackageStorageCredentialResponseBodyCredential) Validate() error {
	return dara.Validate(s)
}
