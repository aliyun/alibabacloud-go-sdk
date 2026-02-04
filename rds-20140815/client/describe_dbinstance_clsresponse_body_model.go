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
	// example:
	//
	// AES_256_GCM
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// acs:kms:cn-hangzhou:123456789:key/xxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// KMS
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// example:
	//
	// D0073A98-52F1-3075-8256-3943F*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
