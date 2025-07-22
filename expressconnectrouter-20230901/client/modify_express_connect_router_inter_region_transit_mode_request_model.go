// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterInterRegionTransitModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyExpressConnectRouterInterRegionTransitModeRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest
	GetEcrId() *string
	SetTransitModeList(v []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) *ModifyExpressConnectRouterInterRegionTransitModeRequest
	GetTransitModeList() []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList
}

type ModifyExpressConnectRouterInterRegionTransitModeRequest struct {
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
	// The cross-region forwarding modes.
	TransitModeList []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList `json:"TransitModeList,omitempty" xml:"TransitModeList,omitempty" type:"Repeated"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) GetTransitModeList() []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	return s.TransitModeList
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetTransitModeList(v []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.TransitModeList = v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList struct {
	// The cross-domain forwarding mode of the ECR. Valid values:
	//
	// 	- **ECMP**: the load balancing mode.
	//
	// 	- **NearBy**: the nearby forwarding mode.
	//
	// example:
	//
	// ECMP
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the ECR.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) GetMode() *string {
	return s.Mode
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetMode(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.Mode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetRegionId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) Validate() error {
	return dara.Validate(s)
}
