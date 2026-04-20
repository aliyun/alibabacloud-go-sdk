// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorksAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPoints(v int32) *AddWorksAuthorizationRequest
	GetAuthPoints() *int32
	SetAuthorizeScope(v int32) *AddWorksAuthorizationRequest
	GetAuthorizeScope() *int32
	SetAuthorizedId(v string) *AddWorksAuthorizationRequest
	GetAuthorizedId() *string
	SetExpireDay(v string) *AddWorksAuthorizationRequest
	GetExpireDay() *string
	SetResourceId(v string) *AddWorksAuthorizationRequest
	GetResourceId() *string
	SetResourceType(v string) *AddWorksAuthorizationRequest
	GetResourceType() *string
}

type AddWorksAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthPoints *int32 `json:"AuthPoints,omitempty" xml:"AuthPoints,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	AuthorizeScope *int32 `json:"AuthorizeScope,omitempty" xml:"AuthorizeScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ASDAS-WDAS****ASDA
	AuthorizedId *string `json:"AuthorizedId,omitempty" xml:"AuthorizedId,omitempty"`
	// example:
	//
	// 2099-12-31
	ExpireDay *string `json:"ExpireDay,omitempty" xml:"ExpireDay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// al*************7ufv
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dashboard
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AddWorksAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorksAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AddWorksAuthorizationRequest) GetAuthPoints() *int32 {
	return s.AuthPoints
}

func (s *AddWorksAuthorizationRequest) GetAuthorizeScope() *int32 {
	return s.AuthorizeScope
}

func (s *AddWorksAuthorizationRequest) GetAuthorizedId() *string {
	return s.AuthorizedId
}

func (s *AddWorksAuthorizationRequest) GetExpireDay() *string {
	return s.ExpireDay
}

func (s *AddWorksAuthorizationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddWorksAuthorizationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddWorksAuthorizationRequest) SetAuthPoints(v int32) *AddWorksAuthorizationRequest {
	s.AuthPoints = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetAuthorizeScope(v int32) *AddWorksAuthorizationRequest {
	s.AuthorizeScope = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetAuthorizedId(v string) *AddWorksAuthorizationRequest {
	s.AuthorizedId = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetExpireDay(v string) *AddWorksAuthorizationRequest {
	s.ExpireDay = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetResourceId(v string) *AddWorksAuthorizationRequest {
	s.ResourceId = &v
	return s
}

func (s *AddWorksAuthorizationRequest) SetResourceType(v string) *AddWorksAuthorizationRequest {
	s.ResourceType = &v
	return s
}

func (s *AddWorksAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
