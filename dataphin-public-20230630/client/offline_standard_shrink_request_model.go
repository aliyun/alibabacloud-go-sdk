// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfflineCommandShrink(v string) *OfflineStandardShrinkRequest
	GetOfflineCommandShrink() *string
	SetOpTenantId(v int64) *OfflineStandardShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *OfflineStandardShrinkRequest
	GetOpUserId() *string
}

type OfflineStandardShrinkRequest struct {
	// The offline instruction.
	//
	// This parameter is required.
	OfflineCommandShrink *string `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty"`
	// The tenant ID.
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

func (s OfflineStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *OfflineStandardShrinkRequest) GetOfflineCommandShrink() *string {
	return s.OfflineCommandShrink
}

func (s *OfflineStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflineStandardShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *OfflineStandardShrinkRequest) SetOfflineCommandShrink(v string) *OfflineStandardShrinkRequest {
	s.OfflineCommandShrink = &v
	return s
}

func (s *OfflineStandardShrinkRequest) SetOpTenantId(v int64) *OfflineStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflineStandardShrinkRequest) SetOpUserId(v string) *OfflineStandardShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *OfflineStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
