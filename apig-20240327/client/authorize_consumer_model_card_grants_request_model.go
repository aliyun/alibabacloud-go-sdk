// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeConsumerModelCardGrantsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelCardIds(v []*string) *AuthorizeConsumerModelCardGrantsRequest
	GetModelCardIds() []*string
	SetModelGroupIds(v []*string) *AuthorizeConsumerModelCardGrantsRequest
	GetModelGroupIds() []*string
	SetPrincipalId(v string) *AuthorizeConsumerModelCardGrantsRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *AuthorizeConsumerModelCardGrantsRequest
	GetPrincipalType() *string
}

type AuthorizeConsumerModelCardGrantsRequest struct {
	// The collection of selected model card IDs for this API. This parameter can be empty or omitted, but at least one of modelCardIds and modelGroupIds must be non-empty.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
	// The actual model group IDs of the current instance. modelCardIds can be empty only when group licenses are granted. Group licenses allow zero intersection with the current API, but principal relationships with both target types empty cannot be saved.
	ModelGroupIds []*string `json:"modelGroupIds,omitempty" xml:"modelGroupIds,omitempty" type:"Repeated"`
	// The immutable ID of the authorization principal, not the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// cs-d00000000000000000000
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// The type of the authorization principal. Valid values: Consumer and ConsumerGroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// Consumer
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
}

func (s AuthorizeConsumerModelCardGrantsRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeConsumerModelCardGrantsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeConsumerModelCardGrantsRequest) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *AuthorizeConsumerModelCardGrantsRequest) GetModelGroupIds() []*string {
	return s.ModelGroupIds
}

func (s *AuthorizeConsumerModelCardGrantsRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *AuthorizeConsumerModelCardGrantsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AuthorizeConsumerModelCardGrantsRequest) SetModelCardIds(v []*string) *AuthorizeConsumerModelCardGrantsRequest {
	s.ModelCardIds = v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsRequest) SetModelGroupIds(v []*string) *AuthorizeConsumerModelCardGrantsRequest {
	s.ModelGroupIds = v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsRequest) SetPrincipalId(v string) *AuthorizeConsumerModelCardGrantsRequest {
	s.PrincipalId = &v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsRequest) SetPrincipalType(v string) *AuthorizeConsumerModelCardGrantsRequest {
	s.PrincipalType = &v
	return s
}

func (s *AuthorizeConsumerModelCardGrantsRequest) Validate() error {
	return dara.Validate(s)
}
