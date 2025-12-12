// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyCronDetailRequest
	GetCurrentPage() *int32
	SetExtend(v string) *DescribePropertyCronDetailRequest
	GetExtend() *string
	SetNextToken(v string) *DescribePropertyCronDetailRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribePropertyCronDetailRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribePropertyCronDetailRequest
	GetRemark() *string
	SetSource(v string) *DescribePropertyCronDetailRequest
	GetSource() *string
	SetUseNextToken(v bool) *DescribePropertyCronDetailRequest
	GetUseNextToken() *bool
	SetUser(v string) *DescribePropertyCronDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyCronDetailRequest
	GetUuid() *string
}

type DescribePropertyCronDetailRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy search by path to the scheduled task is supported. If you want to use fuzzy search, set the parameter to **1**. If you set the parameter to a different value or leave the parameter empty, fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend    *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The path to the scheduled task.
	//
	// example:
	//
	// /etc/cron.d/root
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UseNextToken *bool   `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The username of the account that runs the scheduled task.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyCronDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyCronDetailRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribePropertyCronDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyCronDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyCronDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyCronDetailRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePropertyCronDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribePropertyCronDetailRequest) GetUser() *string {
	return s.User
}

func (s *DescribePropertyCronDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyCronDetailRequest) SetCurrentPage(v int32) *DescribePropertyCronDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetExtend(v string) *DescribePropertyCronDetailRequest {
	s.Extend = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetNextToken(v string) *DescribePropertyCronDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetPageSize(v int32) *DescribePropertyCronDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetRemark(v string) *DescribePropertyCronDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetSource(v string) *DescribePropertyCronDetailRequest {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetUseNextToken(v bool) *DescribePropertyCronDetailRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetUser(v string) *DescribePropertyCronDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetUuid(v string) *DescribePropertyCronDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) Validate() error {
	return dara.Validate(s)
}
