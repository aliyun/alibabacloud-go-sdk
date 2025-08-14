// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitSnapshotJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitSnapshotJobResponseBody
	GetRequestId() *string
}

type SubmitSnapshotJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSnapshotJobResponseBody) SetJobId(v string) *SubmitSnapshotJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBody) SetRequestId(v string) *SubmitSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBody) Validate() error {
	return dara.Validate(s)
}
