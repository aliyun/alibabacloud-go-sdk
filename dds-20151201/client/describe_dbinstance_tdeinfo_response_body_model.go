// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptionKey(v string) *DescribeDBInstanceTDEInfoResponseBody
	GetEncryptionKey() *string
	SetEncryptorName(v string) *DescribeDBInstanceTDEInfoResponseBody
	GetEncryptorName() *string
	SetRequestId(v string) *DescribeDBInstanceTDEInfoResponseBody
	GetRequestId() *string
	SetRoleARN(v string) *DescribeDBInstanceTDEInfoResponseBody
	GetRoleARN() *string
	SetTDEStatus(v string) *DescribeDBInstanceTDEInfoResponseBody
	GetTDEStatus() *string
}

type DescribeDBInstanceTDEInfoResponseBody struct {
	// The custom key of the instance.
	//
	// Currently, only the following regions support Bring Your Own Key (BYOK), which allows you to manage and own encryption keys:
	//
	// - China (Hangzhou)
	//
	// - China (Shanghai)
	//
	// - China (Beijing)
	//
	// - China (Shenzhen)
	//
	// - Hong Kong (China)
	//
	// - Singapore
	//
	// - Malaysia (Kuala Lumpur)
	//
	// > If BYOK is supported, you can manage and own the key, and the system returns your custom key. If BYOK is not supported, you cannot manage the key, and the system returns the string `NoActiveBYOK`.
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption algorithm.
	//
	// example:
	//
	// aes-256-cbc
	EncryptorName *string `json:"EncryptorName,omitempty" xml:"EncryptorName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F4DD0E29-361B-42F2-9301-B0048CCCE5D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The global resource descriptor ARN (Alibaba Cloud Resource Name) of the role pending authorization.
	//
	// example:
	//
	// acs:ram::123456789012****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The TDE enabling status. Valid values:
	//
	// - **enabled**: TDE is enabled.
	//
	// - **disabled**: TDE is disabled.
	//
	// > If the TDE status is disabled, the **RoleARN**, **EncryptionKey**, and **EncryptorName*	- parameters are not returned.
	//
	// example:
	//
	// enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBInstanceTDEInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEInfoResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceTDEInfoResponseBody) GetEncryptorName() *string {
	return s.EncryptorName
}

func (s *DescribeDBInstanceTDEInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceTDEInfoResponseBody) GetRoleARN() *string {
	return s.RoleARN
}

func (s *DescribeDBInstanceTDEInfoResponseBody) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetEncryptionKey(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetEncryptorName(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.EncryptorName = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetRoleARN(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.RoleARN = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) SetTDEStatus(v string) *DescribeDBInstanceTDEInfoResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBInstanceTDEInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
