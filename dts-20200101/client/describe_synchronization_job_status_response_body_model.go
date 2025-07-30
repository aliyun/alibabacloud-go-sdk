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
	// The UNIX timestamp generated when the latest data record was synchronized.
	//
	// >  You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1610616144
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// Indicates whether full data synchronization is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataInitialization *string `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// The status of full data synchronization.
	DataInitializationStatus *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data synchronization.
	DataSynchronizationStatus *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The synchronization latency, in seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The synchronization delay, in milliseconds.
	//
	// example:
	//
	// 506
	DelayMillis *int64 `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the data synchronization instance expires. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// >  This parameter is returned only if the return value of the **PayType*	- parameter is **PrePaid**.
	//
	// example:
	//
	// 2021-03-07T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method of the data synchronization instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data synchronization instance.
	Performance *DescribeSynchronizationJobStatusResponseBodyPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck status.
	PrecheckStatus *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DACDF659-AFC6-4DC8-ADB8-4569419A4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeSynchronizationJobStatusResponseBodySourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of the data synchronization task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Initializing**: The task is performing initial synchronization.
	//
	// 	- **InitializeFailed**: Initial synchronization failed.
	//
	// 	- **Synchronizing**: The task is synchronizing data.
	//
	// 	- **Failed**: The task failed to synchronize data.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Modifying**: The objects in the task are being modified.
	//
	// 	- **Finished**: The task is completed.
	//
	// example:
	//
	// synchronizing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether schema synchronization is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	StructureInitialization *string `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The status of schema synchronization.
	StructureInitializationStatus *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The specification of the data synchronization instance.
	//
	// example:
	//
	// large
	SynchronizationJobClass *string `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// dtstest
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	// The objects that are synchronized by the task.
	SynchronizationObjects []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus struct {
	// The error message returned if full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of full data synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been synchronized during full data synchronization.
	//
	// example:
	//
	// 200001
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of full data synchronization. Valid values:
	//
	// 	- **NotStarted**: Full data synchronization is not started.
	//
	// 	- **Migrating**: Full data synchronization is in progress.
	//
	// 	- **Failed**: Full data synchronization failed.
	//
	// 	- **Finished**: Full data synchronization is completed.
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
	// The UNIX timestamp generated when the latest data record was synchronized.
	//
	// example:
	//
	// 1610709865
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
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
	// 856
	DelayMillis *int64 `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	// The error message returned if incremental data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of incremental data synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of incremental data synchronization. Valid values:
	//
	// 	- **NotStarted**: Incremental data synchronization is not started.
	//
	// 	- **Migrating**: Incremental data synchronization is in progress.
	//
	// 	- **Failed**: Incremental data synchronization failed.
	//
	// 	- **Finished**: Incremental data synchronization is completed.
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
	// The ID of the destination instance.
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
	// The data traffic that is synchronized per second. Unit: MB/s.
	//
	// example:
	//
	// 1
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
	// The number of times SQL statements are synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
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
	// The result of each precheck item.
	Detail []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The precheck progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**: The task passed the precheck.
	//
	// 	- **Failed**: The task failed to pass the precheck.
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
	return dara.Validate(s)
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail struct {
	// The precheck result. Valid values:
	//
	// 	- **Success**: The task passed the precheck.
	//
	// 	- **Failed**: The task failed to pass the precheck.
	//
	// example:
	//
	// Success
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The method to fix the precheck failure.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is **Failed**.
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
	// The ID of the source instance.
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
	// The error message returned if schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables whose schemas have been synchronized.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of schema synchronization. Valid values:
	//
	// 	- **NotStarted**: Schema synchronization is not started.
	//
	// 	- **Migrating**: Schema synchronization is in progress.
	//
	// 	- **Failed**: Schema synchronization failed.
	//
	// 	- **Finished**: Schema synchronization is completed.
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
	// The database name that is used in the destination instance.
	//
	// example:
	//
	// newdtstestdatabase
	NewSchemaName *string `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	// The name of the synchronized database.
	//
	// example:
	//
	// dtstestdatabase
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The source tables that are excluded from the data synchronization task.
	TableExcludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	// The tables that are synchronized by the task.
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
	return dara.Validate(s)
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
	// The name of the synchronized table.
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
