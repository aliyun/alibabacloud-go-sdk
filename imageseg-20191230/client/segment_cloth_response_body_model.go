// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentClothResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentClothResponseBodyData) *SegmentClothResponseBody
	GetData() *SegmentClothResponseBodyData
	SetRequestId(v string) *SegmentClothResponseBody
	GetRequestId() *string
}

type SegmentClothResponseBody struct {
	Data *SegmentClothResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BCE049A3-FE69-41CF-A870-5970156388A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentClothResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBody) GetData() *SegmentClothResponseBodyData {
	return s.Data
}

func (s *SegmentClothResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentClothResponseBody) SetData(v *SegmentClothResponseBodyData) *SegmentClothResponseBody {
	s.Data = v
	return s
}

func (s *SegmentClothResponseBody) SetRequestId(v string) *SegmentClothResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentClothResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentClothResponseBodyData struct {
	Elements []*SegmentClothResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s SegmentClothResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBodyData) GetElements() []*SegmentClothResponseBodyDataElements {
	return s.Elements
}

func (s *SegmentClothResponseBodyData) SetElements(v []*SegmentClothResponseBodyDataElements) *SegmentClothResponseBodyData {
	s.Elements = v
	return s
}

func (s *SegmentClothResponseBodyData) Validate() error {
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

type SegmentClothResponseBodyDataElements struct {
	ClassUrl map[string]*string `json:"ClassUrl,omitempty" xml:"ClassUrl,omitempty"`
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/clothingsegmentation-2020-06-17-16-54-40-688c84cbbd-hnqfq/2020-6-18/invi__015924459307821000041_IIVHoM.png?Expires=1592447730&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Hy8pn3IQj8nuKN0LEaC57cee9L****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentClothResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *SegmentClothResponseBodyDataElements) GetClassUrl() map[string]*string {
	return s.ClassUrl
}

func (s *SegmentClothResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentClothResponseBodyDataElements) SetClassUrl(v map[string]*string) *SegmentClothResponseBodyDataElements {
	s.ClassUrl = v
	return s
}

func (s *SegmentClothResponseBodyDataElements) SetImageURL(v string) *SegmentClothResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *SegmentClothResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
