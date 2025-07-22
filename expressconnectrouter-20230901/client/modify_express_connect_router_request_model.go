// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyExpressConnectRouterRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyExpressConnectRouterRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyExpressConnectRouterRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyExpressConnectRouterRequest
	GetEcrId() *string
	SetName(v string) *ModifyExpressConnectRouterRequest
	GetName() *string
}

type ModifyExpressConnectRouterRequest struct {
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
	// The description of the ECR.
	//
	// >  The description can be empty or 0 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The name of the ECR.
	//
	// >  The name must be 0 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyExpressConnectRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectRouterRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyExpressConnectRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyExpressConnectRouterRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyExpressConnectRouterRequest) GetName() *string {
	return s.Name
}

func (s *ModifyExpressConnectRouterRequest) SetClientToken(v string) *ModifyExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDescription(v string) *ModifyExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDryRun(v bool) *ModifyExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetEcrId(v string) *ModifyExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetName(v string) *ModifyExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) Validate() error {
	return dara.Validate(s)
}
