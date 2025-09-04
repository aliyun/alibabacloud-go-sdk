// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetMasterUid(v int64) *DataValue
	GetMasterUid() *int64
	SetCInstanceId(v string) *DataValue
	GetCInstanceId() *string
	SetAccessKey(v string) *DataValue
	GetAccessKey() *string
	SetUserName(v string) *DataValue
	GetUserName() *string
	SetPassword(v string) *DataValue
	GetPassword() *string
	SetDeleted(v int64) *DataValue
	GetDeleted() *int64
	SetCreateTimestamp(v int64) *DataValue
	GetCreateTimestamp() *int64
	SetRemark(v string) *DataValue
	GetRemark() *string
}

type DataValue struct {
	// The Alibaba Cloud account ID or Resource Access Management (RAM) user to which the AccessKey pair that is used to create the static username and password belongs.
	//
	// example:
	//
	// 1565************1
	MasterUid *int64 `json:"masterUid,omitempty" xml:"masterUid,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-uqm******03
	CInstanceId *string `json:"cInstanceId,omitempty" xml:"cInstanceId,omitempty"`
	// The AccessKey ID that is used to create the static username and password.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The static username.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ6cjc2djAwMzpMVEFJNX*******ZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The static password.
	//
	// example:
	//
	// OUYwQzM2QjZBRkUxNDRFM***************MzZCNzdDQzoxNjcxNDMwMzkyODI1
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The timestamp that indicates when the static username and password were deleted. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	Deleted *int64 `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// The timestamp that indicates when the static username and password were created. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// ***环境
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetMasterUid() *int64 {
	return s.MasterUid
}

func (s *DataValue) GetCInstanceId() *string {
	return s.CInstanceId
}

func (s *DataValue) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DataValue) GetUserName() *string {
	return s.UserName
}

func (s *DataValue) GetPassword() *string {
	return s.Password
}

func (s *DataValue) GetDeleted() *int64 {
	return s.Deleted
}

func (s *DataValue) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DataValue) GetRemark() *string {
	return s.Remark
}

func (s *DataValue) SetMasterUid(v int64) *DataValue {
	s.MasterUid = &v
	return s
}

func (s *DataValue) SetCInstanceId(v string) *DataValue {
	s.CInstanceId = &v
	return s
}

func (s *DataValue) SetAccessKey(v string) *DataValue {
	s.AccessKey = &v
	return s
}

func (s *DataValue) SetUserName(v string) *DataValue {
	s.UserName = &v
	return s
}

func (s *DataValue) SetPassword(v string) *DataValue {
	s.Password = &v
	return s
}

func (s *DataValue) SetDeleted(v int64) *DataValue {
	s.Deleted = &v
	return s
}

func (s *DataValue) SetCreateTimestamp(v int64) *DataValue {
	s.CreateTimestamp = &v
	return s
}

func (s *DataValue) SetRemark(v string) *DataValue {
	s.Remark = &v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
