// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeSandboxInstancesRequest
	GetBackupPlanId() *string
	SetInstanceId(v string) *DescribeSandboxInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v string) *DescribeSandboxInstancesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeSandboxInstancesRequest
	GetPageSize() *string
}

type DescribeSandboxInstancesRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/437215.html) operation to obtain the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](https://help.aliyun.com/document_detail/193091.html) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [CreateSandboxInstance](https://help.aliyun.com/document_detail/437252.html) operation to obtain the ID of the sandbox instance.
	//
	// example:
	//
	// 1jxxxxnxxx1xc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30\\. This is the default value.
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeSandboxInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSandboxInstancesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSandboxInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeSandboxInstancesRequest) SetBackupPlanId(v string) *DescribeSandboxInstancesRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetInstanceId(v string) *DescribeSandboxInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageNumber(v string) *DescribeSandboxInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageSize(v string) *DescribeSandboxInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) Validate() error {
	return dara.Validate(s)
}
