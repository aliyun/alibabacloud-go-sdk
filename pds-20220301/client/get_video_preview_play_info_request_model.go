// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetVideoPreviewPlayInfoRequest
	GetCategory() *string
	SetDriveId(v string) *GetVideoPreviewPlayInfoRequest
	GetDriveId() *string
	SetFileId(v string) *GetVideoPreviewPlayInfoRequest
	GetFileId() *string
	SetGetMasterUrl(v bool) *GetVideoPreviewPlayInfoRequest
	GetGetMasterUrl() *bool
	SetGetWithoutUrl(v bool) *GetVideoPreviewPlayInfoRequest
	GetGetWithoutUrl() *bool
	SetReTranscode(v bool) *GetVideoPreviewPlayInfoRequest
	GetReTranscode() *bool
	SetShareId(v string) *GetVideoPreviewPlayInfoRequest
	GetShareId() *string
	SetTemplateId(v string) *GetVideoPreviewPlayInfoRequest
	GetTemplateId() *string
	SetUrlExpireSec(v int64) *GetVideoPreviewPlayInfoRequest
	GetUrlExpireSec() *int64
}

type GetVideoPreviewPlayInfoRequest struct {
	// The category. It is the transcoding mode that you want to use. Valid values:
	//
	// 	- live_transcoding: plays a live stream while transcoding is in progress.
	//
	// 	- quick_video: plays a video while transcoding is in progress.
	//
	// 	- offline_audio: plays a piece of audio after the audio is transcoded offline.
	//
	// 	- offline_video: plays a video after the video is transcoded offline.
	//
	// This parameter is required.
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to obtain the URL of the master M3U8 playlist. This parameter is valid only if the category parameter is set to quick_video.
	//
	// example:
	//
	// false
	GetMasterUrl *bool `json:"get_master_url,omitempty" xml:"get_master_url,omitempty"`
	// Specifies whether not to query the playback URL. If you set this parameter to true, only transcoding metadata is returned. The video is not transcoded in the TS format, and the playback URL is not returned. If you set this parameter to false, the playback URL is returned. If the video has not been transcoded by using the template specified by template_id, the transcoding process is triggered. You are charged value-added service fees generated for transcoding.
	//
	// example:
	//
	// true
	GetWithoutUrl *bool `json:"get_without_url,omitempty" xml:"get_without_url,omitempty"`
	// Specifies whether to initiate re-transcoding. If you set this parameter to true, the file is re-transcoded, with a fixed 202 response for retries. Before you use this parameter, contact us to enable it for you.
	//
	// example:
	//
	// true
	ReTranscode *bool `json:"re_transcode,omitempty" xml:"re_transcode,omitempty"`
	// The share ID. If you want to share a file, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the definition template. If you leave this parameter empty, all definition templates are available.
	//
	// example:
	//
	// 264_480p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// The validity period of the URL. Unit: seconds. Default value: 900, which is 15 minutes. Maximum value: 14400, which is 4 hours.
	//
	// example:
	//
	// 3600
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetVideoPreviewPlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoRequest) GetCategory() *string {
	return s.Category
}

func (s *GetVideoPreviewPlayInfoRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetVideoPreviewPlayInfoRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoPreviewPlayInfoRequest) GetGetMasterUrl() *bool {
	return s.GetMasterUrl
}

func (s *GetVideoPreviewPlayInfoRequest) GetGetWithoutUrl() *bool {
	return s.GetWithoutUrl
}

func (s *GetVideoPreviewPlayInfoRequest) GetReTranscode() *bool {
	return s.ReTranscode
}

func (s *GetVideoPreviewPlayInfoRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetVideoPreviewPlayInfoRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetVideoPreviewPlayInfoRequest) GetUrlExpireSec() *int64 {
	return s.UrlExpireSec
}

func (s *GetVideoPreviewPlayInfoRequest) SetCategory(v string) *GetVideoPreviewPlayInfoRequest {
	s.Category = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetDriveId(v string) *GetVideoPreviewPlayInfoRequest {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetFileId(v string) *GetVideoPreviewPlayInfoRequest {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetGetMasterUrl(v bool) *GetVideoPreviewPlayInfoRequest {
	s.GetMasterUrl = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetGetWithoutUrl(v bool) *GetVideoPreviewPlayInfoRequest {
	s.GetWithoutUrl = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetReTranscode(v bool) *GetVideoPreviewPlayInfoRequest {
	s.ReTranscode = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetShareId(v string) *GetVideoPreviewPlayInfoRequest {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetTemplateId(v string) *GetVideoPreviewPlayInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetUrlExpireSec(v int64) *GetVideoPreviewPlayInfoRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) Validate() error {
	return dara.Validate(s)
}
