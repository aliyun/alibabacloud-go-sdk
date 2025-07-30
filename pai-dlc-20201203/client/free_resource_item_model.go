// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreeResourceItem interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableNumber(v int64) *FreeResourceItem
	GetAvailableNumber() *int64
	SetClusterID(v string) *FreeResourceItem
	GetClusterID() *string
	SetClusterName(v string) *FreeResourceItem
	GetClusterName() *string
	SetFreeResourceId(v string) *FreeResourceItem
	GetFreeResourceId() *string
	SetGmtCreateTime(v string) *FreeResourceItem
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *FreeResourceItem
	GetGmtModifyTime() *string
	SetRegionID(v string) *FreeResourceItem
	GetRegionID() *string
	SetResourceType(v string) *FreeResourceItem
	GetResourceType() *string
}

type FreeResourceItem struct {
	// example:
	//
	// 2
	AvailableNumber *int64  `json:"AvailableNumber,omitempty" xml:"AvailableNumber,omitempty"`
	ClusterID       *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// freeres-whateversth
	FreeResourceId *string `json:"FreeResourceId,omitempty" xml:"FreeResourceId,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// inner
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// example:
	//
	// cpu
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s FreeResourceItem) String() string {
	return dara.Prettify(s)
}

func (s FreeResourceItem) GoString() string {
	return s.String()
}

func (s *FreeResourceItem) GetAvailableNumber() *int64 {
	return s.AvailableNumber
}

func (s *FreeResourceItem) GetClusterID() *string {
	return s.ClusterID
}

func (s *FreeResourceItem) GetClusterName() *string {
	return s.ClusterName
}

func (s *FreeResourceItem) GetFreeResourceId() *string {
	return s.FreeResourceId
}

func (s *FreeResourceItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *FreeResourceItem) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *FreeResourceItem) GetRegionID() *string {
	return s.RegionID
}

func (s *FreeResourceItem) GetResourceType() *string {
	return s.ResourceType
}

func (s *FreeResourceItem) SetAvailableNumber(v int64) *FreeResourceItem {
	s.AvailableNumber = &v
	return s
}

func (s *FreeResourceItem) SetClusterID(v string) *FreeResourceItem {
	s.ClusterID = &v
	return s
}

func (s *FreeResourceItem) SetClusterName(v string) *FreeResourceItem {
	s.ClusterName = &v
	return s
}

func (s *FreeResourceItem) SetFreeResourceId(v string) *FreeResourceItem {
	s.FreeResourceId = &v
	return s
}

func (s *FreeResourceItem) SetGmtCreateTime(v string) *FreeResourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *FreeResourceItem) SetGmtModifyTime(v string) *FreeResourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *FreeResourceItem) SetRegionID(v string) *FreeResourceItem {
	s.RegionID = &v
	return s
}

func (s *FreeResourceItem) SetResourceType(v string) *FreeResourceItem {
	s.ResourceType = &v
	return s
}

func (s *FreeResourceItem) Validate() error {
	return dara.Validate(s)
}
