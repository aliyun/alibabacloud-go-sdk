// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionNames(v []*string) *AssociateResourceShareRequest
	GetPermissionNames() []*string
	SetResourceArns(v []*string) *AssociateResourceShareRequest
	GetResourceArns() []*string
	SetResourceProperties(v []*AssociateResourceShareRequestResourceProperties) *AssociateResourceShareRequest
	GetResourceProperties() []*AssociateResourceShareRequestResourceProperties
	SetResourceShareId(v string) *AssociateResourceShareRequest
	GetResourceShareId() *string
	SetResources(v []*AssociateResourceShareRequestResources) *AssociateResourceShareRequest
	GetResources() []*AssociateResourceShareRequestResources
	SetTargetProperties(v []*AssociateResourceShareRequestTargetProperties) *AssociateResourceShareRequest
	GetTargetProperties() []*AssociateResourceShareRequestTargetProperties
	SetTargets(v []*string) *AssociateResourceShareRequest
	GetTargets() []*string
}

type AssociateResourceShareRequest struct {
	// The information about the permissions. If you do not configure this parameter, the system automatically associates the default permission for the specified resource type with the resource share. For more information, see [Permission library](https://help.aliyun.com/document_detail/465474.html).
	PermissionNames    []*string                                          `json:"PermissionNames,omitempty" xml:"PermissionNames,omitempty" type:"Repeated"`
	ResourceArns       []*string                                          `json:"ResourceArns,omitempty" xml:"ResourceArns,omitempty" type:"Repeated"`
	ResourceProperties []*AssociateResourceShareRequestResourceProperties `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty" type:"Repeated"`
	// The ID of the resource share.
	//
	// This parameter is required.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The information about the resources.
	Resources []*AssociateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The properties of the principal.
	//
	// >  This parameter is available only when you specify an Alibaba Cloud service as a principal.
	TargetProperties []*AssociateResourceShareRequestTargetProperties `json:"TargetProperties,omitempty" xml:"TargetProperties,omitempty" type:"Repeated"`
	// The information about the principals.
	//
	// example:
	//
	// 172050525300****
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s AssociateResourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequest) GetPermissionNames() []*string {
	return s.PermissionNames
}

func (s *AssociateResourceShareRequest) GetResourceArns() []*string {
	return s.ResourceArns
}

func (s *AssociateResourceShareRequest) GetResourceProperties() []*AssociateResourceShareRequestResourceProperties {
	return s.ResourceProperties
}

func (s *AssociateResourceShareRequest) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *AssociateResourceShareRequest) GetResources() []*AssociateResourceShareRequestResources {
	return s.Resources
}

func (s *AssociateResourceShareRequest) GetTargetProperties() []*AssociateResourceShareRequestTargetProperties {
	return s.TargetProperties
}

func (s *AssociateResourceShareRequest) GetTargets() []*string {
	return s.Targets
}

func (s *AssociateResourceShareRequest) SetPermissionNames(v []*string) *AssociateResourceShareRequest {
	s.PermissionNames = v
	return s
}

func (s *AssociateResourceShareRequest) SetResourceArns(v []*string) *AssociateResourceShareRequest {
	s.ResourceArns = v
	return s
}

func (s *AssociateResourceShareRequest) SetResourceProperties(v []*AssociateResourceShareRequestResourceProperties) *AssociateResourceShareRequest {
	s.ResourceProperties = v
	return s
}

func (s *AssociateResourceShareRequest) SetResourceShareId(v string) *AssociateResourceShareRequest {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceShareRequest) SetResources(v []*AssociateResourceShareRequestResources) *AssociateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *AssociateResourceShareRequest) SetTargetProperties(v []*AssociateResourceShareRequestTargetProperties) *AssociateResourceShareRequest {
	s.TargetProperties = v
	return s
}

func (s *AssociateResourceShareRequest) SetTargets(v []*string) *AssociateResourceShareRequest {
	s.Targets = v
	return s
}

func (s *AssociateResourceShareRequest) Validate() error {
	if s.ResourceProperties != nil {
		for _, item := range s.ResourceProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetProperties != nil {
		for _, item := range s.TargetProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateResourceShareRequestResourceProperties struct {
	// example:
	//
	// {"sharePrincipals":true,"shareTagOptions":false}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// acs:vpc:cn-shanghai:103755469187****:vswitch/vsw-uf62b11ue4m8oz2di****
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s AssociateResourceShareRequestResourceProperties) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareRequestResourceProperties) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequestResourceProperties) GetProperty() *string {
	return s.Property
}

func (s *AssociateResourceShareRequestResourceProperties) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *AssociateResourceShareRequestResourceProperties) SetProperty(v string) *AssociateResourceShareRequestResourceProperties {
	s.Property = &v
	return s
}

func (s *AssociateResourceShareRequestResourceProperties) SetResourceArn(v string) *AssociateResourceShareRequestResourceProperties {
	s.ResourceArn = &v
	return s
}

func (s *AssociateResourceShareRequestResourceProperties) Validate() error {
	return dara.Validate(s)
}

type AssociateResourceShareRequestResources struct {
	// The ID of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  Resources.N.ResourceId and Resources.N.ResourceType must be used in pairs.
	//
	// example:
	//
	// vsw-bp183p93qs667muql****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// >  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be used in pairs.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AssociateResourceShareRequestResources) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *AssociateResourceShareRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *AssociateResourceShareRequestResources) SetResourceId(v string) *AssociateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *AssociateResourceShareRequestResources) SetResourceType(v string) *AssociateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

func (s *AssociateResourceShareRequestResources) Validate() error {
	return dara.Validate(s)
}

type AssociateResourceShareRequestTargetProperties struct {
	// The property parameter of the principal. For example, you can specify a parameter that indicates the time range for resource sharing. Valid values of `timeRangeType`:
	//
	// 	- timeRange: a specific time range
	//
	// 	- day: all day
	//
	// >  `TargetProperties.N.TargetId` and `TargetProperties.N.Property` must be used in pairs.
	//
	// example:
	//
	// {
	//
	//     "timeRange":{
	//
	//         "timeRangeType":"timeRange",
	//
	//         "beginAtTime":"00:00",
	//
	//         "timezone":"UTC+8",
	//
	//         "endAtTime":"19:59"
	//
	//     }
	//
	// }
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The ID of the principal.
	//
	// >  `TargetProperties.N.TargetId` and `TargetProperties.N.Property` must be used in pairs.
	//
	// example:
	//
	// 172050525300****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s AssociateResourceShareRequestTargetProperties) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareRequestTargetProperties) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareRequestTargetProperties) GetProperty() *string {
	return s.Property
}

func (s *AssociateResourceShareRequestTargetProperties) GetTargetId() *string {
	return s.TargetId
}

func (s *AssociateResourceShareRequestTargetProperties) SetProperty(v string) *AssociateResourceShareRequestTargetProperties {
	s.Property = &v
	return s
}

func (s *AssociateResourceShareRequestTargetProperties) SetTargetId(v string) *AssociateResourceShareRequestTargetProperties {
	s.TargetId = &v
	return s
}

func (s *AssociateResourceShareRequestTargetProperties) Validate() error {
	return dara.Validate(s)
}
