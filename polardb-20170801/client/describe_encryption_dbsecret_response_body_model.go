// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEncryptionDBSecretResponseBody
	GetDBClusterId() *string
	SetEncryptionDBRegion(v string) *DescribeEncryptionDBSecretResponseBody
	GetEncryptionDBRegion() *string
	SetEncryptionDBStatus(v string) *DescribeEncryptionDBSecretResponseBody
	GetEncryptionDBStatus() *string
	SetEncryptionKey(v string) *DescribeEncryptionDBSecretResponseBody
	GetEncryptionKey() *string
	SetEncryptionKeyStatus(v string) *DescribeEncryptionDBSecretResponseBody
	GetEncryptionKeyStatus() *string
	SetRequestId(v string) *DescribeEncryptionDBSecretResponseBody
	GetRequestId() *string
}

type DescribeEncryptionDBSecretResponseBody struct {
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Enabled
	EncryptionDBRegion *string `json:"EncryptionDBRegion,omitempty" xml:"EncryptionDBRegion,omitempty"`
	// example:
	//
	// cn-beijing
	EncryptionDBStatus *string `json:"EncryptionDBStatus,omitempty" xml:"EncryptionDBStatus,omitempty"`
	// example:
	//
	// 2a4f4ac2-****-****-****-************
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// Disabled
	EncryptionKeyStatus *string `json:"EncryptionKeyStatus,omitempty" xml:"EncryptionKeyStatus,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 944CED46-A6F7-40C6-B6DC-C6E5CC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEncryptionDBSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBSecretResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEncryptionDBSecretResponseBody) GetEncryptionDBRegion() *string {
	return s.EncryptionDBRegion
}

func (s *DescribeEncryptionDBSecretResponseBody) GetEncryptionDBStatus() *string {
	return s.EncryptionDBStatus
}

func (s *DescribeEncryptionDBSecretResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeEncryptionDBSecretResponseBody) GetEncryptionKeyStatus() *string {
	return s.EncryptionKeyStatus
}

func (s *DescribeEncryptionDBSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEncryptionDBSecretResponseBody) SetDBClusterId(v string) *DescribeEncryptionDBSecretResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) SetEncryptionDBRegion(v string) *DescribeEncryptionDBSecretResponseBody {
	s.EncryptionDBRegion = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) SetEncryptionDBStatus(v string) *DescribeEncryptionDBSecretResponseBody {
	s.EncryptionDBStatus = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) SetEncryptionKey(v string) *DescribeEncryptionDBSecretResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) SetEncryptionKeyStatus(v string) *DescribeEncryptionDBSecretResponseBody {
	s.EncryptionKeyStatus = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) SetRequestId(v string) *DescribeEncryptionDBSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEncryptionDBSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
