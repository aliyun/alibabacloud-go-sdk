// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CancelJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CancelJobResponseBody
	GetRequestId() *string
}

type CancelJobResponseBody struct {
	// The ID of the job.
	//
	// example:
	//
	// d1ce4d3efcb549419193f50f1fcd****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 338CA33A-AE83-5DF4-B6F2-C6D3ED8143F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CancelJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelJobResponseBody) SetJobId(v string) *CancelJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CancelJobResponseBody) SetRequestId(v string) *CancelJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelJobResponseBody) Validate() error {
	return dara.Validate(s)
}
