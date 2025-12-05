// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentCount(v int32) *GetJMeterLogsResponseBody
	GetAgentCount() *int32
	SetCode(v string) *GetJMeterLogsResponseBody
	GetCode() *string
	SetLogs(v []map[string]interface{}) *GetJMeterLogsResponseBody
	GetLogs() []map[string]interface{}
	SetMessage(v string) *GetJMeterLogsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetJMeterLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetJMeterLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetJMeterLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJMeterLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetJMeterLogsResponseBody
	GetTotalCount() *int64
}

type GetJMeterLogsResponseBody struct {
	// The number of engines. The AgentCount value must be greater than the PageNumber value.
	//
	// example:
	//
	// 3
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The system status code. If the request was successful, this parameter is left empty.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned entries.
	//
	// example:
	//
	// { "timeTS":1637114804326, "instanceId":0, "level":"INFO", "logger":"org.apache.jmeter.util.JMeterUtils", "sceneId":251546, 	"planId":1501546, "thread":"main", "time":"2021-11-17T10:06Z", "taskId":15015460000, "logText":"Setting Locale to en_EN\\n" }
	Logs []map[string]interface{} `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetJMeterLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsResponseBody) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetJMeterLogsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJMeterLogsResponseBody) GetLogs() []map[string]interface{} {
	return s.Logs
}

func (s *GetJMeterLogsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJMeterLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetJMeterLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetJMeterLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJMeterLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterLogsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetJMeterLogsResponseBody) SetAgentCount(v int32) *GetJMeterLogsResponseBody {
	s.AgentCount = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetCode(v string) *GetJMeterLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetLogs(v []map[string]interface{}) *GetJMeterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetJMeterLogsResponseBody) SetMessage(v string) *GetJMeterLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetPageNumber(v int32) *GetJMeterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetPageSize(v int32) *GetJMeterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetRequestId(v string) *GetJMeterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetSuccess(v bool) *GetJMeterLogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetTotalCount(v int64) *GetJMeterLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetJMeterLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
