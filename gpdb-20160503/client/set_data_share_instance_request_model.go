// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataShareInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*string) *SetDataShareInstanceRequest
	GetInstanceList() []*string
	SetOperationType(v string) *SetDataShareInstanceRequest
	GetOperationType() *string
	SetOwnerId(v int64) *SetDataShareInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetDataShareInstanceRequest
	GetRegionId() *string
}

type SetDataShareInstanceRequest struct {
	// The ID of the AnalyticDB for PostgreSQL instance in Serverless mode.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	InstanceList []*string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// Specifies whether to enable or disable data sharing. Valid values:
	//
	// 	- **add**: enables data sharing.
	//
	// 	- **remove**: disables data sharing.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDataShareInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataShareInstanceRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceRequest) GetInstanceList() []*string {
	return s.InstanceList
}

func (s *SetDataShareInstanceRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetDataShareInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDataShareInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDataShareInstanceRequest) SetInstanceList(v []*string) *SetDataShareInstanceRequest {
	s.InstanceList = v
	return s
}

func (s *SetDataShareInstanceRequest) SetOperationType(v string) *SetDataShareInstanceRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetOwnerId(v int64) *SetDataShareInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceRequest) SetRegionId(v string) *SetDataShareInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *SetDataShareInstanceRequest) Validate() error {
	return dara.Validate(s)
}
