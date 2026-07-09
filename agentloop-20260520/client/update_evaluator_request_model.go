// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v []*string) *UpdateEvaluatorRequest
	GetAnnotations() []*string
	SetConfig(v map[string]interface{}) *UpdateEvaluatorRequest
	GetConfig() map[string]interface{}
	SetDescription(v string) *UpdateEvaluatorRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateEvaluatorRequest
	GetDisplayName() *string
	SetProperties(v map[string]interface{}) *UpdateEvaluatorRequest
	GetProperties() map[string]interface{}
	SetVersion(v string) *UpdateEvaluatorRequest
	GetVersion() *string
	SetVersionDescription(v string) *UpdateEvaluatorRequest
	GetVersionDescription() *string
	SetClientToken(v string) *UpdateEvaluatorRequest
	GetClientToken() *string
}

type UpdateEvaluatorRequest struct {
	// The list of annotation marks.
	//
	// example:
	//
	// ["__en"]
	Annotations []*string `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// The new version configuration. This parameter is typically required when `version` is specified.
	//
	// example:
	//
	// {"prompt":"请更严格地评估任务完成度"}
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
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
	// The evaluator properties.
	//
	// example:
	//
	// {"agentEvaluatorMode":"raw_prompt"}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The new version number. If specified, a new version is created.
	//
	// example:
	//
	// 1.1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The version description.
	//
	// example:
	//
	// 优化评分说明
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
	// The idempotency token. CloudSpec declares this query parameter, but the backend does not currently perform idempotency comparison.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateEvaluatorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorRequest) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorRequest) GetAnnotations() []*string {
	return s.Annotations
}

func (s *UpdateEvaluatorRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateEvaluatorRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEvaluatorRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateEvaluatorRequest) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *UpdateEvaluatorRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateEvaluatorRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateEvaluatorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEvaluatorRequest) SetAnnotations(v []*string) *UpdateEvaluatorRequest {
	s.Annotations = v
	return s
}

func (s *UpdateEvaluatorRequest) SetConfig(v map[string]interface{}) *UpdateEvaluatorRequest {
	s.Config = v
	return s
}

func (s *UpdateEvaluatorRequest) SetDescription(v string) *UpdateEvaluatorRequest {
	s.Description = &v
	return s
}

func (s *UpdateEvaluatorRequest) SetDisplayName(v string) *UpdateEvaluatorRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateEvaluatorRequest) SetProperties(v map[string]interface{}) *UpdateEvaluatorRequest {
	s.Properties = v
	return s
}

func (s *UpdateEvaluatorRequest) SetVersion(v string) *UpdateEvaluatorRequest {
	s.Version = &v
	return s
}

func (s *UpdateEvaluatorRequest) SetVersionDescription(v string) *UpdateEvaluatorRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateEvaluatorRequest) SetClientToken(v string) *UpdateEvaluatorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEvaluatorRequest) Validate() error {
	return dara.Validate(s)
}
