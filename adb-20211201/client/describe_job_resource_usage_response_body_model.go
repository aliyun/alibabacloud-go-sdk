// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeJobResourceUsageResponseBody
	GetCode() *int32
	SetData(v *DescribeJobResourceUsageResponseBodyData) *DescribeJobResourceUsageResponseBody
	GetData() *DescribeJobResourceUsageResponseBodyData
	SetRequestId(v string) *DescribeJobResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeJobResourceUsageResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeJobResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeJobResourceUsageResponseBody) GetData() *DescribeJobResourceUsageResponseBodyData {
	return s.Data
}

func (s *DescribeJobResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobResourceUsageResponseBody) SetCode(v int32) *DescribeJobResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBody) SetData(v *DescribeJobResourceUsageResponseBodyData) *DescribeJobResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeJobResourceUsageResponseBody) SetRequestId(v string) *DescribeJobResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResourceUsageResponseBodyData struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-23T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The AnalyticDB compute unit (ACU) usage of the job resource group.
	JobAcuUsage []*DescribeJobResourceUsageResponseBodyDataJobAcuUsage `json:"JobAcuUsage,omitempty" xml:"JobAcuUsage,omitempty" type:"Repeated"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeJobResourceUsageResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeJobResourceUsageResponseBodyData) GetJobAcuUsage() []*DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	return s.JobAcuUsage
}

func (s *DescribeJobResourceUsageResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeJobResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeJobResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetEndTime(v string) *DescribeJobResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetJobAcuUsage(v []*DescribeJobResourceUsageResponseBodyDataJobAcuUsage) *DescribeJobResourceUsageResponseBodyData {
	s.JobAcuUsage = v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) SetStartTime(v string) *DescribeJobResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResourceUsageResponseBodyDataJobAcuUsage struct {
	// The ACU usage.
	AcuUsageDetail *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail `json:"AcuUsageDetail,omitempty" xml:"AcuUsageDetail,omitempty" type:"Struct"`
	// The end time of the job. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-23T16:00:00Z
	JobEndTime *string `json:"JobEndTime,omitempty" xml:"JobEndTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 1592
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The start time of the job. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-22T16:00:00Z
	JobStartTime *string `json:"JobStartTime,omitempty" xml:"JobStartTime,omitempty"`
	// The name of the job resource group.
	//
	// example:
	//
	// job_default
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GetAcuUsageDetail() *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	return s.AcuUsageDetail
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GetJobEndTime() *string {
	return s.JobEndTime
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GetJobStartTime() *string {
	return s.JobStartTime
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetAcuUsageDetail(v *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.AcuUsageDetail = v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobEndTime(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobEndTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobId(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobId = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetJobStartTime(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.JobStartTime = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) SetResourceGroupName(v string) *DescribeJobResourceUsageResponseBodyDataJobAcuUsage {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail struct {
	// The number of ACUs for the elastic resources.
	//
	// example:
	//
	// 16ACU
	ElasticAcuNumber *float32 `json:"ElasticAcuNumber,omitempty" xml:"ElasticAcuNumber,omitempty"`
	// The number of ACUs for the reserved resources.
	//
	// example:
	//
	// 16ACU
	ReservedAcuNumber *float32 `json:"ReservedAcuNumber,omitempty" xml:"ReservedAcuNumber,omitempty"`
	// The number of spot ACUs.
	//
	// example:
	//
	// 16ACU
	SpotAcuNumber *float32 `json:"SpotAcuNumber,omitempty" xml:"SpotAcuNumber,omitempty"`
	// The percent of spot ACUs.
	//
	// example:
	//
	// 0.9
	SpotAcuPercentage *float32 `json:"SpotAcuPercentage,omitempty" xml:"SpotAcuPercentage,omitempty"`
	// The total number of ACUs.
	//
	// example:
	//
	// 32ACU
	TotalAcuNumber *float32 `json:"TotalAcuNumber,omitempty" xml:"TotalAcuNumber,omitempty"`
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GoString() string {
	return s.String()
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GetElasticAcuNumber() *float32 {
	return s.ElasticAcuNumber
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GetReservedAcuNumber() *float32 {
	return s.ReservedAcuNumber
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GetSpotAcuNumber() *float32 {
	return s.SpotAcuNumber
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GetSpotAcuPercentage() *float32 {
	return s.SpotAcuPercentage
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) GetTotalAcuNumber() *float32 {
	return s.TotalAcuNumber
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetElasticAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.ElasticAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetReservedAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.ReservedAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetSpotAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.SpotAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetSpotAcuPercentage(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.SpotAcuPercentage = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) SetTotalAcuNumber(v float32) *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail {
	s.TotalAcuNumber = &v
	return s
}

func (s *DescribeJobResourceUsageResponseBodyDataJobAcuUsageAcuUsageDetail) Validate() error {
	return dara.Validate(s)
}
