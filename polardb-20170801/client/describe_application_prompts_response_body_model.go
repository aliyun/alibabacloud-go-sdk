// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPromptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeApplicationPromptsResponseBodyItems) *DescribeApplicationPromptsResponseBody
	GetItems() []*DescribeApplicationPromptsResponseBodyItems
	SetPageNumber(v string) *DescribeApplicationPromptsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeApplicationPromptsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeApplicationPromptsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeApplicationPromptsResponseBody
	GetTotalRecordCount() *string
}

type DescribeApplicationPromptsResponseBody struct {
	Items []*DescribeApplicationPromptsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationPromptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPromptsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPromptsResponseBody) GetItems() []*DescribeApplicationPromptsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationPromptsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeApplicationPromptsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeApplicationPromptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationPromptsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeApplicationPromptsResponseBody) SetItems(v []*DescribeApplicationPromptsResponseBodyItems) *DescribeApplicationPromptsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationPromptsResponseBody) SetPageNumber(v string) *DescribeApplicationPromptsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBody) SetPageRecordCount(v string) *DescribeApplicationPromptsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBody) SetRequestId(v string) *DescribeApplicationPromptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBody) SetTotalRecordCount(v string) *DescribeApplicationPromptsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationPromptsResponseBodyItems struct {
	// example:
	//
	// prompt if enabled
	PromptEnabled *int32 `json:"PromptEnabled,omitempty" xml:"PromptEnabled,omitempty"`
	// example:
	//
	// papt-58z96zl691otf356o4
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// example:
	//
	// prompt name
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
	// example:
	//
	// DELETE_RELATIONS_SYSTEM_PROMPT
	PromptType *string `json:"PromptType,omitempty" xml:"PromptType,omitempty"`
	// example:
	//
	// prompt value
	PromptValue *string `json:"PromptValue,omitempty" xml:"PromptValue,omitempty"`
}

func (s DescribeApplicationPromptsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPromptsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPromptsResponseBodyItems) GetPromptEnabled() *int32 {
	return s.PromptEnabled
}

func (s *DescribeApplicationPromptsResponseBodyItems) GetPromptId() *string {
	return s.PromptId
}

func (s *DescribeApplicationPromptsResponseBodyItems) GetPromptName() *string {
	return s.PromptName
}

func (s *DescribeApplicationPromptsResponseBodyItems) GetPromptType() *string {
	return s.PromptType
}

func (s *DescribeApplicationPromptsResponseBodyItems) GetPromptValue() *string {
	return s.PromptValue
}

func (s *DescribeApplicationPromptsResponseBodyItems) SetPromptEnabled(v int32) *DescribeApplicationPromptsResponseBodyItems {
	s.PromptEnabled = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBodyItems) SetPromptId(v string) *DescribeApplicationPromptsResponseBodyItems {
	s.PromptId = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBodyItems) SetPromptName(v string) *DescribeApplicationPromptsResponseBodyItems {
	s.PromptName = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBodyItems) SetPromptType(v string) *DescribeApplicationPromptsResponseBodyItems {
	s.PromptType = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBodyItems) SetPromptValue(v string) *DescribeApplicationPromptsResponseBodyItems {
	s.PromptValue = &v
	return s
}

func (s *DescribeApplicationPromptsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
