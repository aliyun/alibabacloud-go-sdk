// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FailoverDBClusterZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *FailoverDBClusterZonalRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *FailoverDBClusterZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *FailoverDBClusterZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *FailoverDBClusterZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *FailoverDBClusterZonalRequest
	GetResourceOwnerId() *int64
	SetRollBackForDisaster(v bool) *FailoverDBClusterZonalRequest
	GetRollBackForDisaster() *bool
	SetTargetDBNodeId(v string) *FailoverDBClusterZonalRequest
	GetTargetDBNodeId() *string
	SetTargetZoneType(v string) *FailoverDBClusterZonalRequest
	GetTargetZoneType() *string
}

type FailoverDBClusterZonalRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// false
	RollBackForDisaster *bool `json:"RollBackForDisaster,omitempty" xml:"RollBackForDisaster,omitempty"`
	// example:
	//
	// pi-***********
	TargetDBNodeId *string `json:"TargetDBNodeId,omitempty" xml:"TargetDBNodeId,omitempty"`
	// example:
	//
	// Primary
	TargetZoneType *string `json:"TargetZoneType,omitempty" xml:"TargetZoneType,omitempty"`
}

func (s FailoverDBClusterZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterZonalRequest) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FailoverDBClusterZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *FailoverDBClusterZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *FailoverDBClusterZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *FailoverDBClusterZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *FailoverDBClusterZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *FailoverDBClusterZonalRequest) GetRollBackForDisaster() *bool {
	return s.RollBackForDisaster
}

func (s *FailoverDBClusterZonalRequest) GetTargetDBNodeId() *string {
	return s.TargetDBNodeId
}

func (s *FailoverDBClusterZonalRequest) GetTargetZoneType() *string {
	return s.TargetZoneType
}

func (s *FailoverDBClusterZonalRequest) SetClientToken(v string) *FailoverDBClusterZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetDBClusterId(v string) *FailoverDBClusterZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetOwnerAccount(v string) *FailoverDBClusterZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetOwnerId(v int64) *FailoverDBClusterZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetResourceOwnerAccount(v string) *FailoverDBClusterZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetResourceOwnerId(v int64) *FailoverDBClusterZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetRollBackForDisaster(v bool) *FailoverDBClusterZonalRequest {
	s.RollBackForDisaster = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetTargetDBNodeId(v string) *FailoverDBClusterZonalRequest {
	s.TargetDBNodeId = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) SetTargetZoneType(v string) *FailoverDBClusterZonalRequest {
	s.TargetZoneType = &v
	return s
}

func (s *FailoverDBClusterZonalRequest) Validate() error {
	return dara.Validate(s)
}
