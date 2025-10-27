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
	// The data returned.
	Data []*ListObjectScanEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListObjectScanEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListObjectScanEventResponseBodyData struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// hz-new01****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The details of the file.
	Details []*ListObjectScanEventResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// Indicates whether the file can be detected by cloud sandbox. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DisplaySandboxResult *string `json:"DisplaySandboxResult,omitempty" xml:"DisplaySandboxResult,omitempty"`
	// The ID of the alert.
	//
	// example:
	//
	// 911273
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// WebShell
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The path to the file.
	//
	// example:
	//
	// /usr/local****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The timestamp at which the alert was first detected.
	//
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether an alert is generated for the file extracted from the package. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasSubEvent *bool `json:"HasSubEvent,omitempty" xml:"HasSubEvent,omitempty"`
	// The timestamp at which the alert was last detected.
	//
	// example:
	//
	// 1694576692000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The MD5 hash value of the file.
	//
	// example:
	//
	// 5b394b54ca632fe51c4ab4a6dbaf****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The key of the file that is stored in the OSS bucket.
	//
	// example:
	//
	// 1/2023/07/21/10/18/16899059356518bcf6c64-a04e-492d-a421-4ae8b888****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The risk level of the alert. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The SHA-1 hash value of the file.
	//
	// example:
	//
	// 3c01bdbb26f358bab27f267924aa2c9a03fc****
	Sha1 *string `json:"Sha1,omitempty" xml:"Sha1,omitempty"`
	// The SHA-256 hash value of the file.
	//
	// example:
	//
	// 3a6fed5fc11392b3ee9f81caf017b48640d7458766a8eb0382899a605b41****
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
	// The method that is used to detect the malicious file. Valid values:
	//
	// 	- **API**: uses API operations.
	//
	// 	- **OSS**: uses OSS file check.
	//
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
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListObjectScanEventResponseBodyDataDetails struct {
	// The name of the parameter in the file details.
	//
	// example:
	//
	// DownloadUrl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the alert.
	//
	// example:
	//
	// DownloadUrl
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The value type of the parameter in the file details.
	//
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// http://gcx.cn-hangzhou.aliyuncs.com/****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The value of the parameter.
	//
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
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 253
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
