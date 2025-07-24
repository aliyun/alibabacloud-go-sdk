// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSceneItem interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *GlobalSceneItem
	GetDetail() *string
	SetType(v string) *GlobalSceneItem
	GetType() *string
}

type GlobalSceneItem struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GlobalSceneItem) String() string {
	return dara.Prettify(s)
}

func (s GlobalSceneItem) GoString() string {
	return s.String()
}

func (s *GlobalSceneItem) GetDetail() *string {
	return s.Detail
}

func (s *GlobalSceneItem) GetType() *string {
	return s.Type
}

func (s *GlobalSceneItem) SetDetail(v string) *GlobalSceneItem {
	s.Detail = &v
	return s
}

func (s *GlobalSceneItem) SetType(v string) *GlobalSceneItem {
	s.Type = &v
	return s
}

func (s *GlobalSceneItem) Validate() error {
	return dara.Validate(s)
}
