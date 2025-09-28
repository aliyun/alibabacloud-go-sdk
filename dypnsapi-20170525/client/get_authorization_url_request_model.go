// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetAuthorizationUrlRequest
	GetEndDate() *string
	SetOwnerId(v int64) *GetAuthorizationUrlRequest
	GetOwnerId() *int64
	SetPhoneNo(v string) *GetAuthorizationUrlRequest
	GetPhoneNo() *string
	SetResourceOwnerAccount(v string) *GetAuthorizationUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAuthorizationUrlRequest
	GetResourceOwnerId() *int64
	SetSchemeId(v int64) *GetAuthorizationUrlRequest
	GetSchemeId() *int64
}

type GetAuthorizationUrlRequest struct {
	// The authorization end date, which is in the yyyy-MM-dd format. This parameter is required for services of contract type.
	//
	// example:
	//
	// 2020–12–28
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the authorization scenario. You can view the ID of the authorization scenario on the **Authorization Scenario Management*	- page in the **Phone Number Verification Service console**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 234****
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s GetAuthorizationUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetAuthorizationUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAuthorizationUrlRequest) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *GetAuthorizationUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAuthorizationUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAuthorizationUrlRequest) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetAuthorizationUrlRequest) SetEndDate(v string) *GetAuthorizationUrlRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetPhoneNo(v string) *GetAuthorizationUrlRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerAccount(v string) *GetAuthorizationUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetSchemeId(v int64) *GetAuthorizationUrlRequest {
	s.SchemeId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) Validate() error {
	return dara.Validate(s)
}
