// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *ModifyFlowRequest
	GetCategories() []*string
	SetCustSpaceId(v string) *ModifyFlowRequest
	GetCustSpaceId() *string
	SetEndpointUri(v string) *ModifyFlowRequest
	GetEndpointUri() *string
	SetFlowId(v string) *ModifyFlowRequest
	GetFlowId() *string
	SetFlowName(v string) *ModifyFlowRequest
	GetFlowName() *string
	SetOwnerId(v int64) *ModifyFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFlowRequest
	GetResourceOwnerId() *int64
}

type ModifyFlowRequest struct {
	// This parameter is required.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// http://www.***.com
	EndpointUri *string `json:"EndpointUri,omitempty" xml:"EndpointUri,omitempty"`
	// example:
	//
	// 示例值
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
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

func (s ModifyFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRequest) GetCategories() []*string {
	return s.Categories
}

func (s *ModifyFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ModifyFlowRequest) GetEndpointUri() *string {
	return s.EndpointUri
}

func (s *ModifyFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ModifyFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ModifyFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFlowRequest) SetCategories(v []*string) *ModifyFlowRequest {
	s.Categories = v
	return s
}

func (s *ModifyFlowRequest) SetCustSpaceId(v string) *ModifyFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ModifyFlowRequest) SetEndpointUri(v string) *ModifyFlowRequest {
	s.EndpointUri = &v
	return s
}

func (s *ModifyFlowRequest) SetFlowId(v string) *ModifyFlowRequest {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowRequest) SetFlowName(v string) *ModifyFlowRequest {
	s.FlowName = &v
	return s
}

func (s *ModifyFlowRequest) SetOwnerId(v int64) *ModifyFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowRequest) SetResourceOwnerAccount(v string) *ModifyFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFlowRequest) SetResourceOwnerId(v int64) *ModifyFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFlowRequest) Validate() error {
	return dara.Validate(s)
}
