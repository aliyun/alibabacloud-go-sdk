// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ValidateApplicationRequest
	GetApplicationId() *string
	SetClientToken(v string) *ValidateApplicationRequest
	GetClientToken() *string
	SetResourceGroupId(v string) *ValidateApplicationRequest
	GetResourceGroupId() *string
}

type ValidateApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
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

func (s ValidateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateApplicationRequest) GoString() string {
	return s.String()
}

func (s *ValidateApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ValidateApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ValidateApplicationRequest) SetApplicationId(v string) *ValidateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ValidateApplicationRequest) SetClientToken(v string) *ValidateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateApplicationRequest) SetResourceGroupId(v string) *ValidateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValidateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
