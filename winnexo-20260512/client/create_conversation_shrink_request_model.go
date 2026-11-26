// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateConversationShrinkRequest
	GetAppToken() *string
	SetMetadata(v string) *CreateConversationShrinkRequest
	GetMetadata() *string
	SetObjectId(v string) *CreateConversationShrinkRequest
	GetObjectId() *string
	SetOperatingObjectNameShrink(v string) *CreateConversationShrinkRequest
	GetOperatingObjectNameShrink() *string
	SetOpsUserCode(v string) *CreateConversationShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateConversationShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateConversationShrinkRequest
	GetTenantId() *string
}

type CreateConversationShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The reserved field.
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
	// The name of the operation object.
	//
	// example:
	//
	// ["item name"]
	OperatingObjectNameShrink *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
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

func (s *CreateConversationShrinkRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreateConversationShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateConversationShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateConversationShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateConversationShrinkRequest) SetAppToken(v string) *CreateConversationShrinkRequest {
	s.AppToken = &v
	return s
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

func (s *CreateConversationShrinkRequest) SetOpsUserCode(v string) *CreateConversationShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateConversationShrinkRequest) SetOpsUserId(v string) *CreateConversationShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateConversationShrinkRequest) SetTenantId(v string) *CreateConversationShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateConversationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
