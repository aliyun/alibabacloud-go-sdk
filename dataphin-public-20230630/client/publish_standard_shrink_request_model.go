// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStandardShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *PublishStandardShrinkRequest
	GetOpTenantId() *int64
	SetPublishCommandShrink(v string) *PublishStandardShrinkRequest
	GetPublishCommandShrink() *string
}

type PublishStandardShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PublishCommandShrink *string `json:"PublishCommand,omitempty" xml:"PublishCommand,omitempty"`
}

func (s PublishStandardShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishStandardShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishStandardShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *PublishStandardShrinkRequest) GetPublishCommandShrink() *string {
	return s.PublishCommandShrink
}

func (s *PublishStandardShrinkRequest) SetOpTenantId(v int64) *PublishStandardShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *PublishStandardShrinkRequest) SetPublishCommandShrink(v string) *PublishStandardShrinkRequest {
	s.PublishCommandShrink = &v
	return s
}

func (s *PublishStandardShrinkRequest) Validate() error {
	return dara.Validate(s)
}
