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
	// The error code, which is consistent with the HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetScanResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The detailed description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The success flag.
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
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data on the current page.
	Items []*GetScanResultResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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
	// The AccountId parameter provided by the customer.
	//
	// example:
	//
	// accountIdtest123
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The machine-assisted moderation labels.
	//
	// example:
	//
	// porn
	ApiLabels *string `json:"ApiLabels,omitempty" xml:"ApiLabels,omitempty"`
	// The machine-assisted moderation time. The value is a UNIX/POSIX timestamp in milliseconds.
	//
	// example:
	//
	// 1755501226
	ApiRequestTime *string `json:"ApiRequestTime,omitempty" xml:"ApiRequestTime,omitempty"`
	// The machine-assisted moderation risk level.
	//
	// example:
	//
	// high
	ApiRiskLevel *string `json:"ApiRiskLevel,omitempty" xml:"ApiRiskLevel,omitempty"`
	// The machine-assisted moderation service.
	//
	// example:
	//
	// basecheckLine
	ApiService *string `json:"ApiService,omitempty" xml:"ApiService,omitempty"`
	// The machine-assisted moderation task ID.
	//
	// example:
	//
	// img-a1b2c3d4e5f60718
	ApiTaskId *string `json:"ApiTaskId,omitempty" xml:"ApiTaskId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// agent_01
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The attack level. The value is returned based on the configured high and low risk scores. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// none
	AttackLevel *string `json:"AttackLevel,omitempty" xml:"AttackLevel,omitempty"`
	// The Model Studio request ID.
	//
	// example:
	//
	// 93a2f484-ac61-9b2e-9765-88600c83c1ea
	BailianRequestId *string `json:"BailianRequestId,omitempty" xml:"BailianRequestId,omitempty"`
	// The content.
	//
	// example:
	//
	// This is a sample text content for detection
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The data ID.
	//
	// example:
	//
	// 4f27b8cc7c4544cb90b41882a5b36326
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The end time of the segment, in seconds.
	//
	// example:
	//
	// 22
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information.
	//
	// example:
	//
	// multiImage
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The feedback information.
	//
	// example:
	//
	// Violating content was detected, and the review conclusion is a violation
	ExtFeedback *string `json:"ExtFeedback,omitempty" xml:"ExtFeedback,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// {}
	Extra map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The number of frames.
	//
	// example:
	//
	// 20
	FrameCount *int64 `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	// The creation time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The multimodal file URLs.
	GuardFileUrls []*string `json:"GuardFileUrls,omitempty" xml:"GuardFileUrls,omitempty" type:"Repeated"`
	// The multimodal image URLs.
	GuardImageUrls []*string `json:"GuardImageUrls,omitempty" xml:"GuardImageUrls,omitempty" type:"Repeated"`
	// The image labels.
	ImageLabels []map[string]interface{} `json:"ImageLabels,omitempty" xml:"ImageLabels,omitempty" type:"Repeated"`
	// The image service.
	//
	// example:
	//
	// baselineCheck
	ImageService *string `json:"ImageService,omitempty" xml:"ImageService,omitempty"`
	// The image URL.
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The list of image URLs.
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	// The labels.
	//
	// example:
	//
	// nonLabel
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The LiveId parameter provided by the customer.
	//
	// example:
	//
	// liveIdtest123
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	// The risk level of the malicious file.
	//
	// example:
	//
	// high
	MaliciousFileLevel *string `json:"MaliciousFileLevel,omitempty" xml:"MaliciousFileLevel,omitempty"`
	// The risk level of the malicious URL.
	//
	// example:
	//
	// high
	MaliciousUrlLevel *string `json:"MaliciousUrlLevel,omitempty" xml:"MaliciousUrlLevel,omitempty"`
	// Specifies whether the task is exclusively manually reviewed.
	//
	// example:
	//
	// false
	ManualOnly *bool `json:"ManualOnly,omitempty" xml:"ManualOnly,omitempty"`
	// No labels.
	NoLabels []*string `json:"NoLabels,omitempty" xml:"NoLabels,omitempty" type:"Repeated"`
	// The frame capture offset value.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The request source.
	//
	// example:
	//
	// online_test
	RequestFrom *string `json:"RequestFrom,omitempty" xml:"RequestFrom,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request time. The format is YYYY-MM-DD HH:mm:ss.
	//
	// example:
	//
	// 2023-08-11 09:00:19
	RequestTime *string `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	// The resource type.
	//
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The returned collection.
	Result []*GetScanResultResponseBodyDataItemsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The review labels.
	//
	// example:
	//
	// porn
	ReviewLabels *string `json:"ReviewLabels,omitempty" xml:"ReviewLabels,omitempty"`
	// The review status.
	//
	// example:
	//
	// high
	ReviewRiskLevel *string `json:"ReviewRiskLevel,omitempty" xml:"ReviewRiskLevel,omitempty"`
	// The review time. The value is a UNIX/POSIX timestamp in milliseconds.
	//
	// example:
	//
	// 1755501226
	ReviewTime *string `json:"ReviewTime,omitempty" xml:"ReviewTime,omitempty"`
	// The reviewer.
	//
	// example:
	//
	// 1234567890123456
	ReviewUid *string `json:"ReviewUid,omitempty" xml:"ReviewUid,omitempty"`
	// Specifies whether the task has been reviewed.
	//
	// example:
	//
	// false
	Reviewed *bool `json:"Reviewed,omitempty" xml:"Reviewed,omitempty"`
	// The risk level. The value is returned based on the configured high and low risk scores. Valid values:
	//
	// - high: High risk.
	//
	// - medium: Medium risk.
	//
	// - low: Low risk.
	//
	// - none: No risk detected.
	//
	// example:
	//
	// none
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The details of the hit risk.
	//
	// example:
	//
	// Pornographic services
	RiskTips *string `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	// The hit risk keywords.
	//
	// example:
	//
	// Pornography_vulgar words
	RiskWords *string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	// The result details.
	//
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// The score.
	//
	// example:
	//
	// 25
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The sensitivity level. The value is returned based on the configured sensitivity scores. Valid values:
	//
	// - **S1**: Low sensitivity.
	//
	// - **S2**: Medium sensitivity.
	//
	// - **S3**: Relatively high sensitivity.
	//
	// - **S4**: High sensitivity.
	//
	// - **S0**: Non-sensitive.
	//
	// example:
	//
	// S0
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The start time of the segment, in seconds.
	//
	// example:
	//
	// 11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The suggestion.
	//
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// vi_s_EbrXb716LyBpkfwxyX5xyh-1A6RY9
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The text labels.
	TextLabels []map[string]interface{} `json:"TextLabels,omitempty" xml:"TextLabels,omitempty" type:"Repeated"`
	// The thumbnail URL.
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 00:00:40-00:00:42
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The task URL.
	//
	// example:
	//
	// https://www.aliyuncs.com/xxx.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The audio labels.
	VoiceLabels []map[string]interface{} `json:"VoiceLabels,omitempty" xml:"VoiceLabels,omitempty" type:"Repeated"`
	// Specifies whether audio detection is enabled.
	//
	// example:
	//
	// True
	VoiceScanOpened *bool `json:"VoiceScanOpened,omitempty" xml:"VoiceScanOpened,omitempty"`
	// The audio service.
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

func (s *GetScanResultResponseBodyDataItems) GetBailianRequestId() *string {
	return s.BailianRequestId
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

func (s *GetScanResultResponseBodyDataItems) GetExt() *string {
	return s.Ext
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

func (s *GetScanResultResponseBodyDataItems) SetBailianRequestId(v string) *GetScanResultResponseBodyDataItems {
	s.BailianRequestId = &v
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

func (s *GetScanResultResponseBodyDataItems) SetExt(v string) *GetScanResultResponseBodyDataItems {
	s.Ext = &v
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
	// The confidence score. The value ranges from 0 to 100 and is rounded to two decimal places.
	//
	// example:
	//
	// 50.0
	Confidence *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The description of the Label field.
	//
	// example:
	//
	// Suspected pornographic content
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The labels.
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
