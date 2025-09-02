// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *GetDagRequest
	GetDagId() *int64
	SetProjectEnv(v string) *GetDagRequest
	GetProjectEnv() *string
}

type GetDagRequest struct {
	// The ID of the DAG. You can use one of the following method to obtain the ID:
	//
	// 	- Call the [RunCycleDagNodes](https://help.aliyun.com/document_detail/2780209.html) operation and obtain the value of the **Data*	- response parameter.
	//
	// 	- Call the [RunSmokeTest](https://help.aliyun.com/document_detail/2780210.html) operation and obtain the value of the **Data*	- response parameter.
	//
	// 	- Call the [RunManualDagNodes](https://help.aliyun.com/document_detail/2780218.html) operation and obtain the value of the **DagId*	- response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123141452344525
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s GetDagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDagRequest) GoString() string {
	return s.String()
}

func (s *GetDagRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *GetDagRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetDagRequest) SetDagId(v int64) *GetDagRequest {
	s.DagId = &v
	return s
}

func (s *GetDagRequest) SetProjectEnv(v string) *GetDagRequest {
	s.ProjectEnv = &v
	return s
}

func (s *GetDagRequest) Validate() error {
	return dara.Validate(s)
}
