// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckResultDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetOssCheckResultDetailResponseBody
	GetCode() *int32
	SetData(v *GetOssCheckResultDetailResponseBodyData) *GetOssCheckResultDetailResponseBody
	GetData() *GetOssCheckResultDetailResponseBodyData
	SetMsg(v string) *GetOssCheckResultDetailResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetOssCheckResultDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOssCheckResultDetailResponseBody
	GetSuccess() *bool
}

type GetOssCheckResultDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOssCheckResultDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 62E97001-1255-50A9-8E1E-4FD05473D952
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOssCheckResultDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetOssCheckResultDetailResponseBody) GetData() *GetOssCheckResultDetailResponseBodyData {
	return s.Data
}

func (s *GetOssCheckResultDetailResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetOssCheckResultDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckResultDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOssCheckResultDetailResponseBody) SetCode(v int32) *GetOssCheckResultDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBody) SetData(v *GetOssCheckResultDetailResponseBodyData) *GetOssCheckResultDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetOssCheckResultDetailResponseBody) SetMsg(v string) *GetOssCheckResultDetailResponseBody {
	s.Msg = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBody) SetRequestId(v string) *GetOssCheckResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBody) SetSuccess(v bool) *GetOssCheckResultDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckResultDetailResponseBodyData struct {
	// example:
	//
	// oss-tmp
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
	// example:
	//
	// UNFREEZED
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	// example:
	//
	// COPY
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
	LabelDetails       []*GetOssCheckResultDetailResponseBodyDataLabelDetails  `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	LabelDetails2      []*GetOssCheckResultDetailResponseBodyDataLabelDetails2 `json:"LabelDetails2,omitempty" xml:"LabelDetails2,omitempty" type:"Repeated"`
	Labels             []*string                                               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Labels2            []*string                                               `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	ManualFreezeAction *string                                                 `json:"ManualFreezeAction,omitempty" xml:"ManualFreezeAction,omitempty"`
	ManualOperateTime  *string                                                 `json:"ManualOperateTime,omitempty" xml:"ManualOperateTime,omitempty"`
	ManualOperator     *string                                                 `json:"ManualOperator,omitempty" xml:"ManualOperator,omitempty"`
	// example:
	//
	// f6e2e1946f06310c8a0cc443a05819f3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1748396909030.jpg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	// medium
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// example:
	//
	// {}
	ScanResult       *string                                                    `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	ScanServiceInfos []*GetOssCheckResultDetailResponseBodyDataScanServiceInfos `json:"ScanServiceInfos,omitempty" xml:"ScanServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// P_Z7OLMN
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetOssCheckResultDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssCheckResultDetailResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetOssCheckResultDetailResponseBodyData) GetContentType() *string {
	return s.ContentType
}

func (s *GetOssCheckResultDetailResponseBodyData) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *GetOssCheckResultDetailResponseBodyData) GetFreezeStatus() *string {
	return s.FreezeStatus
}

func (s *GetOssCheckResultDetailResponseBodyData) GetFreezeType() *string {
	return s.FreezeType
}

func (s *GetOssCheckResultDetailResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetOssCheckResultDetailResponseBodyData) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *GetOssCheckResultDetailResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetOssCheckResultDetailResponseBodyData) GetLabelDetails() []*GetOssCheckResultDetailResponseBodyDataLabelDetails {
	return s.LabelDetails
}

func (s *GetOssCheckResultDetailResponseBodyData) GetLabelDetails2() []*GetOssCheckResultDetailResponseBodyDataLabelDetails2 {
	return s.LabelDetails2
}

func (s *GetOssCheckResultDetailResponseBodyData) GetLabels() []*string {
	return s.Labels
}

func (s *GetOssCheckResultDetailResponseBodyData) GetLabels2() []*string {
	return s.Labels2
}

func (s *GetOssCheckResultDetailResponseBodyData) GetManualFreezeAction() *string {
	return s.ManualFreezeAction
}

func (s *GetOssCheckResultDetailResponseBodyData) GetManualOperateTime() *string {
	return s.ManualOperateTime
}

func (s *GetOssCheckResultDetailResponseBodyData) GetManualOperator() *string {
	return s.ManualOperator
}

func (s *GetOssCheckResultDetailResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *GetOssCheckResultDetailResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *GetOssCheckResultDetailResponseBodyData) GetObject() *string {
	return s.Object
}

func (s *GetOssCheckResultDetailResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetOssCheckResultDetailResponseBodyData) GetRiskLevel0() *string {
	return s.RiskLevel0
}

func (s *GetOssCheckResultDetailResponseBodyData) GetRiskLevel2() *string {
	return s.RiskLevel2
}

func (s *GetOssCheckResultDetailResponseBodyData) GetScanResult() *string {
	return s.ScanResult
}

func (s *GetOssCheckResultDetailResponseBodyData) GetScanServiceInfos() []*GetOssCheckResultDetailResponseBodyDataScanServiceInfos {
	return s.ScanServiceInfos
}

func (s *GetOssCheckResultDetailResponseBodyData) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetOssCheckResultDetailResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetOssCheckResultDetailResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOssCheckResultDetailResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetOssCheckResultDetailResponseBodyData) SetBucket(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetCode(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetContentType(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ContentType = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetCopyFrom(v string) *GetOssCheckResultDetailResponseBodyData {
	s.CopyFrom = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetFreezeStatus(v string) *GetOssCheckResultDetailResponseBodyData {
	s.FreezeStatus = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetFreezeType(v string) *GetOssCheckResultDetailResponseBodyData {
	s.FreezeType = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetImageUrl(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetIsCopy(v bool) *GetOssCheckResultDetailResponseBodyData {
	s.IsCopy = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetJobName(v string) *GetOssCheckResultDetailResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetLabelDetails(v []*GetOssCheckResultDetailResponseBodyDataLabelDetails) *GetOssCheckResultDetailResponseBodyData {
	s.LabelDetails = v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetLabelDetails2(v []*GetOssCheckResultDetailResponseBodyDataLabelDetails2) *GetOssCheckResultDetailResponseBodyData {
	s.LabelDetails2 = v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetLabels(v []*string) *GetOssCheckResultDetailResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetLabels2(v []*string) *GetOssCheckResultDetailResponseBodyData {
	s.Labels2 = v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetManualFreezeAction(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ManualFreezeAction = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetManualOperateTime(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ManualOperateTime = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetManualOperator(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ManualOperator = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetMd5(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetMsg(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Msg = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetObject(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Object = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetRiskLevel(v string) *GetOssCheckResultDetailResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetRiskLevel0(v string) *GetOssCheckResultDetailResponseBodyData {
	s.RiskLevel0 = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetRiskLevel2(v string) *GetOssCheckResultDetailResponseBodyData {
	s.RiskLevel2 = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetScanResult(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ScanResult = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetScanServiceInfos(v []*GetOssCheckResultDetailResponseBodyDataScanServiceInfos) *GetOssCheckResultDetailResponseBodyData {
	s.ScanServiceInfos = v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetServiceCode(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetServiceName(v string) *GetOssCheckResultDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetTaskId(v string) *GetOssCheckResultDetailResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) SetUrl(v string) *GetOssCheckResultDetailResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckResultDetailResponseBodyDataLabelDetails struct {
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string  `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetOssCheckResultDetailResponseBodyDataLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponseBodyDataLabelDetails) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) GetDescription() *string {
	return s.Description
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) GetLabel() *string {
	return s.Label
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) SetConfidence(v float32) *GetOssCheckResultDetailResponseBodyDataLabelDetails {
	s.Confidence = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) SetDescription(v string) *GetOssCheckResultDetailResponseBodyDataLabelDetails {
	s.Description = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) SetLabel(v string) *GetOssCheckResultDetailResponseBodyDataLabelDetails {
	s.Label = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckResultDetailResponseBodyDataLabelDetails2 struct {
	Confidence  *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string  `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetOssCheckResultDetailResponseBodyDataLabelDetails2) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponseBodyDataLabelDetails2) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) GetDescription() *string {
	return s.Description
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) GetLabel() *string {
	return s.Label
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) SetConfidence(v float32) *GetOssCheckResultDetailResponseBodyDataLabelDetails2 {
	s.Confidence = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) SetDescription(v string) *GetOssCheckResultDetailResponseBodyDataLabelDetails2 {
	s.Description = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) SetLabel(v string) *GetOssCheckResultDetailResponseBodyDataLabelDetails2 {
	s.Label = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataLabelDetails2) Validate() error {
	return dara.Validate(s)
}

type GetOssCheckResultDetailResponseBodyDataScanServiceInfos struct {
	// example:
	//
	// audio_media_detection
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// example:
	//
	// true
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetOssCheckResultDetailResponseBodyDataScanServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckResultDetailResponseBodyDataScanServiceInfos) GoString() string {
	return s.String()
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) SetCopyFrom(v string) *GetOssCheckResultDetailResponseBodyDataScanServiceInfos {
	s.CopyFrom = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) SetIsCopy(v bool) *GetOssCheckResultDetailResponseBodyDataScanServiceInfos {
	s.IsCopy = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) SetServiceCode(v string) *GetOssCheckResultDetailResponseBodyDataScanServiceInfos {
	s.ServiceCode = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) SetServiceName(v string) *GetOssCheckResultDetailResponseBodyDataScanServiceInfos {
	s.ServiceName = &v
	return s
}

func (s *GetOssCheckResultDetailResponseBodyDataScanServiceInfos) Validate() error {
	return dara.Validate(s)
}
