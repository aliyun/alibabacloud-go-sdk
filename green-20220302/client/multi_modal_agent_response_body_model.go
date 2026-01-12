// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MultiModalAgentResponseBody
	GetCode() *string
	SetData(v *MultiModalAgentResponseBodyData) *MultiModalAgentResponseBody
	GetData() *MultiModalAgentResponseBodyData
	SetMessage(v string) *MultiModalAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *MultiModalAgentResponseBody
	GetRequestId() *string
}

type MultiModalAgentResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *MultiModalAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s MultiModalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *MultiModalAgentResponseBody) GetData() *MultiModalAgentResponseBodyData {
	return s.Data
}

func (s *MultiModalAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MultiModalAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalAgentResponseBody) SetCode(v string) *MultiModalAgentResponseBody {
	s.Code = &v
	return s
}

func (s *MultiModalAgentResponseBody) SetData(v *MultiModalAgentResponseBodyData) *MultiModalAgentResponseBody {
	s.Data = v
	return s
}

func (s *MultiModalAgentResponseBody) SetMessage(v string) *MultiModalAgentResponseBody {
	s.Message = &v
	return s
}

func (s *MultiModalAgentResponseBody) SetRequestId(v string) *MultiModalAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalAgentResponseBodyData struct {
	// example:
	//
	// 26769ada6e264e7ba9aa048241e12be9
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Result []*MultiModalAgentResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// high
	RiskLevel *string                               `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Usage     *MultiModalAgentResponseBodyDataUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s MultiModalAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *MultiModalAgentResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *MultiModalAgentResponseBodyData) GetResult() []*MultiModalAgentResponseBodyDataResult {
	return s.Result
}

func (s *MultiModalAgentResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *MultiModalAgentResponseBodyData) GetUsage() *MultiModalAgentResponseBodyDataUsage {
	return s.Usage
}

func (s *MultiModalAgentResponseBodyData) SetDataId(v string) *MultiModalAgentResponseBodyData {
	s.DataId = &v
	return s
}

func (s *MultiModalAgentResponseBodyData) SetResult(v []*MultiModalAgentResponseBodyDataResult) *MultiModalAgentResponseBodyData {
	s.Result = v
	return s
}

func (s *MultiModalAgentResponseBodyData) SetRiskLevel(v string) *MultiModalAgentResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *MultiModalAgentResponseBodyData) SetUsage(v *MultiModalAgentResponseBodyDataUsage) *MultiModalAgentResponseBodyData {
	s.Usage = v
	return s
}

func (s *MultiModalAgentResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MultiModalAgentResponseBodyDataResult struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// TRACER_SLB_ALL_DEST_WEIGHT_0
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s MultiModalAgentResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *MultiModalAgentResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *MultiModalAgentResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *MultiModalAgentResponseBodyDataResult) GetReason() *string {
	return s.Reason
}

func (s *MultiModalAgentResponseBodyDataResult) SetDescription(v string) *MultiModalAgentResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *MultiModalAgentResponseBodyDataResult) SetLabel(v string) *MultiModalAgentResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *MultiModalAgentResponseBodyDataResult) SetReason(v string) *MultiModalAgentResponseBodyDataResult {
	s.Reason = &v
	return s
}

func (s *MultiModalAgentResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type MultiModalAgentResponseBodyDataUsage struct {
	AgentDetail map[string]interface{} `json:"AgentDetail,omitempty" xml:"AgentDetail,omitempty"`
	// example:
	//
	// 10
	ContentLength *int64 `json:"ContentLength,omitempty" xml:"ContentLength,omitempty"`
	// example:
	//
	// 100
	PromptLength *int64 `json:"PromptLength,omitempty" xml:"PromptLength,omitempty"`
}

func (s MultiModalAgentResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *MultiModalAgentResponseBodyDataUsage) GetAgentDetail() map[string]interface{} {
	return s.AgentDetail
}

func (s *MultiModalAgentResponseBodyDataUsage) GetContentLength() *int64 {
	return s.ContentLength
}

func (s *MultiModalAgentResponseBodyDataUsage) GetPromptLength() *int64 {
	return s.PromptLength
}

func (s *MultiModalAgentResponseBodyDataUsage) SetAgentDetail(v map[string]interface{}) *MultiModalAgentResponseBodyDataUsage {
	s.AgentDetail = v
	return s
}

func (s *MultiModalAgentResponseBodyDataUsage) SetContentLength(v int64) *MultiModalAgentResponseBodyDataUsage {
	s.ContentLength = &v
	return s
}

func (s *MultiModalAgentResponseBodyDataUsage) SetPromptLength(v int64) *MultiModalAgentResponseBodyDataUsage {
	s.PromptLength = &v
	return s
}

func (s *MultiModalAgentResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
