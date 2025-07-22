// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppAgentTemplatesRequest
	GetAppId() *string
	SetId(v string) *DescribeAppAgentTemplatesRequest
	GetId() *string
	SetName(v string) *DescribeAppAgentTemplatesRequest
	GetName() *string
	SetPageNum(v int32) *DescribeAppAgentTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppAgentTemplatesRequest
	GetPageSize() *int32
}

type DescribeAppAgentTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123121414122313121313
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 智能体模版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppAgentTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppAgentTemplatesRequest) GetId() *string {
	return s.Id
}

func (s *DescribeAppAgentTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppAgentTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppAgentTemplatesRequest) SetAppId(v string) *DescribeAppAgentTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppAgentTemplatesRequest) SetId(v string) *DescribeAppAgentTemplatesRequest {
	s.Id = &v
	return s
}

func (s *DescribeAppAgentTemplatesRequest) SetName(v string) *DescribeAppAgentTemplatesRequest {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesRequest) SetPageNum(v int32) *DescribeAppAgentTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppAgentTemplatesRequest) SetPageSize(v int32) *DescribeAppAgentTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppAgentTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
