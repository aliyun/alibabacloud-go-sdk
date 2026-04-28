// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *GetVideoPreviewPlayInfoResponseBody
	GetDomainId() *string
	SetDriveId(v string) *GetVideoPreviewPlayInfoResponseBody
	GetDriveId() *string
	SetFileId(v string) *GetVideoPreviewPlayInfoResponseBody
	GetFileId() *string
	SetShareId(v string) *GetVideoPreviewPlayInfoResponseBody
	GetShareId() *string
	SetVideoPreviewPlayInfo(v *VideoPreviewPlayInfo) *GetVideoPreviewPlayInfoResponseBody
	GetVideoPreviewPlayInfo() *VideoPreviewPlayInfo
}

type GetVideoPreviewPlayInfoResponseBody struct {
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
	// The information about video playback.
	VideoPreviewPlayInfo *VideoPreviewPlayInfo `json:"video_preview_play_info,omitempty" xml:"video_preview_play_info,omitempty"`
}

func (s GetVideoPreviewPlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *GetVideoPreviewPlayInfoResponseBody) GetDriveId() *string {
	return s.DriveId
}

func (s *GetVideoPreviewPlayInfoResponseBody) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoPreviewPlayInfoResponseBody) GetShareId() *string {
	return s.ShareId
}

func (s *GetVideoPreviewPlayInfoResponseBody) GetVideoPreviewPlayInfo() *VideoPreviewPlayInfo {
	return s.VideoPreviewPlayInfo
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetDomainId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetDriveId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetFileId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetShareId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetVideoPreviewPlayInfo(v *VideoPreviewPlayInfo) *GetVideoPreviewPlayInfoResponseBody {
	s.VideoPreviewPlayInfo = v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) Validate() error {
	if s.VideoPreviewPlayInfo != nil {
		if err := s.VideoPreviewPlayInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
