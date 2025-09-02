// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualDagInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v string) *ListManualDagInstancesRequest
	GetDagId() *string
	SetProjectEnv(v string) *ListManualDagInstancesRequest
	GetProjectEnv() *string
	SetProjectName(v string) *ListManualDagInstancesRequest
	GetProjectName() *string
}

type ListManualDagInstancesRequest struct {
	// The ID of the directed acyclic graph (DAG) for the manually triggered workflow. You can call the [RunManualDagNodes](https://help.aliyun.com/document_detail/212830.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000001231241
	DagId *string `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The environment of Operation Center. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// RPOD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The name of the workspace to which the manually triggered workflow belongs. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workspace
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListManualDagInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManualDagInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListManualDagInstancesRequest) GetDagId() *string {
	return s.DagId
}

func (s *ListManualDagInstancesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListManualDagInstancesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListManualDagInstancesRequest) SetDagId(v string) *ListManualDagInstancesRequest {
	s.DagId = &v
	return s
}

func (s *ListManualDagInstancesRequest) SetProjectEnv(v string) *ListManualDagInstancesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListManualDagInstancesRequest) SetProjectName(v string) *ListManualDagInstancesRequest {
	s.ProjectName = &v
	return s
}

func (s *ListManualDagInstancesRequest) Validate() error {
	return dara.Validate(s)
}
