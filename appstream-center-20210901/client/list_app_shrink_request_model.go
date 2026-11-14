// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppShrinkRequest
	GetAppId() *string
	SetAppName(v string) *ListAppShrinkRequest
	GetAppName() *string
	SetAppType(v string) *ListAppShrinkRequest
	GetAppType() *string
	SetBizType(v int32) *ListAppShrinkRequest
	GetBizType() *int32
	SetBizTypeListShrink(v string) *ListAppShrinkRequest
	GetBizTypeListShrink() *string
	SetOriginAppType(v string) *ListAppShrinkRequest
	GetOriginAppType() *string
	SetOsType(v string) *ListAppShrinkRequest
	GetOsType() *string
	SetPageNumber(v int32) *ListAppShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppShrinkRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAppShrinkRequest
	GetProductType() *string
	SetProductTypeListShrink(v string) *ListAppShrinkRequest
	GetProductTypeListShrink() *string
	SetSubAppType(v string) *ListAppShrinkRequest
	GetSubAppType() *string
}

type ListAppShrinkRequest struct {
	// The application ID used for exact match. If omitted, applications are not filtered by application ID.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name used for filtering. Contains matching is supported.
	//
	// example:
	//
	// Business management
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The source type of the application. Valid values:
	//
	// - **USER_APP**: user-created applications
	//
	// - **SYS_APP**: system-provided applications
	//
	// - **VIRTUAL_APP**: virtual applications
	//
	// example:
	//
	// USER_APP
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// Filters by a single business type. If specified together with BizTypeList, the intersection is used. Valid values:
	//
	// - **1**: basic products
	//
	// - **2**: personal cloud
	//
	// - **3**: cloud office
	//
	// - **4**: Cloud Flow
	//
	// - **10**: Personal Edition
	//
	// - **20**: Business Edition
	//
	// - **60**: Linggou Team Edition
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The list of business types for filtering. An application is matched if it belongs to any business type in the list. If specified together with BizType, the single-value condition must still be met.
	BizTypeListShrink *string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty"`
	// The form of the application. Valid values:
	//
	// - **CLIENT_BASE**: client applications
	//
	// - **WEB_BASE**: browser applications
	//
	// - **CLOUD_NATIVE_BASE**: cloud-native applications
	//
	// - **PORT_SERVICE_BASE**: port service applications
	//
	// - **AGENT_SERVICE_BASE**: agent applications
	//
	// example:
	//
	// WEB_BASE
	OriginAppType *string `json:"OriginAppType,omitempty" xml:"OriginAppType,omitempty"`
	// The type of operating system used by the application. Valid values:
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
	// The page number. We recommend that you start from 1. If this parameter is omitted, set to 0, or set to 1, the first page is queried.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. If this parameter is omitted or set to a value less than 1, the value 10 is used.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by a single associated product. If specified together with ProductTypeList, the intersection is used. Valid values:
	//
	// - **CloudGame**: cloud gaming
	//
	// - **CloudApp**: Wuying Cloud Application
	//
	// - **CloudBrowser**: secure browser
	//
	// - **CloudDesktop**: Wuying Cloud Computer
	//
	// - **ImageBuilder**: image builder resources
	//
	// - **ImageValidate**: image validation resources
	//
	// - **AndroidCloud**: Android cloud
	//
	// - **GpuRender**: GPU rendering
	//
	// - **AndroidGpuRender**: Android cloud GPU rendering
	//
	// - **WuyingServer**: Enterprise Edition workstation
	//
	// - **GpuProxy**: GPU remote rendering proxy
	//
	// - **GpuCache**: GPU remote rendering cache
	//
	// - **AIAgent**: AI agent service
	//
	// - **HostedDesktop**: hosted computer
	//
	// - **WuyingWorkstation**: Personal Edition container workstation
	//
	// - **WuyingWorkstationBusiness**: Business Edition container workstation
	//
	// - **WuyingWorkstationTeam**: Linggou Team Edition container workstation
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of associated products for filtering. An application is matched if it belongs to any product in the list. If specified together with ProductType, the single-value condition must still be met.
	ProductTypeListShrink *string `json:"ProductTypeList,omitempty" xml:"ProductTypeList,omitempty"`
	// The subtype of the application. Valid values:
	//
	// - **PWA**: progressive web applications
	//
	// - **SAAS**: SaaS applications
	//
	// - **SANDBOX**: sandbox applications
	//
	// - **CONTAINER**: container applications
	//
	// example:
	//
	// SAAS
	SubAppType *string `json:"SubAppType,omitempty" xml:"SubAppType,omitempty"`
}

func (s ListAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAppShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListAppShrinkRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListAppShrinkRequest) GetBizTypeListShrink() *string {
	return s.BizTypeListShrink
}

func (s *ListAppShrinkRequest) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *ListAppShrinkRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListAppShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppShrinkRequest) GetProductTypeListShrink() *string {
	return s.ProductTypeListShrink
}

func (s *ListAppShrinkRequest) GetSubAppType() *string {
	return s.SubAppType
}

func (s *ListAppShrinkRequest) SetAppId(v string) *ListAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListAppShrinkRequest) SetAppName(v string) *ListAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ListAppShrinkRequest) SetAppType(v string) *ListAppShrinkRequest {
	s.AppType = &v
	return s
}

func (s *ListAppShrinkRequest) SetBizType(v int32) *ListAppShrinkRequest {
	s.BizType = &v
	return s
}

func (s *ListAppShrinkRequest) SetBizTypeListShrink(v string) *ListAppShrinkRequest {
	s.BizTypeListShrink = &v
	return s
}

func (s *ListAppShrinkRequest) SetOriginAppType(v string) *ListAppShrinkRequest {
	s.OriginAppType = &v
	return s
}

func (s *ListAppShrinkRequest) SetOsType(v string) *ListAppShrinkRequest {
	s.OsType = &v
	return s
}

func (s *ListAppShrinkRequest) SetPageNumber(v int32) *ListAppShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppShrinkRequest) SetPageSize(v int32) *ListAppShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppShrinkRequest) SetProductType(v string) *ListAppShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppShrinkRequest) SetProductTypeListShrink(v string) *ListAppShrinkRequest {
	s.ProductTypeListShrink = &v
	return s
}

func (s *ListAppShrinkRequest) SetSubAppType(v string) *ListAppShrinkRequest {
	s.SubAppType = &v
	return s
}

func (s *ListAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
