// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommonImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentCommonImageResponseBodyData) *SegmentCommonImageResponseBody
	GetData() *SegmentCommonImageResponseBodyData
	SetRequestId(v string) *SegmentCommonImageResponseBody
	GetRequestId() *string
}

type SegmentCommonImageResponseBody struct {
	Data *SegmentCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1B8BEF02-0672-44CA-A974-4D6FAC392497
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentCommonImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommonImageResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponseBody) GetData() *SegmentCommonImageResponseBodyData {
	return s.Data
}

func (s *SegmentCommonImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentCommonImageResponseBody) SetData(v *SegmentCommonImageResponseBodyData) *SegmentCommonImageResponseBody {
	s.Data = v
	return s
}

func (s *SegmentCommonImageResponseBody) SetRequestId(v string) *SegmentCommonImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentCommonImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentCommonImageResponseBodyData struct {
	// example:
	//
	// http://luban-vgd-invi.oss-cn-hangzhou.aliyuncs.com/upload/result_segmenter/2019-12-20/invi_segmenter_015768355410261076021_Z3t0fc.png?Expires=1577094741&OSSAccessKeyId=LTAI4Fc5SVvzUQ19K1Cz****&Signature=pkaKK3VlfsTR2r%2BYycJzTVEEos****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentCommonImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommonImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentCommonImageResponseBodyData) SetImageURL(v string) *SegmentCommonImageResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentCommonImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
