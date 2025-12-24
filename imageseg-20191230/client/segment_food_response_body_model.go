// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentFoodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentFoodResponseBodyData) *SegmentFoodResponseBody
	GetData() *SegmentFoodResponseBodyData
	SetRequestId(v string) *SegmentFoodResponseBody
	GetRequestId() *string
}

type SegmentFoodResponseBody struct {
	Data *SegmentFoodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 38265D08-AD0F-4752-8E96-D1D9FB96C3D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentFoodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentFoodResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponseBody) GetData() *SegmentFoodResponseBodyData {
	return s.Data
}

func (s *SegmentFoodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentFoodResponseBody) SetData(v *SegmentFoodResponseBodyData) *SegmentFoodResponseBody {
	s.Data = v
	return s
}

func (s *SegmentFoodResponseBody) SetRequestId(v string) *SegmentFoodResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentFoodResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentFoodResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/foodsegmenter-2020-06-17-15-24-00-8658fc85b8-8ds8k/2020-6-18/invi__015924442076191000002_WqJ99N.png?Expires=1592446007&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=5ITSd6ndSuP7pUfoDFpgLPUOGg****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentFoodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentFoodResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentFoodResponseBodyData) SetImageURL(v string) *SegmentFoodResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentFoodResponseBodyData) Validate() error {
	return dara.Validate(s)
}
