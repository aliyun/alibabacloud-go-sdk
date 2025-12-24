// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDCommonImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentHDCommonImageResponseBodyData) *SegmentHDCommonImageResponseBody
	GetData() *SegmentHDCommonImageResponseBodyData
	SetMessage(v string) *SegmentHDCommonImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SegmentHDCommonImageResponseBody
	GetRequestId() *string
}

type SegmentHDCommonImageResponseBody struct {
	Data    *SegmentHDCommonImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EC994171-7964-44B3-85AF-A715CB56747D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDCommonImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDCommonImageResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponseBody) GetData() *SegmentHDCommonImageResponseBodyData {
	return s.Data
}

func (s *SegmentHDCommonImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SegmentHDCommonImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentHDCommonImageResponseBody) SetData(v *SegmentHDCommonImageResponseBodyData) *SegmentHDCommonImageResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDCommonImageResponseBody) SetMessage(v string) *SegmentHDCommonImageResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentHDCommonImageResponseBody) SetRequestId(v string) *SegmentHDCommonImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHDCommonImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentHDCommonImageResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_commoditysegmenter/2020-10-27/invi_commoditysegmenter_016037842193171000000_5pn2QM.png?Expires=1603786019&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=HwUztguGBYXmXGEmuTh%2FL3ztoh****
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s SegmentHDCommonImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDCommonImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SegmentHDCommonImageResponseBodyData) SetImageUrl(v string) *SegmentHDCommonImageResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *SegmentHDCommonImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
