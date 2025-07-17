// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewByAliClusterIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest
	GetClusterIds() *string
	SetGroupName(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest
	GetGroupName() *string
	SetProductCode(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest
	GetProductCode() *string
	SetRegionId(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest
	GetRegionId() *string
}

type AddPrometheusGlobalViewByAliClusterIdsRequest struct {
	// The IDs of clusters. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// cd1d55bef19904324a20ed0ebb86caa5c,c5b48729918ab4745a24482ac29d0973a, c00a94896641449098bf24931e4166003, cd174485c09384060ba542bc1be1185a4
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The name of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// zyGlobalView
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The identifier to identify the service if custom dashboards are created for the specified clusters.
	//
	// example:
	//
	// adcp
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetClusterIds(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.ClusterIds = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetGroupName(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.GroupName = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetProductCode(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.ProductCode = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetRegionId(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) Validate() error {
	return dara.Validate(s)
}
