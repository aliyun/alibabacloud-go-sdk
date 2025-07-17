// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedInstancesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListAuthorizedInstancesForUserResponseBodyInstances) *ListAuthorizedInstancesForUserResponseBody
	GetInstances() []*ListAuthorizedInstancesForUserResponseBodyInstances
	SetRequestId(v string) *ListAuthorizedInstancesForUserResponseBody
	GetRequestId() *string
}

type ListAuthorizedInstancesForUserResponseBody struct {
	Instances []*ListAuthorizedInstancesForUserResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// B7DB89CC-017D-5503-8953-38FFE241A618
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAuthorizedInstancesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedInstancesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedInstancesForUserResponseBody) GetInstances() []*ListAuthorizedInstancesForUserResponseBodyInstances {
	return s.Instances
}

func (s *ListAuthorizedInstancesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedInstancesForUserResponseBody) SetInstances(v []*ListAuthorizedInstancesForUserResponseBodyInstances) *ListAuthorizedInstancesForUserResponseBody {
	s.Instances = v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBody) SetRequestId(v string) *ListAuthorizedInstancesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedInstancesForUserResponseBodyInstances struct {
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// rm-2zex9lrc0gz0****.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// DMS_TEST
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// example:
	//
	// 21****
	InstanceId       *string                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PermissionDetail *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail `json:"PermissionDetail,omitempty" xml:"PermissionDetail,omitempty" type:"Struct"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// user_test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListAuthorizedInstancesForUserResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedInstancesForUserResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetDbType() *string {
	return s.DbType
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetHost() *string {
	return s.Host
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetPermissionDetail() *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail {
	return s.PermissionDetail
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetPort() *string {
	return s.Port
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) GetUserName() *string {
	return s.UserName
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetDbType(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.DbType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetEnvType(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.EnvType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetHost(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.Host = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetInstanceAlias(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.InstanceAlias = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetInstanceId(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetPermissionDetail(v *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.PermissionDetail = v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetPort(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.Port = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetUserId(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) SetUserName(v string) *ListAuthorizedInstancesForUserResponseBodyInstances {
	s.UserName = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail struct {
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

func (s ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) GoString() string {
	return s.String()
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) GetDsType() *string {
	return s.DsType
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) GetPermType() *string {
	return s.PermType
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) SetDsType(v string) *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail {
	s.DsType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) SetExpireDate(v string) *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) SetMessage(v string) *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail {
	s.Message = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) SetPermType(v string) *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail {
	s.PermType = &v
	return s
}

func (s *ListAuthorizedInstancesForUserResponseBodyInstancesPermissionDetail) Validate() error {
	return dara.Validate(s)
}
