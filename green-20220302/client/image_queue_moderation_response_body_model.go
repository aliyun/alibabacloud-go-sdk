// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageQueueModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImageQueueModerationResponseBody
	GetCode() *int32
	SetData(v *ImageQueueModerationResponseBodyData) *ImageQueueModerationResponseBody
	GetData() *ImageQueueModerationResponseBodyData
	SetMsg(v string) *ImageQueueModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *ImageQueueModerationResponseBody
	GetRequestId() *string
}

type ImageQueueModerationResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImageQueueModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageQueueModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImageQueueModerationResponseBody) GetData() *ImageQueueModerationResponseBodyData {
	return s.Data
}

func (s *ImageQueueModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ImageQueueModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageQueueModerationResponseBody) SetCode(v int32) *ImageQueueModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageQueueModerationResponseBody) SetData(v *ImageQueueModerationResponseBodyData) *ImageQueueModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageQueueModerationResponseBody) SetMsg(v string) *ImageQueueModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageQueueModerationResponseBody) SetRequestId(v string) *ImageQueueModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageQueueModerationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageQueueModerationResponseBodyData struct {
	// example:
	//
	// data1234
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// [{"result":[{"confidence":81.22,"label":"violent_explosion"}]}]
	Frame *string `json:"Frame,omitempty" xml:"Frame,omitempty"`
	// example:
	//
	// 1
	FrameNum *int32                                        `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	Result   []*ImageQueueModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageQueueModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ImageQueueModerationResponseBodyData) GetFrame() *string {
	return s.Frame
}

func (s *ImageQueueModerationResponseBodyData) GetFrameNum() *int32 {
	return s.FrameNum
}

func (s *ImageQueueModerationResponseBodyData) GetResult() []*ImageQueueModerationResponseBodyDataResult {
	return s.Result
}

func (s *ImageQueueModerationResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageQueueModerationResponseBodyData) SetDataId(v string) *ImageQueueModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetFrame(v string) *ImageQueueModerationResponseBodyData {
	s.Frame = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetFrameNum(v int32) *ImageQueueModerationResponseBodyData {
	s.FrameNum = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetResult(v []*ImageQueueModerationResponseBodyDataResult) *ImageQueueModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetRiskLevel(v string) *ImageQueueModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) Validate() error {
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

type ImageQueueModerationResponseBodyDataResult struct {
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ImageQueueModerationResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *ImageQueueModerationResponseBodyDataResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageQueueModerationResponseBodyDataResult) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetDescription(v string) *ImageQueueModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetLabel(v string) *ImageQueueModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetRiskLevel(v string) *ImageQueueModerationResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
