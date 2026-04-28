// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoPreviewSubtitleInfo interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *VideoPreviewSubtitleInfo
	GetLanguage() *string
	SetStatus(v string) *VideoPreviewSubtitleInfo
	GetStatus() *string
	SetUrl(v string) *VideoPreviewSubtitleInfo
	GetUrl() *string
}

type VideoPreviewSubtitleInfo struct {
	// The subtitle language.
	//
	// example:
	//
	// en
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The status of the subtitle task.
	//
	// Valid values:
	//
	// 	- finished
	//
	// 	- failed
	//
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The subtitle URL.
	//
	// example:
	//
	// https://example.data.aliyunpds.com/lt/A05EF408DAB5D3F57C94F67658C99C406EFCA7DD/subtitle/subtitle_0.vtt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewSubtitleInfo) String() string {
	return dara.Prettify(s)
}

func (s VideoPreviewSubtitleInfo) GoString() string {
	return s.String()
}

func (s *VideoPreviewSubtitleInfo) GetLanguage() *string {
	return s.Language
}

func (s *VideoPreviewSubtitleInfo) GetStatus() *string {
	return s.Status
}

func (s *VideoPreviewSubtitleInfo) GetUrl() *string {
	return s.Url
}

func (s *VideoPreviewSubtitleInfo) SetLanguage(v string) *VideoPreviewSubtitleInfo {
	s.Language = &v
	return s
}

func (s *VideoPreviewSubtitleInfo) SetStatus(v string) *VideoPreviewSubtitleInfo {
	s.Status = &v
	return s
}

func (s *VideoPreviewSubtitleInfo) SetUrl(v string) *VideoPreviewSubtitleInfo {
	s.Url = &v
	return s
}

func (s *VideoPreviewSubtitleInfo) Validate() error {
	return dara.Validate(s)
}
