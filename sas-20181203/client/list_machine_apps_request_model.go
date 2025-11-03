// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMachineAppsRequest
	GetAppId() *string
	SetAppName(v string) *ListMachineAppsRequest
	GetAppName() *string
	SetAppRegionId(v string) *ListMachineAppsRequest
	GetAppRegionId() *string
	SetAuthVersion(v string) *ListMachineAppsRequest
	GetAuthVersion() *string
	SetCurrentPage(v int32) *ListMachineAppsRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListMachineAppsRequest
	GetLang() *string
	SetPageSize(v int32) *ListMachineAppsRequest
	GetPageSize() *int32
	SetResourceDirectoryUid(v int64) *ListMachineAppsRequest
	GetResourceDirectoryUid() *int64
}

type ListMachineAppsRequest struct {
	// The ID of the SAE application.
	//
	// example:
	//
	// 5b41f4bf-349f-4263-89b1-9234c034****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the SAE application.
	//
	// example:
	//
	// agent-commprice-shop
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	AppRegionId *string `json:"AppRegionId,omitempty" xml:"AppRegionId,omitempty"`
	// The edition of Security Center. Valid values:
	//
	// 	- **6**: Anti-virus.
	//
	// 	- **5**: Advanced.
	//
	// 	- **3**: Enterprise.
	//
	// 	- **7**: Ultimate.
	//
	// 	- **10**: Value-added Plan.
	//
	// example:
	//
	// 7
	AuthVersion *string `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the resource directory.
	//
	// example:
	//
	// 123456
	ResourceDirectoryUid *int64 `json:"ResourceDirectoryUid,omitempty" xml:"ResourceDirectoryUid,omitempty"`
}

func (s ListMachineAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineAppsRequest) GoString() string {
	return s.String()
}

func (s *ListMachineAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMachineAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListMachineAppsRequest) GetAppRegionId() *string {
	return s.AppRegionId
}

func (s *ListMachineAppsRequest) GetAuthVersion() *string {
	return s.AuthVersion
}

func (s *ListMachineAppsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMachineAppsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListMachineAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMachineAppsRequest) GetResourceDirectoryUid() *int64 {
	return s.ResourceDirectoryUid
}

func (s *ListMachineAppsRequest) SetAppId(v string) *ListMachineAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListMachineAppsRequest) SetAppName(v string) *ListMachineAppsRequest {
	s.AppName = &v
	return s
}

func (s *ListMachineAppsRequest) SetAppRegionId(v string) *ListMachineAppsRequest {
	s.AppRegionId = &v
	return s
}

func (s *ListMachineAppsRequest) SetAuthVersion(v string) *ListMachineAppsRequest {
	s.AuthVersion = &v
	return s
}

func (s *ListMachineAppsRequest) SetCurrentPage(v int32) *ListMachineAppsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListMachineAppsRequest) SetLang(v string) *ListMachineAppsRequest {
	s.Lang = &v
	return s
}

func (s *ListMachineAppsRequest) SetPageSize(v int32) *ListMachineAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMachineAppsRequest) SetResourceDirectoryUid(v int64) *ListMachineAppsRequest {
	s.ResourceDirectoryUid = &v
	return s
}

func (s *ListMachineAppsRequest) Validate() error {
	return dara.Validate(s)
}
