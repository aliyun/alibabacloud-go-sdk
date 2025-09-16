// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *DescribeNamespacesShrinkRequest
	GetHa() *bool
	SetInstanceId(v string) *DescribeNamespacesShrinkRequest
	GetInstanceId() *string
	SetNamespace(v string) *DescribeNamespacesShrinkRequest
	GetNamespace() *string
	SetPageIndex(v int32) *DescribeNamespacesShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeNamespacesShrinkRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeNamespacesShrinkRequest
	GetRegion() *string
	SetTagsShrink(v string) *DescribeNamespacesShrinkRequest
	GetTagsShrink() *string
}

type DescribeNamespacesShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// di-590843445844225
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeNamespacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *DescribeNamespacesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNamespacesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespacesShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeNamespacesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeNamespacesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *DescribeNamespacesShrinkRequest) SetHa(v bool) *DescribeNamespacesShrinkRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetInstanceId(v string) *DescribeNamespacesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetNamespace(v string) *DescribeNamespacesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetPageIndex(v int32) *DescribeNamespacesShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetPageSize(v int32) *DescribeNamespacesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetRegion(v string) *DescribeNamespacesShrinkRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) SetTagsShrink(v string) *DescribeNamespacesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *DescribeNamespacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
