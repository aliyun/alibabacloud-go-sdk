// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBizEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfflineCommandShrink(v string) *OfflineBizEntityShrinkRequest
	GetOfflineCommandShrink() *string
	SetOpTenantId(v int64) *OfflineBizEntityShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *OfflineBizEntityShrinkRequest
	GetOpUserId() *string
}

type OfflineBizEntityShrinkRequest struct {
	// Offline request
	//
	// This parameter is required.
	OfflineCommandShrink *string `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty"`
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s OfflineBizEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineBizEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflineBizEntityShrinkRequest) GetOfflineCommandShrink() *string {
	return s.OfflineCommandShrink
}

func (s *OfflineBizEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflineBizEntityShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *OfflineBizEntityShrinkRequest) SetOfflineCommandShrink(v string) *OfflineBizEntityShrinkRequest {
	s.OfflineCommandShrink = &v
	return s
}

func (s *OfflineBizEntityShrinkRequest) SetOpTenantId(v int64) *OfflineBizEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflineBizEntityShrinkRequest) SetOpUserId(v string) *OfflineBizEntityShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *OfflineBizEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
