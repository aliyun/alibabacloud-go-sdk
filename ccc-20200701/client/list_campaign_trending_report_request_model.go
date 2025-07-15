// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCampaignTrendingReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *ListCampaignTrendingReportRequest
	GetCampaignId() *string
	SetEndTime(v int64) *ListCampaignTrendingReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListCampaignTrendingReportRequest
	GetInstanceId() *string
	SetStartTime(v int64) *ListCampaignTrendingReportRequest
	GetStartTime() *int64
}

type ListCampaignTrendingReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6badb397-a8b5-40b6-21019d382a09
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// 2021-10-14 20:59:59
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2021-10-14 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCampaignTrendingReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCampaignTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListCampaignTrendingReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCampaignTrendingReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCampaignTrendingReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCampaignTrendingReportRequest) SetCampaignId(v string) *ListCampaignTrendingReportRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetEndTime(v int64) *ListCampaignTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetInstanceId(v string) *ListCampaignTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetStartTime(v int64) *ListCampaignTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) Validate() error {
	return dara.Validate(s)
}
