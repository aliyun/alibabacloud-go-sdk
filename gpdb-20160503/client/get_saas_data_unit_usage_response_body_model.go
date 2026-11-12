// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasDataUnitUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistribution(v []*GetSaasDataUnitUsageResponseBodyDistribution) *GetSaasDataUnitUsageResponseBody
	GetDistribution() []*GetSaasDataUnitUsageResponseBodyDistribution
	SetEndTime(v string) *GetSaasDataUnitUsageResponseBody
	GetEndTime() *string
	SetRequestId(v string) *GetSaasDataUnitUsageResponseBody
	GetRequestId() *string
	SetServiceId(v string) *GetSaasDataUnitUsageResponseBody
	GetServiceId() *string
	SetStartTime(v string) *GetSaasDataUnitUsageResponseBody
	GetStartTime() *string
	SetSummary(v *GetSaasDataUnitUsageResponseBodySummary) *GetSaasDataUnitUsageResponseBody
	GetSummary() *GetSaasDataUnitUsageResponseBodySummary
	SetTimeSeries(v []*GetSaasDataUnitUsageResponseBodyTimeSeries) *GetSaasDataUnitUsageResponseBody
	GetTimeSeries() []*GetSaasDataUnitUsageResponseBodyTimeSeries
	SetWarnings(v []*GetSaasDataUnitUsageResponseBodyWarnings) *GetSaasDataUnitUsageResponseBody
	GetWarnings() []*GetSaasDataUnitUsageResponseBodyWarnings
}

type GetSaasDataUnitUsageResponseBody struct {
	// The business distribution. The items are sorted in descending order of DataUnit. Items with unknown amounts are placed at the end. Items with the same value are sorted in ascending order of scenario code, project ID, and name. The items are grouped by scenario and project ID. Items without an ID but with a name are grouped by name. Records missing both are grouped as unassociated projects. The sum of all items in the complete result equals the summary. Missing or blank log scenarios are classified as other. Scenarios with values but no configured mapping retain the original values as codes and display names, and are normally included in the total. No empty rows are created for other if there are no records.
	Distribution []*GetSaasDataUnitUsageResponseBodyDistribution `json:"Distribution,omitempty" xml:"Distribution,omitempty" type:"Repeated"`
	// The actual end time of the query.
	//
	// example:
	//
	// 2026-10-20T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// mock_req_success
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SaaS service ID.
	//
	// example:
	//
	// agdb-example
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The actual start time of the query.
	//
	// example:
	//
	// 2026-10-20T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The usage summary.
	Summary *GetSaasDataUnitUsageResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// The time trend. The items are sorted in ascending order of the UTC bucket start time. Zero buckets are padded for complete queries. The value is null for unknown amounts and is not replaced with zero.
	TimeSeries []*GetSaasDataUnitUsageResponseBodyTimeSeries `json:"TimeSeries,omitempty" xml:"TimeSeries,omitempty" type:"Repeated"`
	// The product notifications.
	Warnings []*GetSaasDataUnitUsageResponseBodyWarnings `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
}

func (s GetSaasDataUnitUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBody) GetDistribution() []*GetSaasDataUnitUsageResponseBodyDistribution {
	return s.Distribution
}

func (s *GetSaasDataUnitUsageResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSaasDataUnitUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSaasDataUnitUsageResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasDataUnitUsageResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSaasDataUnitUsageResponseBody) GetSummary() *GetSaasDataUnitUsageResponseBodySummary {
	return s.Summary
}

func (s *GetSaasDataUnitUsageResponseBody) GetTimeSeries() []*GetSaasDataUnitUsageResponseBodyTimeSeries {
	return s.TimeSeries
}

func (s *GetSaasDataUnitUsageResponseBody) GetWarnings() []*GetSaasDataUnitUsageResponseBodyWarnings {
	return s.Warnings
}

func (s *GetSaasDataUnitUsageResponseBody) SetDistribution(v []*GetSaasDataUnitUsageResponseBodyDistribution) *GetSaasDataUnitUsageResponseBody {
	s.Distribution = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetEndTime(v string) *GetSaasDataUnitUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetRequestId(v string) *GetSaasDataUnitUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetServiceId(v string) *GetSaasDataUnitUsageResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetStartTime(v string) *GetSaasDataUnitUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetSummary(v *GetSaasDataUnitUsageResponseBodySummary) *GetSaasDataUnitUsageResponseBody {
	s.Summary = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetTimeSeries(v []*GetSaasDataUnitUsageResponseBodyTimeSeries) *GetSaasDataUnitUsageResponseBody {
	s.TimeSeries = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetWarnings(v []*GetSaasDataUnitUsageResponseBodyWarnings) *GetSaasDataUnitUsageResponseBody {
	s.Warnings = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) Validate() error {
	if s.Distribution != nil {
		for _, item := range s.Distribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeSeries != nil {
		for _, item := range s.TimeSeries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Warnings != nil {
		for _, item := range s.Warnings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSaasDataUnitUsageResponseBodyDistribution struct {
	// The exact non-negative decimal string. Floating-point or exponential formats are not used. The value is null when the amount is incomplete.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// The project attribution at the time of the call. The value is null when missing, which does not affect the total amount.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// project-1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name. For the same scenario and project ID, the latest non-empty name snapshot within the query range is used. If multiple snapshots exist in the same second, the maximum value in lexicographical order is used. The value is null when missing.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Scenario remake
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The business scenario code.
	//
	// example:
	//
	// drama_game_oneshot
	UsageScene *string `json:"UsageScene,omitempty" xml:"UsageScene,omitempty"`
	// The display name of the business scenario. This name is supplemented by a dictionary that maps the original Meta values to Chinese names, and is not an SLS field. UsageScene always retains the original log value. If no Chinese name is available, the original value is used as the name. Blank values are classified as other. Different codes with the same Chinese name are not merged.
	//
	// example:
	//
	// 游戏生成
	UsageSceneName *string `json:"UsageSceneName,omitempty" xml:"UsageSceneName,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyDistribution) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyDistribution) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetUsageScene() *string {
	return s.UsageScene
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetUsageSceneName() *string {
	return s.UsageSceneName
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetProjectId(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.ProjectId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetProjectName(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.ProjectName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetUsageScene(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.UsageScene = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetUsageSceneName(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.UsageSceneName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodySummary struct {
	// The exact non-negative decimal string. Floating-point or exponential formats are not used. The value is null when the amount is incomplete.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodySummary) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodySummary) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodySummary {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodySummary) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodyTimeSeries struct {
	// The exact non-negative decimal string. Floating-point or exponential formats are not used. The value is null when the amount is incomplete.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// The scenario usage within the bucket.
	Scenes []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	// The start time of the UTC bucket.
	//
	// example:
	//
	// 2026-10-20T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeries) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetScenes() []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	return s.Scenes
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetTime() *string {
	return s.Time
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetScenes(v []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.Scenes = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetTime(v string) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.Time = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) Validate() error {
	if s.Scenes != nil {
		for _, item := range s.Scenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSaasDataUnitUsageResponseBodyTimeSeriesScenes struct {
	// The exact non-negative decimal string. Floating-point or exponential formats are not used. The value is null when the amount is incomplete.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// The project attribution at the time of the call. The value is null when missing, which does not affect the total amount.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// project-1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name. For the same scenario and project ID, the latest non-empty name snapshot within the query range is used. If multiple snapshots exist in the same second, the maximum value in lexicographical order is used. The value is null when missing.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Scenario remake
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The business scenario code.
	//
	// example:
	//
	// game_translation
	UsageScene *string `json:"UsageScene,omitempty" xml:"UsageScene,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetUsageScene() *string {
	return s.UsageScene
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetProjectId(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.ProjectId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetProjectName(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.ProjectName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetUsageScene(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.UsageScene = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodyWarnings struct {
	// The notification code. Valid values include DATA_INCOMPLETE and DATA_UNIT_PENDING.
	//
	// example:
	//
	// DATA_INCOMPLETE
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The notification message.
	//
	// example:
	//
	// Partial usage is not ready yet
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyWarnings) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyWarnings) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) GetCode() *string {
	return s.Code
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) GetMessage() *string {
	return s.Message
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) SetCode(v string) *GetSaasDataUnitUsageResponseBodyWarnings {
	s.Code = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) SetMessage(v string) *GetSaasDataUnitUsageResponseBodyWarnings {
	s.Message = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) Validate() error {
	return dara.Validate(s)
}
