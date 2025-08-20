// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkSQLDiagnosisAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetAppId(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetAppId() *string
	SetDiagnosisInfos(v []*Adb4MysqlSparkDiagnosisInfo) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetDiagnosisInfos() []*Adb4MysqlSparkDiagnosisInfo
	SetElapsedTime(v int64) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetElapsedTime() *int64
	SetInnerQueryId(v int64) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetInnerQueryId() *int64
	SetOperatorListSortedByMetrics(v *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetOperatorListSortedByMetrics() *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics
	SetRequestId(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetRequestId() *string
	SetRoot(v *OperatorNode) *DescribeSparkSQLDiagnosisAttributeResponseBody
	GetRoot() *OperatorNode
}

type DescribeSparkSQLDiagnosisAttributeResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The application ID.
	//
	// >  You can call the [ListSparkApps](https://help.aliyun.com/document_detail/612475.html) operation to query a list of Spark application IDs.
	//
	// example:
	//
	// s202411071444hzdvk486d9d2001****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The queried diagnostic information.
	DiagnosisInfos []*Adb4MysqlSparkDiagnosisInfo `json:"DiagnosisInfos,omitempty" xml:"DiagnosisInfos,omitempty" type:"Repeated"`
	// The execution duration of the query. Unit: milliseconds.
	//
	// example:
	//
	// 100
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The ID of the query executed within the Spark application.
	//
	// example:
	//
	// 1
	InnerQueryId *int64 `json:"InnerQueryId,omitempty" xml:"InnerQueryId,omitempty"`
	// The operators sorted by metrics.
	OperatorListSortedByMetrics *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics `json:"OperatorListSortedByMetrics,omitempty" xml:"OperatorListSortedByMetrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Spark execution plan tree.
	Root *OperatorNode `json:"Root,omitempty" xml:"Root,omitempty"`
}

func (s DescribeSparkSQLDiagnosisAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetDiagnosisInfos() []*Adb4MysqlSparkDiagnosisInfo {
	return s.DiagnosisInfos
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetInnerQueryId() *int64 {
	return s.InnerQueryId
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetOperatorListSortedByMetrics() *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics {
	return s.OperatorListSortedByMetrics
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) GetRoot() *OperatorNode {
	return s.Root
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetAppId(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetDiagnosisInfos(v []*Adb4MysqlSparkDiagnosisInfo) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.DiagnosisInfos = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetElapsedTime(v int64) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetInnerQueryId(v int64) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.InnerQueryId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetOperatorListSortedByMetrics(v *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.OperatorListSortedByMetrics = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetRequestId(v string) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) SetRoot(v *OperatorNode) *DescribeSparkSQLDiagnosisAttributeResponseBody {
	s.Root = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics struct {
	// The operators sorted by the execution duration.
	OperatorListSortedByExclusiveTime []*SparkOperatorInfo `json:"OperatorListSortedByExclusiveTime,omitempty" xml:"OperatorListSortedByExclusiveTime,omitempty" type:"Repeated"`
	// The operators sorted by the maximum memory used.
	OperatorListSortedByMaxMemory []*SparkOperatorInfo `json:"OperatorListSortedByMaxMemory,omitempty" xml:"OperatorListSortedByMaxMemory,omitempty" type:"Repeated"`
}

func (s DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) GoString() string {
	return s.String()
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) GetOperatorListSortedByExclusiveTime() []*SparkOperatorInfo {
	return s.OperatorListSortedByExclusiveTime
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) GetOperatorListSortedByMaxMemory() []*SparkOperatorInfo {
	return s.OperatorListSortedByMaxMemory
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) SetOperatorListSortedByExclusiveTime(v []*SparkOperatorInfo) *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics {
	s.OperatorListSortedByExclusiveTime = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) SetOperatorListSortedByMaxMemory(v []*SparkOperatorInfo) *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics {
	s.OperatorListSortedByMaxMemory = v
	return s
}

func (s *DescribeSparkSQLDiagnosisAttributeResponseBodyOperatorListSortedByMetrics) Validate() error {
	return dara.Validate(s)
}
