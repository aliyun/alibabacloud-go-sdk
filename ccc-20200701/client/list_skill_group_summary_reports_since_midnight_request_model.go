// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupSummaryReportsSinceMidnightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSkillGroupSummaryReportsSinceMidnightRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListSkillGroupSummaryReportsSinceMidnightRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillGroupSummaryReportsSinceMidnightRequest
	GetPageSize() *int32
	SetSkillGroups(v string) *ListSkillGroupSummaryReportsSinceMidnightRequest
	GetSkillGroups() *string
}

type ListSkillGroupSummaryReportsSinceMidnightRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return. The value must be in the range of 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The value must be in the range of 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A JSON-formatted string that contains the IDs of the skill groups to query. If this parameter is omitted, the query includes all skill groups.
	//
	// example:
	//
	// ["skillgroup1@ccc-test", "skillgroup2@ccc-test2"]
	SkillGroups *string `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) GetSkillGroups() *string {
	return s.SkillGroups
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) SetInstanceId(v string) *ListSkillGroupSummaryReportsSinceMidnightRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) SetPageNumber(v int32) *ListSkillGroupSummaryReportsSinceMidnightRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) SetPageSize(v int32) *ListSkillGroupSummaryReportsSinceMidnightRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) SetSkillGroups(v string) *ListSkillGroupSummaryReportsSinceMidnightRequest {
	s.SkillGroups = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightRequest) Validate() error {
	return dara.Validate(s)
}
