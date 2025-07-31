// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckFreezeResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetOssCheckFreezeResultResponseBody
	GetCurrentPage() *int32
	SetItems(v []*GetOssCheckFreezeResultResponseBodyItems) *GetOssCheckFreezeResultResponseBody
	GetItems() []*GetOssCheckFreezeResultResponseBodyItems
	SetPageSize(v int32) *GetOssCheckFreezeResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetOssCheckFreezeResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetOssCheckFreezeResultResponseBody
	GetTotalCount() *int64
}

type GetOssCheckFreezeResultResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetOssCheckFreezeResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F0A594BB-FA7A-580F-AE9E-A4188E092823
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 29
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOssCheckFreezeResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOssCheckFreezeResultResponseBody) GetItems() []*GetOssCheckFreezeResultResponseBodyItems {
	return s.Items
}

func (s *GetOssCheckFreezeResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOssCheckFreezeResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckFreezeResultResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetOssCheckFreezeResultResponseBody) SetCurrentPage(v int32) *GetOssCheckFreezeResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBody) SetItems(v []*GetOssCheckFreezeResultResponseBodyItems) *GetOssCheckFreezeResultResponseBody {
	s.Items = v
	return s
}

func (s *GetOssCheckFreezeResultResponseBody) SetPageSize(v int32) *GetOssCheckFreezeResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBody) SetRequestId(v string) *GetOssCheckFreezeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBody) SetTotalCount(v int64) *GetOssCheckFreezeResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckFreezeResultResponseBodyItems struct {
	// example:
	//
	// tmp
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// audio
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// audio_media_detection
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Freeze   *bool   `json:"Freeze,omitempty" xml:"Freeze,omitempty"`
	// example:
	//
	// FREEZED
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// true
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// dhT20X2310
	JobName            *string                                                 `json:"JobName,omitempty" xml:"JobName,omitempty"`
	LabelDetails       []*GetOssCheckFreezeResultResponseBodyItemsLabelDetails `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	Labels             []*string                                               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Labels2            []*string                                               `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	ManualFreezeAction *string                                                 `json:"ManualFreezeAction,omitempty" xml:"ManualFreezeAction,omitempty"`
	ManualOperateTime  *string                                                 `json:"ManualOperateTime,omitempty" xml:"ManualOperateTime,omitempty"`
	ManualOperator     *string                                                 `json:"ManualOperator,omitempty" xml:"ManualOperator,omitempty"`
	// example:
	//
	// 54416c9b159df4a60ae03c04ccb94cb5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1713014531569_958.png.jpeg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// example:
	//
	// F0A594BB-FA7A-580F-AE9E-A4188E092823
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// low
	RiskLevel0 *string `json:"RiskLevel0,omitempty" xml:"RiskLevel0,omitempty"`
	// example:
	//
	// none
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// audio_media_detection_01
	ServiceCode       *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName       *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SysDisposalStatus *string `json:"SysDisposalStatus,omitempty" xml:"SysDisposalStatus,omitempty"`
	// example:
	//
	// P_BT3FHS
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetOssCheckFreezeResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetContentType() *string {
	return s.ContentType
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetFeedback() *string {
	return s.Feedback
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetFreeze() *bool {
	return s.Freeze
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetFreezeStatus() *string {
	return s.FreezeStatus
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetJobName() *string {
	return s.JobName
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetLabelDetails() []*GetOssCheckFreezeResultResponseBodyItemsLabelDetails {
	return s.LabelDetails
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetLabels() []*string {
	return s.Labels
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetLabels2() []*string {
	return s.Labels2
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetManualFreezeAction() *string {
	return s.ManualFreezeAction
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetManualOperateTime() *string {
	return s.ManualOperateTime
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetManualOperator() *string {
	return s.ManualOperator
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetMd5() *string {
	return s.Md5
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetMsg() *string {
	return s.Msg
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetObject() *string {
	return s.Object
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetRiskLevel0() *string {
	return s.RiskLevel0
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetRiskLevel2() *string {
	return s.RiskLevel2
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetSysDisposalStatus() *string {
	return s.SysDisposalStatus
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOssCheckFreezeResultResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetBucket(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetCode(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Code = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetContentType(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetCopyFrom(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.CopyFrom = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetFeedback(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Feedback = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetFreeze(v bool) *GetOssCheckFreezeResultResponseBodyItems {
	s.Freeze = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetFreezeStatus(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.FreezeStatus = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetFreezeType(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.FreezeType = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetImageUrl(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetIsCopy(v bool) *GetOssCheckFreezeResultResponseBodyItems {
	s.IsCopy = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetJobName(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.JobName = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetLabelDetails(v []*GetOssCheckFreezeResultResponseBodyItemsLabelDetails) *GetOssCheckFreezeResultResponseBodyItems {
	s.LabelDetails = v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetLabels(v []*string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Labels = v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetLabels2(v []*string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Labels2 = v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetManualFreezeAction(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ManualFreezeAction = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetManualOperateTime(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ManualOperateTime = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetManualOperator(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ManualOperator = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetMd5(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Md5 = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetMsg(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Msg = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetObject(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Object = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetRequestId(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetRiskLevel(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetRiskLevel0(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.RiskLevel0 = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetRiskLevel2(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.RiskLevel2 = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetScanResult(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ScanResult = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetServiceCode(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetServiceName(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetSysDisposalStatus(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.SysDisposalStatus = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetTaskId(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) SetUrl(v string) *GetOssCheckFreezeResultResponseBodyItems {
	s.Url = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckFreezeResultResponseBodyItemsLabelDetails struct {
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string  `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetOssCheckFreezeResultResponseBodyItemsLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckFreezeResultResponseBodyItemsLabelDetails) GoString() string {
	return s.String()
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) GetDescription() *string {
	return s.Description
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) GetLabel() *string {
	return s.Label
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) SetConfidence(v float32) *GetOssCheckFreezeResultResponseBodyItemsLabelDetails {
	s.Confidence = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) SetDescription(v string) *GetOssCheckFreezeResultResponseBodyItemsLabelDetails {
	s.Description = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) SetLabel(v string) *GetOssCheckFreezeResultResponseBodyItemsLabelDetails {
	s.Label = &v
	return s
}

func (s *GetOssCheckFreezeResultResponseBodyItemsLabelDetails) Validate() error {
	return dara.Validate(s)
}
