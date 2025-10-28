// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListStackResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*ListStackResourcesResponseBodyResources) *ListStackResourcesResponseBody
	GetResources() []*ListStackResourcesResponseBodyResources
}

type ListStackResourcesResponseBody struct {
	// Details about resources.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*ListStackResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListStackResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackResourcesResponseBody) GetResources() []*ListStackResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListStackResourcesResponseBody) SetRequestId(v string) *ListStackResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackResourcesResponseBody) SetResources(v []*ListStackResourcesResponseBodyResources) *ListStackResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListStackResourcesResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackResourcesResponseBodyResources struct {
	// The time when the resource was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2019-08-01T06:01:23
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The logical ID of the resource. The logical ID is the resource name that is defined in the template.
	//
	// example:
	//
	// dummy
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The information about the modules from which the resource is created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *ListStackResourcesResponseBodyResourcesModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The physical ID of the resource.
	//
	// example:
	//
	// d04af923-e6b7-4272-aeaa-47ec9777****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The drift state of the resource in the most recent successful drift detection. Valid values:
	//
	// 	- DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	//
	// 	- MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	//
	// 	- NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	//
	// 	- IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	//
	// example:
	//
	// IN_SYNC
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ROS::Stack
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// example:
	//
	// test-describe-resource
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The state of the resource. Valid values:
	//
	// 	- INIT_COMPLETE: The resource is pending to be created.
	//
	// 	- CREATE_COMPLETE: The resource is created.
	//
	// 	- CREATE_FAILED: The resource failed to be created.
	//
	// 	- CREATE_IN_PROGRESS: The resource is being created.
	//
	// 	- UPDATE_IN_PROGRESS: The resource is being updated.
	//
	// 	- UPDATE_FAILED: The resource failed to be updated.
	//
	// 	- UPDATE_COMPLETE: The resource is updated.
	//
	// 	- DELETE_IN_PROGRESS: The resource is being deleted.
	//
	// 	- DELETE_FAILED: The resource failed to be deleted.
	//
	// 	- DELETE_COMPLETE: The resource is deleted.
	//
	// 	- CHECK_IN_PROGRESS: The resource is being validated.
	//
	// 	- CHECK_FAILED: The resource failed to be validated.
	//
	// 	- CHECK_COMPLETE: The resource is validated.
	//
	// 	- IMPORT_IN_PROGRESS: The resource is being imported.
	//
	// 	- IMPORT_FAILED: The resource failed to be imported.
	//
	// 	- IMPORT_COMPLETE: The resource is imported.
	//
	// example:
	//
	// UPDATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource is in its current state.
	//
	// example:
	//
	// state changed
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the resource was updated. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2019-08-01T06:01:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStackResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStackResourcesResponseBodyResources) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *ListStackResourcesResponseBodyResources) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *ListStackResourcesResponseBodyResources) GetModuleInfo() *ListStackResourcesResponseBodyResourcesModuleInfo {
	return s.ModuleInfo
}

func (s *ListStackResourcesResponseBodyResources) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *ListStackResourcesResponseBodyResources) GetResourceDriftStatus() *string {
	return s.ResourceDriftStatus
}

func (s *ListStackResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListStackResourcesResponseBodyResources) GetStackId() *string {
	return s.StackId
}

func (s *ListStackResourcesResponseBodyResources) GetStackName() *string {
	return s.StackName
}

func (s *ListStackResourcesResponseBodyResources) GetStatus() *string {
	return s.Status
}

func (s *ListStackResourcesResponseBodyResources) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListStackResourcesResponseBodyResources) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListStackResourcesResponseBodyResources) SetCreateTime(v string) *ListStackResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetDriftDetectionTime(v string) *ListStackResourcesResponseBodyResources {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetLogicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetModuleInfo(v *ListStackResourcesResponseBodyResourcesModuleInfo) *ListStackResourcesResponseBodyResources {
	s.ModuleInfo = v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetPhysicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceDriftStatus(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceType(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackId(v string) *ListStackResourcesResponseBodyResources {
	s.StackId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackName(v string) *ListStackResourcesResponseBodyResources {
	s.StackName = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStatus(v string) *ListStackResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStatusReason(v string) *ListStackResourcesResponseBodyResources {
	s.StatusReason = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetUpdateTime(v string) *ListStackResourcesResponseBodyResources {
	s.UpdateTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) Validate() error {
	if s.ModuleInfo != nil {
		if err := s.ModuleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStackResourcesResponseBodyResourcesModuleInfo struct {
	// The concatenated logical IDs of one or more modules that contain the resource. The modules are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from Module B nested within Parent Module A:
	//
	// `moduleA/moduleB`
	//
	// example:
	//
	// moduleA/moduleB
	LogicalIdHierarchy *string `json:"LogicalIdHierarchy,omitempty" xml:"LogicalIdHierarchy,omitempty"`
	// The concatenated types of one or more modules that contain the resource. The module types are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from a module of the `MODULE::ROS::Child::Example` type that is nested within a parent module of the `MODULE::ROS::Parent::Example` type:
	//
	// `MODULE::ROS::Parent::Example/MODULE::ROS::Child::Example`
	//
	// example:
	//
	// MODULE::ROS::Parent::Example/MODULE::ROS::Child::Example
	TypeHierarchy *string `json:"TypeHierarchy,omitempty" xml:"TypeHierarchy,omitempty"`
}

func (s ListStackResourcesResponseBodyResourcesModuleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourcesResponseBodyResourcesModuleInfo) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) GetLogicalIdHierarchy() *string {
	return s.LogicalIdHierarchy
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) GetTypeHierarchy() *string {
	return s.TypeHierarchy
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) SetLogicalIdHierarchy(v string) *ListStackResourcesResponseBodyResourcesModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) SetTypeHierarchy(v string) *ListStackResourcesResponseBodyResourcesModuleInfo {
	s.TypeHierarchy = &v
	return s
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) Validate() error {
	return dara.Validate(s)
}
