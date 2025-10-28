// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlsLogStoreListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QuerySlsLogStoreListRequest
	GetAppId() *string
	SetCurrentPage(v int32) *QuerySlsLogStoreListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QuerySlsLogStoreListRequest
	GetPageSize() *int32
	SetType(v string) *QuerySlsLogStoreListRequest
	GetType() *string
}

type QuerySlsLogStoreListRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// af58edcf-f7eb-****-****-db4e425f****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The number of the page to return. Pages start from Page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 1000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of data that is collected by Log Service. Valid values:
	//
	// 	- file: the file type
	//
	// 	- stdout: the standard output type
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySlsLogStoreListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySlsLogStoreListRequest) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListRequest) GetAppId() *string {
	return s.AppId
}

func (s *QuerySlsLogStoreListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QuerySlsLogStoreListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySlsLogStoreListRequest) GetType() *string {
	return s.Type
}

func (s *QuerySlsLogStoreListRequest) SetAppId(v string) *QuerySlsLogStoreListRequest {
	s.AppId = &v
	return s
}

func (s *QuerySlsLogStoreListRequest) SetCurrentPage(v int32) *QuerySlsLogStoreListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuerySlsLogStoreListRequest) SetPageSize(v int32) *QuerySlsLogStoreListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySlsLogStoreListRequest) SetType(v string) *QuerySlsLogStoreListRequest {
	s.Type = &v
	return s
}

func (s *QuerySlsLogStoreListRequest) Validate() error {
	return dara.Validate(s)
}
