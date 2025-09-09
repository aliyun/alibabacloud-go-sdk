// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeDescriptionRequest
	GetDBClusterId() *string
	SetDBNodeDescription(v string) *ModifyDBNodeDescriptionRequest
	GetDBNodeDescription() *string
	SetDBNodeId(v string) *ModifyDBNodeDescriptionRequest
	GetDBNodeId() *string
	SetOwnerAccount(v string) *ModifyDBNodeDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBNodeDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodeDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NodeDescriptionTest
	DBNodeDescription *string `json:"DBNodeDescription,omitempty" xml:"DBNodeDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pi-*****************
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBNodeDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeDescriptionRequest) GetDBNodeDescription() *string {
	return s.DBNodeDescription
}

func (s *ModifyDBNodeDescriptionRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *ModifyDBNodeDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeDescriptionRequest) SetDBClusterId(v string) *ModifyDBNodeDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetDBNodeDescription(v string) *ModifyDBNodeDescriptionRequest {
	s.DBNodeDescription = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetDBNodeId(v string) *ModifyDBNodeDescriptionRequest {
	s.DBNodeId = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetOwnerAccount(v string) *ModifyDBNodeDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetOwnerId(v int64) *ModifyDBNodeDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBNodeDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
