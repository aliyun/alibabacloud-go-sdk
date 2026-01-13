// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleMetaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v int32) *CreateCustomerModuleMetaInfoRequest
	GetCustomerModuleId() *int32
	SetFeatureString(v string) *CreateCustomerModuleMetaInfoRequest
	GetFeatureString() *string
	SetFeatureTemplate(v string) *CreateCustomerModuleMetaInfoRequest
	GetFeatureTemplate() *string
}

type CreateCustomerModuleMetaInfoRequest struct {
	// Customer model ID.
	//
	// example:
	//
	// 1
	CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
	// A string representation of List<Object>, where Object has the following structure:
	//
	// {
	//
	// "mappingName": "xxx"
	//
	// }
	//
	// example:
	//
	// {\\"mappingName\\":\\"f1\\",\\"fieldName\\":\\"saf_shanghai_new.f1\\",\\"valueType\\":\\"DOUBLE\\",\\"defaultValue\\":\\"\\"}
	FeatureString *string `json:"FeatureString,omitempty" xml:"FeatureString,omitempty"`
	// Version of the feature adopted.
	//
	// example:
	//
	// FINANCE_51
	FeatureTemplate *string `json:"FeatureTemplate,omitempty" xml:"FeatureTemplate,omitempty"`
}

func (s CreateCustomerModuleMetaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleMetaInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleMetaInfoRequest) GetCustomerModuleId() *int32 {
	return s.CustomerModuleId
}

func (s *CreateCustomerModuleMetaInfoRequest) GetFeatureString() *string {
	return s.FeatureString
}

func (s *CreateCustomerModuleMetaInfoRequest) GetFeatureTemplate() *string {
	return s.FeatureTemplate
}

func (s *CreateCustomerModuleMetaInfoRequest) SetCustomerModuleId(v int32) *CreateCustomerModuleMetaInfoRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoRequest) SetFeatureString(v string) *CreateCustomerModuleMetaInfoRequest {
	s.FeatureString = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoRequest) SetFeatureTemplate(v string) *CreateCustomerModuleMetaInfoRequest {
	s.FeatureTemplate = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoRequest) Validate() error {
	return dara.Validate(s)
}
