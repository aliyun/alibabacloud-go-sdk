// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeRouterWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeRouterWeightRequest
	GetDBClusterId() *string
	SetDBNodeWeights(v string) *ModifyDBNodeRouterWeightRequest
	GetDBNodeWeights() *string
	SetOwnerAccount(v string) *ModifyDBNodeRouterWeightRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeRouterWeightRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBNodeRouterWeightRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeRouterWeightRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyDBNodeRouterWeightRequest
	GetSecurityToken() *string
}

type ModifyDBNodeRouterWeightRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all clusters in a region, including their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node weights.
	//
	// example:
	//
	// "pi-xxxxxxxxxxx":"1","pi-xxxxxxx":"1"}
	DBNodeWeights        *string `json:"DBNodeWeights,omitempty" xml:"DBNodeWeights,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyDBNodeRouterWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeRouterWeightRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeRouterWeightRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeRouterWeightRequest) GetDBNodeWeights() *string {
	return s.DBNodeWeights
}

func (s *ModifyDBNodeRouterWeightRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeRouterWeightRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeRouterWeightRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeRouterWeightRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeRouterWeightRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDBNodeRouterWeightRequest) SetDBClusterId(v string) *ModifyDBNodeRouterWeightRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetDBNodeWeights(v string) *ModifyDBNodeRouterWeightRequest {
	s.DBNodeWeights = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetOwnerAccount(v string) *ModifyDBNodeRouterWeightRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetOwnerId(v int64) *ModifyDBNodeRouterWeightRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeRouterWeightRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetResourceOwnerId(v int64) *ModifyDBNodeRouterWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) SetSecurityToken(v string) *ModifyDBNodeRouterWeightRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBNodeRouterWeightRequest) Validate() error {
	return dara.Validate(s)
}
