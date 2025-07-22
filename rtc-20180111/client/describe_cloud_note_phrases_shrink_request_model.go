// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotePhrasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCloudNotePhrasesShrinkRequest
	GetAppId() *string
	SetConditionShrink(v string) *DescribeCloudNotePhrasesShrinkRequest
	GetConditionShrink() *string
	SetPageNum(v int32) *DescribeCloudNotePhrasesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeCloudNotePhrasesShrinkRequest
	GetPageSize() *int32
}

type DescribeCloudNotePhrasesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
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

func (s DescribeCloudNotePhrasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudNotePhrasesShrinkRequest) GetConditionShrink() *string {
	return s.ConditionShrink
}

func (s *DescribeCloudNotePhrasesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCloudNotePhrasesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudNotePhrasesShrinkRequest) SetAppId(v string) *DescribeCloudNotePhrasesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCloudNotePhrasesShrinkRequest) SetConditionShrink(v string) *DescribeCloudNotePhrasesShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeCloudNotePhrasesShrinkRequest) SetPageNum(v int32) *DescribeCloudNotePhrasesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCloudNotePhrasesShrinkRequest) SetPageSize(v int32) *DescribeCloudNotePhrasesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudNotePhrasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
