// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOptimizeTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOptimizeTasksResponseBody
	GetNextToken() *string
	SetOptimizeTasks(v []*ListOptimizeTasksResponseBodyOptimizeTasks) *ListOptimizeTasksResponseBody
	GetOptimizeTasks() []*ListOptimizeTasksResponseBodyOptimizeTasks
	SetRequestId(v string) *ListOptimizeTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOptimizeTasksResponseBody
	GetTotalCount() *int32
}

type ListOptimizeTasksResponseBody struct {
	// The maximum number of results per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of optimization tasks.
	OptimizeTasks []*ListOptimizeTasksResponseBodyOptimizeTasks `json:"optimizeTasks,omitempty" xml:"optimizeTasks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOptimizeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListOptimizeTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOptimizeTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOptimizeTasksResponseBody) GetOptimizeTasks() []*ListOptimizeTasksResponseBodyOptimizeTasks {
	return s.OptimizeTasks
}

func (s *ListOptimizeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOptimizeTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOptimizeTasksResponseBody) SetMaxResults(v int32) *ListOptimizeTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOptimizeTasksResponseBody) SetNextToken(v string) *ListOptimizeTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOptimizeTasksResponseBody) SetOptimizeTasks(v []*ListOptimizeTasksResponseBodyOptimizeTasks) *ListOptimizeTasksResponseBody {
	s.OptimizeTasks = v
	return s
}

func (s *ListOptimizeTasksResponseBody) SetRequestId(v string) *ListOptimizeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOptimizeTasksResponseBody) SetTotalCount(v int32) *ListOptimizeTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOptimizeTasksResponseBody) Validate() error {
	if s.OptimizeTasks != nil {
		for _, item := range s.OptimizeTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOptimizeTasksResponseBodyOptimizeTasks struct {
	// The scheduling checkpoint that stores states such as the general scheduling watermark.
	Checkpoint map[string]interface{} `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// The type-specific configuration. For ExperienceMining and SkillOptimize, optional prompts support only inline Markdown, and the body is stored in inputs.prompt.content (non-empty UTF-8, up to 16 KiB).
	//
	// example:
	//
	// {"ioVersion":2,"inputs":{"experienceLibrary":{"type":"dataset","dataset":"experiences"},"trajectory":{"type":"sls","project":"my-project","logstore":"agent-trajectory","format":"atif"},"prompt":{"type":"inline","format":"markdown","content":"# Guidance\\\\nExtract evidence-backed tool contracts."}},"outputs":{"changePlan":{"type":"dataset","dataset":"experience-plans"}},"operators":["tool_contract"],"application":{"mode":"manual"}}
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-12T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The next scheduling trigger time.
	//
	// example:
	//
	// 2026-08-13T16:00:00Z
	NextTriggerTime *string `json:"nextTriggerTime,omitempty" xml:"nextTriggerTime,omitempty"`
	// The name of the optimization task.
	//
	// example:
	//
	// efficiency-analysis
	OptimizeTaskName *string `json:"optimizeTaskName,omitempty" xml:"optimizeTaskName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The task status.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the optimization task. Valid values: ExperienceMining, SkillOptimize, EfficiencyAnalysis, SOPGeneration, EvalMining, PolicyMining, and AssetMining.
	//
	// example:
	//
	// EfficiencyAnalysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-08-12T10:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListOptimizeTasksResponseBodyOptimizeTasks) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeTasksResponseBodyOptimizeTasks) GoString() string {
	return s.String()
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetCheckpoint() map[string]interface{} {
	return s.Checkpoint
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetNextTriggerTime() *string {
	return s.NextTriggerTime
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetStatus() *string {
	return s.Status
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetType() *string {
	return s.Type
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetCheckpoint(v map[string]interface{}) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.Checkpoint = v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetConfig(v map[string]interface{}) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.Config = v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetCreateTime(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.CreateTime = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetNextTriggerTime(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.NextTriggerTime = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetOptimizeTaskName(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.OptimizeTaskName = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetRegionId(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.RegionId = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetStatus(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.Status = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetType(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.Type = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) SetUpdateTime(v string) *ListOptimizeTasksResponseBodyOptimizeTasks {
	s.UpdateTime = &v
	return s
}

func (s *ListOptimizeTasksResponseBodyOptimizeTasks) Validate() error {
	return dara.Validate(s)
}
