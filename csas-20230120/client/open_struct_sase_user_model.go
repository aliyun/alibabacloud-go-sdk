// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructSaseUser interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeUnix(v int64) *OpenStructSaseUser
	GetCreateTimeUnix() *int64
	SetCustomFields(v []*IdpCustomField) *OpenStructSaseUser
	GetCustomFields() []*IdpCustomField
	SetDepartments(v []*OpenStructSaseDepartment) *OpenStructSaseUser
	GetDepartments() []*OpenStructSaseDepartment
	SetDescription(v string) *OpenStructSaseUser
	GetDescription() *string
	SetEmail(v string) *OpenStructSaseUser
	GetEmail() *string
	SetIdpConfigId(v string) *OpenStructSaseUser
	GetIdpConfigId() *string
	SetLeaveTimeUnix(v int64) *OpenStructSaseUser
	GetLeaveTimeUnix() *int64
	SetLoginTimeUnix(v int64) *OpenStructSaseUser
	GetLoginTimeUnix() *int64
	SetSaseUserId(v string) *OpenStructSaseUser
	GetSaseUserId() *string
	SetSaseUserStatus(v string) *OpenStructSaseUser
	GetSaseUserStatus() *string
	SetSyncTimeUnix(v int64) *OpenStructSaseUser
	GetSyncTimeUnix() *int64
	SetTelephone(v string) *OpenStructSaseUser
	GetTelephone() *string
	SetTitle(v string) *OpenStructSaseUser
	GetTitle() *string
	SetUpdateTimeUnix(v int64) *OpenStructSaseUser
	GetUpdateTimeUnix() *int64
	SetUsername(v string) *OpenStructSaseUser
	GetUsername() *string
	SetWorkStatus(v string) *OpenStructSaseUser
	GetWorkStatus() *string
}

type OpenStructSaseUser struct {
	CreateTimeUnix *int64                      `json:"CreateTimeUnix,omitempty" xml:"CreateTimeUnix,omitempty"`
	CustomFields   []*IdpCustomField           `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	Departments    []*OpenStructSaseDepartment `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	Description    *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Email          *string                     `json:"Email,omitempty" xml:"Email,omitempty"`
	IdpConfigId    *string                     `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	LeaveTimeUnix  *int64                      `json:"LeaveTimeUnix,omitempty" xml:"LeaveTimeUnix,omitempty"`
	LoginTimeUnix  *int64                      `json:"LoginTimeUnix,omitempty" xml:"LoginTimeUnix,omitempty"`
	SaseUserId     *string                     `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SaseUserStatus *string                     `json:"SaseUserStatus,omitempty" xml:"SaseUserStatus,omitempty"`
	SyncTimeUnix   *int64                      `json:"SyncTimeUnix,omitempty" xml:"SyncTimeUnix,omitempty"`
	Telephone      *string                     `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	Title          *string                     `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTimeUnix *int64                      `json:"UpdateTimeUnix,omitempty" xml:"UpdateTimeUnix,omitempty"`
	Username       *string                     `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkStatus     *string                     `json:"WorkStatus,omitempty" xml:"WorkStatus,omitempty"`
}

func (s OpenStructSaseUser) String() string {
	return dara.Prettify(s)
}

func (s OpenStructSaseUser) GoString() string {
	return s.String()
}

func (s *OpenStructSaseUser) GetCreateTimeUnix() *int64 {
	return s.CreateTimeUnix
}

func (s *OpenStructSaseUser) GetCustomFields() []*IdpCustomField {
	return s.CustomFields
}

func (s *OpenStructSaseUser) GetDepartments() []*OpenStructSaseDepartment {
	return s.Departments
}

func (s *OpenStructSaseUser) GetDescription() *string {
	return s.Description
}

func (s *OpenStructSaseUser) GetEmail() *string {
	return s.Email
}

func (s *OpenStructSaseUser) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *OpenStructSaseUser) GetLeaveTimeUnix() *int64 {
	return s.LeaveTimeUnix
}

func (s *OpenStructSaseUser) GetLoginTimeUnix() *int64 {
	return s.LoginTimeUnix
}

func (s *OpenStructSaseUser) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *OpenStructSaseUser) GetSaseUserStatus() *string {
	return s.SaseUserStatus
}

func (s *OpenStructSaseUser) GetSyncTimeUnix() *int64 {
	return s.SyncTimeUnix
}

func (s *OpenStructSaseUser) GetTelephone() *string {
	return s.Telephone
}

func (s *OpenStructSaseUser) GetTitle() *string {
	return s.Title
}

func (s *OpenStructSaseUser) GetUpdateTimeUnix() *int64 {
	return s.UpdateTimeUnix
}

func (s *OpenStructSaseUser) GetUsername() *string {
	return s.Username
}

func (s *OpenStructSaseUser) GetWorkStatus() *string {
	return s.WorkStatus
}

func (s *OpenStructSaseUser) SetCreateTimeUnix(v int64) *OpenStructSaseUser {
	s.CreateTimeUnix = &v
	return s
}

func (s *OpenStructSaseUser) SetCustomFields(v []*IdpCustomField) *OpenStructSaseUser {
	s.CustomFields = v
	return s
}

func (s *OpenStructSaseUser) SetDepartments(v []*OpenStructSaseDepartment) *OpenStructSaseUser {
	s.Departments = v
	return s
}

func (s *OpenStructSaseUser) SetDescription(v string) *OpenStructSaseUser {
	s.Description = &v
	return s
}

func (s *OpenStructSaseUser) SetEmail(v string) *OpenStructSaseUser {
	s.Email = &v
	return s
}

func (s *OpenStructSaseUser) SetIdpConfigId(v string) *OpenStructSaseUser {
	s.IdpConfigId = &v
	return s
}

func (s *OpenStructSaseUser) SetLeaveTimeUnix(v int64) *OpenStructSaseUser {
	s.LeaveTimeUnix = &v
	return s
}

func (s *OpenStructSaseUser) SetLoginTimeUnix(v int64) *OpenStructSaseUser {
	s.LoginTimeUnix = &v
	return s
}

func (s *OpenStructSaseUser) SetSaseUserId(v string) *OpenStructSaseUser {
	s.SaseUserId = &v
	return s
}

func (s *OpenStructSaseUser) SetSaseUserStatus(v string) *OpenStructSaseUser {
	s.SaseUserStatus = &v
	return s
}

func (s *OpenStructSaseUser) SetSyncTimeUnix(v int64) *OpenStructSaseUser {
	s.SyncTimeUnix = &v
	return s
}

func (s *OpenStructSaseUser) SetTelephone(v string) *OpenStructSaseUser {
	s.Telephone = &v
	return s
}

func (s *OpenStructSaseUser) SetTitle(v string) *OpenStructSaseUser {
	s.Title = &v
	return s
}

func (s *OpenStructSaseUser) SetUpdateTimeUnix(v int64) *OpenStructSaseUser {
	s.UpdateTimeUnix = &v
	return s
}

func (s *OpenStructSaseUser) SetUsername(v string) *OpenStructSaseUser {
	s.Username = &v
	return s
}

func (s *OpenStructSaseUser) SetWorkStatus(v string) *OpenStructSaseUser {
	s.WorkStatus = &v
	return s
}

func (s *OpenStructSaseUser) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Departments != nil {
		for _, item := range s.Departments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
