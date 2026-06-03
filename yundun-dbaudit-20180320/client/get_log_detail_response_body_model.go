// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectRows(v int32) *GetLogDetailResponseBody
	GetAffectRows() *int32
	SetAppClientIp(v string) *GetLogDetailResponseBody
	GetAppClientIp() *string
	SetAppUsername(v string) *GetLogDetailResponseBody
	GetAppUsername() *string
	SetCaptureTime(v string) *GetLogDetailResponseBody
	GetCaptureTime() *string
	SetClientIp(v string) *GetLogDetailResponseBody
	GetClientIp() *string
	SetClientMac(v string) *GetLogDetailResponseBody
	GetClientMac() *string
	SetClientOsUser(v string) *GetLogDetailResponseBody
	GetClientOsUser() *string
	SetClientPort(v int32) *GetLogDetailResponseBody
	GetClientPort() *int32
	SetClientProgram(v string) *GetLogDetailResponseBody
	GetClientProgram() *string
	SetDbId(v int32) *GetLogDetailResponseBody
	GetDbId() *int32
	SetDbServer(v string) *GetLogDetailResponseBody
	GetDbServer() *string
	SetDbUser(v string) *GetLogDetailResponseBody
	GetDbUser() *string
	SetExecCostUS(v int32) *GetLogDetailResponseBody
	GetExecCostUS() *int32
	SetFetchCostUS(v int32) *GetLogDetailResponseBody
	GetFetchCostUS() *int32
	SetInstName(v string) *GetLogDetailResponseBody
	GetInstName() *string
	SetRequestId(v string) *GetLogDetailResponseBody
	GetRequestId() *string
	SetResponseCode(v string) *GetLogDetailResponseBody
	GetResponseCode() *string
	SetResponseText(v string) *GetLogDetailResponseBody
	GetResponseText() *string
	SetRiskLevel(v int32) *GetLogDetailResponseBody
	GetRiskLevel() *int32
	SetRuleId(v int32) *GetLogDetailResponseBody
	GetRuleId() *int32
	SetRuleKeyId(v int32) *GetLogDetailResponseBody
	GetRuleKeyId() *int32
	SetRuleName(v string) *GetLogDetailResponseBody
	GetRuleName() *string
	SetRuleType(v int32) *GetLogDetailResponseBody
	GetRuleType() *int32
	SetSchema(v string) *GetLogDetailResponseBody
	GetSchema() *string
	SetSecondarySqlType(v string) *GetLogDetailResponseBody
	GetSecondarySqlType() *string
	SetServerMac(v string) *GetLogDetailResponseBody
	GetServerMac() *string
	SetSessionId(v string) *GetLogDetailResponseBody
	GetSessionId() *string
	SetSessionLoginTime(v string) *GetLogDetailResponseBody
	GetSessionLoginTime() *string
	SetSessionLogoutTime(v string) *GetLogDetailResponseBody
	GetSessionLogoutTime() *string
	SetSqlContent(v string) *GetLogDetailResponseBody
	GetSqlContent() *string
	SetSqlId(v string) *GetLogDetailResponseBody
	GetSqlId() *string
	SetSqlResult(v string) *GetLogDetailResponseBody
	GetSqlResult() *string
	SetSqlType(v string) *GetLogDetailResponseBody
	GetSqlType() *string
	SetSqlTypeName(v string) *GetLogDetailResponseBody
	GetSqlTypeName() *string
	SetTemplateContent(v string) *GetLogDetailResponseBody
	GetTemplateContent() *string
	SetTemplateId(v string) *GetLogDetailResponseBody
	GetTemplateId() *string
	SetTemplateRules(v []*string) *GetLogDetailResponseBody
	GetTemplateRules() []*string
}

type GetLogDetailResponseBody struct {
	// example:
	//
	// 100
	AffectRows *int32 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// example:
	//
	// 8.8.XX.XX
	AppClientIp *string `json:"AppClientIp,omitempty" xml:"AppClientIp,omitempty"`
	// example:
	//
	// zhangsan
	AppUsername *string `json:"AppUsername,omitempty" xml:"AppUsername,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 00163E06****
	ClientMac *string `json:"ClientMac,omitempty" xml:"ClientMac,omitempty"`
	// example:
	//
	// administrator
	ClientOsUser *string `json:"ClientOsUser,omitempty" xml:"ClientOsUser,omitempty"`
	// example:
	//
	// 15629
	ClientPort *int32 `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// example:
	//
	// navicat
	ClientProgram *string `json:"ClientProgram,omitempty" xml:"ClientProgram,omitempty"`
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// 192.168.XX.XX:3306
	DbServer *string `json:"DbServer,omitempty" xml:"DbServer,omitempty"`
	// example:
	//
	// root
	DbUser *string `json:"DbUser,omitempty" xml:"DbUser,omitempty"`
	// example:
	//
	// 10
	ExecCostUS *int32 `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	// example:
	//
	// 10
	FetchCostUS *int32 `json:"FetchCostUS,omitempty" xml:"FetchCostUS,omitempty"`
	// example:
	//
	// orcl
	InstName *string `json:"InstName,omitempty" xml:"InstName,omitempty"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// example:
	//
	// Table \\"your_table\\" doesn\\"t exist
	ResponseText *string `json:"ResponseText,omitempty" xml:"ResponseText,omitempty"`
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// 50****
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 50****
	RuleKeyId *int32  `json:"RuleKeyId,omitempty" xml:"RuleKeyId,omitempty"`
	RuleName  *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 5
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// db_test
	Schema           *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	SecondarySqlType *string `json:"SecondarySqlType,omitempty" xml:"SecondarySqlType,omitempty"`
	// example:
	//
	// 00163E06****
	ServerMac *string `json:"ServerMac,omitempty" xml:"ServerMac,omitempty"`
	// example:
	//
	// 3011610850021000000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	SessionLoginTime *string `json:"SessionLoginTime,omitempty" xml:"SessionLoginTime,omitempty"`
	// example:
	//
	// 2019-06-06 00:00:00
	SessionLogoutTime *string `json:"SessionLogoutTime,omitempty" xml:"SessionLogoutTime,omitempty"`
	// example:
	//
	// select 	- from xxx where name=\\"zhangsan\\"
	SqlContent *string `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
	// example:
	//
	// 1907181552270011186
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// example:
	//
	// [["id","name","age"],["1","zhangsan","20"]]
	SqlResult *string `json:"SqlResult,omitempty" xml:"SqlResult,omitempty"`
	// example:
	//
	// 8
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// SELECT
	SqlTypeName *string `json:"SqlTypeName,omitempty" xml:"SqlTypeName,omitempty"`
	// example:
	//
	// SELECT 	- FROM XXX WHERE NAME=\\"#\\"
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// example:
	//
	// 1000****
	TemplateId    *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateRules []*string `json:"TemplateRules,omitempty" xml:"TemplateRules,omitempty" type:"Repeated"`
}

func (s GetLogDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogDetailResponseBody) GetAffectRows() *int32 {
	return s.AffectRows
}

func (s *GetLogDetailResponseBody) GetAppClientIp() *string {
	return s.AppClientIp
}

func (s *GetLogDetailResponseBody) GetAppUsername() *string {
	return s.AppUsername
}

func (s *GetLogDetailResponseBody) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *GetLogDetailResponseBody) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetLogDetailResponseBody) GetClientMac() *string {
	return s.ClientMac
}

func (s *GetLogDetailResponseBody) GetClientOsUser() *string {
	return s.ClientOsUser
}

func (s *GetLogDetailResponseBody) GetClientPort() *int32 {
	return s.ClientPort
}

func (s *GetLogDetailResponseBody) GetClientProgram() *string {
	return s.ClientProgram
}

func (s *GetLogDetailResponseBody) GetDbId() *int32 {
	return s.DbId
}

func (s *GetLogDetailResponseBody) GetDbServer() *string {
	return s.DbServer
}

func (s *GetLogDetailResponseBody) GetDbUser() *string {
	return s.DbUser
}

func (s *GetLogDetailResponseBody) GetExecCostUS() *int32 {
	return s.ExecCostUS
}

func (s *GetLogDetailResponseBody) GetFetchCostUS() *int32 {
	return s.FetchCostUS
}

func (s *GetLogDetailResponseBody) GetInstName() *string {
	return s.InstName
}

func (s *GetLogDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogDetailResponseBody) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *GetLogDetailResponseBody) GetResponseText() *string {
	return s.ResponseText
}

func (s *GetLogDetailResponseBody) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetLogDetailResponseBody) GetRuleId() *int32 {
	return s.RuleId
}

func (s *GetLogDetailResponseBody) GetRuleKeyId() *int32 {
	return s.RuleKeyId
}

func (s *GetLogDetailResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetLogDetailResponseBody) GetRuleType() *int32 {
	return s.RuleType
}

func (s *GetLogDetailResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *GetLogDetailResponseBody) GetSecondarySqlType() *string {
	return s.SecondarySqlType
}

func (s *GetLogDetailResponseBody) GetServerMac() *string {
	return s.ServerMac
}

func (s *GetLogDetailResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLogDetailResponseBody) GetSessionLoginTime() *string {
	return s.SessionLoginTime
}

func (s *GetLogDetailResponseBody) GetSessionLogoutTime() *string {
	return s.SessionLogoutTime
}

func (s *GetLogDetailResponseBody) GetSqlContent() *string {
	return s.SqlContent
}

func (s *GetLogDetailResponseBody) GetSqlId() *string {
	return s.SqlId
}

func (s *GetLogDetailResponseBody) GetSqlResult() *string {
	return s.SqlResult
}

func (s *GetLogDetailResponseBody) GetSqlType() *string {
	return s.SqlType
}

func (s *GetLogDetailResponseBody) GetSqlTypeName() *string {
	return s.SqlTypeName
}

func (s *GetLogDetailResponseBody) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetLogDetailResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLogDetailResponseBody) GetTemplateRules() []*string {
	return s.TemplateRules
}

func (s *GetLogDetailResponseBody) SetAffectRows(v int32) *GetLogDetailResponseBody {
	s.AffectRows = &v
	return s
}

func (s *GetLogDetailResponseBody) SetAppClientIp(v string) *GetLogDetailResponseBody {
	s.AppClientIp = &v
	return s
}

func (s *GetLogDetailResponseBody) SetAppUsername(v string) *GetLogDetailResponseBody {
	s.AppUsername = &v
	return s
}

func (s *GetLogDetailResponseBody) SetCaptureTime(v string) *GetLogDetailResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientIp(v string) *GetLogDetailResponseBody {
	s.ClientIp = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientMac(v string) *GetLogDetailResponseBody {
	s.ClientMac = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientOsUser(v string) *GetLogDetailResponseBody {
	s.ClientOsUser = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientPort(v int32) *GetLogDetailResponseBody {
	s.ClientPort = &v
	return s
}

func (s *GetLogDetailResponseBody) SetClientProgram(v string) *GetLogDetailResponseBody {
	s.ClientProgram = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbId(v int32) *GetLogDetailResponseBody {
	s.DbId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbServer(v string) *GetLogDetailResponseBody {
	s.DbServer = &v
	return s
}

func (s *GetLogDetailResponseBody) SetDbUser(v string) *GetLogDetailResponseBody {
	s.DbUser = &v
	return s
}

func (s *GetLogDetailResponseBody) SetExecCostUS(v int32) *GetLogDetailResponseBody {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogDetailResponseBody) SetFetchCostUS(v int32) *GetLogDetailResponseBody {
	s.FetchCostUS = &v
	return s
}

func (s *GetLogDetailResponseBody) SetInstName(v string) *GetLogDetailResponseBody {
	s.InstName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRequestId(v string) *GetLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetResponseCode(v string) *GetLogDetailResponseBody {
	s.ResponseCode = &v
	return s
}

func (s *GetLogDetailResponseBody) SetResponseText(v string) *GetLogDetailResponseBody {
	s.ResponseText = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRiskLevel(v int32) *GetLogDetailResponseBody {
	s.RiskLevel = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleId(v int32) *GetLogDetailResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleKeyId(v int32) *GetLogDetailResponseBody {
	s.RuleKeyId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleName(v string) *GetLogDetailResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetRuleType(v int32) *GetLogDetailResponseBody {
	s.RuleType = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSchema(v string) *GetLogDetailResponseBody {
	s.Schema = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSecondarySqlType(v string) *GetLogDetailResponseBody {
	s.SecondarySqlType = &v
	return s
}

func (s *GetLogDetailResponseBody) SetServerMac(v string) *GetLogDetailResponseBody {
	s.ServerMac = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionId(v string) *GetLogDetailResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionLoginTime(v string) *GetLogDetailResponseBody {
	s.SessionLoginTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSessionLogoutTime(v string) *GetLogDetailResponseBody {
	s.SessionLogoutTime = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlContent(v string) *GetLogDetailResponseBody {
	s.SqlContent = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlId(v string) *GetLogDetailResponseBody {
	s.SqlId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlResult(v string) *GetLogDetailResponseBody {
	s.SqlResult = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlType(v string) *GetLogDetailResponseBody {
	s.SqlType = &v
	return s
}

func (s *GetLogDetailResponseBody) SetSqlTypeName(v string) *GetLogDetailResponseBody {
	s.SqlTypeName = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateContent(v string) *GetLogDetailResponseBody {
	s.TemplateContent = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateId(v string) *GetLogDetailResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetLogDetailResponseBody) SetTemplateRules(v []*string) *GetLogDetailResponseBody {
	s.TemplateRules = v
	return s
}

func (s *GetLogDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
