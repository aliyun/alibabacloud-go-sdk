// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodePrototypeV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListFlowNodePrototypeV2Request
	GetBizCode() *string
	SetGroupCode(v string) *ListFlowNodePrototypeV2Request
	GetGroupCode() *string
	SetKeyword(v string) *ListFlowNodePrototypeV2Request
	GetKeyword() *string
	SetOwnerId(v int64) *ListFlowNodePrototypeV2Request
	GetOwnerId() *int64
	SetPageNo(v int64) *ListFlowNodePrototypeV2Request
	GetPageNo() *int64
	SetPageSize(v int64) *ListFlowNodePrototypeV2Request
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListFlowNodePrototypeV2Request
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowNodePrototypeV2Request
	GetResourceOwnerId() *int64
}

type ListFlowNodePrototypeV2Request struct {
	// This parameter is required.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// Core
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// example:
	//
	// 示例值示例值
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFlowNodePrototypeV2Request) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodePrototypeV2Request) GoString() string {
	return s.String()
}

func (s *ListFlowNodePrototypeV2Request) GetBizCode() *string {
	return s.BizCode
}

func (s *ListFlowNodePrototypeV2Request) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ListFlowNodePrototypeV2Request) GetKeyword() *string {
	return s.Keyword
}

func (s *ListFlowNodePrototypeV2Request) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowNodePrototypeV2Request) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListFlowNodePrototypeV2Request) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFlowNodePrototypeV2Request) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowNodePrototypeV2Request) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowNodePrototypeV2Request) SetBizCode(v string) *ListFlowNodePrototypeV2Request {
	s.BizCode = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetGroupCode(v string) *ListFlowNodePrototypeV2Request {
	s.GroupCode = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetKeyword(v string) *ListFlowNodePrototypeV2Request {
	s.Keyword = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetOwnerId(v int64) *ListFlowNodePrototypeV2Request {
	s.OwnerId = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetPageNo(v int64) *ListFlowNodePrototypeV2Request {
	s.PageNo = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetPageSize(v int64) *ListFlowNodePrototypeV2Request {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetResourceOwnerAccount(v string) *ListFlowNodePrototypeV2Request {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) SetResourceOwnerId(v int64) *ListFlowNodePrototypeV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowNodePrototypeV2Request) Validate() error {
	return dara.Validate(s)
}
