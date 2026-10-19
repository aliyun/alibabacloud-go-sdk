// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAdvancedBackupPolicyRequest
	GetDBInstanceId() *string
}

type DescribeAdvancedBackupPolicyRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp18b0934e70****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAdvancedBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAdvancedBackupPolicyRequest) SetDBInstanceId(v string) *DescribeAdvancedBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAdvancedBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
