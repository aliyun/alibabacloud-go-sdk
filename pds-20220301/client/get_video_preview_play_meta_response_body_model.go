// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *GetVideoPreviewPlayMetaResponseBody
	GetDomainId() *string
	SetDriveId(v string) *GetVideoPreviewPlayMetaResponseBody
	GetDriveId() *string
	SetFileId(v string) *GetVideoPreviewPlayMetaResponseBody
	GetFileId() *string
	SetShareId(v string) *GetVideoPreviewPlayMetaResponseBody
	GetShareId() *string
	SetVideoPreviewPlayMeta(v *VideoPreviewPlayMeta) *GetVideoPreviewPlayMetaResponseBody
	GetVideoPreviewPlayMeta() *VideoPreviewPlayMeta
}

type GetVideoPreviewPlayMetaResponseBody struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The preview metadata of the video.
	VideoPreviewPlayMeta *VideoPreviewPlayMeta `json:"video_preview_play_meta,omitempty" xml:"video_preview_play_meta,omitempty"`
}

func (s GetVideoPreviewPlayMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *GetVideoPreviewPlayMetaResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *GetVideoPreviewPlayMetaResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoPreviewPlayMetaResponseBody) GetShareId() *string {
	return s.ShareId
}

func (s *GetVideoPreviewPlayMetaResponseBody) GetVideoPreviewPlayMeta() *VideoPreviewPlayMeta {
	return s.VideoPreviewPlayMeta
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetDomainId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetDriveId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetFileId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetShareId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetVideoPreviewPlayMeta(v *VideoPreviewPlayMeta) *GetVideoPreviewPlayMetaResponseBody {
	s.VideoPreviewPlayMeta = v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) Validate() error {
	if s.VideoPreviewPlayMeta != nil {
		if err := s.VideoPreviewPlayMeta.Validate(); err != nil {
			return err
		}
	}
	return nil
}
