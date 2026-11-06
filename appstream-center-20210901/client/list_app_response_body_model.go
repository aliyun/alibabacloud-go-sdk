// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppResponseBody
	GetCode() *string
	SetCount(v int32) *ListAppResponseBody
	GetCount() *int32
	SetData(v []*ListAppResponseBodyData) *ListAppResponseBody
	GetData() []*ListAppResponseBodyData
	SetMessage(v string) *ListAppResponseBody
	GetMessage() *string
	SetPageNmuber(v int32) *ListAppResponseBody
	GetPageNmuber() *int32
	SetPageSize(v int32) *ListAppResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppResponseBody
	GetSuccess() *bool
}

type ListAppResponseBody struct {
	// The error code. This field is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of application records that meet the query conditions.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The application records on the current page.
	Data []*ListAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error or result message. This field describes the cause when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the current query. The name of this response field is PageNmuber.
	//
	// example:
	//
	// 1
	PageNmuber *int32 `json:"PageNmuber,omitempty" xml:"PageNmuber,omitempty"`
	// The number of records per page in the paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request. You can provide this value when troubleshooting issues.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the current request was processed successfully. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListAppResponseBody) GetData() []*ListAppResponseBodyData {
	return s.Data
}

func (s *ListAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppResponseBody) GetPageNmuber() *int32 {
	return s.PageNmuber
}

func (s *ListAppResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppResponseBody) SetCode(v string) *ListAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResponseBody) SetCount(v int32) *ListAppResponseBody {
	s.Count = &v
	return s
}

func (s *ListAppResponseBody) SetData(v []*ListAppResponseBodyData) *ListAppResponseBody {
	s.Data = v
	return s
}

func (s *ListAppResponseBody) SetMessage(v string) *ListAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppResponseBody) SetPageNmuber(v int32) *ListAppResponseBody {
	s.PageNmuber = &v
	return s
}

func (s *ListAppResponseBody) SetPageSize(v int32) *ListAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppResponseBody) SetRequestId(v string) *ListAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResponseBody) SetSuccess(v bool) *ListAppResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppResponseBodyData struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The URL used to access the application icon.
	//
	// example:
	//
	// https://example.com/apps/business/icon.png
	AppIconUrl *string `json:"AppIconUrl,omitempty" xml:"AppIconUrl,omitempty"`
	// The application ID, which is used for subsequent application queries or unbinding the image.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The display name of the application.
	//
	// example:
	//
	// Business management platform
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The theme color saved for the application. If no theme color is specified when the application is created, #FFFFFF (white) is saved by default. The saved value is returned during a query. This field may not be configured for legacy applications.
	//
	// example:
	//
	// #FFFFFF
	AppThemeColor *string `json:"AppThemeColor,omitempty" xml:"AppThemeColor,omitempty"`
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
	// The business type to which the application belongs. Valid values:
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
	// Indicates whether the current application is not referenced by any image. A value of true indicates that no image references the application. A value of false indicates that the application is still referenced by an image, and the reference must be removed before deletion.
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// The path of the executable file saved when the application is created or modified. This field is returned only if the path is configured for the application. An example is the installation path of a Windows business application.
	//
	// example:
	//
	// C:\\Program Files\\BusinessApp\\BusinessApp.exe
	ExecuteFileAddress *string `json:"ExecuteFileAddress,omitempty" xml:"ExecuteFileAddress,omitempty"`
	// The time when the application was created. The time is in ISO 8601 format, accurate to milliseconds, and the time zone offset is +00:00.
	//
	// example:
	//
	// 2026-08-25T07:41:30.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the application was last modified. The time is in ISO 8601 format, accurate to milliseconds, and the time zone offset is +00:00.
	//
	// example:
	//
	// 2026-08-25T07:41:30.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
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
	// The product to which the resource belongs. Valid values:
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
	// The command-line parameters used when the application is started.
	//
	// example:
	//
	// --lang=zh-CN
	SilenceParam *string `json:"SilenceParam,omitempty" xml:"SilenceParam,omitempty"`
	// The status of the application record. The current list queries only valid and undeleted applications.
	//
	// example:
	//
	// VALID
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s ListAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListAppResponseBodyData) GetAppIconUrl() *string {
	return s.AppIconUrl
}

func (s *ListAppResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListAppResponseBodyData) GetAppThemeColor() *string {
	return s.AppThemeColor
}

func (s *ListAppResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *ListAppResponseBodyData) GetBizType() *int32 {
	return s.BizType
}

func (s *ListAppResponseBodyData) GetDeletable() *bool {
	return s.Deletable
}

func (s *ListAppResponseBodyData) GetExecuteFileAddress() *string {
	return s.ExecuteFileAddress
}

func (s *ListAppResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAppResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAppResponseBodyData) GetOriginAppType() *string {
	return s.OriginAppType
}

func (s *ListAppResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *ListAppResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppResponseBodyData) GetSilenceParam() *string {
	return s.SilenceParam
}

func (s *ListAppResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAppResponseBodyData) GetSubAppType() *string {
	return s.SubAppType
}

func (s *ListAppResponseBodyData) SetAliUid(v int64) *ListAppResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppIconUrl(v string) *ListAppResponseBodyData {
	s.AppIconUrl = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppId(v string) *ListAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppName(v string) *ListAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppThemeColor(v string) *ListAppResponseBodyData {
	s.AppThemeColor = &v
	return s
}

func (s *ListAppResponseBodyData) SetAppType(v string) *ListAppResponseBodyData {
	s.AppType = &v
	return s
}

func (s *ListAppResponseBodyData) SetBizType(v int32) *ListAppResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListAppResponseBodyData) SetDeletable(v bool) *ListAppResponseBodyData {
	s.Deletable = &v
	return s
}

func (s *ListAppResponseBodyData) SetExecuteFileAddress(v string) *ListAppResponseBodyData {
	s.ExecuteFileAddress = &v
	return s
}

func (s *ListAppResponseBodyData) SetGmtCreate(v string) *ListAppResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAppResponseBodyData) SetGmtModified(v string) *ListAppResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAppResponseBodyData) SetOriginAppType(v string) *ListAppResponseBodyData {
	s.OriginAppType = &v
	return s
}

func (s *ListAppResponseBodyData) SetOsType(v string) *ListAppResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListAppResponseBodyData) SetProductType(v string) *ListAppResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListAppResponseBodyData) SetSilenceParam(v string) *ListAppResponseBodyData {
	s.SilenceParam = &v
	return s
}

func (s *ListAppResponseBodyData) SetStatus(v string) *ListAppResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAppResponseBodyData) SetSubAppType(v string) *ListAppResponseBodyData {
	s.SubAppType = &v
	return s
}

func (s *ListAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
