// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplyScene(v string) *GetSmsTemplateResponseBody
	GetApplyScene() *string
	SetAuditInfo(v *GetSmsTemplateResponseBodyAuditInfo) *GetSmsTemplateResponseBody
	GetAuditInfo() *GetSmsTemplateResponseBodyAuditInfo
	SetCode(v string) *GetSmsTemplateResponseBody
	GetCode() *string
	SetCreateDate(v string) *GetSmsTemplateResponseBody
	GetCreateDate() *string
	SetFileUrlList(v *GetSmsTemplateResponseBodyFileUrlList) *GetSmsTemplateResponseBody
	GetFileUrlList() *GetSmsTemplateResponseBodyFileUrlList
	SetIntlType(v int32) *GetSmsTemplateResponseBody
	GetIntlType() *int32
	SetMessage(v string) *GetSmsTemplateResponseBody
	GetMessage() *string
	SetMoreDataFileUrlList(v *GetSmsTemplateResponseBodyMoreDataFileUrlList) *GetSmsTemplateResponseBody
	GetMoreDataFileUrlList() *GetSmsTemplateResponseBodyMoreDataFileUrlList
	SetOrderId(v string) *GetSmsTemplateResponseBody
	GetOrderId() *string
	SetRelatedSignName(v string) *GetSmsTemplateResponseBody
	GetRelatedSignName() *string
	SetRemark(v string) *GetSmsTemplateResponseBody
	GetRemark() *string
	SetRequestId(v string) *GetSmsTemplateResponseBody
	GetRequestId() *string
	SetSignList(v *GetSmsTemplateResponseBodySignList) *GetSmsTemplateResponseBody
	GetSignList() *GetSmsTemplateResponseBodySignList
	SetTemplateCode(v string) *GetSmsTemplateResponseBody
	GetTemplateCode() *string
	SetTemplateContent(v string) *GetSmsTemplateResponseBody
	GetTemplateContent() *string
	SetTemplateName(v string) *GetSmsTemplateResponseBody
	GetTemplateName() *string
	SetTemplateStatus(v string) *GetSmsTemplateResponseBody
	GetTemplateStatus() *string
	SetTemplateTag(v int32) *GetSmsTemplateResponseBody
	GetTemplateTag() *int32
	SetTemplateType(v string) *GetSmsTemplateResponseBody
	GetTemplateType() *string
	SetVariableAttribute(v string) *GetSmsTemplateResponseBody
	GetVariableAttribute() *string
	SetVendorAuditStatus(v map[string]interface{}) *GetSmsTemplateResponseBody
	GetVendorAuditStatus() map[string]interface{}
}

type GetSmsTemplateResponseBody struct {
	// 应用场景内容。
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplyScene *string `json:"ApplyScene,omitempty" xml:"ApplyScene,omitempty"`
	// 审核信息。
	AuditInfo *GetSmsTemplateResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// 请求状态码。取值：
	//
	// 	- OK：代表请求成功。
	//
	// 	- 其他错误码，请参见[API错误码](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 创建短信模板的时间。
	//
	// example:
	//
	// 2024-06-03 10:02:34
	CreateDate  *string                                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FileUrlList *GetSmsTemplateResponseBodyFileUrlList `json:"FileUrlList,omitempty" xml:"FileUrlList,omitempty" type:"Struct"`
	// 国际/港澳台模板类型。当**TemplateType**参数返回值为**3**时，此参数取值：
	//
	// - **0**：短信通知。
	//
	// - **1**：推广短信。
	//
	// - **2**：验证码。
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message             *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	MoreDataFileUrlList *GetSmsTemplateResponseBodyMoreDataFileUrlList `json:"MoreDataFileUrlList,omitempty" xml:"MoreDataFileUrlList,omitempty" type:"Struct"`
	// 工单号。
	//
	// 审核人员查询审核时会用到此参数。您需要审核加急时需要提供此工单号。
	//
	// example:
	//
	// 2003019****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 申请模板时，关联的短信签名。
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// 短信模板申请说明，是模板审核的参考信息之一。
	//
	// example:
	//
	// 申请验证码模板
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 本次调用请求的ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignList  *GetSmsTemplateResponseBodySignList `json:"SignList,omitempty" xml:"SignList,omitempty" type:"Struct"`
	// 短信模板Code。
	//
	// example:
	//
	// SMS_20375****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 短信模板内容。
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// 短信模板名称。
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板审核状态。返回值：
	//
	// - **0**：审核中。
	//
	// - **1**：通过审核。
	//
	// - **2**：未通过审核，会返回审核失败的原因，请参考[短信审核失败的处理建议](https://www.alibabacloud.com/help/en/sms/user-guide/causes-of-application-failures-and-suggestions)，调用[UpdateSmsTemplate](https://www.alibabacloud.com/help/en/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmstemplate)接口或在[模板管理](https://dysms.console.aliyun.com/domestic/text/template)页面修改短信模板。
	//
	// - **10**：取消审核。
	//
	// example:
	//
	// 2
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// 模板标识。取值：
	//
	// - 2：用户自定义创建模板。
	//
	// - 3：系统赠送模板。
	//
	// - 4：测试模板。
	//
	// example:
	//
	// 2
	TemplateTag *int32 `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty"`
	// 短信类型。取值：
	//
	// - **0**：验证码。
	//
	// - **1**：短信通知。
	//
	// - **2**：推广短信。
	//
	// - **3**：国际/港澳台消息。
	//
	// > 仅支持企业认证用户申请推广短信和国际/港澳台消息。个人用户与企业用户权益区别详情请参见[使用须知](https://www.alibabacloud.com/help/en/sms/user-guide/usage-notes)。
	//
	// example:
	//
	// 0
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 模板变量规则。
	//
	// 模板变量规则详情，请参见[示例文档](https://www.alibabacloud.com/help/en/sms/templaterule-template-variable-parameter-filling-example)。
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	VariableAttribute *string `json:"VariableAttribute,omitempty" xml:"VariableAttribute,omitempty"`
	// 各运营商审核状态，仅数字短信会返回该参数。
	//
	//
	// key代表运营商。取值：
	//
	// - MOBILE_VENDOR：中国移动。
	//
	// - TELECOM_VENDOR：中国电信。
	//
	// - UNICOM_VENDOR：中国联通。
	//
	//  value代表审核状态。取值：
	//
	// - 0：审核中。
	//
	// - 1：通过。
	//
	//  - 2：不通过。
	//
	//  - 15：已失效。
	//
	// example:
	//
	// {
	//
	//     "MOBILE_VENDOR": 0,
	//
	//     "TELCOM_VENDOR": 0,
	//
	//     "UNICOM_VENDOR": 0
	//
	//   }
	VendorAuditStatus map[string]interface{} `json:"VendorAuditStatus,omitempty" xml:"VendorAuditStatus,omitempty"`
}

func (s GetSmsTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBody) GetApplyScene() *string {
	return s.ApplyScene
}

func (s *GetSmsTemplateResponseBody) GetAuditInfo() *GetSmsTemplateResponseBodyAuditInfo {
	return s.AuditInfo
}

func (s *GetSmsTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsTemplateResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetSmsTemplateResponseBody) GetFileUrlList() *GetSmsTemplateResponseBodyFileUrlList {
	return s.FileUrlList
}

func (s *GetSmsTemplateResponseBody) GetIntlType() *int32 {
	return s.IntlType
}

func (s *GetSmsTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsTemplateResponseBody) GetMoreDataFileUrlList() *GetSmsTemplateResponseBodyMoreDataFileUrlList {
	return s.MoreDataFileUrlList
}

func (s *GetSmsTemplateResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *GetSmsTemplateResponseBody) GetRelatedSignName() *string {
	return s.RelatedSignName
}

func (s *GetSmsTemplateResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *GetSmsTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsTemplateResponseBody) GetSignList() *GetSmsTemplateResponseBodySignList {
	return s.SignList
}

func (s *GetSmsTemplateResponseBody) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetSmsTemplateResponseBody) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetSmsTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetSmsTemplateResponseBody) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *GetSmsTemplateResponseBody) GetTemplateTag() *int32 {
	return s.TemplateTag
}

func (s *GetSmsTemplateResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetSmsTemplateResponseBody) GetVariableAttribute() *string {
	return s.VariableAttribute
}

func (s *GetSmsTemplateResponseBody) GetVendorAuditStatus() map[string]interface{} {
	return s.VendorAuditStatus
}

func (s *GetSmsTemplateResponseBody) SetApplyScene(v string) *GetSmsTemplateResponseBody {
	s.ApplyScene = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetAuditInfo(v *GetSmsTemplateResponseBodyAuditInfo) *GetSmsTemplateResponseBody {
	s.AuditInfo = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetCode(v string) *GetSmsTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetCreateDate(v string) *GetSmsTemplateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetFileUrlList(v *GetSmsTemplateResponseBodyFileUrlList) *GetSmsTemplateResponseBody {
	s.FileUrlList = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetIntlType(v int32) *GetSmsTemplateResponseBody {
	s.IntlType = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetMessage(v string) *GetSmsTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetMoreDataFileUrlList(v *GetSmsTemplateResponseBodyMoreDataFileUrlList) *GetSmsTemplateResponseBody {
	s.MoreDataFileUrlList = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetOrderId(v string) *GetSmsTemplateResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRelatedSignName(v string) *GetSmsTemplateResponseBody {
	s.RelatedSignName = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRemark(v string) *GetSmsTemplateResponseBody {
	s.Remark = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetRequestId(v string) *GetSmsTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetSignList(v *GetSmsTemplateResponseBodySignList) *GetSmsTemplateResponseBody {
	s.SignList = v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateCode(v string) *GetSmsTemplateResponseBody {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateContent(v string) *GetSmsTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateName(v string) *GetSmsTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateStatus(v string) *GetSmsTemplateResponseBody {
	s.TemplateStatus = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateTag(v int32) *GetSmsTemplateResponseBody {
	s.TemplateTag = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetTemplateType(v string) *GetSmsTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetVariableAttribute(v string) *GetSmsTemplateResponseBody {
	s.VariableAttribute = &v
	return s
}

func (s *GetSmsTemplateResponseBody) SetVendorAuditStatus(v map[string]interface{}) *GetSmsTemplateResponseBody {
	s.VendorAuditStatus = v
	return s
}

func (s *GetSmsTemplateResponseBody) Validate() error {
	if s.AuditInfo != nil {
		if err := s.AuditInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FileUrlList != nil {
		if err := s.FileUrlList.Validate(); err != nil {
			return err
		}
	}
	if s.MoreDataFileUrlList != nil {
		if err := s.MoreDataFileUrlList.Validate(); err != nil {
			return err
		}
	}
	if s.SignList != nil {
		if err := s.SignList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmsTemplateResponseBodyAuditInfo struct {
	// 审核时间。
	//
	// example:
	//
	// 2024-06-03 11:20:34
	AuditDate *string `json:"AuditDate,omitempty" xml:"AuditDate,omitempty"`
	// 审核未通过的原因。
	//
	// example:
	//
	// 模板内容中包含错别字。
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
}

func (s GetSmsTemplateResponseBodyAuditInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyAuditInfo) GetAuditDate() *string {
	return s.AuditDate
}

func (s *GetSmsTemplateResponseBodyAuditInfo) GetRejectInfo() *string {
	return s.RejectInfo
}

func (s *GetSmsTemplateResponseBodyAuditInfo) SetAuditDate(v string) *GetSmsTemplateResponseBodyAuditInfo {
	s.AuditDate = &v
	return s
}

func (s *GetSmsTemplateResponseBodyAuditInfo) SetRejectInfo(v string) *GetSmsTemplateResponseBodyAuditInfo {
	s.RejectInfo = &v
	return s
}

func (s *GetSmsTemplateResponseBodyAuditInfo) Validate() error {
	return dara.Validate(s)
}

type GetSmsTemplateResponseBodyFileUrlList struct {
	FileUrl []*string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateResponseBodyFileUrlList) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponseBodyFileUrlList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyFileUrlList) GetFileUrl() []*string {
	return s.FileUrl
}

func (s *GetSmsTemplateResponseBodyFileUrlList) SetFileUrl(v []*string) *GetSmsTemplateResponseBodyFileUrlList {
	s.FileUrl = v
	return s
}

func (s *GetSmsTemplateResponseBodyFileUrlList) Validate() error {
	return dara.Validate(s)
}

type GetSmsTemplateResponseBodyMoreDataFileUrlList struct {
	MoreDataFileUrl []*string `json:"MoreDataFileUrl,omitempty" xml:"MoreDataFileUrl,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateResponseBodyMoreDataFileUrlList) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponseBodyMoreDataFileUrlList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodyMoreDataFileUrlList) GetMoreDataFileUrl() []*string {
	return s.MoreDataFileUrl
}

func (s *GetSmsTemplateResponseBodyMoreDataFileUrlList) SetMoreDataFileUrl(v []*string) *GetSmsTemplateResponseBodyMoreDataFileUrlList {
	s.MoreDataFileUrl = v
	return s
}

func (s *GetSmsTemplateResponseBodyMoreDataFileUrlList) Validate() error {
	return dara.Validate(s)
}

type GetSmsTemplateResponseBodySignList struct {
	SignList []*string `json:"SignList,omitempty" xml:"SignList,omitempty" type:"Repeated"`
}

func (s GetSmsTemplateResponseBodySignList) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateResponseBodySignList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateResponseBodySignList) GetSignList() []*string {
	return s.SignList
}

func (s *GetSmsTemplateResponseBodySignList) SetSignList(v []*string) *GetSmsTemplateResponseBodySignList {
	s.SignList = v
	return s
}

func (s *GetSmsTemplateResponseBodySignList) Validate() error {
	return dara.Validate(s)
}
