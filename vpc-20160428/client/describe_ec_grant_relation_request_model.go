// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcGrantRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeEcGrantRelationRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeEcGrantRelationRequest
	GetInstanceType() *string
	SetPageNumber(v int64) *DescribeEcGrantRelationRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEcGrantRelationRequest
	GetPageSize() *int64
	SetVbrRegionNo(v string) *DescribeEcGrantRelationRequest
	GetVbrRegionNo() *string
}

type DescribeEcGrantRelationRequest struct {
	// The instance ID of the instance for which you want to query authorization relationships.
	//
	// - If **InstanceType*	- is set to **VBR**, set this parameter to the VBR instance ID.
	//
	// - If **InstanceType*	- is set to **VPC**, set this parameter to the VPC-connected instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp12mw1f8k3jgygk9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of instance for which you want to query authorization relationships. Valid values:
	//
	// - **VBR**: Virtual Border Router (VBR) instance. Queries the VPC-connected instances that have granted authorization to the VBR instance.
	//
	// - **VPC**: virtual private cloud (VPC) instance. Queries the VBR instances to which the VPC-connected instance has granted authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The page number of the list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VBR instance for which you want to query authorization relationships.
	//
	// - If **InstanceType*	- is set to **VBR**, this parameter is required.
	//
	// - If **InstanceType*	- is set to **VPC**, this parameter is not required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionNo *string `json:"VbrRegionNo,omitempty" xml:"VbrRegionNo,omitempty"`
}

func (s DescribeEcGrantRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcGrantRelationRequest) GoString() string {
	return s.String()
}

func (s *DescribeEcGrantRelationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEcGrantRelationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEcGrantRelationRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEcGrantRelationRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEcGrantRelationRequest) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *DescribeEcGrantRelationRequest) SetInstanceId(v string) *DescribeEcGrantRelationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEcGrantRelationRequest) SetInstanceType(v string) *DescribeEcGrantRelationRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeEcGrantRelationRequest) SetPageNumber(v int64) *DescribeEcGrantRelationRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEcGrantRelationRequest) SetPageSize(v int64) *DescribeEcGrantRelationRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEcGrantRelationRequest) SetVbrRegionNo(v string) *DescribeEcGrantRelationRequest {
	s.VbrRegionNo = &v
	return s
}

func (s *DescribeEcGrantRelationRequest) Validate() error {
	return dara.Validate(s)
}
