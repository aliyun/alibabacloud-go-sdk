// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeExpressConnectRouterInterRegionBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest
	GetClientToken() *string
	SetDryRun(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest
	GetDryRun() *bool
	SetEcrId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest
	GetEcrId() *string
}

type SynchronizeExpressConnectRouterInterRegionBandwidthRequest struct {
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
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetClientToken(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetDryRun(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetEcrId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.EcrId = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
