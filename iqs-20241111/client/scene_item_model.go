// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSceneItem interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *SceneItem
	GetDetail() *string
	SetType(v string) *SceneItem
	GetType() *string
}

type SceneItem struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SceneItem) String() string {
	return dara.Prettify(s)
}

func (s SceneItem) GoString() string {
	return s.String()
}

func (s *SceneItem) GetDetail() *string {
	return s.Detail
}

func (s *SceneItem) GetType() *string {
	return s.Type
}

func (s *SceneItem) SetDetail(v string) *SceneItem {
	s.Detail = &v
	return s
}

func (s *SceneItem) SetType(v string) *SceneItem {
	s.Type = &v
	return s
}

func (s *SceneItem) Validate() error {
	return dara.Validate(s)
}
