// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMPULayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyMPULayoutRequest
	GetAppId() *string
	SetAudioMixCount(v int32) *ModifyMPULayoutRequest
	GetAudioMixCount() *int32
	SetLayoutId(v int64) *ModifyMPULayoutRequest
	GetLayoutId() *int64
	SetName(v string) *ModifyMPULayoutRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyMPULayoutRequest
	GetOwnerId() *int64
	SetPanes(v []*ModifyMPULayoutRequestPanes) *ModifyMPULayoutRequest
	GetPanes() []*ModifyMPULayoutRequestPanes
}

type ModifyMPULayoutRequest struct {
	// The ID of the application. Only a single ID is supported. You can create and query application IDs in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The maximum number of audio mixing tracks. Valid values: **0*	- to **6**.
	//
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// The ID of the layout. For more information about how to obtain the layout ID, see [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 10117
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The name of the layout.
	//
	// example:
	//
	// LayoutName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of panes in the layout.
	Panes []*ModifyMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s ModifyMPULayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyMPULayoutRequest) GetAudioMixCount() *int32 {
	return s.AudioMixCount
}

func (s *ModifyMPULayoutRequest) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *ModifyMPULayoutRequest) GetName() *string {
	return s.Name
}

func (s *ModifyMPULayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMPULayoutRequest) GetPanes() []*ModifyMPULayoutRequestPanes {
	return s.Panes
}

func (s *ModifyMPULayoutRequest) SetAppId(v string) *ModifyMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetAudioMixCount(v int32) *ModifyMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetLayoutId(v int64) *ModifyMPULayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetName(v string) *ModifyMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetOwnerId(v int64) *ModifyMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetPanes(v []*ModifyMPULayoutRequestPanes) *ModifyMPULayoutRequest {
	s.Panes = v
	return s
}

func (s *ModifyMPULayoutRequest) Validate() error {
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

type ModifyMPULayoutRequestPanes struct {
	// The height of the pane, represented as a normalized percentage. Valid values: **0.0 < Height ≤ 1.0**.
	//
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies the major pane. Each layout can have only one major pane. Valid values:
	//
	// - **0**: Secondary pane.
	//
	// - **1**: Major pane.
	//
	// example:
	//
	// 0
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// The ID of the pane. Panes are sorted from top left to bottom right, starting from 0.
	//
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// The width of the pane, represented as a normalized percentage. Valid values: **0.0 < Width ≤ 1.0**.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, represented as a normalized percentage. Valid values: **0.0 ≤ X ≤ 1.0**.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, represented as a normalized percentage. Valid values: **0.0 ≤ Y ≤ 1.0**.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order. A value of 0 indicates the bottom layer, 1 indicates the layer above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s ModifyMPULayoutRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s ModifyMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequestPanes) GetHeight() *float32 {
	return s.Height
}

func (s *ModifyMPULayoutRequestPanes) GetMajorPane() *int32 {
	return s.MajorPane
}

func (s *ModifyMPULayoutRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *ModifyMPULayoutRequestPanes) GetWidth() *float32 {
	return s.Width
}

func (s *ModifyMPULayoutRequestPanes) GetX() *float32 {
	return s.X
}

func (s *ModifyMPULayoutRequestPanes) GetY() *float32 {
	return s.Y
}

func (s *ModifyMPULayoutRequestPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *ModifyMPULayoutRequestPanes) SetHeight(v float32) *ModifyMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetMajorPane(v int32) *ModifyMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetPaneId(v int32) *ModifyMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetWidth(v float32) *ModifyMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetX(v float32) *ModifyMPULayoutRequestPanes {
	s.X = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetY(v float32) *ModifyMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetZOrder(v int32) *ModifyMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) Validate() error {
	return dara.Validate(s)
}
