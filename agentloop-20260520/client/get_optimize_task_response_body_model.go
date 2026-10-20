// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpoint(v map[string]interface{}) *GetOptimizeTaskResponseBody
	GetCheckpoint() map[string]interface{}
	SetConfig(v map[string]interface{}) *GetOptimizeTaskResponseBody
	GetConfig() map[string]interface{}
	SetCreateTime(v string) *GetOptimizeTaskResponseBody
	GetCreateTime() *string
	SetNextTriggerTime(v string) *GetOptimizeTaskResponseBody
	GetNextTriggerTime() *string
	SetOptimizeTaskName(v string) *GetOptimizeTaskResponseBody
	GetOptimizeTaskName() *string
	SetRegionId(v string) *GetOptimizeTaskResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetOptimizeTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetOptimizeTaskResponseBody
	GetStatus() *string
	SetType(v string) *GetOptimizeTaskResponseBody
	GetType() *string
	SetUpdateTime(v string) *GetOptimizeTaskResponseBody
	GetUpdateTime() *string
}

type GetOptimizeTaskResponseBody struct {
	// The scheduling checkpoint, which stores states such as the general scheduling watermark.
	Checkpoint map[string]interface{} `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// The type-specific configuration. Optional prompts for ExperienceMining or SkillOptimize support only inline Markdown. The body is saved to inputs.prompt.content, which must be a non-empty UTF-8 string of up to 16 KiB.
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
	// The optimization task name.
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
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task status. Valid values: Enabled, Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The optimization task type.
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

func (s GetOptimizeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOptimizeTaskResponseBody) GetCheckpoint() map[string]interface{} {
	return s.Checkpoint
}

func (s *GetOptimizeTaskResponseBody) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetOptimizeTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOptimizeTaskResponseBody) GetNextTriggerTime() *string {
	return s.NextTriggerTime
}

func (s *GetOptimizeTaskResponseBody) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *GetOptimizeTaskResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOptimizeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOptimizeTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetOptimizeTaskResponseBody) GetType() *string {
	return s.Type
}

func (s *GetOptimizeTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOptimizeTaskResponseBody) SetCheckpoint(v map[string]interface{}) *GetOptimizeTaskResponseBody {
	s.Checkpoint = v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetConfig(v map[string]interface{}) *GetOptimizeTaskResponseBody {
	s.Config = v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetCreateTime(v string) *GetOptimizeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetNextTriggerTime(v string) *GetOptimizeTaskResponseBody {
	s.NextTriggerTime = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetOptimizeTaskName(v string) *GetOptimizeTaskResponseBody {
	s.OptimizeTaskName = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetRegionId(v string) *GetOptimizeTaskResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetRequestId(v string) *GetOptimizeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetStatus(v string) *GetOptimizeTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetType(v string) *GetOptimizeTaskResponseBody {
	s.Type = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) SetUpdateTime(v string) *GetOptimizeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetOptimizeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
