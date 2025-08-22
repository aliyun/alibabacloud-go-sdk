// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnRealTimeLogProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DeleteDcdnRealTimeLogProjectRequest
	GetProjectName() *string
}

type DeleteDcdnRealTimeLogProjectRequest struct {
	// The name of a real-time log delivery project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteDcdnRealTimeLogProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnRealTimeLogProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteDcdnRealTimeLogProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteDcdnRealTimeLogProjectRequest) SetProjectName(v string) *DeleteDcdnRealTimeLogProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteDcdnRealTimeLogProjectRequest) Validate() error {
	return dara.Validate(s)
}
