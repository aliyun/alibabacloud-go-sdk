// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v string) *CreateConversationShrinkRequest
	GetMetadata() *string
	SetObjectId(v string) *CreateConversationShrinkRequest
	GetObjectId() *string
	SetOperatingObjectNameShrink(v string) *CreateConversationShrinkRequest
	GetOperatingObjectNameShrink() *string
	SetTenantId(v string) *CreateConversationShrinkRequest
	GetTenantId() *string
}

type CreateConversationShrinkRequest struct {
	// A reserved field.
	//
	// example:
	//
	// string_value
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The primary key ID of the associated variable.
	//
	// example:
	//
	// 2676
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The operating object name.
	//
	// example:
	//
	// string_value
	OperatingObjectNameShrink *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateConversationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConversationShrinkRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateConversationShrinkRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateConversationShrinkRequest) GetOperatingObjectNameShrink() *string {
	return s.OperatingObjectNameShrink
}

func (s *CreateConversationShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateConversationShrinkRequest) SetMetadata(v string) *CreateConversationShrinkRequest {
	s.Metadata = &v
	return s
}

func (s *CreateConversationShrinkRequest) SetObjectId(v string) *CreateConversationShrinkRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateConversationShrinkRequest) SetOperatingObjectNameShrink(v string) *CreateConversationShrinkRequest {
	s.OperatingObjectNameShrink = &v
	return s
}

func (s *CreateConversationShrinkRequest) SetTenantId(v string) *CreateConversationShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateConversationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
