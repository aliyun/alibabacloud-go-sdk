// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBody
	GetCheckpoint() *string
	SetDataInitialization(v string) *DescribeSynchronizationJobStatusResponseBody
	GetDataInitialization() *string
	SetDataInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationJobStatusResponseBody
	GetDataInitializationStatus() *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus
	SetDataSynchronizationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationJobStatusResponseBody
	GetDataSynchronizationStatus() *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus
	SetDelay(v string) *DescribeSynchronizationJobStatusResponseBody
	GetDelay() *string
	SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBody
	GetDelayMillis() *int64
	SetDestinationEndpoint(v *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) *DescribeSynchronizationJobStatusResponseBody
	GetDestinationEndpoint() *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint
	SetErrCode(v string) *DescribeSynchronizationJobStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSynchronizationJobStatusResponseBody
	GetErrMessage() *string
	SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBody
	GetErrorMessage() *string
	SetExpireTime(v string) *DescribeSynchronizationJobStatusResponseBody
	GetExpireTime() *string
	SetPayType(v string) *DescribeSynchronizationJobStatusResponseBody
	GetPayType() *string
	SetPerformance(v *DescribeSynchronizationJobStatusResponseBodyPerformance) *DescribeSynchronizationJobStatusResponseBody
	GetPerformance() *DescribeSynchronizationJobStatusResponseBodyPerformance
	SetPrecheckStatus(v *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) *DescribeSynchronizationJobStatusResponseBody
	GetPrecheckStatus() *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus
	SetRequestId(v string) *DescribeSynchronizationJobStatusResponseBody
	GetRequestId() *string
	SetSourceEndpoint(v *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) *DescribeSynchronizationJobStatusResponseBody
	GetSourceEndpoint() *DescribeSynchronizationJobStatusResponseBodySourceEndpoint
	SetStatus(v string) *DescribeSynchronizationJobStatusResponseBody
	GetStatus() *string
	SetStructureInitialization(v string) *DescribeSynchronizationJobStatusResponseBody
	GetStructureInitialization() *string
	SetStructureInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationJobStatusResponseBody
	GetStructureInitializationStatus() *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus
	SetSuccess(v string) *DescribeSynchronizationJobStatusResponseBody
	GetSuccess() *string
	SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusResponseBody
	GetSynchronizationDirection() *string
	SetSynchronizationJobClass(v string) *DescribeSynchronizationJobStatusResponseBody
	GetSynchronizationJobClass() *string
	SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusResponseBody
	GetSynchronizationJobId() *string
	SetSynchronizationJobName(v string) *DescribeSynchronizationJobStatusResponseBody
	GetSynchronizationJobName() *string
	SetSynchronizationObjects(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) *DescribeSynchronizationJobStatusResponseBody
	GetSynchronizationObjects() []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects
	SetTaskId(v string) *DescribeSynchronizationJobStatusResponseBody
	GetTaskId() *string
}

type DescribeSynchronizationJobStatusResponseBody struct {
	// The timestamp of the latest synchronized data, in UNIX timestamp format.
	//
	// > You can use a search engine to find a UNIX timestamp converter.
	//
	// example:
	//
	// 1610616144
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// Indicates whether initial full data synchronization was performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DataInitialization *string `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// The status of initial full data synchronization.
	DataInitializationStatus *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data synchronization.
	DataSynchronizationStatus *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The synchronization latency, in seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The synchronization latency, in milliseconds.
	//
	// example:
	//
	// 506
	DelayMillis *int64 `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	// The connection information of the destination instance.
	DestinationEndpoint *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The error code returned when the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned when data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The expiration time of the synchronization instance, in the format of <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// > This parameter is returned only when the value of the **PayType*	- parameter is **PrePaid**.
	//
	// example:
	//
	// 2021-03-07T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method of the synchronization instance. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The overview of the synchronization link.
	Performance *DescribeSynchronizationJobStatusResponseBodyPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck status.
	PrecheckStatus *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DACDF659-AFC6-4DC8-ADB8-4569419A4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The connection information of the source instance.
	SourceEndpoint *DescribeSynchronizationJobStatusResponseBodySourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of the synchronization instance. Valid values:
	//
	// - **notStarted**: not started.
	//
	// - **prechecking**: running a precheck.
	//
	// - **precheckFailed**: precheck failed.
	//
	// - **initializating**: performing initial synchronization.
	//
	// - **initializeFailed**: initial synchronization failed.
	//
	// - **synchronizing**: synchronizing.
	//
	// - **failed**: synchronization failed.
	//
	// - **suspending**: paused.
	//
	// - **modifying**: modifying synchronization objects.
	//
	// - **finished**: completed.
	//
	// example:
	//
	// synchronizing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether initial schema synchronization was performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	StructureInitialization *string `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The status of initial schema synchronization.
	StructureInitializationStatus *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The specification of the synchronization link.
	//
	// example:
	//
	// large
	SynchronizationJobClass *string `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	// The instance ID of the data synchronization instance.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// The name of the synchronization instance.
	//
	// example:
	//
	// MySQL同步
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The synchronization objects.
	SynchronizationObjects []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
	// The ID of the data synchronization task.
	//
	// example:
	//
	// exjk1alb116****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDataInitialization() *string {
	return s.DataInitialization
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDataInitializationStatus() *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDataSynchronizationStatus() *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDelay() *string {
	return s.Delay
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDelayMillis() *int64 {
	return s.DelayMillis
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetDestinationEndpoint() *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetPerformance() *DescribeSynchronizationJobStatusResponseBodyPerformance {
	return s.Performance
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetPrecheckStatus() *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSourceEndpoint() *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetStructureInitialization() *string {
	return s.StructureInitialization
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetStructureInitializationStatus() *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSynchronizationJobClass() *string {
	return s.SynchronizationJobClass
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSynchronizationJobName() *string {
	return s.SynchronizationJobName
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetSynchronizationObjects() []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	return s.SynchronizationObjects
}

func (s *DescribeSynchronizationJobStatusResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBody {
	s.DelayMillis = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetExpireTime(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPayType(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPerformance(v *DescribeSynchronizationJobStatusResponseBodyPerformance) *DescribeSynchronizationJobStatusResponseBody {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSourceEndpoint(v *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationObjects(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationObjects = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetTaskId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) Validate() error {
	if s.DataInitializationStatus != nil {
		if err := s.DataInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronizationStatus != nil {
		if err := s.DataSynchronizationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.Performance != nil {
		if err := s.Performance.Validate(); err != nil {
			return err
		}
	}
	if s.PrecheckStatus != nil {
		if err := s.PrecheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitializationStatus != nil {
		if err := s.StructureInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SynchronizationObjects != nil {
		for _, item := range s.SynchronizationObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus struct {
	// The error message returned when initial full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial full data synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been synchronized during initial full data synchronization.
	//
	// example:
	//
	// 200001
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of initial full data synchronization. Valid values:
	//
	// - **NotStarted**: not started.
	//
	// - **Migrating**: in progress.
	//
	// - **Failed**: failed.
	//
	// - **Finished**: completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus struct {
	// The timestamp of the latest synchronized data, in UNIX timestamp format.
	//
	// example:
	//
	// 1610709865
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The synchronization latency of incremental data synchronization, in seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The synchronization latency of incremental data synchronization, in milliseconds.
	//
	// example:
	//
	// 856
	DelayMillis *int64 `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	// The error message returned when incremental data synchronization failed.
	//
	// example:
	//
	// 任务失败太久无法恢复
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of incremental data synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of incremental data synchronization. Valid values:
	//
	// - **NotStarted**: not started.
	//
	// - **Migrating**: synchronizing.
	//
	// - **Failed**: failed.
	//
	// - **Finished**: completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetDelay() *string {
	return s.Delay
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetDelayMillis() *int64 {
	return s.DelayMillis
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.DelayMillis = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint struct {
	// The database type of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The instance ID of the destination instance.
	//
	// example:
	//
	// rm-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetPort(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyPerformance struct {
	// The data flow rate of synchronization per second, in MB/s.
	//
	// example:
	//
	// 1
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
	// The number of SQL statements synchronized per second, including BEGIN, COMMIT, DML statements (INSERT, DELETE, UPDATE), and DDL statements.
	//
	// example:
	//
	// 100
	RPS *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) GetFLOW() *string {
	return s.FLOW
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) GetRPS() *string {
	return s.RPS
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetFLOW(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.FLOW = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetRPS(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.RPS = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatus struct {
	// The details of each precheck item.
	Detail []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The overall progress of the precheck, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck result. Valid values:
	//
	// - **Success**: passed.
	//
	// - **Failed**: failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GetDetail() []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail struct {
	// The check result. Valid values:
	//
	// - **Success**: passed.
	//
	// - **Failed**: failed.
	//
	// example:
	//
	// Success
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The error message returned when the precheck failed.
	//
	// > This parameter is returned only when the value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The repair method when the precheck failed.
	//
	// > This parameter is returned only when the value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodySourceEndpoint struct {
	// The database type of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The instance ID of the source instance.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetPort(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetUserName(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus struct {
	// The error message returned when initial schema synchronization encountered an exception.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial schema synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have completed initial schema synchronization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of initial schema synchronization. Valid values:
	//
	// - **NotStarted**: not started.
	//
	// - **Migrating**: in progress.
	//
	// - **Failed**: failed.
	//
	// - **Finished**: completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjects struct {
	// The name mapped to the database to be synchronized in the destination database.
	//
	// example:
	//
	// newdtstestdatabase
	NewSchemaName *string `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	// The name of the database to be synchronized.
	//
	// example:
	//
	// dtstestdatabase
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The tables excluded from the database to be synchronized. These tables will not be synchronized.
	TableExcludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	// The tables to be synchronized.
	TableIncludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GetNewSchemaName() *string {
	return s.NewSchemaName
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GetTableExcludes() []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes {
	return s.TableExcludes
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GetTableIncludes() []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes {
	return s.TableIncludes
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetNewSchemaName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.NewSchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.SchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableIncludes = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) Validate() error {
	if s.TableExcludes != nil {
		for _, item := range s.TableExcludes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TableIncludes != nil {
		for _, item := range s.TableIncludes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes struct {
	// The name of the excluded table.
	//
	// example:
	//
	// order
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) SetTableName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes {
	s.TableName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes struct {
	// The name of the table to be synchronized.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) SetTableName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes {
	s.TableName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) Validate() error {
	return dara.Validate(s)
}
