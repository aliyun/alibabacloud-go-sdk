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
	SetAppServiceName(v string) *CreateSmsAppIcpRecordRequest
	GetAppServiceName() *string
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
	// 审核通过日期，示例2025-09-01
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-05-22
	AppApprovalDate *string `json:"AppApprovalDate,omitempty" xml:"AppApprovalDate,omitempty"`
	// ICP备案/许可证号
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AppIcpLicenseNumber *string `json:"AppIcpLicenseNumber,omitempty" xml:"AppIcpLicenseNumber,omitempty"`
	// app-icp备案详情截图osskey
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456/test1719383196031.jpg
	AppIcpRecordPic *string `json:"AppIcpRecordPic,omitempty" xml:"AppIcpRecordPic,omitempty"`
	// 主办单位名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AppPrincipalUnitName *string `json:"AppPrincipalUnitName,omitempty" xml:"AppPrincipalUnitName,omitempty"`
	// app服务名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AppServiceName *string `json:"AppServiceName,omitempty" xml:"AppServiceName,omitempty"`
	// APP应用商店链接
	//
	// This parameter is required.
	//
	// example:
	//
	// https://alicom-ops.alibaba-inc.com/dyorder/audit-domain/rule_manage
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

func (s *CreateSmsAppIcpRecordRequest) GetAppServiceName() *string {
	return s.AppServiceName
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

func (s *CreateSmsAppIcpRecordRequest) SetAppServiceName(v string) *CreateSmsAppIcpRecordRequest {
	s.AppServiceName = &v
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
