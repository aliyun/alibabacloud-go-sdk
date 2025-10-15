// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetSummaryTaskResultResponseBody
	GetCost() *int64
	SetData(v *GetSummaryTaskResultResponseBodyData) *GetSummaryTaskResultResponseBody
	GetData() *GetSummaryTaskResultResponseBodyData
	SetDataType(v string) *GetSummaryTaskResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetSummaryTaskResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetSummaryTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSummaryTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSummaryTaskResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetSummaryTaskResultResponseBody
	GetTime() *string
}

type GetSummaryTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetSummaryTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetSummaryTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetSummaryTaskResultResponseBody) GetData() *GetSummaryTaskResultResponseBodyData {
	return s.Data
}

func (s *GetSummaryTaskResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetSummaryTaskResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetSummaryTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSummaryTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSummaryTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSummaryTaskResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetSummaryTaskResultResponseBody) SetCost(v int64) *GetSummaryTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetData(v *GetSummaryTaskResultResponseBodyData) *GetSummaryTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetDataType(v string) *GetSummaryTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetErrCode(v string) *GetSummaryTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetMessage(v string) *GetSummaryTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetRequestId(v string) *GetSummaryTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetSuccess(v bool) *GetSummaryTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetTime(v string) *GetSummaryTaskResultResponseBody {
	s.Time = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSummaryTaskResultResponseBodyData struct {
	Choices []*GetSummaryTaskResultResponseBodyDataChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1726285125915
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 1202
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32                                     `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *GetSummaryTaskResultResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyData) GetChoices() []*GetSummaryTaskResultResponseBodyDataChoices {
	return s.Choices
}

func (s *GetSummaryTaskResultResponseBodyData) GetCreated() *int64 {
	return s.Created
}

func (s *GetSummaryTaskResultResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetSummaryTaskResultResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetSummaryTaskResultResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSummaryTaskResultResponseBodyData) GetTime() *string {
	return s.Time
}

func (s *GetSummaryTaskResultResponseBodyData) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *GetSummaryTaskResultResponseBodyData) GetUsage() *GetSummaryTaskResultResponseBodyDataUsage {
	return s.Usage
}

func (s *GetSummaryTaskResultResponseBodyData) SetChoices(v []*GetSummaryTaskResultResponseBodyDataChoices) *GetSummaryTaskResultResponseBodyData {
	s.Choices = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetCreated(v int64) *GetSummaryTaskResultResponseBodyData {
	s.Created = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetId(v string) *GetSummaryTaskResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetModelId(v string) *GetSummaryTaskResultResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetRequestId(v string) *GetSummaryTaskResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTime(v string) *GetSummaryTaskResultResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetUsage(v *GetSummaryTaskResultResponseBodyDataUsage) *GetSummaryTaskResultResponseBodyData {
	s.Usage = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) Validate() error {
	if s.Choices != nil {
		for _, item := range s.Choices {
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

type GetSummaryTaskResultResponseBodyDataChoices struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                              `json:"index,omitempty" xml:"index,omitempty"`
	Message *GetSummaryTaskResultResponseBodyDataChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyDataChoices) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoices) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) GetFinishReason() *string {
	return s.FinishReason
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) GetIndex() *int32 {
	return s.Index
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) GetMessage() *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	return s.Message
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetFinishReason(v string) *GetSummaryTaskResultResponseBodyDataChoices {
	s.FinishReason = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetIndex(v int32) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Index = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetMessage(v *GetSummaryTaskResultResponseBodyDataChoicesMessage) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Message = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSummaryTaskResultResponseBodyDataChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) GetContent() *string {
	return s.Content
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) GetRole() *string {
	return s.Role
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) GetToolCalls() []map[string]interface{} {
	return s.ToolCalls
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetContent(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Content = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetRole(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Role = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetToolCalls(v []map[string]interface{}) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.ToolCalls = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) Validate() error {
	return dara.Validate(s)
}

type GetSummaryTaskResultResponseBodyDataUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 100
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 200
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetSummaryTaskResultResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) GetImageTokens() *int32 {
	return s.ImageTokens
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageCount(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageCount = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetInputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetOutputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) Validate() error {
	return dara.Validate(s)
}
