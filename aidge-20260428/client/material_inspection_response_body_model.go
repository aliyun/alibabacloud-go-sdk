// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaterialInspectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MaterialInspectionResponseBody
	GetCode() *string
	SetData(v *MaterialInspectionResponseBodyData) *MaterialInspectionResponseBody
	GetData() *MaterialInspectionResponseBodyData
	SetMessage(v string) *MaterialInspectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *MaterialInspectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MaterialInspectionResponseBody
	GetSuccess() *bool
}

type MaterialInspectionResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MaterialInspectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MaterialInspectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionResponseBody) GoString() string {
	return s.String()
}

func (s *MaterialInspectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *MaterialInspectionResponseBody) GetData() *MaterialInspectionResponseBodyData {
	return s.Data
}

func (s *MaterialInspectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MaterialInspectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MaterialInspectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MaterialInspectionResponseBody) SetCode(v string) *MaterialInspectionResponseBody {
	s.Code = &v
	return s
}

func (s *MaterialInspectionResponseBody) SetData(v *MaterialInspectionResponseBodyData) *MaterialInspectionResponseBody {
	s.Data = v
	return s
}

func (s *MaterialInspectionResponseBody) SetMessage(v string) *MaterialInspectionResponseBody {
	s.Message = &v
	return s
}

func (s *MaterialInspectionResponseBody) SetRequestId(v string) *MaterialInspectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MaterialInspectionResponseBody) SetSuccess(v bool) *MaterialInspectionResponseBody {
	s.Success = &v
	return s
}

func (s *MaterialInspectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MaterialInspectionResponseBodyData struct {
	Result *MaterialInspectionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s MaterialInspectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *MaterialInspectionResponseBodyData) GetResult() *MaterialInspectionResponseBodyDataResult {
	return s.Result
}

func (s *MaterialInspectionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *MaterialInspectionResponseBodyData) SetResult(v *MaterialInspectionResponseBodyDataResult) *MaterialInspectionResponseBodyData {
	s.Result = v
	return s
}

func (s *MaterialInspectionResponseBodyData) SetUsageMap(v map[string]*int64) *MaterialInspectionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *MaterialInspectionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MaterialInspectionResponseBodyDataResult struct {
	// example:
	//
	// 1项规则：1 PASS，所有检测项合规
	Evidence *string `json:"Evidence,omitempty" xml:"Evidence,omitempty"`
	// example:
	//
	// PASS
	OverallResult *string `json:"OverallResult,omitempty" xml:"OverallResult,omitempty"`
	// example:
	//
	// req-001
	ReqId *string                                          `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	Steps []*MaterialInspectionResponseBodyDataResultSteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// example:
	//
	// Stamp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MaterialInspectionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MaterialInspectionResponseBodyDataResult) GetEvidence() *string {
	return s.Evidence
}

func (s *MaterialInspectionResponseBodyDataResult) GetOverallResult() *string {
	return s.OverallResult
}

func (s *MaterialInspectionResponseBodyDataResult) GetReqId() *string {
	return s.ReqId
}

func (s *MaterialInspectionResponseBodyDataResult) GetSteps() []*MaterialInspectionResponseBodyDataResultSteps {
	return s.Steps
}

func (s *MaterialInspectionResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *MaterialInspectionResponseBodyDataResult) SetEvidence(v string) *MaterialInspectionResponseBodyDataResult {
	s.Evidence = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResult) SetOverallResult(v string) *MaterialInspectionResponseBodyDataResult {
	s.OverallResult = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResult) SetReqId(v string) *MaterialInspectionResponseBodyDataResult {
	s.ReqId = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResult) SetSteps(v []*MaterialInspectionResponseBodyDataResultSteps) *MaterialInspectionResponseBodyDataResult {
	s.Steps = v
	return s
}

func (s *MaterialInspectionResponseBodyDataResult) SetType(v string) *MaterialInspectionResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResult) Validate() error {
	if s.Steps != nil {
		for _, item := range s.Steps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MaterialInspectionResponseBodyDataResultSteps struct {
	// example:
	//
	// PASS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// XXX
	StepId *string `json:"StepId,omitempty" xml:"StepId,omitempty"`
}

func (s MaterialInspectionResponseBodyDataResultSteps) String() string {
	return dara.Prettify(s)
}

func (s MaterialInspectionResponseBodyDataResultSteps) GoString() string {
	return s.String()
}

func (s *MaterialInspectionResponseBodyDataResultSteps) GetResult() *string {
	return s.Result
}

func (s *MaterialInspectionResponseBodyDataResultSteps) GetStepId() *string {
	return s.StepId
}

func (s *MaterialInspectionResponseBodyDataResultSteps) SetResult(v string) *MaterialInspectionResponseBodyDataResultSteps {
	s.Result = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResultSteps) SetStepId(v string) *MaterialInspectionResponseBodyDataResultSteps {
	s.StepId = &v
	return s
}

func (s *MaterialInspectionResponseBodyDataResultSteps) Validate() error {
	return dara.Validate(s)
}
