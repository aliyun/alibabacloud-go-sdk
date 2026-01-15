// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppService interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AppService
	GetBizId() *string
	SetCreateTime(v string) *AppService
	GetCreateTime() *string
	SetDeleted(v int32) *AppService
	GetDeleted() *int32
	SetEndTime(v string) *AppService
	GetEndTime() *string
	SetEspBizId(v string) *AppService
	GetEspBizId() *string
	SetGmtModified(v string) *AppService
	GetGmtModified() *string
	SetInstanceBizId(v string) *AppService
	GetInstanceBizId() *string
	SetName(v string) *AppService
	GetName() *string
	SetProfile(v *AppServiceProfile) *AppService
	GetProfile() *AppServiceProfile
	SetServiceType(v string) *AppService
	GetServiceType() *string
	SetServiceTypeText(v string) *AppService
	GetServiceTypeText() *string
	SetSlug(v string) *AppService
	GetSlug() *string
	SetStartTime(v string) *AppService
	GetStartTime() *string
	SetStatus(v string) *AppService
	GetStatus() *string
	SetUserId(v string) *AppService
	GetUserId() *string
}

type AppService struct {
	BizId           *string            `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreateTime      *string            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Deleted         *int32             `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	EndTime         *string            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EspBizId        *string            `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	GmtModified     *string            `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceBizId   *string            `json:"InstanceBizId,omitempty" xml:"InstanceBizId,omitempty"`
	Name            *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Profile         *AppServiceProfile `json:"Profile,omitempty" xml:"Profile,omitempty"`
	ServiceType     *string            `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	ServiceTypeText *string            `json:"ServiceTypeText,omitempty" xml:"ServiceTypeText,omitempty"`
	Slug            *string            `json:"Slug,omitempty" xml:"Slug,omitempty"`
	StartTime       *string            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId          *string            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppService) String() string {
	return dara.Prettify(s)
}

func (s AppService) GoString() string {
	return s.String()
}

func (s *AppService) GetBizId() *string {
	return s.BizId
}

func (s *AppService) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AppService) GetDeleted() *int32 {
	return s.Deleted
}

func (s *AppService) GetEndTime() *string {
	return s.EndTime
}

func (s *AppService) GetEspBizId() *string {
	return s.EspBizId
}

func (s *AppService) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppService) GetInstanceBizId() *string {
	return s.InstanceBizId
}

func (s *AppService) GetName() *string {
	return s.Name
}

func (s *AppService) GetProfile() *AppServiceProfile {
	return s.Profile
}

func (s *AppService) GetServiceType() *string {
	return s.ServiceType
}

func (s *AppService) GetServiceTypeText() *string {
	return s.ServiceTypeText
}

func (s *AppService) GetSlug() *string {
	return s.Slug
}

func (s *AppService) GetStartTime() *string {
	return s.StartTime
}

func (s *AppService) GetStatus() *string {
	return s.Status
}

func (s *AppService) GetUserId() *string {
	return s.UserId
}

func (s *AppService) SetBizId(v string) *AppService {
	s.BizId = &v
	return s
}

func (s *AppService) SetCreateTime(v string) *AppService {
	s.CreateTime = &v
	return s
}

func (s *AppService) SetDeleted(v int32) *AppService {
	s.Deleted = &v
	return s
}

func (s *AppService) SetEndTime(v string) *AppService {
	s.EndTime = &v
	return s
}

func (s *AppService) SetEspBizId(v string) *AppService {
	s.EspBizId = &v
	return s
}

func (s *AppService) SetGmtModified(v string) *AppService {
	s.GmtModified = &v
	return s
}

func (s *AppService) SetInstanceBizId(v string) *AppService {
	s.InstanceBizId = &v
	return s
}

func (s *AppService) SetName(v string) *AppService {
	s.Name = &v
	return s
}

func (s *AppService) SetProfile(v *AppServiceProfile) *AppService {
	s.Profile = v
	return s
}

func (s *AppService) SetServiceType(v string) *AppService {
	s.ServiceType = &v
	return s
}

func (s *AppService) SetServiceTypeText(v string) *AppService {
	s.ServiceTypeText = &v
	return s
}

func (s *AppService) SetSlug(v string) *AppService {
	s.Slug = &v
	return s
}

func (s *AppService) SetStartTime(v string) *AppService {
	s.StartTime = &v
	return s
}

func (s *AppService) SetStatus(v string) *AppService {
	s.Status = &v
	return s
}

func (s *AppService) SetUserId(v string) *AppService {
	s.UserId = &v
	return s
}

func (s *AppService) Validate() error {
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}
