// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeepfakeDetectIntlStreamResponseBody
	GetCode() *string
	SetMessage(v string) *DeepfakeDetectIntlStreamResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeepfakeDetectIntlStreamResponseBody
	GetRequestId() *string
	SetResultObject(v *DeepfakeDetectIntlStreamResponseBodyResultObject) *DeepfakeDetectIntlStreamResponseBody
	GetResultObject() *DeepfakeDetectIntlStreamResponseBodyResultObject
}

type DeepfakeDetectIntlStreamResponseBody struct {
	// Return code
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information.
	ResultObject *DeepfakeDetectIntlStreamResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s DeepfakeDetectIntlStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlStreamResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeepfakeDetectIntlStreamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeepfakeDetectIntlStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeepfakeDetectIntlStreamResponseBody) GetResultObject() *DeepfakeDetectIntlStreamResponseBodyResultObject {
	return s.ResultObject
}

func (s *DeepfakeDetectIntlStreamResponseBody) SetCode(v string) *DeepfakeDetectIntlStreamResponseBody {
	s.Code = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBody) SetMessage(v string) *DeepfakeDetectIntlStreamResponseBody {
	s.Message = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBody) SetRequestId(v string) *DeepfakeDetectIntlStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBody) SetResultObject(v *DeepfakeDetectIntlStreamResponseBodyResultObject) *DeepfakeDetectIntlStreamResponseBody {
	s.ResultObject = v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeepfakeDetectIntlStreamResponseBodyResultObject struct {
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
	// Risk tags. Multiple tags are separated by commas (,), including:
	//
	// - SuspectDeepForgery: Suspected deep forgery
	//
	// - SuspectPSFace: Suspected synthetic attack
	//
	// - SuspectTemple: Suspected template attack
	//
	// - SuspectRemake: Suspected presentation attack
	//
	// example:
	//
	// SuspectDeepForgery
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s DeepfakeDetectIntlStreamResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlStreamResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) GetRiskScore() map[string]*string {
	return s.RiskScore
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) GetRiskTag() *string {
	return s.RiskTag
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) SetResult(v string) *DeepfakeDetectIntlStreamResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) SetRiskScore(v map[string]*string) *DeepfakeDetectIntlStreamResponseBodyResultObject {
	s.RiskScore = v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) SetRiskTag(v string) *DeepfakeDetectIntlStreamResponseBodyResultObject {
	s.RiskTag = &v
	return s
}

func (s *DeepfakeDetectIntlStreamResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
