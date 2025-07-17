// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAliClusterIdsFromPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest
	GetClusterIds() *string
	SetGlobalViewClusterId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetGroupName(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest
	GetGroupName() *string
	SetRegionId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest
	GetRegionId() *string
}

type RemoveAliClusterIdsFromPrometheusGlobalViewRequest struct {
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
	// global-v2-cn-1670100631025794-amaykca4
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

func (s RemoveAliClusterIdsFromPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetClusterIds(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.ClusterIds = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetGroupName(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetRegionId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
