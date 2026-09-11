// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeFullProcessListResponseBody
	GetCode() *string
	SetConfigList(v map[string]interface{}) *DescribeFullProcessListResponseBody
	GetConfigList() map[string]interface{}
	SetDtsJobId(v string) *DescribeFullProcessListResponseBody
	GetDtsJobId() *string
	SetDynamicMessage(v string) *DescribeFullProcessListResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeFullProcessListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeFullProcessListResponseBody
	GetErrMessage() *string
	SetFullProcessList(v []*DescribeFullProcessListResponseBodyFullProcessList) *DescribeFullProcessListResponseBody
	GetFullProcessList() []*DescribeFullProcessListResponseBodyFullProcessList
	SetHttpStatusCode(v int32) *DescribeFullProcessListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeFullProcessListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeFullProcessListResponseBody
	GetSuccess() *bool
}

type DescribeFullProcessListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter does not return a value. The following parameters describe the rate limiting configurations:
	//
	// - **dts.datamove.blaster.qps.max**: the rate of queries per second to the source database.
	//
	// - **dts.datamove.source.rps.max**: the number of rows per second for full data synchronization or migration (RPS).
	//
	// - **dts.datamove.source.bps.max**: the amount of data per second for full data synchronization or migration, in bytes per second.
	//
	// > - When **JobCode*	- is set to **03**, you must set **EnableLimit*	- to **true*	- for the three parameters to take effect.
	//
	// - When **JobCode*	- is set to **04*	- or **07**, you only need to configure **dts.datamove.source.rps.max*	- and **dts.datamove.source.bps.max**.
	//
	// - A value of **-1*	- indicates that rate limiting is disabled.
	//
	// example:
	//
	// {
	//
	//       "dts.datamove.source.rps.max": 5000,
	//
	//       "dts.datamove.source.bps.max": 10485760
	//
	// }
	ConfigList map[string]interface{} `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The ID of the migration, synchronization, or change tracking task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The dynamic error message used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
	//
	// > For example, if **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The details of the Alibaba Cloud Global Accelerator (GA) instance list.
	FullProcessList []*DescribeFullProcessListResponseBodyFullProcessList `json:"FullProcessList,omitempty" xml:"FullProcessList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFullProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFullProcessListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeFullProcessListResponseBody) GetConfigList() map[string]interface{} {
	return s.ConfigList
}

func (s *DescribeFullProcessListResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeFullProcessListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeFullProcessListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeFullProcessListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeFullProcessListResponseBody) GetFullProcessList() []*DescribeFullProcessListResponseBodyFullProcessList {
	return s.FullProcessList
}

func (s *DescribeFullProcessListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeFullProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFullProcessListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeFullProcessListResponseBody) SetCode(v string) *DescribeFullProcessListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetConfigList(v map[string]interface{}) *DescribeFullProcessListResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetDtsJobId(v string) *DescribeFullProcessListResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetDynamicMessage(v string) *DescribeFullProcessListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetErrCode(v string) *DescribeFullProcessListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetErrMessage(v string) *DescribeFullProcessListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetFullProcessList(v []*DescribeFullProcessListResponseBodyFullProcessList) *DescribeFullProcessListResponseBody {
	s.FullProcessList = v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetHttpStatusCode(v int32) *DescribeFullProcessListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetRequestId(v string) *DescribeFullProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) SetSuccess(v bool) *DescribeFullProcessListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeFullProcessListResponseBody) Validate() error {
	if s.FullProcessList != nil {
		for _, item := range s.FullProcessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFullProcessListResponseBodyFullProcessList struct {
	// The details.
	//
	// example:
	//
	// {}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The exception status of the task. Valid values: - **notstarted**: not started. - **checking**: being checked. - **failed**: failed. - **finished**: completed.
	//
	// example:
	//
	// notstarted
	Exception *string `json:"Exception,omitempty" xml:"Exception,omitempty"`
	// The process name.
	//
	// example:
	//
	// universer
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The process type. Valid values:
	//
	// - **1**: trusted
	//
	// - **2**: suspicious
	//
	// - **3**: malicious.
	//
	// example:
	//
	// 1
	ProcessType *string `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// The SQL statement that is being executed.
	//
	// example:
	//
	// test
	RunningSQL *string `json:"RunningSQL,omitempty" xml:"RunningSQL,omitempty"`
	// The status of the log information.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The task ID.
	//
	// example:
	//
	// TaskD4E5F6
	TaskID *string `json:"TaskID,omitempty" xml:"TaskID,omitempty"`
	// The time when the log was collected, in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// example:
	//
	// 1729650129452
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeFullProcessListResponseBodyFullProcessList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullProcessListResponseBodyFullProcessList) GoString() string {
	return s.String()
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetDetail() *string {
	return s.Detail
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetException() *string {
	return s.Exception
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetProcessType() *string {
	return s.ProcessType
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetRunningSQL() *string {
	return s.RunningSQL
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetState() *string {
	return s.State
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetTaskID() *string {
	return s.TaskID
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetDetail(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.Detail = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetException(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.Exception = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetProcessName(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.ProcessName = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetProcessType(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.ProcessType = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetRunningSQL(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.RunningSQL = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetState(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.State = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetTaskID(v string) *DescribeFullProcessListResponseBodyFullProcessList {
	s.TaskID = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) SetTime(v int64) *DescribeFullProcessListResponseBodyFullProcessList {
	s.Time = &v
	return s
}

func (s *DescribeFullProcessListResponseBodyFullProcessList) Validate() error {
	return dara.Validate(s)
}
