// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *QueryAvatarProjectRequest
	GetProjectId() *string
}

type QueryAvatarProjectRequest struct {
	// example:
	//
	// 11111
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s QueryAvatarProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryAvatarProjectRequest) SetProjectId(v string) *QueryAvatarProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryAvatarProjectRequest) Validate() error {
	return dara.Validate(s)
}
