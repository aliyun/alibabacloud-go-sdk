// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppInstanceAggregate interface {
	dara.Model
	String() string
	GoString() string
	SetAiStaffList(v []*AppAiStaff) *AppInstanceAggregate
	GetAiStaffList() []*AppAiStaff
	SetAppOperationAddress(v *AppOperationAddress) *AppInstanceAggregate
	GetAppOperationAddress() *AppOperationAddress
	SetAppServiceList(v []*AppServiceAggregate) *AppInstanceAggregate
	GetAppServiceList() []*AppServiceAggregate
	SetAppSubType(v string) *AppInstanceAggregate
	GetAppSubType() *string
	SetAppType(v string) *AppInstanceAggregate
	GetAppType() *string
	SetBizId(v string) *AppInstanceAggregate
	GetBizId() *string
	SetBuildType(v string) *AppInstanceAggregate
	GetBuildType() *string
	SetCreateTime(v string) *AppInstanceAggregate
	GetCreateTime() *string
	SetDeleted(v int32) *AppInstanceAggregate
	GetDeleted() *int32
	SetDescription(v string) *AppInstanceAggregate
	GetDescription() *string
	SetDesignSpecBizId(v string) *AppInstanceAggregate
	GetDesignSpecBizId() *string
	SetDesignSpecId(v string) *AppInstanceAggregate
	GetDesignSpecId() *string
	SetDomain(v string) *AppInstanceAggregate
	GetDomain() *string
	SetEndTime(v string) *AppInstanceAggregate
	GetEndTime() *string
	SetEspBizId(v string) *AppInstanceAggregate
	GetEspBizId() *string
	SetGmtDelete(v string) *AppInstanceAggregate
	GetGmtDelete() *string
	SetGmtModified(v string) *AppInstanceAggregate
	GetGmtModified() *string
	SetGmtPublish(v string) *AppInstanceAggregate
	GetGmtPublish() *string
	SetIconUrl(v string) *AppInstanceAggregate
	GetIconUrl() *string
	SetName(v string) *AppInstanceAggregate
	GetName() *string
	SetProfile(v *AppInstanceProfile) *AppInstanceAggregate
	GetProfile() *AppInstanceProfile
	SetSiteHost(v string) *AppInstanceAggregate
	GetSiteHost() *string
	SetSlug(v string) *AppInstanceAggregate
	GetSlug() *string
	SetSourceType(v string) *AppInstanceAggregate
	GetSourceType() *string
	SetStartTime(v string) *AppInstanceAggregate
	GetStartTime() *string
	SetStatus(v string) *AppInstanceAggregate
	GetStatus() *string
	SetStatusText(v string) *AppInstanceAggregate
	GetStatusText() *string
	SetThumbnailUrl(v string) *AppInstanceAggregate
	GetThumbnailUrl() *string
	SetUserId(v string) *AppInstanceAggregate
	GetUserId() *string
}

type AppInstanceAggregate struct {
	AiStaffList         []*AppAiStaff          `json:"AiStaffList,omitempty" xml:"AiStaffList,omitempty" type:"Repeated"`
	AppOperationAddress *AppOperationAddress   `json:"AppOperationAddress,omitempty" xml:"AppOperationAddress,omitempty"`
	AppServiceList      []*AppServiceAggregate `json:"AppServiceList,omitempty" xml:"AppServiceList,omitempty" type:"Repeated"`
	AppSubType          *string                `json:"AppSubType,omitempty" xml:"AppSubType,omitempty"`
	AppType             *string                `json:"AppType,omitempty" xml:"AppType,omitempty"`
	BizId               *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BuildType           *string                `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	CreateTime          *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted             *int32                 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Description         *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	DesignSpecBizId     *string                `json:"DesignSpecBizId,omitempty" xml:"DesignSpecBizId,omitempty"`
	DesignSpecId        *string                `json:"DesignSpecId,omitempty" xml:"DesignSpecId,omitempty"`
	Domain              *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime             *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EspBizId            *string                `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	GmtDelete           *string                `json:"GmtDelete,omitempty" xml:"GmtDelete,omitempty"`
	GmtModified         *string                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtPublish          *string                `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	IconUrl             *string                `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	Name                *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile             *AppInstanceProfile    `json:"Profile,omitempty" xml:"Profile,omitempty"`
	SiteHost            *string                `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
	Slug                *string                `json:"Slug,omitempty" xml:"Slug,omitempty"`
	SourceType          *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime           *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// trial,draft,live,refunded,expired,released
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusText   *string `json:"StatusText,omitempty" xml:"StatusText,omitempty"`
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppInstanceAggregate) String() string {
	return dara.Prettify(s)
}

func (s AppInstanceAggregate) GoString() string {
	return s.String()
}

func (s *AppInstanceAggregate) GetAiStaffList() []*AppAiStaff {
	return s.AiStaffList
}

func (s *AppInstanceAggregate) GetAppOperationAddress() *AppOperationAddress {
	return s.AppOperationAddress
}

func (s *AppInstanceAggregate) GetAppServiceList() []*AppServiceAggregate {
	return s.AppServiceList
}

func (s *AppInstanceAggregate) GetAppSubType() *string {
	return s.AppSubType
}

func (s *AppInstanceAggregate) GetAppType() *string {
	return s.AppType
}

func (s *AppInstanceAggregate) GetBizId() *string {
	return s.BizId
}

func (s *AppInstanceAggregate) GetBuildType() *string {
	return s.BuildType
}

func (s *AppInstanceAggregate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AppInstanceAggregate) GetDeleted() *int32 {
	return s.Deleted
}

func (s *AppInstanceAggregate) GetDescription() *string {
	return s.Description
}

func (s *AppInstanceAggregate) GetDesignSpecBizId() *string {
	return s.DesignSpecBizId
}

func (s *AppInstanceAggregate) GetDesignSpecId() *string {
	return s.DesignSpecId
}

func (s *AppInstanceAggregate) GetDomain() *string {
	return s.Domain
}

func (s *AppInstanceAggregate) GetEndTime() *string {
	return s.EndTime
}

func (s *AppInstanceAggregate) GetEspBizId() *string {
	return s.EspBizId
}

func (s *AppInstanceAggregate) GetGmtDelete() *string {
	return s.GmtDelete
}

func (s *AppInstanceAggregate) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppInstanceAggregate) GetGmtPublish() *string {
	return s.GmtPublish
}

func (s *AppInstanceAggregate) GetIconUrl() *string {
	return s.IconUrl
}

func (s *AppInstanceAggregate) GetName() *string {
	return s.Name
}

func (s *AppInstanceAggregate) GetProfile() *AppInstanceProfile {
	return s.Profile
}

func (s *AppInstanceAggregate) GetSiteHost() *string {
	return s.SiteHost
}

func (s *AppInstanceAggregate) GetSlug() *string {
	return s.Slug
}

func (s *AppInstanceAggregate) GetSourceType() *string {
	return s.SourceType
}

func (s *AppInstanceAggregate) GetStartTime() *string {
	return s.StartTime
}

func (s *AppInstanceAggregate) GetStatus() *string {
	return s.Status
}

func (s *AppInstanceAggregate) GetStatusText() *string {
	return s.StatusText
}

func (s *AppInstanceAggregate) GetThumbnailUrl() *string {
	return s.ThumbnailUrl
}

func (s *AppInstanceAggregate) GetUserId() *string {
	return s.UserId
}

func (s *AppInstanceAggregate) SetAiStaffList(v []*AppAiStaff) *AppInstanceAggregate {
	s.AiStaffList = v
	return s
}

func (s *AppInstanceAggregate) SetAppOperationAddress(v *AppOperationAddress) *AppInstanceAggregate {
	s.AppOperationAddress = v
	return s
}

func (s *AppInstanceAggregate) SetAppServiceList(v []*AppServiceAggregate) *AppInstanceAggregate {
	s.AppServiceList = v
	return s
}

func (s *AppInstanceAggregate) SetAppSubType(v string) *AppInstanceAggregate {
	s.AppSubType = &v
	return s
}

func (s *AppInstanceAggregate) SetAppType(v string) *AppInstanceAggregate {
	s.AppType = &v
	return s
}

func (s *AppInstanceAggregate) SetBizId(v string) *AppInstanceAggregate {
	s.BizId = &v
	return s
}

func (s *AppInstanceAggregate) SetBuildType(v string) *AppInstanceAggregate {
	s.BuildType = &v
	return s
}

func (s *AppInstanceAggregate) SetCreateTime(v string) *AppInstanceAggregate {
	s.CreateTime = &v
	return s
}

func (s *AppInstanceAggregate) SetDeleted(v int32) *AppInstanceAggregate {
	s.Deleted = &v
	return s
}

func (s *AppInstanceAggregate) SetDescription(v string) *AppInstanceAggregate {
	s.Description = &v
	return s
}

func (s *AppInstanceAggregate) SetDesignSpecBizId(v string) *AppInstanceAggregate {
	s.DesignSpecBizId = &v
	return s
}

func (s *AppInstanceAggregate) SetDesignSpecId(v string) *AppInstanceAggregate {
	s.DesignSpecId = &v
	return s
}

func (s *AppInstanceAggregate) SetDomain(v string) *AppInstanceAggregate {
	s.Domain = &v
	return s
}

func (s *AppInstanceAggregate) SetEndTime(v string) *AppInstanceAggregate {
	s.EndTime = &v
	return s
}

func (s *AppInstanceAggregate) SetEspBizId(v string) *AppInstanceAggregate {
	s.EspBizId = &v
	return s
}

func (s *AppInstanceAggregate) SetGmtDelete(v string) *AppInstanceAggregate {
	s.GmtDelete = &v
	return s
}

func (s *AppInstanceAggregate) SetGmtModified(v string) *AppInstanceAggregate {
	s.GmtModified = &v
	return s
}

func (s *AppInstanceAggregate) SetGmtPublish(v string) *AppInstanceAggregate {
	s.GmtPublish = &v
	return s
}

func (s *AppInstanceAggregate) SetIconUrl(v string) *AppInstanceAggregate {
	s.IconUrl = &v
	return s
}

func (s *AppInstanceAggregate) SetName(v string) *AppInstanceAggregate {
	s.Name = &v
	return s
}

func (s *AppInstanceAggregate) SetProfile(v *AppInstanceProfile) *AppInstanceAggregate {
	s.Profile = v
	return s
}

func (s *AppInstanceAggregate) SetSiteHost(v string) *AppInstanceAggregate {
	s.SiteHost = &v
	return s
}

func (s *AppInstanceAggregate) SetSlug(v string) *AppInstanceAggregate {
	s.Slug = &v
	return s
}

func (s *AppInstanceAggregate) SetSourceType(v string) *AppInstanceAggregate {
	s.SourceType = &v
	return s
}

func (s *AppInstanceAggregate) SetStartTime(v string) *AppInstanceAggregate {
	s.StartTime = &v
	return s
}

func (s *AppInstanceAggregate) SetStatus(v string) *AppInstanceAggregate {
	s.Status = &v
	return s
}

func (s *AppInstanceAggregate) SetStatusText(v string) *AppInstanceAggregate {
	s.StatusText = &v
	return s
}

func (s *AppInstanceAggregate) SetThumbnailUrl(v string) *AppInstanceAggregate {
	s.ThumbnailUrl = &v
	return s
}

func (s *AppInstanceAggregate) SetUserId(v string) *AppInstanceAggregate {
	s.UserId = &v
	return s
}

func (s *AppInstanceAggregate) Validate() error {
	if s.AiStaffList != nil {
		for _, item := range s.AiStaffList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AppOperationAddress != nil {
		if err := s.AppOperationAddress.Validate(); err != nil {
			return err
		}
	}
	if s.AppServiceList != nil {
		for _, item := range s.AppServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
