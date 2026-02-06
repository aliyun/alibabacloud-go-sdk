// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *CreateFlowRequest
	GetCategories() []*string
	SetCustSpaceId(v string) *CreateFlowRequest
	GetCustSpaceId() *string
	SetEndpointUri(v string) *CreateFlowRequest
	GetEndpointUri() *string
	SetFlowName(v string) *CreateFlowRequest
	GetFlowName() *string
	SetOwnerId(v int64) *CreateFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowRequest
	GetResourceOwnerId() *int64
}

type CreateFlowRequest struct {
	// This parameter is required.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
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

func (s CreateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) GetCategories() []*string {
	return s.Categories
}

func (s *CreateFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateFlowRequest) GetEndpointUri() *string {
	return s.EndpointUri
}

func (s *CreateFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowRequest) SetCategories(v []*string) *CreateFlowRequest {
	s.Categories = v
	return s
}

func (s *CreateFlowRequest) SetCustSpaceId(v string) *CreateFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateFlowRequest) SetEndpointUri(v string) *CreateFlowRequest {
	s.EndpointUri = &v
	return s
}

func (s *CreateFlowRequest) SetFlowName(v string) *CreateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowRequest) SetOwnerId(v int64) *CreateFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowRequest) SetResourceOwnerAccount(v string) *CreateFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowRequest) SetResourceOwnerId(v int64) *CreateFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowRequest) Validate() error {
	return dara.Validate(s)
}
