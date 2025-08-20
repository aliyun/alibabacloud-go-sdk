// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkAppInfo) *GetSparkAppInfoResponseBody
	GetData() *SparkAppInfo
	SetRequestId(v string) *GetSparkAppInfoResponseBody
	GetRequestId() *string
}

type GetSparkAppInfoResponseBody struct {
	// The queried Spark application. Fields in the response parameter:
	//
	// 	- **Data**: the data of the Spark application template.
	//
	// 	- **EstimateExecutionCpuTimeInSeconds**: the amount of time that is required to consume CPU resources for running the Spark application. Unit: milliseconds.
	//
	// 	- **LogRootPath**: the storage path of log files.
	//
	// 	- **LastAttemptId**: the most recent attempt ID.
	//
	// 	- **WebUiAddress**: the web UI URL.
	//
	// 	- **SubmittedTimeInMillis**: the time when the Spark application was submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **StartedTimeInMillis**: the time when the Spark application was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **TerminatedTimeInMillis**: the time when the Spark application was terminated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// 	- **DBClusterId**: the ID of the cluster on which the Spark application runs.
	//
	// 	- **ResourceGroupName**: the name of the job resource group.
	//
	// 	- **DurationInMillis**: the amount of time that is required to run the Spark application. Unit: milliseconds.
	//
	// example:
	//
	// {     \\"name\\": \\"SparkPi\\",     \\"file\\": \\"local:///tmp/spark-examples.jar\\",     \\"className\\": \\"org.apache.spark.examples.SparkPi\\",     \\"args\\": [         \\"1000000\\"     ],     \\"conf\\": {         \\"spark.driver.resourceSpec\\": \\"small\\",         \\"spark.executor.instances\\": 1,         \\"spark.executor.resourceSpec\\": \\"small\\"     } }",
	//
	//       "EstimateExecutionCpuTimeInSeconds" : 100,
	//
	//       "LogRootPath" : "oss://test/logs/driver",
	//
	//       "LastAttemptId" : "s202204291426hzpre60cfabb0000004-0003",
	//
	//       "WebUiAddress" : "https://sparkui.aliyuncs.com/token=xxx",
	//
	//       "SubmittedTimeInMillis" : 1651213645000,
	//
	//       "StartedTimeInMillis" : 1651213645010,
	//
	//       "LastUpdatedTimeInMillis" : 1651213645200,
	//
	//       "TerminatedTimeInMillis" : 1651213645300,
	//
	//       "DBClusterId" : "am-dbclusterid",
	//
	//       "ResourceGroupName" : "spark-rg",
	//
	//       "DurationInMillis" : 100
	//
	//     }
	Data *SparkAppInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppInfoResponseBody) GetData() *SparkAppInfo {
	return s.Data
}

func (s *GetSparkAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppInfoResponseBody) SetData(v *SparkAppInfo) *GetSparkAppInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppInfoResponseBody) SetRequestId(v string) *GetSparkAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
