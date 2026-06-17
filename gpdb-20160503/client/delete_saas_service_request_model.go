// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSaasServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DeleteSaasServiceRequest
	GetServiceId() *string
}

type DeleteSaasServiceRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DeleteSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSaasServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteSaasServiceRequest) SetRegionId(v string) *DeleteSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSaasServiceRequest) SetServiceId(v string) *DeleteSaasServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
