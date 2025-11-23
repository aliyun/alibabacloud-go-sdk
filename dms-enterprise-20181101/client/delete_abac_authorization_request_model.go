// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAbacAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *DeleteAbacAuthorizationRequest
	GetAuthorizationId() *int64
	SetIdentityType(v string) *DeleteAbacAuthorizationRequest
	GetIdentityType() *string
	SetTid(v int64) *DeleteAbacAuthorizationRequest
	GetTid() *int64
}

type DeleteAbacAuthorizationRequest struct {
	// The authorization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123***
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The type of object to which you want to attach the policy.********
	//
	// Valid values:
	//
	// 	- USER
	//
	// 	- ROLE
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteAbacAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAbacAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAbacAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *DeleteAbacAuthorizationRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *DeleteAbacAuthorizationRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteAbacAuthorizationRequest) SetAuthorizationId(v int64) *DeleteAbacAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *DeleteAbacAuthorizationRequest) SetIdentityType(v string) *DeleteAbacAuthorizationRequest {
	s.IdentityType = &v
	return s
}

func (s *DeleteAbacAuthorizationRequest) SetTid(v int64) *DeleteAbacAuthorizationRequest {
	s.Tid = &v
	return s
}

func (s *DeleteAbacAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
