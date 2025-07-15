// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAISubtitleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLiveAISubtitleRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveAISubtitleRequest
	GetRegionId() *string
	SetSubtitleId(v string) *DeleteLiveAISubtitleRequest
	GetSubtitleId() *string
	SetSubtitleName(v string) *DeleteLiveAISubtitleRequest
	GetSubtitleName() *string
}

type DeleteLiveAISubtitleRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the subtitle template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	SubtitleId *string `json:"SubtitleId,omitempty" xml:"SubtitleId,omitempty"`
	// The name of the subtitle template. The name can contain only digits, letters, and hyphens (-). The name cannot start with a hyphen.
	//
	// example:
	//
	// sub01
	SubtitleName *string `json:"SubtitleName,omitempty" xml:"SubtitleName,omitempty"`
}

func (s DeleteLiveAISubtitleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAISubtitleRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveAISubtitleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveAISubtitleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveAISubtitleRequest) GetSubtitleId() *string {
	return s.SubtitleId
}

func (s *DeleteLiveAISubtitleRequest) GetSubtitleName() *string {
	return s.SubtitleName
}

func (s *DeleteLiveAISubtitleRequest) SetOwnerId(v int64) *DeleteLiveAISubtitleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveAISubtitleRequest) SetRegionId(v string) *DeleteLiveAISubtitleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveAISubtitleRequest) SetSubtitleId(v string) *DeleteLiveAISubtitleRequest {
	s.SubtitleId = &v
	return s
}

func (s *DeleteLiveAISubtitleRequest) SetSubtitleName(v string) *DeleteLiveAISubtitleRequest {
	s.SubtitleName = &v
	return s
}

func (s *DeleteLiveAISubtitleRequest) Validate() error {
	return dara.Validate(s)
}
