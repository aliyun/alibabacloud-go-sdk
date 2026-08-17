// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DiduiAreaDeductionResponseBody
	GetCode() *string
	SetData(v *DiduiAreaDeductionResponseBodyData) *DiduiAreaDeductionResponseBody
	GetData() *DiduiAreaDeductionResponseBodyData
	SetMessage(v string) *DiduiAreaDeductionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DiduiAreaDeductionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DiduiAreaDeductionResponseBody
	GetSuccess() *bool
}

type DiduiAreaDeductionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The floor display area inference result.
	Data *DiduiAreaDeductionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiduiAreaDeductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionResponseBody) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DiduiAreaDeductionResponseBody) GetData() *DiduiAreaDeductionResponseBodyData {
	return s.Data
}

func (s *DiduiAreaDeductionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DiduiAreaDeductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiduiAreaDeductionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DiduiAreaDeductionResponseBody) SetCode(v string) *DiduiAreaDeductionResponseBody {
	s.Code = &v
	return s
}

func (s *DiduiAreaDeductionResponseBody) SetData(v *DiduiAreaDeductionResponseBodyData) *DiduiAreaDeductionResponseBody {
	s.Data = v
	return s
}

func (s *DiduiAreaDeductionResponseBody) SetMessage(v string) *DiduiAreaDeductionResponseBody {
	s.Message = &v
	return s
}

func (s *DiduiAreaDeductionResponseBody) SetRequestId(v string) *DiduiAreaDeductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiduiAreaDeductionResponseBody) SetSuccess(v bool) *DiduiAreaDeductionResponseBody {
	s.Success = &v
	return s
}

func (s *DiduiAreaDeductionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiduiAreaDeductionResponseBodyData struct {
	// The number of downstream call attempts.
	//
	// example:
	//
	// 1
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The workflow error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The downstream HTTP status code.
	//
	// example:
	//
	// OK
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The downstream call latency, in milliseconds.
	//
	// example:
	//
	// 1000
	LatencyMs *int64 `json:"LatencyMs,omitempty" xml:"LatencyMs,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The model request ID.
	//
	// example:
	//
	// model_req_id_xx
	ModelRequestId *string `json:"ModelRequestId,omitempty" xml:"ModelRequestId,omitempty"`
	// The business request ID.
	//
	// example:
	//
	// 643913D2-063F-599C-B3DF-B8D415CE171F
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The area calculation result.
	Result *DiduiAreaDeductionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The workflow status.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow business status.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessingCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s DiduiAreaDeductionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionResponseBodyData) GetAttempts() *int32 {
	return s.Attempts
}

func (s *DiduiAreaDeductionResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *DiduiAreaDeductionResponseBodyData) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *DiduiAreaDeductionResponseBodyData) GetLatencyMs() *int64 {
	return s.LatencyMs
}

func (s *DiduiAreaDeductionResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DiduiAreaDeductionResponseBodyData) GetModelRequestId() *string {
	return s.ModelRequestId
}

func (s *DiduiAreaDeductionResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *DiduiAreaDeductionResponseBodyData) GetResult() *DiduiAreaDeductionResponseBodyDataResult {
	return s.Result
}

func (s *DiduiAreaDeductionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DiduiAreaDeductionResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DiduiAreaDeductionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *DiduiAreaDeductionResponseBodyData) SetAttempts(v int32) *DiduiAreaDeductionResponseBodyData {
	s.Attempts = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetCode(v string) *DiduiAreaDeductionResponseBodyData {
	s.Code = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetHttpStatus(v int32) *DiduiAreaDeductionResponseBodyData {
	s.HttpStatus = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetLatencyMs(v int64) *DiduiAreaDeductionResponseBodyData {
	s.LatencyMs = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetMessage(v string) *DiduiAreaDeductionResponseBodyData {
	s.Message = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetModelRequestId(v string) *DiduiAreaDeductionResponseBodyData {
	s.ModelRequestId = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetReqId(v string) *DiduiAreaDeductionResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetResult(v *DiduiAreaDeductionResponseBodyDataResult) *DiduiAreaDeductionResponseBodyData {
	s.Result = v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetStatus(v string) *DiduiAreaDeductionResponseBodyData {
	s.Status = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetSuccess(v bool) *DiduiAreaDeductionResponseBodyData {
	s.Success = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) SetUsageMap(v map[string]*int64) *DiduiAreaDeductionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *DiduiAreaDeductionResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiduiAreaDeductionResponseBodyDataResult struct {
	// The stage 4 area label.
	Stage4AreaLabel *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel `json:"Stage4AreaLabel,omitempty" xml:"Stage4AreaLabel,omitempty" type:"Struct"`
}

func (s DiduiAreaDeductionResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionResponseBodyDataResult) GetStage4AreaLabel() *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel {
	return s.Stage4AreaLabel
}

func (s *DiduiAreaDeductionResponseBodyDataResult) SetStage4AreaLabel(v *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) *DiduiAreaDeductionResponseBodyDataResult {
	s.Stage4AreaLabel = v
	return s
}

func (s *DiduiAreaDeductionResponseBodyDataResult) Validate() error {
	if s.Stage4AreaLabel != nil {
		if err := s.Stage4AreaLabel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel struct {
	// The floor display area, in square meters.
	//
	// example:
	//
	// 1.5334
	AreaM2 *float64 `json:"AreaM2,omitempty" xml:"AreaM2,omitempty"`
	// The left edge length, in meters.
	//
	// example:
	//
	// 1.025
	LeftEdgeM *float64 `json:"LeftEdgeM,omitempty" xml:"LeftEdgeM,omitempty"`
	// The right edge length, in meters.
	//
	// example:
	//
	// 1.419
	RightEdgeM *float64 `json:"RightEdgeM,omitempty" xml:"RightEdgeM,omitempty"`
}

func (s DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) GetAreaM2() *float64 {
	return s.AreaM2
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) GetLeftEdgeM() *float64 {
	return s.LeftEdgeM
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) GetRightEdgeM() *float64 {
	return s.RightEdgeM
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) SetAreaM2(v float64) *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel {
	s.AreaM2 = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) SetLeftEdgeM(v float64) *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel {
	s.LeftEdgeM = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) SetRightEdgeM(v float64) *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel {
	s.RightEdgeM = &v
	return s
}

func (s *DiduiAreaDeductionResponseBodyDataResultStage4AreaLabel) Validate() error {
	return dara.Validate(s)
}
