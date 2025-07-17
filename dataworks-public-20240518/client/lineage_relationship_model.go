// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageRelationship interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *LineageRelationship
	GetCreateTime() *int64
	SetDstEntity(v *LineageEntity) *LineageRelationship
	GetDstEntity() *LineageEntity
	SetId(v string) *LineageRelationship
	GetId() *string
	SetSrcEntity(v *LineageEntity) *LineageRelationship
	GetSrcEntity() *LineageEntity
	SetTask(v *LineageTask) *LineageRelationship
	GetTask() *LineageTask
}

type LineageRelationship struct {
	// example:
	//
	// 1743040581000
	CreateTime *int64         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DstEntity  *LineageEntity `json:"DstEntity,omitempty" xml:"DstEntity,omitempty"`
	// example:
	//
	// maxcompute-table.p.table:custom-table.xxx:custom-sql.123
	Id        *string        `json:"Id,omitempty" xml:"Id,omitempty"`
	SrcEntity *LineageEntity `json:"SrcEntity,omitempty" xml:"SrcEntity,omitempty"`
	Task      *LineageTask   `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s LineageRelationship) String() string {
	return dara.Prettify(s)
}

func (s LineageRelationship) GoString() string {
	return s.String()
}

func (s *LineageRelationship) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *LineageRelationship) GetDstEntity() *LineageEntity {
	return s.DstEntity
}

func (s *LineageRelationship) GetId() *string {
	return s.Id
}

func (s *LineageRelationship) GetSrcEntity() *LineageEntity {
	return s.SrcEntity
}

func (s *LineageRelationship) GetTask() *LineageTask {
	return s.Task
}

func (s *LineageRelationship) SetCreateTime(v int64) *LineageRelationship {
	s.CreateTime = &v
	return s
}

func (s *LineageRelationship) SetDstEntity(v *LineageEntity) *LineageRelationship {
	s.DstEntity = v
	return s
}

func (s *LineageRelationship) SetId(v string) *LineageRelationship {
	s.Id = &v
	return s
}

func (s *LineageRelationship) SetSrcEntity(v *LineageEntity) *LineageRelationship {
	s.SrcEntity = v
	return s
}

func (s *LineageRelationship) SetTask(v *LineageTask) *LineageRelationship {
	s.Task = v
	return s
}

func (s *LineageRelationship) Validate() error {
	return dara.Validate(s)
}
