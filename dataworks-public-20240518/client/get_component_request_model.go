// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *GetComponentRequest
	GetComponentId() *string
	SetProjectId(v int64) *GetComponentRequest
	GetProjectId() *int64
}

type GetComponentRequest struct {
	// The component ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1112312312312
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID. You must configure either this parameter or the ProjectIdentifier parameter to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComponentRequest) GoString() string {
	return s.String()
}

func (s *GetComponentRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *GetComponentRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetComponentRequest) SetComponentId(v string) *GetComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *GetComponentRequest) SetProjectId(v int64) *GetComponentRequest {
	s.ProjectId = &v
	return s
}

func (s *GetComponentRequest) Validate() error {
	return dara.Validate(s)
}
