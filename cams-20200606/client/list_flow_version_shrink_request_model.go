// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListFlowVersionShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ListFlowVersionShrinkRequest
	GetBizExtendShrink() *string
	SetFlowCode(v string) *ListFlowVersionShrinkRequest
	GetFlowCode() *string
	SetOwnerId(v int64) *ListFlowVersionShrinkRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *ListFlowVersionShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListFlowVersionShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListFlowVersionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowVersionShrinkRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListFlowVersionShrinkRequest
	GetStatus() *string
}

type ListFlowVersionShrinkRequest struct {
	// The business tenant code. The default value is ALICOM_OPAAS.
	//
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Business extension information. The default value is an empty collection.
	//
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// The flow code. Find the flow code in the [flow editor](https://chatapp.console.aliyun.com/ChatFlowBuilder).
	//
	// example:
	//
	// 9ccc41**************************
	FlowCode *string `json:"FlowCode,omitempty" xml:"FlowCode,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the flow version. Valid values:
	//
	// - DRAFT: The flow is a draft.
	//
	// - DELETED: The flow is deleted.
	//
	// - ONLINE: The flow is published.
	//
	// - OFFLINE: The flow is unpublished.
	//
	// example:
	//
	// DRAFT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFlowVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFlowVersionShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListFlowVersionShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ListFlowVersionShrinkRequest) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ListFlowVersionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowVersionShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListFlowVersionShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListFlowVersionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowVersionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowVersionShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFlowVersionShrinkRequest) SetBizCode(v string) *ListFlowVersionShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetBizExtendShrink(v string) *ListFlowVersionShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetFlowCode(v string) *ListFlowVersionShrinkRequest {
	s.FlowCode = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetOwnerId(v int64) *ListFlowVersionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetPageNo(v int64) *ListFlowVersionShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetPageSize(v int64) *ListFlowVersionShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetResourceOwnerAccount(v string) *ListFlowVersionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetResourceOwnerId(v int64) *ListFlowVersionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) SetStatus(v string) *ListFlowVersionShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListFlowVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
