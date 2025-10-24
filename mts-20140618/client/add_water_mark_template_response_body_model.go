// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWaterMarkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWaterMarkTemplateResponseBody
	GetRequestId() *string
	SetWaterMarkTemplate(v *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) *AddWaterMarkTemplateResponseBody
	GetWaterMarkTemplate() *AddWaterMarkTemplateResponseBodyWaterMarkTemplate
}

type AddWaterMarkTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 54BB917F-DD35-4F32-BABA-E60E31B21W63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the watermark template.
	WaterMarkTemplate *AddWaterMarkTemplateResponseBodyWaterMarkTemplate `json:"WaterMarkTemplate,omitempty" xml:"WaterMarkTemplate,omitempty" type:"Struct"`
}

func (s AddWaterMarkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWaterMarkTemplateResponseBody) GetWaterMarkTemplate() *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	return s.WaterMarkTemplate
}

func (s *AddWaterMarkTemplateResponseBody) SetRequestId(v string) *AddWaterMarkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBody) SetWaterMarkTemplate(v *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) *AddWaterMarkTemplateResponseBody {
	s.WaterMarkTemplate = v
	return s
}

func (s *AddWaterMarkTemplateResponseBody) Validate() error {
	if s.WaterMarkTemplate != nil {
		if err := s.WaterMarkTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWaterMarkTemplateResponseBodyWaterMarkTemplate struct {
	// The horizontal offset. Unit: pixel.
	//
	// example:
	//
	// 10
	Dx *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The vertical offset. Unit: pixel.
	//
	// example:
	//
	// 5
	Dy *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 30
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The ID of the watermark template. We recommend that you keep this ID for subsequent operation calls.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the watermark template.
	//
	// example:
	//
	// example-watermark-****
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
	RatioRefer *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer `json:"RatioRefer,omitempty" xml:"RatioRefer,omitempty" type:"Struct"`
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
	// The status of the watermark template.
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
	Timeline *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The type of the watermark. Valid values:
	//
	// 	- Image: an image watermark.
	//
	// 	- Text: a text watermark.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The width of the watermark image. Unit: pixel.
	//
	// example:
	//
	// 10
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplate) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetDx() *string {
	return s.Dx
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetDy() *string {
	return s.Dy
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetHeight() *string {
	return s.Height
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetId() *string {
	return s.Id
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetName() *string {
	return s.Name
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetRatioRefer() *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	return s.RatioRefer
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetReferPos() *string {
	return s.ReferPos
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetState() *string {
	return s.State
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetTimeline() *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	return s.Timeline
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetType() *string {
	return s.Type
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) GetWidth() *string {
	return s.Width
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetDx(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Dx = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetDy(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Dy = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetHeight(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Height = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetId(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Id = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetName(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Name = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetRatioRefer(v *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.RatioRefer = v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetReferPos(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.ReferPos = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetState(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.State = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetTimeline(v *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Timeline = v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetType(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Type = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) SetWidth(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Width = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplate) Validate() error {
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

type AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer struct {
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
	// 0.28
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

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetDx() *string {
	return s.Dx
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetDy() *string {
	return s.Dy
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetHeight() *string {
	return s.Height
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetWidth() *string {
	return s.Width
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetDx(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Dx = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetDy(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Dy = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetHeight(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Height = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetWidth(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Width = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) Validate() error {
	return dara.Validate(s)
}

type AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline struct {
	// The display duration of the watermark. Default value: **ToEND**. The default value indicates that the watermark is displayed until the video ends.
	//
	// example:
	//
	// ToEND
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

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) String() string {
	return dara.Prettify(s)
}

func (s AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GoString() string {
	return s.String()
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GetDuration() *string {
	return s.Duration
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GetStart() *string {
	return s.Start
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) SetDuration(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	s.Duration = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) SetStart(v string) *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	s.Start = &v
	return s
}

func (s *AddWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) Validate() error {
	return dara.Validate(s)
}
