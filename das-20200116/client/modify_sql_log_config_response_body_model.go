// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySqlLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySqlLogConfigResponseBody
	GetCode() *string
	SetData(v *ModifySqlLogConfigResponseBodyData) *ModifySqlLogConfigResponseBody
	GetData() *ModifySqlLogConfigResponseBodyData
	SetMessage(v string) *ModifySqlLogConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySqlLogConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySqlLogConfigResponseBody
	GetSuccess() *string
}

type ModifySqlLogConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 403
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ModifySqlLogConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 03C88D8E-1541-518E-8BFF-BEC6589B6334
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// False
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySqlLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySqlLogConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySqlLogConfigResponseBody) GetData() *ModifySqlLogConfigResponseBodyData {
	return s.Data
}

func (s *ModifySqlLogConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySqlLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySqlLogConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySqlLogConfigResponseBody) SetCode(v string) *ModifySqlLogConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySqlLogConfigResponseBody) SetData(v *ModifySqlLogConfigResponseBodyData) *ModifySqlLogConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifySqlLogConfigResponseBody) SetMessage(v string) *ModifySqlLogConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySqlLogConfigResponseBody) SetRequestId(v string) *ModifySqlLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySqlLogConfigResponseBody) SetSuccess(v string) *ModifySqlLogConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySqlLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifySqlLogConfigResponseBodyData struct {
	// Indicates whether the cold data storage is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ColdEnable *bool `json:"ColdEnable,omitempty" xml:"ColdEnable,omitempty"`
	// The number of days for which the SQL Explorer and Audit data is stored in cold storage. The value is calculated by using the following formula: Value of ColdRetention = Value of Retention - Value of HotRetention.``
	//
	// example:
	//
	// 23
	ColdRetention *int32 `json:"ColdRetention,omitempty" xml:"ColdRetention,omitempty"`
	// The time when the cold data storage was enabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1683712800000
	ColdStartTime *int64 `json:"ColdStartTime,omitempty" xml:"ColdStartTime,omitempty"`
	// The collector version. Valid values:
	//
	// 	- **MYSQL_V0**
	//
	// 	- **MYSQL_V1**
	//
	// 	- **MYSQL_V2**
	//
	// 	- **MYSQL_V3**
	//
	// 	- **PG_V1**
	//
	// 	- **rdspg_v1**
	//
	// 	- **polarpg_v1**
	//
	// example:
	//
	// MYSQL_V3
	CollectorVersion *string `json:"CollectorVersion,omitempty" xml:"CollectorVersion,omitempty"`
	// Indicates whether the hot data storage is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HotEnable *bool `json:"HotEnable,omitempty" xml:"HotEnable,omitempty"`
	// The number of days for which the SQL Explorer and Audit data is stored in hot storage.
	//
	// example:
	//
	// 7
	HotRetention *int32 `json:"HotRetention,omitempty" xml:"HotRetention,omitempty"`
	// The time when the hot data storage was enabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1683712800000
	HotStartTime *int64 `json:"HotStartTime,omitempty" xml:"HotStartTime,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	LogFilter *string `json:"LogFilter,omitempty" xml:"LogFilter,omitempty"`
	// Indicates whether the SQL Explorer feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	RequestEnable *bool `json:"RequestEnable,omitempty" xml:"RequestEnable,omitempty"`
	// The time when the SQL Explorer feature was enabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1683712800000
	RequestStartTime *int64 `json:"RequestStartTime,omitempty" xml:"RequestStartTime,omitempty"`
	// The time when DAS Enterprise Edition V1 expired. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1715335200000
	RequestStopTime *int64 `json:"RequestStopTime,omitempty" xml:"RequestStopTime,omitempty"`
	// The total storage duration of the SQL Explorer and Audit data. Unit: day.
	//
	// example:
	//
	// 30
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Indicates whether DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SqlLogEnable *bool `json:"SqlLogEnable,omitempty" xml:"SqlLogEnable,omitempty"`
	// The state of data migration. Valid values:
	//
	// 	- **FINISH**: The historical data is migrated.
	//
	// 	- **RUNNING**: The historical data is being migrated.
	//
	// 	- **FAILURE**: The historical data fails to be migrated.
	//
	// example:
	//
	// FINISH
	SqlLogState *string `json:"SqlLogState,omitempty" xml:"SqlLogState,omitempty"`
	// The time when DAS Enterprise Edition was enabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1683712800000
	SqlLogVisibleTime *int64 `json:"SqlLogVisibleTime,omitempty" xml:"SqlLogVisibleTime,omitempty"`
	// The latest version of DAS Enterprise Edition that supports the database instance. Valid values:
	//
	// 	- **SQL_LOG_V0**: DAS Enterprise Edition V0.
	//
	// 	- **SQL_LOG_V1**: DAS Enterprise Edition V1.
	//
	// 	- **SQL_LOG_V2**: DAS Enterprise Edition V2.
	//
	// 	- **SQL_LOG_V3**: DAS Enterprise Edition V3.
	//
	// 	- **SQL_LOG_NOT_ENABLE**: DAS Enterprise Edition is not enabled.
	//
	// 	- **SQL_LOG_NOT_SUPPORT**: DAS Enterprise Edition is not supported.
	//
	// example:
	//
	// SQL_LOG_V3
	SupportVersion *string `json:"SupportVersion,omitempty" xml:"SupportVersion,omitempty"`
	// The version of DAS Enterprise Edition that is enabled for the database instance. Valid values:
	//
	// 	- **SQL_LOG_V0**: DAS Enterprise Edition V0.
	//
	// 	- **SQL_LOG_V1**: DAS Enterprise Edition V1.
	//
	// 	- **SQL_LOG_V2**: DAS Enterprise Edition V2.
	//
	// 	- **SQL_LOG_V3**: DAS Enterprise Edition V3.
	//
	// 	- **SQL_LOG_NOT_ENABLE**: DAS Enterprise Edition is not enabled.
	//
	// 	- **SQL_LOG_NOT_SUPPORT**: DAS Enterprise Edition is not supported.
	//
	// example:
	//
	// SQL_LOG_V3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifySqlLogConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySqlLogConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySqlLogConfigResponseBodyData) GetColdEnable() *bool {
	return s.ColdEnable
}

func (s *ModifySqlLogConfigResponseBodyData) GetColdRetention() *int32 {
	return s.ColdRetention
}

func (s *ModifySqlLogConfigResponseBodyData) GetColdStartTime() *int64 {
	return s.ColdStartTime
}

func (s *ModifySqlLogConfigResponseBodyData) GetCollectorVersion() *string {
	return s.CollectorVersion
}

func (s *ModifySqlLogConfigResponseBodyData) GetHotEnable() *bool {
	return s.HotEnable
}

func (s *ModifySqlLogConfigResponseBodyData) GetHotRetention() *int32 {
	return s.HotRetention
}

func (s *ModifySqlLogConfigResponseBodyData) GetHotStartTime() *int64 {
	return s.HotStartTime
}

func (s *ModifySqlLogConfigResponseBodyData) GetLogFilter() *string {
	return s.LogFilter
}

func (s *ModifySqlLogConfigResponseBodyData) GetRequestEnable() *bool {
	return s.RequestEnable
}

func (s *ModifySqlLogConfigResponseBodyData) GetRequestStartTime() *int64 {
	return s.RequestStartTime
}

func (s *ModifySqlLogConfigResponseBodyData) GetRequestStopTime() *int64 {
	return s.RequestStopTime
}

func (s *ModifySqlLogConfigResponseBodyData) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifySqlLogConfigResponseBodyData) GetSqlLogEnable() *bool {
	return s.SqlLogEnable
}

func (s *ModifySqlLogConfigResponseBodyData) GetSqlLogState() *string {
	return s.SqlLogState
}

func (s *ModifySqlLogConfigResponseBodyData) GetSqlLogVisibleTime() *int64 {
	return s.SqlLogVisibleTime
}

func (s *ModifySqlLogConfigResponseBodyData) GetSupportVersion() *string {
	return s.SupportVersion
}

func (s *ModifySqlLogConfigResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ModifySqlLogConfigResponseBodyData) SetColdEnable(v bool) *ModifySqlLogConfigResponseBodyData {
	s.ColdEnable = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetColdRetention(v int32) *ModifySqlLogConfigResponseBodyData {
	s.ColdRetention = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetColdStartTime(v int64) *ModifySqlLogConfigResponseBodyData {
	s.ColdStartTime = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetCollectorVersion(v string) *ModifySqlLogConfigResponseBodyData {
	s.CollectorVersion = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetHotEnable(v bool) *ModifySqlLogConfigResponseBodyData {
	s.HotEnable = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetHotRetention(v int32) *ModifySqlLogConfigResponseBodyData {
	s.HotRetention = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetHotStartTime(v int64) *ModifySqlLogConfigResponseBodyData {
	s.HotStartTime = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetLogFilter(v string) *ModifySqlLogConfigResponseBodyData {
	s.LogFilter = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetRequestEnable(v bool) *ModifySqlLogConfigResponseBodyData {
	s.RequestEnable = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetRequestStartTime(v int64) *ModifySqlLogConfigResponseBodyData {
	s.RequestStartTime = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetRequestStopTime(v int64) *ModifySqlLogConfigResponseBodyData {
	s.RequestStopTime = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetRetention(v int32) *ModifySqlLogConfigResponseBodyData {
	s.Retention = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetSqlLogEnable(v bool) *ModifySqlLogConfigResponseBodyData {
	s.SqlLogEnable = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetSqlLogState(v string) *ModifySqlLogConfigResponseBodyData {
	s.SqlLogState = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetSqlLogVisibleTime(v int64) *ModifySqlLogConfigResponseBodyData {
	s.SqlLogVisibleTime = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetSupportVersion(v string) *ModifySqlLogConfigResponseBodyData {
	s.SupportVersion = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) SetVersion(v string) *ModifySqlLogConfigResponseBodyData {
	s.Version = &v
	return s
}

func (s *ModifySqlLogConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
