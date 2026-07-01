// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordId(v int64) *GetSmsSignResponseBody
	GetAppIcpRecordId() *int64
	SetApplyScene(v string) *GetSmsSignResponseBody
	GetApplyScene() *string
	SetAuditInfo(v *GetSmsSignResponseBodyAuditInfo) *GetSmsSignResponseBody
	GetAuditInfo() *GetSmsSignResponseBodyAuditInfo
	SetAuthorizationLetterAuditPass(v bool) *GetSmsSignResponseBody
	GetAuthorizationLetterAuditPass() *bool
	SetAuthorizationLetterId(v int64) *GetSmsSignResponseBody
	GetAuthorizationLetterId() *int64
	SetCode(v string) *GetSmsSignResponseBody
	GetCode() *string
	SetCreateDate(v string) *GetSmsSignResponseBody
	GetCreateDate() *string
	SetFileUrlList(v []*string) *GetSmsSignResponseBody
	GetFileUrlList() []*string
	SetMessage(v string) *GetSmsSignResponseBody
	GetMessage() *string
	SetOrderId(v string) *GetSmsSignResponseBody
	GetOrderId() *string
	SetQualificationId(v int64) *GetSmsSignResponseBody
	GetQualificationId() *int64
	SetRegisterResult(v int32) *GetSmsSignResponseBody
	GetRegisterResult() *int32
	SetRemark(v string) *GetSmsSignResponseBody
	GetRemark() *string
	SetRequestId(v string) *GetSmsSignResponseBody
	GetRequestId() *string
	SetSignCode(v string) *GetSmsSignResponseBody
	GetSignCode() *string
	SetSignIspRegisterDetailList(v []*GetSmsSignResponseBodySignIspRegisterDetailList) *GetSmsSignResponseBody
	GetSignIspRegisterDetailList() []*GetSmsSignResponseBodySignIspRegisterDetailList
	SetSignName(v string) *GetSmsSignResponseBody
	GetSignName() *string
	SetSignStatus(v int64) *GetSmsSignResponseBody
	GetSignStatus() *int64
	SetSignTag(v string) *GetSmsSignResponseBody
	GetSignTag() *string
	SetSignUsage(v string) *GetSmsSignResponseBody
	GetSignUsage() *string
	SetThirdParty(v bool) *GetSmsSignResponseBody
	GetThirdParty() *bool
	SetTrademarkId(v int64) *GetSmsSignResponseBody
	GetTrademarkId() *int64
}

type GetSmsSignResponseBody struct {
	// APP-ICP备案实体id。
	//
	// example:
	//
	// 1000009***123
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// 应用场景内容。
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplyScene *string `json:"ApplyScene,omitempty" xml:"ApplyScene,omitempty"`
	// 审核信息。
	AuditInfo *GetSmsSignResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// 委托授权书审核状态。取值：
	//
	// - true：审核通过。
	//
	// - false：审核未通过（包含审核通过外的其他所有状态）。
	//
	// example:
	//
	// true
	AuthorizationLetterAuditPass *bool `json:"AuthorizationLetterAuditPass,omitempty" xml:"AuthorizationLetterAuditPass,omitempty"`
	// 委托授权书ID。
	//
	// example:
	//
	// 1000********1234
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// 请求状态码。取值：
	//
	// - OK：代表请求成功。
	//
	// - 其他错误码，请参见[API错误码](https://help.aliyun.com/document_detail/101346.html)。
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 短信签名的创建日期和时间。
	//
	// example:
	//
	// 2024-06-03 10:02:34
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// 更多资料信息，补充上传业务证明文件或业务截图文件列表。
	FileUrlList []*string `json:"FileUrlList,omitempty" xml:"FileUrlList,omitempty" type:"Repeated"`
	// 状态码的描述。
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 工单号。
	//
	// 审核人员查询审核时会用到此参数。您需要审核加急时需要提供此工单号。
	//
	// example:
	//
	// 20044156924
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// 资质ID。申请签名时关联的资质ID。
	//
	// example:
	//
	// 2004393****
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// **已废弃，请使用`SignIspRegisterDetailList`查看各运营商实名报备结果。**
	//
	// 签名实名制报备结果。取值：
	//
	// - 0：报备失败。
	//
	// - 1：报备成功。
	//
	// - 2：报备失效。
	//
	// - -1：无状态。
	//
	// 建议您单击查看[更多签名实名制报备内容及建议操作](https://help.aliyun.com/document_detail/2873145.html)。
	//
	// example:
	//
	// 1
	RegisterResult *int32 `json:"RegisterResult,omitempty" xml:"RegisterResult,omitempty"`
	// 短信签名场景说明，长度不超过200个字符。
	//
	// example:
	//
	// 登录场景验证码
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 本次调用请求的ID，是由阿里云为该请求生成的唯一标识符，可用于排查和定位问题。
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 短信签名Code。
	//
	// example:
	//
	// SIGN_100000077042023_17174665*****_ZM2kG
	SignCode *string `json:"SignCode,omitempty" xml:"SignCode,omitempty"`
	// 运营商报备状态列表。获取此参数返回数据需要[更新SDK](https://api.aliyun.com/api-tools/sdk/Dysmsapi?version=2017-05-25&language=java-tea&tab=primer-doc)至4.1.2版本或以上。
	SignIspRegisterDetailList []*GetSmsSignResponseBodySignIspRegisterDetailList `json:"SignIspRegisterDetailList,omitempty" xml:"SignIspRegisterDetailList,omitempty" type:"Repeated"`
	// 短信签名名称。
	//
	// example:
	//
	// 登录验证
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 签名审核状态。取值：
	//
	// - **0**：审核中。
	//
	// - **1**：审核通过。
	//
	// - **2**：审核失败，请在返回参数`AuditInfo.RejectInfo`中查看审核失败原因。
	//
	// - **10**：取消审核。
	//
	// example:
	//
	// 2
	SignStatus *int64 `json:"SignStatus,omitempty" xml:"SignStatus,omitempty"`
	// 签名标识。取值：
	//
	// - 2：用户自定义创建签名。
	//
	// - 3：系统赠送签名。
	//
	// - 4：测试签名。
	//
	// - 5：试用签名。
	//
	// example:
	//
	// 2
	SignTag *string `json:"SignTag,omitempty" xml:"SignTag,omitempty"`
	// 签名使用场景。
	//
	// example:
	//
	// 已注册商标名称。
	SignUsage *string `json:"SignUsage,omitempty" xml:"SignUsage,omitempty"`
	// 签名为自用或他用。
	//
	// - false：自用（默认值）。
	//
	// - true：他用。
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
	// 商标实体id。
	//
	// example:
	//
	// 1000009081***
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
}

func (s GetSmsSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBody) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
}

func (s *GetSmsSignResponseBody) GetApplyScene() *string {
	return s.ApplyScene
}

func (s *GetSmsSignResponseBody) GetAuditInfo() *GetSmsSignResponseBodyAuditInfo {
	return s.AuditInfo
}

func (s *GetSmsSignResponseBody) GetAuthorizationLetterAuditPass() *bool {
	return s.AuthorizationLetterAuditPass
}

func (s *GetSmsSignResponseBody) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *GetSmsSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsSignResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetSmsSignResponseBody) GetFileUrlList() []*string {
	return s.FileUrlList
}

func (s *GetSmsSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsSignResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *GetSmsSignResponseBody) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *GetSmsSignResponseBody) GetRegisterResult() *int32 {
	return s.RegisterResult
}

func (s *GetSmsSignResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *GetSmsSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsSignResponseBody) GetSignCode() *string {
	return s.SignCode
}

func (s *GetSmsSignResponseBody) GetSignIspRegisterDetailList() []*GetSmsSignResponseBodySignIspRegisterDetailList {
	return s.SignIspRegisterDetailList
}

func (s *GetSmsSignResponseBody) GetSignName() *string {
	return s.SignName
}

func (s *GetSmsSignResponseBody) GetSignStatus() *int64 {
	return s.SignStatus
}

func (s *GetSmsSignResponseBody) GetSignTag() *string {
	return s.SignTag
}

func (s *GetSmsSignResponseBody) GetSignUsage() *string {
	return s.SignUsage
}

func (s *GetSmsSignResponseBody) GetThirdParty() *bool {
	return s.ThirdParty
}

func (s *GetSmsSignResponseBody) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *GetSmsSignResponseBody) SetAppIcpRecordId(v int64) *GetSmsSignResponseBody {
	s.AppIcpRecordId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetApplyScene(v string) *GetSmsSignResponseBody {
	s.ApplyScene = &v
	return s
}

func (s *GetSmsSignResponseBody) SetAuditInfo(v *GetSmsSignResponseBodyAuditInfo) *GetSmsSignResponseBody {
	s.AuditInfo = v
	return s
}

func (s *GetSmsSignResponseBody) SetAuthorizationLetterAuditPass(v bool) *GetSmsSignResponseBody {
	s.AuthorizationLetterAuditPass = &v
	return s
}

func (s *GetSmsSignResponseBody) SetAuthorizationLetterId(v int64) *GetSmsSignResponseBody {
	s.AuthorizationLetterId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetCode(v string) *GetSmsSignResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsSignResponseBody) SetCreateDate(v string) *GetSmsSignResponseBody {
	s.CreateDate = &v
	return s
}

func (s *GetSmsSignResponseBody) SetFileUrlList(v []*string) *GetSmsSignResponseBody {
	s.FileUrlList = v
	return s
}

func (s *GetSmsSignResponseBody) SetMessage(v string) *GetSmsSignResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsSignResponseBody) SetOrderId(v string) *GetSmsSignResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetQualificationId(v int64) *GetSmsSignResponseBody {
	s.QualificationId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRegisterResult(v int32) *GetSmsSignResponseBody {
	s.RegisterResult = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRemark(v string) *GetSmsSignResponseBody {
	s.Remark = &v
	return s
}

func (s *GetSmsSignResponseBody) SetRequestId(v string) *GetSmsSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignCode(v string) *GetSmsSignResponseBody {
	s.SignCode = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignIspRegisterDetailList(v []*GetSmsSignResponseBodySignIspRegisterDetailList) *GetSmsSignResponseBody {
	s.SignIspRegisterDetailList = v
	return s
}

func (s *GetSmsSignResponseBody) SetSignName(v string) *GetSmsSignResponseBody {
	s.SignName = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignStatus(v int64) *GetSmsSignResponseBody {
	s.SignStatus = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignTag(v string) *GetSmsSignResponseBody {
	s.SignTag = &v
	return s
}

func (s *GetSmsSignResponseBody) SetSignUsage(v string) *GetSmsSignResponseBody {
	s.SignUsage = &v
	return s
}

func (s *GetSmsSignResponseBody) SetThirdParty(v bool) *GetSmsSignResponseBody {
	s.ThirdParty = &v
	return s
}

func (s *GetSmsSignResponseBody) SetTrademarkId(v int64) *GetSmsSignResponseBody {
	s.TrademarkId = &v
	return s
}

func (s *GetSmsSignResponseBody) Validate() error {
	if s.AuditInfo != nil {
		if err := s.AuditInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SignIspRegisterDetailList != nil {
		for _, item := range s.SignIspRegisterDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmsSignResponseBodyAuditInfo struct {
	// 审核时间。
	//
	// example:
	//
	// 2024-06-03 12:02:34
	AuditDate *string `json:"AuditDate,omitempty" xml:"AuditDate,omitempty"`
	// 审批未通过的原因。
	//
	// example:
	//
	// 通过资质信息，不能判断是否可以申请此签名。
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
}

func (s GetSmsSignResponseBodyAuditInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBodyAuditInfo) GetAuditDate() *string {
	return s.AuditDate
}

func (s *GetSmsSignResponseBodyAuditInfo) GetRejectInfo() *string {
	return s.RejectInfo
}

func (s *GetSmsSignResponseBodyAuditInfo) SetAuditDate(v string) *GetSmsSignResponseBodyAuditInfo {
	s.AuditDate = &v
	return s
}

func (s *GetSmsSignResponseBodyAuditInfo) SetRejectInfo(v string) *GetSmsSignResponseBodyAuditInfo {
	s.RejectInfo = &v
	return s
}

func (s *GetSmsSignResponseBodyAuditInfo) Validate() error {
	return dara.Validate(s)
}

type GetSmsSignResponseBodySignIspRegisterDetailList struct {
	// 运营商类型。取值：
	//
	// - **mobile**：中国移动；
	//
	// - **unicom**：中国联通；
	//
	// - **telecom**：中国电信。
	//
	// example:
	//
	// telecom
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
	// 运营商反馈时间，格式为yyyy-MM-dd HH:mm:ss。
	//
	// example:
	//
	// 2025-06-13 15:55:26
	OperatorCompleteTime *string `json:"OperatorCompleteTime,omitempty" xml:"OperatorCompleteTime,omitempty"`
	// 报备状态。取值：
	//
	// - **0**：报备失败，原因可能为资质信息与工信注册信息不一致或运营商侧无法支持等。建议您登录[短信服务控制台](https://dysms.console.aliyun.com/domestic/text/sign)查看具体失败原因，并依据提示进行操作；
	//
	// - **1**：已报备待验证，当前至少有一个子端口号运营商已返回报备通过，建议您少量多次向不同运营商手机号发送验证码、通知短信进行验证；
	//
	// - **2**：报备失效，签名超过 6 个月无发送记录时，报备结果失效。如您需要重新启用该签名，请在[短信服务控制台](https://dysms.console.aliyun.com/domestic/text/sign)重新发起报备；
	//
	// - **3**：报备成功，当前至少有一个子端口号运营商已返回报备通过，经验证短信发送成功率符合预期，建议您持续关注发送成功率；
	//
	// - **-1**：报备中，当前尚未收到运营商反馈的报备结果，建议您等待签名报备状态变更为“已报备待验证”后再批量发送，当前可少量多次尝试使用该签名发送，观察短信发送效果；
	//
	// - **-2**：未报备，原因可能为当前签名未关联实名资质或关联资质信息不全，建议您修改当前资质或编辑签名绑定其他资质以重新发起报备。
	//
	// 建议您单击查看[更多签名实名制报备内容及建议操作](https://help.aliyun.com/document_detail/2873145.html)。
	//
	// example:
	//
	// 0
	RegisterStatus *int32 `json:"RegisterStatus,omitempty" xml:"RegisterStatus,omitempty"`
	// 报备状态原因列表。
	RegisterStatusReasons []*GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons `json:"RegisterStatusReasons,omitempty" xml:"RegisterStatusReasons,omitempty" type:"Repeated"`
}

func (s GetSmsSignResponseBodySignIspRegisterDetailList) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBodySignIspRegisterDetailList) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) GetOperatorCompleteTime() *string {
	return s.OperatorCompleteTime
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) GetRegisterStatus() *int32 {
	return s.RegisterStatus
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) GetRegisterStatusReasons() []*GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons {
	return s.RegisterStatusReasons
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) SetOperatorCode(v string) *GetSmsSignResponseBodySignIspRegisterDetailList {
	s.OperatorCode = &v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) SetOperatorCompleteTime(v string) *GetSmsSignResponseBodySignIspRegisterDetailList {
	s.OperatorCompleteTime = &v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) SetRegisterStatus(v int32) *GetSmsSignResponseBodySignIspRegisterDetailList {
	s.RegisterStatus = &v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) SetRegisterStatusReasons(v []*GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) *GetSmsSignResponseBodySignIspRegisterDetailList {
	s.RegisterStatusReasons = v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailList) Validate() error {
	if s.RegisterStatusReasons != nil {
		for _, item := range s.RegisterStatusReasons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons struct {
	// 报备状态原因码。取值：
	//
	// - **UNBINDING_QUA**：签名未关联资质；
	//
	// - **BINDING_INCOMPLETE_QUA**：关联资质信息不全；
	//
	// - **NON_REGISTER**：未发起报备；
	//
	// - **REGISTERING**：签名报备中；
	//
	// - **DETECTING**：未发起探测或探测中；
	//
	// - **DETECT_SUCCESS**：报备成功；
	//
	// - **QUALIFICATION_ERROR**：资质原因；
	//
	// - **SIGNATURE_ERROR**：签名原因；
	//
	// - **SIGNATURE_QUALIFICATION_ERROR**：签名与资质关系不符；
	//
	// - **ONE_CODE_MULTIPLE_SIGN**：扩展码原因；
	//
	// - **OTHERS_ERROR**：其他原因；
	//
	// - **REGISTER_TIMEOUT**：报备超时；
	//
	// - **NO_SEND_RECORD**：签名超过6个月无发送记录；
	//
	// - **EXT_CODE_RECYCLE**：扩展码收回。
	//
	// - **SUBPORT_RECYCLE**：子端口被运营商治理。
	//
	// example:
	//
	// REGISTER_TIMEOUT
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 原因说明列表。可能返回0个或者多个原因说明，返回原因码不一定会返回原因说明。
	ReasonDescList []*string `json:"ReasonDescList,omitempty" xml:"ReasonDescList,omitempty" type:"Repeated"`
}

func (s GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) GetReasonDescList() []*string {
	return s.ReasonDescList
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) SetReasonCode(v string) *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons {
	s.ReasonCode = &v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) SetReasonDescList(v []*string) *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons {
	s.ReasonDescList = v
	return s
}

func (s *GetSmsSignResponseBodySignIspRegisterDetailListRegisterStatusReasons) Validate() error {
	return dara.Validate(s)
}
