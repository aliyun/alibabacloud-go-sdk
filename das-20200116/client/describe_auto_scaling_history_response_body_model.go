// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAutoScalingHistoryResponseBody
	GetCode() *string
	SetData(v *DescribeAutoScalingHistoryResponseBodyData) *DescribeAutoScalingHistoryResponseBody
	GetData() *DescribeAutoScalingHistoryResponseBodyData
	SetMessage(v string) *DescribeAutoScalingHistoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAutoScalingHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeAutoScalingHistoryResponseBody
	GetSuccess() *string
}

type DescribeAutoScalingHistoryResponseBody struct {
	// The HTTP status code returned. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The history of auto scaling.
	Data *DescribeAutoScalingHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request was successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutoScalingHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAutoScalingHistoryResponseBody) GetData() *DescribeAutoScalingHistoryResponseBodyData {
	return s.Data
}

func (s *DescribeAutoScalingHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutoScalingHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoScalingHistoryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeAutoScalingHistoryResponseBody) SetCode(v string) *DescribeAutoScalingHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBody) SetData(v *DescribeAutoScalingHistoryResponseBodyData) *DescribeAutoScalingHistoryResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBody) SetMessage(v string) *DescribeAutoScalingHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBody) SetRequestId(v string) *DescribeAutoScalingHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBody) SetSuccess(v string) *DescribeAutoScalingHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoScalingHistoryResponseBodyData struct {
	// The history of automatic bandwidth scaling of ApsaraDB for Redis instances. This feature is not supported.
	Bandwidth []map[string]interface{} `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The history of resource scale-out of ApsaraDB for Redis instances. This feature is not supported.
	Resource []map[string]interface{} `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	// The history of automatic shard scale-out of ApsaraDB for Redis instances. This feature is not supported.
	Shard []map[string]interface{} `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Repeated"`
	// The history of automatic performance scaling.
	SpecHistory []*DescribeAutoScalingHistoryResponseBodyDataSpecHistory `json:"SpecHistory,omitempty" xml:"SpecHistory,omitempty" type:"Repeated"`
	// The history of storage expansion. This feature is not supported.
	Storage []map[string]interface{} `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Repeated"`
}

func (s DescribeAutoScalingHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetBandwidth() []map[string]interface{} {
	return s.Bandwidth
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetResource() []map[string]interface{} {
	return s.Resource
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetShard() []map[string]interface{} {
	return s.Shard
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetSpecHistory() []*DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	return s.SpecHistory
}

func (s *DescribeAutoScalingHistoryResponseBodyData) GetStorage() []map[string]interface{} {
	return s.Storage
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetBandwidth(v []map[string]interface{}) *DescribeAutoScalingHistoryResponseBodyData {
	s.Bandwidth = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetInstanceId(v string) *DescribeAutoScalingHistoryResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetResource(v []map[string]interface{}) *DescribeAutoScalingHistoryResponseBodyData {
	s.Resource = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetShard(v []map[string]interface{}) *DescribeAutoScalingHistoryResponseBodyData {
	s.Shard = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetSpecHistory(v []*DescribeAutoScalingHistoryResponseBodyDataSpecHistory) *DescribeAutoScalingHistoryResponseBodyData {
	s.SpecHistory = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) SetStorage(v []map[string]interface{}) *DescribeAutoScalingHistoryResponseBodyData {
	s.Storage = v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyData) Validate() error {
	if s.SpecHistory != nil {
		for _, item := range s.SpecHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoScalingHistoryResponseBodyDataSpecHistory struct {
	// The error code returned by the scaling task. Valid values:
	//
	// 	- **Insufficient_Balance**: The account has insufficient balance or an unpaid order.
	//
	// 	- **REACH_SPEC_UPPERBOUND**: The instance type reaches the upper limit.
	//
	// 	- **Control_Error_Timeout_Msg**: The management task timed out.
	//
	// 	- **Invoke_Rds_Api_Error_Msg**: Failed to call the ApsaraDB RDS API.
	//
	// example:
	//
	// Insufficient_Balance
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The original number of CPU cores of the instance.
	//
	// example:
	//
	// 4
	OriginCore *int32 `json:"OriginCore,omitempty" xml:"OriginCore,omitempty"`
	// The original instance type.
	//
	// example:
	//
	// mysql.n2.large.2c
	OriginInstanceClass *string `json:"OriginInstanceClass,omitempty" xml:"OriginInstanceClass,omitempty"`
	// The original memory size of the instance. Unit: GB.
	//
	// example:
	//
	// 8
	OriginMemory *float64 `json:"OriginMemory,omitempty" xml:"OriginMemory,omitempty"`
	// The type of the automatic performance scaling task. Valid values:
	//
	// 	- **SCALE_UP**: automatic instance type scale-up task.
	//
	// 	- **SCALE_DOWN**: automatic instance type scale-down task.
	//
	// example:
	//
	// SCALE_UP
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// The destination number of CPU cores of the instance.
	//
	// example:
	//
	// 8
	TargetCore *int32 `json:"TargetCore,omitempty" xml:"TargetCore,omitempty"`
	// The destination instance type.
	//
	// example:
	//
	// mysql.n2.xlarge.2c
	TargetInstanceClass *string `json:"TargetInstanceClass,omitempty" xml:"TargetInstanceClass,omitempty"`
	// The destination memory size of the instance. Unit: GB.
	//
	// example:
	//
	// 16
	TargetMemory *float64 `json:"TargetMemory,omitempty" xml:"TargetMemory,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **true**: The task was successful.
	//
	// 	- **false**: The task failed.
	//
	// example:
	//
	// true
	TaskExcuteStatus *bool `json:"TaskExcuteStatus,omitempty" xml:"TaskExcuteStatus,omitempty"`
	// The time when the task was run. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1684830763000
	TaskTime *int64 `json:"TaskTime,omitempty" xml:"TaskTime,omitempty"`
}

func (s DescribeAutoScalingHistoryResponseBodyDataSpecHistory) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetOriginCore() *int32 {
	return s.OriginCore
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetOriginInstanceClass() *string {
	return s.OriginInstanceClass
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetOriginMemory() *float64 {
	return s.OriginMemory
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetScaleType() *string {
	return s.ScaleType
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetTargetCore() *int32 {
	return s.TargetCore
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetTargetInstanceClass() *string {
	return s.TargetInstanceClass
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetTargetMemory() *float64 {
	return s.TargetMemory
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetTaskExcuteStatus() *bool {
	return s.TaskExcuteStatus
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) GetTaskTime() *int64 {
	return s.TaskTime
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetErrorCode(v string) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetOriginCore(v int32) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.OriginCore = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetOriginInstanceClass(v string) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.OriginInstanceClass = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetOriginMemory(v float64) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.OriginMemory = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetScaleType(v string) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.ScaleType = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetTargetCore(v int32) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.TargetCore = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetTargetInstanceClass(v string) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.TargetInstanceClass = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetTargetMemory(v float64) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.TargetMemory = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetTaskExcuteStatus(v bool) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.TaskExcuteStatus = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) SetTaskTime(v int64) *DescribeAutoScalingHistoryResponseBodyDataSpecHistory {
	s.TaskTime = &v
	return s
}

func (s *DescribeAutoScalingHistoryResponseBodyDataSpecHistory) Validate() error {
	return dara.Validate(s)
}
