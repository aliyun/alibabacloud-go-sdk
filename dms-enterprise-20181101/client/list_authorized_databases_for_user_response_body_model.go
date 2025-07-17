// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDatabasesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*ListAuthorizedDatabasesForUserResponseBodyDatabases) *ListAuthorizedDatabasesForUserResponseBody
	GetDatabases() []*ListAuthorizedDatabasesForUserResponseBodyDatabases
	SetRequestId(v string) *ListAuthorizedDatabasesForUserResponseBody
	GetRequestId() *string
}

type ListAuthorizedDatabasesForUserResponseBody struct {
	Databases []*ListAuthorizedDatabasesForUserResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// example:
	//
	// 012AE0B5-4B52-532F-BD7C-1EE9F182089B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAuthorizedDatabasesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDatabasesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDatabasesForUserResponseBody) GetDatabases() []*ListAuthorizedDatabasesForUserResponseBodyDatabases {
	return s.Databases
}

func (s *ListAuthorizedDatabasesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedDatabasesForUserResponseBody) SetDatabases(v []*ListAuthorizedDatabasesForUserResponseBodyDatabases) *ListAuthorizedDatabasesForUserResponseBody {
	s.Databases = v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBody) SetRequestId(v string) *ListAuthorizedDatabasesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedDatabasesForUserResponseBodyDatabases struct {
	// example:
	//
	// 254****
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// MYSQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 235****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	Logic            *bool                                                                `json:"Logic,omitempty" xml:"Logic,omitempty"`
	PermissionDetail *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail `json:"PermissionDetail,omitempty" xml:"PermissionDetail,omitempty" type:"Struct"`
	// example:
	//
	// poc_testdb
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// poc
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthorizedDatabasesForUserResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDatabasesForUserResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetDbId() *string {
	return s.DbId
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetDbType() *string {
	return s.DbType
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetLogic() *bool {
	return s.Logic
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetPermissionDetail() *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail {
	return s.PermissionDetail
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetSearchName() *string {
	return s.SearchName
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetDbId(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.DbId = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetDbType(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.DbType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetEnvType(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.EnvType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetInstanceId(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetLogic(v bool) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.Logic = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetPermissionDetail(v *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.PermissionDetail = v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetSchemaName(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.SchemaName = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetSearchName(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.SearchName = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) SetUserId(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabases {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail struct {
	// example:
	//
	// DATABASE
	DsType *string `json:"DsType,omitempty" xml:"DsType,omitempty"`
	// example:
	//
	// 2024-12-06 10:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CORRECT
	PermType *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
}

func (s ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) GetDsType() *string {
	return s.DsType
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) GetPermType() *string {
	return s.PermType
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) SetDsType(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail {
	s.DsType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) SetExpireDate(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) SetMessage(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail {
	s.Message = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) SetPermType(v string) *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail {
	s.PermType = &v
	return s
}

func (s *ListAuthorizedDatabasesForUserResponseBodyDatabasesPermissionDetail) Validate() error {
	return dara.Validate(s)
}
