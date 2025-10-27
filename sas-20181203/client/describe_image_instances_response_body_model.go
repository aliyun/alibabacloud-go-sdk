// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageInstanceList(v []*DescribeImageInstancesResponseBodyImageInstanceList) *DescribeImageInstancesResponseBody
	GetImageInstanceList() []*DescribeImageInstancesResponseBodyImageInstanceList
	SetPageInfo(v *DescribeImageInstancesResponseBodyPageInfo) *DescribeImageInstancesResponseBody
	GetPageInfo() *DescribeImageInstancesResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageInstancesResponseBody
	GetRequestId() *string
}

type DescribeImageInstancesResponseBody struct {
	// The information about the images.
	ImageInstanceList []*DescribeImageInstancesResponseBodyImageInstanceList `json:"ImageInstanceList,omitempty" xml:"ImageInstanceList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageInstancesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageInstancesResponseBody) GetImageInstanceList() []*DescribeImageInstancesResponseBodyImageInstanceList {
	return s.ImageInstanceList
}

func (s *DescribeImageInstancesResponseBody) GetPageInfo() *DescribeImageInstancesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageInstancesResponseBody) SetImageInstanceList(v []*DescribeImageInstancesResponseBodyImageInstanceList) *DescribeImageInstancesResponseBody {
	s.ImageInstanceList = v
	return s
}

func (s *DescribeImageInstancesResponseBody) SetPageInfo(v *DescribeImageInstancesResponseBodyPageInfo) *DescribeImageInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageInstancesResponseBody) SetRequestId(v string) *DescribeImageInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageInstancesResponseBody) Validate() error {
	if s.ImageInstanceList != nil {
		for _, item := range s.ImageInstanceList {
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

type DescribeImageInstancesResponseBodyImageInstanceList struct {
	// The number of alerts that are generated for the image.
	//
	// example:
	//
	// 0
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// Indicates whether alerts are generated for the image. Valid values:
	//
	// 	- **YES**
	//
	// 	- **NO**
	//
	// example:
	//
	// NO
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// Indicates whether the image was deployed. Valid values:
	//
	// 	- **0**: The image was not deployed.
	//
	// 	- **1**: The image was deployed.
	//
	// example:
	//
	// 1
	Deployed *int32 `json:"Deployed,omitempty" xml:"Deployed,omitempty"`
	// The digest value of the image.
	//
	// example:
	//
	// a5ccdd9b166b67e02954aa9b618fe19b7968bd56a15463d2ad7f2643ba5b****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The address of the image.
	//
	// example:
	//
	// []
	Endpoints *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The number of baseline risks.
	//
	// example:
	//
	// 0
	HcCount *int32 `json:"HcCount,omitempty" xml:"HcCount,omitempty"`
	// Indicates whether baseline risks exist. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// NO
	HcStatus *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	// The timestamp generated when the image was created. Unit: milliseconds.
	//
	// example:
	//
	// 1600069948849
	ImageCreate *string `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// c20987f18b130f9d144c9828df630417e2a9523148930dc3963e9d0dab30****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image. Unit: MB.
	//
	// example:
	//
	// 1604487690
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The timestamp generated when the image was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1600069948849
	ImageUpdate *string `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The instance ID of the image.
	//
	// example:
	//
	// 39010****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1721363159000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The region ID of the image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the image. Valid values:
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
	// f2b86d20bf0855af6aa268ce90fd****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// sas-script-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// N/A
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- **private**
	//
	// 	- **public**
	//
	// example:
	//
	// private
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// Indicates whether risks exist. Valid values:
	//
	// 	- **NO**
	//
	// 	- **YES**
	//
	// example:
	//
	// NO
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The scan progress of the image. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	ScaProgress *int32 `json:"ScaProgress,omitempty" xml:"ScaProgress,omitempty"`
	// The error code of the image scan result. Valid values:
	//
	// 	- **TASK_NOT_EXISTS**: The image scan task does not exist.
	//
	// 	- **TASK_NOT_SUPPORT_REGION**: The image scan task cannot be performed in the current region.
	//
	// 	- **forbid_create_repeat_task**: The image scan task already exists.
	//
	// example:
	//
	// TASK_NOT_SUPPORT_REGION
	ScaResult *string `json:"ScaResult,omitempty" xml:"ScaResult,omitempty"`
	// The scan status of the image. Valid values:
	//
	// 	- **INIT**: The image scan task is pending startup.
	//
	// 	- **START**: The image scan task is started.
	//
	// 	- **MESSAGE_SEND**: The message about the image scan task is sent.
	//
	// 	- **START_RUN**: The image analysis task is started.
	//
	// 	- **DOWNLOAD**: The image scan result is downloaded.
	//
	// 	- **PRE_ANALYZER**: The image pre-analysis is started.
	//
	// 	- **WEB_SHELL_ANALYZER**: The WebShell analysis of the image is complete.
	//
	// 	- **CVE_ANALYZER**: The Common Vulnerabilities and Exposures (CVE) analysis of the image is complete.
	//
	// 	- **BIN_ANALYZER**: The binary analysis of the image is complete.
	//
	// 	- **OTHER_ANALYZER**: The extended analysis of the image is complete.
	//
	// 	- **SUCCESS**: The image scan task is complete.
	//
	// 	- **PRE_ANALYZER_SUCCESS**: The image pre-analysis is complete.
	//
	// 	- **FAIL**: The image scan task failed.
	//
	// 	- **TIMEOUT**: The image scan task timed out.
	//
	// example:
	//
	// SUCCESS
	ScaStatus *string `json:"ScaStatus,omitempty" xml:"ScaStatus,omitempty"`
	// The usage label of the image.
	//
	// example:
	//
	// PAI
	SourceBizTag *string `json:"SourceBizTag,omitempty" xml:"SourceBizTag,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- **NORMAL**
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// mysql_5.7
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 12f80307-60aa-4efa-863a-56d72fb****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The number of vulnerabilities in the image.
	//
	// example:
	//
	// 0
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	// Indicates whether vulnerabilities exist in the image. Valid values:
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

func (s DescribeImageInstancesResponseBodyImageInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstancesResponseBodyImageInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetDeployed() *int32 {
	return s.Deployed
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetDigest() *string {
	return s.Digest
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetEndpoints() *string {
	return s.Endpoints
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetHcCount() *int32 {
	return s.HcCount
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetHcStatus() *string {
	return s.HcStatus
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetImageCreate() *string {
	return s.ImageCreate
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetImageSize() *string {
	return s.ImageSize
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetImageUpdate() *string {
	return s.ImageUpdate
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRegistryType() *string {
	return s.RegistryType
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRepoType() *string {
	return s.RepoType
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetScaProgress() *int32 {
	return s.ScaProgress
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetScaResult() *string {
	return s.ScaResult
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetScaStatus() *string {
	return s.ScaStatus
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetSourceBizTag() *string {
	return s.SourceBizTag
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetTag() *string {
	return s.Tag
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) GetVulStatus() *string {
	return s.VulStatus
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetAlarmCount(v int32) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.AlarmCount = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetAlarmStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetDeployed(v int32) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Deployed = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetDigest(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Digest = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetEndpoints(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Endpoints = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetHcCount(v int32) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.HcCount = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetHcStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.HcStatus = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetImageCreate(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ImageCreate = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetImageId(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ImageId = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetImageSize(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ImageSize = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetImageUpdate(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ImageUpdate = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetInstanceId(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetLastScanTime(v int64) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRegionId(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRegistryType(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RegistryType = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRepoId(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RepoId = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRepoName(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RepoName = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRepoNamespace(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRepoType(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RepoType = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetRiskStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.RiskStatus = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetScaProgress(v int32) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ScaProgress = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetScaResult(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ScaResult = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetScaStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.ScaStatus = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetSourceBizTag(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.SourceBizTag = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetTag(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Tag = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetUuid(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.Uuid = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetVulCount(v int32) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.VulCount = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) SetVulStatus(v string) *DescribeImageInstancesResponseBodyImageInstanceList {
	s.VulStatus = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyImageInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeImageInstancesResponseBodyPageInfo struct {
	// The number of images returned on the current page.
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
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageInstancesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageInstancesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageInstancesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageInstancesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageInstancesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeImageInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageInstancesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
