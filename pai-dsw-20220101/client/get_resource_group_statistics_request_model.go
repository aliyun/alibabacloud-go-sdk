// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetResourceGroupStatisticsRequest
	GetEndTime() *string
	SetResourceId(v string) *GetResourceGroupStatisticsRequest
	GetResourceId() *string
	SetStartTime(v string) *GetResourceGroupStatisticsRequest
	GetStartTime() *string
	SetWorkspaceIds(v string) *GetResourceGroupStatisticsRequest
	GetWorkspaceIds() *string
}

type GetResourceGroupStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-123456789
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 27218,34956
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s GetResourceGroupStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetResourceGroupStatisticsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceGroupStatisticsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetResourceGroupStatisticsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *GetResourceGroupStatisticsRequest) SetEndTime(v string) *GetResourceGroupStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetResourceId(v string) *GetResourceGroupStatisticsRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetStartTime(v string) *GetResourceGroupStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) SetWorkspaceIds(v string) *GetResourceGroupStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *GetResourceGroupStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
