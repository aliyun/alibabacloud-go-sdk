// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDatabasesToNetworkDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseIds(v []*string) *MoveDatabasesToNetworkDomainRequest
	GetDatabaseIds() []*string
	SetInstanceId(v string) *MoveDatabasesToNetworkDomainRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *MoveDatabasesToNetworkDomainRequest
	GetNetworkDomainId() *string
	SetProjectId(v int64) *MoveDatabasesToNetworkDomainRequest
	GetProjectId() *int64
	SetRegionId(v string) *MoveDatabasesToNetworkDomainRequest
	GetRegionId() *string
}

type MoveDatabasesToNetworkDomainRequest struct {
	// The instance IDs of the databases that you want to move to the network domain.
	DatabaseIds []*string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zpr3h2zo60l
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to which you want to move the databases.
	//
	// > You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to query the parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveDatabasesToNetworkDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveDatabasesToNetworkDomainRequest) GoString() string {
	return s.String()
}

func (s *MoveDatabasesToNetworkDomainRequest) GetDatabaseIds() []*string {
	return s.DatabaseIds
}

func (s *MoveDatabasesToNetworkDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MoveDatabasesToNetworkDomainRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *MoveDatabasesToNetworkDomainRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveDatabasesToNetworkDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveDatabasesToNetworkDomainRequest) SetDatabaseIds(v []*string) *MoveDatabasesToNetworkDomainRequest {
	s.DatabaseIds = v
	return s
}

func (s *MoveDatabasesToNetworkDomainRequest) SetInstanceId(v string) *MoveDatabasesToNetworkDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainRequest) SetNetworkDomainId(v string) *MoveDatabasesToNetworkDomainRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainRequest) SetProjectId(v int64) *MoveDatabasesToNetworkDomainRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainRequest) SetRegionId(v string) *MoveDatabasesToNetworkDomainRequest {
	s.RegionId = &v
	return s
}

func (s *MoveDatabasesToNetworkDomainRequest) Validate() error {
	return dara.Validate(s)
}
