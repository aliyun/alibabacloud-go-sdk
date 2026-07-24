// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPromptTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPromptTestResultResponseBody
	GetRequestId() *string
	SetResult(v []*GetPromptTestResultResponseBodyResult) *GetPromptTestResultResponseBody
	GetResult() []*GetPromptTestResultResponseBodyResult
}

type GetPromptTestResultResponseBody struct {
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result.
	Result []*GetPromptTestResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetPromptTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPromptTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPromptTestResultResponseBody) GetResult() []*GetPromptTestResultResponseBodyResult {
	return s.Result
}

func (s *GetPromptTestResultResponseBody) SetRequestId(v string) *GetPromptTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPromptTestResultResponseBody) SetResult(v []*GetPromptTestResultResponseBodyResult) *GetPromptTestResultResponseBody {
	s.Result = v
	return s
}

func (s *GetPromptTestResultResponseBody) Validate() error {
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

type GetPromptTestResultResponseBodyResult struct {
	// The content.
	//
	// example:
	//
	// 测试文本
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The labels.
	LabelDetails []*GetPromptTestResultResponseBodyResultLabelDetails `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	// The risk level, which is returned based on the configured high and low risk scores. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	//
	//
	// - low: Low risk.
	//
	//  - none: No risk detected.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetPromptTestResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTestResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPromptTestResultResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetPromptTestResultResponseBodyResult) GetLabelDetails() []*GetPromptTestResultResponseBodyResultLabelDetails {
	return s.LabelDetails
}

func (s *GetPromptTestResultResponseBodyResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetPromptTestResultResponseBodyResult) SetContent(v string) *GetPromptTestResultResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetPromptTestResultResponseBodyResult) SetLabelDetails(v []*GetPromptTestResultResponseBodyResultLabelDetails) *GetPromptTestResultResponseBodyResult {
	s.LabelDetails = v
	return s
}

func (s *GetPromptTestResultResponseBodyResult) SetRiskLevel(v string) *GetPromptTestResultResponseBodyResult {
	s.RiskLevel = &v
	return s
}

func (s *GetPromptTestResultResponseBodyResult) Validate() error {
	if s.LabelDetails != nil {
		for _, item := range s.LabelDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPromptTestResultResponseBodyResultLabelDetails struct {
	// The label description.
	//
	// example:
	//
	// 暴恐
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels.
	//
	// example:
	//
	// terrorism
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The reason why the model determined this risk level for the text.
	//
	// example:
	//
	// 该文本涉及暴力恐怖信息
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s GetPromptTestResultResponseBodyResultLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPromptTestResultResponseBodyResultLabelDetails) GoString() string {
	return s.String()
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) GetDescription() *string {
	return s.Description
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) GetLabel() *string {
	return s.Label
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) GetReason() *string {
	return s.Reason
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) SetDescription(v string) *GetPromptTestResultResponseBodyResultLabelDetails {
	s.Description = &v
	return s
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) SetLabel(v string) *GetPromptTestResultResponseBodyResultLabelDetails {
	s.Label = &v
	return s
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) SetReason(v string) *GetPromptTestResultResponseBodyResultLabelDetails {
	s.Reason = &v
	return s
}

func (s *GetPromptTestResultResponseBodyResultLabelDetails) Validate() error {
	return dara.Validate(s)
}
