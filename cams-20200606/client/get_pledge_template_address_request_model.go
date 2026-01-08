// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPledgeTemplateAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetPledgeTemplateAddressRequest
	GetCustSpaceId() *string
	SetIndustryType(v string) *GetPledgeTemplateAddressRequest
	GetIndustryType() *string
	SetOwnerId(v int64) *GetPledgeTemplateAddressRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetPledgeTemplateAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPledgeTemplateAddressRequest
	GetResourceOwnerId() *int64
}

type GetPledgeTemplateAddressRequest struct {
	// example:
	//
	// cams-x***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// it
	IndustryType         *string `json:"IndustryType,omitempty" xml:"IndustryType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPledgeTemplateAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPledgeTemplateAddressRequest) GoString() string {
	return s.String()
}

func (s *GetPledgeTemplateAddressRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetPledgeTemplateAddressRequest) GetIndustryType() *string {
	return s.IndustryType
}

func (s *GetPledgeTemplateAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPledgeTemplateAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPledgeTemplateAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPledgeTemplateAddressRequest) SetCustSpaceId(v string) *GetPledgeTemplateAddressRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPledgeTemplateAddressRequest) SetIndustryType(v string) *GetPledgeTemplateAddressRequest {
	s.IndustryType = &v
	return s
}

func (s *GetPledgeTemplateAddressRequest) SetOwnerId(v int64) *GetPledgeTemplateAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPledgeTemplateAddressRequest) SetResourceOwnerAccount(v string) *GetPledgeTemplateAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPledgeTemplateAddressRequest) SetResourceOwnerId(v int64) *GetPledgeTemplateAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPledgeTemplateAddressRequest) Validate() error {
	return dara.Validate(s)
}
