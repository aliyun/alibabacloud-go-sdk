// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluator(v *GetEvaluatorResponseBodyEvaluator) *GetEvaluatorResponseBody
	GetEvaluator() *GetEvaluatorResponseBodyEvaluator
	SetRequestId(v string) *GetEvaluatorResponseBody
	GetRequestId() *string
}

type GetEvaluatorResponseBody struct {
	// The evaluator details.
	//
	// example:
	//
	// {"name":"trace_task_completion","type":"AGENT","currentVersion":"1.0.0"}
	Evaluator *GetEvaluatorResponseBodyEvaluator `json:"evaluator,omitempty" xml:"evaluator,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEvaluatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluatorResponseBody) GetEvaluator() *GetEvaluatorResponseBodyEvaluator {
	return s.Evaluator
}

func (s *GetEvaluatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEvaluatorResponseBody) SetEvaluator(v *GetEvaluatorResponseBodyEvaluator) *GetEvaluatorResponseBody {
	s.Evaluator = v
	return s
}

func (s *GetEvaluatorResponseBody) SetRequestId(v string) *GetEvaluatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEvaluatorResponseBody) Validate() error {
	if s.Evaluator != nil {
		if err := s.Evaluator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEvaluatorResponseBodyEvaluator struct {
	// The AgentSpace name.
	//
	// example:
	//
	// prod-agentspace
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The list of annotations.
	//
	// example:
	//
	// ["__en"]
	Annotations []*string `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// The configuration of the current version.
	//
	// example:
	//
	// {"prompt":"请评估任务完成度"}
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the evaluator was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The version number returned in the current response.
	//
	// example:
	//
	// 1.0.0
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The evaluator description.
	//
	// example:
	//
	// 判断 Agent 是否完成用户任务
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name.
	//
	// example:
	//
	// 链路任务完成度
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// 1.0.0
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The evaluation metric name.
	//
	// example:
	//
	// agent_task_completion
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The evaluator name.
	//
	// example:
	//
	// trace_task_completion
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The evaluator properties.
	//
	// example:
	//
	// {"agentEvaluatorMode":"raw_prompt"}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The evaluator type.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the evaluator was last updated. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The list of versions.
	//
	// example:
	//
	// [{"version":"1.0.0"}]
	Versions []*GetEvaluatorResponseBodyEvaluatorVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetEvaluatorResponseBodyEvaluator) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorResponseBodyEvaluator) GoString() string {
	return s.String()
}

func (s *GetEvaluatorResponseBodyEvaluator) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetEvaluatorResponseBodyEvaluator) GetAnnotations() []*string {
	return s.Annotations
}

func (s *GetEvaluatorResponseBodyEvaluator) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetEvaluatorResponseBodyEvaluator) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluatorResponseBodyEvaluator) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *GetEvaluatorResponseBodyEvaluator) GetDescription() *string {
	return s.Description
}

func (s *GetEvaluatorResponseBodyEvaluator) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEvaluatorResponseBodyEvaluator) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetEvaluatorResponseBodyEvaluator) GetMetricName() *string {
	return s.MetricName
}

func (s *GetEvaluatorResponseBodyEvaluator) GetName() *string {
	return s.Name
}

func (s *GetEvaluatorResponseBodyEvaluator) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetEvaluatorResponseBodyEvaluator) GetType() *string {
	return s.Type
}

func (s *GetEvaluatorResponseBodyEvaluator) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetEvaluatorResponseBodyEvaluator) GetVersions() []*GetEvaluatorResponseBodyEvaluatorVersions {
	return s.Versions
}

func (s *GetEvaluatorResponseBodyEvaluator) SetAgentSpace(v string) *GetEvaluatorResponseBodyEvaluator {
	s.AgentSpace = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetAnnotations(v []*string) *GetEvaluatorResponseBodyEvaluator {
	s.Annotations = v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetConfig(v map[string]interface{}) *GetEvaluatorResponseBodyEvaluator {
	s.Config = v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetCreatedAt(v int64) *GetEvaluatorResponseBodyEvaluator {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetCurrentVersion(v string) *GetEvaluatorResponseBodyEvaluator {
	s.CurrentVersion = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetDescription(v string) *GetEvaluatorResponseBodyEvaluator {
	s.Description = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetDisplayName(v string) *GetEvaluatorResponseBodyEvaluator {
	s.DisplayName = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetLatestVersion(v string) *GetEvaluatorResponseBodyEvaluator {
	s.LatestVersion = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetMetricName(v string) *GetEvaluatorResponseBodyEvaluator {
	s.MetricName = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetName(v string) *GetEvaluatorResponseBodyEvaluator {
	s.Name = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetProperties(v map[string]interface{}) *GetEvaluatorResponseBodyEvaluator {
	s.Properties = v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetType(v string) *GetEvaluatorResponseBodyEvaluator {
	s.Type = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetUpdatedAt(v int64) *GetEvaluatorResponseBodyEvaluator {
	s.UpdatedAt = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) SetVersions(v []*GetEvaluatorResponseBodyEvaluatorVersions) *GetEvaluatorResponseBodyEvaluator {
	s.Versions = v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluator) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEvaluatorResponseBodyEvaluatorVersions struct {
	// The time when the version was created. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description.
	//
	// example:
	//
	// 初始版本
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s GetEvaluatorResponseBodyEvaluatorVersions) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorResponseBodyEvaluatorVersions) GoString() string {
	return s.String()
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) GetVersion() *string {
	return s.Version
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) SetCreatedAt(v int64) *GetEvaluatorResponseBodyEvaluatorVersions {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) SetVersion(v string) *GetEvaluatorResponseBodyEvaluatorVersions {
	s.Version = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) SetVersionDescription(v string) *GetEvaluatorResponseBodyEvaluatorVersions {
	s.VersionDescription = &v
	return s
}

func (s *GetEvaluatorResponseBodyEvaluatorVersions) Validate() error {
	return dara.Validate(s)
}
