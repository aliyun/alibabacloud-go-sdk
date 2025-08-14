// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDNAJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelDNAJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CancelDNAJobResponseBody
	GetRequestId() *string
}

type CancelDNAJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelDNAJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDNAJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDNAJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CancelDNAJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDNAJobResponseBody) SetJobId(v string) *CancelDNAJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CancelDNAJobResponseBody) SetRequestId(v string) *CancelDNAJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDNAJobResponseBody) Validate() error {
	return dara.Validate(s)
}
