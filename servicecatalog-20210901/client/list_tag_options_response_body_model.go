// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagOptionsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListTagOptionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagOptionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagOptionsResponseBody
	GetRequestId() *string
	SetTagOptionDetails(v []*ListTagOptionsResponseBodyTagOptionDetails) *ListTagOptionsResponseBody
	GetTagOptionDetails() []*ListTagOptionsResponseBodyTagOptionDetails
	SetTotalCount(v int32) *ListTagOptionsResponseBody
	GetTotalCount() *int32
}

type ListTagOptionsResponseBody struct {
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 37C9C1DF-EFFC-5D8A-80D0-8657B1F3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the tag option.
	TagOptionDetails []*ListTagOptionsResponseBodyTagOptionDetails `json:"TagOptionDetails,omitempty" xml:"TagOptionDetails,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagOptionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagOptionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagOptionsResponseBody) GetTagOptionDetails() []*ListTagOptionsResponseBodyTagOptionDetails {
	return s.TagOptionDetails
}

func (s *ListTagOptionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagOptionsResponseBody) SetNextToken(v string) *ListTagOptionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetPageNumber(v int32) *ListTagOptionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetPageSize(v int32) *ListTagOptionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetRequestId(v string) *ListTagOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetTagOptionDetails(v []*ListTagOptionsResponseBodyTagOptionDetails) *ListTagOptionsResponseBody {
	s.TagOptionDetails = v
	return s
}

func (s *ListTagOptionsResponseBody) SetTotalCount(v int32) *ListTagOptionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagOptionsResponseBody) Validate() error {
	if s.TagOptionDetails != nil {
		for _, item := range s.TagOptionDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagOptionsResponseBodyTagOptionDetails struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	//
	// example:
	//
	// 133413081827****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	//
	// example:
	//
	// tag-bp1r3mxq3t****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagOptionsResponseBodyTagOptionDetails) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsResponseBodyTagOptionDetails) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) GetActive() *bool {
	return s.Active
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) GetKey() *string {
	return s.Key
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) GetOwner() *string {
	return s.Owner
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) GetValue() *string {
	return s.Value
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetActive(v bool) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Active = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetKey(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Key = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetOwner(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Owner = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetTagOptionId(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.TagOptionId = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetValue(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Value = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) Validate() error {
	return dara.Validate(s)
}
