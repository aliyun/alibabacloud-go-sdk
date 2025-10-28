// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateParameterConstraintsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterConstraints(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraints) *GetTemplateParameterConstraintsResponseBody
	GetParameterConstraints() []*GetTemplateParameterConstraintsResponseBodyParameterConstraints
	SetRequestId(v string) *GetTemplateParameterConstraintsResponseBody
	GetRequestId() *string
}

type GetTemplateParameterConstraintsResponseBody struct {
	// The constraints of the parameters.
	ParameterConstraints []*GetTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9816785B-BCF8-514D-8B76-C1EC2BC954FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBody) GetParameterConstraints() []*GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	return s.ParameterConstraints
}

func (s *GetTemplateParameterConstraintsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraints) *GetTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBody) Validate() error {
	if s.ParameterConstraints != nil {
		for _, item := range s.ParameterConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	// The values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The names of the associated parameters.
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// The behavior of the parameter. Valid values:
	//
	// 	- NoLimit: No limit is imposed on the value of this parameter.
	//
	// 	- NotSupport: The value of this parameter cannot be queried.
	//
	// 	- QueryError: This parameter failed to be queried.
	//
	// > If AllowedValues is not returned, Behavior and BehaviorReason are returned.
	//
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason why the behavior of the parameter is returned.
	//
	// example:
	//
	// No resource property refer to the parameter
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	// The values that do not conform to the parameter constraints.
	//
	// > If AllowedValues is returned, IllegalValueByParameterConstraints and IllegalValueByRules are returned at the same time.
	IllegalValueByParameterConstraints []interface{} `json:"IllegalValueByParameterConstraints,omitempty" xml:"IllegalValueByParameterConstraints,omitempty" type:"Repeated"`
	// The values that do not match the rules in the template.
	//
	// > If AllowedValues is returned, IllegalValueByParameterConstraints and IllegalValueByRules are returned at the same time.
	IllegalValueByRules []interface{} `json:"IllegalValueByRules,omitempty" xml:"IllegalValueByRules,omitempty" type:"Repeated"`
	// The unsupported resource in the template.
	NotSupportResources []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources `json:"NotSupportResources,omitempty" xml:"NotSupportResources,omitempty" type:"Repeated"`
	// The original constraint information.
	OriginalConstraints []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// The name of the parameter.
	//
	// example:
	//
	// ZoneInfo
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The error that is returned when the request fails.
	QueryErrors []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors `json:"QueryErrors,omitempty" xml:"QueryErrors,omitempty" type:"Repeated"`
	// Query the details of timeout.
	QueryTimeoutDetails []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails `json:"QueryTimeoutDetails,omitempty" xml:"QueryTimeoutDetails,omitempty" type:"Repeated"`
	// The data type of the parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetAllowedValues() []*string {
	return s.AllowedValues
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetAssociationParameterNames() []*string {
	return s.AssociationParameterNames
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetBehavior() *string {
	return s.Behavior
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetBehaviorReason() *string {
	return s.BehaviorReason
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetIllegalValueByParameterConstraints() []interface{} {
	return s.IllegalValueByParameterConstraints
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetIllegalValueByRules() []interface{} {
	return s.IllegalValueByRules
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetNotSupportResources() []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources {
	return s.NotSupportResources
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetOriginalConstraints() []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	return s.OriginalConstraints
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetQueryErrors() []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	return s.QueryErrors
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetQueryTimeoutDetails() []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails {
	return s.QueryTimeoutDetails
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) GetType() *string {
	return s.Type
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetIllegalValueByParameterConstraints(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.IllegalValueByParameterConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetIllegalValueByRules(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.IllegalValueByRules = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetNotSupportResources(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.NotSupportResources = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetQueryErrors(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.QueryErrors = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetQueryTimeoutDetails(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.QueryTimeoutDetails = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) Validate() error {
	if s.NotSupportResources != nil {
		for _, item := range s.NotSupportResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OriginalConstraints != nil {
		for _, item := range s.OriginalConstraints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryErrors != nil {
		for _, item := range s.QueryErrors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.QueryTimeoutDetails != nil {
		for _, item := range s.QueryTimeoutDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources struct {
	// The name of the resource property.
	//
	// example:
	//
	// InstanceName
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) GetPropertyName() *string {
	return s.PropertyName
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) SetPropertyName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources {
	s.PropertyName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) Validate() error {
	return dara.Validate(s)
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	// The values of the parameter.
	AllowedValues []interface{} `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// Behavior of the parameter
	//
	// example:
	//
	// QueryError
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason for the parameter behavior
	//
	// example:
	//
	// No resource property refer to the parameter
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	PropertiesData *string `json:"PropertiesData,omitempty" xml:"PropertiesData,omitempty"`
	// The name of the resource property.
	//
	// example:
	//
	// ZoneId
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	RequestInfo  *string `json:"RequestInfo,omitempty" xml:"RequestInfo,omitempty"`
	// The name of the resource that is defined in the template.
	//
	// example:
	//
	// MyECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetAllowedValues() []interface{} {
	return s.AllowedValues
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetBehavior() *string {
	return s.Behavior
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetBehaviorReason() *string {
	return s.BehaviorReason
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetPropertiesData() *string {
	return s.PropertiesData
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetPropertyName() *string {
	return s.PropertyName
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetRequestInfo() *string {
	return s.RequestInfo
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetBehavior(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.Behavior = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetBehaviorReason(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertiesData(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertiesData = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetRequestInfo(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.RequestInfo = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) Validate() error {
	return dara.Validate(s)
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors struct {
	// The error message.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The resource name.
	//
	// example:
	//
	// MyECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// InstanceType is needed while query DataDisk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetErrorMessage(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) Validate() error {
	return dara.Validate(s)
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails struct {
	// Error message.
	//
	// example:
	//
	// query property SlaveZoneIds.	- in resource rds error, error message: query 8 seconds timeout
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Resource name.
	//
	// example:
	//
	// rds
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Resource type.
	//
	// example:
	//
	// ALIYUN::RDS::DBInstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) SetErrorMessage(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) SetResourceName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails {
	s.ResourceName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryTimeoutDetails) Validate() error {
	return dara.Validate(s)
}
