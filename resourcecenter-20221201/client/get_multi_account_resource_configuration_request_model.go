// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetMultiAccountResourceConfigurationRequest
	GetAccountId() *string
	SetResourceId(v string) *GetMultiAccountResourceConfigurationRequest
	GetResourceId() *string
	SetResourceRegionId(v string) *GetMultiAccountResourceConfigurationRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *GetMultiAccountResourceConfigurationRequest
	GetResourceType() *string
}

type GetMultiAccountResourceConfigurationRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMultiAccountResourceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiAccountResourceConfigurationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetMultiAccountResourceConfigurationRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *GetMultiAccountResourceConfigurationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetMultiAccountResourceConfigurationRequest) SetAccountId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceRegionId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceType(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
