// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkSQLEngineStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkSQLEngineStateResponseBodyData) *GetSparkSQLEngineStateResponseBody
	GetData() *GetSparkSQLEngineStateResponseBodyData
	SetRequestId(v string) *GetSparkSQLEngineStateResponseBody
	GetRequestId() *string
}

type GetSparkSQLEngineStateResponseBody struct {
	// The state information about the Spark SQL engine.
	Data *GetSparkSQLEngineStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxx-xx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkSQLEngineStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkSQLEngineStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponseBody) GetData() *GetSparkSQLEngineStateResponseBodyData {
	return s.Data
}

func (s *GetSparkSQLEngineStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkSQLEngineStateResponseBody) SetData(v *GetSparkSQLEngineStateResponseBodyData) *GetSparkSQLEngineStateResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkSQLEngineStateResponseBody) SetRequestId(v string) *GetSparkSQLEngineStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSparkSQLEngineStateResponseBodyData struct {
	// The ID of the Spark application.
	//
	// example:
	//
	// s202207151211hz0c****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configuration of the Spark application.
	//
	// example:
	//
	// {"key1": "value1", "key2": "value2"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The third-party JAR package.
	//
	// example:
	//
	// oss://test-bucket/test.jar
	Jars *string `json:"Jars,omitempty" xml:"Jars,omitempty"`
	// The maximum number of started Spark executors.
	//
	// example:
	//
	// 3
	MaxExecutor *string `json:"MaxExecutor,omitempty" xml:"MaxExecutor,omitempty"`
	// The minimum number of started Spark executors.
	//
	// example:
	//
	// 1
	MinExecutor *string `json:"MinExecutor,omitempty" xml:"MinExecutor,omitempty"`
	// The slot number of the Spark application.
	//
	// example:
	//
	// 2
	SlotNum *string `json:"SlotNum,omitempty" xml:"SlotNum,omitempty"`
	// The execution state of the application. Valid values:
	//
	// 	- SUBMITTED
	//
	// 	- STARTING
	//
	// 	- RUNNING
	//
	// 	- FAILING
	//
	// 	- FAILED
	//
	// 	- KILLING
	//
	// 	- KILLED
	//
	// 	- SUCCEEDING
	//
	// 	- COMPLETED
	//
	// 	- FATAL
	//
	// 	- UNKNOWN
	//
	// example:
	//
	// COMPLETED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The timestamp when the Spark SQL application was submitted. Unit: milliseconds.
	//
	// example:
	//
	// 1651213645000
	SubmittedTimeInMillis *string `json:"SubmittedTimeInMillis,omitempty" xml:"SubmittedTimeInMillis,omitempty"`
}

func (s GetSparkSQLEngineStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkSQLEngineStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetJars() *string {
	return s.Jars
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetMaxExecutor() *string {
	return s.MaxExecutor
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetMinExecutor() *string {
	return s.MinExecutor
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetSlotNum() *string {
	return s.SlotNum
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetSparkSQLEngineStateResponseBodyData) GetSubmittedTimeInMillis() *string {
	return s.SubmittedTimeInMillis
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetAppId(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetConfig(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetJars(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.Jars = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetMaxExecutor(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.MaxExecutor = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetMinExecutor(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.MinExecutor = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetSlotNum(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.SlotNum = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetState(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) SetSubmittedTimeInMillis(v string) *GetSparkSQLEngineStateResponseBodyData {
	s.SubmittedTimeInMillis = &v
	return s
}

func (s *GetSparkSQLEngineStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
