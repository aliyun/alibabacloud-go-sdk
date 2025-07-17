// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromPrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalViewClusterId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest
	GetGlobalViewClusterId() *string
	SetGroupName(v string) *RemoveSourcesFromPrometheusGlobalViewRequest
	GetGroupName() *string
	SetRegionId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest
	GetRegionId() *string
	SetSourceNames(v string) *RemoveSourcesFromPrometheusGlobalViewRequest
	GetSourceNames() *string
}

type RemoveSourcesFromPrometheusGlobalViewRequest struct {
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
	// The list of custom data sources. You can specify multiple data sources and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// localPrometheusClusterName,testCumterPrometheusName
	SourceNames *string `json:"SourceNames,omitempty" xml:"SourceNames,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) GetSourceNames() *string {
	return s.SourceNames
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetGroupName(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetRegionId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetSourceNames(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.SourceNames = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
