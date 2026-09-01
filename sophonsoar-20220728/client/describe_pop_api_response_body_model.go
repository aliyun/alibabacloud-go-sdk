// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePopApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribePopApiResponseBody
	GetApiName() *string
	SetOpenApiMetaList(v []*DescribePopApiResponseBodyOpenApiMetaList) *DescribePopApiResponseBody
	GetOpenApiMetaList() []*DescribePopApiResponseBodyOpenApiMetaList
	SetPopCode(v string) *DescribePopApiResponseBody
	GetPopCode() *string
	SetRequestId(v string) *DescribePopApiResponseBody
	GetRequestId() *string
	SetVersion(v string) *DescribePopApiResponseBody
	GetVersion() *string
}

type DescribePopApiResponseBody struct {
	// The name of the API.
	//
	// example:
	//
	// AddAssetCleanConfig
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// List of API information.
	OpenApiMetaList []*DescribePopApiResponseBodyOpenApiMetaList `json:"OpenApiMetaList,omitempty" xml:"OpenApiMetaList,omitempty" type:"Repeated"`
	// The POP code of the Alibaba Cloud product API.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// The ID of the request. Alibaba Cloud generates this unique ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 1A01B0BA-****-5813-****-A5DA15FA95AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version number of the API.
	//
	// example:
	//
	// 2019-09-10
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePopApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePopApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBody) GetApiName() *string {
	return s.ApiName
}

func (s *DescribePopApiResponseBody) GetOpenApiMetaList() []*DescribePopApiResponseBodyOpenApiMetaList {
	return s.OpenApiMetaList
}

func (s *DescribePopApiResponseBody) GetPopCode() *string {
	return s.PopCode
}

func (s *DescribePopApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePopApiResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribePopApiResponseBody) SetApiName(v string) *DescribePopApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribePopApiResponseBody) SetOpenApiMetaList(v []*DescribePopApiResponseBodyOpenApiMetaList) *DescribePopApiResponseBody {
	s.OpenApiMetaList = v
	return s
}

func (s *DescribePopApiResponseBody) SetPopCode(v string) *DescribePopApiResponseBody {
	s.PopCode = &v
	return s
}

func (s *DescribePopApiResponseBody) SetRequestId(v string) *DescribePopApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePopApiResponseBody) SetVersion(v string) *DescribePopApiResponseBody {
	s.Version = &v
	return s
}

func (s *DescribePopApiResponseBody) Validate() error {
	if s.OpenApiMetaList != nil {
		for _, item := range s.OpenApiMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePopApiResponseBodyOpenApiMetaList struct {
	// The description of the parameter.
	//
	// example:
	//
	// demo parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The example value.
	//
	// example:
	//
	// 12.xx.xx.xx
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// DescribePopApi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// - **true**: required.
	//
	// - **false**: not required.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The serialization method for an array parameter. Valid values:
	//
	// - **repeatList**: An array is serialized in the XXX.N format. Example: Instance.1=i-instance1&\\&Instance.2=i-instance2.
	//
	// - **simple**: An array is serialized as a comma-separated string. Example: i-instance1,i-instance2.
	//
	// - **spaceDelimited**: An array is serialized as a space-separated string. Example: i-instance1 i-instance2.
	//
	// - **pipeDelimited**: An array is serialized as a pipe-separated string. Example: i-instance1|i-instance2.
	//
	// - **json**: An array is serialized in JSON format. Example: ["i-instance1","i-instance2"].
	//
	// - **flat**: An array is serialized in the XXX.N format. Example: Instance.1=i-instance1\\&Instance.2=i-instance2.
	//
	// example:
	//
	// json
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// - **string**: a string.
	//
	// - **boolean**: a Boolean value.
	//
	// - **integer**: an integer.
	//
	// - **long**: a long integer.
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePopApiResponseBodyOpenApiMetaList) String() string {
	return dara.Prettify(s)
}

func (s DescribePopApiResponseBodyOpenApiMetaList) GoString() string {
	return s.String()
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetDescription() *string {
	return s.Description
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetName() *string {
	return s.Name
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetRequired() *bool {
	return s.Required
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetStyle() *string {
	return s.Style
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) GetType() *string {
	return s.Type
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetDescription(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Description = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetExampleValue(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.ExampleValue = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetName(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Name = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetRequired(v bool) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Required = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetStyle(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Style = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) SetType(v string) *DescribePopApiResponseBodyOpenApiMetaList {
	s.Type = &v
	return s
}

func (s *DescribePopApiResponseBodyOpenApiMetaList) Validate() error {
	return dara.Validate(s)
}
