// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManualDagInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v string) *GetManualDagInstancesRequest
	GetDagId() *string
	SetProjectEnv(v string) *GetManualDagInstancesRequest
	GetProjectEnv() *string
	SetProjectName(v string) *GetManualDagInstancesRequest
	GetProjectName() *string
}

type GetManualDagInstancesRequest struct {
	// The ID of the directed acyclic graph (DAG) for the manually triggered workflow. You can call the [CreateManualDag](https://help.aliyun.com/document_detail/189728.html) operation to query the ID.
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
	// The name of the workspace to which the manually triggered workflow belongs. You can log on to the DataWorks console and go to the Workspace Settings panel to query the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_workspace
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetManualDagInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManualDagInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetManualDagInstancesRequest) GetDagId() *string {
	return s.DagId
}

func (s *GetManualDagInstancesRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetManualDagInstancesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetManualDagInstancesRequest) SetDagId(v string) *GetManualDagInstancesRequest {
	s.DagId = &v
	return s
}

func (s *GetManualDagInstancesRequest) SetProjectEnv(v string) *GetManualDagInstancesRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetManualDagInstancesRequest) SetProjectName(v string) *GetManualDagInstancesRequest {
	s.ProjectName = &v
	return s
}

func (s *GetManualDagInstancesRequest) Validate() error {
	return dara.Validate(s)
}
