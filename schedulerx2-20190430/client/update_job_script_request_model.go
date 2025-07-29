// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateJobScriptRequest
	GetGroupId() *string
	SetJobId(v int64) *UpdateJobScriptRequest
	GetJobId() *int64
	SetNamespace(v string) *UpdateJobScriptRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *UpdateJobScriptRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *UpdateJobScriptRequest
	GetRegionId() *string
	SetScriptContent(v string) *UpdateJobScriptRequest
	GetScriptContent() *string
	SetVersionDescription(v string) *UpdateJobScriptRequest
	GetVersionDescription() *string
}

type UpdateJobScriptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 301
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScriptContent      *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateJobScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobScriptRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobScriptRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateJobScriptRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateJobScriptRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateJobScriptRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *UpdateJobScriptRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateJobScriptRequest) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *UpdateJobScriptRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateJobScriptRequest) SetGroupId(v string) *UpdateJobScriptRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateJobScriptRequest) SetJobId(v int64) *UpdateJobScriptRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobScriptRequest) SetNamespace(v string) *UpdateJobScriptRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateJobScriptRequest) SetNamespaceSource(v string) *UpdateJobScriptRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateJobScriptRequest) SetRegionId(v string) *UpdateJobScriptRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateJobScriptRequest) SetScriptContent(v string) *UpdateJobScriptRequest {
	s.ScriptContent = &v
	return s
}

func (s *UpdateJobScriptRequest) SetVersionDescription(v string) *UpdateJobScriptRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateJobScriptRequest) Validate() error {
	return dara.Validate(s)
}
