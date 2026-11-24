// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterInterRegionTransitModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest
	GetClientToken() *string
	SetDryRun(v bool) *DescribeExpressConnectRouterInterRegionTransitModeRequest
	GetDryRun() *bool
	SetEcrId(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest
	GetEcrId() *string
	SetVersion(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest
	GetVersion() *string
}

type DescribeExpressConnectRouterInterRegionTransitModeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run and does not modify the EPS resource.
	//
	// - **false**: sends a normal request. This is the default value.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetVersion(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.Version = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) Validate() error {
	return dara.Validate(s)
}
