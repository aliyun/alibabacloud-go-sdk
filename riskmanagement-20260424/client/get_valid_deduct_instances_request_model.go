// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidDeductInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetValidDeductInstancesRequest
	GetRegionId() *string
	SetSdkRequest(v *GetValidDeductInstancesRequestSdkRequest) *GetValidDeductInstancesRequest
	GetSdkRequest() *GetValidDeductInstancesRequestSdkRequest
}

type GetValidDeductInstancesRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-fuzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request parameters.
	SdkRequest *GetValidDeductInstancesRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s GetValidDeductInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetValidDeductInstancesRequest) GetSdkRequest() *GetValidDeductInstancesRequestSdkRequest {
	return s.SdkRequest
}

func (s *GetValidDeductInstancesRequest) SetRegionId(v string) *GetValidDeductInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *GetValidDeductInstancesRequest) SetSdkRequest(v *GetValidDeductInstancesRequestSdkRequest) *GetValidDeductInstancesRequest {
	s.SdkRequest = v
	return s
}

func (s *GetValidDeductInstancesRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetValidDeductInstancesRequestSdkRequest struct {
	// The resource plan instance ID. You can call QueryResourcePackageInstances to query the ID.
	//
	// example:
	//
	// sas_cspm_dp_cn-***80001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource plan name code. Valid values:
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
	// The resource plan status. Default value: valid. This parameter cannot be modified.
	//
	// example:
	//
	// Available
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetValidDeductInstancesRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetValidDeductInstancesRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *GetValidDeductInstancesRequestSdkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetValidDeductInstancesRequestSdkRequest) GetModules() *string {
	return s.Modules
}

func (s *GetValidDeductInstancesRequestSdkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetValidDeductInstancesRequestSdkRequest) SetInstanceId(v string) *GetValidDeductInstancesRequestSdkRequest {
	s.InstanceId = &v
	return s
}

func (s *GetValidDeductInstancesRequestSdkRequest) SetModules(v string) *GetValidDeductInstancesRequestSdkRequest {
	s.Modules = &v
	return s
}

func (s *GetValidDeductInstancesRequestSdkRequest) SetStatus(v int32) *GetValidDeductInstancesRequestSdkRequest {
	s.Status = &v
	return s
}

func (s *GetValidDeductInstancesRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}
