// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetLogListRequest
	GetBeginDate() *string
	SetClientIpList(v []*string) *GetLogListRequest
	GetClientIpList() []*string
	SetClientProgramList(v []*string) *GetLogListRequest
	GetClientProgramList() []*string
	SetDbHostList(v []*string) *GetLogListRequest
	GetDbHostList() []*string
	SetDbId(v int32) *GetLogListRequest
	GetDbId() *int32
	SetDbUserList(v []*string) *GetLogListRequest
	GetDbUserList() []*string
	SetEndDate(v string) *GetLogListRequest
	GetEndDate() *string
	SetInstanceId(v string) *GetLogListRequest
	GetInstanceId() *string
	SetIsSuccess(v string) *GetLogListRequest
	GetIsSuccess() *string
	SetLang(v string) *GetLogListRequest
	GetLang() *string
	SetMaxAffectRows(v int32) *GetLogListRequest
	GetMaxAffectRows() *int32
	SetMaxExecCostUS(v int32) *GetLogListRequest
	GetMaxExecCostUS() *int32
	SetMinAffectRows(v int32) *GetLogListRequest
	GetMinAffectRows() *int32
	SetMinExecCostUS(v int32) *GetLogListRequest
	GetMinExecCostUS() *int32
	SetPageNumber(v int32) *GetLogListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetLogListRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetLogListRequest
	GetRegionId() *string
	SetRiskLevelList(v []*string) *GetLogListRequest
	GetRiskLevelList() []*string
	SetRuleName(v string) *GetLogListRequest
	GetRuleName() *string
	SetRuleTypeList(v []*string) *GetLogListRequest
	GetRuleTypeList() []*string
	SetSessionId(v string) *GetLogListRequest
	GetSessionId() *string
	SetSqlId(v string) *GetLogListRequest
	GetSqlId() *string
	SetSqlKey(v string) *GetLogListRequest
	GetSqlKey() *string
	SetSqlTypeList(v []*string) *GetLogListRequest
	GetSqlTypeList() []*string
	SetTemplateId(v string) *GetLogListRequest
	GetTemplateId() *string
}

type GetLogListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 00:00:00
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// 111.164.XX.XX
	ClientIpList []*string `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	// example:
	//
	// navicat
	ClientProgramList []*string `json:"ClientProgramList,omitempty" xml:"ClientProgramList,omitempty" type:"Repeated"`
	// example:
	//
	// 192.168.XX.XX
	DbHostList []*string `json:"DbHostList,omitempty" xml:"DbHostList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// root
	DbUserList []*string `json:"DbUserList,omitempty" xml:"DbUserList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-06 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dbaudit-cn-78v1gc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsSuccess *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1000
	MaxAffectRows *int32 `json:"MaxAffectRows,omitempty" xml:"MaxAffectRows,omitempty"`
	// example:
	//
	// 20000
	MaxExecCostUS *int32 `json:"MaxExecCostUS,omitempty" xml:"MaxExecCostUS,omitempty"`
	// example:
	//
	// 10
	MinAffectRows *int32 `json:"MinAffectRows,omitempty" xml:"MinAffectRows,omitempty"`
	// example:
	//
	// 100
	MinExecCostUS *int32 `json:"MinExecCostUS,omitempty" xml:"MinExecCostUS,omitempty"`
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1
	RiskLevelList []*string `json:"RiskLevelList,omitempty" xml:"RiskLevelList,omitempty" type:"Repeated"`
	RuleName      *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 5
	RuleTypeList []*string `json:"RuleTypeList,omitempty" xml:"RuleTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 3011610850021000000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1907181552270011186
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// example:
	//
	// select
	SqlKey *string `json:"SqlKey,omitempty" xml:"SqlKey,omitempty"`
	// example:
	//
	// 8
	SqlTypeList []*string `json:"SqlTypeList,omitempty" xml:"SqlTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 1000****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetLogListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogListRequest) GoString() string {
	return s.String()
}

func (s *GetLogListRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetLogListRequest) GetClientIpList() []*string {
	return s.ClientIpList
}

func (s *GetLogListRequest) GetClientProgramList() []*string {
	return s.ClientProgramList
}

func (s *GetLogListRequest) GetDbHostList() []*string {
	return s.DbHostList
}

func (s *GetLogListRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetLogListRequest) GetDbUserList() []*string {
	return s.DbUserList
}

func (s *GetLogListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetLogListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLogListRequest) GetIsSuccess() *string {
	return s.IsSuccess
}

func (s *GetLogListRequest) GetLang() *string {
	return s.Lang
}

func (s *GetLogListRequest) GetMaxAffectRows() *int32 {
	return s.MaxAffectRows
}

func (s *GetLogListRequest) GetMaxExecCostUS() *int32 {
	return s.MaxExecCostUS
}

func (s *GetLogListRequest) GetMinAffectRows() *int32 {
	return s.MinAffectRows
}

func (s *GetLogListRequest) GetMinExecCostUS() *int32 {
	return s.MinExecCostUS
}

func (s *GetLogListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetLogListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLogListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLogListRequest) GetRiskLevelList() []*string {
	return s.RiskLevelList
}

func (s *GetLogListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *GetLogListRequest) GetRuleTypeList() []*string {
	return s.RuleTypeList
}

func (s *GetLogListRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetLogListRequest) GetSqlId() *string {
	return s.SqlId
}

func (s *GetLogListRequest) GetSqlKey() *string {
	return s.SqlKey
}

func (s *GetLogListRequest) GetSqlTypeList() []*string {
	return s.SqlTypeList
}

func (s *GetLogListRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLogListRequest) SetBeginDate(v string) *GetLogListRequest {
	s.BeginDate = &v
	return s
}

func (s *GetLogListRequest) SetClientIpList(v []*string) *GetLogListRequest {
	s.ClientIpList = v
	return s
}

func (s *GetLogListRequest) SetClientProgramList(v []*string) *GetLogListRequest {
	s.ClientProgramList = v
	return s
}

func (s *GetLogListRequest) SetDbHostList(v []*string) *GetLogListRequest {
	s.DbHostList = v
	return s
}

func (s *GetLogListRequest) SetDbId(v int32) *GetLogListRequest {
	s.DbId = &v
	return s
}

func (s *GetLogListRequest) SetDbUserList(v []*string) *GetLogListRequest {
	s.DbUserList = v
	return s
}

func (s *GetLogListRequest) SetEndDate(v string) *GetLogListRequest {
	s.EndDate = &v
	return s
}

func (s *GetLogListRequest) SetInstanceId(v string) *GetLogListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLogListRequest) SetIsSuccess(v string) *GetLogListRequest {
	s.IsSuccess = &v
	return s
}

func (s *GetLogListRequest) SetLang(v string) *GetLogListRequest {
	s.Lang = &v
	return s
}

func (s *GetLogListRequest) SetMaxAffectRows(v int32) *GetLogListRequest {
	s.MaxAffectRows = &v
	return s
}

func (s *GetLogListRequest) SetMaxExecCostUS(v int32) *GetLogListRequest {
	s.MaxExecCostUS = &v
	return s
}

func (s *GetLogListRequest) SetMinAffectRows(v int32) *GetLogListRequest {
	s.MinAffectRows = &v
	return s
}

func (s *GetLogListRequest) SetMinExecCostUS(v int32) *GetLogListRequest {
	s.MinExecCostUS = &v
	return s
}

func (s *GetLogListRequest) SetPageNumber(v int32) *GetLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLogListRequest) SetPageSize(v int32) *GetLogListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLogListRequest) SetRegionId(v string) *GetLogListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLogListRequest) SetRiskLevelList(v []*string) *GetLogListRequest {
	s.RiskLevelList = v
	return s
}

func (s *GetLogListRequest) SetRuleName(v string) *GetLogListRequest {
	s.RuleName = &v
	return s
}

func (s *GetLogListRequest) SetRuleTypeList(v []*string) *GetLogListRequest {
	s.RuleTypeList = v
	return s
}

func (s *GetLogListRequest) SetSessionId(v string) *GetLogListRequest {
	s.SessionId = &v
	return s
}

func (s *GetLogListRequest) SetSqlId(v string) *GetLogListRequest {
	s.SqlId = &v
	return s
}

func (s *GetLogListRequest) SetSqlKey(v string) *GetLogListRequest {
	s.SqlKey = &v
	return s
}

func (s *GetLogListRequest) SetSqlTypeList(v []*string) *GetLogListRequest {
	s.SqlTypeList = v
	return s
}

func (s *GetLogListRequest) SetTemplateId(v string) *GetLogListRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLogListRequest) Validate() error {
	return dara.Validate(s)
}
