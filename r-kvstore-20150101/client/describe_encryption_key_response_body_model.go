// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *DescribeEncryptionKeyResponseBody
	GetCreator() *string
	SetDeleteDate(v string) *DescribeEncryptionKeyResponseBody
	GetDeleteDate() *string
	SetDescription(v string) *DescribeEncryptionKeyResponseBody
	GetDescription() *string
	SetEncryptionKey(v string) *DescribeEncryptionKeyResponseBody
	GetEncryptionKey() *string
	SetEncryptionKeyStatus(v string) *DescribeEncryptionKeyResponseBody
	GetEncryptionKeyStatus() *string
	SetEncryptionName(v string) *DescribeEncryptionKeyResponseBody
	GetEncryptionName() *string
	SetKeyUsage(v string) *DescribeEncryptionKeyResponseBody
	GetKeyUsage() *string
	SetMaterialExpireTime(v string) *DescribeEncryptionKeyResponseBody
	GetMaterialExpireTime() *string
	SetOrigin(v string) *DescribeEncryptionKeyResponseBody
	GetOrigin() *string
	SetRequestId(v string) *DescribeEncryptionKeyResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *DescribeEncryptionKeyResponseBody
	GetRoleArn() *string
}

type DescribeEncryptionKeyResponseBody struct {
	// The ID of the Alibaba Cloud account that is used to create the custom key.
	//
	// example:
	//
	// 17649847********
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the custom key is expected to be deleted. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > If the return value is an empty string, the custom key cannot be automatically deleted.
	//
	// example:
	//
	// 2021-09-24T18:22:03Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the custom key. By default, an empty string is returned.
	//
	// example:
	//
	// testkey
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom key.
	//
	// example:
	//
	// ad463061-992d-4195-8a94-ed63********
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The state of the custom key. Valid values:
	//
	// 	- **Enabled**: The custom key is available.
	//
	// 	- **Disabled**: The custom key is unavailable.
	//
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// The encryption algorithm.
	//
	// example:
	//
	// AES-CTR-256
	EncryptionName *string `json:"EncryptionName,omitempty" xml:"EncryptionName,omitempty"`
	// The purpose of the custom key. A value of `ENCRYPT/DECRYPT` indicates encryption and decryption.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time when the custom key expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > If the return value is an empty string, the custom key does not expire.
	//
	// example:
	//
	// 2021-09-24T18:22:03Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The source of the custom key. A value of **Aliyun_KMS*	- indicates [Key Management Service (KMS)](https://help.aliyun.com/document_detail/28935.html) of Alibaba Cloud.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9A931CE5-C926-5E09-B0EC-6299C4A6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the Resource Access Management (RAM) role to which you want to grant permissions.
	//
	// example:
	//
	// acs:ram::123456789012****:role/AliyunRdsInstanceEncryptionDefaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s DescribeEncryptionKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionKeyResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *DescribeEncryptionKeyResponseBody) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *DescribeEncryptionKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeEncryptionKeyResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeEncryptionKeyResponseBody) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeEncryptionKeyResponseBody) GetEncryptionName() *string {
	return s.EncryptionName
}

func (s *DescribeEncryptionKeyResponseBody) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *DescribeEncryptionKeyResponseBody) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *DescribeEncryptionKeyResponseBody) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeEncryptionKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEncryptionKeyResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeEncryptionKeyResponseBody) SetCreator(v string) *DescribeEncryptionKeyResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetDeleteDate(v string) *DescribeEncryptionKeyResponseBody {
	s.DeleteDate = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetDescription(v string) *DescribeEncryptionKeyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetEncryptionKey(v string) *DescribeEncryptionKeyResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetEncryptionKeyStatus(v string) *DescribeEncryptionKeyResponseBody {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetEncryptionName(v string) *DescribeEncryptionKeyResponseBody {
	s.EncryptionName = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetKeyUsage(v string) *DescribeEncryptionKeyResponseBody {
	s.KeyUsage = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetMaterialExpireTime(v string) *DescribeEncryptionKeyResponseBody {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetOrigin(v string) *DescribeEncryptionKeyResponseBody {
	s.Origin = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetRequestId(v string) *DescribeEncryptionKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) SetRoleArn(v string) *DescribeEncryptionKeyResponseBody {
	s.RoleArn = &v
	return s
}

func (s *DescribeEncryptionKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
