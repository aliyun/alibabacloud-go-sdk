// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListFlowNodeGroupRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListFlowNodeGroupRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListFlowNodeGroupRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListFlowNodeGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowNodeGroupRequest
	GetResourceOwnerId() *int64
}

type ListFlowNodeGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFlowNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *ListFlowNodeGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowNodeGroupRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListFlowNodeGroupRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFlowNodeGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowNodeGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowNodeGroupRequest) SetOwnerId(v int64) *ListFlowNodeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFlowNodeGroupRequest) SetPageNo(v int64) *ListFlowNodeGroupRequest {
	s.PageNo = &v
	return s
}

func (s *ListFlowNodeGroupRequest) SetPageSize(v int64) *ListFlowNodeGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodeGroupRequest) SetResourceOwnerAccount(v string) *ListFlowNodeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowNodeGroupRequest) SetResourceOwnerId(v int64) *ListFlowNodeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
