// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSynchronizationJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSynchronizationJobsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSynchronizationJobsResponseBody
	GetRequestId() *string
	SetSynchronizationInstances(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstances) *DescribeSynchronizationJobsResponseBody
	GetSynchronizationInstances() []*DescribeSynchronizationJobsResponseBodySynchronizationInstances
	SetTotalRecordCount(v int64) *DescribeSynchronizationJobsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeSynchronizationJobsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of records that can be displayed on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92E1E99D-5224-4AD3-8C94-23A3516B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of data synchronization instances and the details of each instance.
	SynchronizationInstances []*DescribeSynchronizationJobsResponseBodySynchronizationInstances `json:"SynchronizationInstances,omitempty" xml:"SynchronizationInstances,omitempty" type:"Repeated"`
	// The total number of data synchronization instances that meet the specified conditions under the Alibaba Cloud account.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSynchronizationJobsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSynchronizationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationJobsResponseBody) GetSynchronizationInstances() []*DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	return s.SynchronizationInstances
}

func (s *DescribeSynchronizationJobsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeSynchronizationJobsResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetRequestId(v string) *DescribeSynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetSynchronizationInstances(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstances) *DescribeSynchronizationJobsResponseBody {
	s.SynchronizationInstances = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) Validate() error {
	if s.SynchronizationInstances != nil {
		for _, item := range s.SynchronizationInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstances struct {
	// The time when the synchronization task was created, in the format of <i>yyyy-MM-dd HH:mm:ss</i>.0 (UTC+8).
	//
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether initial full data synchronization is performed. Valid values:
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
	DataInitializationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data synchronization.
	//
	// > This parameter set and its contained parameters have been discontinued.
	DataSynchronizationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The synchronization latency, in seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection information of the destination instance.
	DestinationEndpoint *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The error message returned when data synchronization fails.
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
	// 2021-07-07T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the synchronization instance was created, in the format of <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2021-06-28T09:36:32Z
	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	// The time when the synchronization task was created, in the format of <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// example:
	//
	// 2021-06-28T09:34:53Z
	JobCreateTime *string `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
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
	// The overview information of the synchronization link.
	Performance *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck status.
	PrecheckStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The connection information of the source instance.
	SourceEndpoint *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of the synchronization instance. Valid values:
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running the precheck.
	//
	// - **PrecheckFailed**: precheck failed.
	//
	// - **Initializating**: performing initial synchronization.
	//
	// - **InitializeFailed**: initial synchronization failed.
	//
	// - **Synchronizing**: synchronizing.
	//
	// - **Failed**: synchronization failed.
	//
	// - **Suspending**: paused.
	//
	// - **Modifying**: modifying synchronization objects.
	//
	// - **Finished**: completed.
	//
	// example:
	//
	// synchronizing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether initial schema synchronization is performed. Valid values:
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
	StructureInitializationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
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
	SynchronizationObjects []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
	// The tag collection.
	Tags []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetDataInitialization() *string {
	return s.DataInitialization
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetDataInitializationStatus() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetDataSynchronizationStatus() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetDelay() *string {
	return s.Delay
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetDestinationEndpoint() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetInstanceCreateTime() *string {
	return s.InstanceCreateTime
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetJobCreateTime() *string {
	return s.JobCreateTime
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetPayType() *string {
	return s.PayType
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetPerformance() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	return s.Performance
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetPrecheckStatus() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSourceEndpoint() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetStructureInitialization() *string {
	return s.StructureInitialization
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetStructureInitializationStatus() *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSynchronizationJobClass() *string {
	return s.SynchronizationJobClass
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSynchronizationJobName() *string {
	return s.SynchronizationJobName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetSynchronizationObjects() []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	return s.SynchronizationObjects
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) GetTags() []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags {
	return s.Tags
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataSynchronizationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDestinationEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetExpireTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetInstanceCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetJobCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPayType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPerformance(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPrecheckStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSourceEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationDirection(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationObjects(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationObjects = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetTags(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Tags = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) Validate() error {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus struct {
	// The error message returned when initial full data synchronization fails.
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

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus struct {
	// The synchronization latency of incremental data synchronization.
	//
	// > This parameter has been discontinued.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The error message returned when incremental data synchronization fails.
	//
	// > This parameter has been discontinued.
	//
	// example:
	//
	// 任务失败太久无法恢复
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of incremental data synchronization.
	//
	// > This parameter has been discontinued.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of incremental data synchronization.
	//
	// > This parameter has been discontinued.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GetDelay() *string {
	return s.Delay
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint struct {
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

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetPort(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetUserName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance struct {
	// The volume of data synchronized per second, in MB/s.
	//
	// example:
	//
	// 1
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
	// The number of SQL statements synchronized per second, including BEGIN, COMMIT, DML statements (INSERT, DELETE, and UPDATE), and DDL statements.
	//
	// example:
	//
	// 100
	RPS *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) GetFLOW() *string {
	return s.FLOW
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) GetRPS() *string {
	return s.RPS
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetFLOW(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.FLOW = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetRPS(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.RPS = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus struct {
	// The execution details of each precheck item.
	Detail []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
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

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GetDetail() []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) Validate() error {
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

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail struct {
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
	// The error message returned when the precheck fails.
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
	// The repair method when the precheck fails.
	//
	// > This parameter is returned only when the value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint struct {
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

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetPort(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetUserName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus struct {
	// The error message returned when initial schema synchronization fails.
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
	// The number of tables for which initial schema synchronization has been completed.
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

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects struct {
	// The name of the database to which the objects are mapped in the destination database.
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
	// The tables that are excluded from the database to be synchronized. These tables are not synchronized.
	TableExcludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	// The tables to be synchronized.
	TableIncludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GetNewSchemaName() *string {
	return s.NewSchemaName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GetTableExcludes() []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes {
	return s.TableExcludes
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GetTableIncludes() []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes {
	return s.TableIncludes
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetNewSchemaName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.NewSchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.SchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableIncludes = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) Validate() error {
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

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes struct {
	// The name of the excluded table.
	//
	// example:
	//
	// order
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) SetTableName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes {
	s.TableName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes struct {
	// The name of the table to be synchronized.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) GetTableName() *string {
	return s.TableName
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) SetTableName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes {
	s.TableName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) Validate() error {
	return dara.Validate(s)
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) SetKey(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) SetValue(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags {
	s.Value = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) Validate() error {
	return dara.Validate(s)
}
