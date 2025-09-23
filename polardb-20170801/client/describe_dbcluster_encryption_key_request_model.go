// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEncryptionKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBClusterEncryptionKeyRequest
	GetClientToken() *string
	SetDBClusterId(v string) *DescribeDBClusterEncryptionKeyRequest
	GetDBClusterId() *string
}

type DescribeDBClusterEncryptionKeyRequest struct {
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// pc-***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterEncryptionKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEncryptionKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEncryptionKeyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBClusterEncryptionKeyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterEncryptionKeyRequest) SetClientToken(v string) *DescribeDBClusterEncryptionKeyRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyRequest) SetDBClusterId(v string) *DescribeDBClusterEncryptionKeyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEncryptionKeyRequest) Validate() error {
	return dara.Validate(s)
}
