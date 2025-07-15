// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcAccessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccurateQuery(v bool) *DescribeVpcAccessesRequest
	GetAccurateQuery() *bool
	SetInstanceId(v string) *DescribeVpcAccessesRequest
	GetInstanceId() *string
	SetName(v string) *DescribeVpcAccessesRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeVpcAccessesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcAccessesRequest
	GetPageSize() *int32
	SetPort(v string) *DescribeVpcAccessesRequest
	GetPort() *string
	SetSecurityToken(v string) *DescribeVpcAccessesRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeVpcAccessesRequestTag) *DescribeVpcAccessesRequest
	GetTag() []*DescribeVpcAccessesRequestTag
	SetVpcAccessId(v string) *DescribeVpcAccessesRequest
	GetVpcAccessId() *string
	SetVpcId(v string) *DescribeVpcAccessesRequest
	GetVpcId() *string
}

type DescribeVpcAccessesRequest struct {
	// Whether to conduct precise queries
	//
	// example:
	//
	// false
	AccurateQuery *bool `json:"AccurateQuery,omitempty" xml:"AccurateQuery,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 10.199.26.***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the authorization. The name must be unique.
	//
	// example:
	//
	// wuying-edm-svc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service port.
	//
	// example:
	//
	// 8080
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The port number.
	Tag []*DescribeVpcAccessesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC authorization.
	//
	// example:
	//
	// vpc-*****ssds24
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf657qec7lx42paw3qxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcAccessesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesRequest) GetAccurateQuery() *bool {
	return s.AccurateQuery
}

func (s *DescribeVpcAccessesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpcAccessesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeVpcAccessesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcAccessesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcAccessesRequest) GetPort() *string {
	return s.Port
}

func (s *DescribeVpcAccessesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVpcAccessesRequest) GetTag() []*DescribeVpcAccessesRequestTag {
	return s.Tag
}

func (s *DescribeVpcAccessesRequest) GetVpcAccessId() *string {
	return s.VpcAccessId
}

func (s *DescribeVpcAccessesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcAccessesRequest) SetAccurateQuery(v bool) *DescribeVpcAccessesRequest {
	s.AccurateQuery = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetInstanceId(v string) *DescribeVpcAccessesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetName(v string) *DescribeVpcAccessesRequest {
	s.Name = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetPageNumber(v int32) *DescribeVpcAccessesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetPageSize(v int32) *DescribeVpcAccessesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetPort(v string) *DescribeVpcAccessesRequest {
	s.Port = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetSecurityToken(v string) *DescribeVpcAccessesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetTag(v []*DescribeVpcAccessesRequestTag) *DescribeVpcAccessesRequest {
	s.Tag = v
	return s
}

func (s *DescribeVpcAccessesRequest) SetVpcAccessId(v string) *DescribeVpcAccessesRequest {
	s.VpcAccessId = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetVpcId(v string) *DescribeVpcAccessesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcAccessesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcAccessesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// appname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcAccessesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcAccessesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcAccessesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcAccessesRequestTag) SetKey(v string) *DescribeVpcAccessesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVpcAccessesRequestTag) SetValue(v string) *DescribeVpcAccessesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVpcAccessesRequestTag) Validate() error {
	return dara.Validate(s)
}
