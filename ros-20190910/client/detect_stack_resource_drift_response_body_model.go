// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackResourceDriftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActualProperties(v string) *DetectStackResourceDriftResponseBody
	GetActualProperties() *string
	SetDriftDetectionTime(v string) *DetectStackResourceDriftResponseBody
	GetDriftDetectionTime() *string
	SetExpectedProperties(v string) *DetectStackResourceDriftResponseBody
	GetExpectedProperties() *string
	SetLogicalResourceId(v string) *DetectStackResourceDriftResponseBody
	GetLogicalResourceId() *string
	SetPhysicalResourceId(v string) *DetectStackResourceDriftResponseBody
	GetPhysicalResourceId() *string
	SetPropertyDifferences(v []*DetectStackResourceDriftResponseBodyPropertyDifferences) *DetectStackResourceDriftResponseBody
	GetPropertyDifferences() []*DetectStackResourceDriftResponseBodyPropertyDifferences
	SetRequestId(v string) *DetectStackResourceDriftResponseBody
	GetRequestId() *string
	SetResourceDriftStatus(v string) *DetectStackResourceDriftResponseBody
	GetResourceDriftStatus() *string
	SetResourceType(v string) *DetectStackResourceDriftResponseBody
	GetResourceType() *string
	SetStackId(v string) *DetectStackResourceDriftResponseBody
	GetStackId() *string
}

type DetectStackResourceDriftResponseBody struct {
	// The actual JSON-formatted resource properties.
	//
	// example:
	//
	// {"ScalingRuleName": "test1"}
	ActualProperties *string `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
	// The time when the resource drift detection was initiated.
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
	// The logical ID of the resource that is defined in the template.
	//
	// example:
	//
	// ScalingRule
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// example:
	//
	// asr-2ze4zzc3kf9yz1kd****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The property drifts of the resource.
	PropertyDifferences []*DetectStackResourceDriftResponseBodyPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The drift status of the resource. Valid values:
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
	// The type of the resource.
	//
	// example:
	//
	// ALIYUN::ESS::ScalingRule
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackResourceDriftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectStackResourceDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBody) GetActualProperties() *string {
	return s.ActualProperties
}

func (s *DetectStackResourceDriftResponseBody) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *DetectStackResourceDriftResponseBody) GetExpectedProperties() *string {
	return s.ExpectedProperties
}

func (s *DetectStackResourceDriftResponseBody) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *DetectStackResourceDriftResponseBody) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *DetectStackResourceDriftResponseBody) GetPropertyDifferences() []*DetectStackResourceDriftResponseBodyPropertyDifferences {
	return s.PropertyDifferences
}

func (s *DetectStackResourceDriftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectStackResourceDriftResponseBody) GetResourceDriftStatus() *string {
	return s.ResourceDriftStatus
}

func (s *DetectStackResourceDriftResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DetectStackResourceDriftResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *DetectStackResourceDriftResponseBody) SetActualProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ActualProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetDriftDetectionTime(v string) *DetectStackResourceDriftResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetExpectedProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ExpectedProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetLogicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPhysicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPropertyDifferences(v []*DetectStackResourceDriftResponseBodyPropertyDifferences) *DetectStackResourceDriftResponseBody {
	s.PropertyDifferences = v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetRequestId(v string) *DetectStackResourceDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceDriftStatus(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceType(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetStackId(v string) *DetectStackResourceDriftResponseBody {
	s.StackId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectStackResourceDriftResponseBodyPropertyDifferences struct {
	// The actual value of the resource property.
	//
	// example:
	//
	// test1
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The drift type of the resource property. Valid values:
	//
	// 	- ADD: The property value has been added to a resource property whose data type was Array or List.
	//
	// 	- REMOVE: The property has been deleted from the current resource configuration.
	//
	// 	- NOT_EQUAL: The current property value differs from the expected value defined in the stack template.
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

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) String() string {
	return dara.Prettify(s)
}

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) GetActualValue() *string {
	return s.ActualValue
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) GetDifferenceType() *string {
	return s.DifferenceType
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) GetExpectedValue() *string {
	return s.ExpectedValue
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) GetPropertyPath() *string {
	return s.PropertyPath
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetActualValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetDifferenceType(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetExpectedValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetPropertyPath(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.PropertyPath = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) Validate() error {
	return dara.Validate(s)
}
