// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTemplateGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *SetDefaultTemplateGroupConsoleRequest
	GetGroupId() *string
	SetGroupSymbol(v string) *SetDefaultTemplateGroupConsoleRequest
	GetGroupSymbol() *string
	SetOwnerId(v int64) *SetDefaultTemplateGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetDefaultTemplateGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetDefaultTemplateGroupConsoleRequest
	GetResourceOwnerId() *int64
}

type SetDefaultTemplateGroupConsoleRequest struct {
	// This parameter is required.
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupSymbol          *string `json:"GroupSymbol,omitempty" xml:"GroupSymbol,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetDefaultTemplateGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTemplateGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultTemplateGroupConsoleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetDefaultTemplateGroupConsoleRequest) GetGroupSymbol() *string {
	return s.GroupSymbol
}

func (s *SetDefaultTemplateGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDefaultTemplateGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultTemplateGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetDefaultTemplateGroupConsoleRequest) SetGroupId(v string) *SetDefaultTemplateGroupConsoleRequest {
	s.GroupId = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleRequest) SetGroupSymbol(v string) *SetDefaultTemplateGroupConsoleRequest {
	s.GroupSymbol = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleRequest) SetOwnerId(v int64) *SetDefaultTemplateGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleRequest) SetResourceOwnerAccount(v string) *SetDefaultTemplateGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleRequest) SetResourceOwnerId(v int64) *SetDefaultTemplateGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultTemplateGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
