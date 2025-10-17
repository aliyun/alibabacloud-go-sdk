// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStreamingOutTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppStreamingOutTemplatesRequest
	GetAppId() *string
	SetCondition(v *DescribeAppStreamingOutTemplatesRequestCondition) *DescribeAppStreamingOutTemplatesRequest
	GetCondition() *DescribeAppStreamingOutTemplatesRequestCondition
	SetPageNum(v int32) *DescribeAppStreamingOutTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppStreamingOutTemplatesRequest
	GetPageSize() *int32
}

type DescribeAppStreamingOutTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId     *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Condition *DescribeAppStreamingOutTemplatesRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppStreamingOutTemplatesRequest) GetCondition() *DescribeAppStreamingOutTemplatesRequestCondition {
	return s.Condition
}

func (s *DescribeAppStreamingOutTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppStreamingOutTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetAppId(v string) *DescribeAppStreamingOutTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetCondition(v *DescribeAppStreamingOutTemplatesRequestCondition) *DescribeAppStreamingOutTemplatesRequest {
	s.Condition = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetPageNum(v int32) *DescribeAppStreamingOutTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetPageSize(v int32) *DescribeAppStreamingOutTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppStreamingOutTemplatesRequestCondition struct {
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Bj6D****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) GetName() *string {
	return s.Name
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) SetName(v string) *DescribeAppStreamingOutTemplatesRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) SetTemplateId(v string) *DescribeAppStreamingOutTemplatesRequestCondition {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) Validate() error {
	return dara.Validate(s)
}
