// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPreCheckDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedTime(v int64) *QueryPreCheckDatabaseResponseBody
	GetCompletedTime() *int64
	SetCreatedTime(v int64) *QueryPreCheckDatabaseResponseBody
	GetCreatedTime() *int64
	SetDescription(v string) *QueryPreCheckDatabaseResponseBody
	GetDescription() *string
	SetProgress(v int32) *QueryPreCheckDatabaseResponseBody
	GetProgress() *int32
	SetRequestId(v string) *QueryPreCheckDatabaseResponseBody
	GetRequestId() *string
	SetResult(v string) *QueryPreCheckDatabaseResponseBody
	GetResult() *string
	SetUpdatedTime(v int64) *QueryPreCheckDatabaseResponseBody
	GetUpdatedTime() *int64
}

type QueryPreCheckDatabaseResponseBody struct {
	// The time when the precheck task was complete.
	//
	// example:
	//
	// 1657524396
	CompletedTime *int64 `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// The time when the precheck task was started.
	//
	// example:
	//
	// 1660448660
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The status of the precheck task. Valid values:
	//
	// 	- **completed**: complete
	//
	// 	- **created**: started
	//
	// 	- **error**: failed
	//
	// example:
	//
	// completed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The precheck progress in percentage. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the precheck task. The value is a JSON string that contains the following fields:
	//
	// 	- **instanceId**: the ID of the server that hosts the database
	//
	// 	- **checkTime**: the precheck time
	//
	// 	- **sourceType**: the database type
	//
	// 	- **results**: the precheck item and result
	//
	//     	- **item**: the precheck item
	//
	//     	- **result**: the precheck result
	//
	// > The following section describes the precheck items:
	//
	// 	- MSSQL
	//
	//     	- **OSS_INTERNAL_ENDPOINT_CONNECTIVITY**: OSS connectivity check
	//
	//     	- **SERVICE_CONNECTIVITY**: control network connectivity check
	//
	//     	- **SQL_SERVER_DB_IN_SIMPLE_RECOVERY_MODE**: recovery mode check
	//
	//     	- **SQL_SERVER_DB_NOT_ONLINE**: SQL Server database status check
	//
	// 	- ORACLE
	//
	//     	- **OSS_INTERNAL_ENDPOINT_CONNECTIVITY**: OSS connectivity check
	//
	//     	- **SERVICE_CONNECTIVITY**: control network connectivity check
	//
	//     	- **ORACLE_INSTANCE_STATUS**: Oracle instance status check
	//
	//     	- **ORACLE_DB_STATUS**: Oracle database status check
	//
	//     	- **ARCHIVELOG**: archive mode check
	//
	// 	- MYSQL
	//
	//     	- **OSS_INTERNAL_ENDPOINT_CONNECTIVITY**: OSS connectivity check
	//
	//     	- **SERVICE_CONNECTIVITY**: control network connectivity check
	//
	//     	- **MYSQL_VERSION**: Supports full backup version checking
	//
	//     	- **MYSQL_BINLOG**: BINLOG check
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "instanceId": "i-wz91if83t97xgtn2****",
	//
	//         "checkTime": 1671245753,
	//
	//         "sourceType": "MSSQL",
	//
	//         "results":
	//
	//         [
	//
	//             {
	//
	//                 "item": "OSS_INTERNAL_ENDPOINT_CONNECTIVITY",
	//
	//                 "result": "PASSED"
	//
	//             },
	//
	//             {
	//
	//                 "item": "SERVICE_CONNECTIVITY",
	//
	//                 "result": "PASSED"
	//
	//             },
	//
	//             {
	//
	//                 "item": "SQL_SERVER_DB_IN_SIMPLE_RECOVERY_MODE",
	//
	//                 "result": "WARNING"
	//
	//             },
	//
	//             {
	//
	//                 "item": "SQL_SERVER_DB_NOT_ONLINE",
	//
	//                 "result": "PASSED"
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// ]
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the precheck task was last updated.
	//
	// example:
	//
	// 1671084106
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s QueryPreCheckDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPreCheckDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPreCheckDatabaseResponseBody) GetCompletedTime() *int64 {
	return s.CompletedTime
}

func (s *QueryPreCheckDatabaseResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *QueryPreCheckDatabaseResponseBody) GetDescription() *string {
	return s.Description
}

func (s *QueryPreCheckDatabaseResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *QueryPreCheckDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPreCheckDatabaseResponseBody) GetResult() *string {
	return s.Result
}

func (s *QueryPreCheckDatabaseResponseBody) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *QueryPreCheckDatabaseResponseBody) SetCompletedTime(v int64) *QueryPreCheckDatabaseResponseBody {
	s.CompletedTime = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetCreatedTime(v int64) *QueryPreCheckDatabaseResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetDescription(v string) *QueryPreCheckDatabaseResponseBody {
	s.Description = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetProgress(v int32) *QueryPreCheckDatabaseResponseBody {
	s.Progress = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetRequestId(v string) *QueryPreCheckDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetResult(v string) *QueryPreCheckDatabaseResponseBody {
	s.Result = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) SetUpdatedTime(v int64) *QueryPreCheckDatabaseResponseBody {
	s.UpdatedTime = &v
	return s
}

func (s *QueryPreCheckDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
