// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackOperationRisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMissingPolicyActions(v []*string) *ListStackOperationRisksResponseBody
	GetMissingPolicyActions() []*string
	SetRequestId(v string) *ListStackOperationRisksResponseBody
	GetRequestId() *string
	SetRiskResources(v []*ListStackOperationRisksResponseBodyRiskResources) *ListStackOperationRisksResponseBody
	GetRiskResources() []*ListStackOperationRisksResponseBodyRiskResources
}

type ListStackOperationRisksResponseBody struct {
	// The operations on which the permissions are not granted to the Alibaba Cloud account of the caller.
	MissingPolicyActions []*string `json:"MissingPolicyActions,omitempty" xml:"MissingPolicyActions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 72108E7A-E874-4A5E-B22C-A61E94AD12CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that are at risk.
	RiskResources []*ListStackOperationRisksResponseBodyRiskResources `json:"RiskResources,omitempty" xml:"RiskResources,omitempty" type:"Repeated"`
}

func (s ListStackOperationRisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackOperationRisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBody) GetMissingPolicyActions() []*string {
	return s.MissingPolicyActions
}

func (s *ListStackOperationRisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackOperationRisksResponseBody) GetRiskResources() []*ListStackOperationRisksResponseBodyRiskResources {
	return s.RiskResources
}

func (s *ListStackOperationRisksResponseBody) SetMissingPolicyActions(v []*string) *ListStackOperationRisksResponseBody {
	s.MissingPolicyActions = v
	return s
}

func (s *ListStackOperationRisksResponseBody) SetRequestId(v string) *ListStackOperationRisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBody) SetRiskResources(v []*ListStackOperationRisksResponseBodyRiskResources) *ListStackOperationRisksResponseBody {
	s.RiskResources = v
	return s
}

func (s *ListStackOperationRisksResponseBody) Validate() error {
	if s.RiskResources != nil {
		for _, item := range s.RiskResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackOperationRisksResponseBodyRiskResources struct {
	// The error code that is returned when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	//
	// example:
	//
	// NoPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logical ID of the resource. The logical ID is the resource name that is defined in the template.
	//
	// example:
	//
	// MySG
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The error message that is returned when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	//
	// example:
	//
	// You are not authorized to complete this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The physical ID of the resource. The physical ID is the actual ID of the resource.
	//
	// example:
	//
	// sg-bp1dpioafqphedg9****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The cause of the risk.
	//
	// example:
	//
	// There are some ECS instances (i-bp18el96s4wq635e****) depending on the security group.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	//
	// example:
	//
	// DF4296CF-F45F-4845-A72B-BE617601DB25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ALIYUN::ECS::SecurityGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the risk. Valid values:
	//
	// 	- Referenced: The resource is referenced by other resources.
	//
	// 	- MaybeReferenced: The resource may be referenced by other resources.
	//
	// 	- AdditionalRiskCheckRequired: An additional risk detection is required for a nested stack.
	//
	// 	- OperationIgnored: The operation does not take effect for the resource.
	//
	// example:
	//
	// Referenced
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s ListStackOperationRisksResponseBodyRiskResources) String() string {
	return dara.Prettify(s)
}

func (s ListStackOperationRisksResponseBodyRiskResources) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetCode() *string {
	return s.Code
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetMessage() *string {
	return s.Message
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetReason() *string {
	return s.Reason
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListStackOperationRisksResponseBodyRiskResources) GetRiskType() *string {
	return s.RiskType
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetCode(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Code = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetLogicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetMessage(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Message = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetPhysicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetReason(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Reason = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRequestId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetResourceType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRiskType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RiskType = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) Validate() error {
	return dara.Validate(s)
}
