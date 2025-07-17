// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLineageRelationshipShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstEntityShrink(v string) *CreateLineageRelationshipShrinkRequest
	GetDstEntityShrink() *string
	SetSrcEntityShrink(v string) *CreateLineageRelationshipShrinkRequest
	GetSrcEntityShrink() *string
	SetTaskShrink(v string) *CreateLineageRelationshipShrinkRequest
	GetTaskShrink() *string
}

type CreateLineageRelationshipShrinkRequest struct {
	DstEntityShrink *string `json:"DstEntity,omitempty" xml:"DstEntity,omitempty"`
	SrcEntityShrink *string `json:"SrcEntity,omitempty" xml:"SrcEntity,omitempty"`
	// The task information.
	TaskShrink *string `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s CreateLineageRelationshipShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLineageRelationshipShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLineageRelationshipShrinkRequest) GetDstEntityShrink() *string {
	return s.DstEntityShrink
}

func (s *CreateLineageRelationshipShrinkRequest) GetSrcEntityShrink() *string {
	return s.SrcEntityShrink
}

func (s *CreateLineageRelationshipShrinkRequest) GetTaskShrink() *string {
	return s.TaskShrink
}

func (s *CreateLineageRelationshipShrinkRequest) SetDstEntityShrink(v string) *CreateLineageRelationshipShrinkRequest {
	s.DstEntityShrink = &v
	return s
}

func (s *CreateLineageRelationshipShrinkRequest) SetSrcEntityShrink(v string) *CreateLineageRelationshipShrinkRequest {
	s.SrcEntityShrink = &v
	return s
}

func (s *CreateLineageRelationshipShrinkRequest) SetTaskShrink(v string) *CreateLineageRelationshipShrinkRequest {
	s.TaskShrink = &v
	return s
}

func (s *CreateLineageRelationshipShrinkRequest) Validate() error {
	return dara.Validate(s)
}
