// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDSkyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentHDSkyResponseBodyData) *SegmentHDSkyResponseBody
	GetData() *SegmentHDSkyResponseBodyData
	SetRequestId(v string) *SegmentHDSkyResponseBody
	GetRequestId() *string
}

type SegmentHDSkyResponseBody struct {
	Data *SegmentHDSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1173F38F-B4F4-4A07-AB2E-D490C01347E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDSkyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDSkyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponseBody) GetData() *SegmentHDSkyResponseBodyData {
	return s.Data
}

func (s *SegmentHDSkyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentHDSkyResponseBody) SetData(v *SegmentHDSkyResponseBodyData) *SegmentHDSkyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDSkyResponseBody) SetRequestId(v string) *SegmentHDSkyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHDSkyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SegmentHDSkyResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/sky-segmentation-hd/res/1173F38F-B4F4-4A07-AB2E-D490C01347E5_0d56_20201027-061858.jpg?Expires=1603781339&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2F8%2Bj%2FWruWOMqDezwpnJOkcNJD****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDSkyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHDSkyResponseBodyData) SetImageURL(v string) *SegmentHDSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentHDSkyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
