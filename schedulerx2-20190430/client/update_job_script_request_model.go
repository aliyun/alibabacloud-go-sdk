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
	// The application ID. You can obtain the application ID on the Applications page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the ID on the Tasks page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 301
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespaces page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The script content.
	//
	// example:
	//
	// #!/bin/bash
	//
	// # The following are predefined variables provided by the system. You can use them to obtain information about the job run.
	//
	// echo "Job parameters: #{schedulerx.jobParameters}"
	//
	// echo "Shard index: #{schedulerx.shardingId}"
	//
	// echo "Shard parameters: #{schedulerx.shardingParameters}"
	//
	// echo "Total number of shards: #{schedulerx.shardingNum}"
	//
	// echo "Current retry count: #{schedulerx.attempt}"
	//
	// echo "Trigger type: #{schedulerx.triggerType}"
	//
	// echo "Scheduled timestamp: #{schedulerx.scheduleTime}"
	//
	// echo "Data timestamp: #{schedulerx.dataTime}"
	//
	// # The output of the last line will be returned as the result
	//
	// echo "hello world"
	//
	// # exit 1 indicates failure
	//
	// exit 0
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// The description of the script version.
	//
	// example:
	//
	// Print job running information
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
