// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBroadcastSceneTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCoverURL(v string) *BroadcastSceneTemplate
	GetCoverURL() *string
	SetCreateTime(v string) *BroadcastSceneTemplate
	GetCreateTime() *string
	SetDesc(v string) *BroadcastSceneTemplate
	GetDesc() *string
	SetId(v string) *BroadcastSceneTemplate
	GetId() *string
	SetModifiedTime(v string) *BroadcastSceneTemplate
	GetModifiedTime() *string
	SetName(v string) *BroadcastSceneTemplate
	GetName() *string
	SetPreviewURL(v string) *BroadcastSceneTemplate
	GetPreviewURL() *string
	SetRatio(v string) *BroadcastSceneTemplate
	GetRatio() *string
	SetShortVideoURL(v string) *BroadcastSceneTemplate
	GetShortVideoURL() *string
	SetTags(v []*string) *BroadcastSceneTemplate
	GetTags() []*string
	SetThumbnailURL(v string) *BroadcastSceneTemplate
	GetThumbnailURL() *string
}

type BroadcastSceneTemplate struct {
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CSNSAsOIDZQU2/cover.jpg
	CoverURL *string `json:"coverURL,omitempty" xml:"coverURL,omitempty"`
	// example:
	//
	// 2026-01-06T07:00:02Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// BS1tneDiuOOjJmI2qOHGw1urA
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2026-01-06T07:00:02Z
	ModifiedTime *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CSNSAsOIDZQU2/result.mp4
	PreviewURL *string `json:"previewURL,omitempty" xml:"previewURL,omitempty"`
	// example:
	//
	// 9:16
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CSNSAsOIDZQU2/result_preview.mp4
	ShortVideoURL *string   `json:"shortVideoURL,omitempty" xml:"shortVideoURL,omitempty"`
	Tags          []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// https://online-avatar-property.oss-cn-beijing.aliyuncs.com/aigc_material/OUTPUT_BROADCAST_SHORT_VIDEO/Mt.CSNSAsOIDZQU2/thumbnail.jpg
	ThumbnailURL *string `json:"thumbnailURL,omitempty" xml:"thumbnailURL,omitempty"`
}

func (s BroadcastSceneTemplate) String() string {
	return dara.Prettify(s)
}

func (s BroadcastSceneTemplate) GoString() string {
	return s.String()
}

func (s *BroadcastSceneTemplate) GetCoverURL() *string {
	return s.CoverURL
}

func (s *BroadcastSceneTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BroadcastSceneTemplate) GetDesc() *string {
	return s.Desc
}

func (s *BroadcastSceneTemplate) GetId() *string {
	return s.Id
}

func (s *BroadcastSceneTemplate) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *BroadcastSceneTemplate) GetName() *string {
	return s.Name
}

func (s *BroadcastSceneTemplate) GetPreviewURL() *string {
	return s.PreviewURL
}

func (s *BroadcastSceneTemplate) GetRatio() *string {
	return s.Ratio
}

func (s *BroadcastSceneTemplate) GetShortVideoURL() *string {
	return s.ShortVideoURL
}

func (s *BroadcastSceneTemplate) GetTags() []*string {
	return s.Tags
}

func (s *BroadcastSceneTemplate) GetThumbnailURL() *string {
	return s.ThumbnailURL
}

func (s *BroadcastSceneTemplate) SetCoverURL(v string) *BroadcastSceneTemplate {
	s.CoverURL = &v
	return s
}

func (s *BroadcastSceneTemplate) SetCreateTime(v string) *BroadcastSceneTemplate {
	s.CreateTime = &v
	return s
}

func (s *BroadcastSceneTemplate) SetDesc(v string) *BroadcastSceneTemplate {
	s.Desc = &v
	return s
}

func (s *BroadcastSceneTemplate) SetId(v string) *BroadcastSceneTemplate {
	s.Id = &v
	return s
}

func (s *BroadcastSceneTemplate) SetModifiedTime(v string) *BroadcastSceneTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *BroadcastSceneTemplate) SetName(v string) *BroadcastSceneTemplate {
	s.Name = &v
	return s
}

func (s *BroadcastSceneTemplate) SetPreviewURL(v string) *BroadcastSceneTemplate {
	s.PreviewURL = &v
	return s
}

func (s *BroadcastSceneTemplate) SetRatio(v string) *BroadcastSceneTemplate {
	s.Ratio = &v
	return s
}

func (s *BroadcastSceneTemplate) SetShortVideoURL(v string) *BroadcastSceneTemplate {
	s.ShortVideoURL = &v
	return s
}

func (s *BroadcastSceneTemplate) SetTags(v []*string) *BroadcastSceneTemplate {
	s.Tags = v
	return s
}

func (s *BroadcastSceneTemplate) SetThumbnailURL(v string) *BroadcastSceneTemplate {
	s.ThumbnailURL = &v
	return s
}

func (s *BroadcastSceneTemplate) Validate() error {
	return dara.Validate(s)
}
