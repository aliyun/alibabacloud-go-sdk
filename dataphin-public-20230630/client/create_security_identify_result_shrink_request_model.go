// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIdentifyResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommandShrink(v string) *CreateSecurityIdentifyResultShrinkRequest
	GetCreateCommandShrink() *string
	SetOpTenantId(v int64) *CreateSecurityIdentifyResultShrinkRequest
	GetOpTenantId() *int64
}

type CreateSecurityIdentifyResultShrinkRequest struct {
	// This parameter is required.
	CreateCommandShrink *string `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateSecurityIdentifyResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIdentifyResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityIdentifyResultShrinkRequest) GetCreateCommandShrink() *string {
	return s.CreateCommandShrink
}

func (s *CreateSecurityIdentifyResultShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateSecurityIdentifyResultShrinkRequest) SetCreateCommandShrink(v string) *CreateSecurityIdentifyResultShrinkRequest {
	s.CreateCommandShrink = &v
	return s
}

func (s *CreateSecurityIdentifyResultShrinkRequest) SetOpTenantId(v int64) *CreateSecurityIdentifyResultShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateSecurityIdentifyResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
