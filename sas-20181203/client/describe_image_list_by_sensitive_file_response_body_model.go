// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListBySensitiveFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageListBySensitiveFileResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeImageListBySensitiveFileResponseBody
	GetHttpStatusCode() *int32
	SetImageInfos(v []*DescribeImageListBySensitiveFileResponseBodyImageInfos) *DescribeImageListBySensitiveFileResponseBody
	GetImageInfos() []*DescribeImageListBySensitiveFileResponseBodyImageInfos
	SetMessage(v string) *DescribeImageListBySensitiveFileResponseBody
	GetMessage() *string
	SetPageInfo(v *DescribeImageListBySensitiveFileResponseBodyPageInfo) *DescribeImageListBySensitiveFileResponseBody
	GetPageInfo() *DescribeImageListBySensitiveFileResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageListBySensitiveFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageListBySensitiveFileResponseBody
	GetSuccess() *bool
}

type DescribeImageListBySensitiveFileResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The information about the images.
	ImageInfos []*DescribeImageListBySensitiveFileResponseBodyImageInfos `json:"ImageInfos,omitempty" xml:"ImageInfos,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *DescribeImageListBySensitiveFileResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E0C5C07F-1576-509A-AE44-1C36B8445B37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageListBySensitiveFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListBySensitiveFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetImageInfos() []*DescribeImageListBySensitiveFileResponseBodyImageInfos {
	return s.ImageInfos
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetPageInfo() *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageListBySensitiveFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetCode(v string) *DescribeImageListBySensitiveFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetHttpStatusCode(v int32) *DescribeImageListBySensitiveFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetImageInfos(v []*DescribeImageListBySensitiveFileResponseBodyImageInfos) *DescribeImageListBySensitiveFileResponseBody {
	s.ImageInfos = v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetMessage(v string) *DescribeImageListBySensitiveFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetPageInfo(v *DescribeImageListBySensitiveFileResponseBodyPageInfo) *DescribeImageListBySensitiveFileResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetRequestId(v string) *DescribeImageListBySensitiveFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) SetSuccess(v bool) *DescribeImageListBySensitiveFileResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListBySensitiveFileResponseBodyImageInfos struct {
	// The image digest.
	//
	// example:
	//
	// v005
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649814050000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The image instance ID.
	//
	// example:
	//
	// cri-a595qp31knh9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1649814050000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The ID of the region in which the image instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// opa-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// libssh2
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sensitive file status. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: ignored
	//
	// 	- **2**: false positive
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag that is added to the image.
	//
	// example:
	//
	// nuxeo6
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The image UUID.
	//
	// example:
	//
	// f58681174f944623345379e23b7b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeImageListBySensitiveFileResponseBodyImageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListBySensitiveFileResponseBodyImageInfos) GoString() string {
	return s.String()
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetDigest(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.Digest = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetFirstScanTime(v int64) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetInstanceId(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetLastScanTime(v int64) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetRegionId(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetRepoName(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.RepoName = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetRepoNamespace(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetRiskLevel(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetStatus(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.Status = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetTag(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.Tag = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) SetUuid(v string) *DescribeImageListBySensitiveFileResponseBodyImageInfos {
	s.Uuid = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyImageInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListBySensitiveFileResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The key of the last data entry.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAGYXFWIAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM1Mzc3Njc4MzA2ODY5NmI2YTY1Nzg2NTcxNjE2NDc4NjE=
	LastRowKey *string `json:"LastRowKey,omitempty" xml:"LastRowKey,omitempty"`
	// The number of entries returned per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 83
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageListBySensitiveFileResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListBySensitiveFileResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) GetLastRowKey() *string {
	return s.LastRowKey
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) SetCount(v int32) *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) SetLastRowKey(v string) *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	s.LastRowKey = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageListBySensitiveFileResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageListBySensitiveFileResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
