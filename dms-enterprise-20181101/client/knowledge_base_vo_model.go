// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseVO interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *KnowledgeBaseVO
	GetCategory() *string
	SetConfidence(v float64) *KnowledgeBaseVO
	GetConfidence() *float64
	SetDbId(v int32) *KnowledgeBaseVO
	GetDbId() *int32
	SetDbName(v string) *KnowledgeBaseVO
	GetDbName() *string
	SetDescription(v string) *KnowledgeBaseVO
	GetDescription() *string
	SetEntityId(v int64) *KnowledgeBaseVO
	GetEntityId() *int64
	SetEnv(v string) *KnowledgeBaseVO
	GetEnv() *string
	SetExpr(v string) *KnowledgeBaseVO
	GetExpr() *string
	SetGmtCreate(v string) *KnowledgeBaseVO
	GetGmtCreate() *string
	SetInstanceName(v string) *KnowledgeBaseVO
	GetInstanceName() *string
	SetIsDelete(v bool) *KnowledgeBaseVO
	GetIsDelete() *bool
	SetKnowledgeId(v string) *KnowledgeBaseVO
	GetKnowledgeId() *string
	SetKnowledgeType(v string) *KnowledgeBaseVO
	GetKnowledgeType() *string
	SetLevelType(v string) *KnowledgeBaseVO
	GetLevelType() *string
	SetName(v string) *KnowledgeBaseVO
	GetName() *string
	SetOldDescription(v string) *KnowledgeBaseVO
	GetOldDescription() *string
	SetOldSummary(v string) *KnowledgeBaseVO
	GetOldSummary() *string
	SetParseDesc(v string) *KnowledgeBaseVO
	GetParseDesc() *string
	SetReasoningLogic(v string) *KnowledgeBaseVO
	GetReasoningLogic() *string
	SetRelationType(v string) *KnowledgeBaseVO
	GetRelationType() *string
	SetShowType(v string) *KnowledgeBaseVO
	GetShowType() *string
	SetSummary(v string) *KnowledgeBaseVO
	GetSummary() *string
	SetTableName(v string) *KnowledgeBaseVO
	GetTableName() *string
}

type KnowledgeBaseVO struct {
	Category       *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Confidence     *float64 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	DbId           *int32   `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName         *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	EntityId       *int64   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Env            *string  `json:"Env,omitempty" xml:"Env,omitempty"`
	Expr           *string  `json:"Expr,omitempty" xml:"Expr,omitempty"`
	GmtCreate      *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceName   *string  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDelete       *bool    `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	KnowledgeId    *string  `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	KnowledgeType  *string  `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	LevelType      *string  `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	Name           *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OldDescription *string  `json:"OldDescription,omitempty" xml:"OldDescription,omitempty"`
	OldSummary     *string  `json:"OldSummary,omitempty" xml:"OldSummary,omitempty"`
	ParseDesc      *string  `json:"ParseDesc,omitempty" xml:"ParseDesc,omitempty"`
	ReasoningLogic *string  `json:"ReasoningLogic,omitempty" xml:"ReasoningLogic,omitempty"`
	RelationType   *string  `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	ShowType       *string  `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
	Summary        *string  `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TableName      *string  `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s KnowledgeBaseVO) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseVO) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseVO) GetCategory() *string {
	return s.Category
}

func (s *KnowledgeBaseVO) GetConfidence() *float64 {
	return s.Confidence
}

func (s *KnowledgeBaseVO) GetDbId() *int32 {
	return s.DbId
}

func (s *KnowledgeBaseVO) GetDbName() *string {
	return s.DbName
}

func (s *KnowledgeBaseVO) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBaseVO) GetEntityId() *int64 {
	return s.EntityId
}

func (s *KnowledgeBaseVO) GetEnv() *string {
	return s.Env
}

func (s *KnowledgeBaseVO) GetExpr() *string {
	return s.Expr
}

func (s *KnowledgeBaseVO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *KnowledgeBaseVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *KnowledgeBaseVO) GetIsDelete() *bool {
	return s.IsDelete
}

func (s *KnowledgeBaseVO) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *KnowledgeBaseVO) GetKnowledgeType() *string {
	return s.KnowledgeType
}

func (s *KnowledgeBaseVO) GetLevelType() *string {
	return s.LevelType
}

func (s *KnowledgeBaseVO) GetName() *string {
	return s.Name
}

func (s *KnowledgeBaseVO) GetOldDescription() *string {
	return s.OldDescription
}

func (s *KnowledgeBaseVO) GetOldSummary() *string {
	return s.OldSummary
}

func (s *KnowledgeBaseVO) GetParseDesc() *string {
	return s.ParseDesc
}

func (s *KnowledgeBaseVO) GetReasoningLogic() *string {
	return s.ReasoningLogic
}

func (s *KnowledgeBaseVO) GetRelationType() *string {
	return s.RelationType
}

func (s *KnowledgeBaseVO) GetShowType() *string {
	return s.ShowType
}

func (s *KnowledgeBaseVO) GetSummary() *string {
	return s.Summary
}

func (s *KnowledgeBaseVO) GetTableName() *string {
	return s.TableName
}

func (s *KnowledgeBaseVO) SetCategory(v string) *KnowledgeBaseVO {
	s.Category = &v
	return s
}

func (s *KnowledgeBaseVO) SetConfidence(v float64) *KnowledgeBaseVO {
	s.Confidence = &v
	return s
}

func (s *KnowledgeBaseVO) SetDbId(v int32) *KnowledgeBaseVO {
	s.DbId = &v
	return s
}

func (s *KnowledgeBaseVO) SetDbName(v string) *KnowledgeBaseVO {
	s.DbName = &v
	return s
}

func (s *KnowledgeBaseVO) SetDescription(v string) *KnowledgeBaseVO {
	s.Description = &v
	return s
}

func (s *KnowledgeBaseVO) SetEntityId(v int64) *KnowledgeBaseVO {
	s.EntityId = &v
	return s
}

func (s *KnowledgeBaseVO) SetEnv(v string) *KnowledgeBaseVO {
	s.Env = &v
	return s
}

func (s *KnowledgeBaseVO) SetExpr(v string) *KnowledgeBaseVO {
	s.Expr = &v
	return s
}

func (s *KnowledgeBaseVO) SetGmtCreate(v string) *KnowledgeBaseVO {
	s.GmtCreate = &v
	return s
}

func (s *KnowledgeBaseVO) SetInstanceName(v string) *KnowledgeBaseVO {
	s.InstanceName = &v
	return s
}

func (s *KnowledgeBaseVO) SetIsDelete(v bool) *KnowledgeBaseVO {
	s.IsDelete = &v
	return s
}

func (s *KnowledgeBaseVO) SetKnowledgeId(v string) *KnowledgeBaseVO {
	s.KnowledgeId = &v
	return s
}

func (s *KnowledgeBaseVO) SetKnowledgeType(v string) *KnowledgeBaseVO {
	s.KnowledgeType = &v
	return s
}

func (s *KnowledgeBaseVO) SetLevelType(v string) *KnowledgeBaseVO {
	s.LevelType = &v
	return s
}

func (s *KnowledgeBaseVO) SetName(v string) *KnowledgeBaseVO {
	s.Name = &v
	return s
}

func (s *KnowledgeBaseVO) SetOldDescription(v string) *KnowledgeBaseVO {
	s.OldDescription = &v
	return s
}

func (s *KnowledgeBaseVO) SetOldSummary(v string) *KnowledgeBaseVO {
	s.OldSummary = &v
	return s
}

func (s *KnowledgeBaseVO) SetParseDesc(v string) *KnowledgeBaseVO {
	s.ParseDesc = &v
	return s
}

func (s *KnowledgeBaseVO) SetReasoningLogic(v string) *KnowledgeBaseVO {
	s.ReasoningLogic = &v
	return s
}

func (s *KnowledgeBaseVO) SetRelationType(v string) *KnowledgeBaseVO {
	s.RelationType = &v
	return s
}

func (s *KnowledgeBaseVO) SetShowType(v string) *KnowledgeBaseVO {
	s.ShowType = &v
	return s
}

func (s *KnowledgeBaseVO) SetSummary(v string) *KnowledgeBaseVO {
	s.Summary = &v
	return s
}

func (s *KnowledgeBaseVO) SetTableName(v string) *KnowledgeBaseVO {
	s.TableName = &v
	return s
}

func (s *KnowledgeBaseVO) Validate() error {
	return dara.Validate(s)
}
