// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagJson(v string) *UpdateWorkflowDagRequest
	GetDagJson() *string
	SetGroupId(v string) *UpdateWorkflowDagRequest
	GetGroupId() *string
	SetNamespace(v string) *UpdateWorkflowDagRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *UpdateWorkflowDagRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *UpdateWorkflowDagRequest
	GetRegionId() *string
	SetWorkflowId(v string) *UpdateWorkflowDagRequest
	GetWorkflowId() *string
}

type UpdateWorkflowDagRequest struct {
	// The definition of the workflow\\"s directed acyclic graph (DAG), including nodes and edges, as a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"nodes":[{"id":2300691},{"id":10518089},{"id":1758851}],"edges":[{"source":10518089,"target":1758851},{"source":10518089,"target":2300691}]}
	DagJson *string `json:"DagJson,omitempty" xml:"DagJson,omitempty"`
	// The Application Group ID. You can find this ID on the **Application Management*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The Namespace ID. You can obtain the ID on the **Namespaces*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required only for specific third-party integrations.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowDagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDagRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDagRequest) GetDagJson() *string {
	return s.DagJson
}

func (s *UpdateWorkflowDagRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateWorkflowDagRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateWorkflowDagRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *UpdateWorkflowDagRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWorkflowDagRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *UpdateWorkflowDagRequest) SetDagJson(v string) *UpdateWorkflowDagRequest {
	s.DagJson = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetGroupId(v string) *UpdateWorkflowDagRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetNamespace(v string) *UpdateWorkflowDagRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetNamespaceSource(v string) *UpdateWorkflowDagRequest {
	s.NamespaceSource = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetRegionId(v string) *UpdateWorkflowDagRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkflowDagRequest) SetWorkflowId(v string) *UpdateWorkflowDagRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowDagRequest) Validate() error {
	return dara.Validate(s)
}
