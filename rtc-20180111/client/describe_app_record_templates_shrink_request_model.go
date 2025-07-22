// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppRecordTemplatesShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *DescribeAppRecordTemplatesShrinkRequest
	GetClientToken() *string
	SetConditionShrink(v string) *DescribeAppRecordTemplatesShrinkRequest
	GetConditionShrink() *string
	SetPageNum(v int32) *DescribeAppRecordTemplatesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppRecordTemplatesShrinkRequest
	GetPageSize() *int32
}

type DescribeAppRecordTemplatesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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

func (s DescribeAppRecordTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppRecordTemplatesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAppRecordTemplatesShrinkRequest) GetConditionShrink() *string {
	return s.ConditionShrink
}

func (s *DescribeAppRecordTemplatesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppRecordTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppRecordTemplatesShrinkRequest) SetAppId(v string) *DescribeAppRecordTemplatesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRecordTemplatesShrinkRequest) SetClientToken(v string) *DescribeAppRecordTemplatesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAppRecordTemplatesShrinkRequest) SetConditionShrink(v string) *DescribeAppRecordTemplatesShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeAppRecordTemplatesShrinkRequest) SetPageNum(v int32) *DescribeAppRecordTemplatesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppRecordTemplatesShrinkRequest) SetPageSize(v int32) *DescribeAppRecordTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppRecordTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
