// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOssCheckResultResponseBody
	GetCurrentPage() *int32
	SetItems(v []*ListOssCheckResultResponseBodyItems) *ListOssCheckResultResponseBody
	GetItems() []*ListOssCheckResultResponseBodyItems
	SetPageSize(v int32) *ListOssCheckResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOssCheckResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOssCheckResultResponseBody
	GetTotalCount() *int64
}

type ListOssCheckResultResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data on the current page.
	Items []*ListOssCheckResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 13
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOssCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssCheckResultResponseBody) GetItems() []*ListOssCheckResultResponseBodyItems {
	return s.Items
}

func (s *ListOssCheckResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOssCheckResultResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOssCheckResultResponseBody) SetCurrentPage(v int32) *ListOssCheckResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetItems(v []*ListOssCheckResultResponseBodyItems) *ListOssCheckResultResponseBody {
	s.Items = v
	return s
}

func (s *ListOssCheckResultResponseBody) SetPageSize(v int32) *ListOssCheckResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetRequestId(v string) *ListOssCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOssCheckResultResponseBody) SetTotalCount(v int64) *ListOssCheckResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOssCheckResultResponseBody) Validate() error {
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

type ListOssCheckResultResponseBodyItems struct {
	// The OSS bucket.
	//
	// example:
	//
	// tmp
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The error code, which is consistent with the HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The audio and video detection type.
	//
	// example:
	//
	// audio
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The primary service.
	//
	// example:
	//
	// audio_media_detection
	CopyFrom *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	// The freeze status.
	//
	// example:
	//
	// FREEZED
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	// The freeze type.
	//
	// example:
	//
	// ACL
	FreezeType *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// http://www.aliyuncs.com/test.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Indicates whether the task is copied.
	//
	// example:
	//
	// true
	IsCopy *bool `json:"IsCopy,omitempty" xml:"IsCopy,omitempty"`
	// The task name.
	//
	// example:
	//
	// dhT20X2310
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The list of labels hit by video frames.
	LabelDetails []*ListOssCheckResultResponseBodyItemsLabelDetails `json:"LabelDetails,omitempty" xml:"LabelDetails,omitempty" type:"Repeated"`
	// The list of labels hit by video audio.
	LabelDetails2 []*ListOssCheckResultResponseBodyItemsLabelDetails2 `json:"LabelDetails2,omitempty" xml:"LabelDetails2,omitempty" type:"Repeated"`
	// The image labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The text labels.
	Labels2 []*string `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	// The MD5 hash of the file.
	//
	// example:
	//
	// 54416c9b159df4a60ae03c04ccb94cb5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The object name.
	//
	// example:
	//
	// 1713014531569_958.png.jpeg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The image risk level.
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The overall risk level.
	//
	// example:
	//
	// low
	RiskLevel0 *string `json:"RiskLevel0,omitempty" xml:"RiskLevel0,omitempty"`
	// The text risk level.
	//
	// example:
	//
	// none
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// The scan result details.
	//
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// The service code.
	//
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The service name.
	//
	// example:
	//
	// 服务名称
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// EP6TI7_au_Zo25ITvCbkocNuF801QOQX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task URL.
	//
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListOssCheckResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *ListOssCheckResultResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *ListOssCheckResultResponseBodyItems) GetContentType() *string {
	return s.ContentType
}

func (s *ListOssCheckResultResponseBodyItems) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *ListOssCheckResultResponseBodyItems) GetFreezeStatus() *string {
	return s.FreezeStatus
}

func (s *ListOssCheckResultResponseBodyItems) GetFreezeType() *string {
	return s.FreezeType
}

func (s *ListOssCheckResultResponseBodyItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListOssCheckResultResponseBodyItems) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *ListOssCheckResultResponseBodyItems) GetJobName() *string {
	return s.JobName
}

func (s *ListOssCheckResultResponseBodyItems) GetLabelDetails() []*ListOssCheckResultResponseBodyItemsLabelDetails {
	return s.LabelDetails
}

func (s *ListOssCheckResultResponseBodyItems) GetLabelDetails2() []*ListOssCheckResultResponseBodyItemsLabelDetails2 {
	return s.LabelDetails2
}

func (s *ListOssCheckResultResponseBodyItems) GetLabels() []*string {
	return s.Labels
}

func (s *ListOssCheckResultResponseBodyItems) GetLabels2() []*string {
	return s.Labels2
}

func (s *ListOssCheckResultResponseBodyItems) GetMd5() *string {
	return s.Md5
}

func (s *ListOssCheckResultResponseBodyItems) GetMsg() *string {
	return s.Msg
}

func (s *ListOssCheckResultResponseBodyItems) GetObject() *string {
	return s.Object
}

func (s *ListOssCheckResultResponseBodyItems) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListOssCheckResultResponseBodyItems) GetRiskLevel0() *string {
	return s.RiskLevel0
}

func (s *ListOssCheckResultResponseBodyItems) GetRiskLevel2() *string {
	return s.RiskLevel2
}

func (s *ListOssCheckResultResponseBodyItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *ListOssCheckResultResponseBodyItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListOssCheckResultResponseBodyItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListOssCheckResultResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOssCheckResultResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *ListOssCheckResultResponseBodyItems) SetBucket(v string) *ListOssCheckResultResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetCode(v string) *ListOssCheckResultResponseBodyItems {
	s.Code = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetContentType(v string) *ListOssCheckResultResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetCopyFrom(v string) *ListOssCheckResultResponseBodyItems {
	s.CopyFrom = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetFreezeStatus(v string) *ListOssCheckResultResponseBodyItems {
	s.FreezeStatus = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetFreezeType(v string) *ListOssCheckResultResponseBodyItems {
	s.FreezeType = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetImageUrl(v string) *ListOssCheckResultResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetIsCopy(v bool) *ListOssCheckResultResponseBodyItems {
	s.IsCopy = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetJobName(v string) *ListOssCheckResultResponseBodyItems {
	s.JobName = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabelDetails(v []*ListOssCheckResultResponseBodyItemsLabelDetails) *ListOssCheckResultResponseBodyItems {
	s.LabelDetails = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabelDetails2(v []*ListOssCheckResultResponseBodyItemsLabelDetails2) *ListOssCheckResultResponseBodyItems {
	s.LabelDetails2 = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabels(v []*string) *ListOssCheckResultResponseBodyItems {
	s.Labels = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetLabels2(v []*string) *ListOssCheckResultResponseBodyItems {
	s.Labels2 = v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetMd5(v string) *ListOssCheckResultResponseBodyItems {
	s.Md5 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetMsg(v string) *ListOssCheckResultResponseBodyItems {
	s.Msg = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetObject(v string) *ListOssCheckResultResponseBodyItems {
	s.Object = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel0(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel0 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetRiskLevel2(v string) *ListOssCheckResultResponseBodyItems {
	s.RiskLevel2 = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetScanResult(v string) *ListOssCheckResultResponseBodyItems {
	s.ScanResult = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetServiceCode(v string) *ListOssCheckResultResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetServiceName(v string) *ListOssCheckResultResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetTaskId(v string) *ListOssCheckResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) SetUrl(v string) *ListOssCheckResultResponseBodyItems {
	s.Url = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItems) Validate() error {
	if s.LabelDetails != nil {
		for _, item := range s.LabelDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LabelDetails2 != nil {
		for _, item := range s.LabelDetails2 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOssCheckResultResponseBodyItemsLabelDetails struct {
	// The description of the label.
	//
	// example:
	//
	// 影音娱乐类
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label hit by the video frame.
	//
	// example:
	//
	// logo_streaming
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ListOssCheckResultResponseBodyItemsLabelDetails) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultResponseBodyItemsLabelDetails) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails) GetDescription() *string {
	return s.Description
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails) GetLabel() *string {
	return s.Label
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails) SetDescription(v string) *ListOssCheckResultResponseBodyItemsLabelDetails {
	s.Description = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails) SetLabel(v string) *ListOssCheckResultResponseBodyItemsLabelDetails {
	s.Label = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails) Validate() error {
	return dara.Validate(s)
}

type ListOssCheckResultResponseBodyItemsLabelDetails2 struct {
	// The description of the label.
	//
	// example:
	//
	// 辱骂内容
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label hit by the audio.
	//
	// example:
	//
	// abuse
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ListOssCheckResultResponseBodyItemsLabelDetails2) String() string {
	return dara.Prettify(s)
}

func (s ListOssCheckResultResponseBodyItemsLabelDetails2) GoString() string {
	return s.String()
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails2) GetDescription() *string {
	return s.Description
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails2) GetLabel() *string {
	return s.Label
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails2) SetDescription(v string) *ListOssCheckResultResponseBodyItemsLabelDetails2 {
	s.Description = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails2) SetLabel(v string) *ListOssCheckResultResponseBodyItemsLabelDetails2 {
	s.Label = &v
	return s
}

func (s *ListOssCheckResultResponseBodyItemsLabelDetails2) Validate() error {
	return dara.Validate(s)
}
