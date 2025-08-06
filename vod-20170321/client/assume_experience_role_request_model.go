// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeExperienceRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCookie(v string) *AssumeExperienceRoleRequest
	GetCookie() *string
	SetData(v string) *AssumeExperienceRoleRequest
	GetData() *string
	SetOwnerId(v int64) *AssumeExperienceRoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssumeExperienceRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssumeExperienceRoleRequest
	GetResourceOwnerId() *int64
}

type AssumeExperienceRoleRequest struct {
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// This parameter is required.
	Data                 *string `json:"Data,omitempty" xml:"Data,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssumeExperienceRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeExperienceRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeExperienceRoleRequest) GetCookie() *string {
	return s.Cookie
}

func (s *AssumeExperienceRoleRequest) GetData() *string {
	return s.Data
}

func (s *AssumeExperienceRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssumeExperienceRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssumeExperienceRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssumeExperienceRoleRequest) SetCookie(v string) *AssumeExperienceRoleRequest {
	s.Cookie = &v
	return s
}

func (s *AssumeExperienceRoleRequest) SetData(v string) *AssumeExperienceRoleRequest {
	s.Data = &v
	return s
}

func (s *AssumeExperienceRoleRequest) SetOwnerId(v int64) *AssumeExperienceRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AssumeExperienceRoleRequest) SetResourceOwnerAccount(v string) *AssumeExperienceRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssumeExperienceRoleRequest) SetResourceOwnerId(v int64) *AssumeExperienceRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssumeExperienceRoleRequest) Validate() error {
	return dara.Validate(s)
}
