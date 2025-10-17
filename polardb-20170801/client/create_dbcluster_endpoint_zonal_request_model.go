// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *CreateDBClusterEndpointZonalRequest
	GetAutoAddNewNodes() *string
	SetClientToken(v string) *CreateDBClusterEndpointZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateDBClusterEndpointZonalRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *CreateDBClusterEndpointZonalRequest
	GetDBEndpointDescription() *string
	SetEndpointConfig(v string) *CreateDBClusterEndpointZonalRequest
	GetEndpointConfig() *string
	SetEndpointType(v string) *CreateDBClusterEndpointZonalRequest
	GetEndpointType() *string
	SetNodes(v string) *CreateDBClusterEndpointZonalRequest
	GetNodes() *string
	SetOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBClusterEndpointZonalRequest
	GetOwnerId() *int64
	SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointZonalRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointZonalRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *CreateDBClusterEndpointZonalRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBClusterEndpointZonalRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *CreateDBClusterEndpointZonalRequest
	GetSccMode() *string
}

type CreateDBClusterEndpointZonalRequest struct {
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// example:
	//
	// {"ConsistLevel": "1","DistributedTransaction": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
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
	// ReadOnly
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s CreateDBClusterEndpointZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointZonalRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointZonalRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *CreateDBClusterEndpointZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBClusterEndpointZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBClusterEndpointZonalRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *CreateDBClusterEndpointZonalRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *CreateDBClusterEndpointZonalRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateDBClusterEndpointZonalRequest) GetNodes() *string {
	return s.Nodes
}

func (s *CreateDBClusterEndpointZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBClusterEndpointZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBClusterEndpointZonalRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *CreateDBClusterEndpointZonalRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *CreateDBClusterEndpointZonalRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *CreateDBClusterEndpointZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBClusterEndpointZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterEndpointZonalRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *CreateDBClusterEndpointZonalRequest) SetAutoAddNewNodes(v string) *CreateDBClusterEndpointZonalRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetClientToken(v string) *CreateDBClusterEndpointZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetDBClusterId(v string) *CreateDBClusterEndpointZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetDBEndpointDescription(v string) *CreateDBClusterEndpointZonalRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetEndpointConfig(v string) *CreateDBClusterEndpointZonalRequest {
	s.EndpointConfig = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetEndpointType(v string) *CreateDBClusterEndpointZonalRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetNodes(v string) *CreateDBClusterEndpointZonalRequest {
	s.Nodes = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetOwnerId(v int64) *CreateDBClusterEndpointZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointZonalRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointZonalRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetReadWriteMode(v string) *CreateDBClusterEndpointZonalRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetResourceOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetResourceOwnerId(v int64) *CreateDBClusterEndpointZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetSccMode(v string) *CreateDBClusterEndpointZonalRequest {
	s.SccMode = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) Validate() error {
	return dara.Validate(s)
}
