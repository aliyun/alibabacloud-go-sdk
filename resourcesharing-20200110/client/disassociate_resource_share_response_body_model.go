// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourceShareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateResourceShareResponseBody
	GetRequestId() *string
	SetResourceShareAssociations(v []*DisassociateResourceShareResponseBodyResourceShareAssociations) *DisassociateResourceShareResponseBody
	GetResourceShareAssociations() []*DisassociateResourceShareResponseBodyResourceShareAssociations
}

type DisassociateResourceShareResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 95230BC9-A8E8-4493-96BD-4F0C758E37F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the entities that are associated with the resource share.
	ResourceShareAssociations []*DisassociateResourceShareResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s DisassociateResourceShareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceShareResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateResourceShareResponseBody) GetResourceShareAssociations() []*DisassociateResourceShareResponseBodyResourceShareAssociations {
	return s.ResourceShareAssociations
}

func (s *DisassociateResourceShareResponseBody) SetRequestId(v string) *DisassociateResourceShareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateResourceShareResponseBody) SetResourceShareAssociations(v []*DisassociateResourceShareResponseBodyResourceShareAssociations) *DisassociateResourceShareResponseBody {
	s.ResourceShareAssociations = v
	return s
}

func (s *DisassociateResourceShareResponseBody) Validate() error {
	return dara.Validate(s)
}

type DisassociateResourceShareResponseBodyResourceShareAssociations struct {
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
	// Disassociating
	AssociationStatus *string `json:"AssociationStatus,omitempty" xml:"AssociationStatus,omitempty"`
	// The cause of the disassociation failure.
	//
	// example:
	//
	// The Resources is invalid.
	AssociationStatusMessage *string `json:"AssociationStatusMessage,omitempty" xml:"AssociationStatusMessage,omitempty"`
	// The association type. Valid values:
	//
	// 	- Resource
	//
	// 	- Target
	//
	// example:
	//
	// Target
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The time when the disassociation of the entity was performed. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the resource was disassociated from the resource share.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was disassociated from the resource share.
	//
	// example:
	//
	// 2020-12-04T09:40:41.250Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the resource.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the resource directory, folder, member, or Alibaba Cloud service.
	//
	// example:
	//
	// 172050525300****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of AssociationType is Resource, the value of this parameter is the type of the resource. For more information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// 	- If the value of AssociationType is Target, the value of this parameter is Account.
	//
	// example:
	//
	// Account
	EntityType  *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
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
	TargetProperty *string `json:"TargetProperty,omitempty" xml:"TargetProperty,omitempty"`
	// The time when the disassociation of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the disassociation of the resource was updated.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the disassociation of the principal was updated.
	//
	// example:
	//
	// 2020-12-04T09:40:45.556Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DisassociateResourceShareResponseBodyResourceShareAssociations) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourceShareResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetAssociationStatus() *string {
	return s.AssociationStatus
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetAssociationStatusMessage() *string {
	return s.AssociationStatusMessage
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetAssociationType() *string {
	return s.AssociationType
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetEntityId() *string {
	return s.EntityId
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetEntityType() *string {
	return s.EntityType
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetTargetProperty() *string {
	return s.TargetProperty
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetAssociationType(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetCreateTime(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetEntityId(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetEntityType(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetResourceArn(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceArn = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareId(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetResourceShareName(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetTargetProperty(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) SetUpdateTime(v string) *DisassociateResourceShareResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

func (s *DisassociateResourceShareResponseBodyResourceShareAssociations) Validate() error {
	return dara.Validate(s)
}
