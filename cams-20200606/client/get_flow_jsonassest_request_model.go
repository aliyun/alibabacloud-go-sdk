// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowJSONAssestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetFlowJSONAssestRequest
	GetCustSpaceId() *string
	SetFlowId(v string) *GetFlowJSONAssestRequest
	GetFlowId() *string
	SetOwnerId(v int64) *GetFlowJSONAssestRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetFlowJSONAssestRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFlowJSONAssestRequest
	GetResourceOwnerId() *int64
}

type GetFlowJSONAssestRequest struct {
	// The Space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the
	//
	// <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)<props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)
	//
	// page.
	//
	// example:
	//
	// cams-kei****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The Flow ID returned when you created the Flow by calling the [CreateFlow](https://help.aliyun.com/document_detail/2638742.html) operation.
	//
	// example:
	//
	// 17b6****************************
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetFlowJSONAssestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowJSONAssestRequest) GoString() string {
	return s.String()
}

func (s *GetFlowJSONAssestRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetFlowJSONAssestRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowJSONAssestRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFlowJSONAssestRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFlowJSONAssestRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFlowJSONAssestRequest) SetCustSpaceId(v string) *GetFlowJSONAssestRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetFlowJSONAssestRequest) SetFlowId(v string) *GetFlowJSONAssestRequest {
	s.FlowId = &v
	return s
}

func (s *GetFlowJSONAssestRequest) SetOwnerId(v int64) *GetFlowJSONAssestRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFlowJSONAssestRequest) SetResourceOwnerAccount(v string) *GetFlowJSONAssestRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFlowJSONAssestRequest) SetResourceOwnerId(v int64) *GetFlowJSONAssestRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFlowJSONAssestRequest) Validate() error {
	return dara.Validate(s)
}
