// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *ListWorkersShrinkRequest
	GetAgentType() *string
	SetCredential(v string) *ListWorkersShrinkRequest
	GetCredential() *string
	SetGroupShrink(v string) *ListWorkersShrinkRequest
	GetGroupShrink() *string
	SetInstanceId(v string) *ListWorkersShrinkRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListWorkersShrinkRequest
	GetMaxResults() *int32
	SetMcp(v string) *ListWorkersShrinkRequest
	GetMcp() *string
	SetModelName(v string) *ListWorkersShrinkRequest
	GetModelName() *string
	SetModelProvider(v string) *ListWorkersShrinkRequest
	GetModelProvider() *string
	SetNameLike(v string) *ListWorkersShrinkRequest
	GetNameLike() *string
	SetNextToken(v string) *ListWorkersShrinkRequest
	GetNextToken() *string
	SetTemplateShrink(v string) *ListWorkersShrinkRequest
	GetTemplateShrink() *string
	SetVersionCode(v string) *ListWorkersShrinkRequest
	GetVersionCode() *string
}

type ListWorkersShrinkRequest struct {
	AgentType   *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Credential  *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	GroupShrink *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// This parameter is required.
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Mcp            *string `json:"Mcp,omitempty" xml:"Mcp,omitempty"`
	ModelName      *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider  *string `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
	NameLike       *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TemplateShrink *string `json:"Template,omitempty" xml:"Template,omitempty"`
	VersionCode    *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListWorkersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkersShrinkRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *ListWorkersShrinkRequest) GetCredential() *string {
	return s.Credential
}

func (s *ListWorkersShrinkRequest) GetGroupShrink() *string {
	return s.GroupShrink
}

func (s *ListWorkersShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkersShrinkRequest) GetMcp() *string {
	return s.Mcp
}

func (s *ListWorkersShrinkRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListWorkersShrinkRequest) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *ListWorkersShrinkRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListWorkersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkersShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *ListWorkersShrinkRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListWorkersShrinkRequest) SetAgentType(v string) *ListWorkersShrinkRequest {
	s.AgentType = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetCredential(v string) *ListWorkersShrinkRequest {
	s.Credential = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetGroupShrink(v string) *ListWorkersShrinkRequest {
	s.GroupShrink = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetInstanceId(v string) *ListWorkersShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetMaxResults(v int32) *ListWorkersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetMcp(v string) *ListWorkersShrinkRequest {
	s.Mcp = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetModelName(v string) *ListWorkersShrinkRequest {
	s.ModelName = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetModelProvider(v string) *ListWorkersShrinkRequest {
	s.ModelProvider = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetNameLike(v string) *ListWorkersShrinkRequest {
	s.NameLike = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetNextToken(v string) *ListWorkersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetTemplateShrink(v string) *ListWorkersShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *ListWorkersShrinkRequest) SetVersionCode(v string) *ListWorkersShrinkRequest {
	s.VersionCode = &v
	return s
}

func (s *ListWorkersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
