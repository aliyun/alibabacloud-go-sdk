// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteBackupPolicyRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteBackupPolicyRequest
	GetRegionId() *string
}

type DeleteBackupPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupPolicyRequest) SetDBInstanceId(v string) *DeleteBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) SetRegionId(v string) *DeleteBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
