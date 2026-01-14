// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumbers(v []*string) *PageRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *PageRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *PageRequest
	GetPageNo() *int64
	SetPageSize(v int64) *PageRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *PageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PageRequest
	GetResourceOwnerId() *int64
}

type PageRequest struct {
	Numbers []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PageRequest) String() string {
	return dara.Prettify(s)
}

func (s PageRequest) GoString() string {
	return s.String()
}

func (s *PageRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *PageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PageRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PageRequest) SetNumbers(v []*string) *PageRequest {
	s.Numbers = v
	return s
}

func (s *PageRequest) SetOwnerId(v int64) *PageRequest {
	s.OwnerId = &v
	return s
}

func (s *PageRequest) SetPageNo(v int64) *PageRequest {
	s.PageNo = &v
	return s
}

func (s *PageRequest) SetPageSize(v int64) *PageRequest {
	s.PageSize = &v
	return s
}

func (s *PageRequest) SetResourceOwnerAccount(v string) *PageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PageRequest) SetResourceOwnerId(v int64) *PageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PageRequest) Validate() error {
	return dara.Validate(s)
}
