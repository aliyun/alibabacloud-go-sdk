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
	NonExistWids *QueryWaterMarkTemplateListResponseBodyNonExistWids `json:"NonExistWids,omitempty" xml:"NonExistWids,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 17079AF5-6276-51A9-B755-D26594C93F3C
	RequestId             *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Dx         *string                                                                                 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy         *string                                                                                 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height     *string                                                                                 `json:"Height,omitempty" xml:"Height,omitempty"`
	Id         *string                                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RatioRefer *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateRatioRefer `json:"RatioRefer,omitempty" xml:"RatioRefer,omitempty" type:"Struct"`
	ReferPos   *string                                                                                 `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	State      *string                                                                                 `json:"State,omitempty" xml:"State,omitempty"`
	Timeline   *QueryWaterMarkTemplateListResponseBodyWaterMarkTemplateListWaterMarkTemplateTimeline   `json:"Timeline,omitempty" xml:"Timeline,omitempty" type:"Struct"`
	Type       *string                                                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Width      *string                                                                                 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Dx     *string `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy     *string `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *string `json:"Width,omitempty" xml:"Width,omitempty"`
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
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
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
