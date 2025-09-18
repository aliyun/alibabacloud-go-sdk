// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForceDeleteExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ForceDeleteExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *ForceDeleteExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *ForceDeleteExpressConnectRouterRequest
	GetEcrId() *string
	SetVersion(v string) *ForceDeleteExpressConnectRouterRequest
	GetVersion() *string
}

type ForceDeleteExpressConnectRouterRequest struct {
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
	// ecr-fu8rszhgv7623c****
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ForceDeleteExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ForceDeleteExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ForceDeleteExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ForceDeleteExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
}

func (s *ForceDeleteExpressConnectRouterRequest) SetClientToken(v string) *ForceDeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetDryRun(v bool) *ForceDeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetEcrId(v string) *ForceDeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetVersion(v string) *ForceDeleteExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
