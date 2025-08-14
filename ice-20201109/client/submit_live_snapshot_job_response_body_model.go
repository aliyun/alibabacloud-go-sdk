// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitLiveSnapshotJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitLiveSnapshotJobResponseBody
	GetRequestId() *string
}

type SubmitLiveSnapshotJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitLiveSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitLiveSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLiveSnapshotJobResponseBody) SetJobId(v string) *SubmitLiveSnapshotJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitLiveSnapshotJobResponseBody) SetRequestId(v string) *SubmitLiveSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLiveSnapshotJobResponseBody) Validate() error {
	return dara.Validate(s)
}
