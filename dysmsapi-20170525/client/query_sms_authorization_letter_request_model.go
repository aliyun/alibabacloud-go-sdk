// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAuthorizationLetterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationLetterIdList(v []*int64) *QuerySmsAuthorizationLetterRequest
	GetAuthorizationLetterIdList() []*int64
	SetOrganizationCode(v string) *QuerySmsAuthorizationLetterRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *QuerySmsAuthorizationLetterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySmsAuthorizationLetterRequest
	GetSignName() *string
	SetState(v string) *QuerySmsAuthorizationLetterRequest
	GetState() *string
	SetStatus(v string) *QuerySmsAuthorizationLetterRequest
	GetStatus() *string
}

type QuerySmsAuthorizationLetterRequest struct {
	// The list of letter of authorization IDs.
	AuthorizationLetterIdList []*int64 `json:"AuthorizationLetterIdList,omitempty" xml:"AuthorizationLetterIdList,omitempty" type:"Repeated"`
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

func (s QuerySmsAuthorizationLetterRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsAuthorizationLetterRequest) GetAuthorizationLetterIdList() []*int64 {
	return s.AuthorizationLetterIdList
}

func (s *QuerySmsAuthorizationLetterRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySmsAuthorizationLetterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsAuthorizationLetterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsAuthorizationLetterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsAuthorizationLetterRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsAuthorizationLetterRequest) GetState() *string {
	return s.State
}

func (s *QuerySmsAuthorizationLetterRequest) GetStatus() *string {
	return s.Status
}

func (s *QuerySmsAuthorizationLetterRequest) SetAuthorizationLetterIdList(v []*int64) *QuerySmsAuthorizationLetterRequest {
	s.AuthorizationLetterIdList = v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetOrganizationCode(v string) *QuerySmsAuthorizationLetterRequest {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetOwnerId(v int64) *QuerySmsAuthorizationLetterRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetResourceOwnerAccount(v string) *QuerySmsAuthorizationLetterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetResourceOwnerId(v int64) *QuerySmsAuthorizationLetterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetSignName(v string) *QuerySmsAuthorizationLetterRequest {
	s.SignName = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetState(v string) *QuerySmsAuthorizationLetterRequest {
	s.State = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) SetStatus(v string) *QuerySmsAuthorizationLetterRequest {
	s.Status = &v
	return s
}

func (s *QuerySmsAuthorizationLetterRequest) Validate() error {
	return dara.Validate(s)
}
