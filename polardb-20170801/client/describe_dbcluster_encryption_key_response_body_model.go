// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEncryptionKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionKeyList(v []*DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) *DescribeDBClusterEncryptionKeyResponseBody
	GetEncryptionKeyList() []*DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList
	SetRequestId(v string) *DescribeDBClusterEncryptionKeyResponseBody
	GetRequestId() *string
}

type DescribeDBClusterEncryptionKeyResponseBody struct {
	EncryptionKeyList []*DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList `json:"EncryptionKeyList,omitempty" xml:"EncryptionKeyList,omitempty" type:"Repeated"`
	// example:
	//
	// 137ECCC0-920E-5B3B-9F8E-B81632108BBE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterEncryptionKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEncryptionKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEncryptionKeyResponseBody) GetEncryptionKeyList() []*DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	return s.EncryptionKeyList
}

func (s *DescribeDBClusterEncryptionKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterEncryptionKeyResponseBody) SetEncryptionKeyList(v []*DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) *DescribeDBClusterEncryptionKeyResponseBody {
	s.EncryptionKeyList = v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBody) SetRequestId(v string) *DescribeDBClusterEncryptionKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList struct {
	// example:
	//
	// alias/your_default_key
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 1****1
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2026-05-08T08:14:16Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// example:
	//
	// Description of the key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 51858179-afb3-4369-8329-*********
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// Enabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// example:
	//
	// ServiceKey
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// ENCRYPT/DECRYPT
	KeyUsage *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	// example:
	//
	// 2025-10-18T08:14:16Z
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty"`
	// example:
	//
	// Aliyun_KMS
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// DiskEncryption
	UsedBy *string `json:"UsedBy,omitempty" xml:"UsedBy,omitempty"`
}

func (s DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetKeyType() *string {
	return s.KeyType
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetMaterialExpireTime() *string {
	return s.MaterialExpireTime
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) GetUsedBy() *string {
	return s.UsedBy
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetAliasName(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.AliasName = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetCreator(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.Creator = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetDeleteDate(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.DeleteDate = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetDescription(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.Description = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetEncryptionKey(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetEncryptionKeyStatus(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetKeyType(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.KeyType = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetKeyUsage(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.KeyUsage = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetMaterialExpireTime(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetOrigin(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.Origin = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) SetUsedBy(v string) *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList {
	s.UsedBy = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyResponseBodyEncryptionKeyList) Validate() error {
	return dara.Validate(s)
}
