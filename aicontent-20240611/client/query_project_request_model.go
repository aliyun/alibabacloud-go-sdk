// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *QueryProjectRequest
	GetProjectId() *string
}

type QueryProjectRequest struct {
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s QueryProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectRequest) SetProjectId(v string) *QueryProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectRequest) Validate() error {
	return dara.Validate(s)
}
