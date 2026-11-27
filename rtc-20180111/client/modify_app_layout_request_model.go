// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppLayoutRequest
	GetAppId() *string
	SetClientToken(v string) *ModifyAppLayoutRequest
	GetClientToken() *string
	SetLayout(v *ModifyAppLayoutRequestLayout) *ModifyAppLayoutRequest
	GetLayout() *ModifyAppLayoutRequestLayout
}

type ModifyAppLayoutRequest struct {
	// Application ID. Create and query through the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Idempotent parameter.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Layout.
	//
	// This parameter is required.
	Layout *ModifyAppLayoutRequestLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
}

func (s ModifyAppLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppLayoutRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppLayoutRequest) GetLayout() *ModifyAppLayoutRequestLayout {
	return s.Layout
}

func (s *ModifyAppLayoutRequest) SetAppId(v string) *ModifyAppLayoutRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppLayoutRequest) SetClientToken(v string) *ModifyAppLayoutRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppLayoutRequest) SetLayout(v *ModifyAppLayoutRequestLayout) *ModifyAppLayoutRequest {
	s.Layout = v
	return s
}

func (s *ModifyAppLayoutRequest) Validate() error {
	if s.Layout != nil {
		if err := s.Layout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppLayoutRequestLayout struct {
	// Layout ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123121231313
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// Layout name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of panes.
	Panes []*ModifyAppLayoutRequestLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s ModifyAppLayoutRequestLayout) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutRequestLayout) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutRequestLayout) GetLayoutId() *string {
	return s.LayoutId
}

func (s *ModifyAppLayoutRequestLayout) GetName() *string {
	return s.Name
}

func (s *ModifyAppLayoutRequestLayout) GetPanes() []*ModifyAppLayoutRequestLayoutPanes {
	return s.Panes
}

func (s *ModifyAppLayoutRequestLayout) SetLayoutId(v string) *ModifyAppLayoutRequestLayout {
	s.LayoutId = &v
	return s
}

func (s *ModifyAppLayoutRequestLayout) SetName(v string) *ModifyAppLayoutRequestLayout {
	s.Name = &v
	return s
}

func (s *ModifyAppLayoutRequestLayout) SetPanes(v []*ModifyAppLayoutRequestLayoutPanes) *ModifyAppLayoutRequestLayout {
	s.Panes = v
	return s
}

func (s *ModifyAppLayoutRequestLayout) Validate() error {
	if s.Panes != nil {
		for _, item := range s.Panes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAppLayoutRequestLayoutPanes struct {
	// Pane height, normalized percentage. Height ranges from 0.0\<Height≤1.0.
	//
	// example:
	//
	// 0.25
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Pane ID, starting from 0.
	//
	// example:
	//
	// 0
	PaneId *int64 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// Pane width, normalized percentage. Width ranges from 0.0\<Width≤1.0.
	//
	// example:
	//
	// 0.25
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// X coordinate, normalized percentage. The X coordinate ranges from 0.0 to 1.0.
	//
	// example:
	//
	// 0.25
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// Y coordinate, normalized percentage. The Y coordinate ranges from 0.0 to 1.0.
	//
	// example:
	//
	// 0.25
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
	// Stacking order. 0 is the bottom layer, 1 is above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s ModifyAppLayoutRequestLayoutPanes) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLayoutRequestLayoutPanes) GoString() string {
	return s.String()
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetPaneId() *int64 {
	return s.PaneId
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetX() *float64 {
	return s.X
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetY() *float64 {
	return s.Y
}

func (s *ModifyAppLayoutRequestLayoutPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetHeight(v float64) *ModifyAppLayoutRequestLayoutPanes {
	s.Height = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetPaneId(v int64) *ModifyAppLayoutRequestLayoutPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetWidth(v float64) *ModifyAppLayoutRequestLayoutPanes {
	s.Width = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetX(v float64) *ModifyAppLayoutRequestLayoutPanes {
	s.X = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetY(v float64) *ModifyAppLayoutRequestLayoutPanes {
	s.Y = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) SetZOrder(v int32) *ModifyAppLayoutRequestLayoutPanes {
	s.ZOrder = &v
	return s
}

func (s *ModifyAppLayoutRequestLayoutPanes) Validate() error {
	return dara.Validate(s)
}
