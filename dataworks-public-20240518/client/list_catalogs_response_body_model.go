// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCatalogsResponseBodyPagingInfo) *ListCatalogsResponseBody
	GetPagingInfo() *ListCatalogsResponseBodyPagingInfo
	SetRequestId(v string) *ListCatalogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCatalogsResponseBody
	GetSuccess() *bool
}

type ListCatalogsResponseBody struct {
	// The pagination result.
	PagingInfo *ListCatalogsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 317CD7D0-AB36-XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) GetPagingInfo() *ListCatalogsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCatalogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCatalogsResponseBody) SetPagingInfo(v *ListCatalogsResponseBodyPagingInfo) *ListCatalogsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCatalogsResponseBody) SetRequestId(v string) *ListCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCatalogsResponseBody) SetSuccess(v bool) *ListCatalogsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCatalogsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCatalogsResponseBodyPagingInfo struct {
	// The catalog.
	Catalogs []*Catalog `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCatalogsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBodyPagingInfo) GetCatalogs() []*Catalog {
	return s.Catalogs
}

func (s *ListCatalogsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCatalogsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCatalogsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCatalogsResponseBodyPagingInfo) SetCatalogs(v []*Catalog) *ListCatalogsResponseBodyPagingInfo {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBodyPagingInfo) SetPageNumber(v int32) *ListCatalogsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCatalogsResponseBodyPagingInfo) SetPageSize(v int32) *ListCatalogsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCatalogsResponseBodyPagingInfo) SetTotalCount(v int64) *ListCatalogsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCatalogsResponseBodyPagingInfo) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
