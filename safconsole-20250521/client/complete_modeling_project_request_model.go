// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteModelingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CompleteModelingProjectRequest
	GetProjectId() *int64
}

type CompleteModelingProjectRequest struct {
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CompleteModelingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteModelingProjectRequest) GoString() string {
	return s.String()
}

func (s *CompleteModelingProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CompleteModelingProjectRequest) SetProjectId(v int64) *CompleteModelingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *CompleteModelingProjectRequest) Validate() error {
	return dara.Validate(s)
}
