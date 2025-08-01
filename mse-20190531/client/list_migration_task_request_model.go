// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListMigrationTaskRequest
	GetAcceptLanguage() *string
	SetOriginInstanceName(v string) *ListMigrationTaskRequest
	GetOriginInstanceName() *string
	SetPageNum(v int64) *ListMigrationTaskRequest
	GetPageNum() *int64
	SetPageSize(v int64) *ListMigrationTaskRequest
	GetPageSize() *int64
	SetRequestPars(v string) *ListMigrationTaskRequest
	GetRequestPars() *string
}

type ListMigrationTaskRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The name of the source instance.
	//
	// example:
	//
	// whdc
	OriginInstanceName *string `json:"OriginInstanceName,omitempty" xml:"OriginInstanceName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s ListMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationTaskRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListMigrationTaskRequest) GetOriginInstanceName() *string {
	return s.OriginInstanceName
}

func (s *ListMigrationTaskRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListMigrationTaskRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMigrationTaskRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListMigrationTaskRequest) SetAcceptLanguage(v string) *ListMigrationTaskRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListMigrationTaskRequest) SetOriginInstanceName(v string) *ListMigrationTaskRequest {
	s.OriginInstanceName = &v
	return s
}

func (s *ListMigrationTaskRequest) SetPageNum(v int64) *ListMigrationTaskRequest {
	s.PageNum = &v
	return s
}

func (s *ListMigrationTaskRequest) SetPageSize(v int64) *ListMigrationTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationTaskRequest) SetRequestPars(v string) *ListMigrationTaskRequest {
	s.RequestPars = &v
	return s
}

func (s *ListMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
