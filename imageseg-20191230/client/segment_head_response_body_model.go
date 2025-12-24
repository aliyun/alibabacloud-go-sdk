// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHeadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentHeadResponseBodyData) *SegmentHeadResponseBody
	GetData() *SegmentHeadResponseBodyData
	SetRequestId(v string) *SegmentHeadResponseBody
	GetRequestId() *string
}

type SegmentHeadResponseBody struct {
	Data *SegmentHeadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 89BBDA42-E1CA-426E-9A46-C703D8DBB1E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHeadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBody) GetData() *SegmentHeadResponseBodyData {
	return s.Data
}

func (s *SegmentHeadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentHeadResponseBody) SetData(v *SegmentHeadResponseBodyData) *SegmentHeadResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHeadResponseBody) SetRequestId(v string) *SegmentHeadResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHeadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentHeadResponseBodyData struct {
	Elements []*SegmentHeadResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentHeadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBodyData) GetElements() []*SegmentHeadResponseBodyDataElements {
	return s.Elements
}

func (s *SegmentHeadResponseBodyData) SetElements(v []*SegmentHeadResponseBodyDataElements) *SegmentHeadResponseBodyData {
	s.Elements = v
	return s
}

func (s *SegmentHeadResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SegmentHeadResponseBodyDataElements struct {
	// example:
	//
	// 180
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_headsegmenter/2020-6-2/invi_headsegmenter_015910809094981099086_TAamhR.png?Expires=1591082709&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=xuUE%2FbYeB9LpR18VXawOVeutQU****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 116
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 445
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 102
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SegmentHeadResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponseBodyDataElements) GetHeight() *int32 {
	return s.Height
}

func (s *SegmentHeadResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHeadResponseBodyDataElements) GetWidth() *int32 {
	return s.Width
}

func (s *SegmentHeadResponseBodyDataElements) GetX() *int32 {
	return s.X
}

func (s *SegmentHeadResponseBodyDataElements) GetY() *int32 {
	return s.Y
}

func (s *SegmentHeadResponseBodyDataElements) SetHeight(v int32) *SegmentHeadResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetImageURL(v string) *SegmentHeadResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetWidth(v int32) *SegmentHeadResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetX(v int32) *SegmentHeadResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) SetY(v int32) *SegmentHeadResponseBodyDataElements {
	s.Y = &v
	return s
}

func (s *SegmentHeadResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
