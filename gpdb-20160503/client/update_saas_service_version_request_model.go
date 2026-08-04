// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaasServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateSaasServiceVersionRequest
	GetRegionId() *string
	SetServiceId(v string) *UpdateSaasServiceVersionRequest
	GetServiceId() *string
}

type UpdateSaasServiceVersionRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agdb-xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateSaasServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaasServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSaasServiceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSaasServiceVersionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateSaasServiceVersionRequest) SetRegionId(v string) *UpdateSaasServiceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSaasServiceVersionRequest) SetServiceId(v string) *UpdateSaasServiceVersionRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateSaasServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
