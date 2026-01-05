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
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QuerySmsAppIcpRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C9955E63-8BFF-101D-80A1-E6998DFEFF1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 审核通过日期，示例2025-09-01
	//
	// example:
	//
	// 2020-01-01
	AppApprovalDate *string `json:"AppApprovalDate,omitempty" xml:"AppApprovalDate,omitempty"`
	// ICP备案/许可证号
	//
	// example:
	//
	// 123
	AppIcpLicenseNumber *string `json:"AppIcpLicenseNumber,omitempty" xml:"AppIcpLicenseNumber,omitempty"`
	// app-icp备案材料id
	//
	// example:
	//
	// 51
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// app-icp备案截图图片Osskey（给签名传工单用）
	//
	// example:
	//
	// 10000025*****02/ac181696-****-49c6-90dc-50689267aa00_mhsjd8b8_17*****662348.jpeg
	AppIcpRecordPic *string `json:"AppIcpRecordPic,omitempty" xml:"AppIcpRecordPic,omitempty"`
	// app-icp备案截图url地址
	//
	// example:
	//
	// https://alicom-fc-media.oss-cn-zhangjiakou.aliyuncs.com/100000****50802/afde****-496d-46e4-899d-b43758****8_mhk9oz0p_176224****542.png?Expires=1762****6&OSSAccessKeyId=bypFN****73PsLI&Signature=BygI9X****h7%2FXmFIo****FB2c%3D
	AppIcpRecordPicUrl *string `json:"AppIcpRecordPicUrl,omitempty" xml:"AppIcpRecordPicUrl,omitempty"`
	// 主办单位名称
	//
	// example:
	//
	// 示例值示例值
	AppPrincipalUnitName *string `json:"AppPrincipalUnitName,omitempty" xml:"AppPrincipalUnitName,omitempty"`
	// app服务名称
	//
	// example:
	//
	// 示例值
	AppServiceName *string `json:"AppServiceName,omitempty" xml:"AppServiceName,omitempty"`
	// APP应用商店链接
	//
	// example:
	//
	// https://test
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

func (s *QuerySmsAppIcpRecordResponseBodyData) GetAppServiceName() *string {
	return s.AppServiceName
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

func (s *QuerySmsAppIcpRecordResponseBodyData) SetAppServiceName(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.AppServiceName = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) SetDomain(v string) *QuerySmsAppIcpRecordResponseBodyData {
	s.Domain = &v
	return s
}

func (s *QuerySmsAppIcpRecordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
