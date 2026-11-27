// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTokenUsageDetailDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBizObjects(v []*TokenUsageDetailDTOBizObjects) *TokenUsageDetailDTO
	GetBizObjects() []*TokenUsageDetailDTOBizObjects
	SetNamespace(v string) *TokenUsageDetailDTO
	GetNamespace() *string
	SetOccurredAt(v int64) *TokenUsageDetailDTO
	GetOccurredAt() *int64
	SetScopeConfig(v *TokenUsageDetailDTOScopeConfig) *TokenUsageDetailDTO
	GetScopeConfig() *TokenUsageDetailDTOScopeConfig
	SetScopeType(v string) *TokenUsageDetailDTO
	GetScopeType() *string
	SetSessionId(v string) *TokenUsageDetailDTO
	GetSessionId() *string
	SetStatus(v string) *TokenUsageDetailDTO
	GetStatus() *string
	SetTotalJobs(v int32) *TokenUsageDetailDTO
	GetTotalJobs() *int32
	SetTotalTokens(v int64) *TokenUsageDetailDTO
	GetTotalTokens() *int64
	SetTriggerSource(v string) *TokenUsageDetailDTO
	GetTriggerSource() *string
	SetUsageType(v string) *TokenUsageDetailDTO
	GetUsageType() *string
	SetWorkspace(v string) *TokenUsageDetailDTO
	GetWorkspace() *string
}

type TokenUsageDetailDTO struct {
	// The associated items of the intelligent assistant, such as jobs and deployments.
	BizObjects []*TokenUsageDetailDTOBizObjects `json:"BizObjects,omitempty" xml:"BizObjects,omitempty" type:"Repeated"`
	// namespace
	//
	// example:
	//
	// test-default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The time when the request is initiated, in epoch milliseconds.
	//
	// example:
	//
	// 1790651783392
	OccurredAt *int64 `json:"OccurredAt,omitempty" xml:"OccurredAt,omitempty"`
	// The inspection configuration.
	ScopeConfig *TokenUsageDetailDTOScopeConfig `json:"ScopeConfig,omitempty" xml:"ScopeConfig,omitempty" type:"Struct"`
	// The inspection type.
	//
	// example:
	//
	// TAG
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The session ID shared across multiple requests in the same session.
	//
	// example:
	//
	// conv_1790651642778_azcvjyh
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The request status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of jobs covered by the inspection.
	//
	// example:
	//
	// 12
	TotalJobs *int32 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// The number of tokens consumed by this request.
	//
	// example:
	//
	// 1287
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// The request trigger type, such as cron or manual.
	//
	// example:
	//
	// CRON
	TriggerSource *string `json:"TriggerSource,omitempty" xml:"TriggerSource,omitempty"`
	// The token consumption scenario.
	//
	// example:
	//
	// ai_assistant
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	// workspace
	//
	// example:
	//
	// c278e83dff5b41
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s TokenUsageDetailDTO) String() string {
	return dara.Prettify(s)
}

func (s TokenUsageDetailDTO) GoString() string {
	return s.String()
}

func (s *TokenUsageDetailDTO) GetBizObjects() []*TokenUsageDetailDTOBizObjects {
	return s.BizObjects
}

func (s *TokenUsageDetailDTO) GetNamespace() *string {
	return s.Namespace
}

func (s *TokenUsageDetailDTO) GetOccurredAt() *int64 {
	return s.OccurredAt
}

func (s *TokenUsageDetailDTO) GetScopeConfig() *TokenUsageDetailDTOScopeConfig {
	return s.ScopeConfig
}

func (s *TokenUsageDetailDTO) GetScopeType() *string {
	return s.ScopeType
}

func (s *TokenUsageDetailDTO) GetSessionId() *string {
	return s.SessionId
}

func (s *TokenUsageDetailDTO) GetStatus() *string {
	return s.Status
}

func (s *TokenUsageDetailDTO) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *TokenUsageDetailDTO) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *TokenUsageDetailDTO) GetTriggerSource() *string {
	return s.TriggerSource
}

func (s *TokenUsageDetailDTO) GetUsageType() *string {
	return s.UsageType
}

func (s *TokenUsageDetailDTO) GetWorkspace() *string {
	return s.Workspace
}

func (s *TokenUsageDetailDTO) SetBizObjects(v []*TokenUsageDetailDTOBizObjects) *TokenUsageDetailDTO {
	s.BizObjects = v
	return s
}

func (s *TokenUsageDetailDTO) SetNamespace(v string) *TokenUsageDetailDTO {
	s.Namespace = &v
	return s
}

func (s *TokenUsageDetailDTO) SetOccurredAt(v int64) *TokenUsageDetailDTO {
	s.OccurredAt = &v
	return s
}

func (s *TokenUsageDetailDTO) SetScopeConfig(v *TokenUsageDetailDTOScopeConfig) *TokenUsageDetailDTO {
	s.ScopeConfig = v
	return s
}

func (s *TokenUsageDetailDTO) SetScopeType(v string) *TokenUsageDetailDTO {
	s.ScopeType = &v
	return s
}

func (s *TokenUsageDetailDTO) SetSessionId(v string) *TokenUsageDetailDTO {
	s.SessionId = &v
	return s
}

func (s *TokenUsageDetailDTO) SetStatus(v string) *TokenUsageDetailDTO {
	s.Status = &v
	return s
}

func (s *TokenUsageDetailDTO) SetTotalJobs(v int32) *TokenUsageDetailDTO {
	s.TotalJobs = &v
	return s
}

func (s *TokenUsageDetailDTO) SetTotalTokens(v int64) *TokenUsageDetailDTO {
	s.TotalTokens = &v
	return s
}

func (s *TokenUsageDetailDTO) SetTriggerSource(v string) *TokenUsageDetailDTO {
	s.TriggerSource = &v
	return s
}

func (s *TokenUsageDetailDTO) SetUsageType(v string) *TokenUsageDetailDTO {
	s.UsageType = &v
	return s
}

func (s *TokenUsageDetailDTO) SetWorkspace(v string) *TokenUsageDetailDTO {
	s.Workspace = &v
	return s
}

func (s *TokenUsageDetailDTO) Validate() error {
	if s.BizObjects != nil {
		for _, item := range s.BizObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TokenUsageDetailDTOBizObjects struct {
	// The ID of the associated item.
	//
	// example:
	//
	// 4c9102f7-51d2-4494-9ea2-468e3c407ad8
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the associated item.
	//
	// example:
	//
	// deployment
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The name of the associated item.
	//
	// example:
	//
	// deployemnt-draft
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TokenUsageDetailDTOBizObjects) String() string {
	return dara.Prettify(s)
}

func (s TokenUsageDetailDTOBizObjects) GoString() string {
	return s.String()
}

func (s *TokenUsageDetailDTOBizObjects) GetId() *string {
	return s.Id
}

func (s *TokenUsageDetailDTOBizObjects) GetKind() *string {
	return s.Kind
}

func (s *TokenUsageDetailDTOBizObjects) GetName() *string {
	return s.Name
}

func (s *TokenUsageDetailDTOBizObjects) SetId(v string) *TokenUsageDetailDTOBizObjects {
	s.Id = &v
	return s
}

func (s *TokenUsageDetailDTOBizObjects) SetKind(v string) *TokenUsageDetailDTOBizObjects {
	s.Kind = &v
	return s
}

func (s *TokenUsageDetailDTOBizObjects) SetName(v string) *TokenUsageDetailDTOBizObjects {
	s.Name = &v
	return s
}

func (s *TokenUsageDetailDTOBizObjects) Validate() error {
	return dara.Validate(s)
}

type TokenUsageDetailDTOScopeConfig struct {
	// The list of deployments selected for inspection.
	DeploymentIds []*string `json:"DeploymentIds,omitempty" xml:"DeploymentIds,omitempty" type:"Repeated"`
	// The tag configuration selected for inspection.
	Tags map[string][]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TokenUsageDetailDTOScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s TokenUsageDetailDTOScopeConfig) GoString() string {
	return s.String()
}

func (s *TokenUsageDetailDTOScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *TokenUsageDetailDTOScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *TokenUsageDetailDTOScopeConfig) SetDeploymentIds(v []*string) *TokenUsageDetailDTOScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *TokenUsageDetailDTOScopeConfig) SetTags(v map[string][]*string) *TokenUsageDetailDTOScopeConfig {
	s.Tags = v
	return s
}

func (s *TokenUsageDetailDTOScopeConfig) Validate() error {
	return dara.Validate(s)
}
