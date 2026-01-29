// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListApplicationsResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListApplicationsResponseBody
	GetCurrentPage() *int32
	SetData(v *ListApplicationsResponseBodyData) *ListApplicationsResponseBody
	GetData() *ListApplicationsResponseBodyData
	SetErrorCode(v string) *ListApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListApplicationsResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListApplicationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApplicationsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListApplicationsResponseBody
	GetTotalSize() *int32
}

type ListApplicationsResponseBody struct {
	// The HTTP status code. Take note of the following rules:
	//
	// - **2xx**: The call was successful.
	//
	// - **3xx**: The call was redirected.
	//
	// - **4xx**: The call failed.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The queried applications.
	Data *ListApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// The ID of the request.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of applications.
	//
	// example:
	//
	// 2
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApplicationsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListApplicationsResponseBody) GetData() *ListApplicationsResponseBodyData {
	return s.Data
}

func (s *ListApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApplicationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApplicationsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListApplicationsResponseBody) SetCode(v string) *ListApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationsResponseBody) SetCurrentPage(v int32) *ListApplicationsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsResponseBody) SetData(v *ListApplicationsResponseBodyData) *ListApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationsResponseBody) SetErrorCode(v string) *ListApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListApplicationsResponseBody) SetMessage(v string) *ListApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationsResponseBody) SetPageSize(v int32) *ListApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetSuccess(v bool) *ListApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalSize(v int32) *ListApplicationsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyData struct {
	// The queried applications.
	Applications []*ListApplicationsResponseBodyDataApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of records in each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of applications.
	//
	// example:
	//
	// 2
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyData) GetApplications() []*ListApplicationsResponseBodyDataApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListApplicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListApplicationsResponseBodyData) SetApplications(v []*ListApplicationsResponseBodyDataApplications) *ListApplicationsResponseBodyData {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBodyData) SetCurrentPage(v int32) *ListApplicationsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetPageSize(v int32) *ListApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsResponseBodyData) SetTotalSize(v int32) *ListApplicationsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListApplicationsResponseBodyData) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyDataApplications struct {
	// Indicates whether the application is being deleted. Valid values:
	//
	// 	- **true**: The application is being deleted.
	//
	// 	- **false**: The application is not being deleted.
	//
	// example:
	//
	// false
	AppDeletingStatus *bool `json:"AppDeletingStatus,omitempty" xml:"AppDeletingStatus,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// description
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application ID.
	//
	// example:
	//
	// f7730764-d88f-4b9a-8d8e-cd8efbfe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// Image
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The base app ID. Only gray-release applications have this property.
	//
	// example:
	//
	// xxx-xxx-xx-xxx
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The gray-release application list of this application.
	Children []*ListApplicationsResponseBodyDataApplicationsChildren `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// The CPU specifications that are required for each instance. Unit: millicores. This parameter cannot be set to 0. Valid values:
	//
	// 	- **500**
	//
	// 	- **1000**
	//
	// 	- **2000**
	//
	// 	- **4000**
	//
	// 	- **8000**
	//
	// 	- **16000**
	//
	// 	- **32000**
	//
	// example:
	//
	// 1000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 20
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// If the idle mode is enabled.
	//
	// example:
	//
	// false
	EnableIdle *string `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/sae-serverless-demo/sae-demo:microservice-java-provider-v1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 2
	Instances  *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	IsStateful *bool  `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The memory size that is required by each instance. Unit: MB. This parameter cannot be set to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 2000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24576*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// example:
	//
	// 1024
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The application has enabled MSE or not.
	//
	// example:
	//
	// true
	MseEnabled *bool `json:"MseEnabled,omitempty" xml:"MseEnabled,omitempty"`
	// The name space of MSE:
	//
	// - default: the free edition.
	//
	// - sae-pro: the professional edition.
	//
	// - sae-ent: the enterprise eiditon.
	//
	// example:
	//
	// sae-ent
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// demo
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The application edition.
	//
	// - lite: the lightweight edition.
	//
	// - std: the standard edition.
	//
	// - pro: the professional edition.
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The package URL.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of running instances.
	//
	// example:
	//
	// 2
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// The tags of the application.
	Tags []*ListApplicationsResponseBodyDataApplicationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListApplicationsResponseBodyDataApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplications) GetAppDeletingStatus() *bool {
	return s.AppDeletingStatus
}

func (s *ListApplicationsResponseBodyDataApplications) GetAppDescription() *string {
	return s.AppDescription
}

func (s *ListApplicationsResponseBodyDataApplications) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsResponseBodyDataApplications) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsResponseBodyDataApplications) GetAppType() *string {
	return s.AppType
}

func (s *ListApplicationsResponseBodyDataApplications) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *ListApplicationsResponseBodyDataApplications) GetChildren() []*ListApplicationsResponseBodyDataApplicationsChildren {
	return s.Children
}

func (s *ListApplicationsResponseBodyDataApplications) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListApplicationsResponseBodyDataApplications) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ListApplicationsResponseBodyDataApplications) GetEnableIdle() *string {
	return s.EnableIdle
}

func (s *ListApplicationsResponseBodyDataApplications) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListApplicationsResponseBodyDataApplications) GetInstances() *int32 {
	return s.Instances
}

func (s *ListApplicationsResponseBodyDataApplications) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *ListApplicationsResponseBodyDataApplications) GetMem() *int32 {
	return s.Mem
}

func (s *ListApplicationsResponseBodyDataApplications) GetMseEnabled() *bool {
	return s.MseEnabled
}

func (s *ListApplicationsResponseBodyDataApplications) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *ListApplicationsResponseBodyDataApplications) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationsResponseBodyDataApplications) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListApplicationsResponseBodyDataApplications) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *ListApplicationsResponseBodyDataApplications) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *ListApplicationsResponseBodyDataApplications) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *ListApplicationsResponseBodyDataApplications) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationsResponseBodyDataApplications) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListApplicationsResponseBodyDataApplications) GetRunningInstances() *int32 {
	return s.RunningInstances
}

func (s *ListApplicationsResponseBodyDataApplications) GetTags() []*ListApplicationsResponseBodyDataApplicationsTags {
	return s.Tags
}

func (s *ListApplicationsResponseBodyDataApplications) GetVpcId() *string {
	return s.VpcId
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppDeletingStatus(v bool) *ListApplicationsResponseBodyDataApplications {
	s.AppDeletingStatus = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppDescription(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppDescription = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppId(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppName(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetAppType(v string) *ListApplicationsResponseBodyDataApplications {
	s.AppType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetBaseAppId(v string) *ListApplicationsResponseBodyDataApplications {
	s.BaseAppId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetChildren(v []*ListApplicationsResponseBodyDataApplicationsChildren) *ListApplicationsResponseBodyDataApplications {
	s.Children = v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetCpu(v int32) *ListApplicationsResponseBodyDataApplications {
	s.Cpu = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetDiskSize(v int32) *ListApplicationsResponseBodyDataApplications {
	s.DiskSize = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetEnableIdle(v string) *ListApplicationsResponseBodyDataApplications {
	s.EnableIdle = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetImageUrl(v string) *ListApplicationsResponseBodyDataApplications {
	s.ImageUrl = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetInstances(v int32) *ListApplicationsResponseBodyDataApplications {
	s.Instances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetIsStateful(v bool) *ListApplicationsResponseBodyDataApplications {
	s.IsStateful = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetMem(v int32) *ListApplicationsResponseBodyDataApplications {
	s.Mem = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetMseEnabled(v bool) *ListApplicationsResponseBodyDataApplications {
	s.MseEnabled = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetMseNamespaceId(v string) *ListApplicationsResponseBodyDataApplications {
	s.MseNamespaceId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNamespaceId(v string) *ListApplicationsResponseBodyDataApplications {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNamespaceName(v string) *ListApplicationsResponseBodyDataApplications {
	s.NamespaceName = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetNewSaeVersion(v string) *ListApplicationsResponseBodyDataApplications {
	s.NewSaeVersion = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetPackageUrl(v string) *ListApplicationsResponseBodyDataApplications {
	s.PackageUrl = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetProgrammingLanguage(v string) *ListApplicationsResponseBodyDataApplications {
	s.ProgrammingLanguage = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetRegionId(v string) *ListApplicationsResponseBodyDataApplications {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetResourceType(v string) *ListApplicationsResponseBodyDataApplications {
	s.ResourceType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetRunningInstances(v int32) *ListApplicationsResponseBodyDataApplications {
	s.RunningInstances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetTags(v []*ListApplicationsResponseBodyDataApplicationsTags) *ListApplicationsResponseBodyDataApplications {
	s.Tags = v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) SetVpcId(v string) *ListApplicationsResponseBodyDataApplications {
	s.VpcId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplications) Validate() error {
	if s.Children != nil {
		for _, item := range s.Children {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyDataApplicationsChildren struct {
	// If is deleting this application.
	//
	// example:
	//
	// false
	AppDeletingStatus *bool `json:"AppDeletingStatus,omitempty" xml:"AppDeletingStatus,omitempty"`
	// The application description.
	//
	// example:
	//
	// Test
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application ID.
	//
	// example:
	//
	// xxx-xxx-xxx-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// app1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The way to deploy applications.
	//
	// example:
	//
	// Image
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The base application ID.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The CPU sepcification.
	//
	// example:
	//
	// 2000
	Cpu        *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnableIdle *string `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 2
	Instances  *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	IsStateful *bool  `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The memory specification.
	//
	// example:
	//
	// 2048
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// If this application has enabled MSE.
	//
	// example:
	//
	// true
	MseEnabled *bool `json:"MseEnabled,omitempty" xml:"MseEnabled,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// demo
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The application edition.
	//
	// - lite: the lightweight edition.
	//
	// - std: the standard edition.
	//
	// - pro: the professional edition.
	//
	// example:
	//
	// pro
	NewSaeVersion *string `json:"NewSaeVersion,omitempty" xml:"NewSaeVersion,omitempty"`
	// The programming language of this application.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of instances in running state.
	//
	// example:
	//
	// 2
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// If the scale rule is enabled.
	//
	// example:
	//
	// false
	ScaleRuleEnabled *bool `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The type of the scale rule.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The application tag.
	Tags []*ListApplicationsResponseBodyDataApplicationsChildrenTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyDataApplicationsChildren) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplicationsChildren) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetAppDeletingStatus() *bool {
	return s.AppDeletingStatus
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetAppDescription() *string {
	return s.AppDescription
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetAppType() *string {
	return s.AppType
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetEnableIdle() *string {
	return s.EnableIdle
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetInstances() *int32 {
	return s.Instances
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetMem() *int32 {
	return s.Mem
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetMseEnabled() *bool {
	return s.MseEnabled
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetNewSaeVersion() *string {
	return s.NewSaeVersion
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetRunningInstances() *int32 {
	return s.RunningInstances
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) GetTags() []*ListApplicationsResponseBodyDataApplicationsChildrenTags {
	return s.Tags
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetAppDeletingStatus(v bool) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.AppDeletingStatus = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetAppDescription(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.AppDescription = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetAppId(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetAppName(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetAppType(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.AppType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetBaseAppId(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.BaseAppId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetCpu(v int32) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.Cpu = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetEnableIdle(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.EnableIdle = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetInstances(v int32) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.Instances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetIsStateful(v bool) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.IsStateful = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetMem(v int32) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.Mem = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetMseEnabled(v bool) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.MseEnabled = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetNamespaceId(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.NamespaceId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetNamespaceName(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.NamespaceName = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetNewSaeVersion(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.NewSaeVersion = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetProgrammingLanguage(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.ProgrammingLanguage = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetRegionId(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetResourceType(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.ResourceType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetRunningInstances(v int32) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.RunningInstances = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetScaleRuleEnabled(v bool) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetScaleRuleType(v string) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.ScaleRuleType = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) SetTags(v []*ListApplicationsResponseBodyDataApplicationsChildrenTags) *ListApplicationsResponseBodyDataApplicationsChildren {
	s.Tags = v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildren) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyDataApplicationsChildrenTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListApplicationsResponseBodyDataApplicationsChildrenTags) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplicationsChildrenTags) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplicationsChildrenTags) GetKey() *string {
	return s.Key
}

func (s *ListApplicationsResponseBodyDataApplicationsChildrenTags) GetValue() *string {
	return s.Value
}

func (s *ListApplicationsResponseBodyDataApplicationsChildrenTags) SetKey(v string) *ListApplicationsResponseBodyDataApplicationsChildrenTags {
	s.Key = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildrenTags) SetValue(v string) *ListApplicationsResponseBodyDataApplicationsChildrenTags {
	s.Value = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsChildrenTags) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyDataApplicationsTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListApplicationsResponseBodyDataApplicationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyDataApplicationsTags) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) GetKey() *string {
	return s.Key
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) GetValue() *string {
	return s.Value
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) SetKey(v string) *ListApplicationsResponseBodyDataApplicationsTags {
	s.Key = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) SetValue(v string) *ListApplicationsResponseBodyDataApplicationsTags {
	s.Value = &v
	return s
}

func (s *ListApplicationsResponseBodyDataApplicationsTags) Validate() error {
	return dara.Validate(s)
}
