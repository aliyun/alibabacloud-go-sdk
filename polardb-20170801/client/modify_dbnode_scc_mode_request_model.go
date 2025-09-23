// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeSccModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBNodeSccModeRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *ModifyDBNodeSccModeRequest
	GetDBNodeId() *string
	SetOwnerAccount(v string) *ModifyDBNodeSccModeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeSccModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBNodeSccModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeSccModeRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *ModifyDBNodeSccModeRequest
	GetSccMode() *string
}

type ModifyDBNodeSccModeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pi-****************
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// SCC
	//
	// This parameter is required.
	//
	// example:
	//
	// OFF
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s ModifyDBNodeSccModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeSccModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeSccModeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeSccModeRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *ModifyDBNodeSccModeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeSccModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeSccModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeSccModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeSccModeRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *ModifyDBNodeSccModeRequest) SetDBClusterId(v string) *ModifyDBNodeSccModeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetDBNodeId(v string) *ModifyDBNodeSccModeRequest {
	s.DBNodeId = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetOwnerAccount(v string) *ModifyDBNodeSccModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetOwnerId(v int64) *ModifyDBNodeSccModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeSccModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetResourceOwnerId(v int64) *ModifyDBNodeSccModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) SetSccMode(v string) *ModifyDBNodeSccModeRequest {
	s.SccMode = &v
	return s
}

func (s *ModifyDBNodeSccModeRequest) Validate() error {
	return dara.Validate(s)
}
