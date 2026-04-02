// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUsageInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetCompletionTokens(v int32) *UsageInfoDTO
	GetCompletionTokens() *int32
	SetImageCount(v int32) *UsageInfoDTO
	GetImageCount() *int32
	SetPromptTokens(v int32) *UsageInfoDTO
	GetPromptTokens() *int32
	SetTotalTokens(v int32) *UsageInfoDTO
	GetTotalTokens() *int32
	SetVideoCount(v int32) *UsageInfoDTO
	GetVideoCount() *int32
	SetVideoDuration(v int32) *UsageInfoDTO
	GetVideoDuration() *int32
}

type UsageInfoDTO struct {
	// example:
	//
	// 50
	CompletionTokens *int32 `json:"completionTokens,omitempty" xml:"completionTokens,omitempty"`
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 100
	PromptTokens *int32 `json:"promptTokens,omitempty" xml:"promptTokens,omitempty"`
	// example:
	//
	// 150
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	// example:
	//
	// 0
	VideoCount *int32 `json:"videoCount,omitempty" xml:"videoCount,omitempty"`
	// example:
	//
	// 0
	VideoDuration *int32 `json:"videoDuration,omitempty" xml:"videoDuration,omitempty"`
}

func (s UsageInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s UsageInfoDTO) GoString() string {
	return s.String()
}

func (s *UsageInfoDTO) GetCompletionTokens() *int32 {
	return s.CompletionTokens
}

func (s *UsageInfoDTO) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *UsageInfoDTO) GetPromptTokens() *int32 {
	return s.PromptTokens
}

func (s *UsageInfoDTO) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *UsageInfoDTO) GetVideoCount() *int32 {
	return s.VideoCount
}

func (s *UsageInfoDTO) GetVideoDuration() *int32 {
	return s.VideoDuration
}

func (s *UsageInfoDTO) SetCompletionTokens(v int32) *UsageInfoDTO {
	s.CompletionTokens = &v
	return s
}

func (s *UsageInfoDTO) SetImageCount(v int32) *UsageInfoDTO {
	s.ImageCount = &v
	return s
}

func (s *UsageInfoDTO) SetPromptTokens(v int32) *UsageInfoDTO {
	s.PromptTokens = &v
	return s
}

func (s *UsageInfoDTO) SetTotalTokens(v int32) *UsageInfoDTO {
	s.TotalTokens = &v
	return s
}

func (s *UsageInfoDTO) SetVideoCount(v int32) *UsageInfoDTO {
	s.VideoCount = &v
	return s
}

func (s *UsageInfoDTO) SetVideoDuration(v int32) *UsageInfoDTO {
	s.VideoDuration = &v
	return s
}

func (s *UsageInfoDTO) Validate() error {
	return dara.Validate(s)
}
