// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidDeductInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetValidDeductInstancesRequest
	GetInstanceId() *string
	SetModules(v string) *GetValidDeductInstancesRequest
	GetModules() *string
	SetStatus(v int32) *GetValidDeductInstancesRequest
	GetStatus() *int32
}

type GetValidDeductInstancesRequest struct {
	// Instance ID of the resource plan instance. You can call [QueryResourcePackageInstances]() to query instance ID.
	//
	// example:
	//
	// sas_cspm_dp_cn-***80001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The code of the resource plan. Valid values:
	//
	// - Vulnerability resource plan: **sas_vul_dp_cn**
	//
	// - CSPM resource plan: **sas_cspm_dp_cn**
	//
	// - Anti-virus Edition resource plan: **sas_viruspackage_dp_cn**
	//
	// example:
	//
	// sas_vul_dp_cn
	Modules *string `json:"Modules,omitempty" xml:"Modules,omitempty"`
	// The status of the resource plan. The default value is valid. This parameter does not support modification.
	//
	// example:
	//
	// Available
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetValidDeductInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetValidDeductInstancesRequest) GetModules() *string {
	return s.Modules
}

func (s *GetValidDeductInstancesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetValidDeductInstancesRequest) SetInstanceId(v string) *GetValidDeductInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetValidDeductInstancesRequest) SetModules(v string) *GetValidDeductInstancesRequest {
	s.Modules = &v
	return s
}

func (s *GetValidDeductInstancesRequest) SetStatus(v int32) *GetValidDeductInstancesRequest {
	s.Status = &v
	return s
}

func (s *GetValidDeductInstancesRequest) Validate() error {
	return dara.Validate(s)
}
