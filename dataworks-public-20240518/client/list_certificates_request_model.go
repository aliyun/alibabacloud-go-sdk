// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateUser(v string) *ListCertificatesRequest
	GetCreateUser() *string
	SetEndCreateTime(v int64) *ListCertificatesRequest
	GetEndCreateTime() *int64
	SetName(v string) *ListCertificatesRequest
	GetName() *string
	SetOrder(v string) *ListCertificatesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListCertificatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCertificatesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCertificatesRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListCertificatesRequest
	GetSortBy() *string
	SetStartCreateTime(v int64) *ListCertificatesRequest
	GetStartCreateTime() *int64
}

type ListCertificatesRequest struct {
	// The ID of the user who created the certificate files.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The time when the certificate file was created. You can call this operation to query the files that are created before the time. Unit: milliseconds.
	//
	// example:
	//
	// 1593877765000
	EndCreateTime *int64 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	// The name of the certificate file. Fuzzy match by file name is supported.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order in which you want to sort the certificate files. Valid values: Desc: descending order ASC: ascending order Default value: Asc
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace to which the certificate file belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field used to sort the certificate files. Valid values: CreateTime Id Name Default value: Id
	//
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The time when the certificate file was created. You can call this operation to query the files that are created after the time. Unit: milliseconds.
	//
	// example:
	//
	// 1730217600000
	StartCreateTime *int64 `json:"StartCreateTime,omitempty" xml:"StartCreateTime,omitempty"`
}

func (s ListCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListCertificatesRequest) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListCertificatesRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *ListCertificatesRequest) GetName() *string {
	return s.Name
}

func (s *ListCertificatesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCertificatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCertificatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCertificatesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCertificatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCertificatesRequest) GetStartCreateTime() *int64 {
	return s.StartCreateTime
}

func (s *ListCertificatesRequest) SetCreateUser(v string) *ListCertificatesRequest {
	s.CreateUser = &v
	return s
}

func (s *ListCertificatesRequest) SetEndCreateTime(v int64) *ListCertificatesRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListCertificatesRequest) SetName(v string) *ListCertificatesRequest {
	s.Name = &v
	return s
}

func (s *ListCertificatesRequest) SetOrder(v string) *ListCertificatesRequest {
	s.Order = &v
	return s
}

func (s *ListCertificatesRequest) SetPageNumber(v int32) *ListCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesRequest) SetPageSize(v int32) *ListCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCertificatesRequest) SetProjectId(v int64) *ListCertificatesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCertificatesRequest) SetSortBy(v string) *ListCertificatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCertificatesRequest) SetStartCreateTime(v int64) *ListCertificatesRequest {
	s.StartCreateTime = &v
	return s
}

func (s *ListCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
