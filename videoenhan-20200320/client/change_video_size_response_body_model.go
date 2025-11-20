// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVideoSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChangeVideoSizeResponseBodyData) *ChangeVideoSizeResponseBody
	GetData() *ChangeVideoSizeResponseBodyData
	SetMessage(v string) *ChangeVideoSizeResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChangeVideoSizeResponseBody
	GetRequestId() *string
}

type ChangeVideoSizeResponseBody struct {
	Data    *ChangeVideoSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C00C5A32-9F54-44F0-9778-0968DD9BF22A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeVideoSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeVideoSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseBody) GetData() *ChangeVideoSizeResponseBodyData {
	return s.Data
}

func (s *ChangeVideoSizeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeVideoSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeVideoSizeResponseBody) SetData(v *ChangeVideoSizeResponseBodyData) *ChangeVideoSizeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeVideoSizeResponseBody) SetMessage(v string) *ChangeVideoSizeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeVideoSizeResponseBody) SetRequestId(v string) *ChangeVideoSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeVideoSizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeVideoSizeResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-crop/2020-07-24-20/ZTZslWcU.jpg?Expires=1595597077&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=5cq1NNBEzS64U7RTXRBGlo7WPy****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-crop/2020-07-24-20/ZTZslWcU.mp4?Expires=1595597077&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=ZyvD9AXCT2IUFkVJngQdbXMwX6****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ChangeVideoSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeVideoSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeResponseBodyData) GetVideoCoverUrl() *string {
	return s.VideoCoverUrl
}

func (s *ChangeVideoSizeResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *ChangeVideoSizeResponseBodyData) SetVideoCoverUrl(v string) *ChangeVideoSizeResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *ChangeVideoSizeResponseBodyData) SetVideoUrl(v string) *ChangeVideoSizeResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *ChangeVideoSizeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
