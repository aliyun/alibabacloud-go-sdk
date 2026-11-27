// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLayoutsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppLayoutsShrinkRequest
	GetAppId() *string
	SetConditionShrink(v string) *DescribeAppLayoutsShrinkRequest
	GetConditionShrink() *string
	SetPageNum(v int32) *DescribeAppLayoutsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppLayoutsShrinkRequest
	GetPageSize() *int32
}

type DescribeAppLayoutsShrinkRequest struct {
	// The application ID. Create and query it using the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The query conditions.
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppLayoutsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppLayoutsShrinkRequest) GetConditionShrink() *string {
	return s.ConditionShrink
}

func (s *DescribeAppLayoutsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppLayoutsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppLayoutsShrinkRequest) SetAppId(v string) *DescribeAppLayoutsShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppLayoutsShrinkRequest) SetConditionShrink(v string) *DescribeAppLayoutsShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeAppLayoutsShrinkRequest) SetPageNum(v int32) *DescribeAppLayoutsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppLayoutsShrinkRequest) SetPageSize(v int32) *DescribeAppLayoutsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppLayoutsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
