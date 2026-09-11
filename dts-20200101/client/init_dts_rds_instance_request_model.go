// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitDtsRdsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *InitDtsRdsInstanceRequest
	GetDtsInstanceId() *string
	SetEndpointCenId(v string) *InitDtsRdsInstanceRequest
	GetEndpointCenId() *string
	SetEndpointInstanceId(v string) *InitDtsRdsInstanceRequest
	GetEndpointInstanceId() *string
	SetEndpointInstanceType(v string) *InitDtsRdsInstanceRequest
	GetEndpointInstanceType() *string
	SetEndpointRegion(v string) *InitDtsRdsInstanceRequest
	GetEndpointRegion() *string
	SetRegionId(v string) *InitDtsRdsInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *InitDtsRdsInstanceRequest
	GetResourceGroupId() *string
}

type InitDtsRdsInstanceRequest struct {
	// The instance ID of the synchronization node.
	//
	// example:
	//
	// dtszvxa4qmot6p****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The instance ID of the CEN instance. This parameter is required if the unit node is a self-managed MySQL database connected through CEN.
	//
	// > You must specify either this parameter or the ApsaraDB RDS for MySQL-related parameters (**EndpointRegion*	- and **EndpointInstanceId**).
	//
	// example:
	//
	// cen-9kqshqum*******
	EndpointCenId *string `json:"EndpointCenId,omitempty" xml:"EndpointCenId,omitempty"`
	// The instance ID of the ApsaraDB RDS for MySQL instance. This parameter is required if the unit node is an ApsaraDB RDS for MySQL instance.
	//
	// > - You must also specify the **EndpointRegion*	- parameter.
	//
	// - You must specify either this parameter or **EndpointCenId**.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	EndpointInstanceId *string `json:"EndpointInstanceId,omitempty" xml:"EndpointInstanceId,omitempty"`
	// The instance type of the unit node. Valid values:
	//
	// - **RDS**: ApsaraDB RDS for MySQL instance.
	//
	// - **CEN**: self-managed MySQL database connected through CEN.
	//
	// example:
	//
	// RDS
	EndpointInstanceType *string `json:"EndpointInstanceType,omitempty" xml:"EndpointInstanceType,omitempty"`
	// The region in which the ApsaraDB RDS for MySQL instance resides. This parameter is required if the unit node is an ApsaraDB RDS for MySQL instance.
	//
	// > - You must also specify the **EndpointInstanceId*	- parameter.
	//
	// - You must specify either this parameter or **EndpointCenId**.
	//
	// example:
	//
	// cn-hangzhou
	EndpointRegion *string `json:"EndpointRegion,omitempty" xml:"EndpointRegion,omitempty"`
	// The region in which the active geo-redundancy database cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. This is a global parameter and does not need to be specified for this operation.
	//
	// example:
	//
	// 资源组ID，全局参数，当前API无需传入。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s InitDtsRdsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s InitDtsRdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *InitDtsRdsInstanceRequest) GetEndpointCenId() *string {
	return s.EndpointCenId
}

func (s *InitDtsRdsInstanceRequest) GetEndpointInstanceId() *string {
	return s.EndpointInstanceId
}

func (s *InitDtsRdsInstanceRequest) GetEndpointInstanceType() *string {
	return s.EndpointInstanceType
}

func (s *InitDtsRdsInstanceRequest) GetEndpointRegion() *string {
	return s.EndpointRegion
}

func (s *InitDtsRdsInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitDtsRdsInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InitDtsRdsInstanceRequest) SetDtsInstanceId(v string) *InitDtsRdsInstanceRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointCenId(v string) *InitDtsRdsInstanceRequest {
	s.EndpointCenId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointInstanceId(v string) *InitDtsRdsInstanceRequest {
	s.EndpointInstanceId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointInstanceType(v string) *InitDtsRdsInstanceRequest {
	s.EndpointInstanceType = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointRegion(v string) *InitDtsRdsInstanceRequest {
	s.EndpointRegion = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetRegionId(v string) *InitDtsRdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetResourceGroupId(v string) *InitDtsRdsInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
