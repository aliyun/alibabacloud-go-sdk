// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAppIcpRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppApprovalDate(v string) *CreateSmsAppIcpRecordRequest
	GetAppApprovalDate() *string
	SetAppIcpLicenseNumber(v string) *CreateSmsAppIcpRecordRequest
	GetAppIcpLicenseNumber() *string
	SetAppIcpRecordPic(v string) *CreateSmsAppIcpRecordRequest
	GetAppIcpRecordPic() *string
	SetAppPrincipalUnitName(v string) *CreateSmsAppIcpRecordRequest
	GetAppPrincipalUnitName() *string
	SetAppRuntimePic(v string) *CreateSmsAppIcpRecordRequest
	GetAppRuntimePic() *string
	SetAppServiceName(v string) *CreateSmsAppIcpRecordRequest
	GetAppServiceName() *string
	SetAppStoreDownloadPic(v string) *CreateSmsAppIcpRecordRequest
	GetAppStoreDownloadPic() *string
	SetDomain(v string) *CreateSmsAppIcpRecordRequest
	GetDomain() *string
	SetOwnerId(v int64) *CreateSmsAppIcpRecordRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateSmsAppIcpRecordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsAppIcpRecordRequest
	GetResourceOwnerId() *int64
}

type CreateSmsAppIcpRecordRequest struct {
	// The ICP filing approval date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-05-22
	AppApprovalDate *string `json:"AppApprovalDate,omitempty" xml:"AppApprovalDate,omitempty"`
	// The ICP record/license number. The number must not exceed 15 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 浙B2-20080101
	AppIcpLicenseNumber *string `json:"AppIcpLicenseNumber,omitempty" xml:"AppIcpLicenseNumber,omitempty"`
	// The fileKey for the screenshot of your app\\"s ICP filing details.
	//
	// 1. To look up your ICP filing, go to the [MIIT service platform](https://beian.miit.gov.cn/#/Integrated/recordQuery), search for your filing, and open its details page.
	//
	// 2. The ICP filing screenshot must be uploaded to OSS and meet the following requirements:
	//
	// - The filename cannot contain Chinese characters or special characters.
	//
	// - The file must be an image in `jpg`, `png`, `gif`, or `jpeg` format. The file size cannot exceed 5 MB.
	//
	// - The screenshot must show the full URL.
	//
	// - For **Record Type**, select "APP".
	//
	// - The **principal unit name*	- must be identical to the company or institution name on the qualification documents associated with the signature.
	//
	// - The screenshot must clearly show the complete ICP record/license number.
	//
	// - The **service name*	- must be identical to the **signature name**.
	//
	// 3. To obtain the fileKey, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	AppIcpRecordPic *string `json:"AppIcpRecordPic,omitempty" xml:"AppIcpRecordPic,omitempty"`
	// The principal unit name from your ICP filing. The name must not exceed 50 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云计算有限公司
	AppPrincipalUnitName *string `json:"AppPrincipalUnitName,omitempty" xml:"AppPrincipalUnitName,omitempty"`
	// APP实际运行截图osskey
	//
	// example:
	//
	// 示例值示例值示例值
	AppRuntimePic *string `json:"AppRuntimePic,omitempty" xml:"AppRuntimePic,omitempty"`
	// The service name from your ICP filing. The name must not exceed 15 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	AppServiceName *string `json:"AppServiceName,omitempty" xml:"AppServiceName,omitempty"`
	// APP应用商店下载截图osskey
	//
	// example:
	//
	// 示例值示例值示例值
	AppStoreDownloadPic *string `json:"AppStoreDownloadPic,omitempty" xml:"AppStoreDownloadPic,omitempty"`
	// The app store link.
	//
	// > - The link must start with `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://apps.apple.com/cn/app/阿里云/id981011420
	Domain               *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSmsAppIcpRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAppIcpRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsAppIcpRecordRequest) GetAppApprovalDate() *string {
	return s.AppApprovalDate
}

func (s *CreateSmsAppIcpRecordRequest) GetAppIcpLicenseNumber() *string {
	return s.AppIcpLicenseNumber
}

func (s *CreateSmsAppIcpRecordRequest) GetAppIcpRecordPic() *string {
	return s.AppIcpRecordPic
}

func (s *CreateSmsAppIcpRecordRequest) GetAppPrincipalUnitName() *string {
	return s.AppPrincipalUnitName
}

func (s *CreateSmsAppIcpRecordRequest) GetAppRuntimePic() *string {
	return s.AppRuntimePic
}

func (s *CreateSmsAppIcpRecordRequest) GetAppServiceName() *string {
	return s.AppServiceName
}

func (s *CreateSmsAppIcpRecordRequest) GetAppStoreDownloadPic() *string {
	return s.AppStoreDownloadPic
}

func (s *CreateSmsAppIcpRecordRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateSmsAppIcpRecordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsAppIcpRecordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsAppIcpRecordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsAppIcpRecordRequest) SetAppApprovalDate(v string) *CreateSmsAppIcpRecordRequest {
	s.AppApprovalDate = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppIcpLicenseNumber(v string) *CreateSmsAppIcpRecordRequest {
	s.AppIcpLicenseNumber = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppIcpRecordPic(v string) *CreateSmsAppIcpRecordRequest {
	s.AppIcpRecordPic = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppPrincipalUnitName(v string) *CreateSmsAppIcpRecordRequest {
	s.AppPrincipalUnitName = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppRuntimePic(v string) *CreateSmsAppIcpRecordRequest {
	s.AppRuntimePic = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppServiceName(v string) *CreateSmsAppIcpRecordRequest {
	s.AppServiceName = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetAppStoreDownloadPic(v string) *CreateSmsAppIcpRecordRequest {
	s.AppStoreDownloadPic = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetDomain(v string) *CreateSmsAppIcpRecordRequest {
	s.Domain = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetOwnerId(v int64) *CreateSmsAppIcpRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetResourceOwnerAccount(v string) *CreateSmsAppIcpRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) SetResourceOwnerId(v int64) *CreateSmsAppIcpRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsAppIcpRecordRequest) Validate() error {
	return dara.Validate(s)
}
