// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabaseSummaryModel interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DatabaseSummaryModel
	GetCreateTime() *string
	SetDescription(v string) *DatabaseSummaryModel
	GetDescription() *string
	SetOwner(v string) *DatabaseSummaryModel
	GetOwner() *string
	SetSchemaName(v string) *DatabaseSummaryModel
	GetSchemaName() *string
	SetUpdateTime(v string) *DatabaseSummaryModel
	GetUpdateTime() *string
}

type DatabaseSummaryModel struct {
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SchemaName  *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DatabaseSummaryModel) String() string {
	return dara.Prettify(s)
}

func (s DatabaseSummaryModel) GoString() string {
	return s.String()
}

func (s *DatabaseSummaryModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DatabaseSummaryModel) GetDescription() *string {
	return s.Description
}

func (s *DatabaseSummaryModel) GetOwner() *string {
	return s.Owner
}

func (s *DatabaseSummaryModel) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DatabaseSummaryModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DatabaseSummaryModel) SetCreateTime(v string) *DatabaseSummaryModel {
	s.CreateTime = &v
	return s
}

func (s *DatabaseSummaryModel) SetDescription(v string) *DatabaseSummaryModel {
	s.Description = &v
	return s
}

func (s *DatabaseSummaryModel) SetOwner(v string) *DatabaseSummaryModel {
	s.Owner = &v
	return s
}

func (s *DatabaseSummaryModel) SetSchemaName(v string) *DatabaseSummaryModel {
	s.SchemaName = &v
	return s
}

func (s *DatabaseSummaryModel) SetUpdateTime(v string) *DatabaseSummaryModel {
	s.UpdateTime = &v
	return s
}

func (s *DatabaseSummaryModel) Validate() error {
	return dara.Validate(s)
}
