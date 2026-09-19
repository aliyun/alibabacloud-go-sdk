// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateCache interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *PublicTemplateCache
	GetCreatedTime() *string
	SetImageDigest(v string) *PublicTemplateCache
	GetImageDigest() *string
	SetImageSizeMB(v int32) *PublicTemplateCache
	GetImageSizeMB() *int32
	SetProgress(v int32) *PublicTemplateCache
	GetProgress() *int32
	SetReadyTime(v string) *PublicTemplateCache
	GetReadyTime() *string
	SetStatus(v string) *PublicTemplateCache
	GetStatus() *string
	SetStatusReason(v string) *PublicTemplateCache
	GetStatusReason() *string
	SetTeamID(v string) *PublicTemplateCache
	GetTeamID() *string
	SetTemplateID(v string) *PublicTemplateCache
	GetTemplateID() *string
}

type PublicTemplateCache struct {
	CreatedTime  *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	ImageDigest  *string `json:"imageDigest,omitempty" xml:"imageDigest,omitempty"`
	ImageSizeMB  *int32  `json:"imageSizeMB,omitempty" xml:"imageSizeMB,omitempty"`
	Progress     *int32  `json:"progress,omitempty" xml:"progress,omitempty"`
	ReadyTime    *string `json:"readyTime,omitempty" xml:"readyTime,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	TeamID       *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	TemplateID   *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s PublicTemplateCache) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateCache) GoString() string {
	return s.String()
}

func (s *PublicTemplateCache) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *PublicTemplateCache) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *PublicTemplateCache) GetImageSizeMB() *int32 {
	return s.ImageSizeMB
}

func (s *PublicTemplateCache) GetProgress() *int32 {
	return s.Progress
}

func (s *PublicTemplateCache) GetReadyTime() *string {
	return s.ReadyTime
}

func (s *PublicTemplateCache) GetStatus() *string {
	return s.Status
}

func (s *PublicTemplateCache) GetStatusReason() *string {
	return s.StatusReason
}

func (s *PublicTemplateCache) GetTeamID() *string {
	return s.TeamID
}

func (s *PublicTemplateCache) GetTemplateID() *string {
	return s.TemplateID
}

func (s *PublicTemplateCache) SetCreatedTime(v string) *PublicTemplateCache {
	s.CreatedTime = &v
	return s
}

func (s *PublicTemplateCache) SetImageDigest(v string) *PublicTemplateCache {
	s.ImageDigest = &v
	return s
}

func (s *PublicTemplateCache) SetImageSizeMB(v int32) *PublicTemplateCache {
	s.ImageSizeMB = &v
	return s
}

func (s *PublicTemplateCache) SetProgress(v int32) *PublicTemplateCache {
	s.Progress = &v
	return s
}

func (s *PublicTemplateCache) SetReadyTime(v string) *PublicTemplateCache {
	s.ReadyTime = &v
	return s
}

func (s *PublicTemplateCache) SetStatus(v string) *PublicTemplateCache {
	s.Status = &v
	return s
}

func (s *PublicTemplateCache) SetStatusReason(v string) *PublicTemplateCache {
	s.StatusReason = &v
	return s
}

func (s *PublicTemplateCache) SetTeamID(v string) *PublicTemplateCache {
	s.TeamID = &v
	return s
}

func (s *PublicTemplateCache) SetTemplateID(v string) *PublicTemplateCache {
	s.TemplateID = &v
	return s
}

func (s *PublicTemplateCache) Validate() error {
	return dara.Validate(s)
}
