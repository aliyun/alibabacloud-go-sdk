// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStreamingOutTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppStreamingOutTemplatesShrinkRequest
	GetAppId() *string
	SetConditionShrink(v string) *DescribeAppStreamingOutTemplatesShrinkRequest
	GetConditionShrink() *string
	SetPageNum(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest
	GetPageSize() *int32
}

type DescribeAppStreamingOutTemplatesShrinkRequest struct {
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

func (s DescribeAppStreamingOutTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) GetConditionShrink() *string {
	return s.ConditionShrink
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetAppId(v string) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetConditionShrink(v string) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetPageNum(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetPageSize(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
