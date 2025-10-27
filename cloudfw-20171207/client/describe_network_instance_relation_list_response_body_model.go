// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceRelationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInstanceList(v []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) *DescribeNetworkInstanceRelationListResponseBody
	GetNetworkInstanceList() []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList
	SetRequestId(v string) *DescribeNetworkInstanceRelationListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworkInstanceRelationListResponseBody
	GetTotalCount() *int32
}

type DescribeNetworkInstanceRelationListResponseBody struct {
	NetworkInstanceList []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList `json:"NetworkInstanceList,omitempty" xml:"NetworkInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 284FF89D-4F70-546F-8EF6-77E0A530****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponseBody) GetNetworkInstanceList() []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	return s.NetworkInstanceList
}

func (s *DescribeNetworkInstanceRelationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInstanceRelationListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworkInstanceRelationListResponseBody) SetNetworkInstanceList(v []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) *DescribeNetworkInstanceRelationListResponseBody {
	s.NetworkInstanceList = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBody) SetRequestId(v string) *DescribeNetworkInstanceRelationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBody) SetTotalCount(v int32) *DescribeNetworkInstanceRelationListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBody) Validate() error {
	if s.NetworkInstanceList != nil {
		for _, item := range s.NetworkInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList struct {
	AssociatedCen []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen `json:"AssociatedCen,omitempty" xml:"AssociatedCen,omitempty" type:"Repeated"`
	// example:
	//
	// cen
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// example:
	//
	// vpc-2vcwfqbrh4kr****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// example:
	//
	// vpc-test
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// example:
	//
	// VPC
	NetworkInstanceType     *string                                                                                      `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	PeerNetworkInstanceList []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList `json:"PeerNetworkInstanceList,omitempty" xml:"PeerNetworkInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetAssociatedCen() []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	return s.AssociatedCen
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetPeerNetworkInstanceList() []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	return s.PeerNetworkInstanceList
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetAssociatedCen(v []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.AssociatedCen = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetConnectType(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.ConnectType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetNetworkInstanceId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetNetworkInstanceName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetNetworkInstanceType(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetPeerNetworkInstanceList(v []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.PeerNetworkInstanceList = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) SetRegionNo(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList {
	s.RegionNo = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceList) Validate() error {
	if s.AssociatedCen != nil {
		for _, item := range s.AssociatedCen {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PeerNetworkInstanceList != nil {
		for _, item := range s.PeerNetworkInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen struct {
	AttachmentId      *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	AttachmentName    *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	CenId             *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenName           *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GetCenId() *string {
	return s.CenId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GetCenName() *string {
	return s.CenName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) SetAttachmentId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	s.AttachmentId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) SetAttachmentName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	s.AttachmentName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) SetCenId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	s.CenId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) SetCenName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	s.CenName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) SetTransitRouterType(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListAssociatedCen) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList struct {
	AssociatedCen []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen `json:"AssociatedCen,omitempty" xml:"AssociatedCen,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-dsf232d****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// example:
	//
	// vpc-test
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// example:
	//
	// vpc
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GetAssociatedCen() []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	return s.AssociatedCen
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) SetAssociatedCen(v []*DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	s.AssociatedCen = v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) SetNetworkInstanceId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) SetNetworkInstanceName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) SetNetworkInstanceType(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) SetRegionNo(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList {
	s.RegionNo = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceList) Validate() error {
	if s.AssociatedCen != nil {
		for _, item := range s.AssociatedCen {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen struct {
	AttachmentId      *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	AttachmentName    *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	CenId             *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenName           *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GetCenId() *string {
	return s.CenId
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GetCenName() *string {
	return s.CenName
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) SetAttachmentId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	s.AttachmentId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) SetAttachmentName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	s.AttachmentName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) SetCenId(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	s.CenId = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) SetCenName(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	s.CenName = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) SetTransitRouterType(v string) *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeNetworkInstanceRelationListResponseBodyNetworkInstanceListPeerNetworkInstanceListAssociatedCen) Validate() error {
	return dara.Validate(s)
}
