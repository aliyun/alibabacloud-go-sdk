// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartDBNodeZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *RestartDBNodeZonalRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *RestartDBNodeZonalRequest
	GetDBNodeId() *string
	SetOwnerAccount(v string) *RestartDBNodeZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestartDBNodeZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestartDBNodeZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestartDBNodeZonalRequest
	GetResourceOwnerId() *int64
}

type RestartDBNodeZonalRequest struct {
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pi-*************
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartDBNodeZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeZonalRequest) GoString() string {
	return s.String()
}

func (s *RestartDBNodeZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartDBNodeZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBNodeZonalRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *RestartDBNodeZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestartDBNodeZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestartDBNodeZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestartDBNodeZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestartDBNodeZonalRequest) SetClientToken(v string) *RestartDBNodeZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetDBClusterId(v string) *RestartDBNodeZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetDBNodeId(v string) *RestartDBNodeZonalRequest {
	s.DBNodeId = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetOwnerAccount(v string) *RestartDBNodeZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetOwnerId(v int64) *RestartDBNodeZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetResourceOwnerAccount(v string) *RestartDBNodeZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBNodeZonalRequest) SetResourceOwnerId(v int64) *RestartDBNodeZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBNodeZonalRequest) Validate() error {
	return dara.Validate(s)
}
