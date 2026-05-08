// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectIdList(v []*string) *ListAvatarProjectRequest
	GetProjectIdList() []*string
}

type ListAvatarProjectRequest struct {
	ProjectIdList []*string `json:"projectIdList,omitempty" xml:"projectIdList,omitempty" type:"Repeated"`
}

func (s ListAvatarProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarProjectRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectRequest) GetProjectIdList() []*string {
	return s.ProjectIdList
}

func (s *ListAvatarProjectRequest) SetProjectIdList(v []*string) *ListAvatarProjectRequest {
	s.ProjectIdList = v
	return s
}

func (s *ListAvatarProjectRequest) Validate() error {
	return dara.Validate(s)
}
