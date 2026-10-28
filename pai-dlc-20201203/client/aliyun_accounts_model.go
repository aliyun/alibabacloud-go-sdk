// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunAccounts interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *AliyunAccounts
	GetAliyunUid() *string
	SetEmployeeId(v string) *AliyunAccounts
	GetEmployeeId() *string
	SetGmtCreateTime(v string) *AliyunAccounts
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *AliyunAccounts
	GetGmtModifyTime() *string
}

type AliyunAccounts struct {
	// The UID of the Alibaba Cloud account.
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The employee ID of the internal user.
	EmployeeId *string `json:"EmployeeId,omitempty" xml:"EmployeeId,omitempty"`
	// The creation time of the record.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time of the record.
	//
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s AliyunAccounts) String() string {
	return dara.Prettify(s)
}

func (s AliyunAccounts) GoString() string {
	return s.String()
}

func (s *AliyunAccounts) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *AliyunAccounts) GetEmployeeId() *string {
	return s.EmployeeId
}

func (s *AliyunAccounts) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *AliyunAccounts) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *AliyunAccounts) SetAliyunUid(v string) *AliyunAccounts {
	s.AliyunUid = &v
	return s
}

func (s *AliyunAccounts) SetEmployeeId(v string) *AliyunAccounts {
	s.EmployeeId = &v
	return s
}

func (s *AliyunAccounts) SetGmtCreateTime(v string) *AliyunAccounts {
	s.GmtCreateTime = &v
	return s
}

func (s *AliyunAccounts) SetGmtModifyTime(v string) *AliyunAccounts {
	s.GmtModifyTime = &v
	return s
}

func (s *AliyunAccounts) Validate() error {
	return dara.Validate(s)
}
