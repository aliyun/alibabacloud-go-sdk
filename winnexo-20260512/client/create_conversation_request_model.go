// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateConversationRequest
	GetAppToken() *string
	SetMetadata(v string) *CreateConversationRequest
	GetMetadata() *string
	SetObjectId(v string) *CreateConversationRequest
	GetObjectId() *string
	SetOperatingObjectName(v []interface{}) *CreateConversationRequest
	GetOperatingObjectName() []interface{}
	SetOpsUserCode(v string) *CreateConversationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateConversationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateConversationRequest
	GetTenantId() *string
}

type CreateConversationRequest struct {
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
	OperatingObjectName []interface{} `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty" type:"Repeated"`
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

func (s CreateConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateConversationRequest) GetAppToken() *string {
	return s.AppToken
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

func (s *CreateConversationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateConversationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateConversationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateConversationRequest) SetAppToken(v string) *CreateConversationRequest {
	s.AppToken = &v
	return s
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

func (s *CreateConversationRequest) SetOpsUserCode(v string) *CreateConversationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateConversationRequest) SetOpsUserId(v string) *CreateConversationRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateConversationRequest) SetTenantId(v string) *CreateConversationRequest {
	s.TenantId = &v
	return s
}

func (s *CreateConversationRequest) Validate() error {
	return dara.Validate(s)
}
