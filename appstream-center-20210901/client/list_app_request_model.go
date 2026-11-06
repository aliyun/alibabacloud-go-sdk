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
	// Specifies the application ID for exact match. If omitted, applications are not filtered by application ID.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name used to filter applications. Partial match is supported.
	//
	// example:
	//
	// Business management
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The source type of the application. Valid values:
	//
	// - **USER_APP**: user-created application
	//
	// - **SYS_APP**: system-provided application
	//
	// - **VIRTUAL_APP**: virtual application
	//
	// example:
	//
	// USER_APP
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Filters by a single business type. If specified simultaneously with BizTypeList, the intersection is used. Valid values:
	//
	// - **1**: basic product
	//
	// - **2**: personal cloud
	//
	// - **3**: cloud office
	//
	// - **4**: cloud flow
	//
	// - **10**: personal edition
	//
	// - **20**: business edition
	//
	// - **60**: Linggou team edition
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of business types used to filter applications. An application is matched if it matches any business type in the list. If specified simultaneously with BizType, the single-value condition must still be met.
	BizTypeList []*int32 `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The form of the application. Valid values:
	//
	// - **CLIENT_BASE**: client application
	//
	// - **WEB_BASE**: browser application
	//
	// - **CLOUD_NATIVE_BASE**: cloud-native application
	//
	// - **PORT_SERVICE_BASE**: port service application
	//
	// - **AGENT_SERVICE_BASE**: agent application
	//
	// example:
	//
	// WEB_BASE
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// The type of the operating system used by the application. Valid values:
	//
	// - **Unknown**: unknown operating system
	//
	// - **Windows**: Windows operating system
	//
	// - **Linux**: Linux operating system
	//
	// - **Android**: Android operating system
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The page number for the query. Query pages starting from 1. If omitted, set to 0, or set to 1, the first page is queried.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records returned per page. If omitted or set to a value less than 1, the value is processed as 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by a single associated product. If specified simultaneously with ProductTypeList, the intersection is used. Valid values:
	//
	// - **CloudGame**: cloud game
	//
	// - **CloudApp**: Wuying cloud application
	//
	// - **CloudBrowser**: secure browser
	//
	// - **CloudDesktop**: Wuying cloud computer
	//
	// - **ImageBuilder**: image building resource
	//
	// - **ImageValidate**: image validation resource
	//
	// - **AndroidCloud**: Android cloud
	//
	// - **GpuRender**: GPU rendering
	//
	// - **AndroidGpuRender**: Android cloud GPU rendering
	//
	// - **WuyingServer**: enterprise edition workstation
	//
	// - **GpuProxy**: GPU remote rendering proxy
	//
	// - **GpuCache**: GPU remote rendering cache
	//
	// - **AIAgent**: AI agent service
	//
	// - **HostedDesktop**: hosted computer
	//
	// - **WuyingWorkstation**: personal edition container workstation
	//
	// - **WuyingWorkstationBusiness**: business edition container workstation
	//
	// - **WuyingWorkstationTeam**: Linggou team edition container workstation
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of associated products used to filter applications. An application is matched if it matches any product in the list. If specified simultaneously with ProductType, the single-value condition must still be met.
	ProductTypeList []*string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty" type:"Repeated"`
	// The subtype of the application. Valid values:
	//
	// - **PWA**: progressive web application
	//
	// - **SAAS**: software as a service (SaaS) application
	//
	// - **SANDBOX**: sandbox application
	//
	// - **CONTAINER**: container application
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
