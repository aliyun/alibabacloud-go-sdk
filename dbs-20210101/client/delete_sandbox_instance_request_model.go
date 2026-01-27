// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DeleteSandboxInstanceRequest
	GetBackupPlanId() *string
	SetInstanceId(v string) *DeleteSandboxInstanceRequest
	GetInstanceId() *string
	SetZoneId(v string) *DeleteSandboxInstanceRequest
	GetZoneId() *string
}

type DeleteSandboxInstanceRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/437215.html) operation to query the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](https://help.aliyun.com/document_detail/193091.html) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [DescribeSandboxInstances](https://help.aliyun.com/document_detail/437257.html) operation to query the ID of the sandbox instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1jxxxxnxxx1xc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteSandboxInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DeleteSandboxInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSandboxInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteSandboxInstanceRequest) SetBackupPlanId(v string) *DeleteSandboxInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) SetInstanceId(v string) *DeleteSandboxInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) SetZoneId(v string) *DeleteSandboxInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) Validate() error {
	return dara.Validate(s)
}
