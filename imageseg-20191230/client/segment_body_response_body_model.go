// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentBodyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentBodyResponseBodyData) *SegmentBodyResponseBody
	GetData() *SegmentBodyResponseBodyData
	SetRequestId(v string) *SegmentBodyResponseBody
	GetRequestId() *string
}

type SegmentBodyResponseBody struct {
	Data *SegmentBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 30EDCEEA-2806-44C6-AF0B-0988849106FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentBodyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBody) GetData() *SegmentBodyResponseBodyData {
	return s.Data
}

func (s *SegmentBodyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentBodyResponseBody) SetData(v *SegmentBodyResponseBodyData) *SegmentBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentBodyResponseBody) SetRequestId(v string) *SegmentBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentBodyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentBodyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_humansegmenter/2021-3-31/invi_humansegmenter_016171823500001081370_Ej0WwO.jpg?Expires=1617184150&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=ZwaWXpAOMzHar%2B1wVO7zeSD83r****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentBodyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentBodyResponseBodyData) SetImageURL(v string) *SegmentBodyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentBodyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
