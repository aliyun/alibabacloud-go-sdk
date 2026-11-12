// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesTrafficPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *ListInstancesTrafficPackagesRequest
	GetInstanceIds() *string
	SetRegionId(v string) *ListInstancesTrafficPackagesRequest
	GetRegionId() *string
}

type ListInstancesTrafficPackagesRequest struct {
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc******","2ad1ae67295445f598017499dc******"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesTrafficPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesTrafficPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListInstancesTrafficPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesTrafficPackagesRequest) SetInstanceIds(v string) *ListInstancesTrafficPackagesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesTrafficPackagesRequest) SetRegionId(v string) *ListInstancesTrafficPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesTrafficPackagesRequest) Validate() error {
	return dara.Validate(s)
}
