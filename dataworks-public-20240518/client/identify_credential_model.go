// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdentifyCredential interface {
	dara.Model
	String() string
	GoString() string
	SetDataSource(v *IdentifyCredentialDataSource) *IdentifyCredential
	GetDataSource() *IdentifyCredentialDataSource
	SetProjectId(v string) *IdentifyCredential
	GetProjectId() *string
	SetUserId(v string) *IdentifyCredential
	GetUserId() *string
	SetUserType(v string) *IdentifyCredential
	GetUserType() *string
}

type IdentifyCredential struct {
	// The data source.
	DataSource *IdentifyCredentialDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The project ID. This parameter is optional.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The user UID. If the `UserType` is `Role`, you must prefix this value with `ROLE_`.
	//
	// example:
	//
	// ROLE_300888674340307309
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user type. Supported values are:
	//
	// - primary account: `PrimaryAccount`
	//
	// - subaccount: `SubAccount`
	//
	// - role: `Role`
	//
	// example:
	//
	// primaryAccount
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s IdentifyCredential) String() string {
	return dara.Prettify(s)
}

func (s IdentifyCredential) GoString() string {
	return s.String()
}

func (s *IdentifyCredential) GetDataSource() *IdentifyCredentialDataSource {
	return s.DataSource
}

func (s *IdentifyCredential) GetProjectId() *string {
	return s.ProjectId
}

func (s *IdentifyCredential) GetUserId() *string {
	return s.UserId
}

func (s *IdentifyCredential) GetUserType() *string {
	return s.UserType
}

func (s *IdentifyCredential) SetDataSource(v *IdentifyCredentialDataSource) *IdentifyCredential {
	s.DataSource = v
	return s
}

func (s *IdentifyCredential) SetProjectId(v string) *IdentifyCredential {
	s.ProjectId = &v
	return s
}

func (s *IdentifyCredential) SetUserId(v string) *IdentifyCredential {
	s.UserId = &v
	return s
}

func (s *IdentifyCredential) SetUserType(v string) *IdentifyCredential {
	s.UserType = &v
	return s
}

func (s *IdentifyCredential) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IdentifyCredentialDataSource struct {
	// The data source instance ID.
	//
	// example:
	//
	// 710007423244
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The data source instance name.
	//
	// example:
	//
	// rm-2zez82ho69yex7s7g
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The data source password.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The user role for the data source. Valid values are:
	//
	// - `Admin`
	//
	// - `RegularUser`
	//
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The data source type. Supported values are:
	//
	// - `hive`
	//
	// - `lindorm_for_engine`
	//
	// - `starrocks`
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The data source user name.
	//
	// example:
	//
	// tom
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s IdentifyCredentialDataSource) String() string {
	return dara.Prettify(s)
}

func (s IdentifyCredentialDataSource) GoString() string {
	return s.String()
}

func (s *IdentifyCredentialDataSource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *IdentifyCredentialDataSource) GetInstanceName() *string {
	return s.InstanceName
}

func (s *IdentifyCredentialDataSource) GetPassword() *string {
	return s.Password
}

func (s *IdentifyCredentialDataSource) GetRole() *string {
	return s.Role
}

func (s *IdentifyCredentialDataSource) GetType() *string {
	return s.Type
}

func (s *IdentifyCredentialDataSource) GetUserName() *string {
	return s.UserName
}

func (s *IdentifyCredentialDataSource) SetInstanceId(v string) *IdentifyCredentialDataSource {
	s.InstanceId = &v
	return s
}

func (s *IdentifyCredentialDataSource) SetInstanceName(v string) *IdentifyCredentialDataSource {
	s.InstanceName = &v
	return s
}

func (s *IdentifyCredentialDataSource) SetPassword(v string) *IdentifyCredentialDataSource {
	s.Password = &v
	return s
}

func (s *IdentifyCredentialDataSource) SetRole(v string) *IdentifyCredentialDataSource {
	s.Role = &v
	return s
}

func (s *IdentifyCredentialDataSource) SetType(v string) *IdentifyCredentialDataSource {
	s.Type = &v
	return s
}

func (s *IdentifyCredentialDataSource) SetUserName(v string) *IdentifyCredentialDataSource {
	s.UserName = &v
	return s
}

func (s *IdentifyCredentialDataSource) Validate() error {
	return dara.Validate(s)
}
