// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescribeNamespacesRequest(v *DescribeNamespacesRequestDescribeNamespacesRequest) *DescribeNamespacesRequest
	GetDescribeNamespacesRequest() *DescribeNamespacesRequestDescribeNamespacesRequest
}

type DescribeNamespacesRequest struct {
	// This parameter is required.
	DescribeNamespacesRequest *DescribeNamespacesRequestDescribeNamespacesRequest `json:"DescribeNamespacesRequest,omitempty" xml:"DescribeNamespacesRequest,omitempty" type:"Struct"`
}

func (s DescribeNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) GetDescribeNamespacesRequest() *DescribeNamespacesRequestDescribeNamespacesRequest {
	return s.DescribeNamespacesRequest
}

func (s *DescribeNamespacesRequest) SetDescribeNamespacesRequest(v *DescribeNamespacesRequestDescribeNamespacesRequest) *DescribeNamespacesRequest {
	s.DescribeNamespacesRequest = v
	return s
}

func (s *DescribeNamespacesRequest) Validate() error {
	if s.DescribeNamespacesRequest != nil {
		if err := s.DescribeNamespacesRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNamespacesRequestDescribeNamespacesRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 223493C7-FCA9-13D4-B75B-AF8B32F4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ns-1
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
	Region *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	Tags   []*DescribeNamespacesRequestDescribeNamespacesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequestDescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetHa() *bool {
	return s.Ha
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) GetTags() []*DescribeNamespacesRequestDescribeNamespacesRequestTags {
	return s.Tags
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetHa(v bool) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetInstanceId(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetNamespace(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetPageIndex(v int32) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetRegion(v string) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Region = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) SetTags(v []*DescribeNamespacesRequestDescribeNamespacesRequestTags) *DescribeNamespacesRequestDescribeNamespacesRequest {
	s.Tags = v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequest) Validate() error {
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

type DescribeNamespacesRequestDescribeNamespacesRequestTags struct {
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesRequestDescribeNamespacesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequestDescribeNamespacesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) SetKey(v string) *DescribeNamespacesRequestDescribeNamespacesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) SetValue(v string) *DescribeNamespacesRequestDescribeNamespacesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeNamespacesRequestDescribeNamespacesRequestTags) Validate() error {
	return dara.Validate(s)
}
