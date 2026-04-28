// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTag interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ImageTag
	GetCount() *int64
	SetCoverFileCategory(v string) *ImageTag
	GetCoverFileCategory() *string
	SetCoverFileId(v string) *ImageTag
	GetCoverFileId() *string
	SetCoverOverallScore(v float32) *ImageTag
	GetCoverOverallScore() *float32
	SetCoverTagConfidence(v float32) *ImageTag
	GetCoverTagConfidence() *float32
	SetCoverUrl(v string) *ImageTag
	GetCoverUrl() *string
	SetName(v string) *ImageTag
	GetName() *string
}

type ImageTag struct {
	// The number of files in the group.
	//
	// example:
	//
	// 1
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The category of the cover image.
	//
	// example:
	//
	// image
	CoverFileCategory *string `json:"cover_file_category,omitempty" xml:"cover_file_category,omitempty"`
	// The ID of the cover file.
	//
	// example:
	//
	// 5d79206586bb5dd69fb34c349282718146c55da7
	CoverFileId *string `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	// The score of the cover image.
	//
	// example:
	//
	// 0.736
	CoverOverallScore *float32 `json:"cover_overall_score,omitempty" xml:"cover_overall_score,omitempty"`
	// The confidence level of the cover image tag.
	//
	// example:
	//
	// 1
	CoverTagConfidence *float32 `json:"cover_tag_confidence,omitempty" xml:"cover_tag_confidence,omitempty"`
	// The URL of the cover image.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	CoverUrl *string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// The name of the tag.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImageTag) String() string {
	return dara.Prettify(s)
}

func (s ImageTag) GoString() string {
	return s.String()
}

func (s *ImageTag) GetCount() *int64 {
	return s.Count
}

func (s *ImageTag) GetCoverFileCategory() *string {
	return s.CoverFileCategory
}

func (s *ImageTag) GetCoverFileId() *string {
	return s.CoverFileId
}

func (s *ImageTag) GetCoverOverallScore() *float32 {
	return s.CoverOverallScore
}

func (s *ImageTag) GetCoverTagConfidence() *float32 {
	return s.CoverTagConfidence
}

func (s *ImageTag) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ImageTag) GetName() *string {
	return s.Name
}

func (s *ImageTag) SetCount(v int64) *ImageTag {
	s.Count = &v
	return s
}

func (s *ImageTag) SetCoverFileCategory(v string) *ImageTag {
	s.CoverFileCategory = &v
	return s
}

func (s *ImageTag) SetCoverFileId(v string) *ImageTag {
	s.CoverFileId = &v
	return s
}

func (s *ImageTag) SetCoverOverallScore(v float32) *ImageTag {
	s.CoverOverallScore = &v
	return s
}

func (s *ImageTag) SetCoverTagConfidence(v float32) *ImageTag {
	s.CoverTagConfidence = &v
	return s
}

func (s *ImageTag) SetCoverUrl(v string) *ImageTag {
	s.CoverUrl = &v
	return s
}

func (s *ImageTag) SetName(v string) *ImageTag {
	s.Name = &v
	return s
}

func (s *ImageTag) Validate() error {
	return dara.Validate(s)
}
