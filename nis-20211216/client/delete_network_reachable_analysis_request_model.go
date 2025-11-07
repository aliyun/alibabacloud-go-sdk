// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkReachableAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkReachableAnalysisIds(v []*string) *DeleteNetworkReachableAnalysisRequest
	GetNetworkReachableAnalysisIds() []*string
	SetRegionId(v string) *DeleteNetworkReachableAnalysisRequest
	GetRegionId() *string
}

type DeleteNetworkReachableAnalysisRequest struct {
	// The IDs of the tasks for analyzing network reachability.
	//
	// This parameter is required.
	NetworkReachableAnalysisIds []*string `json:"NetworkReachableAnalysisIds,omitempty" xml:"NetworkReachableAnalysisIds,omitempty" type:"Repeated"`
	// The ID of the region for which you want to delete a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisRequest) GetNetworkReachableAnalysisIds() []*string {
	return s.NetworkReachableAnalysisIds
}

func (s *DeleteNetworkReachableAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkReachableAnalysisRequest) SetNetworkReachableAnalysisIds(v []*string) *DeleteNetworkReachableAnalysisRequest {
	s.NetworkReachableAnalysisIds = v
	return s
}

func (s *DeleteNetworkReachableAnalysisRequest) SetRegionId(v string) *DeleteNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
