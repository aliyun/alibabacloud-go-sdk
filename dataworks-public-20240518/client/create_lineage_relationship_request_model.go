// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLineageRelationshipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDstEntity(v *LineageEntity) *CreateLineageRelationshipRequest
	GetDstEntity() *LineageEntity
	SetSrcEntity(v *LineageEntity) *CreateLineageRelationshipRequest
	GetSrcEntity() *LineageEntity
	SetTask(v *LineageTask) *CreateLineageRelationshipRequest
	GetTask() *LineageTask
}

type CreateLineageRelationshipRequest struct {
	DstEntity *LineageEntity `json:"DstEntity,omitempty" xml:"DstEntity,omitempty"`
	SrcEntity *LineageEntity `json:"SrcEntity,omitempty" xml:"SrcEntity,omitempty"`
	// The task information.
	Task *LineageTask `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s CreateLineageRelationshipRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLineageRelationshipRequest) GoString() string {
	return s.String()
}

func (s *CreateLineageRelationshipRequest) GetDstEntity() *LineageEntity {
	return s.DstEntity
}

func (s *CreateLineageRelationshipRequest) GetSrcEntity() *LineageEntity {
	return s.SrcEntity
}

func (s *CreateLineageRelationshipRequest) GetTask() *LineageTask {
	return s.Task
}

func (s *CreateLineageRelationshipRequest) SetDstEntity(v *LineageEntity) *CreateLineageRelationshipRequest {
	s.DstEntity = v
	return s
}

func (s *CreateLineageRelationshipRequest) SetSrcEntity(v *LineageEntity) *CreateLineageRelationshipRequest {
	s.SrcEntity = v
	return s
}

func (s *CreateLineageRelationshipRequest) SetTask(v *LineageTask) *CreateLineageRelationshipRequest {
	s.Task = v
	return s
}

func (s *CreateLineageRelationshipRequest) Validate() error {
	return dara.Validate(s)
}
