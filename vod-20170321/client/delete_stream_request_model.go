// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *DeleteStreamRequest
	GetJobIds() *string
	SetVideoId(v string) *DeleteStreamRequest
	GetVideoId() *string
}

type DeleteStreamRequest struct {
	// The job IDs for deleting media streams.
	//
	// 	- Separate multiple IDs with commas (,). A maximum of 20 IDs can be specified for one video.
	//
	// 	- You can obtain job IDs from the PlayInfo parameter that is returned after you call the [GetPlayInfo](https://help.aliyun.com/document_detail/56124.html) operation. Each media stream has a unique job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35eb4dbda18c49cc0025df374b46****
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95948ddba24446b6aed5db985e78****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *DeleteStreamRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *DeleteStreamRequest) SetJobIds(v string) *DeleteStreamRequest {
	s.JobIds = &v
	return s
}

func (s *DeleteStreamRequest) SetVideoId(v string) *DeleteStreamRequest {
	s.VideoId = &v
	return s
}

func (s *DeleteStreamRequest) Validate() error {
	return dara.Validate(s)
}
