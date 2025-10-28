// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetStackResourceResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetStackResourceResponseBody
	GetDescription() *string
	SetDriftDetectionTime(v string) *GetStackResourceResponseBody
	GetDriftDetectionTime() *string
	SetLogicalResourceId(v string) *GetStackResourceResponseBody
	GetLogicalResourceId() *string
	SetMetadata(v map[string]interface{}) *GetStackResourceResponseBody
	GetMetadata() map[string]interface{}
	SetModuleInfo(v *GetStackResourceResponseBodyModuleInfo) *GetStackResourceResponseBody
	GetModuleInfo() *GetStackResourceResponseBodyModuleInfo
	SetPhysicalResourceId(v string) *GetStackResourceResponseBody
	GetPhysicalResourceId() *string
	SetRequestId(v string) *GetStackResourceResponseBody
	GetRequestId() *string
	SetResourceAttributes(v []map[string]interface{}) *GetStackResourceResponseBody
	GetResourceAttributes() []map[string]interface{}
	SetResourceDriftStatus(v string) *GetStackResourceResponseBody
	GetResourceDriftStatus() *string
	SetResourceType(v string) *GetStackResourceResponseBody
	GetResourceType() *string
	SetStackId(v string) *GetStackResourceResponseBody
	GetStackId() *string
	SetStackName(v string) *GetStackResourceResponseBody
	GetStackName() *string
	SetStatus(v string) *GetStackResourceResponseBody
	GetStatus() *string
	SetStatusReason(v string) *GetStackResourceResponseBody
	GetStatusReason() *string
	SetUpdateTime(v string) *GetStackResourceResponseBody
	GetUpdateTime() *string
}

type GetStackResourceResponseBody struct {
	// The resource type.
	//
	// example:
	//
	// 2019-08-01T06:01:23
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The reason why the resource is in its current state.
	//
	// example:
	//
	// no description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The time when the resource was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// WebServer
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The list of the resource properties.
	//
	// example:
	//
	// {"key": "value"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The information about the modules from which the resource is created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *GetStackResourceResponseBodyModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The metadata.
	//
	// example:
	//
	// d04af923-e6b7-4272-aeaa-47ec9777****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the resource in the last successful drift detection. Valid values:
	//
	// 	- DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	//
	// 	- MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	//
	// 	- NOT_CHECKED: ROS has not checked whether the actual configuration of the resource differs from its expected template configuration.
	//
	// 	- IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	ResourceAttributes []map[string]interface{} `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Repeated"`
	// The time when the last successful drift detection was performed on the stack.
	//
	// example:
	//
	// IN_SYNC
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The logical ID of the resource defined in the template.
	//
	// example:
	//
	// ALIYUN::ROS::WaitConditionHandle
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// efdf5c10-96a5-4fd7-ab89-68e7baa2****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack.
	//
	// example:
	//
	// test-describe-resource
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// state changed
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The name of the stack.
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or letter.
	//
	// example:
	//
	// 2019-08-01T06:01:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetStackResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackResourceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetStackResourceResponseBody) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackResourceResponseBody) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GetStackResourceResponseBody) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetStackResourceResponseBody) GetModuleInfo() *GetStackResourceResponseBodyModuleInfo {
	return s.ModuleInfo
}

func (s *GetStackResourceResponseBody) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *GetStackResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResourceResponseBody) GetResourceAttributes() []map[string]interface{} {
	return s.ResourceAttributes
}

func (s *GetStackResourceResponseBody) GetResourceDriftStatus() *string {
	return s.ResourceDriftStatus
}

func (s *GetStackResourceResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetStackResourceResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *GetStackResourceResponseBody) GetStackName() *string {
	return s.StackName
}

func (s *GetStackResourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetStackResourceResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetStackResourceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetStackResourceResponseBody) SetCreateTime(v string) *GetStackResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDescription(v string) *GetStackResourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDriftDetectionTime(v string) *GetStackResourceResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetLogicalResourceId(v string) *GetStackResourceResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetMetadata(v map[string]interface{}) *GetStackResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *GetStackResourceResponseBody) SetModuleInfo(v *GetStackResourceResponseBodyModuleInfo) *GetStackResourceResponseBody {
	s.ModuleInfo = v
	return s
}

func (s *GetStackResourceResponseBody) SetPhysicalResourceId(v string) *GetStackResourceResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetRequestId(v string) *GetStackResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceAttributes(v []map[string]interface{}) *GetStackResourceResponseBody {
	s.ResourceAttributes = v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceDriftStatus(v string) *GetStackResourceResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceType(v string) *GetStackResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackId(v string) *GetStackResourceResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackName(v string) *GetStackResourceResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStatus(v string) *GetStackResourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStatusReason(v string) *GetStackResourceResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResourceResponseBody) SetUpdateTime(v string) *GetStackResourceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetStackResourceResponseBody) Validate() error {
	if s.ModuleInfo != nil {
		if err := s.ModuleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackResourceResponseBodyModuleInfo struct {
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

func (s GetStackResourceResponseBodyModuleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResourceResponseBodyModuleInfo) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponseBodyModuleInfo) GetLogicalIdHierarchy() *string {
	return s.LogicalIdHierarchy
}

func (s *GetStackResourceResponseBodyModuleInfo) GetTypeHierarchy() *string {
	return s.TypeHierarchy
}

func (s *GetStackResourceResponseBodyModuleInfo) SetLogicalIdHierarchy(v string) *GetStackResourceResponseBodyModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *GetStackResourceResponseBodyModuleInfo) SetTypeHierarchy(v string) *GetStackResourceResponseBodyModuleInfo {
	s.TypeHierarchy = &v
	return s
}

func (s *GetStackResourceResponseBodyModuleInfo) Validate() error {
	return dara.Validate(s)
}
