// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteCustomTemplateConsoleRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DeleteCustomTemplateConsoleRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCustomTemplateConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCustomTemplateConsoleRequest
	GetResourceOwnerId() *int64
	SetTemplateIds(v string) *DeleteCustomTemplateConsoleRequest
	GetTemplateIds() *string
}

type DeleteCustomTemplateConsoleRequest struct {
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s DeleteCustomTemplateConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateConsoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateConsoleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteCustomTemplateConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomTemplateConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCustomTemplateConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCustomTemplateConsoleRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *DeleteCustomTemplateConsoleRequest) SetGroupId(v string) *DeleteCustomTemplateConsoleRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteCustomTemplateConsoleRequest) SetOwnerId(v int64) *DeleteCustomTemplateConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomTemplateConsoleRequest) SetResourceOwnerAccount(v string) *DeleteCustomTemplateConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomTemplateConsoleRequest) SetResourceOwnerId(v int64) *DeleteCustomTemplateConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomTemplateConsoleRequest) SetTemplateIds(v string) *DeleteCustomTemplateConsoleRequest {
	s.TemplateIds = &v
	return s
}

func (s *DeleteCustomTemplateConsoleRequest) Validate() error {
	return dara.Validate(s)
}
