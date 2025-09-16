// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineWithAssetsCodeVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildId(v int64) *CreateRoutineWithAssetsCodeVersionShrinkRequest
	GetBuildId() *int64
	SetCodeDescription(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest
	GetCodeDescription() *string
	SetConfOptionsShrink(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest
	GetConfOptionsShrink() *string
	SetExtraInfo(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest
	GetExtraInfo() *string
	SetName(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest
	GetName() *string
}

type CreateRoutineWithAssetsCodeVersionShrinkRequest struct {
	BuildId           *int64  `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	CodeDescription   *string `json:"CodeDescription,omitempty" xml:"CodeDescription,omitempty"`
	ConfOptionsShrink *string `json:"ConfOptions,omitempty" xml:"ConfOptions,omitempty"`
	ExtraInfo         *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) GetBuildId() *int64 {
	return s.BuildId
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) GetCodeDescription() *string {
	return s.CodeDescription
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) GetConfOptionsShrink() *string {
	return s.ConfOptionsShrink
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) SetBuildId(v int64) *CreateRoutineWithAssetsCodeVersionShrinkRequest {
	s.BuildId = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) SetCodeDescription(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest {
	s.CodeDescription = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) SetConfOptionsShrink(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest {
	s.ConfOptionsShrink = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) SetExtraInfo(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest {
	s.ExtraInfo = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) SetName(v string) *CreateRoutineWithAssetsCodeVersionShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
