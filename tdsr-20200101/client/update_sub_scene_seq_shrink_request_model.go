// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneSeqShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneId(v string) *UpdateSubSceneSeqShrinkRequest
	GetSceneId() *string
	SetSortSubSceneIdsShrink(v string) *UpdateSubSceneSeqShrinkRequest
	GetSortSubSceneIdsShrink() *string
}

type UpdateSubSceneSeqShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sgyuyewyew****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SortSubSceneIdsShrink *string `json:"SortSubSceneIds,omitempty" xml:"SortSubSceneIds,omitempty"`
}

func (s UpdateSubSceneSeqShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneSeqShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneSeqShrinkRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateSubSceneSeqShrinkRequest) GetSortSubSceneIdsShrink() *string {
	return s.SortSubSceneIdsShrink
}

func (s *UpdateSubSceneSeqShrinkRequest) SetSceneId(v string) *UpdateSubSceneSeqShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateSubSceneSeqShrinkRequest) SetSortSubSceneIdsShrink(v string) *UpdateSubSceneSeqShrinkRequest {
	s.SortSubSceneIdsShrink = &v
	return s
}

func (s *UpdateSubSceneSeqShrinkRequest) Validate() error {
	return dara.Validate(s)
}
