// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsAppIcpRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySmsAppIcpRecordResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySmsAppIcpRecordResponseBody
	GetCode() *string
	SetData(v []*QuerySmsAppIcpRecordResponseBodyData) *QuerySmsAppIcpRecordResponseBody
	GetData() []*QuerySmsAppIcpRecordResponseBodyData
	SetMessage(v string) *QuerySmsAppIcpRecordResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySmsAppIcpRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySmsAppIcpRecordResponseBody
	GetSuccess() *bool
}

type QuerySmsAppIcpRecordResponseBody struct {
	// The access denial details.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - OK indicates a successful request.
	//
	// - For other error codes, see the [error code list](https://help.aliyun.com/document_detail/101346.htm).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list of APP-ICP record entity details.
	Data []*QuerySmsAppIcpRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C9955E63-8BFF-101D-80A1-E6998DFEFF1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the API call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySmsAppIcpRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAppIcpRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmsAppIcpRecordResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySmsAppIcpRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySmsAppIcpRecordResponseBody) GetData() []*QuerySmsAppIcpRecordResponseBodyData {
	return s.Data
}

func (s *QuerySmsAppIcpRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySmsAppIcpRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmsAppIcpRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySmsAppIcpRecordResponseBody) SetAccessDeniedDetail(v string) *QuerySmsAppIcpRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) SetCode(v string) *QuerySmsAppIcpRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) SetData(v []*QuerySmsAppIcpRecordResponseBodyData) *QuerySmsAppIcpRecordResponseBody {
	s.Data = v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) SetMessage(v string) *QuerySmsAppIcpRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) SetRequestId(v string) *QuerySmsAppIcpRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) SetSuccess(v bool) *QuerySmsAppIcpRecordResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBody) Validate() error {
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

type QuerySmsAppIcpRecordResponseBodyData struct {
	// The approval date.
	//
	// example:
	//
	// 2020-01-01
	AppApprovalDate *string `json:"AppApprovalDate,omitempty" xml:"AppApprovalDate,omitempty"`
	// The icp filing/license number.
	//
	// example:
	//
	// 123
	AppIcpLicenseNumber *string `json:"AppIcpLicenseNumber,omitempty" xml:"AppIcpLicenseNumber,omitempty"`
	// The ID of the APP-ICP record material.
	//
	// example:
	//
	// 51
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// The OSS fileKey for the APP-ICP record screenshot.
	//
	// example:
	//
	// 10000025*****02/ac181696-****-49c6-90dc-50689267aa00_mhsjd8b8_17*****662348.jpeg
	AppIcpRecordPic *string `json:"AppIcpRecordPic,omitempty" xml:"AppIcpRecordPic,omitempty"`
	// The URL of the APP-ICP record screenshot.
	//
	// example:
	//
	// https://alicom-fc-media.oss-cn-zhangjiakou.aliyuncs.com/100000****50802/afde****-496d-46e4-899d-b43758****8_mhk9oz0p_176224****542.png?Expires=1762****6&OSSAccessKeyId=bypFN****73PsLI&Signature=BygI9X****h7%2FXmFIo****FB2c%3D
	AppIcpRecordPicUrl *string `json:"AppIcpRecordPicUrl,omitempty" xml:"AppIcpRecordPicUrl,omitempty"`
	// The hosting unit name.
	//
	// example:
	//
	// 阿里云
	AppPrincipalUnitName *string `json:"AppPrincipalUnitName,omitempty" xml:"AppPrincipalUnitName,omitempty"`
	// APP实际运行截图Osskey
	//
	// example:
	//
	// 示例值示例值
	AppRuntimePic *string `json:"AppRuntimePic,omitempty" xml:"AppRuntimePic,omitempty"`
	// APP实际运行截图url地址
	//
	// example:
	//
	// 示例值示例值示例值
	AppRuntimePicUrl *string `json:"AppRuntimePicUrl,omitempty" xml:"AppRuntimePicUrl,omitempty"`
	// The app service name.
	//
	// example:
	//
	// 测试
	AppServiceName *string `json:"AppServiceName,omitempty" xml:"AppServiceName,omitempty"`
	// APP应用商店下载截图Osskey
	//
	// example:
	//
	// 示例值示例值
	AppStoreDownloadPic *string `json:"AppStoreDownloadPic,omitempty" xml:"AppStoreDownloadPic,omitempty"`
	// APP应用商店下载截图url地址
	//
	// example:
	//
	// 示例值示例值示例值
	AppStoreDownloadPicUrl *string `json:"AppStoreDownloadPicUrl,omitempty" xml:"AppStoreDownloadPicUrl,omitempty"`
	// The app store link.
	//
	// example:
	//
	// https://apps.apple.com/cn/app/阿里云/id981011420
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s QuerySmsAppIcpRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsAppIcpRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppApprovalDate() *string {
	return s.AppApprovalDate
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppIcpLicenseNumber() *string {
	return s.AppIcpLicenseNumber
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppIcpRecordPic() *string {
	return s.AppIcpRecordPic
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppIcpRecordPicUrl() *string {
	return s.AppIcpRecordPicUrl
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppPrincipalUnitName() *string {
	return s.AppPrincipalUnitName
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppRuntimePic() *string {
	return s.AppRuntimePic
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppRuntimePicUrl() *string {
	return s.AppRuntimePicUrl
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppServiceName() *string {
	return s.AppServiceName
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppStoreDownloadPic() *string {
	return s.AppStoreDownloadPic
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppStoreDownloadPicUrl() *string {
	return s.AppStoreDownloadPicUrl
}

func (s *QuerySmsAppIcpRecordResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppApprovalDate(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppApprovalDate = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppIcpLicenseNumber(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppIcpLicenseNumber = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppIcpRecordId(v int64) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppIcpRecordId = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppIcpRecordPic(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppIcpRecordPic = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppIcpRecordPicUrl(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppIcpRecordPicUrl = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppPrincipalUnitName(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppPrincipalUnitName = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppRuntimePic(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppRuntimePic = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppRuntimePicUrl(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppRuntimePicUrl = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppServiceName(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppServiceName = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppStoreDownloadPic(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppStoreDownloadPic = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppStoreDownloadPicUrl(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppStoreDownloadPicUrl = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetDomain(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.Domain = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
