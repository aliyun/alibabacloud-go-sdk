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
	SetEncryptionKeyList(v []*DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) *DescribeDBInstanceEncryptionKeyResponseBody
	GetEncryptionKeyList() []*DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList
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
	// The user who created the key.
	//
	// example:
	//
	// 1443*****9604
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The scheduled time at which the key is deleted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-08T08:14:16Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// Description of the key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 5306d1b6-7fd3-42d9-9511-xxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The details about the key.
	EncryptionKeyList []*DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList `json:"EncryptionKeyList,omitempty" xml:"EncryptionKeyList,omitempty" type:"Repeated"`
	// The status of the key. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// The purpose of the key.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time at which the key expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-18T08:14:16Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The source of the key.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3BC2768E-DEDA-40FC-BBE9-6B884F3626AF
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

func (s *DescribeDBInstanceEncryptionKeyResponseBody) GetEncryptionKeyList() []*DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	return s.EncryptionKeyList
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

func (s *DescribeDBInstanceEncryptionKeyResponseBody) SetEncryptionKeyList(v []*DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) *DescribeDBInstanceEncryptionKeyResponseBody {
	s.EncryptionKeyList = v
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
	if s.EncryptionKeyList != nil {
		for _, item := range s.EncryptionKeyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList struct {
	// The alias of the key.
	//
	// example:
	//
	// alias/xxx
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The user who created the key.
	//
	// example:
	//
	// 1443*****9604
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The scheduled time at which the key is deleted. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-08T08:14:16Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The description of the key.
	//
	// example:
	//
	// Description of the key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 5306d1b6-7fd3-42d9-9511-xxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The status of the key. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// The type of the key. Valid values:
	//
	// 	- **CMK**
	//
	// 	- **ServiceKey**
	//
	// example:
	//
	// ServiceKey
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The purpose of the key.
	//
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// The time at which the key expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-18T08:14:16Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// The source of the key.
	//
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **Master**: primary instance
	//
	// 	- **slave**: read-only instance
	//
	// example:
	//
	// Master
	UsedBy *string `json:"UsedBy,omitempty" xml:"UsedBy,omitempty"`
}

func (s DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) GetUsedBy() *string {
	return s.UsedBy
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetAliasName(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.AliasName = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetCreator(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.Creator = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetDeleteDate(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.DeleteDate = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetDescription(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetEncryptionKey(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetEncryptionKeyStatus(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetKeyType(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.KeyType = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetKeyUsage(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.KeyUsage = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetMaterialExpireTime(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetOrigin(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.Origin = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) SetUsedBy(v string) *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList {
	s.UsedBy = &v
	return s
}

func (s *DescribeDBInstanceEncryptionKeyResponseBodyEncryptionKeyList) Validate() error {
	return dara.Validate(s)
}
