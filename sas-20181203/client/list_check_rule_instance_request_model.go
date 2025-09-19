// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRuleInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCheckRuleInstanceRequest
	GetCurrentPage() *int32
	SetInstanceList(v []*ListCheckRuleInstanceRequestInstanceList) *ListCheckRuleInstanceRequest
	GetInstanceList() []*ListCheckRuleInstanceRequestInstanceList
	SetPageSize(v int32) *ListCheckRuleInstanceRequest
	GetPageSize() *int32
	SetRuleId(v int64) *ListCheckRuleInstanceRequest
	GetRuleId() *int64
}

type ListCheckRuleInstanceRequest struct {
	// The page number of the current page when performing a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Instance list.
	InstanceList []*ListCheckRuleInstanceRequestInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The maximum number of items per page in a paginated query. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Rule ID.
	//
	// > You can call the [LisCheckRule](https://help.aliyun.com/document_detail/2590599.html) interface to get this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListCheckRuleInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRuleInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListCheckRuleInstanceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckRuleInstanceRequest) GetInstanceList() []*ListCheckRuleInstanceRequestInstanceList {
	return s.InstanceList
}

func (s *ListCheckRuleInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckRuleInstanceRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListCheckRuleInstanceRequest) SetCurrentPage(v int32) *ListCheckRuleInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckRuleInstanceRequest) SetInstanceList(v []*ListCheckRuleInstanceRequestInstanceList) *ListCheckRuleInstanceRequest {
	s.InstanceList = v
	return s
}

func (s *ListCheckRuleInstanceRequest) SetPageSize(v int32) *ListCheckRuleInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckRuleInstanceRequest) SetRuleId(v int64) *ListCheckRuleInstanceRequest {
	s.RuleId = &v
	return s
}

func (s *ListCheckRuleInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type ListCheckRuleInstanceRequestInstanceList struct {
	// Asset instance ID.
	//
	// > Call the [ListCheckInstanceResult](~~ListCheckInstanceResult~~) interface to get this parameter.
	//
	// example:
	//
	// i-wz9fdluqx20mp2x7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the asset is located.
	//
	// > Call the [ListCheckInstanceResult](~~ListCheckInstanceResult~~) interface to get this parameter.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCheckRuleInstanceRequestInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRuleInstanceRequestInstanceList) GoString() string {
	return s.String()
}

func (s *ListCheckRuleInstanceRequestInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCheckRuleInstanceRequestInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCheckRuleInstanceRequestInstanceList) SetInstanceId(v string) *ListCheckRuleInstanceRequestInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListCheckRuleInstanceRequestInstanceList) SetRegionId(v string) *ListCheckRuleInstanceRequestInstanceList {
	s.RegionId = &v
	return s
}

func (s *ListCheckRuleInstanceRequestInstanceList) Validate() error {
	return dara.Validate(s)
}
