// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachExpressConnectRouterChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildInstanceId(v string) *DetachExpressConnectRouterChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceType(v string) *DetachExpressConnectRouterChildInstanceRequest
	GetChildInstanceType() *string
	SetClientToken(v string) *DetachExpressConnectRouterChildInstanceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DetachExpressConnectRouterChildInstanceRequest
	GetDryRun() *bool
	SetEcrId(v string) *DetachExpressConnectRouterChildInstanceRequest
	GetEcrId() *string
	SetVersion(v string) *DetachExpressConnectRouterChildInstanceRequest
	GetVersion() *string
}

type DetachExpressConnectRouterChildInstanceRequest struct {
	// The VBR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-j6cwxhgg0s5nuephp****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The type of the network instance. Set the value to **VBR**.
	//
	// This parameter is required.
	//
	// example:
	//
	// VBR
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
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
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DetachExpressConnectRouterChildInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DetachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetVersion(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.Version = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}
