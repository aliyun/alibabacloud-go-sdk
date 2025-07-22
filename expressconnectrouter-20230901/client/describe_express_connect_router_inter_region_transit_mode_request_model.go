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
}

type DescribeExpressConnectRouterInterRegionTransitModeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// FF9nMec/RZ6H9oqFn1pvyir/SLRlxCCyHJonbGzqL01hiM6Jb3wJowdHvjCfog7ww1b9rSHMRFJnrUBfVba68TJg==
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

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) Validate() error {
	return dara.Validate(s)
}
