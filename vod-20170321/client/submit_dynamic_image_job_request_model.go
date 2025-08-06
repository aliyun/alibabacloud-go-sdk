// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageTemplateId(v string) *SubmitDynamicImageJobRequest
	GetDynamicImageTemplateId() *string
	SetOverrideParams(v string) *SubmitDynamicImageJobRequest
	GetOverrideParams() *string
	SetVideoId(v string) *SubmitDynamicImageJobRequest
	GetVideoId() *string
}

type SubmitDynamicImageJobRequest struct {
	// The ID of the frame animation template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a443dc52ef10abc4794d700*****
	DynamicImageTemplateId *string `json:"DynamicImageTemplateId,omitempty" xml:"DynamicImageTemplateId,omitempty"`
	// The override parameter. Specify the value in the JSON format. For more information, see [Parameters for media processing](https://help.aliyun.com/document_detail/98618.html). You can use this parameter to override configurations in the animated image template. For more information, see the "DynamicImageTemplateConfig: the configurations of an animated sticker template" section of the [Basic data types](https://help.aliyun.com/document_detail/52839.html) topic.
	//
	// example:
	//
	// {"Watermarks":[{"Content":"UserID: 666**","WatermarkId":"8ca03c884944bd05efccc312367****"}]}
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD](https://vod.console.aliyun.com) console. In the left-side navigation pane, choose **Media Files*	- > **Audio/Video**. On the Video and Audio page, view the ID of the media file. This method is applicable to files that are uploaded by using the ApsaraVideo VOD console.
	//
	// 	- Obtain the value of VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload media files.
	//
	// 	- Obtain the value of VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation after you upload media files.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7d2fbc3e273441bdb0e08e55f8****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitDynamicImageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequest) GetDynamicImageTemplateId() *string {
	return s.DynamicImageTemplateId
}

func (s *SubmitDynamicImageJobRequest) GetOverrideParams() *string {
	return s.OverrideParams
}

func (s *SubmitDynamicImageJobRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitDynamicImageJobRequest) SetDynamicImageTemplateId(v string) *SubmitDynamicImageJobRequest {
	s.DynamicImageTemplateId = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetOverrideParams(v string) *SubmitDynamicImageJobRequest {
	s.OverrideParams = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetVideoId(v string) *SubmitDynamicImageJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) Validate() error {
	return dara.Validate(s)
}
