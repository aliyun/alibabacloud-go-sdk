// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaDNADeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitMediaDNADeleteJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitMediaDNADeleteJobResponseBody
	GetRequestId() *string
}

type SubmitMediaDNADeleteJobResponseBody struct {
	// The ID of the job.
	//
	// example:
	//
	// 6805B2EC-CE87-****-8FF6-9C0E97719A26
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// e5b1a2e7bee******b632c2710b9423f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaDNADeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaDNADeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitMediaDNADeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitMediaDNADeleteJobResponseBody) SetJobId(v string) *SubmitMediaDNADeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobResponseBody) SetRequestId(v string) *SubmitMediaDNADeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
