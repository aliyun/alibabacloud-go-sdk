// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceId() *string
	SetIdleTime(v int32) *ModifyDBInstanceConfigRequest
	GetIdleTime() *int32
	SetResourceGroupId(v string) *ModifyDBInstanceConfigRequest
	GetResourceGroupId() *string
	SetServerlessResource(v int32) *ModifyDBInstanceConfigRequest
	GetServerlessResource() *int32
}

type ModifyDBInstanceConfigRequest struct {
	// The description of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The wait period for the instance that has no traffic to become idle. Minimum value: 60. Default value: 600. Unit: seconds.
	//
	// example:
	//
	// 600
	IdleTime *int32 `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The threshold of computing resources. Valid values: 8 to 32. Unit: AnalyticDB Compute Units (ACUs).
	//
	// example:
	//
	// 32
	ServerlessResource *int32 `json:"ServerlessResource,omitempty" xml:"ServerlessResource,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConfigRequest) GetIdleTime() *int32 {
	return s.IdleTime
}

func (s *ModifyDBInstanceConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceConfigRequest) GetServerlessResource() *int32 {
	return s.ServerlessResource
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetIdleTime(v int32) *ModifyDBInstanceConfigRequest {
	s.IdleTime = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceGroupId(v string) *ModifyDBInstanceConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetServerlessResource(v int32) *ModifyDBInstanceConfigRequest {
	s.ServerlessResource = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
