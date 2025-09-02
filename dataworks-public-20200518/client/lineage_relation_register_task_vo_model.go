// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLineageRelationRegisterTaskVO interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *LineageRelationRegisterTaskVO
	GetAttributes() map[string]*string
	SetCreateTimestamp(v int64) *LineageRelationRegisterTaskVO
	GetCreateTimestamp() *int64
	SetInputEntities(v []*LineageEntityVO) *LineageRelationRegisterTaskVO
	GetInputEntities() []*LineageEntityVO
	SetName(v string) *LineageRelationRegisterTaskVO
	GetName() *string
	SetOutputEntities(v []*LineageEntityVO) *LineageRelationRegisterTaskVO
	GetOutputEntities() []*LineageEntityVO
	SetQualifiedName(v string) *LineageRelationRegisterTaskVO
	GetQualifiedName() *string
}

type LineageRelationRegisterTaskVO struct {
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 1684327487964
	CreateTimestamp *int64             `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	InputEntities   []*LineageEntityVO `json:"InputEntities,omitempty" xml:"InputEntities,omitempty" type:"Repeated"`
	// example:
	//
	// pai-task name
	Name           *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputEntities []*LineageEntityVO `json:"OutputEntities,omitempty" xml:"OutputEntities,omitempty" type:"Repeated"`
	// example:
	//
	// pai_dlcjob-task.12304
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s LineageRelationRegisterTaskVO) String() string {
	return dara.Prettify(s)
}

func (s LineageRelationRegisterTaskVO) GoString() string {
	return s.String()
}

func (s *LineageRelationRegisterTaskVO) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *LineageRelationRegisterTaskVO) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *LineageRelationRegisterTaskVO) GetInputEntities() []*LineageEntityVO {
	return s.InputEntities
}

func (s *LineageRelationRegisterTaskVO) GetName() *string {
	return s.Name
}

func (s *LineageRelationRegisterTaskVO) GetOutputEntities() []*LineageEntityVO {
	return s.OutputEntities
}

func (s *LineageRelationRegisterTaskVO) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *LineageRelationRegisterTaskVO) SetAttributes(v map[string]*string) *LineageRelationRegisterTaskVO {
	s.Attributes = v
	return s
}

func (s *LineageRelationRegisterTaskVO) SetCreateTimestamp(v int64) *LineageRelationRegisterTaskVO {
	s.CreateTimestamp = &v
	return s
}

func (s *LineageRelationRegisterTaskVO) SetInputEntities(v []*LineageEntityVO) *LineageRelationRegisterTaskVO {
	s.InputEntities = v
	return s
}

func (s *LineageRelationRegisterTaskVO) SetName(v string) *LineageRelationRegisterTaskVO {
	s.Name = &v
	return s
}

func (s *LineageRelationRegisterTaskVO) SetOutputEntities(v []*LineageEntityVO) *LineageRelationRegisterTaskVO {
	s.OutputEntities = v
	return s
}

func (s *LineageRelationRegisterTaskVO) SetQualifiedName(v string) *LineageRelationRegisterTaskVO {
	s.QualifiedName = &v
	return s
}

func (s *LineageRelationRegisterTaskVO) Validate() error {
	return dara.Validate(s)
}
