// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTempFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelimiter(v string) *ListTempFilesRequest
	GetDelimiter() *string
	SetInstanceId(v string) *ListTempFilesRequest
	GetInstanceId() *string
	SetName(v string) *ListTempFilesRequest
	GetName() *string
	SetOrder(v string) *ListTempFilesRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListTempFilesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTempFilesRequest
	GetPageSize() *int64
	SetPrefix(v string) *ListTempFilesRequest
	GetPrefix() *string
	SetSortBy(v string) *ListTempFilesRequest
	GetSortBy() *string
}

type ListTempFilesRequest struct {
	// example:
	//
	// /
	Delimiter *string `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// filename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// somedir/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListTempFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTempFilesRequest) GoString() string {
	return s.String()
}

func (s *ListTempFilesRequest) GetDelimiter() *string {
	return s.Delimiter
}

func (s *ListTempFilesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTempFilesRequest) GetName() *string {
	return s.Name
}

func (s *ListTempFilesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTempFilesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTempFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTempFilesRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListTempFilesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTempFilesRequest) SetDelimiter(v string) *ListTempFilesRequest {
	s.Delimiter = &v
	return s
}

func (s *ListTempFilesRequest) SetInstanceId(v string) *ListTempFilesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTempFilesRequest) SetName(v string) *ListTempFilesRequest {
	s.Name = &v
	return s
}

func (s *ListTempFilesRequest) SetOrder(v string) *ListTempFilesRequest {
	s.Order = &v
	return s
}

func (s *ListTempFilesRequest) SetPageNumber(v int64) *ListTempFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTempFilesRequest) SetPageSize(v int64) *ListTempFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTempFilesRequest) SetPrefix(v string) *ListTempFilesRequest {
	s.Prefix = &v
	return s
}

func (s *ListTempFilesRequest) SetSortBy(v string) *ListTempFilesRequest {
	s.SortBy = &v
	return s
}

func (s *ListTempFilesRequest) Validate() error {
	return dara.Validate(s)
}
