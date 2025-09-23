// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationStatus(v string) *ListResourceShareAssociationsRequest
	GetAssociationStatus() *string
	SetAssociationType(v string) *ListResourceShareAssociationsRequest
	GetAssociationType() *string
	SetMaxResults(v int32) *ListResourceShareAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceShareAssociationsRequest
	GetNextToken() *string
	SetResourceArn(v string) *ListResourceShareAssociationsRequest
	GetResourceArn() *string
	SetResourceId(v string) *ListResourceShareAssociationsRequest
	GetResourceId() *string
	SetResourceShareIds(v []*string) *ListResourceShareAssociationsRequest
	GetResourceShareIds() []*string
	SetTarget(v string) *ListResourceShareAssociationsRequest
	GetTarget() *string
}

type ListResourceShareAssociationsRequest struct {
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
	// The association type. Valid values:
	//
	// 	- Resource
	//
	// 	- Target
	//
	// This parameter is required.
	//
	// example:
	//
	// Resource
	AssociationType *string `json:"AssociationType,omitempty" xml:"AssociationType,omitempty"`
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The `token` that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the resource.
	//
	// >  This parameter is unavailable if you set the `AssociationType` parameter to `Target`.
	//
	// example:
	//
	// vsw-bp183p93qs667muql****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The IDs of the resource shares.
	//
	// Valid values of N: 1 to 5. This indicates that a maximum of five resource shares can be specified at a time.
	//
	// example:
	//
	// rs-6GRmdD3X****
	ResourceShareIds []*string `json:"ResourceShareIds,omitempty" xml:"ResourceShareIds,omitempty" type:"Repeated"`
	// The ID of the principal.
	//
	// >  This parameter is unavailable if you set the `AssociationType` parameter to `Resource`.
	//
	// example:
	//
	// 172050525300****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ListResourceShareAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceShareAssociationsRequest) GetAssociationStatus() *string {
	return s.AssociationStatus
}

func (s *ListResourceShareAssociationsRequest) GetAssociationType() *string {
	return s.AssociationType
}

func (s *ListResourceShareAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceShareAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceShareAssociationsRequest) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourceShareAssociationsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceShareAssociationsRequest) GetResourceShareIds() []*string {
	return s.ResourceShareIds
}

func (s *ListResourceShareAssociationsRequest) GetTarget() *string {
	return s.Target
}

func (s *ListResourceShareAssociationsRequest) SetAssociationStatus(v string) *ListResourceShareAssociationsRequest {
	s.AssociationStatus = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetAssociationType(v string) *ListResourceShareAssociationsRequest {
	s.AssociationType = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetMaxResults(v int32) *ListResourceShareAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetNextToken(v string) *ListResourceShareAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceArn(v string) *ListResourceShareAssociationsRequest {
	s.ResourceArn = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceId(v string) *ListResourceShareAssociationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetResourceShareIds(v []*string) *ListResourceShareAssociationsRequest {
	s.ResourceShareIds = v
	return s
}

func (s *ListResourceShareAssociationsRequest) SetTarget(v string) *ListResourceShareAssociationsRequest {
	s.Target = &v
	return s
}

func (s *ListResourceShareAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
