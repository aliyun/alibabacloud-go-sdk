// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkReachableAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkReachableAnalysisIdsShrink(v string) *DeleteNetworkReachableAnalysisShrinkRequest
	GetNetworkReachableAnalysisIdsShrink() *string
	SetRegionId(v string) *DeleteNetworkReachableAnalysisShrinkRequest
	GetRegionId() *string
}

type DeleteNetworkReachableAnalysisShrinkRequest struct {
	// The IDs of the tasks for analyzing network reachability.
	//
	// This parameter is required.
	NetworkReachableAnalysisIdsShrink *string `json:"NetworkReachableAnalysisIds,omitempty" xml:"NetworkReachableAnalysisIds,omitempty"`
	// The ID of the region for which you want to delete a task for analyzing network reachability.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) GetNetworkReachableAnalysisIdsShrink() *string {
	return s.NetworkReachableAnalysisIdsShrink
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) SetNetworkReachableAnalysisIdsShrink(v string) *DeleteNetworkReachableAnalysisShrinkRequest {
	s.NetworkReachableAnalysisIdsShrink = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) SetRegionId(v string) *DeleteNetworkReachableAnalysisShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
