// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSparkAppInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SparkAppInfo
	GetAppId() *string
	SetAppName(v string) *SparkAppInfo
	GetAppName() *string
	SetDBClusterId(v string) *SparkAppInfo
	GetDBClusterId() *string
	SetDetail(v *Detail) *SparkAppInfo
	GetDetail() *Detail
	SetMessage(v string) *SparkAppInfo
	GetMessage() *string
	SetPriority(v string) *SparkAppInfo
	GetPriority() *string
	SetState(v string) *SparkAppInfo
	GetState() *string
}

type SparkAppInfo struct {
	// The ID of the Spark application.
	//
	// example:
	//
	// s202207151211hz0c****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the Spark application.
	//
	// example:
	//
	// SparkTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The database ID.
	//
	// example:
	//
	// amv-23xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the Spark application.
	Detail *Detail `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The alert message returned, such as task execution failure or insufficient resources. If no alert occurs, null is returned.
	//
	// example:
	//
	// WARN: Disk is full.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The priority of the Spark application.
	//
	// example:
	//
	// NORMAL
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The state of the Spark application. Valid values:
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
	// SUBMITTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SparkAppInfo) String() string {
	return dara.Prettify(s)
}

func (s SparkAppInfo) GoString() string {
	return s.String()
}

func (s *SparkAppInfo) GetAppId() *string {
	return s.AppId
}

func (s *SparkAppInfo) GetAppName() *string {
	return s.AppName
}

func (s *SparkAppInfo) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *SparkAppInfo) GetDetail() *Detail {
	return s.Detail
}

func (s *SparkAppInfo) GetMessage() *string {
	return s.Message
}

func (s *SparkAppInfo) GetPriority() *string {
	return s.Priority
}

func (s *SparkAppInfo) GetState() *string {
	return s.State
}

func (s *SparkAppInfo) SetAppId(v string) *SparkAppInfo {
	s.AppId = &v
	return s
}

func (s *SparkAppInfo) SetAppName(v string) *SparkAppInfo {
	s.AppName = &v
	return s
}

func (s *SparkAppInfo) SetDBClusterId(v string) *SparkAppInfo {
	s.DBClusterId = &v
	return s
}

func (s *SparkAppInfo) SetDetail(v *Detail) *SparkAppInfo {
	s.Detail = v
	return s
}

func (s *SparkAppInfo) SetMessage(v string) *SparkAppInfo {
	s.Message = &v
	return s
}

func (s *SparkAppInfo) SetPriority(v string) *SparkAppInfo {
	s.Priority = &v
	return s
}

func (s *SparkAppInfo) SetState(v string) *SparkAppInfo {
	s.State = &v
	return s
}

func (s *SparkAppInfo) Validate() error {
	if s.Detail != nil {
		if err := s.Detail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
