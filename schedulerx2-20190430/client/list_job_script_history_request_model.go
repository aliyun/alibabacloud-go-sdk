// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobScriptHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListJobScriptHistoryRequest
	GetGroupId() *string
	SetJobId(v int64) *ListJobScriptHistoryRequest
	GetJobId() *int64
	SetNamespace(v string) *ListJobScriptHistoryRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListJobScriptHistoryRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ListJobScriptHistoryRequest
	GetRegionId() *string
}

type ListJobScriptHistoryRequest struct {
	// The application ID. You can obtain the application ID on the Applications page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the Tasks page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespaces page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
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
}

func (s ListJobScriptHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobScriptHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListJobScriptHistoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListJobScriptHistoryRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobScriptHistoryRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListJobScriptHistoryRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListJobScriptHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobScriptHistoryRequest) SetGroupId(v string) *ListJobScriptHistoryRequest {
	s.GroupId = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetJobId(v int64) *ListJobScriptHistoryRequest {
	s.JobId = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetNamespace(v string) *ListJobScriptHistoryRequest {
	s.Namespace = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetNamespaceSource(v string) *ListJobScriptHistoryRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListJobScriptHistoryRequest) SetRegionId(v string) *ListJobScriptHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobScriptHistoryRequest) Validate() error {
	return dara.Validate(s)
}
