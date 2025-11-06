// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppViewTemplatesRequest
	GetAppId() *string
	SetCondition(v *DescribeAppViewTemplatesRequestCondition) *DescribeAppViewTemplatesRequest
	GetCondition() *DescribeAppViewTemplatesRequestCondition
	SetPageNum(v int32) *DescribeAppViewTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppViewTemplatesRequest
	GetPageSize() *int32
}

type DescribeAppViewTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId     *string                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Condition *DescribeAppViewTemplatesRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppViewTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppViewTemplatesRequest) GetCondition() *DescribeAppViewTemplatesRequestCondition {
	return s.Condition
}

func (s *DescribeAppViewTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppViewTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppViewTemplatesRequest) SetAppId(v string) *DescribeAppViewTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppViewTemplatesRequest) SetCondition(v *DescribeAppViewTemplatesRequestCondition) *DescribeAppViewTemplatesRequest {
	s.Condition = v
	return s
}

func (s *DescribeAppViewTemplatesRequest) SetPageNum(v int32) *DescribeAppViewTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppViewTemplatesRequest) SetPageSize(v int32) *DescribeAppViewTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppViewTemplatesRequest) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppViewTemplatesRequestCondition struct {
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Bj6D****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppViewTemplatesRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewTemplatesRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeAppViewTemplatesRequestCondition) GetName() *string {
	return s.Name
}

func (s *DescribeAppViewTemplatesRequestCondition) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppViewTemplatesRequestCondition) SetName(v string) *DescribeAppViewTemplatesRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeAppViewTemplatesRequestCondition) SetTemplateId(v string) *DescribeAppViewTemplatesRequestCondition {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppViewTemplatesRequestCondition) Validate() error {
	return dara.Validate(s)
}
