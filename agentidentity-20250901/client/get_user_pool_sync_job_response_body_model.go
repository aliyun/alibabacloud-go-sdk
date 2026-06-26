// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserPoolSyncJobResponseBody
	GetRequestId() *string
	SetSynchronizationJob(v *GetUserPoolSyncJobResponseBodySynchronizationJob) *GetUserPoolSyncJobResponseBody
	GetSynchronizationJob() *GetUserPoolSyncJobResponseBodySynchronizationJob
}

type GetUserPoolSyncJobResponseBody struct {
	RequestId          *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJob *GetUserPoolSyncJobResponseBodySynchronizationJob `json:"SynchronizationJob,omitempty" xml:"SynchronizationJob,omitempty" type:"Struct"`
}

func (s GetUserPoolSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPoolSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPoolSyncJobResponseBody) GetSynchronizationJob() *GetUserPoolSyncJobResponseBodySynchronizationJob {
	return s.SynchronizationJob
}

func (s *GetUserPoolSyncJobResponseBody) SetRequestId(v string) *GetUserPoolSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBody) SetSynchronizationJob(v *GetUserPoolSyncJobResponseBodySynchronizationJob) *GetUserPoolSyncJobResponseBody {
	s.SynchronizationJob = v
	return s
}

func (s *GetUserPoolSyncJobResponseBody) Validate() error {
	if s.SynchronizationJob != nil {
		if err := s.SynchronizationJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPoolSyncJobResponseBodySynchronizationJob struct {
	EndTime              *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage         *string                                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IdentityProviderType *string                                                     `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	JobSummary           *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary `json:"JobSummary,omitempty" xml:"JobSummary,omitempty" type:"Struct"`
	StartTime            *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	SynchronizationJobId *string                                                     `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	TriggerType          *string                                                     `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetUserPoolSyncJobResponseBodySynchronizationJob) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolSyncJobResponseBodySynchronizationJob) GoString() string {
	return s.String()
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetEndTime() *string {
	return s.EndTime
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetJobSummary() *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary {
	return s.JobSummary
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetStartTime() *string {
	return s.StartTime
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetStatus() *string {
	return s.Status
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetEndTime(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.EndTime = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetErrorMessage(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetIdentityProviderType(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.IdentityProviderType = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetJobSummary(v *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.JobSummary = v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetStartTime(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.StartTime = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetStatus(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.Status = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetSynchronizationJobId(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) SetTriggerType(v string) *GetUserPoolSyncJobResponseBodySynchronizationJob {
	s.TriggerType = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJob) Validate() error {
	if s.JobSummary != nil {
		if err := s.JobSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary struct {
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	Deleted *string `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Same    *string `json:"Same,omitempty" xml:"Same,omitempty"`
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) GoString() string {
	return s.String()
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) GetCreated() *string {
	return s.Created
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) GetDeleted() *string {
	return s.Deleted
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) GetSame() *string {
	return s.Same
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) GetUpdated() *string {
	return s.Updated
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) SetCreated(v string) *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary {
	s.Created = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) SetDeleted(v string) *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary {
	s.Deleted = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) SetSame(v string) *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary {
	s.Same = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) SetUpdated(v string) *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary {
	s.Updated = &v
	return s
}

func (s *GetUserPoolSyncJobResponseBodySynchronizationJobJobSummary) Validate() error {
	return dara.Validate(s)
}
