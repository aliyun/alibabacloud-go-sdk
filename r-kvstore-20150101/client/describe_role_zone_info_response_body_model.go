// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody
	GetNode() *DescribeRoleZoneInfoResponseBodyNode
	SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRoleZoneInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody
	GetTotalCount() *int32
}

type DescribeRoleZoneInfoResponseBody struct {
	Node *DescribeRoleZoneInfoResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 224B97FB-A275-4EAC-86E9-8922FEA2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) GetNode() *DescribeRoleZoneInfoResponseBodyNode {
	return s.Node
}

func (s *DescribeRoleZoneInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRoleZoneInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRoleZoneInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoleZoneInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRoleZoneInfoResponseBody) SetNode(v *DescribeRoleZoneInfoResponseBodyNode) *DescribeRoleZoneInfoResponseBody {
	s.Node = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageNumber(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetPageSize(v int32) *DescribeRoleZoneInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetTotalCount(v int32) *DescribeRoleZoneInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRoleZoneInfoResponseBodyNode struct {
	NodeInfo []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNode) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNode) GetNodeInfo() []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	return s.NodeInfo
}

func (s *DescribeRoleZoneInfoResponseBodyNode) SetNodeInfo(v []*DescribeRoleZoneInfoResponseBodyNodeNodeInfo) *DescribeRoleZoneInfoResponseBodyNode {
	s.NodeInfo = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNode) Validate() error {
	if s.NodeInfo != nil {
		for _, item := range s.NodeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRoleZoneInfoResponseBodyNodeNodeInfo struct {
	CurrentBandWidth       *int64  `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	CurrentMinorVersion    *string `json:"CurrentMinorVersion,omitempty" xml:"CurrentMinorVersion,omitempty"`
	CustinsId              *string `json:"CustinsId,omitempty" xml:"CustinsId,omitempty"`
	DefaultBandWidth       *int64  `json:"DefaultBandWidth,omitempty" xml:"DefaultBandWidth,omitempty"`
	InsName                *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	InsType                *int32  `json:"InsType,omitempty" xml:"InsType,omitempty"`
	IsLatestVersion        *int32  `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	IsOpenBandWidthService *bool   `json:"IsOpenBandWidthService,omitempty" xml:"IsOpenBandWidthService,omitempty"`
	NodeId                 *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType               *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Role                   *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCurrentBandWidth() *int64 {
	return s.CurrentBandWidth
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCurrentMinorVersion() *string {
	return s.CurrentMinorVersion
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetCustinsId() *string {
	return s.CustinsId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetDefaultBandWidth() *int64 {
	return s.DefaultBandWidth
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetInsName() *string {
	return s.InsName
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetInsType() *int32 {
	return s.InsType
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetIsLatestVersion() *int32 {
	return s.IsLatestVersion
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetIsOpenBandWidthService() *bool {
	return s.IsOpenBandWidthService
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetRole() *string {
	return s.Role
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCurrentMinorVersion(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CurrentMinorVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetCustinsId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.CustinsId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetDefaultBandWidth(v int64) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.DefaultBandWidth = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetInsType(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.InsType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsLatestVersion(v int32) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetIsOpenBandWidthService(v bool) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.IsOpenBandWidthService = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetRole(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.Role = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyNodeNodeInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyNodeNodeInfo) Validate() error {
	return dara.Validate(s)
}
