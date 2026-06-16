// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ManualModerationResultResponseBody
	GetCode() *int32
	SetData(v *ManualModerationResultResponseBodyData) *ManualModerationResultResponseBody
	GetData() *ManualModerationResultResponseBodyData
	SetMessage(v string) *ManualModerationResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *ManualModerationResultResponseBody
	GetRequestId() *string
}

type ManualModerationResultResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ManualModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ManualModerationResultResponseBody) GetData() *ManualModerationResultResponseBodyData {
	return s.Data
}

func (s *ManualModerationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ManualModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManualModerationResultResponseBody) SetCode(v int32) *ManualModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *ManualModerationResultResponseBody) SetData(v *ManualModerationResultResponseBodyData) *ManualModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *ManualModerationResultResponseBody) SetMessage(v string) *ManualModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *ManualModerationResultResponseBody) SetRequestId(v string) *ManualModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManualModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ManualModerationResultResponseBodyData struct {
	// The value of the dataId parameter that you specified in the API request. If you did not specify this parameter in the request, this field is not returned.
	//
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The detailed results of the labels.
	Result []*ManualModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The risk level. The value is returned based on the specified high-risk and low-risk scores. Valid values include the following:
	//
	// - high: high risk
	//
	// - low: low risk
	//
	// - none: no risk detected
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The task ID.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ManualModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ManualModerationResultResponseBodyData) GetResult() []*ManualModerationResultResponseBodyDataResult {
	return s.Result
}

func (s *ManualModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ManualModerationResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ManualModerationResultResponseBodyData) SetDataId(v string) *ManualModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetResult(v []*ManualModerationResultResponseBodyDataResult) *ManualModerationResultResponseBodyData {
	s.Result = v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetRiskLevel(v string) *ManualModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ManualModerationResultResponseBodyData) SetTaskId(v string) *ManualModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ManualModerationResultResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ManualModerationResultResponseBodyDataResult struct {
	// The description of the label.
	//
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk label.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ManualModerationResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ManualModerationResultResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ManualModerationResultResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *ManualModerationResultResponseBodyDataResult) SetDescription(v string) *ManualModerationResultResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ManualModerationResultResponseBodyDataResult) SetLabel(v string) *ManualModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ManualModerationResultResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
