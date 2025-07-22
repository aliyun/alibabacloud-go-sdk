// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GrantInstanceToExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *GrantInstanceToExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *GrantInstanceToExpressConnectRouterRequest
	GetEcrId() *string
	SetEcrOwnerAliUid(v int64) *GrantInstanceToExpressConnectRouterRequest
	GetEcrOwnerAliUid() *int64
	SetInstanceId(v string) *GrantInstanceToExpressConnectRouterRequest
	GetInstanceId() *string
	SetInstanceRegionId(v string) *GrantInstanceToExpressConnectRouterRequest
	GetInstanceRegionId() *string
	SetInstanceType(v string) *GrantInstanceToExpressConnectRouterRequest
	GetInstanceType() *string
}

type GrantInstanceToExpressConnectRouterRequest struct {
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
	// The ID of the Alibaba Cloud account that owns the ECR to which you want to grant permissions.
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
}

func (s GrantInstanceToExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetEcrOwnerAliUid() *int64 {
	return s.EcrOwnerAliUid
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *GrantInstanceToExpressConnectRouterRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetClientToken(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetDryRun(v bool) *GrantInstanceToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceRegionId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceType(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
