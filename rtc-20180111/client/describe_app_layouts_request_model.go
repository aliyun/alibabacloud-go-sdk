// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLayoutsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppLayoutsRequest
	GetAppId() *string
	SetCondition(v *DescribeAppLayoutsRequestCondition) *DescribeAppLayoutsRequest
	GetCondition() *DescribeAppLayoutsRequestCondition
	SetPageNum(v int32) *DescribeAppLayoutsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppLayoutsRequest
	GetPageSize() *int32
}

type DescribeAppLayoutsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId     *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Condition *DescribeAppLayoutsRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppLayoutsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppLayoutsRequest) GetCondition() *DescribeAppLayoutsRequestCondition {
	return s.Condition
}

func (s *DescribeAppLayoutsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppLayoutsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppLayoutsRequest) SetAppId(v string) *DescribeAppLayoutsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppLayoutsRequest) SetCondition(v *DescribeAppLayoutsRequestCondition) *DescribeAppLayoutsRequest {
	s.Condition = v
	return s
}

func (s *DescribeAppLayoutsRequest) SetPageNum(v int32) *DescribeAppLayoutsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppLayoutsRequest) SetPageSize(v int32) *DescribeAppLayoutsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppLayoutsRequest) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppLayoutsRequestCondition struct {
	// example:
	//
	// 167466539798442****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAppLayoutsRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsRequestCondition) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeAppLayoutsRequestCondition) GetName() *string {
	return s.Name
}

func (s *DescribeAppLayoutsRequestCondition) SetLayoutId(v string) *DescribeAppLayoutsRequestCondition {
	s.LayoutId = &v
	return s
}

func (s *DescribeAppLayoutsRequestCondition) SetName(v string) *DescribeAppLayoutsRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeAppLayoutsRequestCondition) Validate() error {
	return dara.Validate(s)
}
