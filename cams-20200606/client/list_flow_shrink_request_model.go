// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListFlowShrinkRequest
	GetCustSpaceId() *string
	SetFlowName(v string) *ListFlowShrinkRequest
	GetFlowName() *string
	SetOwnerId(v int64) *ListFlowShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListFlowShrinkRequest
	GetPageShrink() *string
	SetResourceOwnerAccount(v string) *ListFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type ListFlowShrinkRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 99948484
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow that you want to query. If FlowName is left empty, the information about all Flows is queried.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The returned pages.
	PageShrink           *string `json:"Page,omitempty" xml:"Page,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFlowShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListFlowShrinkRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowShrinkRequest) SetCustSpaceId(v string) *ListFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListFlowShrinkRequest) SetFlowName(v string) *ListFlowShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowShrinkRequest) SetOwnerId(v int64) *ListFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFlowShrinkRequest) SetPageShrink(v string) *ListFlowShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListFlowShrinkRequest) SetResourceOwnerAccount(v string) *ListFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowShrinkRequest) SetResourceOwnerId(v int64) *ListFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
