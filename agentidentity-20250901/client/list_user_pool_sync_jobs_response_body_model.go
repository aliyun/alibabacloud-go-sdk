// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPoolSyncJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUserPoolSyncJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserPoolSyncJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserPoolSyncJobsResponseBody
	GetRequestId() *string
	SetSynchronizationJobs(v []*ListUserPoolSyncJobsResponseBodySynchronizationJobs) *ListUserPoolSyncJobsResponseBody
	GetSynchronizationJobs() []*ListUserPoolSyncJobsResponseBodySynchronizationJobs
	SetTotalCount(v int32) *ListUserPoolSyncJobsResponseBody
	GetTotalCount() *int32
}

type ListUserPoolSyncJobsResponseBody struct {
	MaxResults          *int32                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobs []*ListUserPoolSyncJobsResponseBodySynchronizationJobs `json:"SynchronizationJobs,omitempty" xml:"SynchronizationJobs,omitempty" type:"Repeated"`
	TotalCount          *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserPoolSyncJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolSyncJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPoolSyncJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserPoolSyncJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserPoolSyncJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPoolSyncJobsResponseBody) GetSynchronizationJobs() []*ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	return s.SynchronizationJobs
}

func (s *ListUserPoolSyncJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserPoolSyncJobsResponseBody) SetMaxResults(v int32) *ListUserPoolSyncJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBody) SetNextToken(v string) *ListUserPoolSyncJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBody) SetRequestId(v string) *ListUserPoolSyncJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBody) SetSynchronizationJobs(v []*ListUserPoolSyncJobsResponseBodySynchronizationJobs) *ListUserPoolSyncJobsResponseBody {
	s.SynchronizationJobs = v
	return s
}

func (s *ListUserPoolSyncJobsResponseBody) SetTotalCount(v int32) *ListUserPoolSyncJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBody) Validate() error {
	if s.SynchronizationJobs != nil {
		for _, item := range s.SynchronizationJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPoolSyncJobsResponseBodySynchronizationJobs struct {
	EndTime              *string                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage         *string                                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IdentityProviderType *string                                                        `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	JobSummary           *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary `json:"JobSummary,omitempty" xml:"JobSummary,omitempty" type:"Struct"`
	StartTime            *string                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	SynchronizationJobId *string                                                        `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	TriggerType          *string                                                        `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListUserPoolSyncJobsResponseBodySynchronizationJobs) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolSyncJobsResponseBodySynchronizationJobs) GoString() string {
	return s.String()
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetEndTime() *string {
	return s.EndTime
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetJobSummary() *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary {
	return s.JobSummary
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetStartTime() *string {
	return s.StartTime
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetStatus() *string {
	return s.Status
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetEndTime(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.EndTime = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetErrorMessage(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetIdentityProviderType(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.IdentityProviderType = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetJobSummary(v *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.JobSummary = v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetStartTime(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.StartTime = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetStatus(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.Status = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetSynchronizationJobId(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.SynchronizationJobId = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) SetTriggerType(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobs {
	s.TriggerType = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobs) Validate() error {
	if s.JobSummary != nil {
		if err := s.JobSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary struct {
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	Deleted *string `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Same    *string `json:"Same,omitempty" xml:"Same,omitempty"`
	Updated *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) String() string {
	return dara.Prettify(s)
}

func (s ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) GoString() string {
	return s.String()
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) GetCreated() *string {
	return s.Created
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) GetDeleted() *string {
	return s.Deleted
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) GetSame() *string {
	return s.Same
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) GetUpdated() *string {
	return s.Updated
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) SetCreated(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary {
	s.Created = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) SetDeleted(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary {
	s.Deleted = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) SetSame(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary {
	s.Same = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) SetUpdated(v string) *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary {
	s.Updated = &v
	return s
}

func (s *ListUserPoolSyncJobsResponseBodySynchronizationJobsJobSummary) Validate() error {
	return dara.Validate(s)
}
