// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAudienceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *AddCustomAudienceUserRequest
	GetAdAccountId() *string
	SetBatchLastFlag(v bool) *AddCustomAudienceUserRequest
	GetBatchLastFlag() *bool
	SetCustSpaceId(v string) *AddCustomAudienceUserRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *AddCustomAudienceUserRequest
	GetCustomAudienceId() *string
	SetEstimatedNumTotal(v int64) *AddCustomAudienceUserRequest
	GetEstimatedNumTotal() *int64
	SetOwnerId(v int64) *AddCustomAudienceUserRequest
	GetOwnerId() *int64
	SetPageId(v string) *AddCustomAudienceUserRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *AddCustomAudienceUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddCustomAudienceUserRequest
	GetResourceOwnerId() *int64
	SetUsers(v []*AddCustomAudienceUserRequestUsers) *AddCustomAudienceUserRequest
	GetUsers() []*AddCustomAudienceUserRequestUsers
}

type AddCustomAudienceUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3939**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// example:
	//
	// false
	BatchLastFlag *bool `json:"BatchLastFlag,omitempty" xml:"BatchLastFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 399**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// example:
	//
	// 26
	EstimatedNumTotal *int64 `json:"EstimatedNumTotal,omitempty" xml:"EstimatedNumTotal,omitempty"`
	OwnerId           *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239**
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	Users []*AddCustomAudienceUserRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s AddCustomAudienceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAudienceUserRequest) GoString() string {
	return s.String()
}

func (s *AddCustomAudienceUserRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *AddCustomAudienceUserRequest) GetBatchLastFlag() *bool {
	return s.BatchLastFlag
}

func (s *AddCustomAudienceUserRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddCustomAudienceUserRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *AddCustomAudienceUserRequest) GetEstimatedNumTotal() *int64 {
	return s.EstimatedNumTotal
}

func (s *AddCustomAudienceUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCustomAudienceUserRequest) GetPageId() *string {
	return s.PageId
}

func (s *AddCustomAudienceUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddCustomAudienceUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddCustomAudienceUserRequest) GetUsers() []*AddCustomAudienceUserRequestUsers {
	return s.Users
}

func (s *AddCustomAudienceUserRequest) SetAdAccountId(v string) *AddCustomAudienceUserRequest {
	s.AdAccountId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetBatchLastFlag(v bool) *AddCustomAudienceUserRequest {
	s.BatchLastFlag = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetCustSpaceId(v string) *AddCustomAudienceUserRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetCustomAudienceId(v string) *AddCustomAudienceUserRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetEstimatedNumTotal(v int64) *AddCustomAudienceUserRequest {
	s.EstimatedNumTotal = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetOwnerId(v int64) *AddCustomAudienceUserRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetPageId(v string) *AddCustomAudienceUserRequest {
	s.PageId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetResourceOwnerAccount(v string) *AddCustomAudienceUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetResourceOwnerId(v int64) *AddCustomAudienceUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddCustomAudienceUserRequest) SetUsers(v []*AddCustomAudienceUserRequestUsers) *AddCustomAudienceUserRequest {
	s.Users = v
	return s
}

func (s *AddCustomAudienceUserRequest) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddCustomAudienceUserRequestUsers struct {
	// example:
	//
	// 示例值示例值示例值
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s AddCustomAudienceUserRequestUsers) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAudienceUserRequestUsers) GoString() string {
	return s.String()
}

func (s *AddCustomAudienceUserRequestUsers) GetEmail() *string {
	return s.Email
}

func (s *AddCustomAudienceUserRequestUsers) GetPhone() *string {
	return s.Phone
}

func (s *AddCustomAudienceUserRequestUsers) SetEmail(v string) *AddCustomAudienceUserRequestUsers {
	s.Email = &v
	return s
}

func (s *AddCustomAudienceUserRequestUsers) SetPhone(v string) *AddCustomAudienceUserRequestUsers {
	s.Phone = &v
	return s
}

func (s *AddCustomAudienceUserRequestUsers) Validate() error {
	return dara.Validate(s)
}
