// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppViewTemplatesShrinkRequest
	GetAppId() *string
	SetConditionShrink(v string) *DescribeAppViewTemplatesShrinkRequest
	GetConditionShrink() *string
	SetPageNum(v int32) *DescribeAppViewTemplatesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppViewTemplatesShrinkRequest
	GetPageSize() *int32
}

type DescribeAppViewTemplatesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppViewTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppViewTemplatesShrinkRequest) GetConditionShrink() *string {
	return s.ConditionShrink
}

func (s *DescribeAppViewTemplatesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppViewTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppViewTemplatesShrinkRequest) SetAppId(v string) *DescribeAppViewTemplatesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppViewTemplatesShrinkRequest) SetConditionShrink(v string) *DescribeAppViewTemplatesShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeAppViewTemplatesShrinkRequest) SetPageNum(v int32) *DescribeAppViewTemplatesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppViewTemplatesShrinkRequest) SetPageSize(v int32) *DescribeAppViewTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppViewTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
