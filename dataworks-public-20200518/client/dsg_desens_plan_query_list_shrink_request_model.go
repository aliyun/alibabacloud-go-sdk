// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanQueryListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *DsgDesensPlanQueryListShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *DsgDesensPlanQueryListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgDesensPlanQueryListShrinkRequest
	GetPageSize() *int32
	SetRuleName(v string) *DsgDesensPlanQueryListShrinkRequest
	GetRuleName() *string
	SetSceneId(v int64) *DsgDesensPlanQueryListShrinkRequest
	GetSceneId() *int64
	SetStatus(v int32) *DsgDesensPlanQueryListShrinkRequest
	GetStatus() *int32
	SetColumnsShrink(v string) *DsgDesensPlanQueryListShrinkRequest
	GetColumnsShrink() *string
	SetDataType(v string) *DsgDesensPlanQueryListShrinkRequest
	GetDataType() *string
	SetEmptyNotDesesn(v string) *DsgDesensPlanQueryListShrinkRequest
	GetEmptyNotDesesn() *string
}

type DsgDesensPlanQueryListShrinkRequest struct {
	// The owner of the data masking rule.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the sensitive field.
	//
	// example:
	//
	// phone
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the level-2 data masking scenario. You can call the [DsgSceneQuerySceneListByName](https://help.aliyun.com/document_detail/2786322.html) operation to query the list of IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// example:
	//
	// 1
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ColumnsShrink  *string `json:"columns,omitempty" xml:"columns,omitempty"`
	DataType       *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	EmptyNotDesesn *string `json:"emptyNotDesesn,omitempty" xml:"emptyNotDesesn,omitempty"`
}

func (s DsgDesensPlanQueryListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetDataType() *string {
	return s.DataType
}

func (s *DsgDesensPlanQueryListShrinkRequest) GetEmptyNotDesesn() *string {
	return s.EmptyNotDesesn
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetOwner(v string) *DsgDesensPlanQueryListShrinkRequest {
	s.Owner = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetPageNumber(v int32) *DsgDesensPlanQueryListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetPageSize(v int32) *DsgDesensPlanQueryListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetRuleName(v string) *DsgDesensPlanQueryListShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetSceneId(v int64) *DsgDesensPlanQueryListShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetStatus(v int32) *DsgDesensPlanQueryListShrinkRequest {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetColumnsShrink(v string) *DsgDesensPlanQueryListShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetDataType(v string) *DsgDesensPlanQueryListShrinkRequest {
	s.DataType = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) SetEmptyNotDesesn(v string) *DsgDesensPlanQueryListShrinkRequest {
	s.EmptyNotDesesn = &v
	return s
}

func (s *DsgDesensPlanQueryListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
