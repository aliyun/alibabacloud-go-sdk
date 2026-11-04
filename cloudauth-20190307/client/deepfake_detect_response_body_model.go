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
	// The return code. A value of 200 indicates success, and other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeepfakeDetectResponseBodyResultObject struct {
	// The facial attribute output fields.
	//
	// example:
	//
	// {\\"age\\":49,\\"env\\":\\"other\\",\\"expression\\":0,\\"gender\\":1,\\"glasses\\":1,\\"hat\\":0,\\"mask\\":0,\\"others\\":0,\\"smoking\\":0,\\"tattoo\\":0}
	FaceAttributeInfo *string `json:"FaceAttributeInfo,omitempty" xml:"FaceAttributeInfo,omitempty"`
	// The risk result. Valid values:
	//
	// - **0**: low risk.
	//
	// - **1**: high risk.
	//
	// - **2**: suspicious.
	//
	// example:
	//
	// 1
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The map of risk scores.
	RiskScore map[string]*string `json:"RiskScore,omitempty" xml:"RiskScore,omitempty"`
	// The risk tags. Multiple tags are separated by commas (,). Valid values:
	//
	// - SuspectDeepForgery: suspected deepfake.
	//
	// - SuspectPSFace: suspected synthetic attack.
	//
	// - SuspectTemple: suspected black market attack.
	//
	// - SuspectRemake: suspected recaptured face.
	//
	// example:
	//
	// SuspectDeepForgery, SuspectPSFace
	RiskTag *string `json:"RiskTag,omitempty" xml:"RiskTag,omitempty"`
}

func (s DeepfakeDetectResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectResponseBodyResultObject) GetFaceAttributeInfo() *string {
	return s.FaceAttributeInfo
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

func (s *DeepfakeDetectResponseBodyResultObject) SetFaceAttributeInfo(v string) *DeepfakeDetectResponseBodyResultObject {
	s.FaceAttributeInfo = &v
	return s
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
