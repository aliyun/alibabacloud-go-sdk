// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilterDocumentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnd(v []*GetFilterDocumentListRequestAnd) *GetFilterDocumentListRequest
	GetAnd() []*GetFilterDocumentListRequestAnd
	SetDocIdList(v []*string) *GetFilterDocumentListRequest
	GetDocIdList() []*string
	SetLibraryId(v string) *GetFilterDocumentListRequest
	GetLibraryId() *string
	SetOr(v []*GetFilterDocumentListRequestOr) *GetFilterDocumentListRequest
	GetOr() []*GetFilterDocumentListRequestOr
	SetPage(v int32) *GetFilterDocumentListRequest
	GetPage() *int32
	SetPageSize(v int32) *GetFilterDocumentListRequest
	GetPageSize() *int32
	SetStatus(v []*string) *GetFilterDocumentListRequest
	GetStatus() []*string
}

type GetFilterDocumentListRequest struct {
	And       []*GetFilterDocumentListRequestAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	DocIdList []*string                          `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*GetFilterDocumentListRequestOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s GetFilterDocumentListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequest) GetAnd() []*GetFilterDocumentListRequestAnd {
	return s.And
}

func (s *GetFilterDocumentListRequest) GetDocIdList() []*string {
	return s.DocIdList
}

func (s *GetFilterDocumentListRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetFilterDocumentListRequest) GetOr() []*GetFilterDocumentListRequestOr {
	return s.Or
}

func (s *GetFilterDocumentListRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetFilterDocumentListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFilterDocumentListRequest) GetStatus() []*string {
	return s.Status
}

func (s *GetFilterDocumentListRequest) SetAnd(v []*GetFilterDocumentListRequestAnd) *GetFilterDocumentListRequest {
	s.And = v
	return s
}

func (s *GetFilterDocumentListRequest) SetDocIdList(v []*string) *GetFilterDocumentListRequest {
	s.DocIdList = v
	return s
}

func (s *GetFilterDocumentListRequest) SetLibraryId(v string) *GetFilterDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetOr(v []*GetFilterDocumentListRequestOr) *GetFilterDocumentListRequest {
	s.Or = v
	return s
}

func (s *GetFilterDocumentListRequest) SetPage(v int32) *GetFilterDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetPageSize(v int32) *GetFilterDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetStatus(v []*string) *GetFilterDocumentListRequest {
	s.Status = v
	return s
}

func (s *GetFilterDocumentListRequest) Validate() error {
	return dara.Validate(s)
}

type GetFilterDocumentListRequestAnd struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestAnd) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListRequestAnd) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestAnd) GetBoost() *float32 {
	return s.Boost
}

func (s *GetFilterDocumentListRequestAnd) GetKey() *string {
	return s.Key
}

func (s *GetFilterDocumentListRequestAnd) GetOperator() *string {
	return s.Operator
}

func (s *GetFilterDocumentListRequestAnd) GetValue() *string {
	return s.Value
}

func (s *GetFilterDocumentListRequestAnd) SetBoost(v float32) *GetFilterDocumentListRequestAnd {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetKey(v string) *GetFilterDocumentListRequestAnd {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetOperator(v string) *GetFilterDocumentListRequestAnd {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetValue(v string) *GetFilterDocumentListRequestAnd {
	s.Value = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) Validate() error {
	return dara.Validate(s)
}

type GetFilterDocumentListRequestOr struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestOr) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListRequestOr) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestOr) GetBoost() *float32 {
	return s.Boost
}

func (s *GetFilterDocumentListRequestOr) GetKey() *string {
	return s.Key
}

func (s *GetFilterDocumentListRequestOr) GetOperator() *string {
	return s.Operator
}

func (s *GetFilterDocumentListRequestOr) GetValue() *string {
	return s.Value
}

func (s *GetFilterDocumentListRequestOr) SetBoost(v float32) *GetFilterDocumentListRequestOr {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetKey(v string) *GetFilterDocumentListRequestOr {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetOperator(v string) *GetFilterDocumentListRequestOr {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetValue(v string) *GetFilterDocumentListRequestOr {
	s.Value = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) Validate() error {
	return dara.Validate(s)
}
