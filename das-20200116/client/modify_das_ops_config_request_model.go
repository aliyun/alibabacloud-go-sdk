// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDasOpsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifyDasOpsConfigRequest
	GetEnable() *bool
	SetFilters(v []*ModifyDasOpsConfigRequestFilters) *ModifyDasOpsConfigRequest
	GetFilters() []*ModifyDasOpsConfigRequestFilters
	SetInstanceId(v string) *ModifyDasOpsConfigRequest
	GetInstanceId() *string
}

type ModifyDasOpsConfigRequest struct {
	// Specifies whether to enable DAS Alibaba Cloud Managed Services. Valid values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable. Shutdown only turns off the feature but does not unsubscribe from the service. To unsubscribe, go to the unsubscription management page.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// A reserved parameter.
	Filters []*ModifyDasOpsConfigRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-uf6079bda570****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyDasOpsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasOpsConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDasOpsConfigRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyDasOpsConfigRequest) GetFilters() []*ModifyDasOpsConfigRequestFilters {
	return s.Filters
}

func (s *ModifyDasOpsConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDasOpsConfigRequest) SetEnable(v bool) *ModifyDasOpsConfigRequest {
	s.Enable = &v
	return s
}

func (s *ModifyDasOpsConfigRequest) SetFilters(v []*ModifyDasOpsConfigRequestFilters) *ModifyDasOpsConfigRequest {
	s.Filters = v
	return s
}

func (s *ModifyDasOpsConfigRequest) SetInstanceId(v string) *ModifyDasOpsConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDasOpsConfigRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDasOpsConfigRequestFilters struct {
	// The filter parameter.
	//
	// > For supported filter parameters and their values, see **Supplementary description of request parameters**.
	//
	// example:
	//
	// None
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// select
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDasOpsConfigRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ModifyDasOpsConfigRequestFilters) GoString() string {
	return s.String()
}

func (s *ModifyDasOpsConfigRequestFilters) GetKey() *string {
	return s.Key
}

func (s *ModifyDasOpsConfigRequestFilters) GetValue() *string {
	return s.Value
}

func (s *ModifyDasOpsConfigRequestFilters) SetKey(v string) *ModifyDasOpsConfigRequestFilters {
	s.Key = &v
	return s
}

func (s *ModifyDasOpsConfigRequestFilters) SetValue(v string) *ModifyDasOpsConfigRequestFilters {
	s.Value = &v
	return s
}

func (s *ModifyDasOpsConfigRequestFilters) Validate() error {
	return dara.Validate(s)
}
