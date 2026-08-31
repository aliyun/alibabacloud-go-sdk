// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetAttributesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetAssetAttributesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetAssetAttributesShrinkRequest
	GetOpUserId() *string
	SetQueryCommandShrink(v string) *GetAssetAttributesShrinkRequest
	GetQueryCommandShrink() *string
}

type GetAssetAttributesShrinkRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The query instruction.
	//
	// This parameter is required.
	QueryCommandShrink *string `json:"QueryCommand,omitempty" xml:"QueryCommand,omitempty"`
}

func (s GetAssetAttributesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAssetAttributesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetAssetAttributesShrinkRequest) GetQueryCommandShrink() *string {
	return s.QueryCommandShrink
}

func (s *GetAssetAttributesShrinkRequest) SetOpTenantId(v int64) *GetAssetAttributesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAssetAttributesShrinkRequest) SetOpUserId(v string) *GetAssetAttributesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetAssetAttributesShrinkRequest) SetQueryCommandShrink(v string) *GetAssetAttributesShrinkRequest {
	s.QueryCommandShrink = &v
	return s
}

func (s *GetAssetAttributesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
