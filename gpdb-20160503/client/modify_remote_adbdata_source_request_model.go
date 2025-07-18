// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRemoteADBDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *ModifyRemoteADBDataSourceRequest
	GetDataSourceId() *string
	SetDataSourceName(v string) *ModifyRemoteADBDataSourceRequest
	GetDataSourceName() *string
	SetLocalDBInstanceId(v string) *ModifyRemoteADBDataSourceRequest
	GetLocalDBInstanceId() *string
	SetOwnerId(v int64) *ModifyRemoteADBDataSourceRequest
	GetOwnerId() *int64
	SetUserName(v string) *ModifyRemoteADBDataSourceRequest
	GetUserName() *string
	SetUserPassword(v string) *ModifyRemoteADBDataSourceRequest
	GetUserPassword() *string
}

type ModifyRemoteADBDataSourceRequest struct {
	// Service ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Specified dataSourceName.
	//
	// example:
	//
	// test
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The ID of the local data instance being used.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-test
	LocalDBInstanceId *string `json:"LocalDBInstanceId,omitempty" xml:"LocalDBInstanceId,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// New user name.
	//
	// This parameter is required.
	//
	// example:
	//
	// newUserName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// New user password, which must be transmitted in encrypted form.
	//
	// This parameter is required.
	//
	// example:
	//
	// newUserPassword
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
}

func (s ModifyRemoteADBDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRemoteADBDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyRemoteADBDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ModifyRemoteADBDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ModifyRemoteADBDataSourceRequest) GetLocalDBInstanceId() *string {
	return s.LocalDBInstanceId
}

func (s *ModifyRemoteADBDataSourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRemoteADBDataSourceRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifyRemoteADBDataSourceRequest) GetUserPassword() *string {
	return s.UserPassword
}

func (s *ModifyRemoteADBDataSourceRequest) SetDataSourceId(v string) *ModifyRemoteADBDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) SetDataSourceName(v string) *ModifyRemoteADBDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) SetLocalDBInstanceId(v string) *ModifyRemoteADBDataSourceRequest {
	s.LocalDBInstanceId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) SetOwnerId(v int64) *ModifyRemoteADBDataSourceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) SetUserName(v string) *ModifyRemoteADBDataSourceRequest {
	s.UserName = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) SetUserPassword(v string) *ModifyRemoteADBDataSourceRequest {
	s.UserPassword = &v
	return s
}

func (s *ModifyRemoteADBDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
