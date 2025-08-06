// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteTemplateGroupConsoleRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DeleteTemplateGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTemplateGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTemplateGroupConsoleRequest
	GetResourceOwnerId() *int64
}

type DeleteTemplateGroupConsoleRequest struct {
	// This parameter is required.
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteTemplateGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateGroupConsoleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteTemplateGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTemplateGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTemplateGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTemplateGroupConsoleRequest) SetGroupId(v string) *DeleteTemplateGroupConsoleRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteTemplateGroupConsoleRequest) SetOwnerId(v int64) *DeleteTemplateGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTemplateGroupConsoleRequest) SetResourceOwnerAccount(v string) *DeleteTemplateGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTemplateGroupConsoleRequest) SetResourceOwnerId(v int64) *DeleteTemplateGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTemplateGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
