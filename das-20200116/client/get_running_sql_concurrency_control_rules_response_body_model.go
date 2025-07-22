// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningSqlConcurrencyControlRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetCode() *string
	SetData(v *GetRunningSqlConcurrencyControlRulesResponseBodyData) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetData() *GetRunningSqlConcurrencyControlRulesResponseBodyData
	SetMessage(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRunningSqlConcurrencyControlRulesResponseBody
	GetSuccess() *string
}

type GetRunningSqlConcurrencyControlRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetRunningSqlConcurrencyControlRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetData() *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	return s.Data
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetCode(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetData(v *GetRunningSqlConcurrencyControlRulesResponseBodyData) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetMessage(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetRequestId(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) SetSuccess(v string) *GetRunningSqlConcurrencyControlRulesResponseBody {
	s.Success = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRunningSqlConcurrencyControlRulesResponseBodyData struct {
	// The returned data.
	List *GetRunningSqlConcurrencyControlRulesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) GetList() *GetRunningSqlConcurrencyControlRulesResponseBodyDataList {
	return s.List
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetList(v *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.List = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) SetTotal(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataList struct {
	RunningRules []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules `json:"runningRules,omitempty" xml:"runningRules,omitempty" type:"Repeated"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) GetRunningRules() []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	return s.RunningRules
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) SetRunningRules(v []*GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) *GetRunningSqlConcurrencyControlRulesResponseBodyDataList {
	s.RunningRules = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules struct {
	// The duration within which the SQL throttling rule takes effect. Unit: seconds.
	//
	// > The throttling rule takes effect only within this duration.
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
	// The hash value of the SQL keywords. The hash value is calculated based on the SQL keywords that are contained in the SQL statements to which the throttling rule is applied.
	//
	// example:
	//
	// b0b8aceeb43baea87b219c81767b****
	KeywordsHash *string `json:"KeywordsHash,omitempty" xml:"KeywordsHash,omitempty"`
	// The maximum number of concurrent SQL statements. The value is a positive integer.
	//
	// > If the number of concurrent SQL statements that contain the specified keywords reaches this upper limit, the throttling rule is triggered.
	//
	// example:
	//
	// 2
	MaxConcurrency *string `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The keywords contained in the SQL statements to which the throttling rule was applied.
	//
	// > SQL keywords are separated by tildes (~). If the number of concurrent SQL statements that contain all the specified SQL keywords reaches the specified upper limit, the throttling rule is triggered.
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
	// The time when the throttling rule started to take effect. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1608888296000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the throttling rule. The value of **Open*	- indicates that the throttling rule is in effect.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// testxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetConcurrencyControlTime() *int64 {
	return s.ConcurrencyControlTime
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetKeywordsHash() *string {
	return s.KeywordsHash
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetMaxConcurrency() *string {
	return s.MaxConcurrency
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetSqlKeywords() *string {
	return s.SqlKeywords
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetSqlType() *string {
	return s.SqlType
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetStatus() *string {
	return s.Status
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) GetUserId() *string {
	return s.UserId
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetConcurrencyControlTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ConcurrencyControlTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetInstanceId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.InstanceId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetItemId(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.ItemId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetKeywordsHash(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.KeywordsHash = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetMaxConcurrency(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.MaxConcurrency = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlKeywords(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlKeywords = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetSqlType(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.SqlType = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStartTime(v int64) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.StartTime = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetStatus(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.Status = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) SetUserId(v string) *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules {
	s.UserId = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponseBodyDataListRunningRules) Validate() error {
	return dara.Validate(s)
}
