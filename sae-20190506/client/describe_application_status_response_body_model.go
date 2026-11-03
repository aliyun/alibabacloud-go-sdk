// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationStatusResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationStatusResponseBodyData) *DescribeApplicationStatusResponseBody
	GetData() *DescribeApplicationStatusResponseBodyData
	SetErrorCode(v string) *DescribeApplicationStatusResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationStatusResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationStatusResponseBody
	GetTraceId() *string
}

type DescribeApplicationStatusResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DescribeApplicationStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidApplication.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If an exception occurs, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application status information was successfully retrieved. Valid values:
	//
	// - **true**: The information was successfully retrieved.
	//
	// - **false**: The information failed to be retrieved.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationStatusResponseBody) GetData() *DescribeApplicationStatusResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationStatusResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationStatusResponseBody) SetCode(v string) *DescribeApplicationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetData(v *DescribeApplicationStatusResponseBodyData) *DescribeApplicationStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetErrorCode(v string) *DescribeApplicationStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetMessage(v string) *DescribeApplicationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetRequestId(v string) *DescribeApplicationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetSuccess(v bool) *DescribeApplicationStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) SetTraceId(v string) *DescribeApplicationStatusResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationStatusResponseBodyData struct {
	// The ID of the current application.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether Application Real-Time Monitoring Service (ARMS) Pro is activated. Valid values:
	//
	// - **true**: ARMS Pro is activated.
	//
	// - **false**: ARMS Pro is not activated.
	//
	// example:
	//
	// false
	ArmsAdvancedEnabled *string `json:"ArmsAdvancedEnabled,omitempty" xml:"ArmsAdvancedEnabled,omitempty"`
	// The metadata information of the application in ARMS.
	//
	// example:
	//
	// {"appId":"0099b7be-5f5b-4512-a7fc-56049ef1****","licenseKey":"d5cgdt5pu0@7303f55292a****"}
	ArmsApmInfo *string `json:"ArmsApmInfo,omitempty" xml:"ArmsApmInfo,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 1563373372746
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current status of the application. Valid values:
	//
	// - **RUNNING**: The application is running.
	//
	// - **STOPPED**: The application is stopped.
	//
	// - **UNKNOWN**: The application status is unknown.
	//
	// example:
	//
	// RUNNING
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// Specifies whether the Serverless App Engine (SAE) agent is enabled. Valid values:
	//
	// - **true**: The agent is enabled.
	//
	// - **false**: The agent is not enabled.
	//
	// example:
	//
	// false
	EnableAgent *bool `json:"EnableAgent,omitempty" xml:"EnableAgent,omitempty"`
	// The file size limit. Valid values: 0 to 10240. Unit: KB.
	//
	// example:
	//
	// 10240
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" xml:"FileSizeLimit,omitempty"`
	// The ID of the most recently executed change order. If no change order has been executed or the change order information has expired, this field is empty.
	//
	// example:
	//
	// 1ccc2339-fc19-49aa-bda0-1e7b8497****
	LastChangeOrderId *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	// Specifies whether the most recent change order is being executed. Valid values:
	//
	// - **true**: The most recent change order is being executed.
	//
	// - **false**: The most recent change order is not being executed.
	//
	// example:
	//
	// false
	LastChangeOrderRunning *bool `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	// The status of the most recent change order. Valid values:
	//
	// - **READY**: The most recent change order is ready.
	//
	// - **RUNNING**: The most recent change order is being executed.
	//
	// - **SUCCESS**: The most recent change order is successful.
	//
	// - **FAIL**: The most recent change order failed.
	//
	// - **ABORT**: The most recent change order is aborted.
	//
	// - **WAIT_BATCH_CONFIRM**: The most recent change order is waiting for manual batch confirmation.
	//
	// - **AUTO_BATCH_WAIT**: The most recent change order is in the automatic batch waiting state.
	//
	// - **SYSTEM_FAIL**: A system fault occurred.
	//
	// - **WAIT_APPROVAL**: The most recent change order is waiting for approval.
	//
	// - **APPROVED**: The most recent change order is approved and waiting to be executed.
	//
	// example:
	//
	// SUCCESS
	LastChangeOrderStatus *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	// The number of running instances of the current application.
	//
	// example:
	//
	// 1
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// The sub-status of the change order, which is used to determine whether an exception occurred during the release process. Valid values:
	//
	// - **NORMAL**: The release is normal.
	//
	// - **RUNNING_BUT_HAS_ERROR**: An exception occurred during the release. For example, an error occurs during a batch release and requires a manual rollback. In this case, the change order is still being executed because it cannot be completed despite the error.
	//
	// example:
	//
	// NORMAL
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
}

func (s DescribeApplicationStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationStatusResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationStatusResponseBodyData) GetArmsAdvancedEnabled() *string {
	return s.ArmsAdvancedEnabled
}

func (s *DescribeApplicationStatusResponseBodyData) GetArmsApmInfo() *string {
	return s.ArmsApmInfo
}

func (s *DescribeApplicationStatusResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApplicationStatusResponseBodyData) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *DescribeApplicationStatusResponseBodyData) GetEnableAgent() *bool {
	return s.EnableAgent
}

func (s *DescribeApplicationStatusResponseBodyData) GetFileSizeLimit() *int64 {
	return s.FileSizeLimit
}

func (s *DescribeApplicationStatusResponseBodyData) GetLastChangeOrderId() *string {
	return s.LastChangeOrderId
}

func (s *DescribeApplicationStatusResponseBodyData) GetLastChangeOrderRunning() *bool {
	return s.LastChangeOrderRunning
}

func (s *DescribeApplicationStatusResponseBodyData) GetLastChangeOrderStatus() *string {
	return s.LastChangeOrderStatus
}

func (s *DescribeApplicationStatusResponseBodyData) GetRunningInstances() *int32 {
	return s.RunningInstances
}

func (s *DescribeApplicationStatusResponseBodyData) GetSubStatus() *string {
	return s.SubStatus
}

func (s *DescribeApplicationStatusResponseBodyData) SetAppId(v string) *DescribeApplicationStatusResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetArmsAdvancedEnabled(v string) *DescribeApplicationStatusResponseBodyData {
	s.ArmsAdvancedEnabled = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetArmsApmInfo(v string) *DescribeApplicationStatusResponseBodyData {
	s.ArmsApmInfo = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetCreateTime(v string) *DescribeApplicationStatusResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetCurrentStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetEnableAgent(v bool) *DescribeApplicationStatusResponseBodyData {
	s.EnableAgent = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetFileSizeLimit(v int64) *DescribeApplicationStatusResponseBodyData {
	s.FileSizeLimit = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderId(v string) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderId = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderRunning(v bool) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderRunning = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetLastChangeOrderStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.LastChangeOrderStatus = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetRunningInstances(v int32) *DescribeApplicationStatusResponseBodyData {
	s.RunningInstances = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) SetSubStatus(v string) *DescribeApplicationStatusResponseBodyData {
	s.SubStatus = &v
	return s
}

func (s *DescribeApplicationStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
