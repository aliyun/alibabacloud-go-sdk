// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*ListSkillsResponseBodySkills) *ListSkillsResponseBody
	GetSkills() []*ListSkillsResponseBodySkills
	SetTotalCount(v int64) *ListSkillsResponseBody
	GetTotalCount() *int64
}

type ListSkillsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of skill details.
	Skills []*ListSkillsResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The total number of query results.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) GetSkills() []*ListSkillsResponseBodySkills {
	return s.Skills
}

func (s *ListSkillsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) SetSkills(v []*ListSkillsResponseBodySkills) *ListSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListSkillsResponseBody) SetTotalCount(v int64) *ListSkillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillsResponseBodySkills struct {
	// The API key of the skill.
	//
	// example:
	//
	// akm-98f66829***
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The author.
	//
	// example:
	//
	// Li***
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// The currently effective version number. If no version is effective, an empty value is returned.
	//
	// example:
	//
	// 1.0.0
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	// The description of the skill.
	//
	// example:
	//
	// This skill is used for****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// name****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the skill is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The environment variables.
	EnvVars map[string]*string `json:"EnvVars,omitempty" xml:"EnvVars,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-04-28T10:32:53Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The URL of the skill icon.
	//
	// example:
	//
	// https://***-***-****
	SkillIconUrl *string `json:"SkillIconUrl,omitempty" xml:"SkillIconUrl,omitempty"`
	// The unique identifier of the skill.
	//
	// example:
	//
	// s-04rj8mzqj1fu****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The name in the SKILL.md file.
	//
	// example:
	//
	// name****
	SkillName     *string                                      `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	SkillVersions []*ListSkillsResponseBodySkillsSkillVersions `json:"SkillVersions,omitempty" xml:"SkillVersions,omitempty" type:"Repeated"`
	// The slug identifier of the skill. This value is user-defined and unique within the tenant.
	//
	// example:
	//
	// find-skills****
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The source marketplace code.
	//
	// example:
	//
	// CLAWHUB
	SourceMarket *string `json:"SourceMarket,omitempty" xml:"SourceMarket,omitempty"`
	// The source marketplace name.
	//
	// example:
	//
	// ClawHub
	SourceMarketName *string `json:"SourceMarketName,omitempty" xml:"SourceMarketName,omitempty"`
	// The supply type.
	//
	// example:
	//
	// TENANT
	SupplierType *string `json:"SupplierType,omitempty" xml:"SupplierType,omitempty"`
}

func (s ListSkillsResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodySkills) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodySkills) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListSkillsResponseBodySkills) GetAuthor() *string {
	return s.Author
}

func (s *ListSkillsResponseBodySkills) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListSkillsResponseBodySkills) GetDescription() *string {
	return s.Description
}

func (s *ListSkillsResponseBodySkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSkillsResponseBodySkills) GetEnable() *bool {
	return s.Enable
}

func (s *ListSkillsResponseBodySkills) GetEnvVars() map[string]*string {
	return s.EnvVars
}

func (s *ListSkillsResponseBodySkills) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListSkillsResponseBodySkills) GetSkillIconUrl() *string {
	return s.SkillIconUrl
}

func (s *ListSkillsResponseBodySkills) GetSkillId() *string {
	return s.SkillId
}

func (s *ListSkillsResponseBodySkills) GetSkillName() *string {
	return s.SkillName
}

func (s *ListSkillsResponseBodySkills) GetSkillVersions() []*ListSkillsResponseBodySkillsSkillVersions {
	return s.SkillVersions
}

func (s *ListSkillsResponseBodySkills) GetSlug() *string {
	return s.Slug
}

func (s *ListSkillsResponseBodySkills) GetSourceMarket() *string {
	return s.SourceMarket
}

func (s *ListSkillsResponseBodySkills) GetSourceMarketName() *string {
	return s.SourceMarketName
}

func (s *ListSkillsResponseBodySkills) GetSupplierType() *string {
	return s.SupplierType
}

func (s *ListSkillsResponseBodySkills) SetApiKey(v string) *ListSkillsResponseBodySkills {
	s.ApiKey = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetAuthor(v string) *ListSkillsResponseBodySkills {
	s.Author = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetDefaultVersion(v string) *ListSkillsResponseBodySkills {
	s.DefaultVersion = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetDescription(v string) *ListSkillsResponseBodySkills {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetDisplayName(v string) *ListSkillsResponseBodySkills {
	s.DisplayName = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetEnable(v bool) *ListSkillsResponseBodySkills {
	s.Enable = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetEnvVars(v map[string]*string) *ListSkillsResponseBodySkills {
	s.EnvVars = v
	return s
}

func (s *ListSkillsResponseBodySkills) SetGmtCreated(v string) *ListSkillsResponseBodySkills {
	s.GmtCreated = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillIconUrl(v string) *ListSkillsResponseBodySkills {
	s.SkillIconUrl = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillId(v string) *ListSkillsResponseBodySkills {
	s.SkillId = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillName(v string) *ListSkillsResponseBodySkills {
	s.SkillName = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillVersions(v []*ListSkillsResponseBodySkillsSkillVersions) *ListSkillsResponseBodySkills {
	s.SkillVersions = v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSlug(v string) *ListSkillsResponseBodySkills {
	s.Slug = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSourceMarket(v string) *ListSkillsResponseBodySkills {
	s.SourceMarket = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSourceMarketName(v string) *ListSkillsResponseBodySkills {
	s.SourceMarketName = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSupplierType(v string) *ListSkillsResponseBodySkills {
	s.SupplierType = &v
	return s
}

func (s *ListSkillsResponseBodySkills) Validate() error {
	if s.SkillVersions != nil {
		for _, item := range s.SkillVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillsResponseBodySkillsSkillVersions struct {
	ChangeLog              *string `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty"`
	CreatedAt              *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	PublishStatus          *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	SecurityScanFailReason *string `json:"SecurityScanFailReason,omitempty" xml:"SecurityScanFailReason,omitempty"`
	SecurityScanScore      *int32  `json:"SecurityScanScore,omitempty" xml:"SecurityScanScore,omitempty"`
	SecurityScanStatus     *string `json:"SecurityScanStatus,omitempty" xml:"SecurityScanStatus,omitempty"`
	Version                *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListSkillsResponseBodySkillsSkillVersions) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodySkillsSkillVersions) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetSecurityScanFailReason() *string {
	return s.SecurityScanFailReason
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetSecurityScanScore() *int32 {
	return s.SecurityScanScore
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetSecurityScanStatus() *string {
	return s.SecurityScanStatus
}

func (s *ListSkillsResponseBodySkillsSkillVersions) GetVersion() *string {
	return s.Version
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetChangeLog(v string) *ListSkillsResponseBodySkillsSkillVersions {
	s.ChangeLog = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetCreatedAt(v int64) *ListSkillsResponseBodySkillsSkillVersions {
	s.CreatedAt = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetPublishStatus(v string) *ListSkillsResponseBodySkillsSkillVersions {
	s.PublishStatus = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetSecurityScanFailReason(v string) *ListSkillsResponseBodySkillsSkillVersions {
	s.SecurityScanFailReason = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetSecurityScanScore(v int32) *ListSkillsResponseBodySkillsSkillVersions {
	s.SecurityScanScore = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetSecurityScanStatus(v string) *ListSkillsResponseBodySkillsSkillVersions {
	s.SecurityScanStatus = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) SetVersion(v string) *ListSkillsResponseBodySkillsSkillVersions {
	s.Version = &v
	return s
}

func (s *ListSkillsResponseBodySkillsSkillVersions) Validate() error {
	return dara.Validate(s)
}
