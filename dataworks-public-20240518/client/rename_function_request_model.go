// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RenameFunctionRequest
	GetId() *int64
	SetName(v string) *RenameFunctionRequest
	GetName() *string
	SetProjectId(v int64) *RenameFunctionRequest
	GetProjectId() *int64
}

type RenameFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameFunctionRequest) GoString() string {
	return s.String()
}

func (s *RenameFunctionRequest) GetId() *int64 {
	return s.Id
}

func (s *RenameFunctionRequest) GetName() *string {
	return s.Name
}

func (s *RenameFunctionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameFunctionRequest) SetId(v int64) *RenameFunctionRequest {
	s.Id = &v
	return s
}

func (s *RenameFunctionRequest) SetName(v string) *RenameFunctionRequest {
	s.Name = &v
	return s
}

func (s *RenameFunctionRequest) SetProjectId(v int64) *RenameFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameFunctionRequest) Validate() error {
	return dara.Validate(s)
}
