// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListFlowsRequestFilter) *ListFlowsRequest
	GetFilter() []*ListFlowsRequestFilter
	SetMaxResults(v string) *ListFlowsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListFlowsRequest
	GetNextToken() *string
	SetTag(v []*ListFlowsRequestTag) *ListFlowsRequest
	GetTag() []*ListFlowsRequestTag
}

type ListFlowsRequest struct {
	// The filter conditions.
	Filter []*ListFlowsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// AAAAAVY3rYiv9VoUJQSiCitgjgTlo5Q2D1qyq9rPPhxWJPl6j8bgHiGAwZWnCMJPepC+WRjTa7fpTA0Diy2AQ4aWoPY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies the object tags to which the rule applies. You can specify multiple tags.
	Tag []*ListFlowsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) GetFilter() []*ListFlowsRequestFilter {
	return s.Filter
}

func (s *ListFlowsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListFlowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowsRequest) GetTag() []*ListFlowsRequestTag {
	return s.Tag
}

func (s *ListFlowsRequest) SetFilter(v []*ListFlowsRequestFilter) *ListFlowsRequest {
	s.Filter = v
	return s
}

func (s *ListFlowsRequest) SetMaxResults(v string) *ListFlowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFlowsRequest) SetNextToken(v string) *ListFlowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFlowsRequest) SetTag(v []*ListFlowsRequestTag) *ListFlowsRequest {
	s.Tag = v
	return s
}

func (s *ListFlowsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowsRequestFilter struct {
	// The name of the filter field.
	//
	// example:
	//
	// FlowId\\FlowName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListFlowsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListFlowsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListFlowsRequestFilter) GetValues() []*string {
	return s.Values
}

func (s *ListFlowsRequestFilter) SetName(v string) *ListFlowsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListFlowsRequestFilter) SetValues(v []*string) *ListFlowsRequestFilter {
	s.Values = v
	return s
}

func (s *ListFlowsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListFlowsRequestTag struct {
	// The tag key. You can filter the cluster list by tag. You can specify up to 20 tag pairs. The numeric value n for each tag pair must be unique and must be a consecutive integer starting from 1. The value of Tag.N.Key corresponds to Tag.N.Value.
	//
	// > The tag key can be up to 64 characters in length and cannot start with `aliyun`, `acs:`, `http://`, or `https://`.
	//
	// example:
	//
	// CreateBy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The authentication content.
	//
	// example:
	//
	// zhangsan
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFlowsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsRequestTag) GoString() string {
	return s.String()
}

func (s *ListFlowsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListFlowsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListFlowsRequestTag) SetKey(v string) *ListFlowsRequestTag {
	s.Key = &v
	return s
}

func (s *ListFlowsRequestTag) SetValue(v string) *ListFlowsRequestTag {
	s.Value = &v
	return s
}

func (s *ListFlowsRequestTag) Validate() error {
	return dara.Validate(s)
}
