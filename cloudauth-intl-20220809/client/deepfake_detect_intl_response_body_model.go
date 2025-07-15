// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeepfakeDetectIntlResponseBody
	GetCode() *string
	SetMessage(v string) *DeepfakeDetectIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeepfakeDetectIntlResponseBody
	GetRequestId() *string
	SetResultObject(v *DeepfakeDetectIntlResponseBodyResultObject) *DeepfakeDetectIntlResponseBody
	GetResultObject() *DeepfakeDetectIntlResponseBodyResultObject
}

type DeepfakeDetectIntlResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DeepfakeDetectIntlResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepfakeDetectIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeepfakeDetectIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeepfakeDetectIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeepfakeDetectIntlResponseBody) GetResultObject() *DeepfakeDetectIntlResponseBodyResultObject {
	return s.ResultObject
}

func (s *DeepfakeDetectIntlResponseBody) SetCode(v string) *DeepfakeDetectIntlResponseBody {
	s.Code = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetMessage(v string) *DeepfakeDetectIntlResponseBody {
	s.Message = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetRequestId(v string) *DeepfakeDetectIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) SetResultObject(v *DeepfakeDetectIntlResponseBodyResultObject) *DeepfakeDetectIntlResponseBody {
	s.ResultObject = v
	return s
}

func (s *DeepfakeDetectIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeepfakeDetectIntlResponseBodyResultObject struct {
	// example:
	//
	// 1
	Result    *string            `json:"Result,omitempty" xml:"Result,omitempty"`
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// example:
	//
	// SuspectDeepForgery,SuspectWarterMark
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s DeepfakeDetectIntlResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetResult(v string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetRiskScore(v map[string]*string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) SetRiskTag(v string) *DeepfakeDetectIntlResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *DeepfakeDetectIntlResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
