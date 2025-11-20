// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateHumanAnimeStyleVideoResponseBodyData) *GenerateHumanAnimeStyleVideoResponseBody
	GetData() *GenerateHumanAnimeStyleVideoResponseBodyData
	SetMessage(v string) *GenerateHumanAnimeStyleVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateHumanAnimeStyleVideoResponseBody
	GetRequestId() *string
}

type GenerateHumanAnimeStyleVideoResponseBody struct {
	Data    *GenerateHumanAnimeStyleVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// d21a2afa-4d52-4bca-803b-e65028146603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) GetData() *GenerateHumanAnimeStyleVideoResponseBodyData {
	return s.Data
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetData(v *GenerateHumanAnimeStyleVideoResponseBodyData) *GenerateHumanAnimeStyleVideoResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetMessage(v string) *GenerateHumanAnimeStyleVideoResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) SetRequestId(v string) *GenerateHumanAnimeStyleVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateHumanAnimeStyleVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSR****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GenerateHumanAnimeStyleVideoResponseBodyData) SetVideoUrl(v string) *GenerateHumanAnimeStyleVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
