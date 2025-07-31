// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *GetServiceConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetServiceConfigRequest
	GetServiceCode() *string
}

type GetServiceConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetServiceConfigRequest) SetRegionId(v string) *GetServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceConfigRequest) SetResourceType(v string) *GetServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceConfigRequest) SetServiceCode(v string) *GetServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
