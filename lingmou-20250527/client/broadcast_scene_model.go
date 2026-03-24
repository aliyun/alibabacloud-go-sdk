// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBroadcastScene interface {
	dara.Model
	String() string
	GoString() string
	SetClipInfo(v string) *BroadcastScene
	GetClipInfo() *string
	SetCoverURL(v string) *BroadcastScene
	GetCoverURL() *string
	SetCreateTime(v string) *BroadcastScene
	GetCreateTime() *string
	SetDownloadURL(v string) *BroadcastScene
	GetDownloadURL() *string
	SetId(v string) *BroadcastScene
	GetId() *string
	SetModifiedTime(v string) *BroadcastScene
	GetModifiedTime() *string
	SetName(v string) *BroadcastScene
	GetName() *string
	SetPreviewURL(v string) *BroadcastScene
	GetPreviewURL() *string
	SetRatio(v string) *BroadcastScene
	GetRatio() *string
	SetRemainSeconds(v int64) *BroadcastScene
	GetRemainSeconds() *int64
	SetShortVideoURL(v string) *BroadcastScene
	GetShortVideoURL() *string
	SetStatus(v string) *BroadcastScene
	GetStatus() *string
	SetThumbnailURL(v string) *BroadcastScene
	GetThumbnailURL() *string
	SetVersion(v int64) *BroadcastScene
	GetVersion() *int64
}

type BroadcastScene struct {
	// example:
	//
	// {}
	ClipInfo *string `json:"clipInfo,omitempty" xml:"clipInfo,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/BS1CQEYXYQW4MQU2/cover.jpg
	CoverURL *string `json:"coverURL,omitempty" xml:"coverURL,omitempty"`
	// example:
	//
	// 2026-03-24T11:21:27.691732
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/BS1CQEYXYQW4MQU2/result.mp4
	DownloadURL *string `json:"downloadURL,omitempty" xml:"downloadURL,omitempty"`
	// example:
	//
	// BS1WgG5zb-N1GI8nId3r6wo8g
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2026-03-24T11:21:27.691732
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/BS1CQEYXYQW4MQU2/preview.mp4
	PreviewURL *string `json:"previewURL,omitempty" xml:"previewURL,omitempty"`
	// example:
	//
	// 9:16
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// example:
	//
	// 300
	RemainSeconds *int64 `json:"remainSeconds,omitempty" xml:"remainSeconds,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/BS1CQEYXYQW4MQU2/result.mp4
	ShortVideoURL *string `json:"shortVideoURL,omitempty" xml:"shortVideoURL,omitempty"`
	// example:
	//
	// DRAFT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/BS1CQEYXYQW4MQU2/thumbnail.jpg
	ThumbnailURL *string `json:"thumbnailURL,omitempty" xml:"thumbnailURL,omitempty"`
	// example:
	//
	// 0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s BroadcastScene) String() string {
	return dara.Prettify(s)
}

func (s BroadcastScene) GoString() string {
	return s.String()
}

func (s *BroadcastScene) GetClipInfo() *string {
	return s.ClipInfo
}

func (s *BroadcastScene) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BroadcastScene) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BroadcastScene) GetDownloadURL() *string {
	return s.DownloadURL
}

func (s *BroadcastScene) GetId() *string {
	return s.Id
}

func (s *BroadcastScene) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BroadcastScene) GetName() *string {
	return s.Name
}

func (s *BroadcastScene) GetPreviewURL() *string {
	return s.PreviewURL
}

func (s *BroadcastScene) GetRatio() *string {
	return s.Ratio
}

func (s *BroadcastScene) GetRemainSeconds() *int64 {
	return s.RemainSeconds
}

func (s *BroadcastScene) GetShortVideoURL() *string {
	return s.ShortVideoURL
}

func (s *BroadcastScene) GetStatus() *string {
	return s.Status
}

func (s *BroadcastScene) GetThumbnailURL() *string {
	return s.ThumbnailURL
}

func (s *BroadcastScene) GetVersion() *int64 {
	return s.Version
}

func (s *BroadcastScene) SetClipInfo(v string) *BroadcastScene {
	s.ClipInfo = &v
	return s
}

func (s *BroadcastScene) SetCoverURL(v string) *BroadcastScene {
	s.CoverURL = &v
	return s
}

func (s *BroadcastScene) SetCreateTime(v string) *BroadcastScene {
	s.CreateTime = &v
	return s
}

func (s *BroadcastScene) SetDownloadURL(v string) *BroadcastScene {
	s.DownloadURL = &v
	return s
}

func (s *BroadcastScene) SetId(v string) *BroadcastScene {
	s.Id = &v
	return s
}

func (s *BroadcastScene) SetModifiedTime(v string) *BroadcastScene {
	s.ModifiedTime = &v
	return s
}

func (s *BroadcastScene) SetName(v string) *BroadcastScene {
	s.Name = &v
	return s
}

func (s *BroadcastScene) SetPreviewURL(v string) *BroadcastScene {
	s.PreviewURL = &v
	return s
}

func (s *BroadcastScene) SetRatio(v string) *BroadcastScene {
	s.Ratio = &v
	return s
}

func (s *BroadcastScene) SetRemainSeconds(v int64) *BroadcastScene {
	s.RemainSeconds = &v
	return s
}

func (s *BroadcastScene) SetShortVideoURL(v string) *BroadcastScene {
	s.ShortVideoURL = &v
	return s
}

func (s *BroadcastScene) SetStatus(v string) *BroadcastScene {
	s.Status = &v
	return s
}

func (s *BroadcastScene) SetThumbnailURL(v string) *BroadcastScene {
	s.ThumbnailURL = &v
	return s
}

func (s *BroadcastScene) SetVersion(v int64) *BroadcastScene {
	s.Version = &v
	return s
}

func (s *BroadcastScene) Validate() error {
	return dara.Validate(s)
}
