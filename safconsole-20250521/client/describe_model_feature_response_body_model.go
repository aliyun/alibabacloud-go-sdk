// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeModelFeatureResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeModelFeatureResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *DescribeModelFeatureResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeModelFeatureResponseBodyResultObject) *DescribeModelFeatureResponseBody
	GetResultObject() []*DescribeModelFeatureResponseBodyResultObject
	SetSuccess(v bool) *DescribeModelFeatureResponseBody
	GetSuccess() *bool
}

type DescribeModelFeatureResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject []*DescribeModelFeatureResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModelFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelFeatureResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeModelFeatureResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeModelFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelFeatureResponseBody) GetResultObject() []*DescribeModelFeatureResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeModelFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModelFeatureResponseBody) SetCode(v int64) *DescribeModelFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModelFeatureResponseBody) SetHttpStatusCode(v int64) *DescribeModelFeatureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModelFeatureResponseBody) SetRequestId(v string) *DescribeModelFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelFeatureResponseBody) SetResultObject(v []*DescribeModelFeatureResponseBodyResultObject) *DescribeModelFeatureResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeModelFeatureResponseBody) SetSuccess(v bool) *DescribeModelFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModelFeatureResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelFeatureResponseBodyResultObject struct {
	// Default value of the feature.
	//
	// example:
	//
	// 0.1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Feature mapping name.
	//
	// example:
	//
	// f1
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// Feature type.
	//
	// example:
	//
	// double
	FeatureType *string `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// Feature name.
	//
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeModelFeatureResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelFeatureResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeModelFeatureResponseBodyResultObject) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeModelFeatureResponseBodyResultObject) GetFeatureName() *string {
	return s.FeatureName
}

func (s *DescribeModelFeatureResponseBodyResultObject) GetFeatureType() *string {
	return s.FeatureType
}

func (s *DescribeModelFeatureResponseBodyResultObject) GetName() *string {
	return s.Name
}

func (s *DescribeModelFeatureResponseBodyResultObject) SetDefaultValue(v string) *DescribeModelFeatureResponseBodyResultObject {
	s.DefaultValue = &v
	return s
}

func (s *DescribeModelFeatureResponseBodyResultObject) SetFeatureName(v string) *DescribeModelFeatureResponseBodyResultObject {
	s.FeatureName = &v
	return s
}

func (s *DescribeModelFeatureResponseBodyResultObject) SetFeatureType(v string) *DescribeModelFeatureResponseBodyResultObject {
	s.FeatureType = &v
	return s
}

func (s *DescribeModelFeatureResponseBodyResultObject) SetName(v string) *DescribeModelFeatureResponseBodyResultObject {
	s.Name = &v
	return s
}

func (s *DescribeModelFeatureResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
