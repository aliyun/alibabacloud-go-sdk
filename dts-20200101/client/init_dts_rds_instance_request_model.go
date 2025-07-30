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
	// The ID of the data synchronization task.
	//
	// example:
	//
	// dtszvxa4qmot6p****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// If the node is a self-managed MySQL database that is connected over CEN, you must specify the ID of the CEN instance.
	//
	// > You must specify the **EndpointRegion*	- and **EndpointInstanceId*	- parameters or the EndpointCenId parameter based on the type of the node.
	//
	// example:
	//
	// cen-9kqshqum*******
	EndpointCenId *string `json:"EndpointCenId,omitempty" xml:"EndpointCenId,omitempty"`
	// If the node is an ApsaraDB RDS for MySQL instance, you must specify the ID of the ApsaraDB RDS for MySQL instance.
	//
	// > 	- You must also specify the **EndpointRegion*	- parameter.
	//
	// >	- You must specify the EndpointInstanceId parameter or the **EndpointCenId*	- parameter based on the type of the node.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	EndpointInstanceId *string `json:"EndpointInstanceId,omitempty" xml:"EndpointInstanceId,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **RDS**: an ApsaraDB RDS for MySQL instance
	//
	// 	- **CEN**: a self-managed MySQL database that is connected over CEN
	//
	// example:
	//
	// RDS
	EndpointInstanceType *string `json:"EndpointInstanceType,omitempty" xml:"EndpointInstanceType,omitempty"`
	// If the node is an ApsaraDB RDS for MySQL instance, you must specify the region in which the ApsaraDB RDS for MySQL instance resides.
	//
	// > 	- You must also specify the **EndpointInstanceId*	- parameter.
	//
	// >	- You must specify the EndpointRegion parameter or the **EndpointCenId*	- parameter based on the type of the node.
	//
	// example:
	//
	// cn-hangzhou
	EndpointRegion *string `json:"EndpointRegion,omitempty" xml:"EndpointRegion,omitempty"`
	// The ID of the region in which the active geo-redundancy database cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
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
