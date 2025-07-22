// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlRulesHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetCode() *string
	SetData(v *GetSqlConcurrencyControlRulesHistoryResponseBodyData) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetData() *GetSqlConcurrencyControlRulesHistoryResponseBodyData
	SetMessage(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody
	GetSuccess() *string
}

type GetSqlConcurrencyControlRulesHistoryResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetSqlConcurrencyControlRulesHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
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

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetData() *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	return s.Data
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetCode(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetData(v *GetSqlConcurrencyControlRulesHistoryResponseBodyData) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetMessage(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetRequestId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) SetSuccess(v string) *GetSqlConcurrencyControlRulesHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyData struct {
	// The list of the queried throttling rules.
	List *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) GetList() *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList {
	return s.List
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetList(v *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.List = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) SetTotal(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataList struct {
	Rules []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) GetRules() []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	return s.Rules
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) SetRules(v []*GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList {
	s.Rules = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules struct {
	// The duration within which the SQL throttling rule takes effect. Unit: seconds.
	//
	// >  The throttling rule takes effect only within this duration.
	//
	// example:
	//
	// 600
	ConcurrencyControlTime *int64 `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the throttling rule that is applied to the instance.
	//
	// example:
	//
	// 16
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The hash value of the SQL keywords. The SQL keywords are contained in the SQL statements to which the throttling rule is applied.
	//
	// example:
	//
	// b0b8aceeb43baea87b219c81767b****
	KeywordsHash *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	// The maximum number of concurrent SQL statements. Set this parameter to a positive integer.
	//
	// >  When the number of concurrent SQL statements that contain the specified keywords reaches this upper limit, the throttling rule is triggered.
	//
	// example:
	//
	// 2
	MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The keywords that are used to identify the SQL statements that need to be throttled.
	//
	// > SQL keywords are separated with tildes (~). When the number of concurrent SQL statements that contain all the specified SQL keywords reaches the specified upper limit, the throttling rule is triggered.
	//
	// example:
	//
	// call~open~api~test~4~from~POP
	SqlKeywords *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
	// The type of the SQL statements. Valid values:
	//
	// 	- **SELECT**
	//
	// 	- **UPDATE**
	//
	// 	- **DELETE**
	//
	// example:
	//
	// SELECT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the throttling rule. Valid values:
	//
	// 	- **Open**: The throttling rule is in effect.
	//
	// 	- **Closed**: The throttling rule was in effect.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// testxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetConcurrencyControlTime() *int64 {
	return s.ConcurrencyControlTime
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetKeywordsHash() *string {
	return s.KeywordsHash
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetMaxConcurrency() *int64 {
	return s.MaxConcurrency
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetSqlKeywords() *string {
	return s.SqlKeywords
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetSqlType() *string {
	return s.SqlType
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetStatus() *string {
	return s.Status
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) GetUserId() *string {
	return s.UserId
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetConcurrencyControlTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetInstanceId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.InstanceId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetItemId(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.ItemId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetKeywordsHash(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetMaxConcurrency(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlKeywords(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetSqlType(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.SqlType = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStartTime(v int64) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.StartTime = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetStatus(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.Status = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) SetUserId(v string) *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules {
	s.UserId = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponseBodyDataListRules) Validate() error {
	return dara.Validate(s)
}
