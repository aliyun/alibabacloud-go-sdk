// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkSQLEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartSparkSQLEngineResponseBodyData) *StartSparkSQLEngineResponseBody
	GetData() *StartSparkSQLEngineResponseBodyData
	SetRequestId(v string) *StartSparkSQLEngineResponseBody
	GetRequestId() *string
}

type StartSparkSQLEngineResponseBody struct {
	// The returned data.
	Data *StartSparkSQLEngineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartSparkSQLEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSparkSQLEngineResponseBody) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponseBody) GetData() *StartSparkSQLEngineResponseBodyData {
	return s.Data
}

func (s *StartSparkSQLEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSparkSQLEngineResponseBody) SetData(v *StartSparkSQLEngineResponseBodyData) *StartSparkSQLEngineResponseBody {
	s.Data = v
	return s
}

func (s *StartSparkSQLEngineResponseBody) SetRequestId(v string) *StartSparkSQLEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSparkSQLEngineResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartSparkSQLEngineResponseBodyData struct {
	// The ID of the Spark job.
	//
	// example:
	//
	// s202301xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the Spark application.
	//
	// example:
	//
	// SQLEngine1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The state of the Spark SQL engine. Valid values:
	//
	// 	- SUBMITTED
	//
	// 	- STARTING
	//
	// 	- RUNNING
	//
	// 	- FAILED
	//
	// example:
	//
	// SUBMITTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s StartSparkSQLEngineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartSparkSQLEngineResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartSparkSQLEngineResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *StartSparkSQLEngineResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *StartSparkSQLEngineResponseBodyData) GetState() *string {
	return s.State
}

func (s *StartSparkSQLEngineResponseBodyData) SetAppId(v string) *StartSparkSQLEngineResponseBodyData {
	s.AppId = &v
	return s
}

func (s *StartSparkSQLEngineResponseBodyData) SetAppName(v string) *StartSparkSQLEngineResponseBodyData {
	s.AppName = &v
	return s
}

func (s *StartSparkSQLEngineResponseBodyData) SetState(v string) *StartSparkSQLEngineResponseBodyData {
	s.State = &v
	return s
}

func (s *StartSparkSQLEngineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
