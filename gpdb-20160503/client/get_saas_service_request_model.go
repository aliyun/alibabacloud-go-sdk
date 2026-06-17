// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSaasServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *GetSaasServiceRequest
	GetServiceId() *string
}

type GetSaasServiceRequest struct {
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

func (s GetSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *GetSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSaasServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasServiceRequest) SetRegionId(v string) *GetSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetSaasServiceRequest) SetServiceId(v string) *GetSaasServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
