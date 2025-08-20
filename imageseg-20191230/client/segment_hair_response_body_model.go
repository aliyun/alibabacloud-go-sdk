// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentHairResponseBodyData) *SegmentHairResponseBody
	GetData() *SegmentHairResponseBodyData
	SetRequestId(v string) *SegmentHairResponseBody
	GetRequestId() *string
}

type SegmentHairResponseBody struct {
	Data *SegmentHairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentHairResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBody) GetData() *SegmentHairResponseBodyData {
	return s.Data
}

func (s *SegmentHairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentHairResponseBody) SetData(v *SegmentHairResponseBodyData) *SegmentHairResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHairResponseBody) SetRequestId(v string) *SegmentHairResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHairResponseBody) Validate() error {
	return dara.Validate(s)
}

type SegmentHairResponseBodyData struct {
	Elements []*SegmentHairResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentHairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentHairResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBodyData) GetElements() []*SegmentHairResponseBodyDataElements {
	return s.Elements
}

func (s *SegmentHairResponseBodyData) SetElements(v []*SegmentHairResponseBodyDataElements) *SegmentHairResponseBodyData {
	s.Elements = v
	return s
}

func (s *SegmentHairResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SegmentHairResponseBodyDataElements struct {
	// example:
	//
	// 180
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_HeadSegmenter/2021-12-31/invi_HeadSegmenter_016409228383064285967296_iPLUXA.png?Expires=1640924638&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=wpKOqSar1bYvGmlTMryfEH2Q9I****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 113
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 446
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 102
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SegmentHairResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s SegmentHairResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHairResponseBodyDataElements) GetHeight() *int32 {
	return s.Height
}

func (s *SegmentHairResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHairResponseBodyDataElements) GetWidth() *int32 {
	return s.Width
}

func (s *SegmentHairResponseBodyDataElements) GetX() *int32 {
	return s.X
}

func (s *SegmentHairResponseBodyDataElements) GetY() *int32 {
	return s.Y
}

func (s *SegmentHairResponseBodyDataElements) SetHeight(v int32) *SegmentHairResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetImageURL(v string) *SegmentHairResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetWidth(v int32) *SegmentHairResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetX(v int32) *SegmentHairResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) SetY(v int32) *SegmentHairResponseBodyDataElements {
	s.Y = &v
	return s
}

func (s *SegmentHairResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
