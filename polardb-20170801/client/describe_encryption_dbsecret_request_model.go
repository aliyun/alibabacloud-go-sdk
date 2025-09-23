// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEncryptionDBSecretRequest
	GetDBClusterId() *string
}

type DescribeEncryptionDBSecretRequest struct {
	// example:
	//
	// pc-wz9fb5nn44u1d****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeEncryptionDBSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBSecretRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEncryptionDBSecretRequest) SetDBClusterId(v string) *DescribeEncryptionDBSecretRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEncryptionDBSecretRequest) Validate() error {
	return dara.Validate(s)
}
