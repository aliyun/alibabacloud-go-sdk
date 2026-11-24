// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectRouterRegionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterRegionRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterRegionRequest
	GetEcrId() *string
	SetVersion(v string) *DescribeExpressConnectRouterRegionRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterRegionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the token on your client to ensure that it is unique among different requests. The client token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether this request is a dry run. Valid values:
	//
	// - **true**: sends a check request and does not modify the EPS resource.
	//
	// - **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-fu8rszhgv7623c****
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeExpressConnectRouterRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterRegionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterRegionRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterRegionRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeExpressConnectRouterRegionRequest) SetClientToken(v string) *DescribeExpressConnectRouterRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRegionRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetEcrId(v string) *DescribeExpressConnectRouterRegionRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetVersion(v string) *DescribeExpressConnectRouterRegionRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) Validate() error {
	return dara.Validate(s)
}
