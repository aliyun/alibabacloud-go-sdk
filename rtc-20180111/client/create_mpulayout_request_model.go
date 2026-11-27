// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMPULayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMPULayoutRequest
	GetAppId() *string
	SetAudioMixCount(v int32) *CreateMPULayoutRequest
	GetAudioMixCount() *int32
	SetName(v string) *CreateMPULayoutRequest
	GetName() *string
	SetOwnerId(v int64) *CreateMPULayoutRequest
	GetOwnerId() *int64
	SetPanes(v []*CreateMPULayoutRequestPanes) *CreateMPULayoutRequest
	GetPanes() []*CreateMPULayoutRequestPanes
}

type CreateMPULayoutRequest struct {
	// The application ID. Only a single ID is supported. You can create and query application IDs in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The maximum number of audio mixes. Valid values: **0*	- to **6**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// The name of the layout.
	//
	// example:
	//
	// LayoutName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of panes in the layout.
	//
	// This parameter is required.
	Panes []*CreateMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s CreateMPULayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMPULayoutRequest) GetAudioMixCount() *int32 {
	return s.AudioMixCount
}

func (s *CreateMPULayoutRequest) GetName() *string {
	return s.Name
}

func (s *CreateMPULayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMPULayoutRequest) GetPanes() []*CreateMPULayoutRequestPanes {
	return s.Panes
}

func (s *CreateMPULayoutRequest) SetAppId(v string) *CreateMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetAudioMixCount(v int32) *CreateMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *CreateMPULayoutRequest) SetName(v string) *CreateMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *CreateMPULayoutRequest) SetOwnerId(v int64) *CreateMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetPanes(v []*CreateMPULayoutRequestPanes) *CreateMPULayoutRequest {
	s.Panes = v
	return s
}

func (s *CreateMPULayoutRequest) Validate() error {
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

type CreateMPULayoutRequestPanes struct {
	// The height of the pane, represented as a normalized percentage. Valid values: **0.0 < Height ≤ 1.0**.
	//
	// example:
	//
	// 0.25
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies the main pane. Each layout can have only one main pane.
	//
	// - **0**: secondary pane.
	//
	//
	//
	// - **1**: main pane.
	//
	// example:
	//
	// 1
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// The pane ID. Panes are sorted from top left to bottom right, starting from 0.
	//
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// The width of the pane, represented as a normalized percentage. Valid values: **0.0 < Width ≤ 1.0**.
	//
	// example:
	//
	// 0.25
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, represented as a normalized percentage. Valid values: **0.0 ≤ X ≤ 1.0**.
	//
	// example:
	//
	// 0.25
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, represented as a normalized percentage. Valid values: **0.0 ≤ Y ≤ 1.0**.
	//
	// example:
	//
	// 0.25
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s CreateMPULayoutRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s CreateMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequestPanes) GetHeight() *float32 {
	return s.Height
}

func (s *CreateMPULayoutRequestPanes) GetMajorPane() *int32 {
	return s.MajorPane
}

func (s *CreateMPULayoutRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *CreateMPULayoutRequestPanes) GetWidth() *float32 {
	return s.Width
}

func (s *CreateMPULayoutRequestPanes) GetX() *float32 {
	return s.X
}

func (s *CreateMPULayoutRequestPanes) GetY() *float32 {
	return s.Y
}

func (s *CreateMPULayoutRequestPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *CreateMPULayoutRequestPanes) SetHeight(v float32) *CreateMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetMajorPane(v int32) *CreateMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetPaneId(v int32) *CreateMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetWidth(v float32) *CreateMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetX(v float32) *CreateMPULayoutRequestPanes {
	s.X = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetY(v float32) *CreateMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetZOrder(v int32) *CreateMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) Validate() error {
	return dara.Validate(s)
}
