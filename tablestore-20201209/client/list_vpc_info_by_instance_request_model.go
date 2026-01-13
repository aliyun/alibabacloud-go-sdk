// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *ListVpcInfoByInstanceRequest
	GetInstanceName() *string
	SetPageNum(v int64) *ListVpcInfoByInstanceRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListVpcInfoByInstanceRequest
	GetPageSize() *int64
}

type ListVpcInfoByInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mkdd-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 8
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVpcInfoByInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVpcInfoByInstanceRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListVpcInfoByInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVpcInfoByInstanceRequest) SetInstanceName(v string) *ListVpcInfoByInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) SetPageNum(v int64) *ListVpcInfoByInstanceRequest {
	s.PageNum = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) SetPageSize(v int64) *ListVpcInfoByInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListVpcInfoByInstanceRequest) Validate() error {
	return dara.Validate(s)
}
