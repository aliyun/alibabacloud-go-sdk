// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeDataStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeDataStatsResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeDataStatsResponseBodyData) *GetQueryOptimizeDataStatsResponseBody
	GetData() *GetQueryOptimizeDataStatsResponseBodyData
	SetMessage(v string) *GetQueryOptimizeDataStatsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeDataStatsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeDataStatsResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeDataStatsResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetQueryOptimizeDataStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeDataStatsResponseBody) GetData() *GetQueryOptimizeDataStatsResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeDataStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeDataStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeDataStatsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetCode(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetData(v *GetQueryOptimizeDataStatsResponseBodyData) *GetQueryOptimizeDataStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetMessage(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetRequestId(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) SetSuccess(v string) *GetQueryOptimizeDataStatsResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueryOptimizeDataStatsResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The information about the SQL templates.
	List []*GetQueryOptimizeDataStatsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) GetList() []*GetQueryOptimizeDataStatsResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetExtra(v string) *GetQueryOptimizeDataStatsResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetList(v []*GetQueryOptimizeDataStatsResponseBodyDataList) *GetQueryOptimizeDataStatsResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeDataStatsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeDataStatsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) SetTotal(v int64) *GetQueryOptimizeDataStatsResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueryOptimizeDataStatsResponseBodyDataList struct {
	// The average lock wait time. Unit: seconds.
	//
	// example:
	//
	// 0.1
	AvgLockTime *float64 `json:"AvgLockTime,omitempty" xml:"AvgLockTime,omitempty"`
	// The average query execution time. Unit: seconds.
	//
	// example:
	//
	// 1.1
	AvgQueryTime *float64 `json:"AvgQueryTime,omitempty" xml:"AvgQueryTime,omitempty"`
	// The average number of rows affected by the SQL statement.
	//
	// > A value of -1 indicates that this parameter is not collected.
	//
	// example:
	//
	// 100.1
	AvgRowsAffected *float64 `json:"AvgRowsAffected,omitempty" xml:"AvgRowsAffected,omitempty"`
	// The average number of scanned rows.
	//
	// example:
	//
	// 100.1
	AvgRowsExamined *float64 `json:"AvgRowsExamined,omitempty" xml:"AvgRowsExamined,omitempty"`
	// The average number of returned rows.
	//
	// example:
	//
	// 100.1
	AvgRowsSent *float64 `json:"AvgRowsSent,omitempty" xml:"AvgRowsSent,omitempty"`
	// The number of times that the SQL template is executed.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the database to which the SQL template belongs.
	//
	// example:
	//
	// testdb01
	Dbname *string `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The longest lock wait time. Unit: seconds.
	//
	// example:
	//
	// 0.1
	MaxLockTime *float64 `json:"MaxLockTime,omitempty" xml:"MaxLockTime,omitempty"`
	// The longest query execution time. Unit: seconds.
	//
	// example:
	//
	// 1.1
	MaxQueryTime *float64 `json:"MaxQueryTime,omitempty" xml:"MaxQueryTime,omitempty"`
	// The largest number of rows affected by the SQL template.
	//
	// > A value of -1 indicates that this parameter is not collected.
	//
	// example:
	//
	// 10000
	MaxRowsAffected *int64 `json:"MaxRowsAffected,omitempty" xml:"MaxRowsAffected,omitempty"`
	// The largest number of scanned rows.
	//
	// example:
	//
	// 100000
	MaxRowsExamined *int64 `json:"MaxRowsExamined,omitempty" xml:"MaxRowsExamined,omitempty"`
	// The largest number of returned rows.
	//
	// example:
	//
	// 10000
	MaxRowsSent *int64 `json:"MaxRowsSent,omitempty" xml:"MaxRowsSent,omitempty"`
	// The SQL template.
	//
	// example:
	//
	// select 1
	Psql *string `json:"Psql,omitempty" xml:"Psql,omitempty"`
	// The information about the rules.
	RuleList []*GetQueryOptimizeDataStatsResponseBodyDataListRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	// The SQL template ID.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The sample query that took the longest time to execute.
	//
	// example:
	//
	// select 2
	SqlSample *string `json:"SqlSample,omitempty" xml:"SqlSample,omitempty"`
	// The type of the SQL statement.
	//
	// example:
	//
	// INSERT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The account of the database.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetAvgLockTime() *float64 {
	return s.AvgLockTime
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetAvgQueryTime() *float64 {
	return s.AvgQueryTime
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetAvgRowsAffected() *float64 {
	return s.AvgRowsAffected
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetAvgRowsExamined() *float64 {
	return s.AvgRowsExamined
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetAvgRowsSent() *float64 {
	return s.AvgRowsSent
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetCount() *int32 {
	return s.Count
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetDbname() *string {
	return s.Dbname
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetMaxLockTime() *float64 {
	return s.MaxLockTime
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetMaxQueryTime() *float64 {
	return s.MaxQueryTime
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetMaxRowsAffected() *int64 {
	return s.MaxRowsAffected
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetMaxRowsExamined() *int64 {
	return s.MaxRowsExamined
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetMaxRowsSent() *int64 {
	return s.MaxRowsSent
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetPsql() *string {
	return s.Psql
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetRuleList() []*GetQueryOptimizeDataStatsResponseBodyDataListRuleList {
	return s.RuleList
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetSqlSample() *string {
	return s.SqlSample
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetSqlType() *string {
	return s.SqlType
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) GetUser() *string {
	return s.User
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgLockTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgLockTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgQueryTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgQueryTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsAffected(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsAffected = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsExamined(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsExamined = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetAvgRowsSent(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.AvgRowsSent = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetCount(v int32) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetInstanceId(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxLockTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxLockTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxQueryTime(v float64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxQueryTime = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsAffected(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsAffected = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsExamined(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsExamined = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetMaxRowsSent(v int64) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.MaxRowsSent = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetPsql(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.Psql = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetRuleList(v []*GetQueryOptimizeDataStatsResponseBodyDataListRuleList) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.RuleList = v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlSample(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlSample = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetSqlType(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.SqlType = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) SetUser(v string) *GetQueryOptimizeDataStatsResponseBodyDataList {
	s.User = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataList) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueryOptimizeDataStatsResponseBodyDataListRuleList struct {
	// The rule name.
	//
	// example:
	//
	// DAS_NOT_IMPORTANT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **Predefined**
	//
	// 	- **UserDefined**
	//
	// example:
	//
	// Predefined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeDataStatsResponseBodyDataListRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeDataStatsResponseBodyDataListRuleList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) GetName() *string {
	return s.Name
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) GetType() *string {
	return s.Type
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) SetName(v string) *GetQueryOptimizeDataStatsResponseBodyDataListRuleList {
	s.Name = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) SetType(v string) *GetQueryOptimizeDataStatsResponseBodyDataListRuleList {
	s.Type = &v
	return s
}

func (s *GetQueryOptimizeDataStatsResponseBodyDataListRuleList) Validate() error {
	return dara.Validate(s)
}
