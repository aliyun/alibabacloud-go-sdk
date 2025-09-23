// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *ModifyAIDBClusterDescriptionRequest
	GetDBClusterDescription() *string
	SetDBClusterId(v string) *ModifyAIDBClusterDescriptionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyAIDBClusterDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAIDBClusterDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAIDBClusterDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAIDBClusterDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyAIDBClusterDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAIDBClusterDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterDescriptionRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *ModifyAIDBClusterDescriptionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAIDBClusterDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAIDBClusterDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAIDBClusterDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAIDBClusterDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAIDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyAIDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyAIDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyAIDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyAIDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAIDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAIDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
