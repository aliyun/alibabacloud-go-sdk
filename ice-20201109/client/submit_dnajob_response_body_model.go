// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitDNAJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitDNAJobResponseBody
	GetRequestId() *string
}

type SubmitDNAJobResponseBody struct {
	// The ID of the media fingerprint analysis job. We recommend that you save this ID for subsequent calls of other operations.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDNAJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDNAJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDNAJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDNAJobResponseBody) SetJobId(v string) *SubmitDNAJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDNAJobResponseBody) SetRequestId(v string) *SubmitDNAJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDNAJobResponseBody) Validate() error {
	return dara.Validate(s)
}
