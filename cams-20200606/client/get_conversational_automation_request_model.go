// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationalAutomationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetConversationalAutomationRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetConversationalAutomationRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetConversationalAutomationRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetConversationalAutomationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetConversationalAutomationRequest
	GetResourceOwnerId() *int64
}

type GetConversationalAutomationRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account or the instance ID of the customer of Alibaba Cloud.
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-3ie***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number of the enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86130000***
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetConversationalAutomationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConversationalAutomationRequest) GoString() string {
	return s.String()
}

func (s *GetConversationalAutomationRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetConversationalAutomationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetConversationalAutomationRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetConversationalAutomationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetConversationalAutomationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetConversationalAutomationRequest) SetCustSpaceId(v string) *GetConversationalAutomationRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetConversationalAutomationRequest) SetOwnerId(v int64) *GetConversationalAutomationRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConversationalAutomationRequest) SetPhoneNumber(v string) *GetConversationalAutomationRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetConversationalAutomationRequest) SetResourceOwnerAccount(v string) *GetConversationalAutomationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConversationalAutomationRequest) SetResourceOwnerId(v int64) *GetConversationalAutomationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConversationalAutomationRequest) Validate() error {
	return dara.Validate(s)
}
