// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppLayoutRequest
	GetAppId() *string
	SetClientToken(v string) *DeleteAppLayoutRequest
	GetClientToken() *string
	SetLayout(v *DeleteAppLayoutRequestLayout) *DeleteAppLayoutRequest
	GetLayout() *DeleteAppLayoutRequestLayout
}

type DeleteAppLayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Layout      *DeleteAppLayoutRequestLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
}

func (s DeleteAppLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppLayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppLayoutRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppLayoutRequest) GetLayout() *DeleteAppLayoutRequestLayout {
	return s.Layout
}

func (s *DeleteAppLayoutRequest) SetAppId(v string) *DeleteAppLayoutRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppLayoutRequest) SetClientToken(v string) *DeleteAppLayoutRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppLayoutRequest) SetLayout(v *DeleteAppLayoutRequestLayout) *DeleteAppLayoutRequest {
	s.Layout = v
	return s
}

func (s *DeleteAppLayoutRequest) Validate() error {
	if s.Layout != nil {
		if err := s.Layout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAppLayoutRequestLayout struct {
	// This parameter is required.
	//
	// example:
	//
	// 167466539798442****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s DeleteAppLayoutRequestLayout) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLayoutRequestLayout) GoString() string {
	return s.String()
}

func (s *DeleteAppLayoutRequestLayout) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DeleteAppLayoutRequestLayout) SetLayoutId(v string) *DeleteAppLayoutRequestLayout {
	s.LayoutId = &v
	return s
}

func (s *DeleteAppLayoutRequestLayout) Validate() error {
	return dara.Validate(s)
}
