// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteGlobalResourceRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteGlobalResourceRequest
	GetRegionId() *string
	SetResourceName(v string) *DeleteGlobalResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *DeleteGlobalResourceRequest
	GetResourceType() *string
}

type DeleteGlobalResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
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

func (s DeleteGlobalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteGlobalResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGlobalResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *DeleteGlobalResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteGlobalResourceRequest) SetClusterId(v string) *DeleteGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetRegionId(v string) *DeleteGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceName(v string) *DeleteGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteGlobalResourceRequest) SetResourceType(v string) *DeleteGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteGlobalResourceRequest) Validate() error {
	return dara.Validate(s)
}
