// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *GetProjectRequest
	GetProjectId() *string
}

type GetProjectRequest struct {
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetProjectRequest) SetProjectId(v string) *GetProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRequest) Validate() error {
	return dara.Validate(s)
}
