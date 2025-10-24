// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWaterMarkTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNonExistWids(v *QueryWaterMarkTemplateListResponseBodyNonExistWids) *QueryWaterMarkTemplateListResponseBody
	GetNonExistWids() *QueryWaterMarkTemplateListResponseBodyNonExistWids
	SetRequestId(v string) *QueryWaterMarkTemplateListResponseBody
	GetRequestId() *string
	SetWaterMarkTemplateList(v *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) *QueryWaterMarkTemplateListResponseBody
	GetWaterMarkTemplateList() *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList
}

type QueryWaterMarkTemplateListResponseBody struct {
	// The IDs of the templates that do not exist.
	NonExistWids *QueryWaterMarkTemplateListResponseBodyNonExistWids `json:"NonExistWids,omitempty" xml:"NonExistWids,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 17079AF5-6276-51A9-B755-D26594C93F3C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the watermark templates.
	WaterMarkTemplateList *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList `json:"WaterMarkTemplateList,omitempty" xml:"WaterMarkTemplateList,omitempty" type:"Struct"`
}

func (s QueryWaterMarkTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBody) GetNonExistWids() *QueryWaterMarkTemplateListResponseBodyNonExistWids {
	return s.NonExistWids
}

func (s *QueryWaterMarkTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWaterMarkTemplateListResponseBody) GetWaterMarkTemplateList() *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList {
	return s.WaterMarkTemplateList
}

func (s *QueryWaterMarkTemplateListResponseBody) SetNonExistWids(v *QueryWaterMarkTemplateListResponseBodyNonExistWids) *QueryWaterMarkTemplateListResponseBody {
	s.NonExistWids = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBody) SetRequestId(v string) *QueryWaterMarkTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBody) SetWaterMarkTemplateList(v *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) *QueryWaterMarkTemplateListResponseBody {
	s.WaterMarkTemplateList = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBody) Validate() error {
	if s.NonExistWids != nil {
		if err := s.NonExistWids.Validate(); err != nil {
			return err
		}
	}
	if s.WaterMarkTemplateList != nil {
		if err := s.WaterMarkTemplateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWaterMarkTemplateListResponseBodyNonExistWids struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s QueryWaterMarkTemplateListResponseBodyNonExistWids) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBodyNonExistWids) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBodyNonExistWids) GetString_() []*string {
	return s.String_
}

func (s *QueryWaterMarkTemplateListResponseBodyNonExistWids) SetString_(v []*string) *QueryWaterMarkTemplateListResponseBodyNonExistWids {
	s.String_ = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyNonExistWids) Validate() error {
	return dara.Validate(s)
}

type QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList struct {
	WaterMarkTemplate []*QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate `json:"WaterMarkTemplate,omitempty" xml:"WaterMarkTemplate,omitempty" type:"Repeated"`
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) GetWaterMarkTemplate() []*QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	return s.WaterMarkTemplate
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) SetWaterMarkTemplate(v []*QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList {
	s.WaterMarkTemplate = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateList) Validate() error {
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

type QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate struct {
	// The horizontal offset. Unit: pixel.
	//
	// example:
	//
	// 100
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset. Unit: pixel.
	//
	// example:
	//
	// 100
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 8
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the watermark template.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the watermark template.
	//
	// example:
	//
	// example-watermark
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the Height, Width, Dx, and Dy parameters relative to the reference edges. If the values of the Height, Width, Dx, and Dy parameters are decimals between 0 and 1, the values are calculated by referring to the following edges in sequence:
	//
	// 	- **Width**: the width edge.
	//
	// 	- **Height**: the height edge.
	//
	// 	- **Long**: the long edge.
	//
	// 	- **Short**: the short edge.
	RatioRefer *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer `json:"RatioRefer,omitempty" xml:"RatioRefer,omitempty" type:"Struct"`
	// The position of the watermark. Valid values:
	//
	// 	- **TopRight**: the upper-right corner.
	//
	// 	- **TopLeft**: the upper-left corner.
	//
	// 	- **BottomRight**: the lower-right corner.
	//
	// 	- **BottomLeft**: the lower-left corner.
	//
	// example:
	//
	// TopRight
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The status of the watermark template. Valid values: Valid values:
	//
	// 	- **Normal**: The watermark template is normal.
	//
	// 	- **Deleted**: The watermark template is deleted.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timeline of the watermark.
	Timeline *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The type of the watermark. Valid values:
	//
	// 	- Image: an image watermark.
	//
	// 	- Text: a text watermark.
	//
	// > Only watermarks of the **Image*	- type are supported.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The width of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 8
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetDx() *string {
	return s.Dx
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetDy() *string {
	return s.Dy
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetHeight() *string {
	return s.Height
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetId() *string {
	return s.Id
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetName() *string {
	return s.Name
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetRatioRefer() *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	return s.RatioRefer
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetReferPos() *string {
	return s.ReferPos
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetState() *string {
	return s.State
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetTimeline() *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	return s.Timeline
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetType() *string {
	return s.Type
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) GetWidth() *string {
	return s.Width
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetDx(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Dx = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetDy(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Dy = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetHeight(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Height = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetId(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Id = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetName(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Name = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetRatioRefer(v *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.RatioRefer = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetReferPos(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.ReferPos = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetState(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.State = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetTimeline(v *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Timeline = v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetType(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Type = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) SetWidth(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate {
	s.Width = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplate) Validate() error {
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

type QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer struct {
	// The horizontal offset of the watermark relative to the output video image. Default value: **0**. The default value indicates no offset. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the vertical offset. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the horizontal offset to the width of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// 0.51
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset of the watermark relative to the output video image. Default value: **0**. The default value indicates no offset. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the vertical offset. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the vertical offset to the height of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// 0.4
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the watermark image in the output video. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the height of the watermark image. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the height of the watermark image to the height of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// 0.33
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The width of the watermark image in the output video. The value can be an integer or a decimal.
	//
	// 	- **Integer**: the width of the watermark image. This indicates the absolute position. Unit: pixel.
	//
	// 	- **Decimal**: the ratio of the width of the watermark image to the width of the output video. The ratio varies based on the size of the video. Four decimal places are supported, such as 0.9999. More decimal places are discarded.
	//
	// example:
	//
	// 0.36
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetDx() *string {
	return s.Dx
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetDy() *string {
	return s.Dy
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetHeight() *string {
	return s.Height
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) GetWidth() *string {
	return s.Width
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetDx(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Dx = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetDy(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Dy = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetHeight(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Height = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) SetWidth(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer {
	s.Width = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer) Validate() error {
	return dara.Validate(s)
}

type QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline struct {
	// The display duration of the watermark. Default value: **ToEND**. The default value indicates that the watermark is displayed until the video ends.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The beginning of the time range during which the watermark is displayed.
	//
	// 	- Unit: seconds.
	//
	// 	- Default value: **0**.
	//
	// example:
	//
	// 0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) String() string {
	return dara.Prettify(s)
}

func (s QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GoString() string {
	return s.String()
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GetDuration() *string {
	return s.Duration
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) GetStart() *string {
	return s.Start
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) SetDuration(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	s.Duration = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) SetStart(v string) *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline {
	s.Start = &v
	return s
}

func (s *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline) Validate() error {
	return dara.Validate(s)
}
