// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeepfakeDetectResponseBody
	GetCode() *string
	SetMessage(v string) *DeepfakeDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeepfakeDetectResponseBody
	GetRequestId() *string
	SetResultObject(v *DeepfakeDetectResponseBodyResultObject) *DeepfakeDetectResponseBody
	GetResultObject() *DeepfakeDetectResponseBodyResultObject
}

type DeepfakeDetectResponseBody struct {
	// Return code: 200 indicates success, others indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *DeepfakeDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepfakeDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeepfakeDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeepfakeDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeepfakeDetectResponseBody) GetResultObject() *DeepfakeDetectResponseBodyResultObject {
	return s.ResultObject
}

func (s *DeepfakeDetectResponseBody) SetCode(v string) *DeepfakeDetectResponseBody {
	s.Code = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetMessage(v string) *DeepfakeDetectResponseBody {
	s.Message = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetRequestId(v string) *DeepfakeDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeDetectResponseBody) SetResultObject(v *DeepfakeDetectResponseBodyResultObject) *DeepfakeDetectResponseBody {
	s.ResultObject = v
	return s
}

func (s *DeepfakeDetectResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeepfakeDetectResponseBodyResultObject struct {
	// Risk result:
	//
	// - **0**: Low risk
	//
	// - **1**: High risk
	//
	// - **2**: Suspicious
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// Risk score map.
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// Risk tags. Multiple tags are separated by commas (,). Includes:
	//
	// - Suspected deep forgery  SuspectDeepForgery
	//
	// - Suspected synthetic attack  SuspectPSFace
	//
	// - Suspected watermark  SuspectWarterMark
	//
	// - Suspected black industry attack  SuspectTemple
	//
	// - Suspected generated face  SuspectAIGC Face
	//
	// - Suspected rephotographed face  SuspectRemake
	//
	// example:
	//
	// SuspectDeepForgery,SuspectWarterMark
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s DeepfakeDetectResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *DeepfakeDetectResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *DeepfakeDetectResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *DeepfakeDetectResponseBodyResultObject) SetResult(v string) *DeepfakeDetectResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DeepfakeDetectResponseBodyResultObject) SetRiskScore(v map[string]*string) *DeepfakeDetectResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *DeepfakeDetectResponseBodyResultObject) SetRiskTag(v string) *DeepfakeDetectResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *DeepfakeDetectResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
