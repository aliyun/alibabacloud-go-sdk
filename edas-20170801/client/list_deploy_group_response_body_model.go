// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeployGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeployGroupResponseBody
	GetCode() *int32
	SetDeployGroupList(v *ListDeployGroupResponseBodyDeployGroupList) *ListDeployGroupResponseBody
	GetDeployGroupList() *ListDeployGroupResponseBodyDeployGroupList
	SetMessage(v string) *ListDeployGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeployGroupResponseBody
	GetRequestId() *string
}

type ListDeployGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the instance group in which the application is deployed.
	DeployGroupList *ListDeployGroupResponseBodyDeployGroupList `json:"DeployGroupList,omitempty" xml:"DeployGroupList,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3FDE-DS9R-*********************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeployGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeployGroupResponseBody) GetDeployGroupList() *ListDeployGroupResponseBodyDeployGroupList {
	return s.DeployGroupList
}

func (s *ListDeployGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeployGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeployGroupResponseBody) SetCode(v int32) *ListDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetDeployGroupList(v *ListDeployGroupResponseBodyDeployGroupList) *ListDeployGroupResponseBody {
	s.DeployGroupList = v
	return s
}

func (s *ListDeployGroupResponseBody) SetMessage(v string) *ListDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeployGroupResponseBody) SetRequestId(v string) *ListDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployGroupResponseBody) Validate() error {
	if s.DeployGroupList != nil {
		if err := s.DeployGroupList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeployGroupResponseBodyDeployGroupList struct {
	DeployGroup []*ListDeployGroupResponseBodyDeployGroupListDeployGroup `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" type:"Repeated"`
}

func (s ListDeployGroupResponseBodyDeployGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupList) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupList) GetDeployGroup() []*ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	return s.DeployGroup
}

func (s *ListDeployGroupResponseBodyDeployGroupList) SetDeployGroup(v []*ListDeployGroupResponseBodyDeployGroupListDeployGroup) *ListDeployGroupResponseBodyDeployGroupList {
	s.DeployGroup = v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupList) Validate() error {
	if s.DeployGroup != nil {
		for _, item := range s.DeployGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeployGroupResponseBodyDeployGroupListDeployGroup struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the change process for application deployment in the instance group.
	//
	// example:
	//
	// changeorder_a**_*******_**
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The name of the basic component.
	//
	// example:
	//
	// k8s-sc-consumer-****
	BaseComponentMetaName *string `json:"BaseComponentMetaName,omitempty" xml:"BaseComponentMetaName,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 0d247b93-8d62-4e34-****-************
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// doc-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running.
	//
	// example:
	//
	// 400
	CpuLimit *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// The number of CPU cores requested for each application instance when the application is running. Unit: cores. Value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1
	CpuRequest *string `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// The time when the application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627695779
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Container Service for Kubernetes (ACK) cluster.
	//
	// example:
	//
	// c66e65950db****cba92f17434df1****
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The name of the deployment.
	//
	// example:
	//
	// test
	DeploymentName *string `json:"DeploymentName,omitempty" xml:"DeploymentName,omitempty"`
	// The ID of the ACK cluster.
	//
	// example:
	//
	// 497806cb-****-6a7
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 8
	EphemeralStorageLimit *string `json:"EphemeralStorageLimit,omitempty" xml:"EphemeralStorageLimit,omitempty"`
	// The minimum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 4
	EphemeralStorageRequest *string `json:"EphemeralStorageRequest,omitempty" xml:"EphemeralStorageRequest,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// 577f4c50-16ee-43d8-****-************
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// _DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the instance group. Valid values:
	//
	// 	- 0: default group.
	//
	// 	- 1: Canary release is disabled for traffic management.
	//
	// 	- 2: Canary release is enabled for traffic management.
	//
	// example:
	//
	// 1
	GroupType *int32 `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The tag.
	//
	// example:
	//
	// test
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The time when the application was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1587888503825
	LastUpdateTime *int64 `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is running. Unit: MB. Value 0 indicates that no limit is set on the memory size.
	//
	// example:
	//
	// 0
	MemoryLimit *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The size of memory requested for each application instance when the application is running. Unit: MB. Value 0 indicates that no limit is set on the memory size.
	//
	// example:
	//
	// 512
	MemoryRequest *string `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	// The namespace.
	//
	// example:
	//
	// ping****est
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The external download URL of the deployment package.
	PackagePublicUrl *string `json:"PackagePublicUrl,omitempty" xml:"PackagePublicUrl,omitempty"`
	// The URL of the deployment package.
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// E
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The version of the deployment package that was used to deploy an application in the instance group.
	//
	// example:
	//
	// a7d48fe8-ad8f-****-89bd-74cc1ee6****
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The post-start script.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The pre-stop script.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The version of the application. The value progressively increases in the range of 0 to 7.
	//
	// example:
	//
	// 2
	Reversion *string `json:"Reversion,omitempty" xml:"Reversion,omitempty"`
	// The ID of the application deployed in the ACK cluster in Enterprise Distributed Application Service (EDAS).
	//
	// example:
	//
	// 53dd85cc-25b4-4d0e-****-6bf5465****4
	Selector *string `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The state of the application instance group. Valid values:
	//
	// 	- 0: ready
	//
	// 	- 1: in progress
	//
	// 	- 2: successful
	//
	// 	- 3: failed
	//
	// 	- 6: terminated
	//
	// 	- 10: failed due to a system exception
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The throttling policy. This parameter is reserved.
	//
	// example:
	//
	// RollingUpdate
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The time when the application was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627695779
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the vServer group of the Internet-facing SLB instance associated with the instance group.
	//
	// example:
	//
	// rsp-cige6******
	VExtServerGroupId *string `json:"VExtServerGroupId,omitempty" xml:"VExtServerGroupId,omitempty"`
	// The ID of the vServer group of the internal-facing Server Load Balancer (SLB) instance associated with the instance group.
	//
	// example:
	//
	// rsp-cige6******
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) String() string {
	return dara.Prettify(s)
}

func (s ListDeployGroupResponseBodyDeployGroupListDeployGroup) GoString() string {
	return s.String()
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetAppId() *string {
	return s.AppId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetBaseComponentMetaName() *string {
	return s.BaseComponentMetaName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCpuRequest() *string {
	return s.CpuRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEnv() *string {
	return s.Env
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEphemeralStorageLimit() *string {
	return s.EphemeralStorageLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetEphemeralStorageRequest() *string {
	return s.EphemeralStorageRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetGroupType() *int32 {
	return s.GroupType
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetLabels() *string {
	return s.Labels
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetLastUpdateTime() *int64 {
	return s.LastUpdateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetMemoryRequest() *string {
	return s.MemoryRequest
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetNameSpace() *string {
	return s.NameSpace
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackagePublicUrl() *string {
	return s.PackagePublicUrl
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPostStart() *string {
	return s.PostStart
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetPreStop() *string {
	return s.PreStop
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetReversion() *string {
	return s.Reversion
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetSelector() *string {
	return s.Selector
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetStatus() *string {
	return s.Status
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetStrategy() *string {
	return s.Strategy
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetVExtServerGroupId() *string {
	return s.VExtServerGroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetAppVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.AppVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetBaseComponentMetaName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.BaseComponentMetaName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetClusterName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.ClusterName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCpuRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CpuRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCreateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CreateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetCsClusterId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.CsClusterId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetDeploymentName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.DeploymentName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEnv(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Env = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEphemeralStorageLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.EphemeralStorageLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetEphemeralStorageRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.EphemeralStorageRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupName(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupName = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetGroupType(v int32) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.GroupType = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLabels(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Labels = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetLastUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.LastUpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryLimit(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryLimit = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetMemoryRequest(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.MemoryRequest = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetNameSpace(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.NameSpace = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackagePublicUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackagePublicUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageUrl(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageUrl = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPackageVersionId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PackageVersionId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPostStart(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PostStart = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetPreStop(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.PreStop = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetReversion(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Reversion = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetSelector(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Selector = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStatus(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Status = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetStrategy(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.Strategy = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetUpdateTime(v int64) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.UpdateTime = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVExtServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VExtServerGroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) SetVServerGroupId(v string) *ListDeployGroupResponseBodyDeployGroupListDeployGroup {
	s.VServerGroupId = &v
	return s
}

func (s *ListDeployGroupResponseBodyDeployGroupListDeployGroup) Validate() error {
	return dara.Validate(s)
}
