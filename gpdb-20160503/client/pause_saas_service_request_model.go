// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *PauseSaasServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *PauseSaasServiceRequest
	GetServiceId() *string
}

type PauseSaasServiceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s PauseSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *PauseSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PauseSaasServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *PauseSaasServiceRequest) SetRegionId(v string) *PauseSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *PauseSaasServiceRequest) SetServiceId(v string) *PauseSaasServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *PauseSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
