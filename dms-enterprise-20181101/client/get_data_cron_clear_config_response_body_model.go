// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCronClearConfig(v *GetDataCronClearConfigResponseBodyDataCronClearConfig) *GetDataCronClearConfigResponseBody
	GetDataCronClearConfig() *GetDataCronClearConfigResponseBodyDataCronClearConfig
	SetErrorCode(v string) *GetDataCronClearConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCronClearConfigResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataCronClearConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCronClearConfigResponseBody
	GetSuccess() *bool
}

type GetDataCronClearConfigResponseBody struct {
	// Data configuration.
	DataCronClearConfig *GetDataCronClearConfigResponseBodyDataCronClearConfig `json:"DataCronClearConfig,omitempty" xml:"DataCronClearConfig,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataCronClearConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCronClearConfigResponseBody) GetDataCronClearConfig() *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	return s.DataCronClearConfig
}

func (s *GetDataCronClearConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCronClearConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCronClearConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCronClearConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCronClearConfigResponseBody) SetDataCronClearConfig(v *GetDataCronClearConfigResponseBodyDataCronClearConfig) *GetDataCronClearConfigResponseBody {
	s.DataCronClearConfig = v
	return s
}

func (s *GetDataCronClearConfigResponseBody) SetErrorCode(v string) *GetDataCronClearConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCronClearConfigResponseBody) SetErrorMessage(v string) *GetDataCronClearConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCronClearConfigResponseBody) SetRequestId(v string) *GetDataCronClearConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCronClearConfigResponseBody) SetSuccess(v bool) *GetDataCronClearConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCronClearConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataCronClearConfigResponseBodyDataCronClearConfig struct {
	// The number of times that the task is run.
	//
	// example:
	//
	// 2
	CronCallTimes *string `json:"CronCallTimes,omitempty" xml:"CronCallTimes,omitempty"`
	// The crontab expression that you can use to run the task at a specified time. For more information, see [Crontab expression](https://help.aliyun.com/document_detail/206581.html).
	//
	// example:
	//
	// 0 0 23 ? 	- 7,1
	CronFormat *string `json:"CronFormat,omitempty" xml:"CronFormat,omitempty"`
	// The time when the task was last run.
	//
	// example:
	//
	// 2022-11-18 00:00:00
	CronLastCallStartTime *string `json:"CronLastCallStartTime,omitempty" xml:"CronLastCallStartTime,omitempty"`
	// The time when the task is run next time. This parameter is displayed only when the status of the scheduled task is SUCCESS.
	//
	// example:
	//
	// 2022-11-19 00:00:00
	CronNextCallTime *string `json:"CronNextCallTime,omitempty" xml:"CronNextCallTime,omitempty"`
	// The status of the scheduled task. If this parameter is empty, it indicates the task is not run. Valid values:
	//
	// 	- PAUSE: The task is suspended.
	//
	// 	- WAITING: The task is waiting to be run.
	//
	// 	- SUCCESS: The task is complete.
	//
	// example:
	//
	// PAUSE
	CronStatus *string `json:"CronStatus,omitempty" xml:"CronStatus,omitempty"`
	// The number of times that the Optimize Table statement is automatically exeuted. This parameter is valid only when the value of the OptimizeTableAfterEveryClearTimes parameter is greater than 0.
	//
	// example:
	//
	// 0
	CurrentClearTaskCount *int64 `json:"CurrentClearTaskCount,omitempty" xml:"CurrentClearTaskCount,omitempty"`
	// The execution duration of the task. Unit: hours. If the value is 0, it indicates the duration is not specified.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Specifies whether to enable automatic execution of the OPTIMIZE TABLE statement. Valid values:
	//
	// 	- 0: disables automatic execution
	//
	// 	- A number greater than 0: enables automatic execution. The number specifies the number of times that cleanup operations must be performed before the OPTIMIZE TABLE statement is automatically executed.
	//
	// example:
	//
	// 0
	OptimizeTableAfterEveryClearTimes *int64 `json:"OptimizeTableAfterEveryClearTimes,omitempty" xml:"OptimizeTableAfterEveryClearTimes,omitempty"`
}

func (s GetDataCronClearConfigResponseBodyDataCronClearConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearConfigResponseBodyDataCronClearConfig) GoString() string {
	return s.String()
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCronCallTimes() *string {
	return s.CronCallTimes
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCronFormat() *string {
	return s.CronFormat
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCronLastCallStartTime() *string {
	return s.CronLastCallStartTime
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCronNextCallTime() *string {
	return s.CronNextCallTime
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCronStatus() *string {
	return s.CronStatus
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetCurrentClearTaskCount() *int64 {
	return s.CurrentClearTaskCount
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetDuration() *string {
	return s.Duration
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) GetOptimizeTableAfterEveryClearTimes() *int64 {
	return s.OptimizeTableAfterEveryClearTimes
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCronCallTimes(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CronCallTimes = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCronFormat(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CronFormat = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCronLastCallStartTime(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CronLastCallStartTime = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCronNextCallTime(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CronNextCallTime = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCronStatus(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CronStatus = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetCurrentClearTaskCount(v int64) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.CurrentClearTaskCount = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetDuration(v string) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.Duration = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) SetOptimizeTableAfterEveryClearTimes(v int64) *GetDataCronClearConfigResponseBodyDataCronClearConfig {
	s.OptimizeTableAfterEveryClearTimes = &v
	return s
}

func (s *GetDataCronClearConfigResponseBodyDataCronClearConfig) Validate() error {
	return dara.Validate(s)
}
