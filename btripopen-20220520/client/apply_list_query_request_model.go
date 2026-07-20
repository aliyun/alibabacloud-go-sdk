// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyListQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllApply(v bool) *ApplyListQueryRequest
	GetAllApply() *bool
	SetDepartId(v string) *ApplyListQueryRequest
	GetDepartId() *string
	SetEndTime(v string) *ApplyListQueryRequest
	GetEndTime() *string
	SetGmtModified(v string) *ApplyListQueryRequest
	GetGmtModified() *string
	SetOnlyShangLvApply(v bool) *ApplyListQueryRequest
	GetOnlyShangLvApply() *bool
	SetPage(v int32) *ApplyListQueryRequest
	GetPage() *int32
	SetPageSize(v int32) *ApplyListQueryRequest
	GetPageSize() *int32
	SetStartTime(v string) *ApplyListQueryRequest
	GetStartTime() *string
	SetSubCorpId(v string) *ApplyListQueryRequest
	GetSubCorpId() *string
	SetType(v int32) *ApplyListQueryRequest
	GetType() *int32
	SetUnionNo(v string) *ApplyListQueryRequest
	GetUnionNo() *string
	SetUserId(v string) *ApplyListQueryRequest
	GetUserId() *string
}

type ApplyListQueryRequest struct {
	// example:
	//
	// true
	AllApply *bool `json:"all_apply,omitempty" xml:"all_apply,omitempty"`
	// example:
	//
	// dept1
	DepartId *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 2017-05-01 00:00:00
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 2018-05-01 00:00:00
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// false
	OnlyShangLvApply *bool `json:"only_shang_lv_apply,omitempty" xml:"only_shang_lv_apply,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2017-05-01 00:00:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// btrip123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// abs123
	UnionNo *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// example:
	//
	// user1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyListQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryRequest) GoString() string {
	return s.String()
}

func (s *ApplyListQueryRequest) GetAllApply() *bool {
	return s.AllApply
}

func (s *ApplyListQueryRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyListQueryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ApplyListQueryRequest) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ApplyListQueryRequest) GetOnlyShangLvApply() *bool {
	return s.OnlyShangLvApply
}

func (s *ApplyListQueryRequest) GetPage() *int32 {
	return s.Page
}

func (s *ApplyListQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ApplyListQueryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ApplyListQueryRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyListQueryRequest) GetType() *int32 {
	return s.Type
}

func (s *ApplyListQueryRequest) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyListQueryRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyListQueryRequest) SetAllApply(v bool) *ApplyListQueryRequest {
	s.AllApply = &v
	return s
}

func (s *ApplyListQueryRequest) SetDepartId(v string) *ApplyListQueryRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryRequest) SetEndTime(v string) *ApplyListQueryRequest {
	s.EndTime = &v
	return s
}

func (s *ApplyListQueryRequest) SetGmtModified(v string) *ApplyListQueryRequest {
	s.GmtModified = &v
	return s
}

func (s *ApplyListQueryRequest) SetOnlyShangLvApply(v bool) *ApplyListQueryRequest {
	s.OnlyShangLvApply = &v
	return s
}

func (s *ApplyListQueryRequest) SetPage(v int32) *ApplyListQueryRequest {
	s.Page = &v
	return s
}

func (s *ApplyListQueryRequest) SetPageSize(v int32) *ApplyListQueryRequest {
	s.PageSize = &v
	return s
}

func (s *ApplyListQueryRequest) SetStartTime(v string) *ApplyListQueryRequest {
	s.StartTime = &v
	return s
}

func (s *ApplyListQueryRequest) SetSubCorpId(v string) *ApplyListQueryRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyListQueryRequest) SetType(v int32) *ApplyListQueryRequest {
	s.Type = &v
	return s
}

func (s *ApplyListQueryRequest) SetUnionNo(v string) *ApplyListQueryRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyListQueryRequest) SetUserId(v string) *ApplyListQueryRequest {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryRequest) Validate() error {
	return dara.Validate(s)
}
