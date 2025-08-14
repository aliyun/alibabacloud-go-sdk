// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSampleResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateSampleResponseBodyResultObject) *CreateSampleResponseBody
	GetResultObject() *CreateSampleResponseBodyResultObject
}

type CreateSampleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *CreateSampleResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s CreateSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleResponseBody) GetResultObject() *CreateSampleResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateSampleResponseBody) SetRequestId(v string) *CreateSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleResponseBody) SetResultObject(v *CreateSampleResponseBodyResultObject) *CreateSampleResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateSampleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateSampleResponseBodyResultObject struct {
	// Number of colored groups
	//
	// example:
	//
	// 10
	CommunityCount *int32 `json:"communityCount,omitempty" xml:"communityCount,omitempty"`
	// Number of failed samples
	//
	// example:
	//
	// 1
	FailCount *int32 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// Recall probability
	//
	// example:
	//
	// 2.5%
	RecallProbability *string `json:"recallProbability,omitempty" xml:"recallProbability,omitempty"`
	// Risk density
	//
	// example:
	//
	// 1.5%
	RiskDensity *string `json:"riskDensity,omitempty" xml:"riskDensity,omitempty"`
	// Number of samples
	//
	// example:
	//
	// 100
	SampleCount *int32 `json:"sampleCount,omitempty" xml:"sampleCount,omitempty"`
	// Number of successful samples
	//
	// example:
	//
	// 99
	SuccessCount *int32 `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s CreateSampleResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateSampleResponseBodyResultObject) GetCommunityCount() *int32 {
	return s.CommunityCount
}

func (s *CreateSampleResponseBodyResultObject) GetFailCount() *int32 {
	return s.FailCount
}

func (s *CreateSampleResponseBodyResultObject) GetRecallProbability() *string {
	return s.RecallProbability
}

func (s *CreateSampleResponseBodyResultObject) GetRiskDensity() *string {
	return s.RiskDensity
}

func (s *CreateSampleResponseBodyResultObject) GetSampleCount() *int32 {
	return s.SampleCount
}

func (s *CreateSampleResponseBodyResultObject) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateSampleResponseBodyResultObject) SetCommunityCount(v int32) *CreateSampleResponseBodyResultObject {
	s.CommunityCount = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) SetFailCount(v int32) *CreateSampleResponseBodyResultObject {
	s.FailCount = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) SetRecallProbability(v string) *CreateSampleResponseBodyResultObject {
	s.RecallProbability = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) SetRiskDensity(v string) *CreateSampleResponseBodyResultObject {
	s.RiskDensity = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) SetSampleCount(v int32) *CreateSampleResponseBodyResultObject {
	s.SampleCount = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) SetSuccessCount(v int32) *CreateSampleResponseBodyResultObject {
	s.SuccessCount = &v
	return s
}

func (s *CreateSampleResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
