// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageIds(v string) *DeleteDynamicImageRequest
	GetDynamicImageIds() *string
	SetVideoId(v string) *DeleteDynamicImageRequest
	GetVideoId() *string
}

type DeleteDynamicImageRequest struct {
	// The IDs of the animated stickers.
	//
	// - Separate multiple IDs with commas (,). You can specify a maximum of 10 IDs.
	//
	// - If you do not set this parameter, the system finds the video specified by the VideoId parameter and deletes the information about the animated stickers associated with the video. If more than 10 animated stickers are associated with the video specified by the VideoId parameter, the deletion request is denied.
	//
	// example:
	//
	// beafec3834a4e52ea52042a4****,8281c8519847fd8970e79e80b6****
	DynamicImageIds *string `json:"DynamicImageIds,omitempty" xml:"DynamicImageIds,omitempty"`
	// The ID of the video associated with the animated stickers whose information you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2321077d460b028700ef6c2f4d****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteDynamicImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageRequest) GetDynamicImageIds() *string {
	return s.DynamicImageIds
}

func (s *DeleteDynamicImageRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *DeleteDynamicImageRequest) SetDynamicImageIds(v string) *DeleteDynamicImageRequest {
	s.DynamicImageIds = &v
	return s
}

func (s *DeleteDynamicImageRequest) SetVideoId(v string) *DeleteDynamicImageRequest {
	s.VideoId = &v
	return s
}

func (s *DeleteDynamicImageRequest) Validate() error {
	return dara.Validate(s)
}
