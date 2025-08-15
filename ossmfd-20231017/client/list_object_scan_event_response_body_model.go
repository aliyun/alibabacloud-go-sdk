// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectScanEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListObjectScanEventResponseBodyData) *ListObjectScanEventResponseBody
	GetData() []*ListObjectScanEventResponseBodyData
	SetPageInfo(v *ListObjectScanEventResponseBodyPageInfo) *ListObjectScanEventResponseBody
	GetPageInfo() *ListObjectScanEventResponseBodyPageInfo
	SetRequestId(v string) *ListObjectScanEventResponseBody
	GetRequestId() *string
}

type ListObjectScanEventResponseBody struct {
	Data     []*ListObjectScanEventResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListObjectScanEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// example:
	//
	// E99D386F-5546-5BCD-BADD-F4EF4B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListObjectScanEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventResponseBody) GetData() []*ListObjectScanEventResponseBodyData {
	return s.Data
}

func (s *ListObjectScanEventResponseBody) GetPageInfo() *ListObjectScanEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListObjectScanEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListObjectScanEventResponseBody) SetData(v []*ListObjectScanEventResponseBodyData) *ListObjectScanEventResponseBody {
	s.Data = v
	return s
}

func (s *ListObjectScanEventResponseBody) SetPageInfo(v *ListObjectScanEventResponseBodyPageInfo) *ListObjectScanEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListObjectScanEventResponseBody) SetRequestId(v string) *ListObjectScanEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListObjectScanEventResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListObjectScanEventResponseBodyData struct {
	// example:
	//
	// testBucket******
	BucketName *string                                       `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Details    []*ListObjectScanEventResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DisplaySandboxResult *string `json:"DisplaySandboxResult,omitempty" xml:"DisplaySandboxResult,omitempty"`
	// example:
	//
	// 1
	EventId   *int64  `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// /usr/local****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// true
	HasSubEvent *bool `json:"HasSubEvent,omitempty" xml:"HasSubEvent,omitempty"`
	// example:
	//
	// 1694576692000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// e991e62f484bb6accd676e57a9******
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 3c01bdbb26f358bab27f267924aa2c9a03fc****
	Sha1 *string `json:"Sha1,omitempty" xml:"Sha1,omitempty"`
	// example:
	//
	// 3a6fed5fc11392b3ee9f81caf017b48640d7458766a8eb0382899a605b41****
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListObjectScanEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *ListObjectScanEventResponseBodyData) GetDetails() []*ListObjectScanEventResponseBodyDataDetails {
	return s.Details
}

func (s *ListObjectScanEventResponseBodyData) GetDisplaySandboxResult() *string {
	return s.DisplaySandboxResult
}

func (s *ListObjectScanEventResponseBodyData) GetEventId() *int64 {
	return s.EventId
}

func (s *ListObjectScanEventResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *ListObjectScanEventResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *ListObjectScanEventResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListObjectScanEventResponseBodyData) GetHasSubEvent() *bool {
	return s.HasSubEvent
}

func (s *ListObjectScanEventResponseBodyData) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListObjectScanEventResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *ListObjectScanEventResponseBodyData) GetOssKey() *string {
	return s.OssKey
}

func (s *ListObjectScanEventResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListObjectScanEventResponseBodyData) GetSha1() *string {
	return s.Sha1
}

func (s *ListObjectScanEventResponseBodyData) GetSha256() *string {
	return s.Sha256
}

func (s *ListObjectScanEventResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListObjectScanEventResponseBodyData) SetBucketName(v string) *ListObjectScanEventResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetDetails(v []*ListObjectScanEventResponseBodyDataDetails) *ListObjectScanEventResponseBodyData {
	s.Details = v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetDisplaySandboxResult(v string) *ListObjectScanEventResponseBodyData {
	s.DisplaySandboxResult = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetEventId(v int64) *ListObjectScanEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetEventName(v string) *ListObjectScanEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetFilePath(v string) *ListObjectScanEventResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetFirstTime(v int64) *ListObjectScanEventResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetHasSubEvent(v bool) *ListObjectScanEventResponseBodyData {
	s.HasSubEvent = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetLastTime(v int64) *ListObjectScanEventResponseBodyData {
	s.LastTime = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetMd5(v string) *ListObjectScanEventResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetOssKey(v string) *ListObjectScanEventResponseBodyData {
	s.OssKey = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetRiskLevel(v string) *ListObjectScanEventResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetSha1(v string) *ListObjectScanEventResponseBodyData {
	s.Sha1 = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetSha256(v string) *ListObjectScanEventResponseBodyData {
	s.Sha256 = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) SetSource(v string) *ListObjectScanEventResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListObjectScanEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListObjectScanEventResponseBodyDataDetails struct {
	// example:
	//
	// DownloadUrl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DownloadUrl
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// http://gcx.cn-hangzhou.aliyuncs.com/****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// http://gcx.cn-hangzhou.aliyuncs.com/****
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s ListObjectScanEventResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventResponseBodyDataDetails) GetName() *string {
	return s.Name
}

func (s *ListObjectScanEventResponseBodyDataDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *ListObjectScanEventResponseBodyDataDetails) GetType() *string {
	return s.Type
}

func (s *ListObjectScanEventResponseBodyDataDetails) GetValue() *string {
	return s.Value
}

func (s *ListObjectScanEventResponseBodyDataDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *ListObjectScanEventResponseBodyDataDetails) SetName(v string) *ListObjectScanEventResponseBodyDataDetails {
	s.Name = &v
	return s
}

func (s *ListObjectScanEventResponseBodyDataDetails) SetNameDisplay(v string) *ListObjectScanEventResponseBodyDataDetails {
	s.NameDisplay = &v
	return s
}

func (s *ListObjectScanEventResponseBodyDataDetails) SetType(v string) *ListObjectScanEventResponseBodyDataDetails {
	s.Type = &v
	return s
}

func (s *ListObjectScanEventResponseBodyDataDetails) SetValue(v string) *ListObjectScanEventResponseBodyDataDetails {
	s.Value = &v
	return s
}

func (s *ListObjectScanEventResponseBodyDataDetails) SetValueDisplay(v string) *ListObjectScanEventResponseBodyDataDetails {
	s.ValueDisplay = &v
	return s
}

func (s *ListObjectScanEventResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}

type ListObjectScanEventResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 187
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListObjectScanEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListObjectScanEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListObjectScanEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListObjectScanEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListObjectScanEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListObjectScanEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListObjectScanEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListObjectScanEventResponseBodyPageInfo) SetPageSize(v int32) *ListObjectScanEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListObjectScanEventResponseBodyPageInfo) SetTotalCount(v int32) *ListObjectScanEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListObjectScanEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
