// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSaasServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckSaasServiceVersionRequest
	GetRegionId() *string
	SetServiceId(v string) *CheckSaasServiceVersionRequest
	GetServiceId() *string
}

type CheckSaasServiceVersionRequest struct {
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
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CheckSaasServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSaasServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *CheckSaasServiceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckSaasServiceVersionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CheckSaasServiceVersionRequest) SetRegionId(v string) *CheckSaasServiceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *CheckSaasServiceVersionRequest) SetServiceId(v string) *CheckSaasServiceVersionRequest {
	s.ServiceId = &v
	return s
}

func (s *CheckSaasServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
