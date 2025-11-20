// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterpolateVideoFrameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *InterpolateVideoFrameResponseBodyData) *InterpolateVideoFrameResponseBody
	GetData() *InterpolateVideoFrameResponseBodyData
	SetMessage(v string) *InterpolateVideoFrameResponseBody
	GetMessage() *string
	SetRequestId(v string) *InterpolateVideoFrameResponseBody
	GetRequestId() *string
}

type InterpolateVideoFrameResponseBody struct {
	Data    *InterpolateVideoFrameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7EF85B5B-FB44-4C3E-9B8F-08C6CD912CEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InterpolateVideoFrameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InterpolateVideoFrameResponseBody) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponseBody) GetData() *InterpolateVideoFrameResponseBodyData {
	return s.Data
}

func (s *InterpolateVideoFrameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InterpolateVideoFrameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InterpolateVideoFrameResponseBody) SetData(v *InterpolateVideoFrameResponseBodyData) *InterpolateVideoFrameResponseBody {
	s.Data = v
	return s
}

func (s *InterpolateVideoFrameResponseBody) SetMessage(v string) *InterpolateVideoFrameResponseBody {
	s.Message = &v
	return s
}

func (s *InterpolateVideoFrameResponseBody) SetRequestId(v string) *InterpolateVideoFrameResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterpolateVideoFrameResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InterpolateVideoFrameResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-interp/20-12-22/mRsPNVunG7717nne_20-12-22-07-29-51.mp4?Expires=1608624020&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=NFjSLll8E7E1tkuLPFyTpr6ULi****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s InterpolateVideoFrameResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InterpolateVideoFrameResponseBodyData) GoString() string {
	return s.String()
}

func (s *InterpolateVideoFrameResponseBodyData) GetVideoURL() *string {
	return s.VideoURL
}

func (s *InterpolateVideoFrameResponseBodyData) SetVideoURL(v string) *InterpolateVideoFrameResponseBodyData {
	s.VideoURL = &v
	return s
}

func (s *InterpolateVideoFrameResponseBodyData) Validate() error {
	return dara.Validate(s)
}
