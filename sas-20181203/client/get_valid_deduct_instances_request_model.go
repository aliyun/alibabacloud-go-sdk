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
	// Resource package instance ID, can be queried through [QueryResourcePackageInstances]().
	//
	// example:
	//
	// sas_cspm_dp_cn-***80001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Resource package name code, values:
	//
	// - Vulnerability resource package: **sas_vul_dp_cn**
	//
	// - CSPM resource package: **sas_cspm_dp_cn**
	//
	// - Anti-virus resource package: **sas_viruspackage_dp_cn**
	//
	// example:
	//
	// sas_vul_dp_cn
	Modules *string `json:"Modules,omitempty" xml:"Modules,omitempty"`
	// Resource package status, default is valid (valid), not modifiable.
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
