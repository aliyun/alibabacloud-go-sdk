// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiAppRiskEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAiAppRiskEventRequest
	GetAppId() *string
	SetEndTime(v string) *ListAiAppRiskEventRequest
	GetEndTime() *string
	SetRegionId(v string) *ListAiAppRiskEventRequest
	GetRegionId() *string
	SetStartTime(v string) *ListAiAppRiskEventRequest
	GetStartTime() *string
}

type ListAiAppRiskEventRequest struct {
	// The application ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2026-01-02 16:08:38
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-01-01 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAiAppRiskEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiAppRiskEventRequest) GoString() string {
	return s.String()
}

func (s *ListAiAppRiskEventRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAiAppRiskEventRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAiAppRiskEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAiAppRiskEventRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAiAppRiskEventRequest) SetAppId(v string) *ListAiAppRiskEventRequest {
	s.AppId = &v
	return s
}

func (s *ListAiAppRiskEventRequest) SetEndTime(v string) *ListAiAppRiskEventRequest {
	s.EndTime = &v
	return s
}

func (s *ListAiAppRiskEventRequest) SetRegionId(v string) *ListAiAppRiskEventRequest {
	s.RegionId = &v
	return s
}

func (s *ListAiAppRiskEventRequest) SetStartTime(v string) *ListAiAppRiskEventRequest {
	s.StartTime = &v
	return s
}

func (s *ListAiAppRiskEventRequest) Validate() error {
	return dara.Validate(s)
}
