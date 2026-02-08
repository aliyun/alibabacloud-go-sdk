// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIdsShrink(v string) *DeleteMaterialTaskShrinkRequest
	GetTaskIdsShrink() *string
}

type DeleteMaterialTaskShrinkRequest struct {
	// This parameter is required.
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DeleteMaterialTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialTaskShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *DeleteMaterialTaskShrinkRequest) SetTaskIdsShrink(v string) *DeleteMaterialTaskShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *DeleteMaterialTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
