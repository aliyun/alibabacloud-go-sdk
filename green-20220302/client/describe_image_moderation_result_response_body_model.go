// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModerationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeImageModerationResultResponseBody
	GetCode() *int32
	SetData(v *DescribeImageModerationResultResponseBodyData) *DescribeImageModerationResultResponseBody
	GetData() *DescribeImageModerationResultResponseBodyData
	SetMsg(v string) *DescribeImageModerationResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeImageModerationResultResponseBody
	GetRequestId() *string
}

type DescribeImageModerationResultResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image moderation results.
	Data *DescribeImageModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2881AD4F-638B-52A3-BA20-F74C5B1CEAE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageModerationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeImageModerationResultResponseBody) GetData() *DescribeImageModerationResultResponseBodyData {
	return s.Data
}

func (s *DescribeImageModerationResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeImageModerationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageModerationResultResponseBody) SetCode(v int32) *DescribeImageModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetData(v *DescribeImageModerationResultResponseBodyData) *DescribeImageModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetMsg(v string) *DescribeImageModerationResultResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeImageModerationResultResponseBody) SetRequestId(v string) *DescribeImageModerationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageModerationResultResponseBodyData struct {
	// The value of dataId that is specified in the API request. If this parameter is not specified in the API request, this field is not available in the response.
	//
	// example:
	//
	// 2a5389eb-4ff8-4584-ac99-644e2a539aa1
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The information about the captured frames.
	//
	// example:
	//
	// [{"result":[{"confidence":81.22,"label":"violent_explosion"}]}]
	Frame *string `json:"Frame,omitempty" xml:"Frame,omitempty"`
	// The number of frames.
	//
	// example:
	//
	// 1
	FrameNum     *int32  `json:"FrameNum,omitempty" xml:"FrameNum,omitempty"`
	ManualTaskId *string `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	// The reqId field returned by the Image Async Moderation API.
	//
	// example:
	//
	// B0963D30-BAB4-562F-9ED0-7A23AEC51C7C
	ReqId *string `json:"ReqId,omitempty" xml:"ReqId,omitempty"`
	// The results of image moderation parameters such as the label parameter and the confidence parameter.
	Result []*DescribeImageModerationResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Risk Level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageModerationResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *DescribeImageModerationResultResponseBodyData) GetFrame() *string {
	return s.Frame
}

func (s *DescribeImageModerationResultResponseBodyData) GetFrameNum() *int32 {
	return s.FrameNum
}

func (s *DescribeImageModerationResultResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *DescribeImageModerationResultResponseBodyData) GetReqId() *string {
	return s.ReqId
}

func (s *DescribeImageModerationResultResponseBodyData) GetResult() []*DescribeImageModerationResultResponseBodyDataResult {
	return s.Result
}

func (s *DescribeImageModerationResultResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageModerationResultResponseBodyData) SetDataId(v string) *DescribeImageModerationResultResponseBodyData {
	s.DataId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetFrame(v string) *DescribeImageModerationResultResponseBodyData {
	s.Frame = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetFrameNum(v int32) *DescribeImageModerationResultResponseBodyData {
	s.FrameNum = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetManualTaskId(v string) *DescribeImageModerationResultResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetReqId(v string) *DescribeImageModerationResultResponseBodyData {
	s.ReqId = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetResult(v []*DescribeImageModerationResultResponseBodyDataResult) *DescribeImageModerationResultResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) SetRiskLevel(v string) *DescribeImageModerationResultResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyData) Validate() error {
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

type DescribeImageModerationResultResponseBodyDataResult struct {
	// The score of the confidence level. Valid values: 0 to 100. The value is accurate to two decimal places.
	//
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the result.
	//
	// example:
	//
	// no risk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels returned after the image moderation.
	//
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Risk Level
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageModerationResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModerationResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DescribeImageModerationResultResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageModerationResultResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *DescribeImageModerationResultResponseBodyDataResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetConfidence(v float32) *DescribeImageModerationResultResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetDescription(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetLabel(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) SetRiskLevel(v string) *DescribeImageModerationResultResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageModerationResultResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
