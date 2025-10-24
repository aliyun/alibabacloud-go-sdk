// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWaterMarkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *SearchWaterMarkTemplateResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *SearchWaterMarkTemplateResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *SearchWaterMarkTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchWaterMarkTemplateResponseBody
	GetTotalCount() *int64
	SetWaterMarkTemplateList(v *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) *SearchWaterMarkTemplateResponseBody
	GetWaterMarkTemplateList() *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList
}

type SearchWaterMarkTemplateResponseBody struct {
	// The width of the watermark image in the output video. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the width of the watermark image. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the width of the watermark image to the width of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The values of the Height, Width, Dx, and Dy parameters relative to the reference edges. If the values of the Height, Width, Dx, and Dy parameters are decimals between 0 and 1, the values are calculated by referring to the following edges in sequence:
	//
	// 	- **Width**: the width edge.
	//
	// 	- **Height**: the height edge.
	//
	// 	- **Long**: the long edge.
	//
	// 	- **Short**: the short edge.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FC029D04-8F47-57FF-A759-23383C15617D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the watermark. Valid values:
	//
	// 	- Image: an image watermark.
	//
	// 	- Text: a text watermark.
	//
	// >  Only watermarks of the **Image*	- types are supported.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The height of the watermark image in the output video. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the height of the watermark image. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the height of the watermark image to the height of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	WaterMarkTemplateList *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList `json:"WaterMarkTemplateList,omitempty" xml:"WaterMarkTemplateList,omitempty" type:"Struct"`
}

func (s SearchWaterMarkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchWaterMarkTemplateResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchWaterMarkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchWaterMarkTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchWaterMarkTemplateResponseBody) GetWaterMarkTemplateList() *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList {
	return s.WaterMarkTemplateList
}

func (s *SearchWaterMarkTemplateResponseBody) SetPageNumber(v int64) *SearchWaterMarkTemplateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBody) SetPageSize(v int64) *SearchWaterMarkTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBody) SetRequestId(v string) *SearchWaterMarkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBody) SetTotalCount(v int64) *SearchWaterMarkTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBody) SetWaterMarkTemplateList(v *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) *SearchWaterMarkTemplateResponseBody {
	s.WaterMarkTemplateList = v
	return s
}

func (s *SearchWaterMarkTemplateResponseBody) Validate() error {
	if s.WaterMarkTemplateList != nil {
		if err := s.WaterMarkTemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList struct {
	WaterMarkTemplate []*SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate `json:"WaterMarkTemplate,omitempty" xml:"WaterMarkTemplate,omitempty" type:"Repeated"`
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) GetWaterMarkTemplate() []*SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	return s.WaterMarkTemplate
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) SetWaterMarkTemplate(v []*SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList {
	s.WaterMarkTemplate = v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateList) Validate() error {
	if s.WaterMarkTemplate != nil {
		for _, item := range s.WaterMarkTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate struct {
	// The name of the watermark template.
	//
	// example:
	//
	// 100
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The values of the Height, Width, Dx, and Dy parameters relative to the reference edges. If the values of the Height, Width, Dx, and Dy parameters are decimals between 0 and 1, the values are calculated by referring to the following edges in sequence:
	//
	// 	- **Width**: the width edge.
	//
	// 	- **Height**: the height edge.
	//
	// 	- **Long**: the long edge.
	//
	// 	- **Short**: the short edge.
	//
	// example:
	//
	// 100
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The vertical offset. Unit: pixel.
	//
	// example:
	//
	// 88c6ca184c0e4578645b665e2a12****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The width of the watermark image in the output video. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the width of the watermark image. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the width of the watermark image to the width of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// example-watermark
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the watermark template. Valid values: Valid values:
	//
	// 	- **Normal**: The watermark template is normal.
	//
	// 	- **Deleted**: The watermark template is deleted.
	RatioRefer *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer `json:"RatioRefer,omitempty" xml:"RatioRefer,omitempty" type:"Struct"`
	// The beginning of the time range during which the watermark is displayed.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **0**.
	//
	// example:
	//
	// TopRight
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The display duration of the watermark. Default value: **ToEND**. The default value indicates that the watermark is displayed until the video ends.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timeline of the watermark.
	Timeline *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The position of the watermark. Valid values:
	//
	// 	- TopRight: the upper-right corner.
	//
	// 	- TopLeft: the upper-left corner.
	//
	// 	- BottomRight: the lower-right corner.
	//
	// 	- BottomLeft: the lower-left corner.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vertical offset. Unit: pixel.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetDx() *string {
	return s.Dx
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetDy() *string {
	return s.Dy
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetHeight() *string {
	return s.Height
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetId() *string {
	return s.Id
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetName() *string {
	return s.Name
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetRatioRefer() *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	return s.RatioRefer
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetReferPos() *string {
	return s.ReferPos
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetState() *string {
	return s.State
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetTimeline() *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	return s.Timeline
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetType() *string {
	return s.Type
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetWidth() *string {
	return s.Width
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetDx(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Dx = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetDy(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Dy = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetHeight(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Height = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetId(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Id = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetName(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Name = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetRatioRefer(v *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.RatioRefer = v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetReferPos(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.ReferPos = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetState(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.State = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetTimeline(v *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Timeline = v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetType(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Type = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetWidth(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Width = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplate) Validate() error {
	if s.RatioRefer != nil {
		if err := s.RatioRefer.Validate(); err != nil {
			return err
		}
	}
	if s.Timeline != nil {
		if err := s.Timeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer struct {
	// The horizontal offset. Unit: pixel.
	//
	// example:
	//
	// 0.51
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The timeline of the watermark.
	//
	// example:
	//
	// 0.2
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 0.33
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The width of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 0.36
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetDx() *string {
	return s.Dx
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetDy() *string {
	return s.Dy
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetHeight() *string {
	return s.Height
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetWidth() *string {
	return s.Width
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetDx(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Dx = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetDy(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Dy = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetHeight(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Height = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetWidth(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Width = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) Validate() error {
	return dara.Validate(s)
}

type SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline struct {
	// The horizontal offset of the watermark relative to the output video image. Default value: **0**. The default value indicates no offset.
	//
	// The value can be an integer or a decimal.
	//
	// 	- **Integer**: the vertical offset. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the horizontal offset to the width of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// ToEND
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) String() string {
	return dara.Prettify(s)
}

func (s SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GoString() string {
	return s.String()
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GetDuration() *string {
	return s.Duration
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GetStart() *string {
	return s.Start
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) SetDuration(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	s.Duration = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) SetStart(v string) *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	s.Start = &v
	return s
}

func (s *SearchWaterMarkTemplateResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) Validate() error {
	return dara.Validate(s)
}
