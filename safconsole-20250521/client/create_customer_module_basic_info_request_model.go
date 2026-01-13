// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleName(v string) *CreateCustomerModuleBasicInfoRequest
	GetCustomerModuleName() *string
	SetDescription(v string) *CreateCustomerModuleBasicInfoRequest
	GetDescription() *string
	SetModuleType(v string) *CreateCustomerModuleBasicInfoRequest
	GetModuleType() *string
	SetStorePath(v string) *CreateCustomerModuleBasicInfoRequest
	GetStorePath() *string
}

type CreateCustomerModuleBasicInfoRequest struct {
	// The display identifier of the model for customers.
	//
	// example:
	//
	// module01
	CustomerModuleName *string `json:"CustomerModuleName,omitempty" xml:"CustomerModuleName,omitempty"`
	// Model description.
	//
	// example:
	//
	// 模型描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Model deployment type.
	//
	// example:
	//
	// PMML
	ModuleType *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	// Model storage path.
	//
	// example:
	//
	// customer/xxxxxxxxx/xxxxxxxx.pmml
	StorePath *string `json:"StorePath,omitempty" xml:"StorePath,omitempty"`
}

func (s CreateCustomerModuleBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleBasicInfoRequest) GetCustomerModuleName() *string {
	return s.CustomerModuleName
}

func (s *CreateCustomerModuleBasicInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomerModuleBasicInfoRequest) GetModuleType() *string {
	return s.ModuleType
}

func (s *CreateCustomerModuleBasicInfoRequest) GetStorePath() *string {
	return s.StorePath
}

func (s *CreateCustomerModuleBasicInfoRequest) SetCustomerModuleName(v string) *CreateCustomerModuleBasicInfoRequest {
	s.CustomerModuleName = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoRequest) SetDescription(v string) *CreateCustomerModuleBasicInfoRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoRequest) SetModuleType(v string) *CreateCustomerModuleBasicInfoRequest {
	s.ModuleType = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoRequest) SetStorePath(v string) *CreateCustomerModuleBasicInfoRequest {
	s.StorePath = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
