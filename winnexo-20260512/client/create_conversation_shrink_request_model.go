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
	// 会话元数据，可含 model 等透传字段（model 需为合法抽象模型名，否则回退默认）
	//
	// example:
	//
	// string_value
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// 关联业务对象ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// operatingObjectName
	//
	// example:
	//
	// string_value
	OperatingObjectNameShrink *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
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
