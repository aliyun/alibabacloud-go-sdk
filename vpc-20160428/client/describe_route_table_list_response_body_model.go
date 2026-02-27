// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteTableListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteTableListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteTableListResponseBody
	GetRequestId() *string
	SetRouterTableList(v *DescribeRouteTableListResponseBodyRouterTableList) *DescribeRouteTableListResponseBody
	GetRouterTableList() *DescribeRouteTableListResponseBodyRouterTableList
	SetTotalCount(v int32) *DescribeRouteTableListResponseBody
	GetTotalCount() *int32
}

type DescribeRouteTableListResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	// DC668356-BCB4-42FD-9BC3-FA2B2E04B634
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouterTableList *DescribeRouteTableListResponseBodyRouterTableList `json:"RouterTableList,omitempty" xml:"RouterTableList,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteTableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteTableListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteTableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteTableListResponseBody) GetRouterTableList() *DescribeRouteTableListResponseBodyRouterTableList {
	return s.RouterTableList
}

func (s *DescribeRouteTableListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteTableListResponseBody) SetPageNumber(v int32) *DescribeRouteTableListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetPageSize(v int32) *DescribeRouteTableListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetRequestId(v string) *DescribeRouteTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetRouterTableList(v *DescribeRouteTableListResponseBodyRouterTableList) *DescribeRouteTableListResponseBody {
	s.RouterTableList = v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetTotalCount(v int32) *DescribeRouteTableListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) Validate() error {
	if s.RouterTableList != nil {
		if err := s.RouterTableList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTableListResponseBodyRouterTableList struct {
	RouterTableListType []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListType `json:"RouterTableListType,omitempty" xml:"RouterTableListType,omitempty" type:"Repeated"`
}

func (s DescribeRouteTableListResponseBodyRouterTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableList) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableList) GetRouterTableListType() []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	return s.RouterTableListType
}

func (s *DescribeRouteTableListResponseBodyRouterTableList) SetRouterTableListType(v []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) *DescribeRouteTableListResponseBodyRouterTableList {
	s.RouterTableListType = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableList) Validate() error {
	if s.RouterTableListType != nil {
		for _, item := range s.RouterTableListType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTableListResponseBodyRouterTableListRouterTableListType struct {
	AssociateType          *string                                                                         `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	CreationTime           *string                                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description            *string                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GatewayIds             *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds `json:"GatewayIds,omitempty" xml:"GatewayIds,omitempty" type:"Struct"`
	OwnerId                *int64                                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId        *string                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RoutePropagationEnable *bool                                                                           `json:"RoutePropagationEnable,omitempty" xml:"RoutePropagationEnable,omitempty"`
	RouteTableId           *string                                                                         `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	RouteTableName         *string                                                                         `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	RouteTableType         *string                                                                         `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	RouterId               *string                                                                         `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	RouterType             *string                                                                         `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	Status                 *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchIds             *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId                  *string                                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetAssociateType() *string {
	return s.AssociateType
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetGatewayIds() *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds {
	return s.GatewayIds
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRoutePropagationEnable() *bool {
	return s.RoutePropagationEnable
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetTags() *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags {
	return s.Tags
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetVSwitchIds() *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetAssociateType(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.AssociateType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetCreationTime(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetDescription(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.Description = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetGatewayIds(v *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.GatewayIds = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetOwnerId(v int64) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetResourceGroupId(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRoutePropagationEnable(v bool) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RoutePropagationEnable = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRouteTableId(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRouteTableName(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RouteTableName = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRouteTableType(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RouteTableType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRouterId(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RouterId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetRouterType(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.RouterType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetStatus(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.Status = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetTags(v *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.Tags = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetVSwitchIds(v *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.VSwitchIds = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) SetVpcId(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType {
	s.VpcId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListType) Validate() error {
	if s.GatewayIds != nil {
		if err := s.GatewayIds.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds struct {
	GatewayIds []*string `json:"GatewayIds,omitempty" xml:"GatewayIds,omitempty" type:"Repeated"`
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) GetGatewayIds() []*string {
	return s.GatewayIds
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) SetGatewayIds(v []*string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds {
	s.GatewayIds = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeGatewayIds) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags struct {
	Tag []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) GetTag() []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag {
	return s.Tag
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) SetTag(v []*DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags {
	s.Tag = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) SetKey(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) SetValue(v string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) SetVSwitchId(v []*string) *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouterTableListRouterTableListTypeVSwitchIds) Validate() error {
	return dara.Validate(s)
}
