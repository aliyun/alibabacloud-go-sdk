// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryHotlineNumberRequest
	GetCurrentPage() *int32
	SetDepartmentId(v int64) *QueryHotlineNumberRequest
	GetDepartmentId() *int64
	SetGroupIds(v []*int64) *QueryHotlineNumberRequest
	GetGroupIds() []*int64
	SetHotlineNumber(v string) *QueryHotlineNumberRequest
	GetHotlineNumber() *string
	SetInstanceId(v string) *QueryHotlineNumberRequest
	GetInstanceId() *string
	SetPageSize(v int32) *QueryHotlineNumberRequest
	GetPageSize() *int32
}

type QueryHotlineNumberRequest struct {
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
	GroupIds []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
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

func (s QueryHotlineNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineNumberRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineNumberRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryHotlineNumberRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *QueryHotlineNumberRequest) GetGroupIds() []*int64 {
	return s.GroupIds
}

func (s *QueryHotlineNumberRequest) GetHotlineNumber() *string {
	return s.HotlineNumber
}

func (s *QueryHotlineNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryHotlineNumberRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryHotlineNumberRequest) SetCurrentPage(v int32) *QueryHotlineNumberRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetDepartmentId(v int64) *QueryHotlineNumberRequest {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetGroupIds(v []*int64) *QueryHotlineNumberRequest {
	s.GroupIds = v
	return s
}

func (s *QueryHotlineNumberRequest) SetHotlineNumber(v string) *QueryHotlineNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetInstanceId(v string) *QueryHotlineNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineNumberRequest) SetPageSize(v int32) *QueryHotlineNumberRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineNumberRequest) Validate() error {
	return dara.Validate(s)
}
