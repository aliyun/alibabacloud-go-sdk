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
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *DescribeApplicationStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the status of the application was queried. Valid values:
	//
	// 	- **true**: The status was queried.
	//
	// 	- **false**: The status failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
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
	// The ID of the application.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Indicates whether Application Real-Time Monitoring Service (ARMS) advanced monitoring is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ArmsAdvancedEnabled *string `json:"ArmsAdvancedEnabled,omitempty" xml:"ArmsAdvancedEnabled,omitempty"`
	// The metadata of the application in ARMS.
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
	// 	- **RUNNING**
	//
	// 	- **STOPPED**
	//
	// 	- **UNKNOWN**
	//
	// example:
	//
	// RUNNING
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// Indicates whether SAE agent is enabled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableAgent *bool `json:"EnableAgent,omitempty" xml:"EnableAgent,omitempty"`
	// The file size limit. Unit: KB. Valid values: 0 to 10240.
	//
	// example:
	//
	// 10240
	FileSizeLimit *int64 `json:"FileSizeLimit,omitempty" xml:"FileSizeLimit,omitempty"`
	// The ID of the latest change order that is executed. If no change orders are executed or if change orders expire, this parameter is left empty.
	//
	// example:
	//
	// 1ccc2339-fc19-49aa-bda0-1e7b8497****
	LastChangeOrderId *string `json:"LastChangeOrderId,omitempty" xml:"LastChangeOrderId,omitempty"`
	// Indicates whether the latest change order is being executed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	LastChangeOrderRunning *bool `json:"LastChangeOrderRunning,omitempty" xml:"LastChangeOrderRunning,omitempty"`
	// The status of the latest change order. Valid values:
	//
	// 	- **READY**: The change order is ready.
	//
	// 	- **RUNNING**: The change order is being executed.
	//
	// 	- **SUCCESS**: The change order was executed.
	//
	// 	- **FAIL**: The change order failed to be executed.
	//
	// 	- **ABORT**: The change order is stopped.
	//
	// 	- **WAIT_BATCH_CONFIRM**: The change order is pending execution. You must manually confirm the release batch.
	//
	// 	- **AUTO_BATCH_WAIT**: The change order is pending execution. SAE will automatically confirm the release batch.
	//
	// 	- **SYSTEM_FAIL**: A system exception occurred.
	//
	// 	- **WAIT_APPROVAL**: The change order is pending approval.
	//
	// 	- **APPROVED**: The change order is approved and is pending execution.
	//
	// example:
	//
	// SUCCESS
	LastChangeOrderStatus *string `json:"LastChangeOrderStatus,omitempty" xml:"LastChangeOrderStatus,omitempty"`
	// The number of running instances of the application.
	//
	// example:
	//
	// 1
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// The substatus of the change order. This parameter indicates whether an exception occurred while the change order was being executed. Valid values:
	//
	// 	- **NORMAL**
	//
	// 	- **RUNNING_BUT_HAS_ERROR**: For example, if an error occurs during a phased release, you must manually roll back the application. In this case, the change order cannot be completed because the change order is still being executed.
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
