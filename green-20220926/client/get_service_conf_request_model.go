// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByDefault(v bool) *GetServiceConfRequest
	GetByDefault() *bool
	SetRegionId(v string) *GetServiceConfRequest
	GetRegionId() *string
	SetResourceType(v string) *GetServiceConfRequest
	GetResourceType() *string
	SetScene(v string) *GetServiceConfRequest
	GetScene() *string
	SetServiceCode(v string) *GetServiceConfRequest
	GetServiceCode() *string
}

type GetServiceConfRequest struct {
	// Query default configuration
	//
	// example:
	//
	// False
	ByDefault *bool `json:"ByDefault,omitempty" xml:"ByDefault,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Audit scenario.
	//
	// example:
	//
	// pornographic
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceConfRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConfRequest) GetByDefault() *bool {
	return s.ByDefault
}

func (s *GetServiceConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceConfRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceConfRequest) GetScene() *string {
	return s.Scene
}

func (s *GetServiceConfRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetServiceConfRequest) SetByDefault(v bool) *GetServiceConfRequest {
	s.ByDefault = &v
	return s
}

func (s *GetServiceConfRequest) SetRegionId(v string) *GetServiceConfRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceConfRequest) SetResourceType(v string) *GetServiceConfRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfRequest) SetScene(v string) *GetServiceConfRequest {
	s.Scene = &v
	return s
}

func (s *GetServiceConfRequest) SetServiceCode(v string) *GetServiceConfRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfRequest) Validate() error {
	return dara.Validate(s)
}
