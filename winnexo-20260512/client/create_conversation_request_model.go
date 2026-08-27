// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v string) *CreateConversationRequest
	GetMetadata() *string
	SetObjectId(v string) *CreateConversationRequest
	GetObjectId() *string
	SetOperatingObjectName(v []interface{}) *CreateConversationRequest
	GetOperatingObjectName() []interface{}
	SetTenantId(v string) *CreateConversationRequest
	GetTenantId() *string
}

type CreateConversationRequest struct {
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
	OperatingObjectName []interface{} `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateConversationRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateConversationRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateConversationRequest) GetOperatingObjectName() []interface{} {
	return s.OperatingObjectName
}

func (s *CreateConversationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateConversationRequest) SetMetadata(v string) *CreateConversationRequest {
	s.Metadata = &v
	return s
}

func (s *CreateConversationRequest) SetObjectId(v string) *CreateConversationRequest {
	s.ObjectId = &v
	return s
}

func (s *CreateConversationRequest) SetOperatingObjectName(v []interface{}) *CreateConversationRequest {
	s.OperatingObjectName = v
	return s
}

func (s *CreateConversationRequest) SetTenantId(v string) *CreateConversationRequest {
	s.TenantId = &v
	return s
}

func (s *CreateConversationRequest) Validate() error {
	return dara.Validate(s)
}
