// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *RevokeInstanceFromExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetEcrId() *string
	SetEcrOwnerAliUid(v int64) *RevokeInstanceFromExpressConnectRouterRequest
	GetEcrOwnerAliUid() *int64
	SetInstanceId(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetInstanceId() *string
	SetInstanceRegionId(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetInstanceRegionId() *string
	SetInstanceType(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetInstanceType() *string
	SetVersion(v string) *RevokeInstanceFromExpressConnectRouterRequest
	GetVersion() *string
}

type RevokeInstanceFromExpressConnectRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the ECR from which you want to revoke permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121012345612****
	EcrOwnerAliUid *int64 `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	// The ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VBR**
	//
	// 	- **VPC**
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetEcrOwnerAliUid() *int64 {
	return s.EcrOwnerAliUid
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetClientToken(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetDryRun(v bool) *RevokeInstanceFromExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceRegionId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceType(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetVersion(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
