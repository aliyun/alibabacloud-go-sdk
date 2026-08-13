// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLastEventId(v string) *StreamChatMessageRequest
	GetLastEventId() *string
	SetTenantId(v string) *StreamChatMessageRequest
	GetTenantId() *string
}

type StreamChatMessageRequest struct {
	// 上次接收到的 SSE event id，用于断线续推；不传则从头全量回放
	//
	// example:
	//
	// 5-1683456789012
	LastEventId *string `json:"lastEventId,omitempty" xml:"lastEventId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s StreamChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s StreamChatMessageRequest) GoString() string {
	return s.String()
}

func (s *StreamChatMessageRequest) GetLastEventId() *string {
	return s.LastEventId
}

func (s *StreamChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *StreamChatMessageRequest) SetLastEventId(v string) *StreamChatMessageRequest {
	s.LastEventId = &v
	return s
}

func (s *StreamChatMessageRequest) SetTenantId(v string) *StreamChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *StreamChatMessageRequest) Validate() error {
	return dara.Validate(s)
}
