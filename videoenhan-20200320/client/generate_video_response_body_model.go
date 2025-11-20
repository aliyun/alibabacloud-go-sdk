// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateVideoResponseBodyData) *GenerateVideoResponseBody
	GetData() *GenerateVideoResponseBodyData
	SetMessage(v string) *GenerateVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateVideoResponseBody
	GetRequestId() *string
}

type GenerateVideoResponseBody struct {
	Data    *GenerateVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7CB9B663-3EF8-4C9C-A464-FDA2B5F1E3A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseBody) GetData() *GenerateVideoResponseBodyData {
	return s.Data
}

func (s *GenerateVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateVideoResponseBody) SetData(v *GenerateVideoResponseBodyData) *GenerateVideoResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVideoResponseBody) SetMessage(v string) *GenerateVideoResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateVideoResponseBody) SetRequestId(v string) *GenerateVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-gen/2021-05-07-15/B9MGfwxu.mp4?Expires=1620372653&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=F9flL1n5GPYaae0dLl%2F8D%2Bn4j6****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/video-gen/2021-05-07-15/B9MGfwxu.jpg?Expires=1620372653&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=djBvGvdJu8bd%2FC%2BVHdg1d57U%2Bu****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVideoResponseBodyData) GetVideoCoverUrl() *string {
	return s.VideoCoverUrl
}

func (s *GenerateVideoResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GenerateVideoResponseBodyData) SetVideoCoverUrl(v string) *GenerateVideoResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *GenerateVideoResponseBodyData) SetVideoUrl(v string) *GenerateVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *GenerateVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
