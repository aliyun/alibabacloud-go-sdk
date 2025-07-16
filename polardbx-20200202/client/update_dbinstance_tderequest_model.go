// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpdateDBInstanceTDERequest
	GetDBInstanceName() *string
	SetEncryptionKey(v string) *UpdateDBInstanceTDERequest
	GetEncryptionKey() *string
	SetRegionId(v string) *UpdateDBInstanceTDERequest
	GetRegionId() *string
	SetRoleArn(v string) *UpdateDBInstanceTDERequest
	GetRoleArn() *string
	SetTDEStatus(v int32) *UpdateDBInstanceTDERequest
	GetTDEStatus() *int32
}

type UpdateDBInstanceTDERequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// RkVBNURDMjAtNkQ4QS01OTc5LTk3QUEtRkM1NzU0Nk******
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acs:ram::1406926****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TDEStatus *int32 `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s UpdateDBInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDERequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdateDBInstanceTDERequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *UpdateDBInstanceTDERequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDBInstanceTDERequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateDBInstanceTDERequest) GetTDEStatus() *int32 {
	return s.TDEStatus
}

func (s *UpdateDBInstanceTDERequest) SetDBInstanceName(v string) *UpdateDBInstanceTDERequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetEncryptionKey(v string) *UpdateDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetRegionId(v string) *UpdateDBInstanceTDERequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetRoleArn(v string) *UpdateDBInstanceTDERequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) SetTDEStatus(v int32) *UpdateDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *UpdateDBInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
