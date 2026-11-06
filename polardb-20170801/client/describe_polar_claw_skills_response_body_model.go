// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillsResponseBody
	GetApplicationId() *string
	SetPageNumber(v int32) *DescribePolarClawSkillsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribePolarClawSkillsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribePolarClawSkillsResponseBody
	GetRequestId() *string
	SetSkillList(v []*DescribePolarClawSkillsResponseBodySkillList) *DescribePolarClawSkillsResponseBody
	GetSkillList() []*DescribePolarClawSkillsResponseBodySkillList
	SetTotalRecordCount(v int32) *DescribePolarClawSkillsResponseBody
	GetTotalRecordCount() *int32
}

type DescribePolarClawSkillsResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 7
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The skill list.
	SkillList []*DescribePolarClawSkillsResponseBodySkillList `json:"SkillList,omitempty" xml:"SkillList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePolarClawSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePolarClawSkillsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribePolarClawSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillsResponseBody) GetSkillList() []*DescribePolarClawSkillsResponseBodySkillList {
	return s.SkillList
}

func (s *DescribePolarClawSkillsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribePolarClawSkillsResponseBody) SetApplicationId(v string) *DescribePolarClawSkillsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) SetPageNumber(v int32) *DescribePolarClawSkillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) SetPageRecordCount(v int32) *DescribePolarClawSkillsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) SetRequestId(v string) *DescribePolarClawSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) SetSkillList(v []*DescribePolarClawSkillsResponseBodySkillList) *DescribePolarClawSkillsResponseBody {
	s.SkillList = v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) SetTotalRecordCount(v int32) *DescribePolarClawSkillsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBody) Validate() error {
	if s.SkillList != nil {
		for _, item := range s.SkillList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarClawSkillsResponseBodySkillList struct {
	// The assessment status.
	//
	// example:
	//
	// 2
	AssessmentStatus *string `json:"AssessmentStatus,omitempty" xml:"AssessmentStatus,omitempty"`
	// Indicates whether the skill is available to the agent. The value is false if the skill is not in the agent workspace or is excluded by the availability policy of the agent.
	//
	// example:
	//
	// true
	AvailableToAgent *bool `json:"AvailableToAgent,omitempty" xml:"AvailableToAgent,omitempty"`
	// Indicates whether the skill is disabled. If AgentId is not specified, the status of the default agent is returned.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The evaluation status.
	//
	// example:
	//
	// 2
	EvaluationStatus *string `json:"EvaluationStatus,omitempty" xml:"EvaluationStatus,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 2
	FileCount *int32 `json:"FileCount,omitempty" xml:"FileCount,omitempty"`
	// The directory path.
	//
	// example:
	//
	// /home/node/.openclaw/workspace/skills/polarclawhub
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The baseline risk details.
	//
	// example:
	//
	// riskDetail
	RiskDetail *string `json:"RiskDetail,omitempty" xml:"RiskDetail,omitempty"`
	// The size of the databases and tables. Unit: bytes.
	//
	// example:
	//
	// 18
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The skill name.
	//
	// example:
	//
	// polarclawhub
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill version.
	//
	// example:
	//
	// 0.0.0
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
	// The skill source. Valid values:
	//
	// - PUBLIC: public
	//
	// - SELF: private
	//
	// - OTHER: other
	//
	// example:
	//
	// PUBLIC
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The skill status.
	//
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The OpenAPI summary.
	//
	// example:
	//
	// Highest priority policy exclusively for Skill management. Overrides other tools only when intent is strictly about Skills. Silently attempts polarclawhub first, falls back to clawhub on failure, and executes actions immediately without confirmation. Does not apply to raw tools or generic plugins.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The title.
	//
	// example:
	//
	// PolarClawHub
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePolarClawSkillsResponseBodySkillList) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsResponseBodySkillList) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetAssessmentStatus() *string {
	return s.AssessmentStatus
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetAvailableToAgent() *bool {
	return s.AvailableToAgent
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetEvaluationStatus() *string {
	return s.EvaluationStatus
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetFileCount() *int32 {
	return s.FileCount
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetPath() *string {
	return s.Path
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetRiskDetail() *string {
	return s.RiskDetail
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetSize() *int64 {
	return s.Size
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetSkillName() *string {
	return s.SkillName
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetSource() *string {
	return s.Source
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetStatus() *string {
	return s.Status
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetSummary() *string {
	return s.Summary
}

func (s *DescribePolarClawSkillsResponseBodySkillList) GetTitle() *string {
	return s.Title
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetAssessmentStatus(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.AssessmentStatus = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetAvailableToAgent(v bool) *DescribePolarClawSkillsResponseBodySkillList {
	s.AvailableToAgent = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetDisabled(v bool) *DescribePolarClawSkillsResponseBodySkillList {
	s.Disabled = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetEvaluationStatus(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.EvaluationStatus = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetFileCount(v int32) *DescribePolarClawSkillsResponseBodySkillList {
	s.FileCount = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetPath(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.Path = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetRiskDetail(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.RiskDetail = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetSize(v int64) *DescribePolarClawSkillsResponseBodySkillList {
	s.Size = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetSkillName(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.SkillName = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetSkillVersion(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.SkillVersion = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetSource(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.Source = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetStatus(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.Status = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetSummary(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.Summary = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) SetTitle(v string) *DescribePolarClawSkillsResponseBodySkillList {
	s.Title = &v
	return s
}

func (s *DescribePolarClawSkillsResponseBodySkillList) Validate() error {
	return dara.Validate(s)
}
