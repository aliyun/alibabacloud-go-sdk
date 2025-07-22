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
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10117
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name    *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Panes   []*ModifyMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type ModifyMPULayoutRequestPanes struct {
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
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
