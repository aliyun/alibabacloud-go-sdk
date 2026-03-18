// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversationListQry interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ConversationListQry
	GetKeyword() *string
	SetPage(v int32) *ConversationListQry
	GetPage() *int32
	SetPageSize(v int32) *ConversationListQry
	GetPageSize() *int32
	SetStatus(v int32) *ConversationListQry
	GetStatus() *int32
}

type ConversationListQry struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ConversationListQry) String() string {
	return dara.Prettify(s)
}

func (s ConversationListQry) GoString() string {
	return s.String()
}

func (s *ConversationListQry) GetKeyword() *string {
	return s.Keyword
}

func (s *ConversationListQry) GetPage() *int32 {
	return s.Page
}

func (s *ConversationListQry) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ConversationListQry) GetStatus() *int32 {
	return s.Status
}

func (s *ConversationListQry) SetKeyword(v string) *ConversationListQry {
	s.Keyword = &v
	return s
}

func (s *ConversationListQry) SetPage(v int32) *ConversationListQry {
	s.Page = &v
	return s
}

func (s *ConversationListQry) SetPageSize(v int32) *ConversationListQry {
	s.PageSize = &v
	return s
}

func (s *ConversationListQry) SetStatus(v int32) *ConversationListQry {
	s.Status = &v
	return s
}

func (s *ConversationListQry) Validate() error {
	return dara.Validate(s)
}
