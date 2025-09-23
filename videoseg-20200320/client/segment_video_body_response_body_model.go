// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentVideoBodyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentVideoBodyResponseBodyData) *SegmentVideoBodyResponseBody
	GetData() *SegmentVideoBodyResponseBodyData
	SetMessage(v string) *SegmentVideoBodyResponseBody
	GetMessage() *string
	SetRequestId(v string) *SegmentVideoBodyResponseBody
	GetRequestId() *string
}

type SegmentVideoBodyResponseBody struct {
	Data    *SegmentVideoBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 49E2CC28-ED1D-4CC5-854D-7D0AE2B20976
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentVideoBodyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentVideoBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponseBody) GetData() *SegmentVideoBodyResponseBodyData {
	return s.Data
}

func (s *SegmentVideoBodyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SegmentVideoBodyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentVideoBodyResponseBody) SetData(v *SegmentVideoBodyResponseBodyData) *SegmentVideoBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentVideoBodyResponseBody) SetMessage(v string) *SegmentVideoBodyResponseBody {
	s.Message = &v
	return s
}

func (s *SegmentVideoBodyResponseBody) SetRequestId(v string) *SegmentVideoBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentVideoBodyResponseBody) Validate() error {
	return dara.Validate(s)
}

type SegmentVideoBodyResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-human-segmentation/D86DDFBC-B8ED-4780-9E6A-E5BA98D7CC9F.mp4?Expires=1584709406&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Fx5HVxvRjAMIjWL2OvhTlOO4cC****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SegmentVideoBodyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentVideoBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SegmentVideoBodyResponseBodyData) SetVideoUrl(v string) *SegmentVideoBodyResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *SegmentVideoBodyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
