// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleInfoResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSampleInfoResponseBodyResultObject) *DescribeSampleInfoResponseBody
	GetResultObject() *DescribeSampleInfoResponseBodyResultObject
}

type DescribeSampleInfoResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeSampleInfoResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeSampleInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleInfoResponseBody) GetResultObject() *DescribeSampleInfoResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleInfoResponseBody) SetRequestId(v string) *DescribeSampleInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleInfoResponseBody) SetResultObject(v *DescribeSampleInfoResponseBodyResultObject) *DescribeSampleInfoResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSampleInfoResponseBodyResultObject struct {
	// Primary key ID
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Sample tags.
	//
	// example:
	//
	// em0102
	SampleTags *string `json:"sampleTags,omitempty" xml:"sampleTags,omitempty"`
	// Sample type
	//
	// example:
	//
	// PHONE
	SampleType *string `json:"sampleType,omitempty" xml:"sampleType,omitempty"`
	// Sample value.
	//
	// example:
	//
	// 17700000000
	SampleValue *string `json:"sampleValue,omitempty" xml:"sampleValue,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1753804800000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeSampleInfoResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleInfoResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetSampleTags() *string {
	return s.SampleTags
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetSampleType() *string {
	return s.SampleType
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetSampleValue() *string {
	return s.SampleValue
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSampleInfoResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetId(v int64) *DescribeSampleInfoResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetSampleTags(v string) *DescribeSampleInfoResponseBodyResultObject {
	s.SampleTags = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetSampleType(v string) *DescribeSampleInfoResponseBodyResultObject {
	s.SampleType = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetSampleValue(v string) *DescribeSampleInfoResponseBodyResultObject {
	s.SampleValue = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetUpdateTime(v string) *DescribeSampleInfoResponseBodyResultObject {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) SetVersion(v int32) *DescribeSampleInfoResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeSampleInfoResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
