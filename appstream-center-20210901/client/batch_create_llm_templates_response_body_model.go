// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateLlmTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchCreateLlmTemplatesResponseBodyData) *BatchCreateLlmTemplatesResponseBody
	GetData() *BatchCreateLlmTemplatesResponseBodyData
	SetRequestId(v string) *BatchCreateLlmTemplatesResponseBody
	GetRequestId() *string
}

type BatchCreateLlmTemplatesResponseBody struct {
	Data *BatchCreateLlmTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCreateLlmTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesResponseBody) GetData() *BatchCreateLlmTemplatesResponseBodyData {
	return s.Data
}

func (s *BatchCreateLlmTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateLlmTemplatesResponseBody) SetData(v *BatchCreateLlmTemplatesResponseBodyData) *BatchCreateLlmTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBody) SetRequestId(v string) *BatchCreateLlmTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchCreateLlmTemplatesResponseBodyData struct {
	LlmTemplateIds []*string                                              `json:"LlmTemplateIds,omitempty" xml:"LlmTemplateIds,omitempty" type:"Repeated"`
	SkippedItems   []*BatchCreateLlmTemplatesResponseBodyDataSkippedItems `json:"SkippedItems,omitempty" xml:"SkippedItems,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s BatchCreateLlmTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesResponseBodyData) GetLlmTemplateIds() []*string {
	return s.LlmTemplateIds
}

func (s *BatchCreateLlmTemplatesResponseBodyData) GetSkippedItems() []*BatchCreateLlmTemplatesResponseBodyDataSkippedItems {
	return s.SkippedItems
}

func (s *BatchCreateLlmTemplatesResponseBodyData) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *BatchCreateLlmTemplatesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *BatchCreateLlmTemplatesResponseBodyData) SetLlmTemplateIds(v []*string) *BatchCreateLlmTemplatesResponseBodyData {
	s.LlmTemplateIds = v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyData) SetSkippedItems(v []*BatchCreateLlmTemplatesResponseBodyDataSkippedItems) *BatchCreateLlmTemplatesResponseBodyData {
	s.SkippedItems = v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyData) SetSuccessCount(v int32) *BatchCreateLlmTemplatesResponseBodyData {
	s.SuccessCount = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyData) SetTotalCount(v int32) *BatchCreateLlmTemplatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyData) Validate() error {
	if s.SkippedItems != nil {
		for _, item := range s.SkippedItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateLlmTemplatesResponseBodyDataSkippedItems struct {
	// example:
	//
	// qwen3.5-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// example:
	//
	// LLM template already exists with same providerTemplateId and llmCode.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s BatchCreateLlmTemplatesResponseBodyDataSkippedItems) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesResponseBodyDataSkippedItems) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesResponseBodyDataSkippedItems) GetLlmCode() *string {
	return s.LlmCode
}

func (s *BatchCreateLlmTemplatesResponseBodyDataSkippedItems) GetReason() *string {
	return s.Reason
}

func (s *BatchCreateLlmTemplatesResponseBodyDataSkippedItems) SetLlmCode(v string) *BatchCreateLlmTemplatesResponseBodyDataSkippedItems {
	s.LlmCode = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyDataSkippedItems) SetReason(v string) *BatchCreateLlmTemplatesResponseBodyDataSkippedItems {
	s.Reason = &v
	return s
}

func (s *BatchCreateLlmTemplatesResponseBodyDataSkippedItems) Validate() error {
	return dara.Validate(s)
}
