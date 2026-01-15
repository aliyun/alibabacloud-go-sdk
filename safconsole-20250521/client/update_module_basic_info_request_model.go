// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModuleBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *UpdateModuleBasicInfoRequest
	GetCustomerModuleId() *int32
	SetCustomerModuleName(v string) *UpdateModuleBasicInfoRequest
	GetCustomerModuleName() *string
	SetDescription(v string) *UpdateModuleBasicInfoRequest
	GetDescription() *string
	SetModuleType(v string) *UpdateModuleBasicInfoRequest
	GetModuleType() *string
	SetStorePath(v string) *UpdateModuleBasicInfoRequest
	GetStorePath() *string
}

type UpdateModuleBasicInfoRequest struct {
	// Customer model ID
	//
	// example:
	//
	// 456
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
	// Model name.
	//
	// example:
	//
	// moduleName
	CustomerModuleName *string `json:"CustomerModuleName,omitempty" xml:"CustomerModuleName,omitempty"`
	// Description.
	//
	// example:
	//
	// 备注
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Module type.
	//
	// example:
	//
	// PMML
	ModuleType *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	// Test address.
	//
	// example:
	//
	// saf/xxxxx/xxx.pmml
	StorePath *string `json:"StorePath,omitempty" xml:"StorePath,omitempty"`
}

func (s UpdateModuleBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModuleBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateModuleBasicInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *UpdateModuleBasicInfoRequest) GetCustomerModuleName() *string {
	return s.CustomerModuleName
}

func (s *UpdateModuleBasicInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModuleBasicInfoRequest) GetModuleType() *string {
	return s.ModuleType
}

func (s *UpdateModuleBasicInfoRequest) GetStorePath() *string {
	return s.StorePath
}

func (s *UpdateModuleBasicInfoRequest) SetCustomerModuleId(v int32) *UpdateModuleBasicInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *UpdateModuleBasicInfoRequest) SetCustomerModuleName(v string) *UpdateModuleBasicInfoRequest {
	s.CustomerModuleName = &v
	return s
}

func (s *UpdateModuleBasicInfoRequest) SetDescription(v string) *UpdateModuleBasicInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateModuleBasicInfoRequest) SetModuleType(v string) *UpdateModuleBasicInfoRequest {
	s.ModuleType = &v
	return s
}

func (s *UpdateModuleBasicInfoRequest) SetStorePath(v string) *UpdateModuleBasicInfoRequest {
	s.StorePath = &v
	return s
}

func (s *UpdateModuleBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
