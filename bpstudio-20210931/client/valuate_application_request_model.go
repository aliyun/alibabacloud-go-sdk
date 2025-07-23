// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ValuateApplicationRequest
	GetApplicationId() *string
	SetClientToken(v string) *ValuateApplicationRequest
	GetClientToken() *string
	SetResourceGroupId(v string) *ValuateApplicationRequest
	GetResourceGroupId() *string
}

type ValuateApplicationRequest struct {
	// The operation that you want to perform. Set the value to ValuateApplication.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group to which the application you want to query belongs.
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

func (s ValuateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ValuateApplicationRequest) GoString() string {
	return s.String()
}

func (s *ValuateApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ValuateApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValuateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ValuateApplicationRequest) SetApplicationId(v string) *ValuateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ValuateApplicationRequest) SetClientToken(v string) *ValuateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateApplicationRequest) SetResourceGroupId(v string) *ValuateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValuateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
