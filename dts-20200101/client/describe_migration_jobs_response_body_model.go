// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeMigrationJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMigrationJobsResponseBody
	GetErrMessage() *string
	SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody
	GetMigrationJobs() *DescribeMigrationJobsResponseBodyMigrationJobs
	SetPageNumber(v int32) *DescribeMigrationJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeMigrationJobsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeMigrationJobsResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeMigrationJobsResponseBody struct {
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
	ErrMessage    *string                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	MigrationJobs *DescribeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries that can be displayed on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0ED4846D-ED60-435D-88C0-7EC0CE4D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data migration instances that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// 300
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMigrationJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMigrationJobsResponseBody) GetMigrationJobs() *DescribeMigrationJobsResponseBodyMigrationJobs {
	return s.MigrationJobs
}

func (s *DescribeMigrationJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMigrationJobsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeMigrationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrationJobsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMigrationJobsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeMigrationJobsResponseBody) SetErrCode(v string) *DescribeMigrationJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetErrMessage(v string) *DescribeMigrationJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody {
	s.MigrationJobs = v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageNumber(v int32) *DescribeMigrationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetRequestId(v string) *DescribeMigrationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetSuccess(v string) *DescribeMigrationJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) Validate() error {
	if s.MigrationJobs != nil {
		if err := s.MigrationJobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobs struct {
	MigrationJob []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob `json:"MigrationJob,omitempty" xml:"MigrationJob,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) GetMigrationJob() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	return s.MigrationJob
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) SetMigrationJob(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) *DescribeMigrationJobsResponseBodyMigrationJobs {
	s.MigrationJob = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) Validate() error {
	if s.MigrationJob != nil {
		for _, item := range s.MigrationJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob struct {
	DataInitialization      *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization      `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty" type:"Struct"`
	DataSynchronization     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization     `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	DestinationEndpoint     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint     `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	InstanceCreateTime      *string                                                                            `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	JobCreateTime           *string                                                                            `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	MigrationJobClass       *string                                                                            `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	MigrationJobID          *string                                                                            `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
	MigrationJobName        *string                                                                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus      *string                                                                            `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	MigrationMode           *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode           `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	MigrationObject         *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject         `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	PayType                 *string                                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Precheck                *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck                `json:"Precheck,omitempty" xml:"Precheck,omitempty" type:"Struct"`
	SourceEndpoint          *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint          `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitialization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty" type:"Struct"`
	Tags                    *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDataInitialization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	return s.DataInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDataSynchronization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDestinationEndpoint() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetInstanceCreateTime() *string {
	return s.InstanceCreateTime
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetJobCreateTime() *string {
	return s.JobCreateTime
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobClass() *string {
	return s.MigrationJobClass
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobID() *string {
	return s.MigrationJobID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobStatus() *string {
	return s.MigrationJobStatus
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationMode() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationObject() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	return s.MigrationObject
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetPrecheck() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	return s.Precheck
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetSourceEndpoint() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetStructureInitialization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetTags() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags {
	return s.Tags
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataSynchronization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataSynchronization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDestinationEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetInstanceCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetJobCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobClass(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPrecheck(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Precheck = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetSourceEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetStructureInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.StructureInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetTags(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Tags = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) Validate() error {
	if s.DataInitialization != nil {
		if err := s.DataInitialization.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronization != nil {
		if err := s.DataSynchronization.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationMode != nil {
		if err := s.MigrationMode.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationObject != nil {
		if err := s.MigrationObject.Validate(); err != nil {
			return err
		}
	}
	if s.Precheck != nil {
		if err := s.Precheck.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitialization != nil {
		if err := s.StructureInitialization.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetDelay() *string {
	return s.Delay
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetDelay(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Delay = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject struct {
	SynchronousObject []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GetSynchronousObject() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	return s.SynchronousObject
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) SetSynchronousObject(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	s.SynchronousObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) Validate() error {
	if s.SynchronousObject != nil {
		for _, item := range s.SynchronousObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject struct {
	DatabaseName  *string                                                                                              `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                              `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetTableList() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	return s.TableList
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetWholeDatabase() *string {
	return s.WholeDatabase
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetTableList(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) Validate() error {
	if s.TableList != nil {
		if err := s.TableList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GetTable() []*string {
	return s.Table
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck struct {
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags struct {
	Tag []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) GetTag() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	return s.Tag
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) SetTag(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags {
	s.Tag = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetKey(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetValue(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) Validate() error {
	return dara.Validate(s)
}
