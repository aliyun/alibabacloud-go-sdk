// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkReachableAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisRequest
	GetNetworkReachableAnalysisId() *string
	SetRegionId(v string) *GetNetworkReachableAnalysisRequest
	GetRegionId() *string
}

type GetNetworkReachableAnalysisRequest struct {
	// The ID of the task for analyzing network reachability. You can call the **CreateNetworkRearchableAnalysis*	- operation to obtain the ID of the task for analyzing network reachability.
	//
	// This parameter is required.
	//
	// example:
	//
	// nra-90eef36a9e6e4662****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The ID of the region for which you want to obtain the result of network reachability analysis.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetNetworkReachableAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkReachableAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisRequest) GetNetworkReachableAnalysisId() *string {
	return s.NetworkReachableAnalysisId
}

func (s *GetNetworkReachableAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNetworkReachableAnalysisRequest) SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisRequest {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *GetNetworkReachableAnalysisRequest) SetRegionId(v string) *GetNetworkReachableAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *GetNetworkReachableAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
