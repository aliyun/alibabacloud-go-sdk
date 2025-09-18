// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *DeleteExpressConnectRouterRequest
	GetEcrId() *string
	SetVersion(v string) *DeleteExpressConnectRouterRequest
	GetVersion() *string
}

type DeleteExpressConnectRouterRequest struct {
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
	// false
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

func (s DeleteExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DeleteExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteExpressConnectRouterRequest) SetClientToken(v string) *DeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetDryRun(v bool) *DeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetEcrId(v string) *DeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetVersion(v string) *DeleteExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
