// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoriesShrink(v string) *CreateFlowShrinkRequest
	GetCategoriesShrink() *string
	SetCustSpaceId(v string) *CreateFlowShrinkRequest
	GetCustSpaceId() *string
	SetEndpointUri(v string) *CreateFlowShrinkRequest
	GetEndpointUri() *string
	SetFlowName(v string) *CreateFlowShrinkRequest
	GetFlowName() *string
	SetOwnerId(v int64) *CreateFlowShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateFlowShrinkRequest struct {
	// This parameter is required.
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// http://www.***.com
	EndpointUri *string `json:"EndpointUri,omitempty" xml:"EndpointUri,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	FlowName             *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowShrinkRequest) GetCategoriesShrink() *string {
	return s.CategoriesShrink
}

func (s *CreateFlowShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateFlowShrinkRequest) GetEndpointUri() *string {
	return s.EndpointUri
}

func (s *CreateFlowShrinkRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowShrinkRequest) SetCategoriesShrink(v string) *CreateFlowShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetCustSpaceId(v string) *CreateFlowShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetEndpointUri(v string) *CreateFlowShrinkRequest {
	s.EndpointUri = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetFlowName(v string) *CreateFlowShrinkRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetOwnerId(v int64) *CreateFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetResourceOwnerAccount(v string) *CreateFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowShrinkRequest) SetResourceOwnerId(v int64) *CreateFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
