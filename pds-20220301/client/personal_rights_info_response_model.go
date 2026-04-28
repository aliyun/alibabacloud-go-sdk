// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalRightsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresTime(v string) *PersonalRightsInfoResponse
	GetExpiresTime() *string
	SetHistoryLatestRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse
	GetHistoryLatestRights() *PersonalRightsInfoResponse
	SetIcon(v string) *PersonalRightsInfoResponse
	GetIcon() *string
	SetIsExpires(v bool) *PersonalRightsInfoResponse
	GetIsExpires() *bool
	SetName(v string) *PersonalRightsInfoResponse
	GetName() *string
	SetOtherRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse
	GetOtherRights() *PersonalRightsInfoResponse
	SetPrivileges(v []*DataBoxPrivileges) *PersonalRightsInfoResponse
	GetPrivileges() []*DataBoxPrivileges
	SetSpuId(v string) *PersonalRightsInfoResponse
	GetSpuId() *string
	SetTitle(v string) *PersonalRightsInfoResponse
	GetTitle() *string
}

type PersonalRightsInfoResponse struct {
	ExpiresTime         *string                     `json:"expires_time,omitempty" xml:"expires_time,omitempty"`
	HistoryLatestRights *PersonalRightsInfoResponse `json:"history_latest_rights,omitempty" xml:"history_latest_rights,omitempty"`
	Icon                *string                     `json:"icon,omitempty" xml:"icon,omitempty"`
	IsExpires           *bool                       `json:"is_expires,omitempty" xml:"is_expires,omitempty"`
	Name                *string                     `json:"name,omitempty" xml:"name,omitempty"`
	OtherRights         *PersonalRightsInfoResponse `json:"other_rights,omitempty" xml:"other_rights,omitempty"`
	Privileges          []*DataBoxPrivileges        `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	SpuId               *string                     `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Title               *string                     `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PersonalRightsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PersonalRightsInfoResponse) GoString() string {
	return s.String()
}

func (s *PersonalRightsInfoResponse) GetExpiresTime() *string {
	return s.ExpiresTime
}

func (s *PersonalRightsInfoResponse) GetHistoryLatestRights() *PersonalRightsInfoResponse {
	return s.HistoryLatestRights
}

func (s *PersonalRightsInfoResponse) GetIcon() *string {
	return s.Icon
}

func (s *PersonalRightsInfoResponse) GetIsExpires() *bool {
	return s.IsExpires
}

func (s *PersonalRightsInfoResponse) GetName() *string {
	return s.Name
}

func (s *PersonalRightsInfoResponse) GetOtherRights() *PersonalRightsInfoResponse {
	return s.OtherRights
}

func (s *PersonalRightsInfoResponse) GetPrivileges() []*DataBoxPrivileges {
	return s.Privileges
}

func (s *PersonalRightsInfoResponse) GetSpuId() *string {
	return s.SpuId
}

func (s *PersonalRightsInfoResponse) GetTitle() *string {
	return s.Title
}

func (s *PersonalRightsInfoResponse) SetExpiresTime(v string) *PersonalRightsInfoResponse {
	s.ExpiresTime = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetHistoryLatestRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse {
	s.HistoryLatestRights = v
	return s
}

func (s *PersonalRightsInfoResponse) SetIcon(v string) *PersonalRightsInfoResponse {
	s.Icon = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetIsExpires(v bool) *PersonalRightsInfoResponse {
	s.IsExpires = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetName(v string) *PersonalRightsInfoResponse {
	s.Name = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetOtherRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse {
	s.OtherRights = v
	return s
}

func (s *PersonalRightsInfoResponse) SetPrivileges(v []*DataBoxPrivileges) *PersonalRightsInfoResponse {
	s.Privileges = v
	return s
}

func (s *PersonalRightsInfoResponse) SetSpuId(v string) *PersonalRightsInfoResponse {
	s.SpuId = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetTitle(v string) *PersonalRightsInfoResponse {
	s.Title = &v
	return s
}

func (s *PersonalRightsInfoResponse) Validate() error {
	if s.HistoryLatestRights != nil {
		if err := s.HistoryLatestRights.Validate(); err != nil {
			return err
		}
	}
	if s.OtherRights != nil {
		if err := s.OtherRights.Validate(); err != nil {
			return err
		}
	}
	if s.Privileges != nil {
		for _, item := range s.Privileges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
