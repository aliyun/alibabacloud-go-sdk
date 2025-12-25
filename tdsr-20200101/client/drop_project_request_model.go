// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *DropProjectRequest
	GetProjectId() *string
}

type DropProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3242****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DropProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DropProjectRequest) GoString() string {
	return s.String()
}

func (s *DropProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DropProjectRequest) SetProjectId(v string) *DropProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DropProjectRequest) Validate() error {
	return dara.Validate(s)
}
