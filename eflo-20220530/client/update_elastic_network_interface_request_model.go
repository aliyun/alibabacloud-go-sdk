// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateElasticNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateElasticNetworkInterfaceRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateElasticNetworkInterfaceRequest
	GetDescription() *string
	SetElasticNetworkInterfaceId(v string) *UpdateElasticNetworkInterfaceRequest
	GetElasticNetworkInterfaceId() *string
	SetRegionId(v string) *UpdateElasticNetworkInterfaceRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *UpdateElasticNetworkInterfaceRequest
	GetSecurityGroupId() *string
}

type UpdateElasticNetworkInterfaceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 3fd79d62-ab1d-11ec-9a53-0242ac110004
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the variable.
	//
	// example:
	//
	// LHICDOSEExternaluserinquiryP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Lingjun Elastic Network Interface ID
	//
	// This parameter is required.
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-wz9fj2s3o21nw2****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UpdateElasticNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateElasticNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateElasticNetworkInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateElasticNetworkInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateElasticNetworkInterfaceRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *UpdateElasticNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateElasticNetworkInterfaceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateElasticNetworkInterfaceRequest) SetClientToken(v string) *UpdateElasticNetworkInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetDescription(v string) *UpdateElasticNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetElasticNetworkInterfaceId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetRegionId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) SetSecurityGroupId(v string) *UpdateElasticNetworkInterfaceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateElasticNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
