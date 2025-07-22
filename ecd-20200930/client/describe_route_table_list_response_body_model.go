// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRouteTableListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRouteTableListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeRouteTableListResponseBody
	GetRequestId() *string
	SetRouteTableList(v []*DescribeRouteTableListResponseBodyRouteTableList) *DescribeRouteTableListResponseBody
	GetRouteTableList() []*DescribeRouteTableListResponseBodyRouteTableList
}

type DescribeRouteTableListResponseBody struct {
	MaxResults     *int32                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteTableList []*DescribeRouteTableListResponseBodyRouteTableList `json:"RouteTableList,omitempty" xml:"RouteTableList,omitempty" type:"Repeated"`
}

func (s DescribeRouteTableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRouteTableListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteTableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteTableListResponseBody) GetRouteTableList() []*DescribeRouteTableListResponseBodyRouteTableList {
	return s.RouteTableList
}

func (s *DescribeRouteTableListResponseBody) SetMaxResults(v int32) *DescribeRouteTableListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetNextToken(v string) *DescribeRouteTableListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetRequestId(v string) *DescribeRouteTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteTableListResponseBody) SetRouteTableList(v []*DescribeRouteTableListResponseBodyRouteTableList) *DescribeRouteTableListResponseBody {
	s.RouteTableList = v
	return s
}

func (s *DescribeRouteTableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteTableListResponseBodyRouteTableList struct {
	AssociateType  *string   `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	RouteTableId   *string   `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	RouteTableType *string   `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	RouterType     *string   `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchIds     []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId          *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRouteTableListResponseBodyRouteTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponseBodyRouteTableList) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetAssociateType() *string {
	return s.AssociateType
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetAssociateType(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.AssociateType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetRouteTableId(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetRouteTableType(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.RouteTableType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetRouterType(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.RouterType = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetStatus(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.Status = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetVSwitchIds(v []*string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.VSwitchIds = v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) SetVpcId(v string) *DescribeRouteTableListResponseBodyRouteTableList {
	s.VpcId = &v
	return s
}

func (s *DescribeRouteTableListResponseBodyRouteTableList) Validate() error {
	return dara.Validate(s)
}
