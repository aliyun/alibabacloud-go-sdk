// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationGroupsResponseBody
	GetCode() *string
	SetData(v []*DescribeApplicationGroupsResponseBodyData) *DescribeApplicationGroupsResponseBody
	GetData() []*DescribeApplicationGroupsResponseBodyData
	SetErrorCode(v string) *DescribeApplicationGroupsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationGroupsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationGroupsResponseBody
	GetTraceId() *string
}

type DescribeApplicationGroupsResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: success
	//
	// - **3xx**: redirection
	//
	// - **4xx**: request error
	//
	// - **5xx**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application group information.
	Data []*DescribeApplicationGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application instance groups were retrieved successfully. Valid values:
	//
	// - **true**: The application instance groups were retrieved successfully.
	//
	// - **false**: The application instance groups failed to be retrieved.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to precisely query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationGroupsResponseBody) GetData() []*DescribeApplicationGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationGroupsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationGroupsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationGroupsResponseBody) SetCode(v string) *DescribeApplicationGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetData(v []*DescribeApplicationGroupsResponseBodyData) *DescribeApplicationGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetErrorCode(v string) *DescribeApplicationGroupsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetMessage(v string) *DescribeApplicationGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetRequestId(v string) *DescribeApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetSuccess(v bool) *DescribeApplicationGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) SetTraceId(v string) *DescribeApplicationGroupsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationGroupsResponseBodyData struct {
	// The application runtime environment version in the HSF framework, such as the Ali-Tomcat container.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The group ID.
	//
	// example:
	//
	// b2a8a925-477a-eswa-b823-d5e22500****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// _DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The group type.
	//
	// example:
	//
	// 0
	GroupType *int32 `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The image address. This parameter is required when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The JDK version on which the deployment package depends. This parameter is not supported when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// Open JDK 8
	Jdk *string `json:"Jdk,omitempty" xml:"Jdk,omitempty"`
	// The application package type. Valid values:
	//
	// - If you deploy the application by using Java, valid values are **FatJar**, **War**, and **Image**.
	//
	// - If you deploy the application by using PHP, valid values are:
	//
	//     - **PhpZip**
	//
	//     - **IMAGE_PHP_5_4**
	//
	//     - **IMAGE_PHP_5_4_ALPINE**
	//
	//     - **IMAGE_PHP_5_5**
	//
	//     - **IMAGE_PHP_5_5_ALPINE**
	//
	//     - **IMAGE_PHP_5_6**
	//
	//     - **IMAGE_PHP_5_6_ALPINE**
	//
	//     - **IMAGE_PHP_7_0**
	//
	//     - **IMAGE_PHP_7_0_ALPINE**
	//
	//     - **IMAGE_PHP_7_1**
	//
	//     - **IMAGE_PHP_7_1_ALPINE**
	//
	//     - **IMAGE_PHP_7_2**
	//
	//     - **IMAGE_PHP_7_2_ALPINE**
	//
	//     - **IMAGE_PHP_7_3**
	//
	//     - **IMAGE_PHP_7_3_ALPINE**
	//
	// example:
	//
	// Image
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The deployment package address. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PhpZip**.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/demo/nginx:latest
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required when **Package Type*	- is set to **FatJar**, **War**, or **PhpZip**. If the application is deployed by using **ImageUrl**, this parameter is automatically generated.
	//
	// example:
	//
	// 1.0.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The version ID of the deployment package.
	//
	// example:
	//
	// 12345678-1234-4abc-8def-1234567890ab
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 10
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The number of running instances.
	//
	// example:
	//
	// 1
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// The Tomcat version on which the deployment package depends. This parameter is not supported when **Package Type*	- is set to **Image**.
	//
	// example:
	//
	// Apache Tomcat 7
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
}

func (s DescribeApplicationGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsResponseBodyData) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DescribeApplicationGroupsResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApplicationGroupsResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApplicationGroupsResponseBodyData) GetGroupType() *int32 {
	return s.GroupType
}

func (s *DescribeApplicationGroupsResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeApplicationGroupsResponseBodyData) GetJdk() *string {
	return s.Jdk
}

func (s *DescribeApplicationGroupsResponseBodyData) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeApplicationGroupsResponseBodyData) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DescribeApplicationGroupsResponseBodyData) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DescribeApplicationGroupsResponseBodyData) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *DescribeApplicationGroupsResponseBodyData) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DescribeApplicationGroupsResponseBodyData) GetRunningInstances() *int32 {
	return s.RunningInstances
}

func (s *DescribeApplicationGroupsResponseBodyData) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DescribeApplicationGroupsResponseBodyData) SetEdasContainerVersion(v string) *DescribeApplicationGroupsResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupId(v string) *DescribeApplicationGroupsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupName(v string) *DescribeApplicationGroupsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetGroupType(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetImageUrl(v string) *DescribeApplicationGroupsResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetJdk(v string) *DescribeApplicationGroupsResponseBodyData {
	s.Jdk = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageType(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageType = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageUrl(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageUrl = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageVersion(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageVersion = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetPackageVersionId(v string) *DescribeApplicationGroupsResponseBodyData {
	s.PackageVersionId = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetReplicas(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.Replicas = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetRunningInstances(v int32) *DescribeApplicationGroupsResponseBodyData {
	s.RunningInstances = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) SetWebContainer(v string) *DescribeApplicationGroupsResponseBodyData {
	s.WebContainer = &v
	return s
}

func (s *DescribeApplicationGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
