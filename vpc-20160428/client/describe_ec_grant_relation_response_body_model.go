// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcGrantRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeEcGrantRelationResponseBody
	GetCount() *int32
	SetEcGrantRelations(v []*DescribeEcGrantRelationResponseBodyEcGrantRelations) *DescribeEcGrantRelationResponseBody
	GetEcGrantRelations() []*DescribeEcGrantRelationResponseBodyEcGrantRelations
	SetPage(v int32) *DescribeEcGrantRelationResponseBody
	GetPage() *int32
	SetPageSize(v int32) *DescribeEcGrantRelationResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEcGrantRelationResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEcGrantRelationResponseBody
	GetTotalCount() *int32
}

type DescribeEcGrantRelationResponseBody struct {
	// The number of query results.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of authorization relationship query results.
	EcGrantRelations []*DescribeEcGrantRelationResponseBodyEcGrantRelations `json:"EcGrantRelations,omitempty" xml:"EcGrantRelations,omitempty" type:"Repeated"`
	// The page number of the list.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6E90F6B-2B41-5AAF-ABEB-236ADBAAD91D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEcGrantRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcGrantRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEcGrantRelationResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeEcGrantRelationResponseBody) GetEcGrantRelations() []*DescribeEcGrantRelationResponseBodyEcGrantRelations {
	return s.EcGrantRelations
}

func (s *DescribeEcGrantRelationResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *DescribeEcGrantRelationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEcGrantRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEcGrantRelationResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEcGrantRelationResponseBody) SetCount(v int32) *DescribeEcGrantRelationResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) SetEcGrantRelations(v []*DescribeEcGrantRelationResponseBodyEcGrantRelations) *DescribeEcGrantRelationResponseBody {
	s.EcGrantRelations = v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) SetPage(v int32) *DescribeEcGrantRelationResponseBody {
	s.Page = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) SetPageSize(v int32) *DescribeEcGrantRelationResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) SetRequestId(v string) *DescribeEcGrantRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) SetTotalCount(v int32) *DescribeEcGrantRelationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBody) Validate() error {
	if s.EcGrantRelations != nil {
		for _, item := range s.EcGrantRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEcGrantRelationResponseBodyEcGrantRelations struct {
	// The Alibaba Cloud account ID of the VPC instance owner in the authorization relationship.
	//
	// example:
	//
	// 1250123456123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the VPC-connected instance granted authorization to the VBR instance.
	//
	// example:
	//
	// 2022-09-02T11:46Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The scope of VBR instances that are granted authorization. Valid values:
	//
	// - **All**: The VPC-connected instance is authorized to all VBR instances under the specified region and Alibaba Cloud account.
	//
	// - **Specify**: The VPC-connected instance is authorized to a specified VBR instance.
	//
	// example:
	//
	// All
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The instance ID of the VPC-connected instance in the authorization relationship.
	//
	// example:
	//
	// vpc-bp1brjuegjc88v3u9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC instance in the authorization relationship.
	//
	// example:
	//
	// VPCname
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The vRouter ID of the VPC instance in the authorization relationship.
	//
	// example:
	//
	// vrt-bp1i0yzgjd8ra05ec****
	InstanceRouterId *string `json:"InstanceRouterId,omitempty" xml:"InstanceRouterId,omitempty"`
	// The region ID of the VPC instance in the authorization relationship.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of the authorization relationship. Valid values:
	//
	// - **Created**: An authorization relationship exists between the VPC-connected instance and the VBR instance.
	//
	// - **Deleted**: No authorization relationship exists between the VPC-connected instance and the VBR instance.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance ID of the VBR instance in the authorization relationship.
	//
	// example:
	//
	// vbr-m5ex0xf63xk8s5bob****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The Alibaba Cloud account ID of the VBR instance owner.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The region ID of the VBR instance in the authorization relationship.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionNo *string `json:"VbrRegionNo,omitempty" xml:"VbrRegionNo,omitempty"`
}

func (s DescribeEcGrantRelationResponseBodyEcGrantRelations) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcGrantRelationResponseBodyEcGrantRelations) GoString() string {
	return s.String()
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetGrantType() *string {
	return s.GrantType
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetInstanceRouterId() *string {
	return s.InstanceRouterId
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetStatus() *string {
	return s.Status
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetVbrOwnerUid() *int64 {
	return s.VbrOwnerUid
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) GetVbrRegionNo() *string {
	return s.VbrRegionNo
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetAliUid(v int64) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.AliUid = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetGmtCreate(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.GmtCreate = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetGrantType(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.GrantType = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetInstanceId(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.InstanceId = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetInstanceName(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.InstanceName = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetInstanceRouterId(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.InstanceRouterId = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetRegionNo(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.RegionNo = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetStatus(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.Status = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetVbrInstanceId(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetVbrOwnerUid(v int64) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.VbrOwnerUid = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) SetVbrRegionNo(v string) *DescribeEcGrantRelationResponseBodyEcGrantRelations {
	s.VbrRegionNo = &v
	return s
}

func (s *DescribeEcGrantRelationResponseBodyEcGrantRelations) Validate() error {
	return dara.Validate(s)
}
