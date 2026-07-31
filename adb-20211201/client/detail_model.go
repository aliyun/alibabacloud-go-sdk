// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *Detail
	GetAppType() *string
	SetDBClusterId(v string) *Detail
	GetDBClusterId() *string
	SetData(v string) *Detail
	GetData() *string
	SetDurationInMillis(v int64) *Detail
	GetDurationInMillis() *int64
	SetEstimateExecutionCpuTimeInSeconds(v int64) *Detail
	GetEstimateExecutionCpuTimeInSeconds() *int64
	SetExecutionDurationInMillis(v int64) *Detail
	GetExecutionDurationInMillis() *int64
	SetLastAttemptId(v string) *Detail
	GetLastAttemptId() *string
	SetLastUpdatedTimeInMillis(v int64) *Detail
	GetLastUpdatedTimeInMillis() *int64
	SetLogRootPath(v string) *Detail
	GetLogRootPath() *string
	SetResourceGroupName(v string) *Detail
	GetResourceGroupName() *string
	SetResourceProvisioningDurationInMillis(v int64) *Detail
	GetResourceProvisioningDurationInMillis() *int64
	SetRunningStartTimeInMillis(v int64) *Detail
	GetRunningStartTimeInMillis() *int64
	SetStartedTimeInMillis(v int64) *Detail
	GetStartedTimeInMillis() *int64
	SetSubmittedTimeInMillis(v int64) *Detail
	GetSubmittedTimeInMillis() *int64
	SetTerminatedTimeInMillis(v int64) *Detail
	GetTerminatedTimeInMillis() *int64
	SetWebUiAddress(v string) *Detail
	GetWebUiAddress() *string
}

type Detail struct {
	// The type of the Spark application.
	//
	// example:
	//
	// BATCH
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The ID of the cluster that runs the Spark application.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The configuration content of the Spark application.
	//
	// example:
	//
	// {     "name": "SparkPi",     "file": "local:///tmp/spark-examples.jar",     "className": "org.apache.spark.examples.SparkPi",     "args": [         "1000000"     ],     "conf": {         "spark.driver.resourceSpec": "small",         "spark.executor.instances": 1,         "spark.executor.resourceSpec": "small"     } }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The execution duration of the Spark application, in milliseconds (ms).
	//
	// example:
	//
	// 100
	DurationInMillis *int64 `json:"DurationInMillis,omitempty" xml:"DurationInMillis,omitempty"`
	// The CPU time consumed by the Spark application, in milliseconds (ms).
	//
	// example:
	//
	// 100
	EstimateExecutionCpuTimeInSeconds *int64 `json:"EstimateExecutionCpuTimeInSeconds,omitempty" xml:"EstimateExecutionCpuTimeInSeconds,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 36000
	ExecutionDurationInMillis *int64 `json:"ExecutionDurationInMillis,omitempty" xml:"ExecutionDurationInMillis,omitempty"`
	// The ID of the last retry of the Spark application.
	//
	// example:
	//
	// s202204291426hzpre60****-0003
	LastAttemptId *string `json:"LastAttemptId,omitempty" xml:"LastAttemptId,omitempty"`
	// The time when the Spark application was last updated. This value is a UNIX timestamp, in milliseconds (ms).
	//
	// example:
	//
	// 1651213645200
	LastUpdatedTimeInMillis *int64 `json:"LastUpdatedTimeInMillis,omitempty" xml:"LastUpdatedTimeInMillis,omitempty"`
	// The storage path of the log file.
	//
	// example:
	//
	// oss://<bucket-name>/logs/driver
	LogRootPath *string `json:"LogRootPath,omitempty" xml:"LogRootPath,omitempty"`
	// The name of the job resource group.
	//
	// example:
	//
	// spark-rg
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The duration of resource provisioning for the application.
	//
	// example:
	//
	// 36000
	ResourceProvisioningDurationInMillis *int64 `json:"ResourceProvisioningDurationInMillis,omitempty" xml:"ResourceProvisioningDurationInMillis,omitempty"`
	// The timestamp when the job started running.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 36000
	RunningStartTimeInMillis *int64 `json:"RunningStartTimeInMillis,omitempty" xml:"RunningStartTimeInMillis,omitempty"`
	// The time when the Spark application was created. This value is a UNIX timestamp, in milliseconds (ms).
	//
	// example:
	//
	// 1651213645010
	StartedTimeInMillis *int64 `json:"StartedTimeInMillis,omitempty" xml:"StartedTimeInMillis,omitempty"`
	// The time when the Spark application was submitted. This value is a UNIX timestamp, in milliseconds (ms).
	//
	// example:
	//
	// 1651213645000
	SubmittedTimeInMillis *int64 `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
	// The time when the Spark application was terminated. This value is a UNIX timestamp, in milliseconds (ms).
	//
	// example:
	//
	// 1651213645300
	TerminatedTimeInMillis *int64 `json:"TerminatedTimeInMillis,omitempty" xml:"TerminatedTimeInMillis,omitempty"`
	// The Web UI address.
	//
	// example:
	//
	// https://adbsparkui-cn-hangzhou.aliyuncs.com/?token=****
	WebUiAddress *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s Detail) String() string {
	return dara.Prettify(s)
}

func (s Detail) GoString() string {
	return s.String()
}

func (s *Detail) GetAppType() *string {
	return s.AppType
}

func (s *Detail) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *Detail) GetData() *string {
	return s.Data
}

func (s *Detail) GetDurationInMillis() *int64 {
	return s.DurationInMillis
}

func (s *Detail) GetEstimateExecutionCpuTimeInSeconds() *int64 {
	return s.EstimateExecutionCpuTimeInSeconds
}

func (s *Detail) GetExecutionDurationInMillis() *int64 {
	return s.ExecutionDurationInMillis
}

func (s *Detail) GetLastAttemptId() *string {
	return s.LastAttemptId
}

func (s *Detail) GetLastUpdatedTimeInMillis() *int64 {
	return s.LastUpdatedTimeInMillis
}

func (s *Detail) GetLogRootPath() *string {
	return s.LogRootPath
}

func (s *Detail) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *Detail) GetResourceProvisioningDurationInMillis() *int64 {
	return s.ResourceProvisioningDurationInMillis
}

func (s *Detail) GetRunningStartTimeInMillis() *int64 {
	return s.RunningStartTimeInMillis
}

func (s *Detail) GetStartedTimeInMillis() *int64 {
	return s.StartedTimeInMillis
}

func (s *Detail) GetSubmittedTimeInMillis() *int64 {
	return s.SubmittedTimeInMillis
}

func (s *Detail) GetTerminatedTimeInMillis() *int64 {
	return s.TerminatedTimeInMillis
}

func (s *Detail) GetWebUiAddress() *string {
	return s.WebUiAddress
}

func (s *Detail) SetAppType(v string) *Detail {
	s.AppType = &v
	return s
}

func (s *Detail) SetDBClusterId(v string) *Detail {
	s.DBClusterId = &v
	return s
}

func (s *Detail) SetData(v string) *Detail {
	s.Data = &v
	return s
}

func (s *Detail) SetDurationInMillis(v int64) *Detail {
	s.DurationInMillis = &v
	return s
}

func (s *Detail) SetEstimateExecutionCpuTimeInSeconds(v int64) *Detail {
	s.EstimateExecutionCpuTimeInSeconds = &v
	return s
}

func (s *Detail) SetExecutionDurationInMillis(v int64) *Detail {
	s.ExecutionDurationInMillis = &v
	return s
}

func (s *Detail) SetLastAttemptId(v string) *Detail {
	s.LastAttemptId = &v
	return s
}

func (s *Detail) SetLastUpdatedTimeInMillis(v int64) *Detail {
	s.LastUpdatedTimeInMillis = &v
	return s
}

func (s *Detail) SetLogRootPath(v string) *Detail {
	s.LogRootPath = &v
	return s
}

func (s *Detail) SetResourceGroupName(v string) *Detail {
	s.ResourceGroupName = &v
	return s
}

func (s *Detail) SetResourceProvisioningDurationInMillis(v int64) *Detail {
	s.ResourceProvisioningDurationInMillis = &v
	return s
}

func (s *Detail) SetRunningStartTimeInMillis(v int64) *Detail {
	s.RunningStartTimeInMillis = &v
	return s
}

func (s *Detail) SetStartedTimeInMillis(v int64) *Detail {
	s.StartedTimeInMillis = &v
	return s
}

func (s *Detail) SetSubmittedTimeInMillis(v int64) *Detail {
	s.SubmittedTimeInMillis = &v
	return s
}

func (s *Detail) SetTerminatedTimeInMillis(v int64) *Detail {
	s.TerminatedTimeInMillis = &v
	return s
}

func (s *Detail) SetWebUiAddress(v string) *Detail {
	s.WebUiAddress = &v
	return s
}

func (s *Detail) Validate() error {
	return dara.Validate(s)
}
