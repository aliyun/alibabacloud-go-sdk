// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScanResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetScanResultResponseBody
	GetCode() *int32
	SetData(v *GetScanResultResponseBodyData) *GetScanResultResponseBody
	GetData() *GetScanResultResponseBodyData
	SetHttpStatusCode(v int32) *GetScanResultResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *GetScanResultResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetScanResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetScanResultResponseBody
	GetSuccess() *bool
}

type GetScanResultResponseBody struct {
	// Error code, consistent with HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *GetScanResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Further description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetScanResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetScanResultResponseBody) GetData() *GetScanResultResponseBodyData {
	return s.Data
}

func (s *GetScanResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetScanResultResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetScanResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScanResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetScanResultResponseBody) SetCode(v int32) *GetScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetScanResultResponseBody) SetData(v *GetScanResultResponseBodyData) *GetScanResultResponseBody {
	s.Data = v
	return s
}

func (s *GetScanResultResponseBody) SetHttpStatusCode(v int32) *GetScanResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetScanResultResponseBody) SetMsg(v string) *GetScanResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetScanResultResponseBody) SetRequestId(v string) *GetScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScanResultResponseBody) SetSuccess(v bool) *GetScanResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetScanResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScanResultResponseBodyData struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Data for the current page.
	Items []*GetScanResultResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetScanResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetScanResultResponseBodyData) GetItems() []*GetScanResultResponseBodyDataItems {
	return s.Items
}

func (s *GetScanResultResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScanResultResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetScanResultResponseBodyData) SetCurrentPage(v int32) *GetScanResultResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetScanResultResponseBodyData) SetItems(v []*GetScanResultResponseBodyDataItems) *GetScanResultResponseBodyData {
	s.Items = v
	return s
}

func (s *GetScanResultResponseBodyData) SetPageSize(v int32) *GetScanResultResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetScanResultResponseBodyData) SetTotalCount(v int64) *GetScanResultResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetScanResultResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScanResultResponseBodyDataItems struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Automated review labels.
	//
	// example:
	//
	// porn
	ApiLabels *string `json:"ApiLabels,omitempty" xml:"ApiLabels,omitempty"`
	// Machine review time.
	//
	// example:
	//
	// 1755501226
	ApiRequestTime *string `json:"ApiRequestTime,omitempty" xml:"ApiRequestTime,omitempty"`
	// Automated review risk level.
	//
	// example:
	//
	// high
	ApiRiskLevel *string `json:"ApiRiskLevel,omitempty" xml:"ApiRiskLevel,omitempty"`
	// Automated review service
	//
	// example:
	//
	// basecheckLine
	ApiService *string `json:"ApiService,omitempty" xml:"ApiService,omitempty"`
	// Automated review task ID.
	//
	// example:
	//
	// xxx
	ApiTaskId *string `json:"ApiTaskId,omitempty" xml:"ApiTaskId,omitempty"`
	// example:
	//
	// agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Attack level, returned based on the set high and low risk scores. The return values include:
	//
	// - high: High risk
	//
	// - medium: Medium risk
	//
	// - low: Low risk
	//
	// - none: No risk detected
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// Content.
	//
	// example:
	//
	// xxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Data Id
	//
	// example:
	//
	// 4f27b8cc7c4544cb90b41882a5b36326
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// Segment end time (in seconds).
	//
	// example:
	//
	// 22
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Feedback information.
	//
	// example:
	//
	// xxx
	ExtFeedback *string `json:"ExtFeedback,omitempty" xml:"ExtFeedback,omitempty"`
	// Additional parameters.
	//
	// example:
	//
	// {}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Frame count.
	//
	// example:
	//
	// 20
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Multimodal file URLs.
	GuardFileUrls []*string `json:"GuardFileUrls,omitempty" xml:"GuardFileUrls,omitempty" type:"Repeated"`
	// Multimodal image URLs.
	GuardImageUrls []*string `json:"GuardImageUrls,omitempty" xml:"GuardImageUrls,omitempty" type:"Repeated"`
	// Image labels.
	ImageLabels []map[string]interface{} `json:"ImageLabels,omitempty" xml:"ImageLabels,omitempty" type:"Repeated"`
	// Image service.
	//
	// example:
	//
	// baselineCheck
	ImageService *string `json:"ImageService,omitempty" xml:"ImageService,omitempty"`
	// URL
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	ImageUrl  *string   `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// Labels.
	//
	// example:
	//
	// nonLabel
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// Malicious file risk level.
	//
	// example:
	//
	// high
	MaliciousFileLevel *string `json:"MaliciousFileLevel,omitempty" xml:"MaliciousFileLevel,omitempty"`
	// Malicious URL risk level.
	//
	// example:
	//
	// high
	MaliciousUrlLevel *string `json:"MaliciousUrlLevel,omitempty" xml:"MaliciousUrlLevel,omitempty"`
	// Whether it is a pure manual review.
	//
	// example:
	//
	// false
	ManualOnly *bool `json:"ManualOnly,omitempty" xml:"ManualOnly,omitempty"`
	// No labels
	NoLabels []*string `json:"NoLabels,omitempty" xml:"NoLabels,omitempty" type:"Repeated"`
	// Frame offset value.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Request source.
	//
	// example:
	//
	// online_test
	RequestFrom *string `json:"RequestFrom,omitempty" xml:"RequestFrom,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request time.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// Resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Return collection.
	Result []*GetScanResultResponseBodyDataItemsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Review labels.
	//
	// example:
	//
	// porn
	ReviewLabels *string `json:"ReviewLabels,omitempty" xml:"ReviewLabels,omitempty"`
	// Review status.
	//
	// example:
	//
	// high
	ReviewRiskLevel *string `json:"ReviewRiskLevel,omitempty" xml:"ReviewRiskLevel,omitempty"`
	// Review time.
	//
	// example:
	//
	// 1755501226
	ReviewTime *string `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	// Reviewer.
	//
	// example:
	//
	// xx
	ReviewUid *string `json:"ReviewUid,omitempty" xml:"ReviewUid,omitempty"`
	// Whether it has been reviewed.
	//
	// example:
	//
	// false
	Reviewed *bool `json:"Reviewed,omitempty" xml:"Reviewed,omitempty"`
	// Risk level, returned based on the set high and low risk scores. The return values include:
	//
	// - high: High risk
	//
	// - medium: Medium risk
	//
	// - low: Low risk
	//
	// - none: No risk detected
	//
	// example:
	//
	// none
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Details of the detected risk.
	//
	// example:
	//
	// 色情服务
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// Keywords of the detected risk.
	//
	// example:
	//
	// 色情_低俗词
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// Details of the result.
	//
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// Score.
	//
	// example:
	//
	// 25
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// Sensitive level, returned based on the set high and low risk scores. The return values include:
	//
	// - **S1**: Indicates low sensitivity.
	//
	// - **S2**: Indicates medium sensitivity.
	//
	// - **S3**: Indicates high sensitivity.
	//
	// - **S4**: Indicates very high sensitivity.
	//
	// - **S0**: Indicates no sensitivity.
	//
	// example:
	//
	// S0
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Segment start time (in seconds).
	//
	// example:
	//
	// 11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Suggestion.
	//
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// Task ID.
	//
	// example:
	//
	// vi_s_EbrXb716LyBpkfwxyX5xyh-1A6RY9
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Text labels.
	TextLabels []map[string]interface{} `json:"TextLabels,omitempty" xml:"TextLabels,omitempty" type:"Repeated"`
	// Thumbnail URL.
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// Timestamp.
	//
	// example:
	//
	// 00:00:40-00:00:42
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// Task URL
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Voice labels.
	VoiceLabels []map[string]interface{} `json:"VoiceLabels,omitempty" xml:"VoiceLabels,omitempty" type:"Repeated"`
	// Whether audio detection is enabled.
	//
	// example:
	//
	// True
	VoiceScanOpened *bool `json:"VoiceScanOpened,omitempty" xml:"VoiceScanOpened,omitempty"`
	// Voice service.
	//
	// example:
	//
	// live_stream_detection
	VoiceService *string `json:"VoiceService,omitempty" xml:"VoiceService,omitempty"`
}

func (s GetScanResultResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyDataItems) GetAccountId() *string {
	return s.AccountId
}

func (s *GetScanResultResponseBodyDataItems) GetApiLabels() *string {
	return s.ApiLabels
}

func (s *GetScanResultResponseBodyDataItems) GetApiRequestTime() *string {
	return s.ApiRequestTime
}

func (s *GetScanResultResponseBodyDataItems) GetApiRiskLevel() *string {
	return s.ApiRiskLevel
}

func (s *GetScanResultResponseBodyDataItems) GetApiService() *string {
	return s.ApiService
}

func (s *GetScanResultResponseBodyDataItems) GetApiTaskId() *string {
	return s.ApiTaskId
}

func (s *GetScanResultResponseBodyDataItems) GetAppId() *string {
	return s.AppId
}

func (s *GetScanResultResponseBodyDataItems) GetAttackLevel() *string {
	return s.AttackLevel
}

func (s *GetScanResultResponseBodyDataItems) GetContent() *string {
	return s.Content
}

func (s *GetScanResultResponseBodyDataItems) GetDataId() *string {
	return s.DataId
}

func (s *GetScanResultResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *GetScanResultResponseBodyDataItems) GetExtFeedback() *string {
	return s.ExtFeedback
}

func (s *GetScanResultResponseBodyDataItems) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *GetScanResultResponseBodyDataItems) GetFrameCount() *int64 {
	return s.FrameCount
}

func (s *GetScanResultResponseBodyDataItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetScanResultResponseBodyDataItems) GetGuardFileUrls() []*string {
	return s.GuardFileUrls
}

func (s *GetScanResultResponseBodyDataItems) GetGuardImageUrls() []*string {
	return s.GuardImageUrls
}

func (s *GetScanResultResponseBodyDataItems) GetImageLabels() []map[string]interface{} {
	return s.ImageLabels
}

func (s *GetScanResultResponseBodyDataItems) GetImageService() *string {
	return s.ImageService
}

func (s *GetScanResultResponseBodyDataItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetScanResultResponseBodyDataItems) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *GetScanResultResponseBodyDataItems) GetLabels() *string {
	return s.Labels
}

func (s *GetScanResultResponseBodyDataItems) GetLiveId() *string {
	return s.LiveId
}

func (s *GetScanResultResponseBodyDataItems) GetMaliciousFileLevel() *string {
	return s.MaliciousFileLevel
}

func (s *GetScanResultResponseBodyDataItems) GetMaliciousUrlLevel() *string {
	return s.MaliciousUrlLevel
}

func (s *GetScanResultResponseBodyDataItems) GetManualOnly() *bool {
	return s.ManualOnly
}

func (s *GetScanResultResponseBodyDataItems) GetNoLabels() []*string {
	return s.NoLabels
}

func (s *GetScanResultResponseBodyDataItems) GetOffset() *int64 {
	return s.Offset
}

func (s *GetScanResultResponseBodyDataItems) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetScanResultResponseBodyDataItems) GetRequestFrom() *string {
	return s.RequestFrom
}

func (s *GetScanResultResponseBodyDataItems) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScanResultResponseBodyDataItems) GetRequestTime() *string {
	return s.RequestTime
}

func (s *GetScanResultResponseBodyDataItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetScanResultResponseBodyDataItems) GetResult() []*GetScanResultResponseBodyDataItemsResult {
	return s.Result
}

func (s *GetScanResultResponseBodyDataItems) GetReviewLabels() *string {
	return s.ReviewLabels
}

func (s *GetScanResultResponseBodyDataItems) GetReviewRiskLevel() *string {
	return s.ReviewRiskLevel
}

func (s *GetScanResultResponseBodyDataItems) GetReviewTime() *string {
	return s.ReviewTime
}

func (s *GetScanResultResponseBodyDataItems) GetReviewUid() *string {
	return s.ReviewUid
}

func (s *GetScanResultResponseBodyDataItems) GetReviewed() *bool {
	return s.Reviewed
}

func (s *GetScanResultResponseBodyDataItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetScanResultResponseBodyDataItems) GetRiskTips() *string {
	return s.RiskTips
}

func (s *GetScanResultResponseBodyDataItems) GetRiskWords() *string {
	return s.RiskWords
}

func (s *GetScanResultResponseBodyDataItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *GetScanResultResponseBodyDataItems) GetScore() *float32 {
	return s.Score
}

func (s *GetScanResultResponseBodyDataItems) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *GetScanResultResponseBodyDataItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetScanResultResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *GetScanResultResponseBodyDataItems) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetScanResultResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScanResultResponseBodyDataItems) GetTextLabels() []map[string]interface{} {
	return s.TextLabels
}

func (s *GetScanResultResponseBodyDataItems) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *GetScanResultResponseBodyDataItems) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *GetScanResultResponseBodyDataItems) GetUrl() *string {
	return s.Url
}

func (s *GetScanResultResponseBodyDataItems) GetVoiceLabels() []map[string]interface{} {
	return s.VoiceLabels
}

func (s *GetScanResultResponseBodyDataItems) GetVoiceScanOpened() *bool {
	return s.VoiceScanOpened
}

func (s *GetScanResultResponseBodyDataItems) GetVoiceService() *string {
	return s.VoiceService
}

func (s *GetScanResultResponseBodyDataItems) SetAccountId(v string) *GetScanResultResponseBodyDataItems {
	s.AccountId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetApiLabels(v string) *GetScanResultResponseBodyDataItems {
	s.ApiLabels = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetApiRequestTime(v string) *GetScanResultResponseBodyDataItems {
	s.ApiRequestTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetApiRiskLevel(v string) *GetScanResultResponseBodyDataItems {
	s.ApiRiskLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetApiService(v string) *GetScanResultResponseBodyDataItems {
	s.ApiService = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetApiTaskId(v string) *GetScanResultResponseBodyDataItems {
	s.ApiTaskId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetAppId(v string) *GetScanResultResponseBodyDataItems {
	s.AppId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetAttackLevel(v string) *GetScanResultResponseBodyDataItems {
	s.AttackLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetContent(v string) *GetScanResultResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetDataId(v string) *GetScanResultResponseBodyDataItems {
	s.DataId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetEndTime(v string) *GetScanResultResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetExtFeedback(v string) *GetScanResultResponseBodyDataItems {
	s.ExtFeedback = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetExtra(v map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.Extra = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetFrameCount(v int64) *GetScanResultResponseBodyDataItems {
	s.FrameCount = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetGmtCreate(v string) *GetScanResultResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetGuardFileUrls(v []*string) *GetScanResultResponseBodyDataItems {
	s.GuardFileUrls = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetGuardImageUrls(v []*string) *GetScanResultResponseBodyDataItems {
	s.GuardImageUrls = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.ImageLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageService(v string) *GetScanResultResponseBodyDataItems {
	s.ImageService = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageUrl(v string) *GetScanResultResponseBodyDataItems {
	s.ImageUrl = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetImageUrls(v []*string) *GetScanResultResponseBodyDataItems {
	s.ImageUrls = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetLabels(v string) *GetScanResultResponseBodyDataItems {
	s.Labels = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetLiveId(v string) *GetScanResultResponseBodyDataItems {
	s.LiveId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetMaliciousFileLevel(v string) *GetScanResultResponseBodyDataItems {
	s.MaliciousFileLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetMaliciousUrlLevel(v string) *GetScanResultResponseBodyDataItems {
	s.MaliciousUrlLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetManualOnly(v bool) *GetScanResultResponseBodyDataItems {
	s.ManualOnly = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetNoLabels(v []*string) *GetScanResultResponseBodyDataItems {
	s.NoLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetOffset(v int64) *GetScanResultResponseBodyDataItems {
	s.Offset = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetPageNum(v int64) *GetScanResultResponseBodyDataItems {
	s.PageNum = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestFrom(v string) *GetScanResultResponseBodyDataItems {
	s.RequestFrom = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestId(v string) *GetScanResultResponseBodyDataItems {
	s.RequestId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRequestTime(v string) *GetScanResultResponseBodyDataItems {
	s.RequestTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetResourceType(v string) *GetScanResultResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetResult(v []*GetScanResultResponseBodyDataItemsResult) *GetScanResultResponseBodyDataItems {
	s.Result = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetReviewLabels(v string) *GetScanResultResponseBodyDataItems {
	s.ReviewLabels = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetReviewRiskLevel(v string) *GetScanResultResponseBodyDataItems {
	s.ReviewRiskLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetReviewTime(v string) *GetScanResultResponseBodyDataItems {
	s.ReviewTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetReviewUid(v string) *GetScanResultResponseBodyDataItems {
	s.ReviewUid = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetReviewed(v bool) *GetScanResultResponseBodyDataItems {
	s.Reviewed = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskLevel(v string) *GetScanResultResponseBodyDataItems {
	s.RiskLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskTips(v string) *GetScanResultResponseBodyDataItems {
	s.RiskTips = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetRiskWords(v string) *GetScanResultResponseBodyDataItems {
	s.RiskWords = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetScanResult(v string) *GetScanResultResponseBodyDataItems {
	s.ScanResult = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetScore(v float32) *GetScanResultResponseBodyDataItems {
	s.Score = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetSensitiveLevel(v string) *GetScanResultResponseBodyDataItems {
	s.SensitiveLevel = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetServiceCode(v string) *GetScanResultResponseBodyDataItems {
	s.ServiceCode = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetStartTime(v string) *GetScanResultResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetSuggestion(v string) *GetScanResultResponseBodyDataItems {
	s.Suggestion = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTaskId(v string) *GetScanResultResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTextLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.TextLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetThumbnail(v string) *GetScanResultResponseBodyDataItems {
	s.Thumbnail = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetTimeStamp(v string) *GetScanResultResponseBodyDataItems {
	s.TimeStamp = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetUrl(v string) *GetScanResultResponseBodyDataItems {
	s.Url = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceLabels(v []map[string]interface{}) *GetScanResultResponseBodyDataItems {
	s.VoiceLabels = v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceScanOpened(v bool) *GetScanResultResponseBodyDataItems {
	s.VoiceScanOpened = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) SetVoiceService(v string) *GetScanResultResponseBodyDataItems {
	s.VoiceService = &v
	return s
}

func (s *GetScanResultResponseBodyDataItems) Validate() error {
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

type GetScanResultResponseBodyDataItemsResult struct {
	// Confidence score, ranging from 0 to 100, with two decimal places.
	//
	// example:
	//
	// 50.0
	Confidence *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// Description of the Label field.
	//
	// example:
	//
	// 疑似色情内容
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Label.
	//
	// example:
	//
	// politics
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetScanResultResponseBodyDataItemsResult) String() string {
	return dara.Prettify(s)
}

func (s GetScanResultResponseBodyDataItemsResult) GoString() string {
	return s.String()
}

func (s *GetScanResultResponseBodyDataItemsResult) GetConfidence() *string {
	return s.Confidence
}

func (s *GetScanResultResponseBodyDataItemsResult) GetDescription() *string {
	return s.Description
}

func (s *GetScanResultResponseBodyDataItemsResult) GetLabel() *string {
	return s.Label
}

func (s *GetScanResultResponseBodyDataItemsResult) SetConfidence(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Confidence = &v
	return s
}

func (s *GetScanResultResponseBodyDataItemsResult) SetDescription(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Description = &v
	return s
}

func (s *GetScanResultResponseBodyDataItemsResult) SetLabel(v string) *GetScanResultResponseBodyDataItemsResult {
	s.Label = &v
	return s
}

func (s *GetScanResultResponseBodyDataItemsResult) Validate() error {
	return dara.Validate(s)
}
