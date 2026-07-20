// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteEditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactInformation(v string) *CustomerNoteEditRequest
	GetContactInformation() *string
	SetContactName(v string) *CustomerNoteEditRequest
	GetContactName() *string
	SetNoteContent(v string) *CustomerNoteEditRequest
	GetNoteContent() *string
	SetNoteId(v int64) *CustomerNoteEditRequest
	GetNoteId() *int64
	SetTouchDate(v int64) *CustomerNoteEditRequest
	GetTouchDate() *int64
}

type CustomerNoteEditRequest struct {
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
	// 日常拜访客户，讨论客户agent建设方案
	NoteContent *string `json:"NoteContent,omitempty" xml:"NoteContent,omitempty"`
	// example:
	//
	// 1629862
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// example:
	//
	// 1784266662000
	TouchDate *int64 `json:"TouchDate,omitempty" xml:"TouchDate,omitempty"`
}

func (s CustomerNoteEditRequest) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteEditRequest) GoString() string {
	return s.String()
}

func (s *CustomerNoteEditRequest) GetContactInformation() *string {
	return s.ContactInformation
}

func (s *CustomerNoteEditRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CustomerNoteEditRequest) GetNoteContent() *string {
	return s.NoteContent
}

func (s *CustomerNoteEditRequest) GetNoteId() *int64 {
	return s.NoteId
}

func (s *CustomerNoteEditRequest) GetTouchDate() *int64 {
	return s.TouchDate
}

func (s *CustomerNoteEditRequest) SetContactInformation(v string) *CustomerNoteEditRequest {
	s.ContactInformation = &v
	return s
}

func (s *CustomerNoteEditRequest) SetContactName(v string) *CustomerNoteEditRequest {
	s.ContactName = &v
	return s
}

func (s *CustomerNoteEditRequest) SetNoteContent(v string) *CustomerNoteEditRequest {
	s.NoteContent = &v
	return s
}

func (s *CustomerNoteEditRequest) SetNoteId(v int64) *CustomerNoteEditRequest {
	s.NoteId = &v
	return s
}

func (s *CustomerNoteEditRequest) SetTouchDate(v int64) *CustomerNoteEditRequest {
	s.TouchDate = &v
	return s
}

func (s *CustomerNoteEditRequest) Validate() error {
	return dara.Validate(s)
}
