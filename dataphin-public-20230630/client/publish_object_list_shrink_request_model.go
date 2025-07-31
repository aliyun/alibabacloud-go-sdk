// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishObjectListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishObjectListShrinkRequest
	GetOpTenantId() *int64
	SetPublishCommandShrink(v string) *PublishObjectListShrinkRequest
	GetPublishCommandShrink() *string
}

type PublishObjectListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PublishCommandShrink *string `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty"`
}

func (s PublishObjectListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishObjectListShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishObjectListShrinkRequest) GetPublishCommandShrink() *string {
	return s.PublishCommandShrink
}

func (s *PublishObjectListShrinkRequest) SetOpTenantId(v int64) *PublishObjectListShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishObjectListShrinkRequest) SetPublishCommandShrink(v string) *PublishObjectListShrinkRequest {
	s.PublishCommandShrink = &v
	return s
}

func (s *PublishObjectListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
