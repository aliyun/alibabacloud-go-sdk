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
	SetEffectiveTime(v string) *ModifyMasterSpecRequest
	GetEffectiveTime() *string
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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective period of the specification change. Valid values:
	//
	// - **Immediately*	- (default): The change takes effect immediately.
	//
	// - **MaintainTime**: The change takes effect during the maintenance window of the instance.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// If you want to change the master node to a MasterAI node, specify this parameter.
	//
	// > - This parameter and MasterCU cannot be specified at the same time.
	//
	// >- Only specific regions and zones support changing the master node to a MasterAI node.
	//
	// >- Only AnalyticDB for PostgreSQL V7.0 Basic Edition instances support MasterAI nodes.
	//
	// >- You can view all valid values of this parameter on the specification change page for the master node.
	//
	// example:
	//
	// ADB.AIMedium.2
	MasterAISpec *string `json:"MasterAISpec,omitempty" xml:"MasterAISpec,omitempty"`
	// The master resources. Valid values:
	//
	// - 2 CU
	//
	// - 4 CU
	//
	// - 8 CU
	//
	// - 16 CU
	//
	// - 32 CU
	//
	// > Master resources greater than 8 CU incur additional fees.
	//
	// example:
	//
	// 8 CU
	MasterCU *int32 `json:"MasterCU,omitempty" xml:"MasterCU,omitempty"`
	// The ID of the resource group to which the instance belongs. For information about how to obtain the resource group ID, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
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

func (s *ModifyMasterSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
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

func (s *ModifyMasterSpecRequest) SetEffectiveTime(v string) *ModifyMasterSpecRequest {
	s.EffectiveTime = &v
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
