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
	// The information about the Spark application. Response parameter description:
	//
	// - **Data**: the Spark application template data.
	//
	// - **EstimateExecutionCpuTimeInSeconds**: the CPU time consumed to execute the Spark application, in milliseconds.
	//
	// - **LogRootPath**: the storage path of the log file.
	//
	// - **LastAttemptId**: the retry ID.
	//
	// - **WebUiAddress**: the web UI address.
	//
	// - **SubmittedTimeInMillis**: the time when the Spark application was submitted. The value is a UNIX timestamp in milliseconds.
	//
	// - **StartedTimeInMillis**: the time when the Spark application was created. The value is a UNIX timestamp in milliseconds.
	//
	// - **LastUpdatedTimeInMillis**: the time when the Spark application was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// - **TerminatedTimeInMillis**: the time when the execution of the Spark application was terminated. The value is a UNIX timestamp in milliseconds.
	//
	// - **DBClusterId**: the ID of the cluster where the Spark application is executed.
	//
	// - **ResourceGroupName**: the name of the job resource group.
	//
	// - **DurationInMillis**: the execution duration of the Spark application, in milliseconds.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
