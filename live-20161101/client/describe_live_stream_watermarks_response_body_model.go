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
	Total         *int32                                                 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	if s.WatermarkList != nil {
		if err := s.WatermarkList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Watermark != nil {
		for _, item := range s.Watermark {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamWatermarksResponseBodyWatermarkListWatermark struct {
	Description  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Height       *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Name         *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	OffsetCorner *string  `json:"OffsetCorner,omitempty" xml:"OffsetCorner,omitempty"`
	PictureUrl   *string  `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	RefHeight    *int32   `json:"RefHeight,omitempty" xml:"RefHeight,omitempty"`
	RefWidth     *int32   `json:"RefWidth,omitempty" xml:"RefWidth,omitempty"`
	RuleCount    *int32   `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	TemplateId   *string  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Transparency *int32   `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	Type         *int32   `json:"Type,omitempty" xml:"Type,omitempty"`
	XOffset      *float32 `json:"XOffset,omitempty" xml:"XOffset,omitempty"`
	YOffset      *float32 `json:"YOffset,omitempty" xml:"YOffset,omitempty"`
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
