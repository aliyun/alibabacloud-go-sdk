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
	// The ID of the animated image template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a443dc52ef10abc4794d700*****
	DynamicImageTemplateId *string `json:"DynamicImageTemplateId,omitempty" xml:"DynamicImageTemplateId,omitempty"`
	// The override parameters in the JSON format. For more information, see [OverrideParams](https://help.aliyun.com/document_detail/98618.html). You can use this parameter to override the parameters in the animated image template. For more information, see [DynamicImageTemplateConfig](https://help.aliyun.com/document_detail/52839.html).
	//
	// example:
	//
	// {"Watermarks":[{"Content":"UserID: 666**","WatermarkId":"8ca03c884944bd05efccc312367****"}]}
	OverrideParams *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Assets*	- > **Audio/Video*	- to view the video ID.
	//
	// - Obtain the video ID from the value of the VideoId parameter returned by the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation when you obtain the upload URL and credential.
	//
	// - After the video is uploaded, obtain the video ID from the value of the VideoId parameter returned by the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation.
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
