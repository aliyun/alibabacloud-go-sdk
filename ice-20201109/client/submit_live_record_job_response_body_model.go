// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveRecordJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitLiveRecordJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitLiveRecordJobResponseBody
	GetRequestId() *string
}

type SubmitLiveRecordJobResponseBody struct {
	// The ID of the recording job.
	//
	// example:
	//
	// ab0e3e76-1e9d-11ed-ba64-0c42a1b73d66
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA98A0C-7870-15FE-B96F-8880BB600A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitLiveRecordJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveRecordJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveRecordJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitLiveRecordJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLiveRecordJobResponseBody) SetJobId(v string) *SubmitLiveRecordJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitLiveRecordJobResponseBody) SetRequestId(v string) *SubmitLiveRecordJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLiveRecordJobResponseBody) Validate() error {
	return dara.Validate(s)
}
