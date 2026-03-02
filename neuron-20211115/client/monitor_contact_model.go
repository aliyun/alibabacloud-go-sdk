// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorContact interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *MonitorContact
	GetEmail() *string
	SetEnterpriseId(v int64) *MonitorContact
	GetEnterpriseId() *int64
	SetGmtCreate(v string) *MonitorContact
	GetGmtCreate() *string
	SetGmtModified(v string) *MonitorContact
	GetGmtModified() *string
	SetId(v int64) *MonitorContact
	GetId() *int64
	SetIsVerify(v int32) *MonitorContact
	GetIsVerify() *int32
	SetName(v string) *MonitorContact
	GetName() *string
	SetPhone(v string) *MonitorContact
	GetPhone() *string
}

type MonitorContact struct {
	// This parameter is required.
	//
	// example:
	//
	// al2i@sina.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2022-05-01T00:00:00.000Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	IsVerify *int32 `json:"isVerify,omitempty" xml:"isVerify,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yani
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15117923456
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s MonitorContact) String() string {
	return dara.Prettify(s)
}

func (s MonitorContact) GoString() string {
	return s.String()
}

func (s *MonitorContact) GetEmail() *string {
	return s.Email
}

func (s *MonitorContact) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *MonitorContact) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MonitorContact) GetGmtModified() *string {
	return s.GmtModified
}

func (s *MonitorContact) GetId() *int64 {
	return s.Id
}

func (s *MonitorContact) GetIsVerify() *int32 {
	return s.IsVerify
}

func (s *MonitorContact) GetName() *string {
	return s.Name
}

func (s *MonitorContact) GetPhone() *string {
	return s.Phone
}

func (s *MonitorContact) SetEmail(v string) *MonitorContact {
	s.Email = &v
	return s
}

func (s *MonitorContact) SetEnterpriseId(v int64) *MonitorContact {
	s.EnterpriseId = &v
	return s
}

func (s *MonitorContact) SetGmtCreate(v string) *MonitorContact {
	s.GmtCreate = &v
	return s
}

func (s *MonitorContact) SetGmtModified(v string) *MonitorContact {
	s.GmtModified = &v
	return s
}

func (s *MonitorContact) SetId(v int64) *MonitorContact {
	s.Id = &v
	return s
}

func (s *MonitorContact) SetIsVerify(v int32) *MonitorContact {
	s.IsVerify = &v
	return s
}

func (s *MonitorContact) SetName(v string) *MonitorContact {
	s.Name = &v
	return s
}

func (s *MonitorContact) SetPhone(v string) *MonitorContact {
	s.Phone = &v
	return s
}

func (s *MonitorContact) Validate() error {
	return dara.Validate(s)
}
