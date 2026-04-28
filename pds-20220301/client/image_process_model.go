// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageProcess interface {
	dara.Model
	String() string
	GoString() string
	SetImageThumbnailProcess(v string) *ImageProcess
	GetImageThumbnailProcess() *string
	SetOfficeThumbnailProcess(v string) *ImageProcess
	GetOfficeThumbnailProcess() *string
	SetVideoThumbnailProcess(v string) *ImageProcess
	GetVideoThumbnailProcess() *string
}

type ImageProcess struct {
	// The thumbnail processing rules for images. For more information, see the "IMG implementation modes" topic of Object Storage Service (OSS). Default value: image/resize,m_fill,h_128,w_128,limit_0.
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The thumbnail processing rules for documents. For a document, the snapshot of one of the pages in the document is used as the thumbnail. This parameter takes effect on this snapshot. Default value: image/resize,m_fill,h_128,w_128,limit_0.
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0
	OfficeThumbnailProcess *string `json:"office_thumbnail_process,omitempty" xml:"office_thumbnail_process,omitempty"`
	// The thumbnail processing rules for videos. For more information, see the "Video snapshots" topic of OSS. Default value: video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto.
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ImageProcess) String() string {
	return dara.Prettify(s)
}

func (s ImageProcess) GoString() string {
	return s.String()
}

func (s *ImageProcess) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *ImageProcess) GetOfficeThumbnailProcess() *string {
	return s.OfficeThumbnailProcess
}

func (s *ImageProcess) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *ImageProcess) SetImageThumbnailProcess(v string) *ImageProcess {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ImageProcess) SetOfficeThumbnailProcess(v string) *ImageProcess {
	s.OfficeThumbnailProcess = &v
	return s
}

func (s *ImageProcess) SetVideoThumbnailProcess(v string) *ImageProcess {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *ImageProcess) Validate() error {
	return dara.Validate(s)
}
