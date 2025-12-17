// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *ListFileUploadRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *ListFileUploadRequest
	GetDmsUnit() *string
	SetFileCategory(v string) *ListFileUploadRequest
	GetFileCategory() *string
	SetFileFrom(v string) *ListFileUploadRequest
	GetFileFrom() *string
	SetFileId(v string) *ListFileUploadRequest
	GetFileId() *string
	SetSessionId(v string) *ListFileUploadRequest
	GetSessionId() *string
	SetSortColumn(v string) *ListFileUploadRequest
	GetSortColumn() *string
	SetSortDirection(v string) *ListFileUploadRequest
	GetSortDirection() *string
}

type ListFileUploadRequest struct {
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// example:
	//
	// TextReport
	FileCategory *string `json:"FileCategory,omitempty" xml:"FileCategory,omitempty"`
	// example:
	//
	// Agent
	FileFrom *string `json:"FileFrom,omitempty" xml:"FileFrom,omitempty"`
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// h8r********4fch
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// gmtCreated
	SortColumn *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s ListFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileUploadRequest) GoString() string {
	return s.String()
}

func (s *ListFileUploadRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *ListFileUploadRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *ListFileUploadRequest) GetFileCategory() *string {
	return s.FileCategory
}

func (s *ListFileUploadRequest) GetFileFrom() *string {
	return s.FileFrom
}

func (s *ListFileUploadRequest) GetFileId() *string {
	return s.FileId
}

func (s *ListFileUploadRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListFileUploadRequest) GetSortColumn() *string {
	return s.SortColumn
}

func (s *ListFileUploadRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *ListFileUploadRequest) SetCallFrom(v string) *ListFileUploadRequest {
	s.CallFrom = &v
	return s
}

func (s *ListFileUploadRequest) SetDmsUnit(v string) *ListFileUploadRequest {
	s.DmsUnit = &v
	return s
}

func (s *ListFileUploadRequest) SetFileCategory(v string) *ListFileUploadRequest {
	s.FileCategory = &v
	return s
}

func (s *ListFileUploadRequest) SetFileFrom(v string) *ListFileUploadRequest {
	s.FileFrom = &v
	return s
}

func (s *ListFileUploadRequest) SetFileId(v string) *ListFileUploadRequest {
	s.FileId = &v
	return s
}

func (s *ListFileUploadRequest) SetSessionId(v string) *ListFileUploadRequest {
	s.SessionId = &v
	return s
}

func (s *ListFileUploadRequest) SetSortColumn(v string) *ListFileUploadRequest {
	s.SortColumn = &v
	return s
}

func (s *ListFileUploadRequest) SetSortDirection(v string) *ListFileUploadRequest {
	s.SortDirection = &v
	return s
}

func (s *ListFileUploadRequest) Validate() error {
	return dara.Validate(s)
}
