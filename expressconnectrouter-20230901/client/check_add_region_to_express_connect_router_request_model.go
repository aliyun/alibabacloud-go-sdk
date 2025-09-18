// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAddRegionToExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CheckAddRegionToExpressConnectRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *CheckAddRegionToExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *CheckAddRegionToExpressConnectRouterRequest
	GetEcrId() *string
	SetFreshRegionId(v string) *CheckAddRegionToExpressConnectRouterRequest
	GetFreshRegionId() *string
	SetVersion(v string) *CheckAddRegionToExpressConnectRouterRequest
	GetVersion() *string
}

type CheckAddRegionToExpressConnectRouterRequest struct {
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
	// 	- true: performs only a dry run.
	//
	// 	- false (default): performs a dry run and performs the actual request.
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
	// The ID of the region for which you want to check whether the CDT service is enabled for the ECR feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	FreshRegionId *string `json:"FreshRegionId,omitempty" xml:"FreshRegionId,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckAddRegionToExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CheckAddRegionToExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *CheckAddRegionToExpressConnectRouterRequest) GetFreshRegionId() *string {
	return s.FreshRegionId
}

func (s *CheckAddRegionToExpressConnectRouterRequest) GetVersion() *string {
	return s.Version
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetClientToken(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetDryRun(v bool) *CheckAddRegionToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetEcrId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetFreshRegionId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.FreshRegionId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetVersion(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.Version = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
