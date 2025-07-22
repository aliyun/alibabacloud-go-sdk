// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppRecordTemplatesRequest
	GetAppId() *string
	SetClientToken(v string) *DescribeAppRecordTemplatesRequest
	GetClientToken() *string
	SetCondition(v *DescribeAppRecordTemplatesRequestCondition) *DescribeAppRecordTemplatesRequest
	GetCondition() *DescribeAppRecordTemplatesRequestCondition
	SetPageNum(v int32) *DescribeAppRecordTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeAppRecordTemplatesRequest
	GetPageSize() *int32
}

type DescribeAppRecordTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string                                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Condition   *DescribeAppRecordTemplatesRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppRecordTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppRecordTemplatesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAppRecordTemplatesRequest) GetCondition() *DescribeAppRecordTemplatesRequestCondition {
	return s.Condition
}

func (s *DescribeAppRecordTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeAppRecordTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppRecordTemplatesRequest) SetAppId(v string) *DescribeAppRecordTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequest) SetClientToken(v string) *DescribeAppRecordTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequest) SetCondition(v *DescribeAppRecordTemplatesRequestCondition) *DescribeAppRecordTemplatesRequest {
	s.Condition = v
	return s
}

func (s *DescribeAppRecordTemplatesRequest) SetPageNum(v int32) *DescribeAppRecordTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequest) SetPageSize(v int32) *DescribeAppRecordTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeAppRecordTemplatesRequestCondition struct {
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ac7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppRecordTemplatesRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesRequestCondition) GetName() *string {
	return s.Name
}

func (s *DescribeAppRecordTemplatesRequestCondition) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeAppRecordTemplatesRequestCondition) SetName(v string) *DescribeAppRecordTemplatesRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequestCondition) SetTemplateId(v string) *DescribeAppRecordTemplatesRequestCondition {
	s.TemplateId = &v
	return s
}

func (s *DescribeAppRecordTemplatesRequestCondition) Validate() error {
	return dara.Validate(s)
}
