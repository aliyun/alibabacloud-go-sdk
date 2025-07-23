// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ReleaseApplicationRequest
	GetApplicationId() *string
	SetClientToken(v string) *ReleaseApplicationRequest
	GetClientToken() *string
	SetResourceGroupId(v string) *ReleaseApplicationRequest
	GetResourceGroupId() *string
}

type ReleaseApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7QSXFQW46ZNGOUDM
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
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ReleaseApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseApplicationRequest) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ReleaseApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ReleaseApplicationRequest) SetApplicationId(v string) *ReleaseApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ReleaseApplicationRequest) SetClientToken(v string) *ReleaseApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseApplicationRequest) SetResourceGroupId(v string) *ReleaseApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ReleaseApplicationRequest) Validate() error {
	return dara.Validate(s)
}
