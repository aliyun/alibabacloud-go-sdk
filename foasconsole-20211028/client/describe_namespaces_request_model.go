// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *DescribeNamespacesRequest
	GetHa() *bool
	SetInstanceId(v string) *DescribeNamespacesRequest
	GetInstanceId() *string
	SetNamespace(v string) *DescribeNamespacesRequest
	GetNamespace() *string
	SetPageIndex(v int32) *DescribeNamespacesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeNamespacesRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeNamespacesRequest
	GetRegion() *string
	SetTags(v []*DescribeNamespacesRequestTags) *DescribeNamespacesRequest
	GetTags() []*DescribeNamespacesRequestTags
}

type DescribeNamespacesRequest struct {
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
	Region *string                          `json:"Region,omitempty" xml:"Region,omitempty"`
	Tags   []*DescribeNamespacesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) GetHa() *bool {
	return s.Ha
}

func (s *DescribeNamespacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNamespacesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespacesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeNamespacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeNamespacesRequest) GetTags() []*DescribeNamespacesRequestTags {
	return s.Tags
}

func (s *DescribeNamespacesRequest) SetHa(v bool) *DescribeNamespacesRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesRequest) SetInstanceId(v string) *DescribeNamespacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesRequest) SetNamespace(v string) *DescribeNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageIndex(v int32) *DescribeNamespacesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesRequest) SetRegion(v string) *DescribeNamespacesRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesRequest) SetTags(v []*DescribeNamespacesRequestTags) *DescribeNamespacesRequest {
	s.Tags = v
	return s
}

func (s *DescribeNamespacesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNamespacesRequestTags struct {
	// example:
	//
	// FLink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNamespacesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNamespacesRequestTags) SetKey(v string) *DescribeNamespacesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesRequestTags) SetValue(v string) *DescribeNamespacesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeNamespacesRequestTags) Validate() error {
	return dara.Validate(s)
}
