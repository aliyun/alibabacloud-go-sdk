// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLayoutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppLayoutShrinkRequest
	GetAppId() *string
	SetClientToken(v string) *DeleteAppLayoutShrinkRequest
	GetClientToken() *string
	SetLayoutShrink(v string) *DeleteAppLayoutShrinkRequest
	GetLayoutShrink() *string
}

type DeleteAppLayoutShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	LayoutShrink *string `json:"Layout,omitempty" xml:"Layout,omitempty"`
}

func (s DeleteAppLayoutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLayoutShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppLayoutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppLayoutShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppLayoutShrinkRequest) GetLayoutShrink() *string {
	return s.LayoutShrink
}

func (s *DeleteAppLayoutShrinkRequest) SetAppId(v string) *DeleteAppLayoutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppLayoutShrinkRequest) SetClientToken(v string) *DeleteAppLayoutShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppLayoutShrinkRequest) SetLayoutShrink(v string) *DeleteAppLayoutShrinkRequest {
	s.LayoutShrink = &v
	return s
}

func (s *DeleteAppLayoutShrinkRequest) Validate() error {
	return dara.Validate(s)
}
