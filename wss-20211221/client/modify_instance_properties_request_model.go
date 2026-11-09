// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstancePropertiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstancePropertiesRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest
	GetInstanceIds() []*string
	SetKey(v string) *ModifyInstancePropertiesRequest
	GetKey() *string
	SetProperties(v []*ModifyInstancePropertiesRequestProperties) *ModifyInstancePropertiesRequest
	GetProperties() []*ModifyInstancePropertiesRequestProperties
	SetResourceType(v string) *ModifyInstancePropertiesRequest
	GetResourceType() *string
	SetValue(v string) *ModifyInstancePropertiesRequest
	GetValue() *string
}

type ModifyInstancePropertiesRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// mdp-0c62ayep0nk4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The key of the attribute.
	//
	// example:
	//
	// PackageUsedUpStrategy
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of instance attributes to modify. You can modify multiple key-value pairs at a time. These pairs are merged with the original Key and Value request parameters. If an attribute with the same name has the same value, it is deduplicated. If the values are different, a parameter error is returned.
	Properties []*ModifyInstancePropertiesRequestProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > This parameter is case-sensitive. Ensure correct spelling.
	//
	// This parameter is required.
	//
	// example:
	//
	// DurationPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The value of the attribute.
	//
	// - PackageUsedUpStrategy: [enum]
	//
	//    - Postpaid (switches to the pay-as-you-go billing method)
	//
	//    - Shutdown (hibernates or shuts down the instance)
	//
	//    - Maintenance (stops the instance and enters O&M mode, and does not allow client connections)
	//
	// example:
	//
	// Postpaid
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstancePropertiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePropertiesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstancePropertiesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyInstancePropertiesRequest) GetKey() *string {
	return s.Key
}

func (s *ModifyInstancePropertiesRequest) GetProperties() []*ModifyInstancePropertiesRequestProperties {
	return s.Properties
}

func (s *ModifyInstancePropertiesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyInstancePropertiesRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyInstancePropertiesRequest) SetInstanceId(v string) *ModifyInstancePropertiesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetInstanceIds(v []*string) *ModifyInstancePropertiesRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetKey(v string) *ModifyInstancePropertiesRequest {
	s.Key = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetProperties(v []*ModifyInstancePropertiesRequestProperties) *ModifyInstancePropertiesRequest {
	s.Properties = v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetResourceType(v string) *ModifyInstancePropertiesRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) SetValue(v string) *ModifyInstancePropertiesRequest {
	s.Value = &v
	return s
}

func (s *ModifyInstancePropertiesRequest) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstancePropertiesRequestProperties struct {
	// The name of the instance attribute. For example, DisableRemainingNotify specifies whether to disable the remaining duration reminder for the client, and PackageUsedUpStrategy specifies the policy applied when the package duration is used up.
	//
	// example:
	//
	// DisableRemainingNotify
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the instance attribute. When Key is set to DisableRemainingNotify, only the string true or false is supported, and ResourceType must be set to DurationPackage. A value of true indicates that the remaining duration reminder for the client is disabled. A value of false indicates that the reminder is retained and triggered based on the original conditions. Administrator notifications continue to be triggered based on the original conditions. For other attributes, pass values based on their respective definitions.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstancePropertiesRequestProperties) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstancePropertiesRequestProperties) GoString() string {
	return s.String()
}

func (s *ModifyInstancePropertiesRequestProperties) GetKey() *string {
	return s.Key
}

func (s *ModifyInstancePropertiesRequestProperties) GetValue() *string {
	return s.Value
}

func (s *ModifyInstancePropertiesRequestProperties) SetKey(v string) *ModifyInstancePropertiesRequestProperties {
	s.Key = &v
	return s
}

func (s *ModifyInstancePropertiesRequestProperties) SetValue(v string) *ModifyInstancePropertiesRequestProperties {
	s.Value = &v
	return s
}

func (s *ModifyInstancePropertiesRequestProperties) Validate() error {
	return dara.Validate(s)
}
