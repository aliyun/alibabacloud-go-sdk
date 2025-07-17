// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRecordingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *SyncRecordingRulesRequest
	GetClusterId() *string
	SetRegionId(v string) *SyncRecordingRulesRequest
	GetRegionId() *string
	SetTargetClusters(v string) *SyncRecordingRulesRequest
	GetTargetClusters() *string
}

type SyncRecordingRulesRequest struct {
	// The ID of the cluster whose aggregation rule you want to synchronize.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region. The destination region can be the same as the source region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of clusters to which you want to synchronize the aggregation rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "cn":[         "c06ca68cd16f14f52bb07772eda***",         "c33dd70a0ac184c1b879d807ab2***",         "c384cf7e4dcb543e6ac8c7d4dd3***"     ],     "us":[         "ce30f833bc4a04a56a06b070319***"     ],     "jp":[      ],     "ap":[      ],     "gov":[      ],     "finance":[      ] }
	TargetClusters *string `json:"TargetClusters,omitempty" xml:"TargetClusters,omitempty"`
}

func (s SyncRecordingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncRecordingRulesRequest) GoString() string {
	return s.String()
}

func (s *SyncRecordingRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *SyncRecordingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SyncRecordingRulesRequest) GetTargetClusters() *string {
	return s.TargetClusters
}

func (s *SyncRecordingRulesRequest) SetClusterId(v string) *SyncRecordingRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *SyncRecordingRulesRequest) SetRegionId(v string) *SyncRecordingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SyncRecordingRulesRequest) SetTargetClusters(v string) *SyncRecordingRulesRequest {
	s.TargetClusters = &v
	return s
}

func (s *SyncRecordingRulesRequest) Validate() error {
	return dara.Validate(s)
}
