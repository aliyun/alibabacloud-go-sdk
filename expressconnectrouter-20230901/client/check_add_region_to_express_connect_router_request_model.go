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
	// You can generate the parameter value on your client and ensure that the value is unique across different requests. The ClientToken parameter only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on this request. Valid values:
	//
	// - true: sends a check request and does not modify EPS resources.
	//
	// - false: sends a normal request. This is the default value.
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
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The region for which you want to check whether the Cloud Data Transfer (CDT) service is enabled for the Express Connect Router (ECR) feature.
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
