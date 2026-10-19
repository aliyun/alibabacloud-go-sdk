// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyAdvancedBackupStatusRequest
	GetDBInstanceId() *string
}

type ModifyAdvancedBackupStatusRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp114f14849d****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyAdvancedBackupStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAdvancedBackupStatusRequest) SetDBInstanceId(v string) *ModifyAdvancedBackupStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAdvancedBackupStatusRequest) Validate() error {
	return dara.Validate(s)
}
