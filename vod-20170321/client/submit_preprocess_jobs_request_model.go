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
	// The preprocessing type. Set the value to **LivePreprocess**. LivePreprocess specifies that the video is preprocessed in the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// LivePreprocess
	PreprocessType *string `json:"PreprocessType,omitempty" xml:"PreprocessType,omitempty"`
	// The ID of the video. You can use one of the following methods to obtain the ID:
	//
	// 	- After you upload a video in the ApsaraVideo VOD console, you can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the ID of the video.
	//
	// 	- Obtain the VideoId from the response to the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation that you call to upload videos.
	//
	// 	- Obtain the VideoId from the response to the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation that you call to query videos.
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
