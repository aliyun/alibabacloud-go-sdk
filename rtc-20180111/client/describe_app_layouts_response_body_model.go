// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLayoutsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayouts(v []*DescribeAppLayoutsResponseBodyLayouts) *DescribeAppLayoutsResponseBody
	GetLayouts() []*DescribeAppLayoutsResponseBodyLayouts
	SetRequestId(v string) *DescribeAppLayoutsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeAppLayoutsResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppLayoutsResponseBody
	GetTotalPage() *int64
}

type DescribeAppLayoutsResponseBody struct {
	Layouts []*DescribeAppLayoutsResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Repeated"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppLayoutsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsResponseBody) GetLayouts() []*DescribeAppLayoutsResponseBodyLayouts {
	return s.Layouts
}

func (s *DescribeAppLayoutsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppLayoutsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppLayoutsResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppLayoutsResponseBody) SetLayouts(v []*DescribeAppLayoutsResponseBodyLayouts) *DescribeAppLayoutsResponseBody {
	s.Layouts = v
	return s
}

func (s *DescribeAppLayoutsResponseBody) SetRequestId(v string) *DescribeAppLayoutsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppLayoutsResponseBody) SetTotalNum(v int64) *DescribeAppLayoutsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppLayoutsResponseBody) SetTotalPage(v int64) *DescribeAppLayoutsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppLayoutsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppLayoutsResponseBodyLayouts struct {
	// example:
	//
	// 167466539798442****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// 测试
	Name  *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Panes []*DescribeAppLayoutsResponseBodyLayoutsPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s DescribeAppLayoutsResponseBodyLayouts) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsResponseBodyLayouts) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsResponseBodyLayouts) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeAppLayoutsResponseBodyLayouts) GetName() *string {
	return s.Name
}

func (s *DescribeAppLayoutsResponseBodyLayouts) GetPanes() []*DescribeAppLayoutsResponseBodyLayoutsPanes {
	return s.Panes
}

func (s *DescribeAppLayoutsResponseBodyLayouts) SetLayoutId(v string) *DescribeAppLayoutsResponseBodyLayouts {
	s.LayoutId = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayouts) SetName(v string) *DescribeAppLayoutsResponseBodyLayouts {
	s.Name = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayouts) SetPanes(v []*DescribeAppLayoutsResponseBodyLayoutsPanes) *DescribeAppLayoutsResponseBodyLayouts {
	s.Panes = v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayouts) Validate() error {
	return dara.Validate(s)
}

type DescribeAppLayoutsResponseBodyLayoutsPanes struct {
	// example:
	//
	// 0.25
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
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

func (s DescribeAppLayoutsResponseBodyLayoutsPanes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLayoutsResponseBodyLayoutsPanes) GoString() string {
	return s.String()
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetHeight() *float32 {
	return s.Height
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetWidth() *float32 {
	return s.Width
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetX() *float32 {
	return s.X
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetY() *float32 {
	return s.Y
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetHeight(v float32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.Height = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetPaneId(v int32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetWidth(v float32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.Width = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetX(v float32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.X = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetY(v float32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.Y = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) SetZOrder(v int32) *DescribeAppLayoutsResponseBodyLayoutsPanes {
	s.ZOrder = &v
	return s
}

func (s *DescribeAppLayoutsResponseBodyLayoutsPanes) Validate() error {
	return dara.Validate(s)
}
