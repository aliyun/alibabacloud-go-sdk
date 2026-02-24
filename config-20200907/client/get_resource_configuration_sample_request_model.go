// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMockOnly(v string) *GetResourceConfigurationSampleRequest
	GetMockOnly() *string
	SetResourceType(v string) *GetResourceConfigurationSampleRequest
	GetResourceType() *string
}

type GetResourceConfigurationSampleRequest struct {
	// Specifies whether to obtain only mock data, which is not real resource data.
	//
	// Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	MockOnly *string `json:"MockOnly,omitempty" xml:"MockOnly,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the resource type, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceConfigurationSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationSampleRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationSampleRequest) GetMockOnly() *string {
	return s.MockOnly
}

func (s *GetResourceConfigurationSampleRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationSampleRequest) SetMockOnly(v string) *GetResourceConfigurationSampleRequest {
	s.MockOnly = &v
	return s
}

func (s *GetResourceConfigurationSampleRequest) SetResourceType(v string) *GetResourceConfigurationSampleRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationSampleRequest) Validate() error {
	return dara.Validate(s)
}
