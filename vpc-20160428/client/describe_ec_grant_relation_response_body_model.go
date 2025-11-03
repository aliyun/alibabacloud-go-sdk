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
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The query results.
	EcGrantRelations []*DescribeEcGrantRelationResponseBodyEcGrantRelations `json:"EcGrantRelations,omitempty" xml:"EcGrantRelations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
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
	// The total number of entries returned.
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
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 1250123456123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when permissions on the VPC were granted to the VBR.
	//
	// example:
	//
	// 2022-09-02T11:46Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The VBRs that have permissions on the VPC. Valid values:
	//
	// 	- **All**: VBRs that reside in the specified region and belong to the specified Alibaba Cloud account all have permissions on the VPC.
	//
	// 	- **Specify**: Only the specified VBR has permissions on the VPC.
	//
	// example:
	//
	// All
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1brjuegjc88v3u9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// VPCname
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the vRouter.
	//
	// example:
	//
	// vrt-bp1i0yzgjd8ra05ec****
	InstanceRouterId *string `json:"InstanceRouterId,omitempty" xml:"InstanceRouterId,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The query result. Valid values:
	//
	// 	- **Created**: The VBR has permissions on the VPC.
	//
	// 	- **Deleted**: The VBR does not have permissions on the VPC.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-m5ex0xf63xk8s5bob****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VBR belongs.
	//
	// example:
	//
	// 1210123456123456
	VbrOwnerUid *int64 `json:"VbrOwnerUid,omitempty" xml:"VbrOwnerUid,omitempty"`
	// The ID of the region where the VBR is deployed.
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
