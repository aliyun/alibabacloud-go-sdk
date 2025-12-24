// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommodityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentCommodityResponseBodyData) *SegmentCommodityResponseBody
	GetData() *SegmentCommodityResponseBodyData
	SetRequestId(v string) *SegmentCommodityResponseBody
	GetRequestId() *string
}

type SegmentCommodityResponseBody struct {
	Data *SegmentCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentCommodityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBody) GetData() *SegmentCommodityResponseBodyData {
	return s.Data
}

func (s *SegmentCommodityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentCommodityResponseBody) SetData(v *SegmentCommodityResponseBodyData) *SegmentCommodityResponseBody {
	s.Data = v
	return s
}

func (s *SegmentCommodityResponseBody) SetRequestId(v string) *SegmentCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentCommodityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentCommodityResponseBodyData struct {
	// example:
	//
	// http://luban-vgd-invi.oss-cn-hangzhou.aliyuncs.com/upload/result_segmenter/2019-12-20/invi_segmenter_015768355410261076021_Z3t0fc.png?Expires=1577094741&OSSAccessKeyId=LTAI4Fc5SVvzUQ19K1Cz****&Signature=pkaKK3VlfsTR2r%2BYycJzTVEEos****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommodityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentCommodityResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentCommodityResponseBodyData) SetImageURL(v string) *SegmentCommodityResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommodityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
