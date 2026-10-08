// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceCLSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *DescribeDBInstanceCLSResponseBody
	GetAlgorithm() *string
	SetEncryptionKey(v string) *DescribeDBInstanceCLSResponseBody
	GetEncryptionKey() *string
	SetEncryptionKeyMode(v string) *DescribeDBInstanceCLSResponseBody
	GetEncryptionKeyMode() *string
	SetRequestId(v string) *DescribeDBInstanceCLSResponseBody
	GetRequestId() *string
	SetWhiteListMode(v bool) *DescribeDBInstanceCLSResponseBody
	GetWhiteListMode() *bool
}

type DescribeDBInstanceCLSResponseBody struct {
	// The encryption algorithm. Valid values:
	//
	// - AES_128_CBC
	//
	// - AES_128_GCM
	//
	// - AES_128_CTR
	//
	// - AES_128_ECB
	//
	// - AES_256_CBC
	//
	// - AES_256_GCM
	//
	// - AES_256_CTR
	//
	// - AES_256_ECB
	//
	// - SM4_128_CBC
	//
	// - SM4_128_GCM
	//
	// - SM4_128_CTR
	//
	// - SM4_128_ECB
	//
	// example:
	//
	// AES_256_GCM
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The custom KMS master key ID.
	//
	// >  This parameter takes effect only when the column encryption key pattern is set to kms_key. If this parameter is not specified, the current column encryption key settings of the database remain unchanged.
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The column encryption key mode. Valid values:
	//
	// - client_key: configures a user-generated random key on the client side.
	//
	// - kms_key: configures a custom key by using Alibaba Cloud Key Management Service (KMS).
	//
	// >  After an instance is configured to use KMS for key management, you can no longer switch back to the client-side random key mode.
	//
	// example:
	//
	// kms_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0073A98-52F1-3075-8256-3943F*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the whitelist mode is enabled.
	//
	// example:
	//
	// true
	WhiteListMode *bool `json:"WhiteListMode,omitempty" xml:"WhiteListMode,omitempty"`
}

func (s DescribeDBInstanceCLSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceCLSResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceCLSResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeDBInstanceCLSResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceCLSResponseBody) GetEncryptionKeyMode() *string {
	return s.EncryptionKeyMode
}

func (s *DescribeDBInstanceCLSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceCLSResponseBody) GetWhiteListMode() *bool {
	return s.WhiteListMode
}

func (s *DescribeDBInstanceCLSResponseBody) SetAlgorithm(v string) *DescribeDBInstanceCLSResponseBody {
	s.Algorithm = &v
	return s
}

func (s *DescribeDBInstanceCLSResponseBody) SetEncryptionKey(v string) *DescribeDBInstanceCLSResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceCLSResponseBody) SetEncryptionKeyMode(v string) *DescribeDBInstanceCLSResponseBody {
	s.EncryptionKeyMode = &v
	return s
}

func (s *DescribeDBInstanceCLSResponseBody) SetRequestId(v string) *DescribeDBInstanceCLSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceCLSResponseBody) SetWhiteListMode(v bool) *DescribeDBInstanceCLSResponseBody {
	s.WhiteListMode = &v
	return s
}

func (s *DescribeDBInstanceCLSResponseBody) Validate() error {
	return dara.Validate(s)
}
