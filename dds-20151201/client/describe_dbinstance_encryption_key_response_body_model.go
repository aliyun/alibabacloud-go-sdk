// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEncryptionKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetCreator() *string
	SetDeleteDate(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetDeleteDate() *string
	SetDescription(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetDescription() *string
	SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetEncryptionKey() *string
	SetEncryptionKeyStatus(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetEncryptionKeyStatus() *string
	SetKeyUsage(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetKeyUsage() *string
	SetMaterialExpireTime(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetMaterialExpireTime() *string
	SetOrigin(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetOrigin() *string
	SetRequestId(v string) *DescribeDBInstanceEncryptionKeyResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceEncryptionKeyResponseBody struct {
	// The UID of the key creator.
	//
	// example:
	//
	// 123456
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The scheduled time when the key for the instance will be deleted. If the parameter is left empty, the key will not be deleted.
	//
	// example:
	//
	// 2020-07-06T18:22:03Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the key for the instance.
	//
	// example:
	//
	// key description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key for the instance.
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// Indicates whether the key for the instance is enabled. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// The purpose of the key for the instance.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The expiration time of the key for the instance. The time is displayed in UTC. If the parameter is left empty, the key for the instance will not expire.
	//
	// example:
	//
	// 2020-07-06T18:22:03Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The source of the key for the instance.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 783C2062-A2D3-4EA8-88AD-E43F990C23BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetCreator(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Creator = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetDeleteDate(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.DeleteDate = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetDescription(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetEncryptionKeyStatus(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetKeyUsage(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.KeyUsage = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetMaterialExpireTime(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetOrigin(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.Origin = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetRequestId(v string) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
