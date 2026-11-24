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
	// A client token to ensure that the request is idempotent.
	//
	// Generate a unique parameter value from your client. The ClientToken parameter can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the request as the **ClientToken**. Each API request has a unique **RequestId**.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run to check the request. No resources are modified.
	//
	// - **false*	- (default): Sends a standard request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ECR instance.
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
