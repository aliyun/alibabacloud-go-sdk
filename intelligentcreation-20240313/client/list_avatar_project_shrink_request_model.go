// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvatarProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectIdListShrink(v string) *ListAvatarProjectShrinkRequest
	GetProjectIdListShrink() *string
}

type ListAvatarProjectShrinkRequest struct {
	ProjectIdListShrink *string `json:"projectIdList,omitempty" xml:"projectIdList,omitempty"`
}

func (s ListAvatarProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvatarProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAvatarProjectShrinkRequest) GetProjectIdListShrink() *string {
	return s.ProjectIdListShrink
}

func (s *ListAvatarProjectShrinkRequest) SetProjectIdListShrink(v string) *ListAvatarProjectShrinkRequest {
	s.ProjectIdListShrink = &v
	return s
}

func (s *ListAvatarProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
