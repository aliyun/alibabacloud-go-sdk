// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendName(v string) *CreateBackendRequest
	GetBackendName() *string
	SetBackendType(v string) *CreateBackendRequest
	GetBackendType() *string
	SetCreateEventBridgeServiceLinkedRole(v bool) *CreateBackendRequest
	GetCreateEventBridgeServiceLinkedRole() *bool
	SetCreateSlr(v bool) *CreateBackendRequest
	GetCreateSlr() *bool
	SetDescription(v string) *CreateBackendRequest
	GetDescription() *string
	SetSecurityToken(v string) *CreateBackendRequest
	GetSecurityToken() *string
	SetTag(v []*CreateBackendRequestTag) *CreateBackendRequest
	GetTag() []*CreateBackendRequestTag
}

type CreateBackendRequest struct {
	// The name of the backend service.
	//
	// This parameter is required.
	//
	// example:
	//
	// testBackendService
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The type of the backend service.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
	// Specifies to create a EventBridge service-linked role.
	//
	// example:
	//
	// true
	CreateEventBridgeServiceLinkedRole *bool `json:"CreateEventBridgeServiceLinkedRole,omitempty" xml:"CreateEventBridgeServiceLinkedRole,omitempty"`
	// Specifies to create a service-linked role.
	//
	// example:
	//
	// true
	CreateSlr *bool `json:"CreateSlr,omitempty" xml:"CreateSlr,omitempty"`
	// The description.
	//
	// example:
	//
	// release data api 411055691504981
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*CreateBackendRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateBackendRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendRequest) GoString() string {
	return s.String()
}

func (s *CreateBackendRequest) GetBackendName() *string {
	return s.BackendName
}

func (s *CreateBackendRequest) GetBackendType() *string {
	return s.BackendType
}

func (s *CreateBackendRequest) GetCreateEventBridgeServiceLinkedRole() *bool {
	return s.CreateEventBridgeServiceLinkedRole
}

func (s *CreateBackendRequest) GetCreateSlr() *bool {
	return s.CreateSlr
}

func (s *CreateBackendRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBackendRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateBackendRequest) GetTag() []*CreateBackendRequestTag {
	return s.Tag
}

func (s *CreateBackendRequest) SetBackendName(v string) *CreateBackendRequest {
	s.BackendName = &v
	return s
}

func (s *CreateBackendRequest) SetBackendType(v string) *CreateBackendRequest {
	s.BackendType = &v
	return s
}

func (s *CreateBackendRequest) SetCreateEventBridgeServiceLinkedRole(v bool) *CreateBackendRequest {
	s.CreateEventBridgeServiceLinkedRole = &v
	return s
}

func (s *CreateBackendRequest) SetCreateSlr(v bool) *CreateBackendRequest {
	s.CreateSlr = &v
	return s
}

func (s *CreateBackendRequest) SetDescription(v string) *CreateBackendRequest {
	s.Description = &v
	return s
}

func (s *CreateBackendRequest) SetSecurityToken(v string) *CreateBackendRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateBackendRequest) SetTag(v []*CreateBackendRequestTag) *CreateBackendRequest {
	s.Tag = v
	return s
}

func (s *CreateBackendRequest) Validate() error {
	return dara.Validate(s)
}

type CreateBackendRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateBackendRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateBackendRequestTag) GoString() string {
	return s.String()
}

func (s *CreateBackendRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateBackendRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateBackendRequestTag) SetKey(v string) *CreateBackendRequestTag {
	s.Key = &v
	return s
}

func (s *CreateBackendRequestTag) SetValue(v string) *CreateBackendRequestTag {
	s.Value = &v
	return s
}

func (s *CreateBackendRequestTag) Validate() error {
	return dara.Validate(s)
}
