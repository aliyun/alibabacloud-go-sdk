// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLayoutListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayouts(v *DescribeSystemLayoutListResponseBodyLayouts) *DescribeSystemLayoutListResponseBody
	GetLayouts() *DescribeSystemLayoutListResponseBodyLayouts
	SetRequestId(v string) *DescribeSystemLayoutListResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeSystemLayoutListResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeSystemLayoutListResponseBody
	GetTotalPage() *int64
}

type DescribeSystemLayoutListResponseBody struct {
	Layouts *DescribeSystemLayoutListResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 2
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeSystemLayoutListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponseBody) GetLayouts() *DescribeSystemLayoutListResponseBodyLayouts {
	return s.Layouts
}

func (s *DescribeSystemLayoutListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSystemLayoutListResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeSystemLayoutListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeSystemLayoutListResponseBody) SetLayouts(v *DescribeSystemLayoutListResponseBodyLayouts) *DescribeSystemLayoutListResponseBody {
	s.Layouts = v
	return s
}

func (s *DescribeSystemLayoutListResponseBody) SetRequestId(v string) *DescribeSystemLayoutListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBody) SetTotalNum(v int64) *DescribeSystemLayoutListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBody) SetTotalPage(v int64) *DescribeSystemLayoutListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBody) Validate() error {
	if s.Layouts != nil {
		if err := s.Layouts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemLayoutListResponseBodyLayouts struct {
	Layout []*DescribeSystemLayoutListResponseBodyLayoutsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Repeated"`
}

func (s DescribeSystemLayoutListResponseBodyLayouts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponseBodyLayouts) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponseBodyLayouts) GetLayout() []*DescribeSystemLayoutListResponseBodyLayoutsLayout {
	return s.Layout
}

func (s *DescribeSystemLayoutListResponseBodyLayouts) SetLayout(v []*DescribeSystemLayoutListResponseBodyLayoutsLayout) *DescribeSystemLayoutListResponseBodyLayouts {
	s.Layout = v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayouts) Validate() error {
	if s.Layout != nil {
		for _, item := range s.Layout {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSystemLayoutListResponseBodyLayoutsLayout struct {
	// AudioMixCount。
	//
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// example:
	//
	// 22
	LayoutId *int64                                                  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Name     *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Panes    *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayout) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) GetAudioMixCount() *int32 {
	return s.AudioMixCount
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) GetLayoutId() *int64 {
	return s.LayoutId
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) GetName() *string {
	return s.Name
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) GetPanes() *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes {
	return s.Panes
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) SetAudioMixCount(v int32) *DescribeSystemLayoutListResponseBodyLayoutsLayout {
	s.AudioMixCount = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) SetLayoutId(v int64) *DescribeSystemLayoutListResponseBodyLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) SetName(v string) *DescribeSystemLayoutListResponseBodyLayoutsLayout {
	s.Name = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) SetPanes(v *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) *DescribeSystemLayoutListResponseBodyLayoutsLayout {
	s.Panes = v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayout) Validate() error {
	if s.Panes != nil {
		if err := s.Panes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes struct {
	Panes []*DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) GetPanes() []*DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	return s.Panes
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) SetPanes(v []*DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes {
	s.Panes = v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanes) Validate() error {
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

type DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes struct {
	// example:
	//
	// 0.25
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// MajorPane。
	//
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
	// 0.25
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.25
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.25
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetHeight() *float32 {
	return s.Height
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetMajorPane() *int32 {
	return s.MajorPane
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetWidth() *float32 {
	return s.Width
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetX() *float32 {
	return s.X
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetY() *float32 {
	return s.Y
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetHeight(v float32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.Height = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetMajorPane(v int32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.MajorPane = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetPaneId(v int32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetWidth(v float32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.Width = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetX(v float32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.X = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetY(v float32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.Y = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) SetZOrder(v int32) *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes {
	s.ZOrder = &v
	return s
}

func (s *DescribeSystemLayoutListResponseBodyLayoutsLayoutPanesPanes) Validate() error {
	return dara.Validate(s)
}
