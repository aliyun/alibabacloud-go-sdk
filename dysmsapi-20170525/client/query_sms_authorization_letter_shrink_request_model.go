// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationLetterIdListShrink(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetAuthorizationLetterIdListShrink() *string
	SetOrganizationCode(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetSignName() *string
	SetState(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetState() *string
	SetStatus(v string) *QuerySmsAuthorizationLetterShrinkRequest
	GetStatus() *string
}

type QuerySmsAuthorizationLetterShrinkRequest struct {
	// The list of letter of authorization IDs.
	AuthorizationLetterIdListShrink *string `json:"AuthorizationLetterIdList,omitempty" xml:"AuthorizationLetterIdList,omitempty"`
	// The unified social credit code of the authorizing party. The length cannot exceed 150 characters.
	//
	// example:
	//
	// 9****************A
	OrganizationCode     *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature name. If the authorization scope includes multiple signatures when you create the letter of authorization, the letters of authorization that contain the signature are returned.
	//
	// example:
	//
	// 菜鸟网络
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The review status of the letter of authorization, which is related to the review status of the signature. Valid values:
	//
	// - **INT**: Pending review. The letter of authorization has been created. After you submit a signature application, it enters the review process.
	//
	// - **PASSED**: Review passed. When a signature in the authorized signature scope of the letter of authorization passes the review, the status of the letter of authorization changes to PASSED.
	//
	// example:
	//
	// PASSED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The availability status of the letter of authorization, which is related to the validity period of the letter of authorization. Valid values:
	//
	// - **VALID**: Available. The letter of authorization is within the validity period.
	//
	// - **INVALID**: Unavailable. The letter of authorization has expired.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySmsAuthorizationLetterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetAuthorizationLetterIdListShrink() *string {
	return s.AuthorizationLetterIdListShrink
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetState() *string {
	return s.State
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetAuthorizationLetterIdListShrink(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.AuthorizationLetterIdListShrink = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetOrganizationCode(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetSignName(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.SignName = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetState(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.State = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) SetStatus(v string) *QuerySmsAuthorizationLetterShrinkRequest {
	s.Status = &v
	return s
}

func (s *QuerySmsAuthorizationLetterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
