// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppEnv(v string) *DeployApplicationRequest
	GetAppEnv() *string
	SetAppId(v string) *DeployApplicationRequest
	GetAppId() *string
	SetBatch(v int32) *DeployApplicationRequest
	GetBatch() *int32
	SetBatchWaitTime(v int32) *DeployApplicationRequest
	GetBatchWaitTime() *int32
	SetBuildPackId(v int64) *DeployApplicationRequest
	GetBuildPackId() *int64
	SetComponentIds(v string) *DeployApplicationRequest
	GetComponentIds() *string
	SetDeployType(v string) *DeployApplicationRequest
	GetDeployType() *string
	SetDesc(v string) *DeployApplicationRequest
	GetDesc() *string
	SetGray(v bool) *DeployApplicationRequest
	GetGray() *bool
	SetGroupId(v string) *DeployApplicationRequest
	GetGroupId() *string
	SetImageUrl(v string) *DeployApplicationRequest
	GetImageUrl() *string
	SetPackageVersion(v string) *DeployApplicationRequest
	GetPackageVersion() *string
	SetReleaseType(v int64) *DeployApplicationRequest
	GetReleaseType() *int64
	SetTrafficControlStrategy(v string) *DeployApplicationRequest
	GetTrafficControlStrategy() *string
	SetWarUrl(v string) *DeployApplicationRequest
	GetWarUrl() *string
}

type DeployApplicationRequest struct {
	// The environment variables of the application. Specify each environment variable by using two key-value pairs. Example: `{"name":"x","value":"y"},{"name":"x2","value":"y2"}`. The `keys` of the two key-value pairs are `name` and `value`.
	//
	// example:
	//
	// [{\\"name\\":\\"env_name_1\\", \\"value\\":\\"env_value_1\\"}, {\\"name\\":\\"env_name_2\\",\\"value\\":\\"env_value_2\\"}]
	AppEnv *string `json:"AppEnv,omitempty" xml:"AppEnv,omitempty"`
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/423162.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-********************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of batches per instance group.
	//
	// 	- If you specify an ID when you set the GroupId parameter, the application is deployed to the specified instance group. The minimum number of batches that can be specified is 1. The maximum number of batches is the maximum number of ECS instances in the Normal state in the instance group. The actual value falls in the range of [1, specified number]. The specified number of batches equals the number of ECS instances in the specified instance group.
	//
	// 	- If you set the GroupId parameter to all, the application is deployed to all instance groups. The minimum number of batches that can be specified is 1. The maximum number of batches is the number of ECS instances in the instance group that has the largest number of ECS instances in the Normal state.
	//
	// example:
	//
	// 1
	Batch *int32 `json:"Batch,omitempty" xml:"Batch,omitempty"`
	// The wait time between deployment batches for the application. Unit: minutes.
	//
	// 	- Default value: 0. If no wait time between deployment batches is needed, set this parameter to 0.
	//
	// 	- Maximum value: 5.
	//
	// If many deployment batches are needed, we recommend that you specify a small value for this parameter. Otherwise, the application deployment is time-consuming.
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The build package number of EDAS Container.
	//
	// 	- You do not need to set the parameter if you do not need to change the EDAS Container version during the deployment.
	//
	// 	- Set the parameter if you need to update the EDAS Container version of the application during the deployment.
	//
	// You can query the build package number by using one of the following methods:
	//
	// 	- Call the ListBuildPack operation. For more information, see [ListBuildPack](https://help.aliyun.com/document_detail/149391.html).
	//
	// 	- Obtain the value in the **Build package number*	- column of the [Release notes for EDAS Container](https://help.aliyun.com/document_detail/92614.html) topic. For example, `59` indicates `EDAS Container 3.5.8`.
	//
	// example:
	//
	// 59
	BuildPackId *int64 `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	// The IDs of the components used by the application. The parameter is not applicable to High-Speed Framework (HSF) applications. You can call the ListComponents operation to query the component IDs. For more information, see [ListComponents](https://help.aliyun.com/document_detail/423223.html).
	//
	// 	- If you have specified the component IDs when you create the application, you do not need to set the parameter when you deploy the application.
	//
	// 	- Set the parameter if you need to update the component versions for the application during the deployment.
	//
	// Valid values for common application components:
	//
	// 	- 4: Apache Tomcat 7.0.91
	//
	// 	- 7: Apache Tomcat 8.5.42
	//
	// 	- 5: OpenJDK 1.8.x
	//
	// 	- 6: OpenJDK 1.7.x
	//
	// For more information, see the Common application parameters section of the [InsertApplication](https://help.aliyun.com/document_detail/423185.html) topic.
	//
	// example:
	//
	// 7
	ComponentIds *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
	// The deployment mode of the application. Valid values: `url` and `image`. The image value is deprecated. You can deploy an application to a Swarm cluster only by using an image.``
	//
	// This parameter is required.
	//
	// example:
	//
	// URL
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The description of the application deployment.
	//
	// example:
	//
	// Deploy by edas pop api
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// Specifies whether canary release is selected as the deployment method. Valid values:
	//
	// 	- true: Canary release is selected.
	//
	//     	- To implement a canary release, specify the GroupId parameter, which specifies the ID of the instance group for the canary release.
	//
	//     	- Canary release can be selected as the deployment method for only one batch.
	//
	//     	- After the canary release is complete, the application is released in regular mode. The Batch parameter specifies the number of batches.
	//
	// 	- false: Single-batch release or phased release is selected.
	//
	// example:
	//
	// true
	Gray *bool `json:"Gray,omitempty" xml:"Gray,omitempty"`
	// The ID of the instance group to which the application is deployed. You can call the ListDeployGroup operation to query the ID of the instance group. For more information, see [ListDeployGroup](https://help.aliyun.com/document_detail/423184.html).
	//
	// Set the parameter to `all` if you want to deploy the application to all instance groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The URL of the application image that is used to deploy the application in a Swarm cluster. We recommend that you use an image that is stored in Alibaba Cloud Container Registry. This parameter is deprecated.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/mw/testapp:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The version of the application deployment package. The value can be up to 64 characters in length. We recommend that you use a timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The mode in which the deployment batches are triggered. Valid values:
	//
	// 	- 0: automatic.
	//
	// 	- 1: You must manually trigger the next batch. You can manually click **Proceed to Next Batch*	- in the console or call the ContinuePipeline operation to proceed to the next batch. We recommend that you choose the automatic mode when you call an API operation to deploy the application. For more information, see [ContinuePipeline](https://help.aliyun.com/document_detail/126990.html).
	//
	// example:
	//
	// 0
	ReleaseType *int64 `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	// The canary release policy. For more information about canary release policies, see [DeployK8sApplication](https://help.aliyun.com/document_detail/423212.html).
	//
	// example:
	//
	// {"http":{"rules":[{"conditionType":"percent","percent":10}]}}
	TrafficControlStrategy *string `json:"TrafficControlStrategy,omitempty" xml:"TrafficControlStrategy,omitempty"`
	// The URL of the application deployment package. The package can be a WAR or JAR package. This parameter is required if you set the **DeployType*	- parameter to `url`. We recommend that you specify the URL of an application deployment package that is stored in an Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// https://edas.oss-cn-hangzhou.aliyuncs.com/demo/hello-edas.war
	WarUrl *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
}

func (s DeployApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) GetAppEnv() *string {
	return s.AppEnv
}

func (s *DeployApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeployApplicationRequest) GetBatch() *int32 {
	return s.Batch
}

func (s *DeployApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DeployApplicationRequest) GetBuildPackId() *int64 {
	return s.BuildPackId
}

func (s *DeployApplicationRequest) GetComponentIds() *string {
	return s.ComponentIds
}

func (s *DeployApplicationRequest) GetDeployType() *string {
	return s.DeployType
}

func (s *DeployApplicationRequest) GetDesc() *string {
	return s.Desc
}

func (s *DeployApplicationRequest) GetGray() *bool {
	return s.Gray
}

func (s *DeployApplicationRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeployApplicationRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DeployApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DeployApplicationRequest) GetReleaseType() *int64 {
	return s.ReleaseType
}

func (s *DeployApplicationRequest) GetTrafficControlStrategy() *string {
	return s.TrafficControlStrategy
}

func (s *DeployApplicationRequest) GetWarUrl() *string {
	return s.WarUrl
}

func (s *DeployApplicationRequest) SetAppEnv(v string) *DeployApplicationRequest {
	s.AppEnv = &v
	return s
}

func (s *DeployApplicationRequest) SetAppId(v string) *DeployApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeployApplicationRequest) SetBatch(v int32) *DeployApplicationRequest {
	s.Batch = &v
	return s
}

func (s *DeployApplicationRequest) SetBatchWaitTime(v int32) *DeployApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployApplicationRequest) SetBuildPackId(v int64) *DeployApplicationRequest {
	s.BuildPackId = &v
	return s
}

func (s *DeployApplicationRequest) SetComponentIds(v string) *DeployApplicationRequest {
	s.ComponentIds = &v
	return s
}

func (s *DeployApplicationRequest) SetDeployType(v string) *DeployApplicationRequest {
	s.DeployType = &v
	return s
}

func (s *DeployApplicationRequest) SetDesc(v string) *DeployApplicationRequest {
	s.Desc = &v
	return s
}

func (s *DeployApplicationRequest) SetGray(v bool) *DeployApplicationRequest {
	s.Gray = &v
	return s
}

func (s *DeployApplicationRequest) SetGroupId(v string) *DeployApplicationRequest {
	s.GroupId = &v
	return s
}

func (s *DeployApplicationRequest) SetImageUrl(v string) *DeployApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *DeployApplicationRequest) SetPackageVersion(v string) *DeployApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployApplicationRequest) SetReleaseType(v int64) *DeployApplicationRequest {
	s.ReleaseType = &v
	return s
}

func (s *DeployApplicationRequest) SetTrafficControlStrategy(v string) *DeployApplicationRequest {
	s.TrafficControlStrategy = &v
	return s
}

func (s *DeployApplicationRequest) SetWarUrl(v string) *DeployApplicationRequest {
	s.WarUrl = &v
	return s
}

func (s *DeployApplicationRequest) Validate() error {
	return dara.Validate(s)
}
