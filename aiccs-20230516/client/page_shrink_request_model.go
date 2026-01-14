// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumbersShrink(v string) *PageShrinkRequest
	GetNumbersShrink() *string
	SetOwnerId(v int64) *PageShrinkRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *PageShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *PageShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *PageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PageShrinkRequest
	GetResourceOwnerId() *int64
}

type PageShrinkRequest struct {
	NumbersShrink *string `json:"Numbers,omitempty" xml:"Numbers,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s PageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PageShrinkRequest) GoString() string {
	return s.String()
}

func (s *PageShrinkRequest) GetNumbersShrink() *string {
	return s.NumbersShrink
}

func (s *PageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PageShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PageShrinkRequest) SetNumbersShrink(v string) *PageShrinkRequest {
	s.NumbersShrink = &v
	return s
}

func (s *PageShrinkRequest) SetOwnerId(v int64) *PageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *PageShrinkRequest) SetPageNo(v int64) *PageShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *PageShrinkRequest) SetPageSize(v int64) *PageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *PageShrinkRequest) SetResourceOwnerAccount(v string) *PageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PageShrinkRequest) SetResourceOwnerId(v int64) *PageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
