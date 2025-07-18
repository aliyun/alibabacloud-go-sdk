// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemoteADBDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceName(v string) *CreateRemoteADBDataSourceRequest
	GetDataSourceName() *string
	SetLocalDBInstanceId(v string) *CreateRemoteADBDataSourceRequest
	GetLocalDBInstanceId() *string
	SetLocalDatabase(v string) *CreateRemoteADBDataSourceRequest
	GetLocalDatabase() *string
	SetManagerUserName(v string) *CreateRemoteADBDataSourceRequest
	GetManagerUserName() *string
	SetManagerUserPassword(v string) *CreateRemoteADBDataSourceRequest
	GetManagerUserPassword() *string
	SetOwnerId(v int64) *CreateRemoteADBDataSourceRequest
	GetOwnerId() *int64
	SetRemoteDBInstanceId(v string) *CreateRemoteADBDataSourceRequest
	GetRemoteDBInstanceId() *string
	SetRemoteDatabase(v string) *CreateRemoteADBDataSourceRequest
	GetRemoteDatabase() *string
	SetUserName(v string) *CreateRemoteADBDataSourceRequest
	GetUserName() *string
	SetUserPassword(v string) *CreateRemoteADBDataSourceRequest
	GetUserPassword() *string
}

type CreateRemoteADBDataSourceRequest struct {
	// Customer-specified DataSourceName.
	//
	// example:
	//
	// test
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Instance ID of the data being used (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-test1
	LocalDBInstanceId *string `json:"LocalDBInstanceId,omitempty" xml:"LocalDBInstanceId,omitempty"`
	// Database name of the data being used (required)
	//
	// This parameter is required.
	//
	// example:
	//
	// db1
	LocalDatabase *string `json:"LocalDatabase,omitempty" xml:"LocalDatabase,omitempty"`
	// Management account of the data-using instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// managerAccount
	ManagerUserName *string `json:"ManagerUserName,omitempty" xml:"ManagerUserName,omitempty"`
	// Password of the management account of the data-using instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// password2
	ManagerUserPassword *string `json:"ManagerUserPassword,omitempty" xml:"ManagerUserPassword,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Instance ID providing the data (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-test2
	RemoteDBInstanceId *string `json:"RemoteDBInstanceId,omitempty" xml:"RemoteDBInstanceId,omitempty"`
	// Database name providing the data (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// db2
	RemoteDatabase *string `json:"RemoteDatabase,omitempty" xml:"RemoteDatabase,omitempty"`
	// Account name of the data-providing instance used for user mapping (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Password of the data-providing instance account used for user mapping.
	//
	// This parameter is required.
	//
	// example:
	//
	// password1
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
}

func (s CreateRemoteADBDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRemoteADBDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateRemoteADBDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateRemoteADBDataSourceRequest) GetLocalDBInstanceId() *string {
	return s.LocalDBInstanceId
}

func (s *CreateRemoteADBDataSourceRequest) GetLocalDatabase() *string {
	return s.LocalDatabase
}

func (s *CreateRemoteADBDataSourceRequest) GetManagerUserName() *string {
	return s.ManagerUserName
}

func (s *CreateRemoteADBDataSourceRequest) GetManagerUserPassword() *string {
	return s.ManagerUserPassword
}

func (s *CreateRemoteADBDataSourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRemoteADBDataSourceRequest) GetRemoteDBInstanceId() *string {
	return s.RemoteDBInstanceId
}

func (s *CreateRemoteADBDataSourceRequest) GetRemoteDatabase() *string {
	return s.RemoteDatabase
}

func (s *CreateRemoteADBDataSourceRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateRemoteADBDataSourceRequest) GetUserPassword() *string {
	return s.UserPassword
}

func (s *CreateRemoteADBDataSourceRequest) SetDataSourceName(v string) *CreateRemoteADBDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetLocalDBInstanceId(v string) *CreateRemoteADBDataSourceRequest {
	s.LocalDBInstanceId = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetLocalDatabase(v string) *CreateRemoteADBDataSourceRequest {
	s.LocalDatabase = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetManagerUserName(v string) *CreateRemoteADBDataSourceRequest {
	s.ManagerUserName = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetManagerUserPassword(v string) *CreateRemoteADBDataSourceRequest {
	s.ManagerUserPassword = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetOwnerId(v int64) *CreateRemoteADBDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetRemoteDBInstanceId(v string) *CreateRemoteADBDataSourceRequest {
	s.RemoteDBInstanceId = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetRemoteDatabase(v string) *CreateRemoteADBDataSourceRequest {
	s.RemoteDatabase = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetUserName(v string) *CreateRemoteADBDataSourceRequest {
	s.UserName = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) SetUserPassword(v string) *CreateRemoteADBDataSourceRequest {
	s.UserPassword = &v
	return s
}

func (s *CreateRemoteADBDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
