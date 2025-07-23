// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeployApplicationRequest
	GetApplicationId() *string
	SetClientToken(v string) *DeployApplicationRequest
	GetClientToken() *string
	SetResourceGroupId(v string) *DeployApplicationRequest
	GetResourceGroupId() *string
}

type DeployApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VVK605ZH00OA4MRT
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeployApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeployApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeployApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeployApplicationRequest) SetApplicationId(v string) *DeployApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeployApplicationRequest) SetClientToken(v string) *DeployApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployApplicationRequest) SetResourceGroupId(v string) *DeployApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeployApplicationRequest) Validate() error {
	return dara.Validate(s)
}
