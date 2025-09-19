// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListByBuildRiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageListByBuildRiskResponseBody
	GetCode() *string
	SetData(v *DescribeImageListByBuildRiskResponseBodyData) *DescribeImageListByBuildRiskResponseBody
	GetData() *DescribeImageListByBuildRiskResponseBodyData
	SetMessage(v string) *DescribeImageListByBuildRiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeImageListByBuildRiskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageListByBuildRiskResponseBody
	GetSuccess() *bool
}

type DescribeImageListByBuildRiskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeImageListByBuildRiskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 52870893-48A7-5A9E-9E05-6253E5B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeImageListByBuildRiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageListByBuildRiskResponseBody) GetData() *DescribeImageListByBuildRiskResponseBodyData {
	return s.Data
}

func (s *DescribeImageListByBuildRiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageListByBuildRiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageListByBuildRiskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageListByBuildRiskResponseBody) SetCode(v string) *DescribeImageListByBuildRiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBody) SetData(v *DescribeImageListByBuildRiskResponseBodyData) *DescribeImageListByBuildRiskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBody) SetMessage(v string) *DescribeImageListByBuildRiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBody) SetRequestId(v string) *DescribeImageListByBuildRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBody) SetSuccess(v bool) *DescribeImageListByBuildRiskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListByBuildRiskResponseBodyData struct {
	// The images.
	List []*DescribeImageListByBuildRiskResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageListByBuildRiskResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s DescribeImageListByBuildRiskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskResponseBodyData) GetList() []*DescribeImageListByBuildRiskResponseBodyDataList {
	return s.List
}

func (s *DescribeImageListByBuildRiskResponseBodyData) GetPageInfo() *DescribeImageListByBuildRiskResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeImageListByBuildRiskResponseBodyData) SetList(v []*DescribeImageListByBuildRiskResponseBodyDataList) *DescribeImageListByBuildRiskResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyData) SetPageInfo(v *DescribeImageListByBuildRiskResponseBodyDataPageInfo) *DescribeImageListByBuildRiskResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListByBuildRiskResponseBodyDataList struct {
	// The digest value of the image.
	//
	// example:
	//
	// a5ccdd9b166b67e02954aa9b618fe19b7968bd56a15463d2ad7f2643ba5b****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1723710827000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The instance ID of the image repository.
	//
	// example:
	//
	// 39010****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1723710827999
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test-tepo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// namespace
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
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the alert event. Valid values:
	//
	// 	- **0**: unhandled.
	//
	// 	- **1**: ignored.
	//
	// 	- **2**: false positive.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 6ec898e6274f942e0e4a053eff1c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeImageListByBuildRiskResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetDigest(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.Digest = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetFirstScanTime(v int64) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetInstanceId(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetLastScanTime(v int64) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetRegionId(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetRepoName(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.RepoName = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetRepoNamespace(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetRiskLevel(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetStatus(v int32) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetTag(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.Tag = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) SetUuid(v string) *DescribeImageListByBuildRiskResponseBodyDataList {
	s.Uuid = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageListByBuildRiskResponseBodyDataPageInfo struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 109
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageListByBuildRiskResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeImageListByBuildRiskResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeImageListByBuildRiskResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) SetTotalCount(v int32) *DescribeImageListByBuildRiskResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
