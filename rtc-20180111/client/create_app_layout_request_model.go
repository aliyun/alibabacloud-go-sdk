// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppLayoutRequest
	GetAppId() *string
	SetClientToken(v string) *CreateAppLayoutRequest
	GetClientToken() *string
	SetLayout(v *CreateAppLayoutRequestLayout) *CreateAppLayoutRequest
	GetLayout() *CreateAppLayoutRequestLayout
}

type CreateAppLayoutRequest struct {
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
	Layout *CreateAppLayoutRequestLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
}

func (s CreateAppLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutRequest) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppLayoutRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppLayoutRequest) GetLayout() *CreateAppLayoutRequestLayout {
	return s.Layout
}

func (s *CreateAppLayoutRequest) SetAppId(v string) *CreateAppLayoutRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppLayoutRequest) SetClientToken(v string) *CreateAppLayoutRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppLayoutRequest) SetLayout(v *CreateAppLayoutRequestLayout) *CreateAppLayoutRequest {
	s.Layout = v
	return s
}

func (s *CreateAppLayoutRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAppLayoutRequestLayout struct {
	// This parameter is required.
	//
	// example:
	//
	// 布局
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Panes []*CreateAppLayoutRequestLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s CreateAppLayoutRequestLayout) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutRequestLayout) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutRequestLayout) GetName() *string {
	return s.Name
}

func (s *CreateAppLayoutRequestLayout) GetPanes() []*CreateAppLayoutRequestLayoutPanes {
	return s.Panes
}

func (s *CreateAppLayoutRequestLayout) SetName(v string) *CreateAppLayoutRequestLayout {
	s.Name = &v
	return s
}

func (s *CreateAppLayoutRequestLayout) SetPanes(v []*CreateAppLayoutRequestLayoutPanes) *CreateAppLayoutRequestLayout {
	s.Panes = v
	return s
}

func (s *CreateAppLayoutRequestLayout) Validate() error {
	return dara.Validate(s)
}

type CreateAppLayoutRequestLayoutPanes struct {
	// example:
	//
	// 0.25
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	PaneId *int64 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0.25
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.25
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.25
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s CreateAppLayoutRequestLayoutPanes) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutRequestLayoutPanes) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutRequestLayoutPanes) GetHeight() *float64 {
	return s.Height
}

func (s *CreateAppLayoutRequestLayoutPanes) GetPaneId() *int64 {
	return s.PaneId
}

func (s *CreateAppLayoutRequestLayoutPanes) GetWidth() *float64 {
	return s.Width
}

func (s *CreateAppLayoutRequestLayoutPanes) GetX() *float64 {
	return s.X
}

func (s *CreateAppLayoutRequestLayoutPanes) GetY() *float64 {
	return s.Y
}

func (s *CreateAppLayoutRequestLayoutPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *CreateAppLayoutRequestLayoutPanes) SetHeight(v float64) *CreateAppLayoutRequestLayoutPanes {
	s.Height = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) SetPaneId(v int64) *CreateAppLayoutRequestLayoutPanes {
	s.PaneId = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) SetWidth(v float64) *CreateAppLayoutRequestLayoutPanes {
	s.Width = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) SetX(v float64) *CreateAppLayoutRequestLayoutPanes {
	s.X = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) SetY(v float64) *CreateAppLayoutRequestLayoutPanes {
	s.Y = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) SetZOrder(v int32) *CreateAppLayoutRequestLayoutPanes {
	s.ZOrder = &v
	return s
}

func (s *CreateAppLayoutRequestLayoutPanes) Validate() error {
	return dara.Validate(s)
}
