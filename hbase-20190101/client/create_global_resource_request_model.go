// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateGlobalResourceRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateGlobalResourceRequest
	GetClusterId() *string
	SetRegionId(v string) *CreateGlobalResourceRequest
	GetRegionId() *string
	SetResourceName(v string) *CreateGlobalResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *CreateGlobalResourceRequest
	GetResourceType() *string
}

type CreateGlobalResourceRequest struct {
	// example:
	//
	// xxxxx-xxxxx-xxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PubPhoenixSLBQueryServerVip
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GLOBAL_VIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateGlobalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGlobalResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateGlobalResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGlobalResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateGlobalResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateGlobalResourceRequest) SetClientToken(v string) *CreateGlobalResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetClusterId(v string) *CreateGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetRegionId(v string) *CreateGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceName(v string) *CreateGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceType(v string) *CreateGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateGlobalResourceRequest) Validate() error {
	return dara.Validate(s)
}
