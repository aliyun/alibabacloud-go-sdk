// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceBindedEndUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNo(v string) *UpdateDeviceBindedEndUserRequest
	GetSerialNo() *string
	SetSourceAdEndUsers(v []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers) *UpdateDeviceBindedEndUserRequest
	GetSourceAdEndUsers() []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers
	SetSourceEndUserIds(v string) *UpdateDeviceBindedEndUserRequest
	GetSourceEndUserIds() *string
	SetTargetAdEndUsers(v []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers) *UpdateDeviceBindedEndUserRequest
	GetTargetAdEndUsers() []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers
	SetTargetEndUserIds(v string) *UpdateDeviceBindedEndUserRequest
	GetTargetEndUserIds() *string
	SetUserType(v string) *UpdateDeviceBindedEndUserRequest
	GetUserType() *string
}

type UpdateDeviceBindedEndUserRequest struct {
	// This parameter is required.
	SerialNo         *string                                             `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SourceAdEndUsers []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers `json:"SourceAdEndUsers,omitempty" xml:"SourceAdEndUsers,omitempty" type:"Repeated"`
	SourceEndUserIds *string                                             `json:"SourceEndUserIds,omitempty" xml:"SourceEndUserIds,omitempty"`
	TargetAdEndUsers []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers `json:"TargetAdEndUsers,omitempty" xml:"TargetAdEndUsers,omitempty" type:"Repeated"`
	TargetEndUserIds *string                                             `json:"TargetEndUserIds,omitempty" xml:"TargetEndUserIds,omitempty"`
	UserType         *string                                             `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *UpdateDeviceBindedEndUserRequest) GetSourceAdEndUsers() []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	return s.SourceAdEndUsers
}

func (s *UpdateDeviceBindedEndUserRequest) GetSourceEndUserIds() *string {
	return s.SourceEndUserIds
}

func (s *UpdateDeviceBindedEndUserRequest) GetTargetAdEndUsers() []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	return s.TargetAdEndUsers
}

func (s *UpdateDeviceBindedEndUserRequest) GetTargetEndUserIds() *string {
	return s.TargetEndUserIds
}

func (s *UpdateDeviceBindedEndUserRequest) GetUserType() *string {
	return s.UserType
}

func (s *UpdateDeviceBindedEndUserRequest) SetSerialNo(v string) *UpdateDeviceBindedEndUserRequest {
	s.SerialNo = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetSourceAdEndUsers(v []*UpdateDeviceBindedEndUserRequestSourceAdEndUsers) *UpdateDeviceBindedEndUserRequest {
	s.SourceAdEndUsers = v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetSourceEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.SourceEndUserIds = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetTargetAdEndUsers(v []*UpdateDeviceBindedEndUserRequestTargetAdEndUsers) *UpdateDeviceBindedEndUserRequest {
	s.TargetAdEndUsers = v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetTargetEndUserIds(v string) *UpdateDeviceBindedEndUserRequest {
	s.TargetEndUserIds = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) SetUserType(v string) *UpdateDeviceBindedEndUserRequest {
	s.UserType = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDeviceBindedEndUserRequestSourceAdEndUsers struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequestSourceAdEndUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequestSourceAdEndUsers) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) GetAdDomain() *string {
	return s.AdDomain
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetAdDomain(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.AdDomain = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetDirectoryId(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) SetEndUserId(v string) *UpdateDeviceBindedEndUserRequestSourceAdEndUsers {
	s.EndUserId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestSourceAdEndUsers) Validate() error {
	return dara.Validate(s)
}

type UpdateDeviceBindedEndUserRequestTargetAdEndUsers struct {
	AdDomain    *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId   *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
}

func (s UpdateDeviceBindedEndUserRequestTargetAdEndUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceBindedEndUserRequestTargetAdEndUsers) GoString() string {
	return s.String()
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) GetAdDomain() *string {
	return s.AdDomain
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetAdDomain(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.AdDomain = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetDirectoryId(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) SetEndUserId(v string) *UpdateDeviceBindedEndUserRequestTargetAdEndUsers {
	s.EndUserId = &v
	return s
}

func (s *UpdateDeviceBindedEndUserRequestTargetAdEndUsers) Validate() error {
	return dara.Validate(s)
}
