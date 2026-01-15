// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppInstance interface {
	dara.Model
	String() string
	GoString() string
	SetAppSubType(v string) *AppInstance
	GetAppSubType() *string
	SetAppType(v string) *AppInstance
	GetAppType() *string
	SetBizId(v string) *AppInstance
	GetBizId() *string
	SetBuildType(v string) *AppInstance
	GetBuildType() *string
	SetDeleted(v int32) *AppInstance
	GetDeleted() *int32
	SetDescription(v string) *AppInstance
	GetDescription() *string
	SetDesignSpecBizId(v string) *AppInstance
	GetDesignSpecBizId() *string
	SetDesignSpecId(v string) *AppInstance
	GetDesignSpecId() *string
	SetDomain(v string) *AppInstance
	GetDomain() *string
	SetEndTime(v string) *AppInstance
	GetEndTime() *string
	SetEspBizId(v string) *AppInstance
	GetEspBizId() *string
	SetGmtCreate(v string) *AppInstance
	GetGmtCreate() *string
	SetGmtDelete(v string) *AppInstance
	GetGmtDelete() *string
	SetGmtModified(v string) *AppInstance
	GetGmtModified() *string
	SetGmtPublish(v string) *AppInstance
	GetGmtPublish() *string
	SetIconUrl(v string) *AppInstance
	GetIconUrl() *string
	SetName(v string) *AppInstance
	GetName() *string
	SetProfile(v *AppInstanceProfile) *AppInstance
	GetProfile() *AppInstanceProfile
	SetSiteHost(v string) *AppInstance
	GetSiteHost() *string
	SetSlug(v string) *AppInstance
	GetSlug() *string
	SetSourceType(v string) *AppInstance
	GetSourceType() *string
	SetStartTime(v string) *AppInstance
	GetStartTime() *string
	SetStatus(v string) *AppInstance
	GetStatus() *string
	SetStatusText(v string) *AppInstance
	GetStatusText() *string
	SetThumbnailUrl(v string) *AppInstance
	GetThumbnailUrl() *string
	SetUserId(v string) *AppInstance
	GetUserId() *string
}

type AppInstance struct {
	AppSubType      *string             `json:"AppSubType,omitempty" xml:"AppSubType,omitempty"`
	AppType         *string             `json:"AppType,omitempty" xml:"AppType,omitempty"`
	BizId           *string             `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BuildType       *string             `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	Deleted         *int32              `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Description     *string             `json:"Description,omitempty" xml:"Description,omitempty"`
	DesignSpecBizId *string             `json:"DesignSpecBizId,omitempty" xml:"DesignSpecBizId,omitempty"`
	DesignSpecId    *string             `json:"DesignSpecId,omitempty" xml:"DesignSpecId,omitempty"`
	Domain          *string             `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime         *string             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EspBizId        *string             `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	GmtCreate       *string             `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtDelete       *string             `json:"GmtDelete,omitempty" xml:"GmtDelete,omitempty"`
	GmtModified     *string             `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtPublish      *string             `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	IconUrl         *string             `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	Name            *string             `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile         *AppInstanceProfile `json:"Profile,omitempty" xml:"Profile,omitempty"`
	SiteHost        *string             `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	Slug            *string             `json:"Slug,omitempty" xml:"Slug,omitempty"`
	SourceType      *string             `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime       *string             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// trial,draft,live,refunded,expired,released
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusText   *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppInstance) String() string {
	return dara.Prettify(s)
}

func (s AppInstance) GoString() string {
	return s.String()
}

func (s *AppInstance) GetAppSubType() *string {
	return s.AppSubType
}

func (s *AppInstance) GetAppType() *string {
	return s.AppType
}

func (s *AppInstance) GetBizId() *string {
	return s.BizId
}

func (s *AppInstance) GetBuildType() *string {
	return s.BuildType
}

func (s *AppInstance) GetDeleted() *int32 {
	return s.Deleted
}

func (s *AppInstance) GetDescription() *string {
	return s.Description
}

func (s *AppInstance) GetDesignSpecBizId() *string {
	return s.DesignSpecBizId
}

func (s *AppInstance) GetDesignSpecId() *string {
	return s.DesignSpecId
}

func (s *AppInstance) GetDomain() *string {
	return s.Domain
}

func (s *AppInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *AppInstance) GetEspBizId() *string {
	return s.EspBizId
}

func (s *AppInstance) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AppInstance) GetGmtDelete() *string {
	return s.GmtDelete
}

func (s *AppInstance) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppInstance) GetGmtPublish() *string {
	return s.GmtPublish
}

func (s *AppInstance) GetIconUrl() *string {
	return s.IconUrl
}

func (s *AppInstance) GetName() *string {
	return s.Name
}

func (s *AppInstance) GetProfile() *AppInstanceProfile {
	return s.Profile
}

func (s *AppInstance) GetSiteHost() *string {
	return s.SiteHost
}

func (s *AppInstance) GetSlug() *string {
	return s.Slug
}

func (s *AppInstance) GetSourceType() *string {
	return s.SourceType
}

func (s *AppInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *AppInstance) GetStatus() *string {
	return s.Status
}

func (s *AppInstance) GetStatusText() *string {
	return s.StatusText
}

func (s *AppInstance) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *AppInstance) GetUserId() *string {
	return s.UserId
}

func (s *AppInstance) SetAppSubType(v string) *AppInstance {
	s.AppSubType = &v
	return s
}

func (s *AppInstance) SetAppType(v string) *AppInstance {
	s.AppType = &v
	return s
}

func (s *AppInstance) SetBizId(v string) *AppInstance {
	s.BizId = &v
	return s
}

func (s *AppInstance) SetBuildType(v string) *AppInstance {
	s.BuildType = &v
	return s
}

func (s *AppInstance) SetDeleted(v int32) *AppInstance {
	s.Deleted = &v
	return s
}

func (s *AppInstance) SetDescription(v string) *AppInstance {
	s.Description = &v
	return s
}

func (s *AppInstance) SetDesignSpecBizId(v string) *AppInstance {
	s.DesignSpecBizId = &v
	return s
}

func (s *AppInstance) SetDesignSpecId(v string) *AppInstance {
	s.DesignSpecId = &v
	return s
}

func (s *AppInstance) SetDomain(v string) *AppInstance {
	s.Domain = &v
	return s
}

func (s *AppInstance) SetEndTime(v string) *AppInstance {
	s.EndTime = &v
	return s
}

func (s *AppInstance) SetEspBizId(v string) *AppInstance {
	s.EspBizId = &v
	return s
}

func (s *AppInstance) SetGmtCreate(v string) *AppInstance {
	s.GmtCreate = &v
	return s
}

func (s *AppInstance) SetGmtDelete(v string) *AppInstance {
	s.GmtDelete = &v
	return s
}

func (s *AppInstance) SetGmtModified(v string) *AppInstance {
	s.GmtModified = &v
	return s
}

func (s *AppInstance) SetGmtPublish(v string) *AppInstance {
	s.GmtPublish = &v
	return s
}

func (s *AppInstance) SetIconUrl(v string) *AppInstance {
	s.IconUrl = &v
	return s
}

func (s *AppInstance) SetName(v string) *AppInstance {
	s.Name = &v
	return s
}

func (s *AppInstance) SetProfile(v *AppInstanceProfile) *AppInstance {
	s.Profile = v
	return s
}

func (s *AppInstance) SetSiteHost(v string) *AppInstance {
	s.SiteHost = &v
	return s
}

func (s *AppInstance) SetSlug(v string) *AppInstance {
	s.Slug = &v
	return s
}

func (s *AppInstance) SetSourceType(v string) *AppInstance {
	s.SourceType = &v
	return s
}

func (s *AppInstance) SetStartTime(v string) *AppInstance {
	s.StartTime = &v
	return s
}

func (s *AppInstance) SetStatus(v string) *AppInstance {
	s.Status = &v
	return s
}

func (s *AppInstance) SetStatusText(v string) *AppInstance {
	s.StatusText = &v
	return s
}

func (s *AppInstance) SetThumbnailUrl(v string) *AppInstance {
	s.ThumbnailUrl = &v
	return s
}

func (s *AppInstance) SetUserId(v string) *AppInstance {
	s.UserId = &v
	return s
}

func (s *AppInstance) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
