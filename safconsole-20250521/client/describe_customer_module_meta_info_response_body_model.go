// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleMetaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeCustomerModuleMetaInfoResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeCustomerModuleMetaInfoResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeCustomerModuleMetaInfoResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeCustomerModuleMetaInfoResponseBodyResultObject) *DescribeCustomerModuleMetaInfoResponseBody
	GetResultObject() *DescribeCustomerModuleMetaInfoResponseBodyResultObject
	SetSuccess(v bool) *DescribeCustomerModuleMetaInfoResponseBody
	GetSuccess() *bool
}

type DescribeCustomerModuleMetaInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 055f1546-f465-4c92-a2da-bfb86abe6f56
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeCustomerModuleMetaInfoResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomerModuleMetaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleMetaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) GetResultObject() *DescribeCustomerModuleMetaInfoResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) SetCode(v int64) *DescribeCustomerModuleMetaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) SetHttpStatusCode(v int64) *DescribeCustomerModuleMetaInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) SetRequestId(v string) *DescribeCustomerModuleMetaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) SetResultObject(v *DescribeCustomerModuleMetaInfoResponseBodyResultObject) *DescribeCustomerModuleMetaInfoResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) SetSuccess(v bool) *DescribeCustomerModuleMetaInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerModuleMetaInfoResponseBodyResultObject struct {
	FeatureList []*DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList `json:"FeatureList,omitempty" xml:"FeatureList,omitempty" type:"Repeated"`
	// example:
	//
	// FINANCE_60
	FeatureTemplate *string `json:"FeatureTemplate,omitempty" xml:"FeatureTemplate,omitempty"`
}

func (s DescribeCustomerModuleMetaInfoResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleMetaInfoResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObject) GetFeatureList() []*DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList {
	return s.FeatureList
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObject) GetFeatureTemplate() *string {
	return s.FeatureTemplate
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObject) SetFeatureList(v []*DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) *DescribeCustomerModuleMetaInfoResponseBodyResultObject {
	s.FeatureList = v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObject) SetFeatureTemplate(v string) *DescribeCustomerModuleMetaInfoResponseBodyResultObject {
	s.FeatureTemplate = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObject) Validate() error {
	if s.FeatureList != nil {
		for _, item := range s.FeatureList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList struct {
	// example:
	//
	// 0.1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// f1
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// example:
	//
	// double
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) GetFeatureName() *string {
	return s.FeatureName
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) GetFeatureType() *string {
	return s.FeatureType
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) GetName() *string {
	return s.Name
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) SetDefaultValue(v string) *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList {
	s.DefaultValue = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) SetFeatureName(v string) *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList {
	s.FeatureName = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) SetFeatureType(v string) *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList {
	s.FeatureType = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) SetName(v string) *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList {
	s.Name = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponseBodyResultObjectFeatureList) Validate() error {
	return dara.Validate(s)
}
