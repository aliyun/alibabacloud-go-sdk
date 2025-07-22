// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLayoutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppLayoutShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppLayoutShrinkRequest
	GetClientToken() *string
	SetLayoutShrink(v string) *ModifyAppLayoutShrinkRequest
	GetLayoutShrink() *string
}

type ModifyAppLayoutShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	LayoutShrink *string `json:"Layout,omitempty" xml:"Layout,omitempty"`
}

func (s ModifyAppLayoutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppLayoutShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppLayoutShrinkRequest) GetLayoutShrink() *string {
	return s.LayoutShrink
}

func (s *ModifyAppLayoutShrinkRequest) SetAppId(v string) *ModifyAppLayoutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppLayoutShrinkRequest) SetClientToken(v string) *ModifyAppLayoutShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppLayoutShrinkRequest) SetLayoutShrink(v string) *ModifyAppLayoutShrinkRequest {
	s.LayoutShrink = &v
	return s
}

func (s *ModifyAppLayoutShrinkRequest) Validate() error {
	return dara.Validate(s)
}
