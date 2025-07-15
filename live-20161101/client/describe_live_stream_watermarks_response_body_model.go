// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamWatermarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamWatermarksResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeLiveStreamWatermarksResponseBody
	GetTotal() *int32
	SetWatermarkList(v *DescribeLiveStreamWatermarksResponseBodyWatermarkList) *DescribeLiveStreamWatermarksResponseBody
	GetWatermarkList() *DescribeLiveStreamWatermarksResponseBodyWatermarkList
}

type DescribeLiveStreamWatermarksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of watermark templates that meet the specified conditions.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Details of the watermark templates.
	WatermarkList *DescribeLiveStreamWatermarksResponseBodyWatermarkList `json:"WatermarkList,omitempty" xml:"WatermarkList,omitempty" type:"Struct"`
}

func (s DescribeLiveStreamWatermarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamWatermarksResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveStreamWatermarksResponseBody) GetWatermarkList() *DescribeLiveStreamWatermarksResponseBodyWatermarkList {
	return s.WatermarkList
}

func (s *DescribeLiveStreamWatermarksResponseBody) SetRequestId(v string) *DescribeLiveStreamWatermarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBody) SetTotal(v int32) *DescribeLiveStreamWatermarksResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBody) SetWatermarkList(v *DescribeLiveStreamWatermarksResponseBodyWatermarkList) *DescribeLiveStreamWatermarksResponseBody {
	s.WatermarkList = v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamWatermarksResponseBodyWatermarkList struct {
	Watermark []*DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamWatermarksResponseBodyWatermarkList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarksResponseBodyWatermarkList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkList) GetWatermark() []*DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	return s.Watermark
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkList) SetWatermark(v []*DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) *DescribeLiveStreamWatermarksResponseBodyWatermarkList {
	s.Watermark = v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark struct {
	// The description of the watermark.
	//
	// example:
	//
	// my watermark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The height of the watermark. Unit: pixels.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the watermark.
	//
	// example:
	//
	// livewatermark****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The position of the watermark.
	//
	// 	- TopLeft: the upper-left corner.
	//
	// 	- TopRight: the upper-right corner.
	//
	// 	- BottomLeft: the lower-left corner.
	//
	// 	- BottomRight: the lower-right corner.
	//
	// example:
	//
	// TopRight
	OffsetCorner *string `json:"OffsetCorner,omitempty" xml:"OffsetCorner,omitempty"`
	// The URL of the watermark image.
	//
	// example:
	//
	// http://example.com
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The height of the background video. Unit: pixels.
	//
	// example:
	//
	// 1080
	RefHeight *int32 `json:"RefHeight,omitempty" xml:"RefHeight,omitempty"`
	// The width of the background video. Unit: pixels.
	//
	// example:
	//
	// 1920
	RefWidth *int32 `json:"RefWidth,omitempty" xml:"RefWidth,omitempty"`
	// The number of watermark rules configured for the domain name.
	//
	// example:
	//
	// 12
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 445409ec-7eaa-4 61d-8f29-4bec2eb9 ****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The transparency of the watermark. A smaller value indicates a more transparent watermark. Valid values: 0 to 255.
	//
	// example:
	//
	// 255
	Transparency *int32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	// The watermark type.
	//
	// 	- 0: image. Only image watermarks are supported.
	//
	// 	- 1: text.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The offset of the watermark along the x-axis. Unit: pixels.
	//
	// >  The value of the RefWidth parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the XOffset parameter indicates the x-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the x-axis is positive toward the right.
	//
	// example:
	//
	// 50.0
	XOffset *float32 `json:"XOffset,omitempty" xml:"XOffset,omitempty"`
	// The offset of the watermark along the y-axis. Unit: pixels.
	//
	// >  The value of the RefHeight parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the YOffset parameter indicates the y-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the y-axis is positive downward.
	//
	// example:
	//
	// 100.0
	YOffset *float32 `json:"YOffset,omitempty" xml:"YOffset,omitempty"`
}

func (s DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetDescription() *string {
	return s.Description
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetHeight() *int32 {
	return s.Height
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetName() *string {
	return s.Name
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetOffsetCorner() *string {
	return s.OffsetCorner
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetPictureUrl() *string {
	return s.PictureUrl
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetRefHeight() *int32 {
	return s.RefHeight
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetRefWidth() *int32 {
	return s.RefWidth
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetTransparency() *int32 {
	return s.Transparency
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetType() *int32 {
	return s.Type
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetXOffset() *float32 {
	return s.XOffset
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) GetYOffset() *float32 {
	return s.YOffset
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetDescription(v string) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.Description = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetHeight(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.Height = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetName(v string) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.Name = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetOffsetCorner(v string) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.OffsetCorner = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetPictureUrl(v string) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.PictureUrl = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetRefHeight(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.RefHeight = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetRefWidth(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.RefWidth = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetRuleCount(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.RuleCount = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetTemplateId(v string) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.TemplateId = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetTransparency(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.Transparency = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetType(v int32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.Type = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetXOffset(v float32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.XOffset = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) SetYOffset(v float32) *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark {
	s.YOffset = &v
	return s
}

func (s *DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark) Validate() error {
	return dara.Validate(s)
}
