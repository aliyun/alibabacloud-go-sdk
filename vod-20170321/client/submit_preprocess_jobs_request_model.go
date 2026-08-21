// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPreprocessType(v string) *SubmitPreprocessJobsRequest
	GetPreprocessType() *string
	SetVideoId(v string) *SubmitPreprocessJobsRequest
	GetVideoId() *string
}

type SubmitPreprocessJobsRequest struct {
	// The preprocessing type. Set the value to **LivePreprocess*	- (video preprocessing for the China Production Studio).
	//
	// This parameter is required.
	//
	// example:
	//
	// LivePreprocess
	PreprocessType *string `json:"PreprocessType,omitempty" xml:"PreprocessType,omitempty"`
	// The video ID. You can obtain the video ID by using one of the following methods:
	//
	// - For videos uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the video ID.
	//
	// - When you upload a video by calling the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation, the video ID is the value of the VideoId parameter in the response.
	//
	// - After the video is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the video ID, which is the value of the VideoId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3e680e618708efbf2cae7cc9312****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitPreprocessJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsRequest) GetPreprocessType() *string {
	return s.PreprocessType
}

func (s *SubmitPreprocessJobsRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitPreprocessJobsRequest) SetPreprocessType(v string) *SubmitPreprocessJobsRequest {
	s.PreprocessType = &v
	return s
}

func (s *SubmitPreprocessJobsRequest) SetVideoId(v string) *SubmitPreprocessJobsRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitPreprocessJobsRequest) Validate() error {
	return dara.Validate(s)
}
