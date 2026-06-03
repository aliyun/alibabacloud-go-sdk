// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetLogListResponseBody
	GetBeginDate() *string
	SetEndDate(v string) *GetLogListResponseBody
	GetEndDate() *string
	SetIncomplete(v string) *GetLogListResponseBody
	GetIncomplete() *string
	SetPageNumber(v int32) *GetLogListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetLogListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetLogListResponseBody
	GetRequestId() *string
	SetResults(v []*GetLogListResponseBodyResults) *GetLogListResponseBody
	GetResults() []*GetLogListResponseBodyResults
	SetTotalCount(v int32) *GetLogListResponseBody
	GetTotalCount() *int32
}

type GetLogListResponseBody struct {
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// true
	Incomplete *string `json:"Incomplete,omitempty" xml:"Incomplete,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*GetLogListResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetLogListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogListResponseBody) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetLogListResponseBody) GetEndDate() *string {
	return s.EndDate
}

func (s *GetLogListResponseBody) GetIncomplete() *string {
	return s.Incomplete
}

func (s *GetLogListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLogListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLogListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogListResponseBody) GetResults() []*GetLogListResponseBodyResults {
	return s.Results
}

func (s *GetLogListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetLogListResponseBody) SetBeginDate(v string) *GetLogListResponseBody {
	s.BeginDate = &v
	return s
}

func (s *GetLogListResponseBody) SetEndDate(v string) *GetLogListResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetLogListResponseBody) SetIncomplete(v string) *GetLogListResponseBody {
	s.Incomplete = &v
	return s
}

func (s *GetLogListResponseBody) SetPageNumber(v int32) *GetLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetLogListResponseBody) SetPageSize(v int32) *GetLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLogListResponseBody) SetRequestId(v string) *GetLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogListResponseBody) SetResults(v []*GetLogListResponseBodyResults) *GetLogListResponseBody {
	s.Results = v
	return s
}

func (s *GetLogListResponseBody) SetTotalCount(v int32) *GetLogListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetLogListResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLogListResponseBodyResults struct {
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
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientIpAlias *string `json:"ClientIpAlias,omitempty" xml:"ClientIpAlias,omitempty"`
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
	// 1000
	ExecCostUS *int32 `json:"ExecCostUS,omitempty" xml:"ExecCostUS,omitempty"`
	// example:
	//
	// 2000
	FetchCostUS *int32 `json:"FetchCostUS,omitempty" xml:"FetchCostUS,omitempty"`
	// example:
	//
	// orcl
	InstName *string `json:"InstName,omitempty" xml:"InstName,omitempty"`
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
	// 0
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
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
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
	// select 	- from ****
	SqlContent *string `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
	// example:
	//
	// 1907181552270011186
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// example:
	//
	// 8
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// example:
	//
	// 1000****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetLogListResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s GetLogListResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetLogListResponseBodyResults) GetAffectRows() *int32 {
	return s.AffectRows
}

func (s *GetLogListResponseBodyResults) GetAppClientIp() *string {
	return s.AppClientIp
}

func (s *GetLogListResponseBodyResults) GetAppUsername() *string {
	return s.AppUsername
}

func (s *GetLogListResponseBodyResults) GetCaptureTime() *string {
	return s.CaptureTime
}

func (s *GetLogListResponseBodyResults) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetLogListResponseBodyResults) GetClientIpAlias() *string {
	return s.ClientIpAlias
}

func (s *GetLogListResponseBodyResults) GetClientMac() *string {
	return s.ClientMac
}

func (s *GetLogListResponseBodyResults) GetClientOsUser() *string {
	return s.ClientOsUser
}

func (s *GetLogListResponseBodyResults) GetClientPort() *int32 {
	return s.ClientPort
}

func (s *GetLogListResponseBodyResults) GetClientProgram() *string {
	return s.ClientProgram
}

func (s *GetLogListResponseBodyResults) GetDbId() *int32 {
	return s.DbId
}

func (s *GetLogListResponseBodyResults) GetDbServer() *string {
	return s.DbServer
}

func (s *GetLogListResponseBodyResults) GetDbUser() *string {
	return s.DbUser
}

func (s *GetLogListResponseBodyResults) GetExecCostUS() *int32 {
	return s.ExecCostUS
}

func (s *GetLogListResponseBodyResults) GetFetchCostUS() *int32 {
	return s.FetchCostUS
}

func (s *GetLogListResponseBodyResults) GetInstName() *string {
	return s.InstName
}

func (s *GetLogListResponseBodyResults) GetResponseCode() *string {
	return s.ResponseCode
}

func (s *GetLogListResponseBodyResults) GetResponseText() *string {
	return s.ResponseText
}

func (s *GetLogListResponseBodyResults) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetLogListResponseBodyResults) GetRuleId() *int32 {
	return s.RuleId
}

func (s *GetLogListResponseBodyResults) GetRuleKeyId() *int32 {
	return s.RuleKeyId
}

func (s *GetLogListResponseBodyResults) GetRuleName() *string {
	return s.RuleName
}

func (s *GetLogListResponseBodyResults) GetRuleType() *int32 {
	return s.RuleType
}

func (s *GetLogListResponseBodyResults) GetSchema() *string {
	return s.Schema
}

func (s *GetLogListResponseBodyResults) GetServerMac() *string {
	return s.ServerMac
}

func (s *GetLogListResponseBodyResults) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLogListResponseBodyResults) GetSessionLoginTime() *string {
	return s.SessionLoginTime
}

func (s *GetLogListResponseBodyResults) GetSessionLogoutTime() *string {
	return s.SessionLogoutTime
}

func (s *GetLogListResponseBodyResults) GetSqlContent() *string {
	return s.SqlContent
}

func (s *GetLogListResponseBodyResults) GetSqlId() *string {
	return s.SqlId
}

func (s *GetLogListResponseBodyResults) GetSqlType() *string {
	return s.SqlType
}

func (s *GetLogListResponseBodyResults) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLogListResponseBodyResults) SetAffectRows(v int32) *GetLogListResponseBodyResults {
	s.AffectRows = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetAppClientIp(v string) *GetLogListResponseBodyResults {
	s.AppClientIp = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetAppUsername(v string) *GetLogListResponseBodyResults {
	s.AppUsername = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetCaptureTime(v string) *GetLogListResponseBodyResults {
	s.CaptureTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientIp(v string) *GetLogListResponseBodyResults {
	s.ClientIp = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientIpAlias(v string) *GetLogListResponseBodyResults {
	s.ClientIpAlias = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientMac(v string) *GetLogListResponseBodyResults {
	s.ClientMac = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientOsUser(v string) *GetLogListResponseBodyResults {
	s.ClientOsUser = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientPort(v int32) *GetLogListResponseBodyResults {
	s.ClientPort = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetClientProgram(v string) *GetLogListResponseBodyResults {
	s.ClientProgram = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbId(v int32) *GetLogListResponseBodyResults {
	s.DbId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbServer(v string) *GetLogListResponseBodyResults {
	s.DbServer = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetDbUser(v string) *GetLogListResponseBodyResults {
	s.DbUser = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetExecCostUS(v int32) *GetLogListResponseBodyResults {
	s.ExecCostUS = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetFetchCostUS(v int32) *GetLogListResponseBodyResults {
	s.FetchCostUS = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetInstName(v string) *GetLogListResponseBodyResults {
	s.InstName = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetResponseCode(v string) *GetLogListResponseBodyResults {
	s.ResponseCode = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetResponseText(v string) *GetLogListResponseBodyResults {
	s.ResponseText = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRiskLevel(v int32) *GetLogListResponseBodyResults {
	s.RiskLevel = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleId(v int32) *GetLogListResponseBodyResults {
	s.RuleId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleKeyId(v int32) *GetLogListResponseBodyResults {
	s.RuleKeyId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleName(v string) *GetLogListResponseBodyResults {
	s.RuleName = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetRuleType(v int32) *GetLogListResponseBodyResults {
	s.RuleType = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSchema(v string) *GetLogListResponseBodyResults {
	s.Schema = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetServerMac(v string) *GetLogListResponseBodyResults {
	s.ServerMac = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionId(v string) *GetLogListResponseBodyResults {
	s.SessionId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionLoginTime(v string) *GetLogListResponseBodyResults {
	s.SessionLoginTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSessionLogoutTime(v string) *GetLogListResponseBodyResults {
	s.SessionLogoutTime = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlContent(v string) *GetLogListResponseBodyResults {
	s.SqlContent = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlId(v string) *GetLogListResponseBodyResults {
	s.SqlId = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetSqlType(v string) *GetLogListResponseBodyResults {
	s.SqlType = &v
	return s
}

func (s *GetLogListResponseBodyResults) SetTemplateId(v string) *GetLogListResponseBodyResults {
	s.TemplateId = &v
	return s
}

func (s *GetLogListResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
