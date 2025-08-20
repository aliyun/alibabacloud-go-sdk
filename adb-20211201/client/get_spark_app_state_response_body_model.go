// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkAppStateResponseBodyData) *GetSparkAppStateResponseBody
	GetData() *GetSparkAppStateResponseBodyData
	SetRequestId(v string) *GetSparkAppStateResponseBody
	GetRequestId() *string
}

type GetSparkAppStateResponseBody struct {
	// The returned data.
	Data *GetSparkAppStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponseBody) GetData() *GetSparkAppStateResponseBodyData {
	return s.Data
}

func (s *GetSparkAppStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppStateResponseBody) SetData(v *GetSparkAppStateResponseBodyData) *GetSparkAppStateResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppStateResponseBody) SetRequestId(v string) *GetSparkAppStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppStateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSparkAppStateResponseBodyData struct {
	// The Spark application ID.
	//
	// example:
	//
	// s202204191546hzpread6a896000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The alert message returned for the operation, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// Insufficient resources.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the application. Valid values:
	//
	// 	- **SUBMITTED**
	//
	// 	- **STARTING**
	//
	// 	- **RUNNING**
	//
	// 	- **FAILING**
	//
	// 	- **FAILED**
	//
	// 	- **KILLING**
	//
	// 	- **KILLED**
	//
	// 	- **SUCCEEDING**
	//
	// 	- **COMPLETED**
	//
	// 	- **FATAL**
	//
	// 	- **UNKNOWN**
	//
	// example:
	//
	// COMPLETED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSparkAppStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppStateResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppStateResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetSparkAppStateResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppStateResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetSparkAppStateResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetSparkAppStateResponseBodyData) SetAppId(v string) *GetSparkAppStateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetAppName(v string) *GetSparkAppStateResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetDBClusterId(v string) *GetSparkAppStateResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetMessage(v string) *GetSparkAppStateResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) SetState(v string) *GetSparkAppStateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSparkAppStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
