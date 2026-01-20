// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowExternalTargets(v bool) *CreateResourceShareRequest
	GetAllowExternalTargets() *bool
	SetPermissionNames(v []*string) *CreateResourceShareRequest
	GetPermissionNames() []*string
	SetResourceArns(v []*string) *CreateResourceShareRequest
	GetResourceArns() []*string
	SetResourceGroupId(v string) *CreateResourceShareRequest
	GetResourceGroupId() *string
	SetResourceProperties(v []*CreateResourceShareRequestResourceProperties) *CreateResourceShareRequest
	GetResourceProperties() []*CreateResourceShareRequestResourceProperties
	SetResourceShareName(v string) *CreateResourceShareRequest
	GetResourceShareName() *string
	SetResources(v []*CreateResourceShareRequestResources) *CreateResourceShareRequest
	GetResources() []*CreateResourceShareRequestResources
	SetTag(v []*CreateResourceShareRequestTag) *CreateResourceShareRequest
	GetTag() []*CreateResourceShareRequestTag
	SetTargetProperties(v []*CreateResourceShareRequestTargetProperties) *CreateResourceShareRequest
	GetTargetProperties() []*CreateResourceShareRequestTargetProperties
	SetTargets(v []*string) *CreateResourceShareRequest
	GetTargets() []*string
}

type CreateResourceShareRequest struct {
	// Specifies whether resources in the resource share can be shared with accounts outside the resource directory. Valid values:
	//
	// 	- false (default): Resources in the resource share can be shared only with accounts in the resource directory.
	//
	// 	- true: Resources in the resource share can be shared with both accounts in the resource directory and accounts outside the resource directory.
	//
	// example:
	//
	// false
	AllowExternalTargets *bool `json:"AllowExternalTargets,omitempty" xml:"AllowExternalTargets,omitempty"`
	// The information about the permissions. If you do not configure this parameter, the system automatically associates the default permission for the specified resource type with the resource share. For more information, see [Permission library](https://help.aliyun.com/document_detail/465474.html).
	PermissionNames []*string `json:"PermissionNames,omitempty" xml:"PermissionNames,omitempty" type:"Repeated"`
	ResourceArns    []*string `json:"ResourceArns,omitempty" xml:"ResourceArns,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz5nlvlak****
	ResourceGroupId    *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceProperties []*CreateResourceShareRequestResourceProperties `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty" type:"Repeated"`
	// The name of the resource share.
	//
	// The name must be 1 to 50 characters in length.
	//
	// The name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The information about the shared resources.
	Resources []*CreateResourceShareRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateResourceShareRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The properties of the principal.
	//
	// >  This parameter is available only when you specify an Alibaba Cloud service as a principal.
	TargetProperties []*CreateResourceShareRequestTargetProperties `json:"TargetProperties,omitempty" xml:"TargetProperties,omitempty" type:"Repeated"`
	// The information about the principals.
	//
	// example:
	//
	// 172050525300****
	Targets []*string `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s CreateResourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequest) GetAllowExternalTargets() *bool {
	return s.AllowExternalTargets
}

func (s *CreateResourceShareRequest) GetPermissionNames() []*string {
	return s.PermissionNames
}

func (s *CreateResourceShareRequest) GetResourceArns() []*string {
	return s.ResourceArns
}

func (s *CreateResourceShareRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateResourceShareRequest) GetResourceProperties() []*CreateResourceShareRequestResourceProperties {
	return s.ResourceProperties
}

func (s *CreateResourceShareRequest) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *CreateResourceShareRequest) GetResources() []*CreateResourceShareRequestResources {
	return s.Resources
}

func (s *CreateResourceShareRequest) GetTag() []*CreateResourceShareRequestTag {
	return s.Tag
}

func (s *CreateResourceShareRequest) GetTargetProperties() []*CreateResourceShareRequestTargetProperties {
	return s.TargetProperties
}

func (s *CreateResourceShareRequest) GetTargets() []*string {
	return s.Targets
}

func (s *CreateResourceShareRequest) SetAllowExternalTargets(v bool) *CreateResourceShareRequest {
	s.AllowExternalTargets = &v
	return s
}

func (s *CreateResourceShareRequest) SetPermissionNames(v []*string) *CreateResourceShareRequest {
	s.PermissionNames = v
	return s
}

func (s *CreateResourceShareRequest) SetResourceArns(v []*string) *CreateResourceShareRequest {
	s.ResourceArns = v
	return s
}

func (s *CreateResourceShareRequest) SetResourceGroupId(v string) *CreateResourceShareRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateResourceShareRequest) SetResourceProperties(v []*CreateResourceShareRequestResourceProperties) *CreateResourceShareRequest {
	s.ResourceProperties = v
	return s
}

func (s *CreateResourceShareRequest) SetResourceShareName(v string) *CreateResourceShareRequest {
	s.ResourceShareName = &v
	return s
}

func (s *CreateResourceShareRequest) SetResources(v []*CreateResourceShareRequestResources) *CreateResourceShareRequest {
	s.Resources = v
	return s
}

func (s *CreateResourceShareRequest) SetTag(v []*CreateResourceShareRequestTag) *CreateResourceShareRequest {
	s.Tag = v
	return s
}

func (s *CreateResourceShareRequest) SetTargetProperties(v []*CreateResourceShareRequestTargetProperties) *CreateResourceShareRequest {
	s.TargetProperties = v
	return s
}

func (s *CreateResourceShareRequest) SetTargets(v []*string) *CreateResourceShareRequest {
	s.Targets = v
	return s
}

func (s *CreateResourceShareRequest) Validate() error {
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
	if s.Tag != nil {
		for _, item := range s.Tag {
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

type CreateResourceShareRequestResourceProperties struct {
	// example:
	//
	// {"sharePrincipals":true,"shareTagOptions":false}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// acs:vpc:cn-shanghai:103755469187****:vswitch/vsw-uf62b11ue4m8oz2di****
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s CreateResourceShareRequestResourceProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareRequestResourceProperties) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequestResourceProperties) GetProperty() *string {
	return s.Property
}

func (s *CreateResourceShareRequestResourceProperties) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *CreateResourceShareRequestResourceProperties) SetProperty(v string) *CreateResourceShareRequestResourceProperties {
	s.Property = &v
	return s
}

func (s *CreateResourceShareRequestResourceProperties) SetResourceArn(v string) *CreateResourceShareRequestResourceProperties {
	s.ResourceArn = &v
	return s
}

func (s *CreateResourceShareRequestResourceProperties) Validate() error {
	return dara.Validate(s)
}

type CreateResourceShareRequestResources struct {
	// The ID of a shared resource.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five shared resources can be specified at a time.
	//
	// >  `Resources.N.ResourceId` and `Resources.N.ResourceType` must be used in pairs.
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

func (s CreateResourceShareRequestResources) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareRequestResources) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateResourceShareRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceShareRequestResources) SetResourceId(v string) *CreateResourceShareRequestResources {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceShareRequestResources) SetResourceType(v string) *CreateResourceShareRequestResources {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceShareRequestResources) Validate() error {
	return dara.Validate(s)
}

type CreateResourceShareRequestTag struct {
	// The tag key.
	//
	// >  The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// >  The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceShareRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateResourceShareRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateResourceShareRequestTag) SetKey(v string) *CreateResourceShareRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceShareRequestTag) SetValue(v string) *CreateResourceShareRequestTag {
	s.Value = &v
	return s
}

func (s *CreateResourceShareRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateResourceShareRequestTargetProperties struct {
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

func (s CreateResourceShareRequestTargetProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceShareRequestTargetProperties) GoString() string {
	return s.String()
}

func (s *CreateResourceShareRequestTargetProperties) GetProperty() *string {
	return s.Property
}

func (s *CreateResourceShareRequestTargetProperties) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateResourceShareRequestTargetProperties) SetProperty(v string) *CreateResourceShareRequestTargetProperties {
	s.Property = &v
	return s
}

func (s *CreateResourceShareRequestTargetProperties) SetTargetId(v string) *CreateResourceShareRequestTargetProperties {
	s.TargetId = &v
	return s
}

func (s *CreateResourceShareRequestTargetProperties) Validate() error {
	return dara.Validate(s)
}
