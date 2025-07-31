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
	// example:
	//
	// 1
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*ListOssCheckResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ListOssCheckResultResponseBodyItems struct {
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
	CopyFrom     *string `json:"CopyFrom,omitempty" xml:"CopyFrom,omitempty"`
	FreezeStatus *string `json:"FreezeStatus,omitempty" xml:"FreezeStatus,omitempty"`
	FreezeType   *string `json:"FreezeType,omitempty" xml:"FreezeType,omitempty"`
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
	JobName *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Labels  []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Labels2 []*string `json:"Labels2,omitempty" xml:"Labels2,omitempty" type:"Repeated"`
	// example:
	//
	// 54416c9b159df4a60ae03c04ccb94cb5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1713014531569_958.png.jpeg
	Object     *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RiskLevel  *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskLevel0 *string `json:"RiskLevel0,omitempty" xml:"RiskLevel0,omitempty"`
	RiskLevel2 *string `json:"RiskLevel2,omitempty" xml:"RiskLevel2,omitempty"`
	// example:
	//
	// {}
	ScanResult *string `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	// example:
	//
	// audio_media_detection_01
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// EP6TI7_au_Zo25ITvCbkocNuF801QOQX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	return dara.Validate(s)
}
