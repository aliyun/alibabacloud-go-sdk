// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MoveFunctionRequest
	GetId() *int64
	SetPath(v string) *MoveFunctionRequest
	GetPath() *string
	SetProjectId(v int64) *MoveFunctionRequest
	GetProjectId() *int64
}

type MoveFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the UDF. You do not need to specify a UDF name in the path.
	//
	// For example, if you want to move the test UDF to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// This parameter indicates the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveFunctionRequest) GoString() string {
	return s.String()
}

func (s *MoveFunctionRequest) GetId() *int64 {
	return s.Id
}

func (s *MoveFunctionRequest) GetPath() *string {
	return s.Path
}

func (s *MoveFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveFunctionRequest) SetId(v int64) *MoveFunctionRequest {
	s.Id = &v
	return s
}

func (s *MoveFunctionRequest) SetPath(v string) *MoveFunctionRequest {
	s.Path = &v
	return s
}

func (s *MoveFunctionRequest) SetProjectId(v int64) *MoveFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveFunctionRequest) Validate() error {
	return dara.Validate(s)
}
