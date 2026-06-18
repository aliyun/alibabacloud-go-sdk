// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryHotlineNumberShrinkRequest
	GetCurrentPage() *int32
	SetDepartmentId(v int64) *QueryHotlineNumberShrinkRequest
	GetDepartmentId() *int64
	SetGroupIdsShrink(v string) *QueryHotlineNumberShrinkRequest
	GetGroupIdsShrink() *string
	SetHotlineNumber(v string) *QueryHotlineNumberShrinkRequest
	GetHotlineNumber() *string
	SetInstanceId(v string) *QueryHotlineNumberShrinkRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryHotlineNumberShrinkRequest
	GetPageSize() *int32
}

type QueryHotlineNumberShrinkRequest struct {
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The department ID.
	//
	// example:
	//
	// 2256****
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// The list of skill groups.
	GroupIdsShrink *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// The hotline number. Fuzzy query is supported.
	//
	// example:
	//
	// 0571****2211
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryHotlineNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryHotlineNumberShrinkRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QueryHotlineNumberShrinkRequest) GetGroupIdsShrink() *string {
	return s.GroupIdsShrink
}

func (s *QueryHotlineNumberShrinkRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *QueryHotlineNumberShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryHotlineNumberShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryHotlineNumberShrinkRequest) SetCurrentPage(v int32) *QueryHotlineNumberShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetDepartmentId(v int64) *QueryHotlineNumberShrinkRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetGroupIdsShrink(v string) *QueryHotlineNumberShrinkRequest {
	s.GroupIdsShrink = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetHotlineNumber(v string) *QueryHotlineNumberShrinkRequest {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetInstanceId(v string) *QueryHotlineNumberShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) SetPageSize(v int32) *QueryHotlineNumberShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
