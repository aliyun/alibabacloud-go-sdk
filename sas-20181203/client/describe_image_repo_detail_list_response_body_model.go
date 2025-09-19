// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageRepoResponses(v []*DescribeImageRepoDetailListResponseBodyImageRepoResponses) *DescribeImageRepoDetailListResponseBody
	GetImageRepoResponses() []*DescribeImageRepoDetailListResponseBodyImageRepoResponses
	SetPageInfo(v *DescribeImageRepoDetailListResponseBodyPageInfo) *DescribeImageRepoDetailListResponseBody
	GetPageInfo() *DescribeImageRepoDetailListResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageRepoDetailListResponseBody
	GetRequestId() *string
}

type DescribeImageRepoDetailListResponseBody struct {
	// The information about image repositories.
	ImageRepoResponses []*DescribeImageRepoDetailListResponseBodyImageRepoResponses `json:"ImageRepoResponses,omitempty" xml:"ImageRepoResponses,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageRepoDetailListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageRepoDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoDetailListResponseBody) GetImageRepoResponses() []*DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	return s.ImageRepoResponses
}

func (s *DescribeImageRepoDetailListResponseBody) GetPageInfo() *DescribeImageRepoDetailListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageRepoDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageRepoDetailListResponseBody) SetImageRepoResponses(v []*DescribeImageRepoDetailListResponseBodyImageRepoResponses) *DescribeImageRepoDetailListResponseBody {
	s.ImageRepoResponses = v
	return s
}

func (s *DescribeImageRepoDetailListResponseBody) SetPageInfo(v *DescribeImageRepoDetailListResponseBodyPageInfo) *DescribeImageRepoDetailListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageRepoDetailListResponseBody) SetRequestId(v string) *DescribeImageRepoDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageRepoDetailListResponseBodyImageRepoResponses struct {
	// The number of alerts that are generated for the image repository.
	//
	// example:
	//
	// 0
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// Indicates whether alerts are generated for the image repository. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// The address of the image repository.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "domains": [
	//
	//                   "****.cn-shenzhen.cr.aliyuncs.com"
	//
	//             ],
	//
	//             "type": "internet"
	//
	//       },
	//
	//       {
	//
	//             "domains": [
	//
	//                   "****.cn-shenzhen.cr.aliyuncs.com"
	//
	//             ],
	//
	//             "type": "intranet"
	//
	//       },
	//
	//       {
	//
	//             "domains": [
	//
	//                   "****.cn-shenzhen.cr.aliyuncs.com"
	//
	//             ],
	//
	//             "type": "vpc"
	//
	//       }
	//
	// ]
	Endpoints *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The number of the images on which risks are detected.
	//
	// example:
	//
	// 2
	HasRiskImageCount *int32 `json:"HasRiskImageCount,omitempty" xml:"HasRiskImageCount,omitempty"`
	// The number of baseline risk items on the image repository.
	//
	// example:
	//
	// 0
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// Indicates whether baseline risk items are detected on the image repository. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// NO
	HcStatus *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	// The number of images.
	//
	// example:
	//
	// 3
	ImageCount *int32 `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// i-uf6fsg6xlmorug5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the image.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the image repository. Valid values:
	//
	// 	- **acr**
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// 	- **CI/CD**
	//
	// example:
	//
	// acr
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-2chlzf47w2rk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// zeus
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// test-dev
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// Indicates whether the image repository is at risk. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The usage label of the image.
	//
	// example:
	//
	// PAI
	SourceBizTag *string `json:"SourceBizTag,omitempty" xml:"SourceBizTag,omitempty"`
	// The number of vulnerabilities detected on the image repository.
	//
	// example:
	//
	// 0
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// Indicates whether vulnerabilities are detected on the image repository. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	VulStatus *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeImageRepoDetailListResponseBodyImageRepoResponses) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoDetailListResponseBodyImageRepoResponses) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetEndpoints() *string {
	return s.Endpoints
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetHasRiskImageCount() *int32 {
	return s.HasRiskImageCount
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetHcCount() *int32 {
	return s.HcCount
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetHcStatus() *string {
	return s.HcStatus
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRegistryType() *string {
	return s.RegistryType
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetSourceBizTag() *string {
	return s.SourceBizTag
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetAlarmCount(v int32) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.AlarmCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetAlarmStatus(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetEndpoints(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.Endpoints = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetHasRiskImageCount(v int32) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.HasRiskImageCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetHcCount(v int32) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.HcCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetHcStatus(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.HcStatus = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetImageCount(v int32) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.ImageCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetInstanceId(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRegionId(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RegionId = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRegistryType(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RegistryType = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRepoId(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RepoId = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRepoName(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RepoName = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRepoNamespace(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetRiskStatus(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.RiskStatus = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetSourceBizTag(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.SourceBizTag = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetVulCount(v int32) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.VulCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) SetVulStatus(v string) *DescribeImageRepoDetailListResponseBodyImageRepoResponses {
	s.VulStatus = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyImageRepoResponses) Validate() error {
	return dara.Validate(s)
}

type DescribeImageRepoDetailListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of image repositories.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageRepoDetailListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoDetailListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) SetCount(v int32) *DescribeImageRepoDetailListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageRepoDetailListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageRepoDetailListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageRepoDetailListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageRepoDetailListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
