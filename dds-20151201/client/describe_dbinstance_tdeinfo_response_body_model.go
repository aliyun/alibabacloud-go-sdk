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
	// 实例的自定义密钥。
	//
	// 目前仅以下地域支持BYOK（Bring Your Own Key，用户可以自行管理和拥有加密密钥）：
	//
	// - 华东1（杭州）
	//
	// - 华东2（上海）
	//
	// - 华北2（北京）
	//
	// - 华南1（深圳）
	//
	// - 中国（香港）
	//
	// - 新加坡
	//
	// - 马来西亚（吉隆坡）
	//
	// > 支持BYOK，用户可以管理且拥有密钥，系统将返回用户的自定义密钥；不支持BYOK，用户不可管理密钥，系统将返回字符串`NoActiveBYOK`。
	//
	// example:
	//
	// 2axxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// 加密算法。
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
	// 指定待授权角色的全局资源描述符ARN（Alibaba Cloud Resource Name）信息。
	//
	// example:
	//
	// acs:ram::123456789012****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// The TDE status. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
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
