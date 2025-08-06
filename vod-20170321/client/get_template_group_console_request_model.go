// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateGroupConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetTemplateGroupConsoleRequest
	GetGroupId() *string
	SetOwnerId(v int64) *GetTemplateGroupConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetTemplateGroupConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetTemplateGroupConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *GetTemplateGroupConsoleRequest
	GetResourceRealOwnerId() *int64
}

type GetTemplateGroupConsoleRequest struct {
	// This parameter is required.
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s GetTemplateGroupConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateGroupConsoleRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateGroupConsoleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetTemplateGroupConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetTemplateGroupConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetTemplateGroupConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetTemplateGroupConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetTemplateGroupConsoleRequest) SetGroupId(v string) *GetTemplateGroupConsoleRequest {
	s.GroupId = &v
	return s
}

func (s *GetTemplateGroupConsoleRequest) SetOwnerId(v int64) *GetTemplateGroupConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTemplateGroupConsoleRequest) SetResourceOwnerAccount(v string) *GetTemplateGroupConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTemplateGroupConsoleRequest) SetResourceOwnerId(v int64) *GetTemplateGroupConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTemplateGroupConsoleRequest) SetResourceRealOwnerId(v int64) *GetTemplateGroupConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetTemplateGroupConsoleRequest) Validate() error {
	return dara.Validate(s)
}
