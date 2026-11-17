// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeConsumerModelCardGrantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelCardIds(v []*string) *RevokeConsumerModelCardGrantsRequest
	GetModelCardIds() []*string
	SetModelGroupIds(v []*string) *RevokeConsumerModelCardGrantsRequest
	GetModelGroupIds() []*string
	SetPrincipalId(v string) *RevokeConsumerModelCardGrantsRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *RevokeConsumerModelCardGrantsRequest
	GetPrincipalType() *string
}

type RevokeConsumerModelCardGrantsRequest struct {
	// The IDs of the model cards for which entitlements are revoked in this API. Revoking non-existent entitlements is a valid no-op.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
	// The IDs of the actual model groups for which entitlements are revoked in this API. If you revoke only group entitlements, modelCardIds can be empty, and the group or principal is not required to exist. At least one of the two target types must be specified. This operation does not delete consumers, keys, or quotas.
	ModelGroupIds []*string `json:"modelGroupIds,omitempty" xml:"modelGroupIds,omitempty" type:"Repeated"`
	// The immutable ID of the principal for which entitlements are revoked. This is not the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// cs-d00000000000000000000
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// The type of the principal for which entitlements are revoked. Valid values:
	//
	// - Consumer
	//
	// - ConsumerGroup
	//
	// This parameter is required.
	//
	// example:
	//
	// Consumer
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
}

func (s RevokeConsumerModelCardGrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeConsumerModelCardGrantsRequest) GoString() string {
	return s.String()
}

func (s *RevokeConsumerModelCardGrantsRequest) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *RevokeConsumerModelCardGrantsRequest) GetModelGroupIds() []*string {
	return s.ModelGroupIds
}

func (s *RevokeConsumerModelCardGrantsRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *RevokeConsumerModelCardGrantsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *RevokeConsumerModelCardGrantsRequest) SetModelCardIds(v []*string) *RevokeConsumerModelCardGrantsRequest {
	s.ModelCardIds = v
	return s
}

func (s *RevokeConsumerModelCardGrantsRequest) SetModelGroupIds(v []*string) *RevokeConsumerModelCardGrantsRequest {
	s.ModelGroupIds = v
	return s
}

func (s *RevokeConsumerModelCardGrantsRequest) SetPrincipalId(v string) *RevokeConsumerModelCardGrantsRequest {
	s.PrincipalId = &v
	return s
}

func (s *RevokeConsumerModelCardGrantsRequest) SetPrincipalType(v string) *RevokeConsumerModelCardGrantsRequest {
	s.PrincipalType = &v
	return s
}

func (s *RevokeConsumerModelCardGrantsRequest) Validate() error {
	return dara.Validate(s)
}
