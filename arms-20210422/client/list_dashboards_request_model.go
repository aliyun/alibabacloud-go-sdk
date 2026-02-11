// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDashboardsRequest
	GetClusterId() *string
	SetClusterType(v string) *ListDashboardsRequest
	GetClusterType() *string
	SetProduct(v string) *ListDashboardsRequest
	GetProduct() *string
	SetRecreateSwitch(v bool) *ListDashboardsRequest
	GetRecreateSwitch() *bool
	SetRegionId(v string) *ListDashboardsRequest
	GetRegionId() *string
	SetTitle(v string) *ListDashboardsRequest
	GetTitle() *string
}

type ListDashboardsRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDashboardsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListDashboardsRequest) GetProduct() *string {
	return s.Product
}

func (s *ListDashboardsRequest) GetRecreateSwitch() *bool {
	return s.RecreateSwitch
}

func (s *ListDashboardsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDashboardsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListDashboardsRequest) SetClusterId(v string) *ListDashboardsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterType(v string) *ListDashboardsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsRequest) SetProduct(v string) *ListDashboardsRequest {
	s.Product = &v
	return s
}

func (s *ListDashboardsRequest) SetRecreateSwitch(v bool) *ListDashboardsRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *ListDashboardsRequest) SetRegionId(v string) *ListDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsRequest) SetTitle(v string) *ListDashboardsRequest {
	s.Title = &v
	return s
}

func (s *ListDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
