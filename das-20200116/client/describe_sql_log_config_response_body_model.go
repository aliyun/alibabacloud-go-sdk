// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlLogConfigResponseBody
	GetCode() *string
	SetData(v *DescribeSqlLogConfigResponseBodyData) *DescribeSqlLogConfigResponseBody
	GetData() *DescribeSqlLogConfigResponseBodyData
	SetMessage(v string) *DescribeSqlLogConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlLogConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlLogConfigResponseBody
	GetSuccess() *string
}

type DescribeSqlLogConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *DescribeSqlLogConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 0A74B755-98B7-59DB-8724-1321B394****
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

func (s DescribeSqlLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlLogConfigResponseBody) GetData() *DescribeSqlLogConfigResponseBodyData {
	return s.Data
}

func (s *DescribeSqlLogConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlLogConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlLogConfigResponseBody) SetCode(v string) *DescribeSqlLogConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBody) SetData(v *DescribeSqlLogConfigResponseBodyData) *DescribeSqlLogConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlLogConfigResponseBody) SetMessage(v string) *DescribeSqlLogConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBody) SetRequestId(v string) *DescribeSqlLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBody) SetSuccess(v string) *DescribeSqlLogConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogConfigResponseBodyData struct {
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
	// The number of days for which the SQL Explorer and Audit data is stored in cold storage.
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
	// true
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
	// The total storage duration of the SQL Explorer and Audit data. The value of this parameter is the sum of the values of **HotRetention*	- and **ColdRetention**. Unit: day.
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
	// A reserved parameter.
	//
	// example:
	//
	// None
	SqlLogSource *string `json:"SqlLogSource,omitempty" xml:"SqlLogSource,omitempty"`
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
	// A reserved parameter.
	//
	// example:
	//
	// None
	SupportMigration *bool `json:"SupportMigration,omitempty" xml:"SupportMigration,omitempty"`
	// The latest version of DAS Enterprise Edition that supports the database instance. Valid values:
	//
	// 	- **SQL_LOG_V0**: DAS Enterprise Edition V0.
	//
	// 	- **SQL_LOG_V1**: DAS Enterprise version V1.
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
	// 	- **SQL_LOG_V1**: DAS Enterprise version V1.
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

func (s DescribeSqlLogConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogConfigResponseBodyData) GetColdEnable() *bool {
	return s.ColdEnable
}

func (s *DescribeSqlLogConfigResponseBodyData) GetColdRetention() *int32 {
	return s.ColdRetention
}

func (s *DescribeSqlLogConfigResponseBodyData) GetColdStartTime() *int64 {
	return s.ColdStartTime
}

func (s *DescribeSqlLogConfigResponseBodyData) GetCollectorVersion() *string {
	return s.CollectorVersion
}

func (s *DescribeSqlLogConfigResponseBodyData) GetHotEnable() *bool {
	return s.HotEnable
}

func (s *DescribeSqlLogConfigResponseBodyData) GetHotRetention() *int32 {
	return s.HotRetention
}

func (s *DescribeSqlLogConfigResponseBodyData) GetHotStartTime() *int64 {
	return s.HotStartTime
}

func (s *DescribeSqlLogConfigResponseBodyData) GetLogFilter() *string {
	return s.LogFilter
}

func (s *DescribeSqlLogConfigResponseBodyData) GetRequestEnable() *bool {
	return s.RequestEnable
}

func (s *DescribeSqlLogConfigResponseBodyData) GetRequestStartTime() *int64 {
	return s.RequestStartTime
}

func (s *DescribeSqlLogConfigResponseBodyData) GetRequestStopTime() *int64 {
	return s.RequestStopTime
}

func (s *DescribeSqlLogConfigResponseBodyData) GetRetention() *int32 {
	return s.Retention
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSqlLogEnable() *bool {
	return s.SqlLogEnable
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSqlLogSource() *string {
	return s.SqlLogSource
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSqlLogState() *string {
	return s.SqlLogState
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSqlLogVisibleTime() *int64 {
	return s.SqlLogVisibleTime
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSupportMigration() *bool {
	return s.SupportMigration
}

func (s *DescribeSqlLogConfigResponseBodyData) GetSupportVersion() *string {
	return s.SupportVersion
}

func (s *DescribeSqlLogConfigResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeSqlLogConfigResponseBodyData) SetColdEnable(v bool) *DescribeSqlLogConfigResponseBodyData {
	s.ColdEnable = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetColdRetention(v int32) *DescribeSqlLogConfigResponseBodyData {
	s.ColdRetention = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetColdStartTime(v int64) *DescribeSqlLogConfigResponseBodyData {
	s.ColdStartTime = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetCollectorVersion(v string) *DescribeSqlLogConfigResponseBodyData {
	s.CollectorVersion = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetHotEnable(v bool) *DescribeSqlLogConfigResponseBodyData {
	s.HotEnable = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetHotRetention(v int32) *DescribeSqlLogConfigResponseBodyData {
	s.HotRetention = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetHotStartTime(v int64) *DescribeSqlLogConfigResponseBodyData {
	s.HotStartTime = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetLogFilter(v string) *DescribeSqlLogConfigResponseBodyData {
	s.LogFilter = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetRequestEnable(v bool) *DescribeSqlLogConfigResponseBodyData {
	s.RequestEnable = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetRequestStartTime(v int64) *DescribeSqlLogConfigResponseBodyData {
	s.RequestStartTime = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetRequestStopTime(v int64) *DescribeSqlLogConfigResponseBodyData {
	s.RequestStopTime = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetRetention(v int32) *DescribeSqlLogConfigResponseBodyData {
	s.Retention = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSqlLogEnable(v bool) *DescribeSqlLogConfigResponseBodyData {
	s.SqlLogEnable = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSqlLogSource(v string) *DescribeSqlLogConfigResponseBodyData {
	s.SqlLogSource = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSqlLogState(v string) *DescribeSqlLogConfigResponseBodyData {
	s.SqlLogState = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSqlLogVisibleTime(v int64) *DescribeSqlLogConfigResponseBodyData {
	s.SqlLogVisibleTime = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSupportMigration(v bool) *DescribeSqlLogConfigResponseBodyData {
	s.SupportMigration = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetSupportVersion(v string) *DescribeSqlLogConfigResponseBodyData {
	s.SupportVersion = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) SetVersion(v string) *DescribeSqlLogConfigResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeSqlLogConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
