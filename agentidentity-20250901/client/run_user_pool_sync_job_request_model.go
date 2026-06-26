// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunUserPoolSyncJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderType(v string) *RunUserPoolSyncJobRequest
	GetIdentityProviderType() *string
	SetMaxSyncUsers(v string) *RunUserPoolSyncJobRequest
	GetMaxSyncUsers() *string
	SetUserPoolName(v string) *RunUserPoolSyncJobRequest
	GetUserPoolName() *string
}

type RunUserPoolSyncJobRequest struct {
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	MaxSyncUsers         *string `json:"MaxSyncUsers,omitempty" xml:"MaxSyncUsers,omitempty"`
	UserPoolName         *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s RunUserPoolSyncJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RunUserPoolSyncJobRequest) GoString() string {
	return s.String()
}

func (s *RunUserPoolSyncJobRequest) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *RunUserPoolSyncJobRequest) GetMaxSyncUsers() *string {
	return s.MaxSyncUsers
}

func (s *RunUserPoolSyncJobRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *RunUserPoolSyncJobRequest) SetIdentityProviderType(v string) *RunUserPoolSyncJobRequest {
	s.IdentityProviderType = &v
	return s
}

func (s *RunUserPoolSyncJobRequest) SetMaxSyncUsers(v string) *RunUserPoolSyncJobRequest {
	s.MaxSyncUsers = &v
	return s
}

func (s *RunUserPoolSyncJobRequest) SetUserPoolName(v string) *RunUserPoolSyncJobRequest {
	s.UserPoolName = &v
	return s
}

func (s *RunUserPoolSyncJobRequest) Validate() error {
	return dara.Validate(s)
}
