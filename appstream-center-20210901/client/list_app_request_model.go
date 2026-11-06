// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppRequest
	GetAppId() *string
	SetAppName(v string) *ListAppRequest
	GetAppName() *string
	SetAppType(v string) *ListAppRequest
	GetAppType() *string
	SetBizType(v int32) *ListAppRequest
	GetBizType() *int32
	SetBizTypeList(v []*int32) *ListAppRequest
	GetBizTypeList() []*int32
	SetOriginAppType(v string) *ListAppRequest
	GetOriginAppType() *string
	SetOsType(v string) *ListAppRequest
	GetOsType() *string
	SetPageNumber(v int32) *ListAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAppRequest
	GetProductType() *string
	SetProductTypeList(v []*string) *ListAppRequest
	GetProductTypeList() []*string
	SetSubAppType(v string) *ListAppRequest
	GetSubAppType() *string
}

type ListAppRequest struct {
	// The application ID for exact query. If this parameter is omitted, applications are not filtered by application ID.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name filter. Contains matching is supported.
	//
	// example:
	//
	// Business management
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The source type of the application. Valid values:
	//
	// - **USER_APP**: User-created applications.
	//
	// - **SYS_APP**: System-provided applications.
	//
	// - **VIRTUAL_APP**: Virtual applications.
	//
	// example:
	//
	// USER_APP
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Filters by a single business type. If specified together with BizTypeList, the intersection is used. Valid values:
	//
	// - **1**: Basic products.
	//
	// - **2**: Personal cloud.
	//
	// - **3**: Cloud office.
	//
	// - **4**: Cloud Flow.
	//
	// - **10**: Personal Edition.
	//
	// - **20**: Business Edition.
	//
	// - **60**: Linggou Team Edition.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of business types for filtering. A match with any business type in the list is considered a match. If specified together with BizType, the single-value condition must still be met.
	BizTypeList []*int32 `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The form of the application. Valid values:
	//
	// - **CLIENT_BASE**: Client applications.
	//
	// - **WEB_BASE**: Browser applications.
	//
	// - **CLOUD_NATIVE_BASE**: Cloud-native applications.
	//
	// - **PORT_SERVICE_BASE**: Port service applications.
	//
	// - **AGENT_SERVICE_BASE**: Agent applications.
	//
	// example:
	//
	// WEB_BASE
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// The type of operating system used by the application. Valid values:
	//
	// - **Unknown**: The operating system is unknown.
	//
	// - **Windows**: Windows operating system.
	//
	// - **Linux**: Linux operating system.
	//
	// - **Android**: Android operating system.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The page number for the query. Start from 1. If this parameter is omitted, set to 0, or set to 1, the first page is queried.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page. If this parameter is omitted or set to a value less than 1, the value 10 is used.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by a single associated product. If specified together with ProductTypeList, the intersection is used. Valid values:
	//
	// - **CloudGame**: Cloud game.
	//
	// - **CloudApp**: Wuying cloud application.
	//
	// - **CloudBrowser**: Secure browser.
	//
	// - **CloudDesktop**: Wuying cloud desktop.
	//
	// - **ImageBuilder**: Image builder resources.
	//
	// - **ImageValidate**: Image validation resources.
	//
	// - **AndroidCloud**: Android cloud.
	//
	// - **GpuRender**: GPU rendering.
	//
	// - **AndroidGpuRender**: Android cloud GPU rendering.
	//
	// - **WuyingServer**: Enterprise Edition workstation.
	//
	// - **GpuProxy**: GPU remote rendering proxy.
	//
	// - **GpuCache**: GPU remote rendering cache.
	//
	// - **AIAgent**: AI Agent service.
	//
	// - **HostedDesktop**: Hosted desktop.
	//
	// - **WuyingWorkstation**: Personal Edition container workstation.
	//
	// - **WuyingWorkstationBusiness**: Business Edition container workstation.
	//
	// - **WuyingWorkstationTeam**: Linggou Team Edition container workstation.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of associated products for filtering. A match with any product in the list is considered a match. If specified together with ProductType, the single-value condition must still be met.
	ProductTypeList []*string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty" type:"Repeated"`
	// The subtype of the application. Valid values:
	//
	// - **PWA**: Progressive web applications.
	//
	// - **SAAS**: Software as a Service (SaaS) applications.
	//
	// - **SANDBOX**: Sandbox applications.
	//
	// - **CONTAINER**: Container applications.
	//
	// example:
	//
	// SAAS
	SubAppType *string `json:"SubAppType,omitempty" xml:"SubAppType,omitempty"`
}

func (s ListAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppRequest) GoString() string {
	return s.String()
}

func (s *ListAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAppRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListAppRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListAppRequest) GetBizTypeList() []*int32 {
	return s.BizTypeList
}

func (s *ListAppRequest) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *ListAppRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppRequest) GetProductTypeList() []*string {
	return s.ProductTypeList
}

func (s *ListAppRequest) GetSubAppType() *string {
	return s.SubAppType
}

func (s *ListAppRequest) SetAppId(v string) *ListAppRequest {
	s.AppId = &v
	return s
}

func (s *ListAppRequest) SetAppName(v string) *ListAppRequest {
	s.AppName = &v
	return s
}

func (s *ListAppRequest) SetAppType(v string) *ListAppRequest {
	s.AppType = &v
	return s
}

func (s *ListAppRequest) SetBizType(v int32) *ListAppRequest {
	s.BizType = &v
	return s
}

func (s *ListAppRequest) SetBizTypeList(v []*int32) *ListAppRequest {
	s.BizTypeList = v
	return s
}

func (s *ListAppRequest) SetOriginAppType(v string) *ListAppRequest {
	s.OriginAppType = &v
	return s
}

func (s *ListAppRequest) SetOsType(v string) *ListAppRequest {
	s.OsType = &v
	return s
}

func (s *ListAppRequest) SetPageNumber(v int32) *ListAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppRequest) SetPageSize(v int32) *ListAppRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppRequest) SetProductType(v string) *ListAppRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppRequest) SetProductTypeList(v []*string) *ListAppRequest {
	s.ProductTypeList = v
	return s
}

func (s *ListAppRequest) SetSubAppType(v string) *ListAppRequest {
	s.SubAppType = &v
	return s
}

func (s *ListAppRequest) Validate() error {
	return dara.Validate(s)
}
