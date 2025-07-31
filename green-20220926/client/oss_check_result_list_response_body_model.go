// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssCheckResultListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthStatus(v string) *OssCheckResultListResponseBody
	GetAuthStatus() *string
	SetCurrentPage(v int32) *OssCheckResultListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*OssCheckResultListResponseBodyItems) *OssCheckResultListResponseBody
	GetItems() []*OssCheckResultListResponseBodyItems
	SetPageSize(v int32) *OssCheckResultListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *OssCheckResultListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *OssCheckResultListResponseBody
	GetTotalCount() *int32
}

type OssCheckResultListResponseBody struct {
	// example:
	//
	// OK
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*OssCheckResultListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s OssCheckResultListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OssCheckResultListResponseBody) GoString() string {
	return s.String()
}

func (s *OssCheckResultListResponseBody) GetAuthStatus() *string {
	return s.AuthStatus
}

func (s *OssCheckResultListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OssCheckResultListResponseBody) GetItems() []*OssCheckResultListResponseBodyItems {
	return s.Items
}

func (s *OssCheckResultListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OssCheckResultListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OssCheckResultListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *OssCheckResultListResponseBody) SetAuthStatus(v string) *OssCheckResultListResponseBody {
	s.AuthStatus = &v
	return s
}

func (s *OssCheckResultListResponseBody) SetCurrentPage(v int32) *OssCheckResultListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *OssCheckResultListResponseBody) SetItems(v []*OssCheckResultListResponseBodyItems) *OssCheckResultListResponseBody {
	s.Items = v
	return s
}

func (s *OssCheckResultListResponseBody) SetPageSize(v int32) *OssCheckResultListResponseBody {
	s.PageSize = &v
	return s
}

func (s *OssCheckResultListResponseBody) SetRequestId(v string) *OssCheckResultListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OssCheckResultListResponseBody) SetTotalCount(v int32) *OssCheckResultListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *OssCheckResultListResponseBody) Validate() error {
	return dara.Validate(s)
}

type OssCheckResultListResponseBodyItems struct {
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
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1713014531569_958.png.jpeg
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
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
	// P_XHDUS
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s OssCheckResultListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s OssCheckResultListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *OssCheckResultListResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *OssCheckResultListResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *OssCheckResultListResponseBodyItems) GetContentType() *string {
	return s.ContentType
}

func (s *OssCheckResultListResponseBodyItems) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *OssCheckResultListResponseBodyItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *OssCheckResultListResponseBodyItems) GetIsCopy() *bool {
	return s.IsCopy
}

func (s *OssCheckResultListResponseBodyItems) GetJobName() *string {
	return s.JobName
}

func (s *OssCheckResultListResponseBodyItems) GetLabels() []*string {
	return s.Labels
}

func (s *OssCheckResultListResponseBodyItems) GetLabels2() []*string {
	return s.Labels2
}

func (s *OssCheckResultListResponseBodyItems) GetMd5() *string {
	return s.Md5
}

func (s *OssCheckResultListResponseBodyItems) GetMsg() *string {
	return s.Msg
}

func (s *OssCheckResultListResponseBodyItems) GetObject() *string {
	return s.Object
}

func (s *OssCheckResultListResponseBodyItems) GetScanResult() *string {
	return s.ScanResult
}

func (s *OssCheckResultListResponseBodyItems) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *OssCheckResultListResponseBodyItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *OssCheckResultListResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *OssCheckResultListResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *OssCheckResultListResponseBodyItems) SetBucket(v string) *OssCheckResultListResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetCode(v string) *OssCheckResultListResponseBodyItems {
	s.Code = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetContentType(v string) *OssCheckResultListResponseBodyItems {
	s.ContentType = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetCopyFrom(v string) *OssCheckResultListResponseBodyItems {
	s.CopyFrom = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetImageUrl(v string) *OssCheckResultListResponseBodyItems {
	s.ImageUrl = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetIsCopy(v bool) *OssCheckResultListResponseBodyItems {
	s.IsCopy = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetJobName(v string) *OssCheckResultListResponseBodyItems {
	s.JobName = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetLabels(v []*string) *OssCheckResultListResponseBodyItems {
	s.Labels = v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetLabels2(v []*string) *OssCheckResultListResponseBodyItems {
	s.Labels2 = v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetMd5(v string) *OssCheckResultListResponseBodyItems {
	s.Md5 = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetMsg(v string) *OssCheckResultListResponseBodyItems {
	s.Msg = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetObject(v string) *OssCheckResultListResponseBodyItems {
	s.Object = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetScanResult(v string) *OssCheckResultListResponseBodyItems {
	s.ScanResult = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetServiceCode(v string) *OssCheckResultListResponseBodyItems {
	s.ServiceCode = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetServiceName(v string) *OssCheckResultListResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetTaskId(v string) *OssCheckResultListResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) SetUrl(v string) *OssCheckResultListResponseBodyItems {
	s.Url = &v
	return s
}

func (s *OssCheckResultListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
