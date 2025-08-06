// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoId(v string) *ListDynamicImageRequest
	GetVideoId() *string
}

type ListDynamicImageRequest struct {
	// The ID of the video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e114f1100590c3193918fd449a****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListDynamicImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicImageRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *ListDynamicImageRequest) SetVideoId(v string) *ListDynamicImageRequest {
	s.VideoId = &v
	return s
}

func (s *ListDynamicImageRequest) Validate() error {
	return dara.Validate(s)
}
