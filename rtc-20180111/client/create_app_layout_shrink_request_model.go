// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLayoutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppLayoutShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *CreateAppLayoutShrinkRequest
	GetClientToken() *string
	SetLayoutShrink(v string) *CreateAppLayoutShrinkRequest
	GetLayoutShrink() *string
}

type CreateAppLayoutShrinkRequest struct {
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

func (s CreateAppLayoutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppLayoutShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppLayoutShrinkRequest) GetLayoutShrink() *string {
	return s.LayoutShrink
}

func (s *CreateAppLayoutShrinkRequest) SetAppId(v string) *CreateAppLayoutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppLayoutShrinkRequest) SetClientToken(v string) *CreateAppLayoutShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppLayoutShrinkRequest) SetLayoutShrink(v string) *CreateAppLayoutShrinkRequest {
	s.LayoutShrink = &v
	return s
}

func (s *CreateAppLayoutShrinkRequest) Validate() error {
	return dara.Validate(s)
}
