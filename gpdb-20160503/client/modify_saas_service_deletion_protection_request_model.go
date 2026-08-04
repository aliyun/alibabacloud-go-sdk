// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySaasServiceDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *ModifySaasServiceDeletionProtectionRequest
	GetDeletionProtection() *bool
	SetRegionId(v string) *ModifySaasServiceDeletionProtectionRequest
	GetRegionId() *string
	SetServiceId(v string) *ModifySaasServiceDeletionProtectionRequest
	GetServiceId() *string
}

type ModifySaasServiceDeletionProtectionRequest struct {
	// Specifies whether to enable deletion protection.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
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

func (s ModifySaasServiceDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySaasServiceDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifySaasServiceDeletionProtectionRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifySaasServiceDeletionProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySaasServiceDeletionProtectionRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ModifySaasServiceDeletionProtectionRequest) SetDeletionProtection(v bool) *ModifySaasServiceDeletionProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionRequest) SetRegionId(v string) *ModifySaasServiceDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionRequest) SetServiceId(v string) *ModifySaasServiceDeletionProtectionRequest {
	s.ServiceId = &v
	return s
}

func (s *ModifySaasServiceDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
