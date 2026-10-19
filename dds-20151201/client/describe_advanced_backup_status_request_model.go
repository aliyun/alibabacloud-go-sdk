// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAdvancedBackupStatusRequest
	GetDBInstanceId() *string
}

type DescribeAdvancedBackupStatusRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1hyjjqztj5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeAdvancedBackupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAdvancedBackupStatusRequest) SetDBInstanceId(v string) *DescribeAdvancedBackupStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAdvancedBackupStatusRequest) Validate() error {
	return dara.Validate(s)
}
