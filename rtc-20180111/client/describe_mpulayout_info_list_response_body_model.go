// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMPULayoutInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayouts(v *DescribeMPULayoutInfoListResponseBodyLayouts) *DescribeMPULayoutInfoListResponseBody
	GetLayouts() *DescribeMPULayoutInfoListResponseBodyLayouts
	SetRequestId(v string) *DescribeMPULayoutInfoListResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeMPULayoutInfoListResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeMPULayoutInfoListResponseBody
	GetTotalPage() *int64
}

type DescribeMPULayoutInfoListResponseBody struct {
	Layouts *DescribeMPULayoutInfoListResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBody) GetLayouts() *DescribeMPULayoutInfoListResponseBodyLayouts {
	return s.Layouts
}

func (s *DescribeMPULayoutInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMPULayoutInfoListResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeMPULayoutInfoListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeMPULayoutInfoListResponseBody) SetLayouts(v *DescribeMPULayoutInfoListResponseBodyLayouts) *DescribeMPULayoutInfoListResponseBody {
	s.Layouts = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetRequestId(v string) *DescribeMPULayoutInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalNum(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalPage(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMPULayoutInfoListResponseBodyLayouts struct {
	Layout []*DescribeMPULayoutInfoListResponseBodyLayoutsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayouts) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayouts) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayouts) GetLayout() []*DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	return s.Layout
}

func (s *DescribeMPULayoutInfoListResponseBodyLayouts) SetLayout(v []*DescribeMPULayoutInfoListResponseBodyLayoutsLayout) *DescribeMPULayoutInfoListResponseBodyLayouts {
	s.Layout = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayouts) Validate() error {
	return dara.Validate(s)
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayout struct {
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name  *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Panes *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GetAudioMixCount() *int32 {
	return s.AudioMixCount
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GetName() *string {
	return s.Name
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GetPanes() *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes {
	return s.Panes
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetAudioMixCount(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.AudioMixCount = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetLayoutId(v int64) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetName(v string) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetPanes(v *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Panes = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) Validate() error {
	return dara.Validate(s)
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes struct {
	Panes []*DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) GetPanes() []*DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	return s.Panes
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) SetPanes(v []*DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes {
	s.Panes = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) Validate() error {
	return dara.Validate(s)
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes struct {
	// example:
	//
	// 0.5
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
	// 0.5
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.5
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.5
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) String() string {
	return dara.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetHeight() *float32 {
	return s.Height
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetMajorPane() *int32 {
	return s.MajorPane
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetWidth() *float32 {
	return s.Width
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetX() *float32 {
	return s.X
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetY() *float32 {
	return s.Y
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetHeight(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Height = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetMajorPane(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.MajorPane = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetPaneId(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetWidth(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Width = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetX(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.X = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetY(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Y = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetZOrder(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.ZOrder = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) Validate() error {
	return dara.Validate(s)
}
