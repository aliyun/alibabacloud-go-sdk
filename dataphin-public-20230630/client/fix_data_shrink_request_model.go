// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *FixDataShrinkRequest
	GetEnv() *string
	SetFixDataCommandShrink(v string) *FixDataShrinkRequest
	GetFixDataCommandShrink() *string
	SetOpTenantId(v int64) *FixDataShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *FixDataShrinkRequest
	GetOpUserId() *string
}

type FixDataShrinkRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD (default): production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The command to rerun downstream nodes to fix data link issues. You can choose to force a rerun.
	//
	// This parameter is required.
	FixDataCommandShrink *string `json:"FixDataCommand,omitempty" xml:"FixDataCommand,omitempty"`
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

func (s FixDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FixDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *FixDataShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *FixDataShrinkRequest) GetFixDataCommandShrink() *string {
	return s.FixDataCommandShrink
}

func (s *FixDataShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *FixDataShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *FixDataShrinkRequest) SetEnv(v string) *FixDataShrinkRequest {
	s.Env = &v
	return s
}

func (s *FixDataShrinkRequest) SetFixDataCommandShrink(v string) *FixDataShrinkRequest {
	s.FixDataCommandShrink = &v
	return s
}

func (s *FixDataShrinkRequest) SetOpTenantId(v int64) *FixDataShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *FixDataShrinkRequest) SetOpUserId(v string) *FixDataShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *FixDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
