// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossAccountsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTargetsShrink(v string) *UpdateCrossAccountsShrinkRequest
	GetCreateTargetsShrink() *string
	SetDeleteTargetsShrink(v string) *UpdateCrossAccountsShrinkRequest
	GetDeleteTargetsShrink() *string
}

type UpdateCrossAccountsShrinkRequest struct {
	CreateTargetsShrink *string `json:"CreateTargets,omitempty" xml:"CreateTargets,omitempty"`
	DeleteTargetsShrink *string `json:"DeleteTargets,omitempty" xml:"DeleteTargets,omitempty"`
}

func (s UpdateCrossAccountsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsShrinkRequest) GetCreateTargetsShrink() *string {
	return s.CreateTargetsShrink
}

func (s *UpdateCrossAccountsShrinkRequest) GetDeleteTargetsShrink() *string {
	return s.DeleteTargetsShrink
}

func (s *UpdateCrossAccountsShrinkRequest) SetCreateTargetsShrink(v string) *UpdateCrossAccountsShrinkRequest {
	s.CreateTargetsShrink = &v
	return s
}

func (s *UpdateCrossAccountsShrinkRequest) SetDeleteTargetsShrink(v string) *UpdateCrossAccountsShrinkRequest {
	s.DeleteTargetsShrink = &v
	return s
}

func (s *UpdateCrossAccountsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
