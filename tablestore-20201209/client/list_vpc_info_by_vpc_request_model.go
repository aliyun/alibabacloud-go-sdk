// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int64) *ListVpcInfoByVpcRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByVpcRequest
	GetPageSize() *int64
	SetVpcId(v string) *ListVpcInfoByVpcRequest
	GetVpcId() *string
}

type ListVpcInfoByVpcRequest struct {
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// VPC ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1***********0mh8
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcInfoByVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByVpcRequest) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByVpcRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListVpcInfoByVpcRequest) SetPageNum(v int64) *ListVpcInfoByVpcRequest {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetPageSize(v int64) *ListVpcInfoByVpcRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) SetVpcId(v string) *ListVpcInfoByVpcRequest {
	s.VpcId = &v
	return s
}

func (s *ListVpcInfoByVpcRequest) Validate() error {
	return dara.Validate(s)
}
