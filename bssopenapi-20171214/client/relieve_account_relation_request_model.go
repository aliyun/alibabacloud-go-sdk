// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelieveAccountRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildUserId(v int64) *RelieveAccountRelationRequest
	GetChildUserId() *int64
	SetParentUserId(v int64) *RelieveAccountRelationRequest
	GetParentUserId() *int64
	SetRelationId(v int64) *RelieveAccountRelationRequest
	GetRelationId() *int64
	SetRelationType(v string) *RelieveAccountRelationRequest
	GetRelationType() *string
	SetRequestId(v string) *RelieveAccountRelationRequest
	GetRequestId() *string
}

type RelieveAccountRelationRequest struct {
	// The UID of the RAM user. You must specify either the RelationId parameter, or the ParentUserId, ChildUserId, and RelationType parameters.
	//
	// example:
	//
	// 1512996702208737
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The UID of the Alibaba Cloud account. You must specify either the RelationId parameter, or the ParentUserId, ChildUserId, and RelationType parameters.
	//
	// example:
	//
	// 1738376485192612
	ParentUserId *int64 `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The ID of the account association. You must specify either the RelationId parameter, or the ParentUserId, ChildUserId, and RelationType parameters.
	//
	// example:
	//
	// RelationId
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type. Only enterprise_group is supported.
	//
	// example:
	//
	// enterprise_group
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The unique ID of the request. This ID is used to troubleshoot issues.
	//
	// This parameter is required.
	//
	// example:
	//
	// request_id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RelieveAccountRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s RelieveAccountRelationRequest) GoString() string {
	return s.String()
}

func (s *RelieveAccountRelationRequest) GetChildUserId() *int64 {
	return s.ChildUserId
}

func (s *RelieveAccountRelationRequest) GetParentUserId() *int64 {
	return s.ParentUserId
}

func (s *RelieveAccountRelationRequest) GetRelationId() *int64 {
	return s.RelationId
}

func (s *RelieveAccountRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *RelieveAccountRelationRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *RelieveAccountRelationRequest) SetChildUserId(v int64) *RelieveAccountRelationRequest {
	s.ChildUserId = &v
	return s
}

func (s *RelieveAccountRelationRequest) SetParentUserId(v int64) *RelieveAccountRelationRequest {
	s.ParentUserId = &v
	return s
}

func (s *RelieveAccountRelationRequest) SetRelationId(v int64) *RelieveAccountRelationRequest {
	s.RelationId = &v
	return s
}

func (s *RelieveAccountRelationRequest) SetRelationType(v string) *RelieveAccountRelationRequest {
	s.RelationType = &v
	return s
}

func (s *RelieveAccountRelationRequest) SetRequestId(v string) *RelieveAccountRelationRequest {
	s.RequestId = &v
	return s
}

func (s *RelieveAccountRelationRequest) Validate() error {
	return dara.Validate(s)
}
