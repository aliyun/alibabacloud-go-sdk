// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliDingGroupMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *ListAliDingGroupMessagesRequest
	GetChatId() *string
	SetDirection(v string) *ListAliDingGroupMessagesRequest
	GetDirection() *string
	SetPageSize(v int32) *ListAliDingGroupMessagesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListAliDingGroupMessagesRequest
	GetTenantId() *string
	SetTime(v string) *ListAliDingGroupMessagesRequest
	GetTime() *string
}

type ListAliDingGroupMessagesRequest struct {
	// The session ID, typically used for JSSDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// cid-example
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The traffic direction. Valid values:
	//
	// - OutBound: outbound.
	//
	// - InBound: inbound.
	//
	// - Both: bidirectional.
	//
	// example:
	//
	// newer
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through `--tenant-id` in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The relationship information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-08T09:00:00+08:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ListAliDingGroupMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliDingGroupMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAliDingGroupMessagesRequest) GetChatId() *string {
	return s.ChatId
}

func (s *ListAliDingGroupMessagesRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListAliDingGroupMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliDingGroupMessagesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAliDingGroupMessagesRequest) GetTime() *string {
	return s.Time
}

func (s *ListAliDingGroupMessagesRequest) SetChatId(v string) *ListAliDingGroupMessagesRequest {
	s.ChatId = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetDirection(v string) *ListAliDingGroupMessagesRequest {
	s.Direction = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetPageSize(v int32) *ListAliDingGroupMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetTenantId(v string) *ListAliDingGroupMessagesRequest {
	s.TenantId = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) SetTime(v string) *ListAliDingGroupMessagesRequest {
	s.Time = &v
	return s
}

func (s *ListAliDingGroupMessagesRequest) Validate() error {
	return dara.Validate(s)
}
