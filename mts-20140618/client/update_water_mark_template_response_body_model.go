// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaterMarkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWaterMarkTemplateResponseBody
	GetRequestId() *string
	SetWaterMarkTemplate(v *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) *UpdateWaterMarkTemplateResponseBody
	GetWaterMarkTemplate() *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate
}

type UpdateWaterMarkTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E558894E-40D9-57C6-B5CC-0F5CDF23614E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the watermark template.
	WaterMarkTemplate *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate `json:"WaterMarkTemplate,omitempty" xml:"WaterMarkTemplate,omitempty" type:"Struct"`
}

func (s UpdateWaterMarkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWaterMarkTemplateResponseBody) GetWaterMarkTemplate() *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	return s.WaterMarkTemplate
}

func (s *UpdateWaterMarkTemplateResponseBody) SetRequestId(v string) *UpdateWaterMarkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBody) SetWaterMarkTemplate(v *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) *UpdateWaterMarkTemplateResponseBody {
	s.WaterMarkTemplate = v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBody) Validate() error {
	if s.WaterMarkTemplate != nil {
		if err := s.WaterMarkTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate struct {
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
	RatioRefer *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer `json:"RatioRefer,omitempty" xml:"RatioRefer,omitempty" type:"Struct"`
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
	// The status of the watermark template. Default value: **Normal**.
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timeline of the watermark.
	Timeline *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	// The type of the watermark. Valid values:
	//
	// 	- Image: an image watermark.
	//
	// 	- Text: a text watermark.
	//
	// > Only watermarks of the Image type are supported.
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

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetDx() *string {
	return s.Dx
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetDy() *string {
	return s.Dy
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetHeight() *string {
	return s.Height
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetId() *string {
	return s.Id
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetName() *string {
	return s.Name
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetRatioRefer() *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	return s.RatioRefer
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetReferPos() *string {
	return s.ReferPos
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetState() *string {
	return s.State
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetTimeline() *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	return s.Timeline
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetType() *string {
	return s.Type
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) GetWidth() *string {
	return s.Width
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetDx(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Dx = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetDy(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Dy = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetHeight(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Height = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetId(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Id = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetName(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Name = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetRatioRefer(v *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.RatioRefer = v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetReferPos(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.ReferPos = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetState(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.State = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetTimeline(v *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Timeline = v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetType(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Type = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) SetWidth(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate {
	s.Width = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplate) Validate() error {
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

type UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer struct {
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

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetDx() *string {
	return s.Dx
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetDy() *string {
	return s.Dy
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetHeight() *string {
	return s.Height
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) GetWidth() *string {
	return s.Width
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetDx(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Dx = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetDy(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Dy = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetHeight(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Height = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) SetWidth(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer {
	s.Width = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateRatioRefer) Validate() error {
	return dara.Validate(s)
}

type UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline struct {
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

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GoString() string {
	return s.String()
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GetDuration() *string {
	return s.Duration
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) GetStart() *string {
	return s.Start
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) SetDuration(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	s.Duration = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) SetStart(v string) *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline {
	s.Start = &v
	return s
}

func (s *UpdateWaterMarkTemplateResponseBodyWaterMarkTemplateTimeline) Validate() error {
	return dara.Validate(s)
}
