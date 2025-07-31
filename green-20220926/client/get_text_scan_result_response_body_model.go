// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextScanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTextScanResultResponseBody
	GetCode() *int32
	SetData(v *GetTextScanResultResponseBodyData) *GetTextScanResultResponseBody
	GetData() *GetTextScanResultResponseBodyData
	SetMsg(v string) *GetTextScanResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetTextScanResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTextScanResultResponseBody
	GetSuccess() *bool
}

type GetTextScanResultResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTextScanResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTextScanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTextScanResultResponseBody) GetData() *GetTextScanResultResponseBodyData {
	return s.Data
}

func (s *GetTextScanResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetTextScanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextScanResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTextScanResultResponseBody) SetCode(v int32) *GetTextScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetData(v *GetTextScanResultResponseBodyData) *GetTextScanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTextScanResultResponseBody) SetMsg(v string) *GetTextScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetRequestId(v string) *GetTextScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextScanResultResponseBody) SetSuccess(v bool) *GetTextScanResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetTextScanResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTextScanResultResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetTextScanResultResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTextScanResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetTextScanResultResponseBodyData) GetItems() []*GetTextScanResultResponseBodyDataItems {
	return s.Items
}

func (s *GetTextScanResultResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTextScanResultResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTextScanResultResponseBodyData) SetCurrentPage(v int32) *GetTextScanResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetItems(v []*GetTextScanResultResponseBodyDataItems) *GetTextScanResultResponseBodyData {
	s.Items = v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetPageSize(v int32) *GetTextScanResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetTextScanResultResponseBodyData) SetTotalCount(v int64) *GetTextScanResultResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetTextScanResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTextScanResultResponseBodyDataItems struct {
	BailianRequestId *string `json:"BailianRequestId,omitempty" xml:"BailianRequestId,omitempty"`
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// miss
	ExtFeedback *string `json:"ExtFeedback,omitempty" xml:"ExtFeedback,omitempty"`
	// example:
	//
	// {}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 2023-07-11 14:21:36
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// nonLabel
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2023-07-11 14:21:36
	RequestTime *string                                         `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	Result      []*GetTextScanResultResponseBodyDataItemsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	RiskLevel   *string                                         `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// 20
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// example:
	//
	// txtwkgb******AsYNXoJswy-1Aa1Qk
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTextScanResultResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyDataItems) GetBailianRequestId() *string {
	return s.BailianRequestId
}

func (s *GetTextScanResultResponseBodyDataItems) GetContent() *string {
	return s.Content
}

func (s *GetTextScanResultResponseBodyDataItems) GetExtFeedback() *string {
	return s.ExtFeedback
}

func (s *GetTextScanResultResponseBodyDataItems) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *GetTextScanResultResponseBodyDataItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTextScanResultResponseBodyDataItems) GetLabels() *string {
	return s.Labels
}

func (s *GetTextScanResultResponseBodyDataItems) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextScanResultResponseBodyDataItems) GetRequestTime() *string {
	return s.RequestTime
}

func (s *GetTextScanResultResponseBodyDataItems) GetResult() []*GetTextScanResultResponseBodyDataItemsResult {
	return s.Result
}

func (s *GetTextScanResultResponseBodyDataItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetTextScanResultResponseBodyDataItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *GetTextScanResultResponseBodyDataItems) GetScore() *float32 {
	return s.Score
}

func (s *GetTextScanResultResponseBodyDataItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetTextScanResultResponseBodyDataItems) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetTextScanResultResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTextScanResultResponseBodyDataItems) SetBailianRequestId(v string) *GetTextScanResultResponseBodyDataItems {
	s.BailianRequestId = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetContent(v string) *GetTextScanResultResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetExtFeedback(v string) *GetTextScanResultResponseBodyDataItems {
	s.ExtFeedback = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetExtra(v map[string]interface{}) *GetTextScanResultResponseBodyDataItems {
	s.Extra = v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetGmtCreate(v string) *GetTextScanResultResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetLabels(v string) *GetTextScanResultResponseBodyDataItems {
	s.Labels = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRequestId(v string) *GetTextScanResultResponseBodyDataItems {
	s.RequestId = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRequestTime(v string) *GetTextScanResultResponseBodyDataItems {
	s.RequestTime = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetResult(v []*GetTextScanResultResponseBodyDataItemsResult) *GetTextScanResultResponseBodyDataItems {
	s.Result = v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetRiskLevel(v string) *GetTextScanResultResponseBodyDataItems {
	s.RiskLevel = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetScanResult(v string) *GetTextScanResultResponseBodyDataItems {
	s.ScanResult = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetScore(v float32) *GetTextScanResultResponseBodyDataItems {
	s.Score = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetServiceCode(v string) *GetTextScanResultResponseBodyDataItems {
	s.ServiceCode = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetSuggestion(v string) *GetTextScanResultResponseBodyDataItems {
	s.Suggestion = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) SetTaskId(v string) *GetTextScanResultResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type GetTextScanResultResponseBodyDataItemsResult struct {
	// example:
	//
	// 25.0
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// political_n
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetTextScanResultResponseBodyDataItemsResult) String() string {
	return dara.Prettify(s)
}

func (s GetTextScanResultResponseBodyDataItemsResult) GoString() string {
	return s.String()
}

func (s *GetTextScanResultResponseBodyDataItemsResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetTextScanResultResponseBodyDataItemsResult) GetDescription() *string {
	return s.Description
}

func (s *GetTextScanResultResponseBodyDataItemsResult) GetLabel() *string {
	return s.Label
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetConfidence(v float32) *GetTextScanResultResponseBodyDataItemsResult {
	s.Confidence = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetDescription(v string) *GetTextScanResultResponseBodyDataItemsResult {
	s.Description = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItemsResult) SetLabel(v string) *GetTextScanResultResponseBodyDataItemsResult {
	s.Label = &v
	return s
}

func (s *GetTextScanResultResponseBodyDataItemsResult) Validate() error {
	return dara.Validate(s)
}
