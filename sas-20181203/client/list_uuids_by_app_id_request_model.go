// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByAppIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListUuidsByAppIdRequest
	GetAppId() *string
	SetAppRegionId(v string) *ListUuidsByAppIdRequest
	GetAppRegionId() *string
	SetCurrentPage(v int32) *ListUuidsByAppIdRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListUuidsByAppIdRequest
	GetLang() *string
	SetPageSize(v int32) *ListUuidsByAppIdRequest
	GetPageSize() *int32
	SetResourceDirectoryUid(v int64) *ListUuidsByAppIdRequest
	GetResourceDirectoryUid() *int64
}

type ListUuidsByAppIdRequest struct {
	// SAE application ID.
	//
	// example:
	//
	// 5b41f4bf-349f-4263-89b1-9234c034****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	AppRegionId *string `json:"AppRegionId,omitempty" xml:"AppRegionId,omitempty"`
	// The page number to display in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for request and response, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of items to display per page in a paginated query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource associated directory UID.
	//
	// example:
	//
	// 123456
	ResourceDirectoryUid *int64 `json:"ResourceDirectoryUid,omitempty" xml:"ResourceDirectoryUid,omitempty"`
}

func (s ListUuidsByAppIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByAppIdRequest) GoString() string {
	return s.String()
}

func (s *ListUuidsByAppIdRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListUuidsByAppIdRequest) GetAppRegionId() *string {
	return s.AppRegionId
}

func (s *ListUuidsByAppIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUuidsByAppIdRequest) GetLang() *string {
	return s.Lang
}

func (s *ListUuidsByAppIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUuidsByAppIdRequest) GetResourceDirectoryUid() *int64 {
	return s.ResourceDirectoryUid
}

func (s *ListUuidsByAppIdRequest) SetAppId(v string) *ListUuidsByAppIdRequest {
	s.AppId = &v
	return s
}

func (s *ListUuidsByAppIdRequest) SetAppRegionId(v string) *ListUuidsByAppIdRequest {
	s.AppRegionId = &v
	return s
}

func (s *ListUuidsByAppIdRequest) SetCurrentPage(v int32) *ListUuidsByAppIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUuidsByAppIdRequest) SetLang(v string) *ListUuidsByAppIdRequest {
	s.Lang = &v
	return s
}

func (s *ListUuidsByAppIdRequest) SetPageSize(v int32) *ListUuidsByAppIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListUuidsByAppIdRequest) SetResourceDirectoryUid(v int64) *ListUuidsByAppIdRequest {
	s.ResourceDirectoryUid = &v
	return s
}

func (s *ListUuidsByAppIdRequest) Validate() error {
	return dara.Validate(s)
}
