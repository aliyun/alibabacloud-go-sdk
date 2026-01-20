// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateResourceShareResponseBody
	GetRequestId() *string
	SetResourceShareAssociations(v []*AssociateResourceShareResponseBodyResourceShareAssociations) *AssociateResourceShareResponseBody
	GetResourceShareAssociations() []*AssociateResourceShareResponseBodyResourceShareAssociations
}

type AssociateResourceShareResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 111FB84A-60A9-403E-9067-E55D7EE95BD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the entities that are associated with the resource share.
	ResourceShareAssociations []*AssociateResourceShareResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s AssociateResourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateResourceShareResponseBody) GetResourceShareAssociations() []*AssociateResourceShareResponseBodyResourceShareAssociations {
	return s.ResourceShareAssociations
}

func (s *AssociateResourceShareResponseBody) SetRequestId(v string) *AssociateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateResourceShareResponseBody) SetResourceShareAssociations(v []*AssociateResourceShareResponseBodyResourceShareAssociations) *AssociateResourceShareResponseBody {
	s.ResourceShareAssociations = v
	return s
}

func (s *AssociateResourceShareResponseBody) Validate() error {
	if s.ResourceShareAssociations != nil {
		for _, item := range s.ResourceShareAssociations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateResourceShareResponseBodyResourceShareAssociations struct {
	// The association status. Valid values:
	//
	// 	- Associating: The entity is being associated.
	//
	// 	- Associated: The entity is associated.
	//
	// 	- Failed: The entity fails to be associated.
	//
	// 	- Disassociating: The entity is being disassociated.
	//
	// 	- Disassociated: The entity is disassociated.
	//
	// >  The system deletes the records of entities in the `Failed` or `Disassociated` state within 48 hours to 96 hours.
	//
	// example:
	//
	// Associating
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The cause of the association failure.
	//
	// example:
	//
	// The reason for the association failure.
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	// The association type. Valid values:
	//
	// 	- Resource
	//
	// 	- Target
	//
	// example:
	//
	// Resource
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The time when the association of the entity was created. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the shared resource was associated with the resource share.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was associated with the resource share.
	//
	// example:
	//
	// 2020-12-04T09:40:41.246Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the shared resource.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the principal.
	//
	// example:
	//
	// vsw-bp183p93qs667muql****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of AssociationType is Resource, the value of this parameter is the type of the shared resource. For information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// 	- If the value of AssociationType is Target, the value of this parameter is `Account`.
	//
	// example:
	//
	// VSwitch
	EntityType  *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// {"sharePrincipals":true,"shareTagOptions":false}
	ResourceProperty *string `json:"ResourceProperty,omitempty" xml:"ResourceProperty,omitempty"`
	// The ID of the resource share.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareId *string `json:"ResourceShareId,omitempty" xml:"ResourceShareId,omitempty"`
	// The name of the resource share.
	//
	// example:
	//
	// test
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The properties of the principal, such as the time range within which the resource is shared.
	//
	// >  This parameter is returned only if the principal is an Alibaba Cloud service.
	//
	// example:
	//
	// {
	//
	//     "plan":{
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
	TargetProperty *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the association of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the association of the shared resource was updated.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the association of the principal was updated.
	//
	// example:
	//
	// 2020-12-04T09:40:41.246Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AssociateResourceShareResponseBodyResourceShareAssociations) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourceShareResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetAssociationStatus() *string {
	return s.AssociationStatus
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetAssociationStatusMessage() *string {
	return s.AssociationStatusMessage
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetAssociationType() *string {
	return s.AssociationType
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetEntityId() *string {
	return s.EntityId
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetEntityType() *string {
	return s.EntityType
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetResourceProperty() *string {
	return s.ResourceProperty
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetTargetProperty() *string {
	return s.TargetProperty
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetAssociationType(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetCreateTime(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetEntityId(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetEntityType(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceArn(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceArn = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceProperty(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceProperty = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareId(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareName(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetTargetProperty(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) SetUpdateTime(v string) *AssociateResourceShareResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

func (s *AssociateResourceShareResponseBodyResourceShareAssociations) Validate() error {
	return dara.Validate(s)
}
