// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawSkillDetailResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawSkillDetailResponseBody
	GetCode() *int32
	SetLatestVersion(v *DescribePolarClawSkillDetailResponseBodyLatestVersion) *DescribePolarClawSkillDetailResponseBody
	GetLatestVersion() *DescribePolarClawSkillDetailResponseBodyLatestVersion
	SetMessage(v string) *DescribePolarClawSkillDetailResponseBody
	GetMessage() *string
	SetOwner(v *DescribePolarClawSkillDetailResponseBodyOwner) *DescribePolarClawSkillDetailResponseBody
	GetOwner() *DescribePolarClawSkillDetailResponseBodyOwner
	SetRequestId(v string) *DescribePolarClawSkillDetailResponseBody
	GetRequestId() *string
	SetSkill(v *DescribePolarClawSkillDetailResponseBodySkill) *DescribePolarClawSkillDetailResponseBody
	GetSkill() *DescribePolarClawSkillDetailResponseBodySkill
}

type DescribePolarClawSkillDetailResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The latest version information.
	LatestVersion *DescribePolarClawSkillDetailResponseBodyLatestVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Skill author information.
	Owner *DescribePolarClawSkillDetailResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The core information of the Skill.
	Skill *DescribePolarClawSkillDetailResponseBodySkill `json:"Skill,omitempty" xml:"Skill,omitempty" type:"Struct"`
}

func (s DescribePolarClawSkillDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawSkillDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawSkillDetailResponseBody) GetLatestVersion() *DescribePolarClawSkillDetailResponseBodyLatestVersion {
	return s.LatestVersion
}

func (s *DescribePolarClawSkillDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawSkillDetailResponseBody) GetOwner() *DescribePolarClawSkillDetailResponseBodyOwner {
	return s.Owner
}

func (s *DescribePolarClawSkillDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawSkillDetailResponseBody) GetSkill() *DescribePolarClawSkillDetailResponseBodySkill {
	return s.Skill
}

func (s *DescribePolarClawSkillDetailResponseBody) SetApplicationId(v string) *DescribePolarClawSkillDetailResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetCode(v int32) *DescribePolarClawSkillDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetLatestVersion(v *DescribePolarClawSkillDetailResponseBodyLatestVersion) *DescribePolarClawSkillDetailResponseBody {
	s.LatestVersion = v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetMessage(v string) *DescribePolarClawSkillDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetOwner(v *DescribePolarClawSkillDetailResponseBodyOwner) *DescribePolarClawSkillDetailResponseBody {
	s.Owner = v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetRequestId(v string) *DescribePolarClawSkillDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) SetSkill(v *DescribePolarClawSkillDetailResponseBodySkill) *DescribePolarClawSkillDetailResponseBody {
	s.Skill = v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBody) Validate() error {
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawSkillDetailResponseBodyLatestVersion struct {
	// The version changelog.
	//
	// example:
	//
	// empty
	Changelog *string `json:"Changelog,omitempty" xml:"Changelog,omitempty"`
	// The version publish timestamp in Unix milliseconds.
	//
	// example:
	//
	// 1767545394459
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribePolarClawSkillDetailResponseBodyLatestVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponseBodyLatestVersion) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) GetChangelog() *string {
	return s.Changelog
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) SetChangelog(v string) *DescribePolarClawSkillDetailResponseBodyLatestVersion {
	s.Changelog = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) SetCreatedAt(v int64) *DescribePolarClawSkillDetailResponseBodyLatestVersion {
	s.CreatedAt = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) SetVersion(v string) *DescribePolarClawSkillDetailResponseBodyLatestVersion {
	s.Version = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyLatestVersion) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawSkillDetailResponseBodyOwner struct {
	// The display name of the author.
	//
	// example:
	//
	// Peter Steinberger
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The account identifier of the author.
	//
	// example:
	//
	// steipete
	Handle *string `json:"Handle,omitempty" xml:"Handle,omitempty"`
	// The profile picture URL.
	//
	// example:
	//
	// https://avatars.githubusercontent.com/u/58493?v=4
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The user ID of the author.
	//
	// example:
	//
	// s179zksw999xz8ms4cy7pb2fr183m5jq
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribePolarClawSkillDetailResponseBodyOwner) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) GetHandle() *string {
	return s.Handle
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) GetImage() *string {
	return s.Image
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) GetUserId() *string {
	return s.UserId
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) SetDisplayName(v string) *DescribePolarClawSkillDetailResponseBodyOwner {
	s.DisplayName = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) SetHandle(v string) *DescribePolarClawSkillDetailResponseBodyOwner {
	s.Handle = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) SetImage(v string) *DescribePolarClawSkillDetailResponseBodyOwner {
	s.Image = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) SetUserId(v string) *DescribePolarClawSkillDetailResponseBodyOwner {
	s.UserId = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodyOwner) Validate() error {
	return dara.Validate(s)
}

type DescribePolarClawSkillDetailResponseBodySkill struct {
	// The first publish timestamp in Unix milliseconds.
	//
	// example:
	//
	// 1767545394459
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The display name.
	//
	// example:
	//
	// Weather
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The Skill identifier.
	//
	// example:
	//
	// weather
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The statistics information.
	Stats *DescribePolarClawSkillDetailResponseBodySkillStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	// The brief description.
	//
	// example:
	//
	// Get current weather and forecasts (no API key required).
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The tag key-value pairs.
	//
	// example:
	//
	// {
	//
	//     "latest": "1.0.0"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The last update timestamp in Unix milliseconds.
	//
	// example:
	//
	// 1778485729679
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribePolarClawSkillDetailResponseBodySkill) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponseBodySkill) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetSlug() *string {
	return s.Slug
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetStats() *DescribePolarClawSkillDetailResponseBodySkillStats {
	return s.Stats
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetSummary() *string {
	return s.Summary
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetCreatedAt(v int64) *DescribePolarClawSkillDetailResponseBodySkill {
	s.CreatedAt = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetDisplayName(v string) *DescribePolarClawSkillDetailResponseBodySkill {
	s.DisplayName = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetSlug(v string) *DescribePolarClawSkillDetailResponseBodySkill {
	s.Slug = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetStats(v *DescribePolarClawSkillDetailResponseBodySkillStats) *DescribePolarClawSkillDetailResponseBodySkill {
	s.Stats = v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetSummary(v string) *DescribePolarClawSkillDetailResponseBodySkill {
	s.Summary = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetTags(v map[string]interface{}) *DescribePolarClawSkillDetailResponseBodySkill {
	s.Tags = v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) SetUpdatedAt(v int64) *DescribePolarClawSkillDetailResponseBodySkill {
	s.UpdatedAt = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkill) Validate() error {
	if s.Stats != nil {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawSkillDetailResponseBodySkillStats struct {
	// The number of comments.
	//
	// example:
	//
	// 6
	Comments *int64 `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The number of downloads.
	//
	// example:
	//
	// 155765
	Downloads *int64 `json:"Downloads,omitempty" xml:"Downloads,omitempty"`
	// The total number of installations of all time.
	//
	// example:
	//
	// 3787
	InstallsAllTime *int64 `json:"InstallsAllTime,omitempty" xml:"InstallsAllTime,omitempty"`
	// The current number of installations.
	//
	// example:
	//
	// 3664
	InstallsCurrent *int64 `json:"InstallsCurrent,omitempty" xml:"InstallsCurrent,omitempty"`
	// The number of stars.
	//
	// example:
	//
	// 404
	Stars *int64 `json:"Stars,omitempty" xml:"Stars,omitempty"`
	// The number of versions.
	//
	// example:
	//
	// 1
	Versions *int64 `json:"Versions,omitempty" xml:"Versions,omitempty"`
}

func (s DescribePolarClawSkillDetailResponseBodySkillStats) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillDetailResponseBodySkillStats) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetComments() *int64 {
	return s.Comments
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetDownloads() *int64 {
	return s.Downloads
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetInstallsAllTime() *int64 {
	return s.InstallsAllTime
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetInstallsCurrent() *int64 {
	return s.InstallsCurrent
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetStars() *int64 {
	return s.Stars
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) GetVersions() *int64 {
	return s.Versions
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetComments(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.Comments = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetDownloads(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.Downloads = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetInstallsAllTime(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.InstallsAllTime = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetInstallsCurrent(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.InstallsCurrent = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetStars(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.Stars = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) SetVersions(v int64) *DescribePolarClawSkillDetailResponseBodySkillStats {
	s.Versions = &v
	return s
}

func (s *DescribePolarClawSkillDetailResponseBodySkillStats) Validate() error {
	return dara.Validate(s)
}
