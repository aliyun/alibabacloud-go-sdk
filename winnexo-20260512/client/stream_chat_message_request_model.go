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
	// The last received SSE event ID, used for resuming delivery after a disconnection. If not specified, the full stream is replayed from the beginning.
	//
	// example:
	//
	// 5-1683456789012
	LastEventId *string `json:"lastEventId,omitempty" xml:"lastEventId,omitempty"`
	// The ID of the effective tenant.
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
