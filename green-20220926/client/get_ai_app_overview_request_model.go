// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetAiAppOverviewRequest
	GetEndTime() *string
	SetRegionId(v string) *GetAiAppOverviewRequest
	GetRegionId() *string
	SetStartTime(v string) *GetAiAppOverviewRequest
	GetStartTime() *string
}

type GetAiAppOverviewRequest struct {
	// The end time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2025-07-09 10:30:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Format: YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2023-08-21 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAiAppOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppOverviewRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppOverviewRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppOverviewRequest) SetEndTime(v string) *GetAiAppOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *GetAiAppOverviewRequest) SetRegionId(v string) *GetAiAppOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppOverviewRequest) SetStartTime(v string) *GetAiAppOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *GetAiAppOverviewRequest) Validate() error {
	return dara.Validate(s)
}
