// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowQueryStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetByDatabase(v interface{}) *DescribeSlowQueryStatsResponseBody
	GetByDatabase() interface{}
	SetByTimeBucket(v interface{}) *DescribeSlowQueryStatsResponseBody
	GetByTimeBucket() interface{}
	SetByUser(v interface{}) *DescribeSlowQueryStatsResponseBody
	GetByUser() interface{}
	SetPercentiles(v map[string]interface{}) *DescribeSlowQueryStatsResponseBody
	GetPercentiles() map[string]interface{}
	SetRequestId(v string) *DescribeSlowQueryStatsResponseBody
	GetRequestId() *string
	SetSummary(v map[string]interface{}) *DescribeSlowQueryStatsResponseBody
	GetSummary() map[string]interface{}
	SetTopQueries(v interface{}) *DescribeSlowQueryStatsResponseBody
	GetTopQueries() interface{}
	SetTopSqlDigests(v interface{}) *DescribeSlowQueryStatsResponseBody
	GetTopSqlDigests() interface{}
}

type DescribeSlowQueryStatsResponseBody struct {
	// An array of slow query statistics, grouped by database.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "maxScanBytes": 271603242215,
	//
	//       "minScanRows": 550584606,
	//
	//       "totalScanRows": 96977924056,
	//
	//       "avgLatency": 10804,
	//
	//       "minLatency": 5138,
	//
	//       "maxLatency": 24746,
	//
	//       "totalScanBytes": 1441239240695,
	//
	//       "queryCount": 29,
	//
	//       "avgScanBytes": 49697904851,
	//
	//       "minScanBytes": 8691897406,
	//
	//       "totalLatency": 313322,
	//
	//       "maxScanRows": 23040660808,
	//
	//       "avgCpuTimeMs": 732085,
	//
	//       "avgScanRows": 3344066346,
	//
	//       "maxCpuTimeMs": 1368932,
	//
	//       "db": "tpcds_1000g",
	//
	//       "totalCpuTimeMs": 21230477,
	//
	//       "minCpuTimeMs": 292711
	//
	//     }
	//
	// ]
	ByDatabase interface{} `json:"ByDatabase,omitempty" xml:"ByDatabase,omitempty"`
	// An array of slow query statistics, grouped by time bucket.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "maxScanBytes": 261400873691,
	//
	//       "minScanRows": 5915582721,
	//
	//       "totalScanRows": 23662330884,
	//
	//       "avgLatency": 6988.5,
	//
	//       "minLatency": 6657,
	//
	//       "maxLatency": 7401,
	//
	//       "totalScanBytes": 1045601376995,
	//
	//       "queryCount": 4,
	//
	//       "avgScanBytes": 261400344248.75,
	//
	//       "minScanBytes": 261400167768,
	//
	//       "maxScanRows": 5915582721,
	//
	//       "avgCpuTimeMs": 603488.25,
	//
	//       "avgScanRows": 5915582721,
	//
	//       "maxCpuTimeMs": 637311,
	//
	//       "timeBucket": "2026-04-15 22:00:00",
	//
	//       "totalCpuTimeMs": 2413953,
	//
	//       "minCpuTimeMs": 567462
	//
	//     }
	//
	// ]
	ByTimeBucket interface{} `json:"ByTimeBucket,omitempty" xml:"ByTimeBucket,omitempty"`
	// An array of slow query statistics, grouped by user.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "maxScanBytes": 279794422882,
	//
	//       "minScanRows": 0,
	//
	//       "totalScanRows": 214219674180,
	//
	//       "avgLatency": 9862.755555555555,
	//
	//       "minLatency": 5051,
	//
	//       "maxLatency": 24746,
	//
	//       "totalScanBytes": 4657224498428,
	//
	//       "queryCount": 45,
	//
	//       "avgScanBytes": 103493877742.84445,
	//
	//       "minScanBytes": 0,
	//
	//       "totalLatency": 443824,
	//
	//       "maxScanRows": 23040660808,
	//
	//       "avgCpuTimeMs": 649451.7777777778,
	//
	//       "avgScanRows": 4760437204,
	//
	//       "maxCpuTimeMs": 1368932,
	//
	//       "distinctSqlDigests": 1,
	//
	//       "user": "admin",
	//
	//       "totalCpuTimeMs": 29225330,
	//
	//       "minCpuTimeMs": 2434
	//
	//     }
	//
	//   ]
	ByUser interface{} `json:"ByUser,omitempty" xml:"ByUser,omitempty"`
	// The percentile statistics for query latency.
	//
	// example:
	//
	// {
	//
	//     "p99": "24746.0",
	//
	//     "p50": "8295.0",
	//
	//     "p95": "23872.0",
	//
	//     "p90": "15794.0",
	//
	//     "p75": "11972.25"
	//
	//   }
	Percentiles map[string]interface{} `json:"Percentiles,omitempty" xml:"Percentiles,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FCEEA97F-XXXX-XXXX-932F-B4BAEA170896
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The summary of slow query statistics.
	//
	// example:
	//
	// {
	//
	//     "totalAllQueries": "4531",
	//
	//     "maxScanBytes": "279794422882",
	//
	//     "totalQueries": "45",
	//
	//     "minScanRows": "0",
	//
	//     "totalScanRows": "214219674180",
	//
	//     "errorQueryCount": "0",
	//
	//     "avgLatency": "9862",
	//
	//     "minLatency": "5051",
	//
	//     "maxLatency": "24746",
	//
	//     "totalScanBytes": "4657224498428",
	//
	//     "avgScanBytes": "103493877742",
	//
	//     "distinctUsers": "1",
	//
	//     "minScanBytes": "0",
	//
	//     "slowQueryRatio": "0.009931582432134187",
	//
	//     "maxScanRows": "23040660808",
	//
	//     "avgCpuTimeMs": "649451",
	//
	//     "avgScanRows": "4760437204",
	//
	//     "maxCpuTimeMs": "1368932",
	//
	//     "totalCpuTimeMs": "29225330"
	//
	//   }
	Summary map[string]interface{} `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// An array of detailed audit records for the top N slow queries.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "peakMemoryBytes": 71731165356,
	//
	//         "planTimesMs": "{\\"plan\\":15, \\"garbage_collect\\":0, \\"lock_tables\\":0, \\"analyze\\":2, \\"rewrite\\":4, \\"fold_const_by_be\\":0, \\"collect_partitions\\":0, \\"optimize\\":3, \\"translate\\":2, \\"init_scan_node\\":-1, \\"finalize_scan_node\\":-1, \\"create_scan_range\\":-1, \\"distribute\\":4}",
	//
	//         "catalog": "internal",
	//
	//         "sqlMode": "ONLY_FULL_GROUP_BY",
	//
	//         "errorCode": 0,
	//
	//         "spillReadBytesFromLocalStorage": -1,
	//
	//         "computeGroup": "test",
	//
	//         "queryId": "921f10bc054d4dbb-903784520a2ab26f",
	//
	//         "parseTimeMs": 0,
	//
	//         "stmtId": 5822,
	//
	//         "hitSqlCache": false,
	//
	//         "chosenMViews": "[]",
	//
	//         "scanBytesFromLocalStorage": 861079963,
	//
	//         "frontendIp": "172.16.17.192",
	//
	//         "handledInFe": false,
	//
	//         "returnRows": 100,
	//
	//         "state": "EOF",
	//
	//         "scanRows": 550584606,
	//
	//         "cpuTimeMs": 1368932,
	//
	//         "scheduleTimesMs": "{\\"schedule_time_ms\\":14, \\"fragment_assign_time_ms\\":0, \\"fragment_serialize_time_ms\\":2, \\"fragment_rpc_phase_1_time_ms\\":11, \\"fragment_rpc_phase_2_time_ms\\":1, \\"fragment_compressed_size_byte\\":142152, \\"fragment_rpc_count\\":6}",
	//
	//         "shuffleSendBytes": 27745167331,
	//
	//         "stmtType": "SELECT",
	//
	//         "sqlHash": "ecf08bbca3e4b33b1630e03a9dc671b4",
	//
	//         "errorMessage": "",
	//
	//         "isQuery": true,
	//
	//         "isNereids": true,
	//
	//         "changedVariables": "{\\"enable_profile\\":\\"true\\", \\"enable_auto_analyze\\":\\"false\\", \\"runtime_filter_wait_time_ms\\":\\"10000\\", \\"sql_converter_service_url\\":\\"http://127.0.0.1:5001/api/v1/convert\\"}",
	//
	//         "scanBytesFromRemoteStorage": 0,
	//
	//         "scanBytes": 11358721363,
	//
	//         "isInternal": false,
	//
	//         "workloadGroup": "normal",
	//
	//         "queriedTablesAndViews": "[\\"internal.tpcds_1000g.item\\", \\"internal.tpcds_1000g.store_sales\\", \\"internal.tpcds_1000g.date_dim\\", \\"internal.tpcds_1000g.store\\"]",
	//
	//         "sqlDigest": "d41d8cd98f00b204e9800998ecf8427e",
	//
	//         "clientIp": "123.56.117.27:33664",
	//
	//         "queryTime": 24746,
	//
	//         "shuffleSendRows": 454538612,
	//
	//         "time": "1776267563848",
	//
	//         "getMetaTimesMs": "{\\"get_partition_version_time_ms\\":3381937, \\"get_partition_version_count_has_data\\":0, \\"get_partition_version_count\\":1, \\"get_table_version_time_ms\\":0, \\"get_table_version_count\\":0}",
	//
	//         "spillWriteBytesFromLocalStorage": -1,
	//
	//         "user": "admin",
	//
	//         "db": "tpcds_1000g",
	//
	//         "stmt": "xxx"
	//
	//     }
	//
	// ]
	TopQueries interface{} `json:"TopQueries,omitempty" xml:"TopQueries,omitempty"`
	// An array of statistics for the top N slow queries, grouped by SQL digest. Available for kernel version 5.0 and later.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "maxScanBytes": 279794422882,
	//
	//       "minScanRows": 0,
	//
	//       "totalScanRows": 214219674180,
	//
	//       "avgLatency": 9862.755555555555,
	//
	//       "minLatency": 5051,
	//
	//       "maxLatency": 24746,
	//
	//       "totalScanBytes": 4657224498428,
	//
	//       "queryCount": 45,
	//
	//       "avgScanBytes": 103493877742.84445,
	//
	//       "minScanBytes": 0,
	//
	//       "totalLatency": 443824,
	//
	//       "sqlDigest": "d41d8cd98f00b204e9800998ecf8427e",
	//
	//       "maxScanRows": 23040660808,
	//
	//       "avgCpuTimeMs": 649451.7777777778,
	//
	//       "avgScanRows": 4760437204,
	//
	//       "maxCpuTimeMs": 1368932,
	//
	//       "sampleStmt": "xxx",
	//
	//       "user": "admin",
	//
	//       "db": "tpcds_1000g",
	//
	//       "totalCpuTimeMs": 29225330,
	//
	//       "minCpuTimeMs": 2434
	//
	//     }
	//
	//   ]
	TopSqlDigests interface{} `json:"TopSqlDigests,omitempty" xml:"TopSqlDigests,omitempty"`
}

func (s DescribeSlowQueryStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatsResponseBody) GetByDatabase() interface{} {
	return s.ByDatabase
}

func (s *DescribeSlowQueryStatsResponseBody) GetByTimeBucket() interface{} {
	return s.ByTimeBucket
}

func (s *DescribeSlowQueryStatsResponseBody) GetByUser() interface{} {
	return s.ByUser
}

func (s *DescribeSlowQueryStatsResponseBody) GetPercentiles() map[string]interface{} {
	return s.Percentiles
}

func (s *DescribeSlowQueryStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlowQueryStatsResponseBody) GetSummary() map[string]interface{} {
	return s.Summary
}

func (s *DescribeSlowQueryStatsResponseBody) GetTopQueries() interface{} {
	return s.TopQueries
}

func (s *DescribeSlowQueryStatsResponseBody) GetTopSqlDigests() interface{} {
	return s.TopSqlDigests
}

func (s *DescribeSlowQueryStatsResponseBody) SetByDatabase(v interface{}) *DescribeSlowQueryStatsResponseBody {
	s.ByDatabase = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetByTimeBucket(v interface{}) *DescribeSlowQueryStatsResponseBody {
	s.ByTimeBucket = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetByUser(v interface{}) *DescribeSlowQueryStatsResponseBody {
	s.ByUser = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetPercentiles(v map[string]interface{}) *DescribeSlowQueryStatsResponseBody {
	s.Percentiles = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetRequestId(v string) *DescribeSlowQueryStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetSummary(v map[string]interface{}) *DescribeSlowQueryStatsResponseBody {
	s.Summary = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetTopQueries(v interface{}) *DescribeSlowQueryStatsResponseBody {
	s.TopQueries = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) SetTopSqlDigests(v interface{}) *DescribeSlowQueryStatsResponseBody {
	s.TopSqlDigests = v
	return s
}

func (s *DescribeSlowQueryStatsResponseBody) Validate() error {
	return dara.Validate(s)
}
