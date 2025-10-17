// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointZonalRequest
	GetAutoAddNewNodes() *string
	SetClientToken(v string) *ModifyDBClusterEndpointZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBEndpointDescription() *string
	SetDBEndpointId(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBEndpointId() *string
	SetEndpointConfig(v string) *ModifyDBClusterEndpointZonalRequest
	GetEndpointConfig() *string
	SetNodes(v string) *ModifyDBClusterEndpointZonalRequest
	GetNodes() *string
	SetOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest
	GetOwnerId() *int64
	SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointZonalRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointZonalRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *ModifyDBClusterEndpointZonalRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *ModifyDBClusterEndpointZonalRequest
	GetSccMode() *string
}

type ModifyDBClusterEndpointZonalRequest struct {
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// example:
	//
	// ReadWrite
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// OFF
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s ModifyDBClusterEndpointZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointZonalRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *ModifyDBClusterEndpointZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *ModifyDBClusterEndpointZonalRequest) GetNodes() *string {
	return s.Nodes
}

func (s *ModifyDBClusterEndpointZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterEndpointZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *ModifyDBClusterEndpointZonalRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *ModifyDBClusterEndpointZonalRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *ModifyDBClusterEndpointZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterEndpointZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *ModifyDBClusterEndpointZonalRequest) SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointZonalRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetClientToken(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBClusterId(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBEndpointDescription(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBEndpointId(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetEndpointConfig(v string) *ModifyDBClusterEndpointZonalRequest {
	s.EndpointConfig = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetNodes(v string) *ModifyDBClusterEndpointZonalRequest {
	s.Nodes = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointZonalRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointZonalRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetReadWriteMode(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetResourceOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetSccMode(v string) *ModifyDBClusterEndpointZonalRequest {
	s.SccMode = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) Validate() error {
	return dara.Validate(s)
}
