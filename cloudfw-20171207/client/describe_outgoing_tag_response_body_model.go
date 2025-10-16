// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOutgoingTagResponseBody
	GetRequestId() *string
	SetTagList(v []*DescribeOutgoingTagResponseBodyTagList) *DescribeOutgoingTagResponseBody
	GetTagList() []*DescribeOutgoingTagResponseBodyTagList
	SetTotalCount(v int32) *DescribeOutgoingTagResponseBody
	GetTotalCount() *int32
}

type DescribeOutgoingTagResponseBody struct {
	// example:
	//
	// B532203E-813B-5BEB-B75B-315E1D08****
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagList   []*DescribeOutgoingTagResponseBodyTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingTagResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingTagResponseBody) GetTagList() []*DescribeOutgoingTagResponseBodyTagList {
	return s.TagList
}

func (s *DescribeOutgoingTagResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingTagResponseBody) SetRequestId(v string) *DescribeOutgoingTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingTagResponseBody) SetTagList(v []*DescribeOutgoingTagResponseBodyTagList) *DescribeOutgoingTagResponseBody {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingTagResponseBody) SetTotalCount(v int32) *DescribeOutgoingTagResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingTagResponseBody) Validate() error {
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingTagResponseBodyTagList struct {
	// example:
	//
	// Trusted
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// example:
	//
	// 0
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test describe
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// example:
	//
	// tag-6833388d18cc****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// test tag
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingTagResponseBodyTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingTagResponseBodyTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingTagResponseBodyTagList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingTagResponseBodyTagList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeOutgoingTagResponseBodyTagList) GetTagDescribe() *string {
	return s.TagDescribe
}

func (s *DescribeOutgoingTagResponseBodyTagList) GetTagId() *string {
	return s.TagId
}

func (s *DescribeOutgoingTagResponseBodyTagList) GetTagName() *string {
	return s.TagName
}

func (s *DescribeOutgoingTagResponseBodyTagList) SetClassId(v string) *DescribeOutgoingTagResponseBodyTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingTagResponseBodyTagList) SetRiskLevel(v int32) *DescribeOutgoingTagResponseBodyTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingTagResponseBodyTagList) SetTagDescribe(v string) *DescribeOutgoingTagResponseBodyTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingTagResponseBodyTagList) SetTagId(v string) *DescribeOutgoingTagResponseBodyTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingTagResponseBodyTagList) SetTagName(v string) *DescribeOutgoingTagResponseBodyTagList {
	s.TagName = &v
	return s
}

func (s *DescribeOutgoingTagResponseBodyTagList) Validate() error {
	return dara.Validate(s)
}
