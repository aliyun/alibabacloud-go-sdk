// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFreeResourceClusterControlItem interface {
	dara.Model
	String() string
	GoString() string
	SetClusterID(v string) *FreeResourceClusterControlItem
	GetClusterID() *string
	SetClusterName(v string) *FreeResourceClusterControlItem
	GetClusterName() *string
	SetCrossClusters(v bool) *FreeResourceClusterControlItem
	GetCrossClusters() *bool
	SetEnableFreeResource(v bool) *FreeResourceClusterControlItem
	GetEnableFreeResource() *bool
	SetFreeResourceClusterControlId(v string) *FreeResourceClusterControlItem
	GetFreeResourceClusterControlId() *string
	SetGmtCreateTime(v string) *FreeResourceClusterControlItem
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *FreeResourceClusterControlItem
	GetGmtModifyTime() *string
	SetRegionID(v string) *FreeResourceClusterControlItem
	GetRegionID() *string
}

type FreeResourceClusterControlItem struct {
	ClusterID          *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName        *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CrossClusters      *bool   `json:"CrossClusters,omitempty" xml:"CrossClusters,omitempty"`
	EnableFreeResource *bool   `json:"EnableFreeResource,omitempty" xml:"EnableFreeResource,omitempty"`
	// example:
	//
	// frcc-whateversth
	FreeResourceClusterControlId *string `json:"FreeResourceClusterControlId,omitempty" xml:"FreeResourceClusterControlId,omitempty"`
	GmtCreateTime                *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime                *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	RegionID                     *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
}

func (s FreeResourceClusterControlItem) String() string {
	return dara.Prettify(s)
}

func (s FreeResourceClusterControlItem) GoString() string {
	return s.String()
}

func (s *FreeResourceClusterControlItem) GetClusterID() *string {
	return s.ClusterID
}

func (s *FreeResourceClusterControlItem) GetClusterName() *string {
	return s.ClusterName
}

func (s *FreeResourceClusterControlItem) GetCrossClusters() *bool {
	return s.CrossClusters
}

func (s *FreeResourceClusterControlItem) GetEnableFreeResource() *bool {
	return s.EnableFreeResource
}

func (s *FreeResourceClusterControlItem) GetFreeResourceClusterControlId() *string {
	return s.FreeResourceClusterControlId
}

func (s *FreeResourceClusterControlItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *FreeResourceClusterControlItem) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *FreeResourceClusterControlItem) GetRegionID() *string {
	return s.RegionID
}

func (s *FreeResourceClusterControlItem) SetClusterID(v string) *FreeResourceClusterControlItem {
	s.ClusterID = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetClusterName(v string) *FreeResourceClusterControlItem {
	s.ClusterName = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetCrossClusters(v bool) *FreeResourceClusterControlItem {
	s.CrossClusters = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetEnableFreeResource(v bool) *FreeResourceClusterControlItem {
	s.EnableFreeResource = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetFreeResourceClusterControlId(v string) *FreeResourceClusterControlItem {
	s.FreeResourceClusterControlId = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetGmtCreateTime(v string) *FreeResourceClusterControlItem {
	s.GmtCreateTime = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetGmtModifyTime(v string) *FreeResourceClusterControlItem {
	s.GmtModifyTime = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetRegionID(v string) *FreeResourceClusterControlItem {
	s.RegionID = &v
	return s
}

func (s *FreeResourceClusterControlItem) Validate() error {
	return dara.Validate(s)
}
