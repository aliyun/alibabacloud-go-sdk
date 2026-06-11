// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSkillVersionViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DownloadSkillVersionViaOssRequest
	GetNamespaceId() *string
	SetSkillName(v string) *DownloadSkillVersionViaOssRequest
	GetSkillName() *string
	SetSkillVersion(v string) *DownloadSkillVersionViaOssRequest
	GetSkillVersion() *string
}

type DownloadSkillVersionViaOssRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customer-service-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	SkillVersion *string `json:"SkillVersion,omitempty" xml:"SkillVersion,omitempty"`
}

func (s DownloadSkillVersionViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadSkillVersionViaOssRequest) GoString() string {
	return s.String()
}

func (s *DownloadSkillVersionViaOssRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DownloadSkillVersionViaOssRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *DownloadSkillVersionViaOssRequest) GetSkillVersion() *string {
	return s.SkillVersion
}

func (s *DownloadSkillVersionViaOssRequest) SetNamespaceId(v string) *DownloadSkillVersionViaOssRequest {
	s.NamespaceId = &v
	return s
}

func (s *DownloadSkillVersionViaOssRequest) SetSkillName(v string) *DownloadSkillVersionViaOssRequest {
	s.SkillName = &v
	return s
}

func (s *DownloadSkillVersionViaOssRequest) SetSkillVersion(v string) *DownloadSkillVersionViaOssRequest {
	s.SkillVersion = &v
	return s
}

func (s *DownloadSkillVersionViaOssRequest) Validate() error {
	return dara.Validate(s)
}
