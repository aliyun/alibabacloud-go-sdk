// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResourceShareAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceShareAssociationsResponseBody
	GetRequestId() *string
	SetResourceShareAssociations(v []*ListResourceShareAssociationsResponseBodyResourceShareAssociations) *ListResourceShareAssociationsResponseBody
	GetResourceShareAssociations() []*ListResourceShareAssociationsResponseBodyResourceShareAssociations
}

type ListResourceShareAssociationsResponseBody struct {
	// The `token` that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11BA57B5-7301-4E2F-BBA5-2AE4C2F4FCDB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the entities.
	ResourceShareAssociations []*ListResourceShareAssociationsResponseBodyResourceShareAssociations `json:"ResourceShareAssociations,omitempty" xml:"ResourceShareAssociations,omitempty" type:"Repeated"`
}

func (s ListResourceShareAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceShareAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceShareAssociationsResponseBody) GetResourceShareAssociations() []*ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	return s.ResourceShareAssociations
}

func (s *ListResourceShareAssociationsResponseBody) SetNextToken(v string) *ListResourceShareAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBody) SetRequestId(v string) *ListResourceShareAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBody) SetResourceShareAssociations(v []*ListResourceShareAssociationsResponseBodyResourceShareAssociations) *ListResourceShareAssociationsResponseBody {
	s.ResourceShareAssociations = v
	return s
}

func (s *ListResourceShareAssociationsResponseBody) Validate() error {
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

type ListResourceShareAssociationsResponseBodyResourceShareAssociations struct {
	// The information about the failure.
	AssociationFailedDetails []*ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails `json:"AssociationFailedDetails,omitempty" xml:"AssociationFailedDetails,omitempty" type:"Repeated"`
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
	// Associated
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
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the shared resource was associated with or disassociated from the resource share.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the principal was associated with or disassociated from the resource share.
	//
	// example:
	//
	// 2020-12-07T07:39:01.818Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the ID of the resource.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the ID of the principal.
	//
	// example:
	//
	// vsw-bp1upw03qyz8n7us9****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of AssociationType is Resource, the value of this parameter is the type of the resource. For information about the types of resources that can be shared, see [Services that work with Resource Sharing](https://help.aliyun.com/document_detail/450526.html).
	//
	// 	- If the value of AssociationType is Target, the value of this parameter is `Account`.
	//
	// example:
	//
	// VSwitch
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the principal is outside the resource directory. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	External    *bool   `json:"External,omitempty" xml:"External,omitempty"`
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
	// example
	ResourceShareName *string `json:"ResourceShareName,omitempty" xml:"ResourceShareName,omitempty"`
	// The properties of the principal, such as the time range within which the resource is shared. Valid values of `timeRangeType`:
	//
	// 	- timeRange: a specific time range
	//
	// 	- day: all day
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
	// The time when the association of the entity was updated. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of `AssociationType` is `Resource`, the value of this parameter is the time when the association of the shared resource was updated.
	//
	// 	- If the value of `AssociationType` is `Target`, the value of this parameter is the time when the association of the principal was updated.
	//
	// example:
	//
	// 2020-12-07T07:39:02.920Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociations) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetAssociationFailedDetails() []*ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	return s.AssociationFailedDetails
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetAssociationStatus() *string {
	return s.AssociationStatus
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetAssociationStatusMessage() *string {
	return s.AssociationStatusMessage
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetAssociationType() *string {
	return s.AssociationType
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetEntityId() *string {
	return s.EntityId
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetExternal() *bool {
	return s.External
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetResourceProperty() *string {
	return s.ResourceProperty
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetResourceShareId() *string {
	return s.ResourceShareId
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetResourceShareName() *string {
	return s.ResourceShareName
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetTargetProperty() *string {
	return s.TargetProperty
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationFailedDetails(v []*ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationFailedDetails = v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationStatus(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationStatusMessage(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationStatusMessage = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetAssociationType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetCreateTime(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetEntityId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.EntityId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetEntityType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.EntityType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetExternal(v bool) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.External = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceArn(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceArn = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceProperty(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceProperty = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceShareId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceShareId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetResourceShareName(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.ResourceShareName = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetTargetProperty(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.TargetProperty = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) SetUpdateTime(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociations {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociations) Validate() error {
	if s.AssociationFailedDetails != nil {
		for _, item := range s.AssociationFailedDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails struct {
	// This parameter is deprecated. The OperationType parameter is used instead.
	//
	// example:
	//
	// None
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The ID of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of AssociationType is Resource, the value of this parameter is the ID of the principal.
	//
	// 	- If the value of AssociationType is Target, the value of this parameter is the ID of the resource.
	//
	// example:
	//
	// 172050525300****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The type of the entity. The value of this parameter depends on the value of the AssociationType parameter:
	//
	// 	- If the value of AssociationType is Resource, the value of this parameter is the type of the resource. For information about the types of resources that can be shared, see Services that work with Resource Sharing.
	//
	// 	- If the value of AssociationType is Target, the value of this parameter is `ResourceDirectory`, `Folder`, `Account`, or `Service`.
	//
	// example:
	//
	// Account
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The failure description.
	//
	// example:
	//
	// You cannot access the specified resource at this time.
	FailureDescription *string `json:"FailureDescription,omitempty" xml:"FailureDescription,omitempty"`
	// The failure cause. Valid values:
	//
	// 	- Unavailable: The resource does not exist.
	//
	// 	- LimitExceeded: The number of principals for the resource exceeds the upper limit.
	//
	// 	- ZonalResourceInaccessible: The resource is unavailable in this region.
	//
	// 	- InternalError: An internal error occurred.
	//
	// 	- UnsupportedOperation: You cannot perform this operation.
	//
	// example:
	//
	// Unavailable
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// The operation type. Valid values:
	//
	// 	- Associate
	//
	// 	- Disassociate
	//
	// example:
	//
	// Associate
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ResourceArn   *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// This parameter is deprecated. The FailureReason parameter is used instead.
	//
	// example:
	//
	// None
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is deprecated. The FailureDescription parameter is used instead.
	//
	// example:
	//
	// None
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetAssociateType() *string {
	return s.AssociateType
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetEntityId() *string {
	return s.EntityId
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetEntityType() *string {
	return s.EntityType
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetFailureDescription() *string {
	return s.FailureDescription
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetOperationType() *string {
	return s.OperationType
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetStatus() *string {
	return s.Status
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetAssociateType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.AssociateType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetEntityId(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.EntityId = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetEntityType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.EntityType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetFailureDescription(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.FailureDescription = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetFailureReason(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.FailureReason = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetOperationType(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.OperationType = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetResourceArn(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.ResourceArn = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetStatus(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.Status = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) SetStatusMessage(v string) *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails {
	s.StatusMessage = &v
	return s
}

func (s *ListResourceShareAssociationsResponseBodyResourceShareAssociationsAssociationFailedDetails) Validate() error {
	return dara.Validate(s)
}
