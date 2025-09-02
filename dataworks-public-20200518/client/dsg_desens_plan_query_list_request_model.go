// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgDesensPlanQueryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *DsgDesensPlanQueryListRequest
	GetOwner() *string
	SetPageNumber(v int32) *DsgDesensPlanQueryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgDesensPlanQueryListRequest
	GetPageSize() *int32
	SetRuleName(v string) *DsgDesensPlanQueryListRequest
	GetRuleName() *string
	SetSceneId(v int32) *DsgDesensPlanQueryListRequest
	GetSceneId() *int32
	SetStatus(v int32) *DsgDesensPlanQueryListRequest
	GetStatus() *int32
}

type DsgDesensPlanQueryListRequest struct {
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
	SceneId *int32 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the data masking rule. Valid values:
	//
	// 	- 0: expired
	//
	// 	- 1: effective
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DsgDesensPlanQueryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgDesensPlanQueryListRequest) GoString() string {
	return s.String()
}

func (s *DsgDesensPlanQueryListRequest) GetOwner() *string {
	return s.Owner
}

func (s *DsgDesensPlanQueryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgDesensPlanQueryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgDesensPlanQueryListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgDesensPlanQueryListRequest) GetSceneId() *int32 {
	return s.SceneId
}

func (s *DsgDesensPlanQueryListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DsgDesensPlanQueryListRequest) SetOwner(v string) *DsgDesensPlanQueryListRequest {
	s.Owner = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetPageNumber(v int32) *DsgDesensPlanQueryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetPageSize(v int32) *DsgDesensPlanQueryListRequest {
	s.PageSize = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetRuleName(v string) *DsgDesensPlanQueryListRequest {
	s.RuleName = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetSceneId(v int32) *DsgDesensPlanQueryListRequest {
	s.SceneId = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) SetStatus(v int32) *DsgDesensPlanQueryListRequest {
	s.Status = &v
	return s
}

func (s *DsgDesensPlanQueryListRequest) Validate() error {
	return dara.Validate(s)
}
