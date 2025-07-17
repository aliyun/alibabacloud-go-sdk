// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAliClusterIdsToPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest
	GetClusterIds() *string
	SetGlobalViewClusterId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetGroupName(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest
	GetGroupName() *string
	SetRegionId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest
	GetRegionId() *string
}

type AddAliClusterIdsToPrometheusGlobalViewRequest struct {
	// The IDs of clusters. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// cd1d55bef19904324a20ed0ebb86caa5c,c5b48729918ab4745a24482ac29d0973a, c00a94896641449098bf24931e4166003, cd174485c09384060ba542bc1be1185a4
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The ID of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// global-v2-cn-1478326682034601-vss8pd0i
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	// The name of the global aggregation instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// zyGlobalView
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetClusterIds(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.ClusterIds = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetGroupName(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetRegionId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
