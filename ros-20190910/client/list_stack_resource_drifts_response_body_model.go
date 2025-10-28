// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackResourceDriftsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListStackResourceDriftsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListStackResourceDriftsResponseBody
	GetRequestId() *string
	SetResourceDrifts(v []*ListStackResourceDriftsResponseBodyResourceDrifts) *ListStackResourceDriftsResponseBody
	GetResourceDrifts() []*ListStackResourceDriftsResponseBodyResourceDrifts
}

type ListStackResourceDriftsResponseBody struct {
	// The query token returned in this call.
	//
	// example:
	//
	// AAAAAdDWBF2****w==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource drifts.
	ResourceDrifts []*ListStackResourceDriftsResponseBodyResourceDrifts `json:"ResourceDrifts,omitempty" xml:"ResourceDrifts,omitempty" type:"Repeated"`
}

func (s ListStackResourceDriftsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStackResourceDriftsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackResourceDriftsResponseBody) GetResourceDrifts() []*ListStackResourceDriftsResponseBodyResourceDrifts {
	return s.ResourceDrifts
}

func (s *ListStackResourceDriftsResponseBody) SetNextToken(v string) *ListStackResourceDriftsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetRequestId(v string) *ListStackResourceDriftsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetResourceDrifts(v []*ListStackResourceDriftsResponseBodyResourceDrifts) *ListStackResourceDriftsResponseBody {
	s.ResourceDrifts = v
	return s
}

func (s *ListStackResourceDriftsResponseBody) Validate() error {
	if s.ResourceDrifts != nil {
		for _, item := range s.ResourceDrifts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackResourceDriftsResponseBodyResourceDrifts struct {
	// The actual JSON-formatted resource properties.
	//
	// example:
	//
	// {"ScalingRuleName": "test1"}
	ActualProperties *string `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
	// The time when the drift detection operation was performed on the resource.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The JSON-formatted resource properties that are defined in the template.
	//
	// example:
	//
	// {"ScalingRuleName": "test2"}
	ExpectedProperties *string `json:"ExpectedProperties,omitempty" xml:"ExpectedProperties,omitempty"`
	// The logical ID of the resource. The logical ID indicates the name of the resource that is defined in the template.
	//
	// example:
	//
	// ScalingRule
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The information about the modules from which the resource was created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The physical ID of the resource.
	//
	// example:
	//
	// asr-2ze4zzc3kf9yz1kd****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The property drifts of the resource.
	PropertyDifferences []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	// The drift state of the resource. Valid values:
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
	// MODIFIED
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ESS::ScalingRule
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetActualProperties() *string {
	return s.ActualProperties
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetExpectedProperties() *string {
	return s.ExpectedProperties
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetModuleInfo() *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo {
	return s.ModuleInfo
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetPropertyDifferences() []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	return s.PropertyDifferences
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetResourceDriftStatus() *string {
	return s.ResourceDriftStatus
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) GetStackId() *string {
	return s.StackId
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetActualProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ActualProperties = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetDriftDetectionTime(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetExpectedProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ExpectedProperties = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetLogicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetModuleInfo(v *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ModuleInfo = v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPhysicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPropertyDifferences(v []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PropertyDifferences = v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceDriftStatus(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceType(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetStackId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.StackId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) Validate() error {
	if s.ModuleInfo != nil {
		if err := s.ModuleInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PropertyDifferences != nil {
		for _, item := range s.PropertyDifferences {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo struct {
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

func (s ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) GetLogicalIdHierarchy() *string {
	return s.LogicalIdHierarchy
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) GetTypeHierarchy() *string {
	return s.TypeHierarchy
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) SetLogicalIdHierarchy(v string) *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) SetTypeHierarchy(v string) *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo {
	s.TypeHierarchy = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) Validate() error {
	return dara.Validate(s)
}

type ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences struct {
	// The actual value of the resource property.
	//
	// example:
	//
	// test1
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The drift type of the resource property. Valid values:
	//
	// 	- ADD: The value is added to a resource property whose data type is Array or List.
	//
	// 	- REMOVE: The property is deleted from the current resource configuration.
	//
	// 	- NOT_EQUAL: The current property value differs from the expected value that is defined in the stack template.
	//
	// example:
	//
	// NOT_EQUAL
	DifferenceType *string `json:"DifferenceType,omitempty" xml:"DifferenceType,omitempty"`
	// The expected value of the resource property that is defined in the template.
	//
	// example:
	//
	// test2
	ExpectedValue *string `json:"ExpectedValue,omitempty" xml:"ExpectedValue,omitempty"`
	// The path of the resource property.
	//
	// example:
	//
	// /ScalingRuleName
	PropertyPath *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) String() string {
	return dara.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GetActualValue() *string {
	return s.ActualValue
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GetDifferenceType() *string {
	return s.DifferenceType
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GetExpectedValue() *string {
	return s.ExpectedValue
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GetPropertyPath() *string {
	return s.PropertyPath
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetActualValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetDifferenceType(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetExpectedValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetPropertyPath(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.PropertyPath = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) Validate() error {
	return dara.Validate(s)
}
