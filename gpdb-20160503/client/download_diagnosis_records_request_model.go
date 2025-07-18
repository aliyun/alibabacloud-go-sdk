// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadDiagnosisRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DownloadDiagnosisRecordsRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DownloadDiagnosisRecordsRequest
	GetDatabase() *string
	SetEndTime(v string) *DownloadDiagnosisRecordsRequest
	GetEndTime() *string
	SetLang(v string) *DownloadDiagnosisRecordsRequest
	GetLang() *string
	SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest
	GetQueryCondition() *string
	SetResourceGroupId(v string) *DownloadDiagnosisRecordsRequest
	GetResourceGroupId() *string
	SetStartTime(v string) *DownloadDiagnosisRecordsRequest
	GetStartTime() *string
	SetUser(v string) *DownloadDiagnosisRecordsRequest
	GetUser() *string
}

type DownloadDiagnosisRecordsRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The end of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-07T07:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the file that contains the query diagnostic information. Valid values:
	//
	// 	- **zh**: simplified Chinese
	//
	// 	- **en**: English
	//
	// 	- **ja**: Japanese
	//
	// 	- **zh-tw**: traditional Chinese
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The filter condition on queries. The value is in the JSON format. Valid values:
	//
	// 	- `{"Type":"maxCost", "Value":"100"}`: filters the top 100 queries that are the most time-consuming.
	//
	// 	- `{"Type":"status","Value":"finished"}`: filters completed queries.
	//
	// 	- `{"Type":"status","Value":"running"}`: filters running queries.
	//
	// 	- `{"Type":"cost","Max":"200"}`: filters the queries that consume less than 200 milliseconds.
	//
	// 	- `{"Type":"cost","Min":"200","Max":"60000"}`: filters the queries that consume 200 milliseconds or more and less than 1 minute.
	//
	// 	- `{"Type":"cost","Min":"60000"}`: filters the queries that consume 1 minute or more.
	//
	// 	- `{"Type":"cost","Min":"30","Max":"50"}`: filters the queries that consume 30 milliseconds or more and less than 50 milliseconds. You can customize a filter condition by setting **Min*	- and **Max**.
	//
	// example:
	//
	// { "Type":"maxCost", "Value":"100" }
	QueryCondition *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-07T06:59Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DownloadDiagnosisRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadDiagnosisRecordsRequest) GoString() string {
	return s.String()
}

func (s *DownloadDiagnosisRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DownloadDiagnosisRecordsRequest) GetDatabase() *string {
	return s.Database
}

func (s *DownloadDiagnosisRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DownloadDiagnosisRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DownloadDiagnosisRecordsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DownloadDiagnosisRecordsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DownloadDiagnosisRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DownloadDiagnosisRecordsRequest) GetUser() *string {
	return s.User
}

func (s *DownloadDiagnosisRecordsRequest) SetDBInstanceId(v string) *DownloadDiagnosisRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetDatabase(v string) *DownloadDiagnosisRecordsRequest {
	s.Database = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetEndTime(v string) *DownloadDiagnosisRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetLang(v string) *DownloadDiagnosisRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetQueryCondition(v string) *DownloadDiagnosisRecordsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetResourceGroupId(v string) *DownloadDiagnosisRecordsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetStartTime(v string) *DownloadDiagnosisRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) SetUser(v string) *DownloadDiagnosisRecordsRequest {
	s.User = &v
	return s
}

func (s *DownloadDiagnosisRecordsRequest) Validate() error {
	return dara.Validate(s)
}
