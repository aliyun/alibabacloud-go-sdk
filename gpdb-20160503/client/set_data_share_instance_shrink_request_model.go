// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataShareInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceListShrink(v string) *SetDataShareInstanceShrinkRequest
	GetInstanceListShrink() *string
	SetOperationType(v string) *SetDataShareInstanceShrinkRequest
	GetOperationType() *string
	SetOwnerId(v int64) *SetDataShareInstanceShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetDataShareInstanceShrinkRequest
	GetRegionId() *string
}

type SetDataShareInstanceShrinkRequest struct {
	// The ID of the AnalyticDB for PostgreSQL instance in Serverless mode.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	InstanceListShrink *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
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

func (s SetDataShareInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataShareInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetDataShareInstanceShrinkRequest) GetInstanceListShrink() *string {
	return s.InstanceListShrink
}

func (s *SetDataShareInstanceShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetDataShareInstanceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDataShareInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDataShareInstanceShrinkRequest) SetInstanceListShrink(v string) *SetDataShareInstanceShrinkRequest {
	s.InstanceListShrink = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOperationType(v string) *SetDataShareInstanceShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetOwnerId(v int64) *SetDataShareInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) SetRegionId(v string) *SetDataShareInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *SetDataShareInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
