// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListKnowledgeFilesRequest
	GetDBClusterId() *string
	SetFileIds(v string) *ListKnowledgeFilesRequest
	GetFileIds() *string
	SetPage(v string) *ListKnowledgeFilesRequest
	GetPage() *string
	SetPageSize(v string) *ListKnowledgeFilesRequest
	GetPageSize() *string
	SetStatus(v string) *ListKnowledgeFilesRequest
	GetStatus() *string
	SetUser(v string) *ListKnowledgeFilesRequest
	GetUser() *string
}

type ListKnowledgeFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// am-bp19aaaaaa****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// [1001,1002,1003]
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// u123
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListKnowledgeFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeFilesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeFilesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListKnowledgeFilesRequest) GetFileIds() *string {
	return s.FileIds
}

func (s *ListKnowledgeFilesRequest) GetPage() *string {
	return s.Page
}

func (s *ListKnowledgeFilesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListKnowledgeFilesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListKnowledgeFilesRequest) GetUser() *string {
	return s.User
}

func (s *ListKnowledgeFilesRequest) SetDBClusterId(v string) *ListKnowledgeFilesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListKnowledgeFilesRequest) SetFileIds(v string) *ListKnowledgeFilesRequest {
	s.FileIds = &v
	return s
}

func (s *ListKnowledgeFilesRequest) SetPage(v string) *ListKnowledgeFilesRequest {
	s.Page = &v
	return s
}

func (s *ListKnowledgeFilesRequest) SetPageSize(v string) *ListKnowledgeFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeFilesRequest) SetStatus(v string) *ListKnowledgeFilesRequest {
	s.Status = &v
	return s
}

func (s *ListKnowledgeFilesRequest) SetUser(v string) *ListKnowledgeFilesRequest {
	s.User = &v
	return s
}

func (s *ListKnowledgeFilesRequest) Validate() error {
	return dara.Validate(s)
}
