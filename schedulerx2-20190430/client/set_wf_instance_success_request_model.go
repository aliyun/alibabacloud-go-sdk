// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWfInstanceSuccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *SetWfInstanceSuccessRequest
	GetGroupId() *string
	SetNamespace(v string) *SetWfInstanceSuccessRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *SetWfInstanceSuccessRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *SetWfInstanceSuccessRequest
	GetRegionId() *string
	SetWfInstanceId(v int64) *SetWfInstanceSuccessRequest
	GetWfInstanceId() *int64
	SetWorkflowId(v int64) *SetWfInstanceSuccessRequest
	GetWorkflowId() *int64
}

type SetWfInstanceSuccessRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
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
	// The workflow instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s SetWfInstanceSuccessRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWfInstanceSuccessRequest) GoString() string {
	return s.String()
}

func (s *SetWfInstanceSuccessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetWfInstanceSuccessRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SetWfInstanceSuccessRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *SetWfInstanceSuccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetWfInstanceSuccessRequest) GetWfInstanceId() *int64 {
	return s.WfInstanceId
}

func (s *SetWfInstanceSuccessRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *SetWfInstanceSuccessRequest) SetGroupId(v string) *SetWfInstanceSuccessRequest {
	s.GroupId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetNamespace(v string) *SetWfInstanceSuccessRequest {
	s.Namespace = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetNamespaceSource(v string) *SetWfInstanceSuccessRequest {
	s.NamespaceSource = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetRegionId(v string) *SetWfInstanceSuccessRequest {
	s.RegionId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetWfInstanceId(v int64) *SetWfInstanceSuccessRequest {
	s.WfInstanceId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) SetWorkflowId(v int64) *SetWfInstanceSuccessRequest {
	s.WorkflowId = &v
	return s
}

func (s *SetWfInstanceSuccessRequest) Validate() error {
	return dara.Validate(s)
}
