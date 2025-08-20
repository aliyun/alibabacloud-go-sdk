// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *KillSparkAppResponseBodyData) *KillSparkAppResponseBody
	GetData() *KillSparkAppResponseBodyData
	SetRequestId(v string) *KillSparkAppResponseBody
	GetRequestId() *string
}

type KillSparkAppResponseBody struct {
	// The returned data.
	Data *KillSparkAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 69D0810B-F9F5-5F4C-A57F-DF36133B63C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillSparkAppResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponseBody) GetData() *KillSparkAppResponseBodyData {
	return s.Data
}

func (s *KillSparkAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillSparkAppResponseBody) SetData(v *KillSparkAppResponseBodyData) *KillSparkAppResponseBody {
	s.Data = v
	return s
}

func (s *KillSparkAppResponseBody) SetRequestId(v string) *KillSparkAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillSparkAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type KillSparkAppResponseBodyData struct {
	// The Spark application ID.
	//
	// example:
	//
	// s202204132018hzprec1ac****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// LAKEHOUSE-1-1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp1c3em7b2e****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// [Advisor] Advisor feature is not available for instance: am-2ze292w4fyglwxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution state of the Spark application. Valid values:
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
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s KillSparkAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KillSparkAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *KillSparkAppResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *KillSparkAppResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *KillSparkAppResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillSparkAppResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *KillSparkAppResponseBodyData) GetState() *string {
	return s.State
}

func (s *KillSparkAppResponseBodyData) SetAppId(v string) *KillSparkAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetAppName(v string) *KillSparkAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetDBClusterId(v string) *KillSparkAppResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetMessage(v string) *KillSparkAppResponseBodyData {
	s.Message = &v
	return s
}

func (s *KillSparkAppResponseBodyData) SetState(v string) *KillSparkAppResponseBodyData {
	s.State = &v
	return s
}

func (s *KillSparkAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
