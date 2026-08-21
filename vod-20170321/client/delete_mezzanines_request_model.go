// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMezzaninesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteMezzaninesRequest
	GetForce() *bool
	SetReferenceIds(v string) *DeleteMezzaninesRequest
	GetReferenceIds() *string
	SetVideoIds(v string) *DeleteMezzaninesRequest
	GetVideoIds() *string
}

type DeleteMezzaninesRequest struct {
	// Specifies whether to force delete the source file. Valid values:
	//
	// - **false*	- (default): No.
	//
	// - **true**: Yes.
	//
	// > If the video transcoding pattern is set to no transcoding or asynchronous transcoding, the source file is used as the original stream for playback and cannot be deleted by default. To force delete the source file of such a video, go to Settings and set this parameter to **true**.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The list of custom IDs. Specify one or more custom IDs separated by commas (,). A maximum of 20 IDs are supported.
	//
	// example:
	//
	// 123-123,1234-1234
	ReferenceIds *string `json:"ReferenceIds,omitempty" xml:"ReferenceIds,omitempty"`
	// The list of audio or video IDs whose source files you want to delete. You can specify a maximum of 20 IDs at a time. Separate multiple IDs with commas (,). You can obtain the IDs by using the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Media Files*	- > **Audio/Video*	- to view the audio or video IDs.
	//
	// - When you call the [CreateUploadVideo](https://help.aliyun.com/document_detail/55407.html) operation to obtain the upload URL and credential, the audio or video ID is the value of the VideoId response parameter.
	//
	// - After the audio or video file is uploaded, you can call the [SearchMedia](https://help.aliyun.com/document_detail/86044.html) operation to query the audio or video ID, which is the value of the VideoId response parameter.
	//
	// example:
	//
	// 23ab850b4f654b6e91d24d8157****,93ab850b4f6f4b6e91d24d81d4****
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s DeleteMezzaninesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMezzaninesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteMezzaninesRequest) GetReferenceIds() *string {
	return s.ReferenceIds
}

func (s *DeleteMezzaninesRequest) GetVideoIds() *string {
	return s.VideoIds
}

func (s *DeleteMezzaninesRequest) SetForce(v bool) *DeleteMezzaninesRequest {
	s.Force = &v
	return s
}

func (s *DeleteMezzaninesRequest) SetReferenceIds(v string) *DeleteMezzaninesRequest {
	s.ReferenceIds = &v
	return s
}

func (s *DeleteMezzaninesRequest) SetVideoIds(v string) *DeleteMezzaninesRequest {
	s.VideoIds = &v
	return s
}

func (s *DeleteMezzaninesRequest) Validate() error {
	return dara.Validate(s)
}
