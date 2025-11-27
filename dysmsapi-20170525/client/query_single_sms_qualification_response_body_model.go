// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySingleSmsQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuerySingleSmsQualificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuerySingleSmsQualificationResponseBody
	GetCode() *string
	SetData(v *QuerySingleSmsQualificationResponseBodyData) *QuerySingleSmsQualificationResponseBody
	GetData() *QuerySingleSmsQualificationResponseBodyData
	SetMessage(v string) *QuerySingleSmsQualificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySingleSmsQualificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySingleSmsQualificationResponseBody
	GetSuccess() *bool
}

type QuerySingleSmsQualificationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySingleSmsQualificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySingleSmsQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuerySingleSmsQualificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySingleSmsQualificationResponseBody) GetData() *QuerySingleSmsQualificationResponseBodyData {
	return s.Data
}

func (s *QuerySingleSmsQualificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySingleSmsQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySingleSmsQualificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySingleSmsQualificationResponseBody) SetAccessDeniedDetail(v string) *QuerySingleSmsQualificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) SetCode(v string) *QuerySingleSmsQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) SetData(v *QuerySingleSmsQualificationResponseBodyData) *QuerySingleSmsQualificationResponseBody {
	s.Data = v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) SetMessage(v string) *QuerySingleSmsQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) SetRequestId(v string) *QuerySingleSmsQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) SetSuccess(v bool) *QuerySingleSmsQualificationResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySingleSmsQualificationResponseBodyData struct {
	// 经办人身份证有效期
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	AdminIDCardExpDate *string `json:"AdminIDCardExpDate,omitempty" xml:"AdminIDCardExpDate,omitempty"`
	// 经办人身份证国徽面，产品需求，要求身份证可以分正反面上传
	//
	// example:
	//
	// https://******.aliyuncs.com/******
	AdminIDCardFrontFace *string `json:"AdminIDCardFrontFace,omitempty" xml:"AdminIDCardFrontFace,omitempty"`
	// 经办人身份证号码
	//
	// example:
	//
	// 511391********5123
	AdminIDCardNo *string `json:"AdminIDCardNo,omitempty" xml:"AdminIDCardNo,omitempty"`
	// 经办人身份证图片地址，正反面合一
	//
	// example:
	//
	// https://******.aliyuncs.com/******
	AdminIDCardPic *string `json:"AdminIDCardPic,omitempty" xml:"AdminIDCardPic,omitempty"`
	// 管理员身份证类型
	//
	// example:
	//
	// identityCard
	AdminIDCardType *string `json:"AdminIDCardType,omitempty" xml:"AdminIDCardType,omitempty"`
	// 经办人姓名
	//
	// example:
	//
	// 示例值示例值
	AdminName *string `json:"AdminName,omitempty" xml:"AdminName,omitempty"`
	// 经办人手机号码
	//
	// example:
	//
	// 137*******
	AdminPhoneNo *string `json:"AdminPhoneNo,omitempty" xml:"AdminPhoneNo,omitempty"`
	// 证件信息
	BusinessLicensePics []*QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics `json:"BusinessLicensePics,omitempty" xml:"BusinessLicensePics,omitempty" type:"Repeated"`
	// 行业类型，在当前模式下是可以用产品线code来区分
	//
	// example:
	//
	// dysms
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 公司名称
	//
	// example:
	//
	// 示例值示例值
	CompanyName *string `json:"CompanyName,omitempty" xml:"CompanyName,omitempty"`
	// 企业类型, COMPANY:公司，政府或者事业单位:NON_PROFIT_ORGANIZATION
	//
	// example:
	//
	// COMPANY
	CompanyType *string `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	// example:
	//
	// 2023-01-01~2033-01-01
	EffTimeStr *string `json:"EffTimeStr,omitempty" xml:"EffTimeStr,omitempty"`
	// 法人身份证号码
	//
	// example:
	//
	// 511391********5123
	LegalPersonIDCardNo *string `json:"LegalPersonIDCardNo,omitempty" xml:"LegalPersonIDCardNo,omitempty"`
	// 法人身份证类型
	//
	// example:
	//
	// identityCard
	LegalPersonIDCardType *string `json:"LegalPersonIDCardType,omitempty" xml:"LegalPersonIDCardType,omitempty"`
	// 法人身份证有效期
	//
	// example:
	//
	// 2023-01-01~2033-01-01
	LegalPersonIdCardEffTime *string `json:"LegalPersonIdCardEffTime,omitempty" xml:"LegalPersonIdCardEffTime,omitempty"`
	// 法人姓名
	//
	// example:
	//
	// 示例值
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// 社会统一信用代码
	//
	// example:
	//
	// 910X********0012
	OrganizationCode *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	// 更多资料
	OtherFiles []*QuerySingleSmsQualificationResponseBodyDataOtherFiles `json:"OtherFiles,omitempty" xml:"OtherFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 10000****
	QualificationGroupId *int64 `json:"QualificationGroupId,omitempty" xml:"QualificationGroupId,omitempty"`
	// 资质名称
	//
	// example:
	//
	// 示例值示例值
	QualificationName *string `json:"QualificationName,omitempty" xml:"QualificationName,omitempty"`
	// 备注
	//
	// example:
	//
	// 示例值示例值
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 当前审核状态
	//
	// example:
	//
	// PASSED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// 是否自用
	//
	// example:
	//
	// false
	UseBySelf *bool `json:"UseBySelf,omitempty" xml:"UseBySelf,omitempty"`
	// example:
	//
	// false
	WhetherShare *bool `json:"WhetherShare,omitempty" xml:"WhetherShare,omitempty"`
	// 乾坤袋工单ID
	//
	// example:
	//
	// 2001****
	WorkOrderId *int64 `json:"WorkOrderId,omitempty" xml:"WorkOrderId,omitempty"`
}

func (s QuerySingleSmsQualificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminIDCardExpDate() *string {
	return s.AdminIDCardExpDate
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminIDCardFrontFace() *string {
	return s.AdminIDCardFrontFace
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminIDCardNo() *string {
	return s.AdminIDCardNo
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminIDCardPic() *string {
	return s.AdminIDCardPic
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminIDCardType() *string {
	return s.AdminIDCardType
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminName() *string {
	return s.AdminName
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetAdminPhoneNo() *string {
	return s.AdminPhoneNo
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetBusinessLicensePics() []*QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics {
	return s.BusinessLicensePics
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetBusinessType() *string {
	return s.BusinessType
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetCompanyName() *string {
	return s.CompanyName
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetCompanyType() *string {
	return s.CompanyType
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetEffTimeStr() *string {
	return s.EffTimeStr
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetLegalPersonIDCardNo() *string {
	return s.LegalPersonIDCardNo
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetLegalPersonIDCardType() *string {
	return s.LegalPersonIDCardType
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetLegalPersonIdCardEffTime() *string {
	return s.LegalPersonIdCardEffTime
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetOtherFiles() []*QuerySingleSmsQualificationResponseBodyDataOtherFiles {
	return s.OtherFiles
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetQualificationGroupId() *int64 {
	return s.QualificationGroupId
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetQualificationName() *string {
	return s.QualificationName
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetState() *string {
	return s.State
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetUseBySelf() *bool {
	return s.UseBySelf
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetWhetherShare() *bool {
	return s.WhetherShare
}

func (s *QuerySingleSmsQualificationResponseBodyData) GetWorkOrderId() *int64 {
	return s.WorkOrderId
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminIDCardExpDate(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminIDCardExpDate = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminIDCardFrontFace(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminIDCardFrontFace = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminIDCardNo(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminIDCardNo = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminIDCardPic(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminIDCardPic = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminIDCardType(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminIDCardType = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminName(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminName = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetAdminPhoneNo(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.AdminPhoneNo = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetBusinessLicensePics(v []*QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) *QuerySingleSmsQualificationResponseBodyData {
	s.BusinessLicensePics = v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetBusinessType(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetCompanyName(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.CompanyName = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetCompanyType(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.CompanyType = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetEffTimeStr(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.EffTimeStr = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetLegalPersonIDCardNo(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.LegalPersonIDCardNo = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetLegalPersonIDCardType(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.LegalPersonIDCardType = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetLegalPersonIdCardEffTime(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.LegalPersonIdCardEffTime = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetLegalPersonName(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.LegalPersonName = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetOrganizationCode(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.OrganizationCode = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetOtherFiles(v []*QuerySingleSmsQualificationResponseBodyDataOtherFiles) *QuerySingleSmsQualificationResponseBodyData {
	s.OtherFiles = v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetQualificationGroupId(v int64) *QuerySingleSmsQualificationResponseBodyData {
	s.QualificationGroupId = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetQualificationName(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.QualificationName = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetRemark(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.Remark = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetState(v string) *QuerySingleSmsQualificationResponseBodyData {
	s.State = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetUseBySelf(v bool) *QuerySingleSmsQualificationResponseBodyData {
	s.UseBySelf = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetWhetherShare(v bool) *QuerySingleSmsQualificationResponseBodyData {
	s.WhetherShare = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) SetWorkOrderId(v int64) *QuerySingleSmsQualificationResponseBodyData {
	s.WorkOrderId = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyData) Validate() error {
	if s.BusinessLicensePics != nil {
		for _, item := range s.BusinessLicensePics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OtherFiles != nil {
		for _, item := range s.OtherFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics struct {
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
	// 文件的完整路径
	//
	// example:
	//
	// https://******.aliyuncs.com/******
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	// example:
	//
	// businessLicense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) GetLicensePic() *string {
	return s.LicensePic
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) GetPicUrl() *string {
	return s.PicUrl
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) GetType() *string {
	return s.Type
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) SetLicensePic(v string) *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics {
	s.LicensePic = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) SetPicUrl(v string) *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics {
	s.PicUrl = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) SetType(v string) *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics {
	s.Type = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyDataBusinessLicensePics) Validate() error {
	return dara.Validate(s)
}

type QuerySingleSmsQualificationResponseBodyDataOtherFiles struct {
	// example:
	//
	// 123456/111.png
	LicensePic *string `json:"LicensePic,omitempty" xml:"LicensePic,omitempty"`
	// 文件的完整路径
	//
	// example:
	//
	// https://******.aliyuncs.com/******
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s QuerySingleSmsQualificationResponseBodyDataOtherFiles) String() string {
	return dara.Prettify(s)
}

func (s QuerySingleSmsQualificationResponseBodyDataOtherFiles) GoString() string {
	return s.String()
}

func (s *QuerySingleSmsQualificationResponseBodyDataOtherFiles) GetLicensePic() *string {
	return s.LicensePic
}

func (s *QuerySingleSmsQualificationResponseBodyDataOtherFiles) GetPicUrl() *string {
	return s.PicUrl
}

func (s *QuerySingleSmsQualificationResponseBodyDataOtherFiles) SetLicensePic(v string) *QuerySingleSmsQualificationResponseBodyDataOtherFiles {
	s.LicensePic = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyDataOtherFiles) SetPicUrl(v string) *QuerySingleSmsQualificationResponseBodyDataOtherFiles {
	s.PicUrl = &v
	return s
}

func (s *QuerySingleSmsQualificationResponseBodyDataOtherFiles) Validate() error {
	return dara.Validate(s)
}
