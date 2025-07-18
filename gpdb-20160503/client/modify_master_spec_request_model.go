// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMasterSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyMasterSpecRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *ModifyMasterSpecRequest
	GetDBInstanceId() *string
	SetMasterAISpec(v string) *ModifyMasterSpecRequest
	GetMasterAISpec() *string
	SetMasterCU(v int32) *ModifyMasterSpecRequest
	GetMasterCU() *int32
	SetResourceGroupId(v string) *ModifyMasterSpecRequest
	GetResourceGroupId() *string
}

type ModifyMasterSpecRequest struct {
	// The description of the instance.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter must be specified if you want to change coordinator nodes to AI coordinator nodes.
	//
	// >-  You cannot specify the MasterAISpec and MasterCU parameters at the same time.
	//
	// >- You can change coordinator nodes to AI coordinator nodes only in specific regions and zones.
	//
	// >- Only AnalyticDB for PostgreSQL V7.0 instances of Basic Edition support AI coordinator nodes.
	//
	// >- You can view the valid values of this parameter on the configuration change page of coordinator nodes.
	//
	// example:
	//
	// ADB.AIMedium.2
	MasterAISpec *string `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	// The specifications of coordinator node resources. Valid values:
	//
	// 	- 2 CU
	//
	// 	- 4 CU
	//
	// 	- 8 CU
	//
	// 	- 16 CU
	//
	// 	- 32 CU
	//
	// >  You are charged for coordinator node resources of more than 8 compute units (CUs).
	//
	// example:
	//
	// 8 CU
	MasterCU *int32 `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyMasterSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMasterSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyMasterSpecRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyMasterSpecRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyMasterSpecRequest) GetMasterAISpec() *string {
	return s.MasterAISpec
}

func (s *ModifyMasterSpecRequest) GetMasterCU() *int32 {
	return s.MasterCU
}

func (s *ModifyMasterSpecRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyMasterSpecRequest) SetDBInstanceDescription(v string) *ModifyMasterSpecRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyMasterSpecRequest) SetDBInstanceId(v string) *ModifyMasterSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyMasterSpecRequest) SetMasterAISpec(v string) *ModifyMasterSpecRequest {
	s.MasterAISpec = &v
	return s
}

func (s *ModifyMasterSpecRequest) SetMasterCU(v int32) *ModifyMasterSpecRequest {
	s.MasterCU = &v
	return s
}

func (s *ModifyMasterSpecRequest) SetResourceGroupId(v string) *ModifyMasterSpecRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyMasterSpecRequest) Validate() error {
	return dara.Validate(s)
}
