// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineBizEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineCommandShrink(v string) *OnlineBizEntityShrinkRequest
	GetOnlineCommandShrink() *string
	SetOpTenantId(v int64) *OnlineBizEntityShrinkRequest
	GetOpTenantId() *int64
}

type OnlineBizEntityShrinkRequest struct {
	// This parameter is required.
	OnlineCommandShrink *string `json:"OnlineCommand,omitempty" xml:"OnlineCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OnlineBizEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineBizEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *OnlineBizEntityShrinkRequest) GetOnlineCommandShrink() *string {
	return s.OnlineCommandShrink
}

func (s *OnlineBizEntityShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OnlineBizEntityShrinkRequest) SetOnlineCommandShrink(v string) *OnlineBizEntityShrinkRequest {
	s.OnlineCommandShrink = &v
	return s
}

func (s *OnlineBizEntityShrinkRequest) SetOpTenantId(v int64) *OnlineBizEntityShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *OnlineBizEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
