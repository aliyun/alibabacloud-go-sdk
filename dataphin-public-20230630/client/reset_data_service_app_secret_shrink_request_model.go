// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataServiceAppSecretShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ResetDataServiceAppSecretShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ResetDataServiceAppSecretShrinkRequest
	GetOpUserId() *string
	SetUpdateCommandShrink(v string) *ResetDataServiceAppSecretShrinkRequest
	GetUpdateCommandShrink() *string
}

type ResetDataServiceAppSecretShrinkRequest struct {
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
	// Reset Data Service Application Key
	//
	// This parameter is required.
	UpdateCommandShrink *string `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty"`
}

func (s ResetDataServiceAppSecretShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ResetDataServiceAppSecretShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ResetDataServiceAppSecretShrinkRequest) GetUpdateCommandShrink() *string {
	return s.UpdateCommandShrink
}

func (s *ResetDataServiceAppSecretShrinkRequest) SetOpTenantId(v int64) *ResetDataServiceAppSecretShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResetDataServiceAppSecretShrinkRequest) SetOpUserId(v string) *ResetDataServiceAppSecretShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ResetDataServiceAppSecretShrinkRequest) SetUpdateCommandShrink(v string) *ResetDataServiceAppSecretShrinkRequest {
	s.UpdateCommandShrink = &v
	return s
}

func (s *ResetDataServiceAppSecretShrinkRequest) Validate() error {
	return dara.Validate(s)
}
