// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitSnapshotJobResponseBody
	GetRequestId() *string
	SetSnapshotJob(v *SubmitSnapshotJobResponseBodySnapshotJob) *SubmitSnapshotJobResponseBody
	GetSnapshotJob() *SubmitSnapshotJobResponseBodySnapshotJob
}

type SubmitSnapshotJobResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-5EB0-4AF6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshot job.
	SnapshotJob *SubmitSnapshotJobResponseBodySnapshotJob `json:"SnapshotJob,omitempty" xml:"SnapshotJob,omitempty" type:"Struct"`
}

func (s SubmitSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSnapshotJobResponseBody) GetSnapshotJob() *SubmitSnapshotJobResponseBodySnapshotJob {
	return s.SnapshotJob
}

func (s *SubmitSnapshotJobResponseBody) SetRequestId(v string) *SubmitSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBody) SetSnapshotJob(v *SubmitSnapshotJobResponseBodySnapshotJob) *SubmitSnapshotJobResponseBody {
	s.SnapshotJob = v
	return s
}

func (s *SubmitSnapshotJobResponseBody) Validate() error {
	if s.SnapshotJob != nil {
		if err := s.SnapshotJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitSnapshotJobResponseBodySnapshotJob struct {
	// The ID of the snapshot job.
	//
	// example:
	//
	// ad90a501b1b94b72374ad0050464****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetJobId(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.JobId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) Validate() error {
	return dara.Validate(s)
}
