// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceProjectWhiteListsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ReplaceProjectWhiteListsShrinkRequest
	GetId() *int64
	SetOpTenantId(v int64) *ReplaceProjectWhiteListsShrinkRequest
	GetOpTenantId() *int64
	SetReplaceCommandShrink(v string) *ReplaceProjectWhiteListsShrinkRequest
	GetReplaceCommandShrink() *string
}

type ReplaceProjectWhiteListsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ReplaceCommandShrink *string `json:"ReplaceCommand,omitempty" xml:"ReplaceCommand,omitempty"`
}

func (s ReplaceProjectWhiteListsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ReplaceProjectWhiteListsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ReplaceProjectWhiteListsShrinkRequest) GetReplaceCommandShrink() *string {
	return s.ReplaceCommandShrink
}

func (s *ReplaceProjectWhiteListsShrinkRequest) SetId(v int64) *ReplaceProjectWhiteListsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ReplaceProjectWhiteListsShrinkRequest) SetOpTenantId(v int64) *ReplaceProjectWhiteListsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ReplaceProjectWhiteListsShrinkRequest) SetReplaceCommandShrink(v string) *ReplaceProjectWhiteListsShrinkRequest {
	s.ReplaceCommandShrink = &v
	return s
}

func (s *ReplaceProjectWhiteListsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
