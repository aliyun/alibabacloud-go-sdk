// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFailoverTestJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateFailoverTestJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateFailoverTestJobResponseBody
	GetRequestId() *string
}

type CreateFailoverTestJobResponseBody struct {
	// The ID of the failover test.
	//
	// example:
	//
	// ftj-xxxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C44F62BE-9CE7-4277-B117-69243F3988BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFailoverTestJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFailoverTestJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFailoverTestJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateFailoverTestJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFailoverTestJobResponseBody) SetJobId(v string) *CreateFailoverTestJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateFailoverTestJobResponseBody) SetRequestId(v string) *CreateFailoverTestJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFailoverTestJobResponseBody) Validate() error {
	return dara.Validate(s)
}
