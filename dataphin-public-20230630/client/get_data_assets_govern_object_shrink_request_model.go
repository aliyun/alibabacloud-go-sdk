// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAssetsGovernObjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandShrink(v string) *GetDataAssetsGovernObjectShrinkRequest
	GetCommandShrink() *string
	SetOpTenantId(v int64) *GetDataAssetsGovernObjectShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetDataAssetsGovernObjectShrinkRequest
	GetOpUserId() *string
}

type GetDataAssetsGovernObjectShrinkRequest struct {
	// The query instruction.
	//
	// This parameter is required.
	CommandShrink *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operation user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetDataAssetsGovernObjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectShrinkRequest) GetCommandShrink() *string {
	return s.CommandShrink
}

func (s *GetDataAssetsGovernObjectShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataAssetsGovernObjectShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetDataAssetsGovernObjectShrinkRequest) SetCommandShrink(v string) *GetDataAssetsGovernObjectShrinkRequest {
	s.CommandShrink = &v
	return s
}

func (s *GetDataAssetsGovernObjectShrinkRequest) SetOpTenantId(v int64) *GetDataAssetsGovernObjectShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataAssetsGovernObjectShrinkRequest) SetOpUserId(v string) *GetDataAssetsGovernObjectShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *GetDataAssetsGovernObjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
