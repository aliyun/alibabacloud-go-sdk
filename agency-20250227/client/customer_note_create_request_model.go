// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactInformation(v string) *CustomerNoteCreateRequest
	GetContactInformation() *string
	SetContactName(v string) *CustomerNoteCreateRequest
	GetContactName() *string
	SetCustomerName(v string) *CustomerNoteCreateRequest
	GetCustomerName() *string
	SetCustomerUid(v string) *CustomerNoteCreateRequest
	GetCustomerUid() *string
	SetNoteContent(v string) *CustomerNoteCreateRequest
	GetNoteContent() *string
	SetTouchDate(v int64) *CustomerNoteCreateRequest
	GetTouchDate() *int64
}

type CustomerNoteCreateRequest struct {
	// example:
	//
	// 13833333333
	ContactInformation *string `json:"ContactInformation,omitempty" xml:"ContactInformation,omitempty"`
	// example:
	//
	// 张三
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// XXXX有限公司
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// 1647796581073291
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// example:
	//
	// 日常拜访客户，讨论客户agent建设方案
	NoteContent *string `json:"NoteContent,omitempty" xml:"NoteContent,omitempty"`
	// example:
	//
	// 1784266662000
	TouchDate *int64 `json:"TouchDate,omitempty" xml:"TouchDate,omitempty"`
}

func (s CustomerNoteCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomerNoteCreateRequest) GetContactInformation() *string {
	return s.ContactInformation
}

func (s *CustomerNoteCreateRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CustomerNoteCreateRequest) GetCustomerName() *string {
	return s.CustomerName
}

func (s *CustomerNoteCreateRequest) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *CustomerNoteCreateRequest) GetNoteContent() *string {
	return s.NoteContent
}

func (s *CustomerNoteCreateRequest) GetTouchDate() *int64 {
	return s.TouchDate
}

func (s *CustomerNoteCreateRequest) SetContactInformation(v string) *CustomerNoteCreateRequest {
	s.ContactInformation = &v
	return s
}

func (s *CustomerNoteCreateRequest) SetContactName(v string) *CustomerNoteCreateRequest {
	s.ContactName = &v
	return s
}

func (s *CustomerNoteCreateRequest) SetCustomerName(v string) *CustomerNoteCreateRequest {
	s.CustomerName = &v
	return s
}

func (s *CustomerNoteCreateRequest) SetCustomerUid(v string) *CustomerNoteCreateRequest {
	s.CustomerUid = &v
	return s
}

func (s *CustomerNoteCreateRequest) SetNoteContent(v string) *CustomerNoteCreateRequest {
	s.NoteContent = &v
	return s
}

func (s *CustomerNoteCreateRequest) SetTouchDate(v int64) *CustomerNoteCreateRequest {
	s.TouchDate = &v
	return s
}

func (s *CustomerNoteCreateRequest) Validate() error {
	return dara.Validate(s)
}
